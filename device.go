package eiscp

import (
	"fmt"
	"io"
	"net"
	"strings"
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
	Host            string
	persistent      bool
	conn            net.Conn
	destinationType DeviceType
	version         byte
	sender          chan Command
	Responses       chan Message
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
		d.Responses = make(chan Message)

		go func(dev Device) {
			fmt.Println("starting persistent listener")
			d.persistentListener()
		}(d)
	}

	return &d, nil
}

// NewReceiver - sugar for NewDevice with Receiver as device type and version 1
// host must be an IPv4 dotted-quad address... for now
func NewReceiver(host string, persistent bool) (*Device, error) {
	return newDevice(host, TypeReceiver, 0x01, persistent)
}

// Close connection
func (d *Device) Close() error {
	if d.persistent {
		fmt.Println("ignoring close on persistent channel")
		return nil
	}
	if d.conn != nil {
		err := d.conn.Close()
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

	conn, err := net.DialTCP("tcp4", nil, &r)
	d.conn = conn
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (d *Device) readSimple() (*Message, error) {
	if d.conn == nil {
		return nil, fmt.Errorf("not connected")
	}

	blocksize := 1024
	bufsiz := 20 * blocksize // NRI needs 9k on my NR-686
	raw := make([]byte, 0, bufsiz)
	tmp := make([]byte, blocksize)
	for {
		d.conn.SetReadDeadline(time.Now().Add(time.Second * 3))
		defer d.conn.SetDeadline(time.Time{})

		n, err := d.conn.Read(tmp)
		if err != nil && err != io.EOF {
			fmt.Printf("cannot read data from device: %s", err.Error())
			return nil, err
		}

		// fmt.Printf("read %d bytes: %s\n", n, string(tmp))
		raw = append(raw, tmp[:n]...)

		// saw EOF or short block, must be done
		if err == io.EOF || n != blocksize {
			break
		}
		// NRI needs this... *facepalm*
		time.Sleep(time.Millisecond * 10)
	}

	var msg Message
	msg.Parse(&raw)
	if !msg.Valid {
		return nil, fmt.Errorf("invalid EISCP message")
	}
	return &msg, nil
}

func (d *Device) readMulti(command string) (*MultiMessage, error) {
	if d.conn == nil {
		return nil, fmt.Errorf("not connected")
	}

	mm := MultiMessage{}
	blocksize := 1500 // should read interface for MTU
	bufsize := 20 * blocksize
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

func (d *Device) persistentListener() error {
	if d.conn == nil {
		return fmt.Errorf("not connected")
	}

	blocksize := 1500
	block := make([]byte, blocksize)
	bufsize := 5 * blocksize
	buf := make([]byte, bufsize)
	var bytesread int
	var msg Message
	for {
		n, err := d.conn.Read(block)
		if err != nil {
			return err
		}

		// copy the read block into the buffer
		for pos, b := range block[:n] {
			buf[bytesread+pos] = b
			// buf[pos] = b
		}
		// if it wasn't a full block, it must be complete, parse and send to channel
		if n != blocksize {
			bytesread = 0
			msg.Parse(&buf)
			if !msg.Valid {
				fmt.Printf("invalid message: %+v", msg)
			}
			d.Responses <- msg
			time.Sleep(time.Millisecond * 5)
			continue
		}
		// otherwise keep reading
		bytesread += n
	}
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

// SetSingle sends a command which will only ever have a single-message response, multiple values are ignored.
// It is better to use SetGetOne in most cases
func (d *Device) setSingle(command, arg string) (*Message, error) {
	err := d.writeCommand(command, arg)
	if err != nil {
		return nil, err
	}

	pulls := 0
	msg, err := d.readSimple()
	if err != nil {
		return nil, err
	}

	// the response given didn't answer the question we asked -- keep digging
	for msg.Command != command {
		// NLS and NLT are common up whenever something is playing
		if msg.Command != "NLT" && msg.Command != "NLS" && msg.Command != "NTM" {
			fmt.Printf("wrong response: [%s] / [%s] %s\n", command, msg.Command, msg.Response)
			pulls++
		}
		if pulls > 5 {
			return nil, fmt.Errorf("too many responses")
		}

		// try again
		msg, err = d.readSimple()
		if err != nil {
			return nil, err
		}
	}
	return msg, nil
}

// SetOnly sends a command and does not check for a response
func (d *Device) SetOnly(command, arg string) error {
	err := d.writeCommand(command, arg)
	if err != nil {
		return err
	}
	return nil
}

// Set sends a command and returns all responses
func (d *Device) SetGetAll(command, arg string) (*MultiMessage, error) {
	err := d.writeCommand(command, arg)
	if err != nil {
		return nil, err
	}

	mm, err := d.readMulti(command)
	if err != nil {
		return nil, err
	}
	return mm, nil
}

func (d *Device) SetGetOne(command, arg string) (*Message, error) {
	mm, err := d.SetGetAll(command, arg)
	if err != nil {
		return nil, err
	}
	if len(mm.Messages) == 0 {
		return nil, fmt.Errorf("no reply")
	}
	return mm.Messages[len(mm.Messages)-1], nil
}
