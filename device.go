package eiscp

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

// DeviceType - device destination code in ISCP
type DeviceType byte

// Destination code
const (
	TypeReceiver DeviceType = 0x31
)

// Device of Onkyo receiver
type Device struct {
	Host             string
	persistent       bool
	conn             *net.TCPConn
	destinationType  DeviceType
	version          byte
	sender           chan Command
	privateResponses chan Message
	Responses        chan Message
	mux              sync.Mutex
}

// just use the NewReceiver shortcut
func newDevice(host string, deviceType DeviceType, iscpVersion byte, persistent bool) (*Device, error) {
	d := Device{
		Host:            host,
		destinationType: deviceType,
		version:         iscpVersion,
		persistent:      persistent,
	}

	err := d.Connect()
	if err != nil {
		return nil, err
	}

	if persistent {
		d.sender = make(chan Command)
		d.Responses = make(chan Message, 50)        // the channel for the application to listen on
		d.privateResponses = make(chan Message, 50) // the channel for SetGetAll/SetGet to use

		go func(dev Device) {
			d.persistentListener()
		}(d)

		go func(dev Device) {
			d.persistentSender()
		}(d)
	}

	return &d, nil
}

// NewReceiver - sugar for NewDevice with Receiver as device type and version 1
func NewReceiver(host string, persistent bool) (*Device, error) {
	return newDevice(host, TypeReceiver, 0x01, persistent)
}

// Close connection
func (d *Device) Close() error {
	if d.conn != nil {
		// d.conn.SetLinger(0)
		err := d.conn.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		d.conn = nil
		return err
	}
	return nil
}

// Connect to an eISCP device by v4 IP address (not host name)
func (d *Device) Connect() error {
	if d.conn != nil {
		fmt.Println("already connected")
		return nil
	}

	// now that I can move data, switch this back to net.Dial to be more flexible
	r := net.TCPAddr{
		IP:   net.ParseIP(d.Host),
		Port: 60128,
	}

	conn, err := net.DialTCP("tcp", nil, &r)
	d.conn = conn
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// read is used for non-persistent connections (e.g. onkyo cli tool)
func (d *Device) read(command string) (*MultiMessage, error) {
	if d.conn == nil {
		return nil, fmt.Errorf("not connected")
	}

	mm := MultiMessage{}
	blocksize := 1024 // should read interface for MTU
	bufsize := 100 * blocksize
	for {
		raw := make([]byte, 0, bufsize)
		tmp := make([]byte, blocksize)
		for {
			d.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
			defer d.conn.SetDeadline(time.Time{})

			n, err := d.conn.Read(tmp)
			if err != nil && err != io.EOF && !strings.Contains(err.Error(), "i/o timeout") {
				fmt.Printf("cannot read data from device: %s", err.Error())
				return nil, err
			} else if err != nil && strings.Contains(err.Error(), "i/o timeout") {
				return &mm, nil
			}
			raw = append(raw, tmp[:n]...)
			if err == io.EOF || n != blocksize {
				break
			}

			// let NRI and others catch up
			time.Sleep(time.Millisecond * 10)
		}
		var msg Message
		msg.Parse(&raw)
		if !msg.Valid {
			return &mm, nil
		}
		// fmt.Printf("got message [%s]: [%s]\n", msg.Command, msg.Response)
		mm.Messages = append(mm.Messages, &msg)
		if msg.Command == command {
			// fmt.Println("got original command, returning")
			return &mm, nil
		}
	}
}

// long-lived clients should use the persistentListner and read from the d.Responses channel
func (d *Device) persistentListener() error {
	if d.conn == nil {
		return fmt.Errorf("not connected")
	}

	blocksize := 1024 // get interface MTU
	block := make([]byte, blocksize)
	bufsize := 100 * blocksize
	buf := make([]byte, bufsize)
	var bytesread int
	var msg Message
	for {
		n, err := d.conn.Read(block)
		if err != nil {
			fmt.Printf("persistentListener error: [%s], resetting\n", err.Error())
			if err := d.Close(); err != nil {
				fmt.Println(err.Error())
				return err
			}
			if err := d.Connect(); err != nil {
				fmt.Println(err.Error())
				return err
			}
			continue
		}

		// copy the read block into the buffer
		for pos, b := range block[:n] {
			buf[bytesread+pos] = b
		}

		// if it wasn't a full block, it must be complete, parse and send to channel
		if n != blocksize {
			bytesread = 0
			msg.Parse(&buf)
			if !msg.Valid {
				fmt.Printf("invalid message: %+v\n", msg)
			}
			d.Responses <- msg
			d.privateResponses <- msg
			continue
		}
		// otherwise keep reading
		bytesread += n
		if bytesread == bufsize {
			fmt.Printf("buffer full, %n:\n%s\n", bytesread, buf)
			bytesread = 0 // trash everything read so far
			continue      // the remainder of this read will be garbage
		}
		// NRI needs this on TX-NR686 (lowest threshold not researched)
		time.Sleep(time.Millisecond * 10)
	}
	return nil
}

func (d *Device) persistentSender() error {
	if d.conn == nil {
		return fmt.Errorf("not connected")
	}

	for cmd := range d.sender {
		d.send(cmd)
	}
	return nil
}

func (d *Device) writeCommand(command, arg string) error {
	c := Command{
		Code:  command,
		Value: arg,
	}
	if d.persistent {
		d.sender <- c
		return nil
	}
	return d.send(c)
}

func (d *Device) send(c Command) error {
	if d.conn == nil {
		// be smart and try to reconnect if possible
		return fmt.Errorf("not connected")
	}

	msg := Message{
		Destination: byte(d.destinationType),
		Version:     d.version,
		raw:         []byte(c.Code + c.Value),
	}
	m := msg.BuildEISCP()
	// fmt.Printf("m: %+v %s\n", m, string(m))
	_, err := d.conn.Write(m)

	return err
}

// SetOnly sends a command and does not check for a response
func (d *Device) SetOnly(command, arg string) error {
	d.mux.Lock()
	defer d.mux.Unlock()

	err := d.writeCommand(command, arg)
	if err != nil {
		return err
	}
	return nil
}

// Set sends a command and returns all responses
func (d *Device) SetGetAll(command, arg string) (*MultiMessage, error) {
	d.mux.Lock()
	defer d.mux.Unlock()

	err := d.writeCommand(command, arg)
	if err != nil {
		return nil, err
	}

	if d.persistent {
		var pmm MultiMessage
		for {
			select {
			case msg := <-d.privateResponses:
				// fmt.Printf("SetGetAll: %+v\n", msg)
				pmm.Messages = append(pmm.Messages, &msg)
				if msg.Command == command {
					return &pmm, nil
				}
			case <-time.After(time.Second * 3):
				return &pmm, fmt.Errorf("timeout reached: %d non-responses received", len(pmm.Messages))
			}
		}
	}

	mm, err := d.read(command)
	if err != nil {
		return nil, err
	}
	return mm, nil
}

func (d *Device) SetGetOne(command, arg string) (*Message, error) {
	// SetGetAll does the requred locking
	mm, err := d.SetGetAll(command, arg)
	if err != nil {
		return nil, err
	}
	if len(mm.Messages) == 0 {
		return nil, fmt.Errorf("no response")
	}
	return mm.Messages[len(mm.Messages)-1], nil
}
