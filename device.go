package eiscp

import (
	"fmt"
	"io"
	"net"
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
	conn            net.Conn
	destinationType DeviceType
	version         byte
	cache           map[string]string
}

// just use the NewReceiver shortcut
func newDevice(host string, deviceType DeviceType, iscpVersion byte) (*Device, error) {
	d := Device{
		Host:            host,
		destinationType: deviceType,
		version:         iscpVersion,
	}
	d.cache = make(map[string]string)
	err := d.Connect()
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// NewReceiver - sugar for NewDevice with Receiver as device type and version 1
// host must be an IPv4 dotted-quad address... for now
func NewReceiver(host string) (*Device, error) {
	return newDevice(host, TypeReceiver, 0x01)
}

// Close connection
func (d *Device) Close() error {
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

func (d *Device) readResponse() (*Message, error) {
	if d.conn == nil {
		return nil, fmt.Errorf("not connected")
	}

	blocksize := 1024
	bufsiz := 20 * blocksize // NRI needs 9k on my NR-686
	raw := make([]byte, 0, bufsiz)
	tmp := make([]byte, blocksize)
	for {
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
	if err := msg.Parse(&raw); err != nil {
		return nil, err
	}
	return &msg, nil
}

func (d *Device) writeCommand(command, arg string) error {
	if d.conn == nil {
		// be smart and try to reconnect if possible
		return fmt.Errorf("not connected")
	}

	msg := Message{
		Destination: byte(d.destinationType),
		Version:     d.version,
		ISCP:        []byte(command + arg),
	}
	req := msg.BuildEISCP()
	// fmt.Printf("req: %+v %s\n", req, string(req))
	_, err := d.conn.Write(req)

	return err
}

// Set is the primary interface to send commands to the device
// only use this directly if a specific command is not already written
func (d *Device) Set(command, arg string) (*Message, error) {
	err := d.writeCommand(command, arg)
	if err != nil {
		return nil, err
	}

	msg, err := d.readResponse()
	if err != nil {
		return nil, err
	}
	if msg == nil {
		return nil, fmt.Errorf("unable to read response")
	}

	// the response given didn't answer the question we asked -- keep digging
	for msg.Command != command {
		if err != nil {
			return nil, err
		}
		// store it, in case we want it later
		d.cache[msg.Command] = msg.Response
		// try again
		msg, err = d.readResponse()
	}
	return msg, nil
}

func (d *Device) SetOnly(command, arg string) error {
	err := d.writeCommand(command, arg)
	if err != nil {
		return err
	}
	return nil
}
