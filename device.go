// Package eiscp provides basic support for eISCP/ISCP protocol
package eiscp

import (
    // "io/ioutil"
	"fmt"
	"net"
)

// Device of Onkyo receiver
type Device struct {
    Host            string
	conn            net.Conn
	destinationType DeviceType
	version         byte
}

// NewDevice - create and connect to eISCP device
// just use the NewReceiver shortcut
func NewDevice(host string, deviceType DeviceType, iscpVersion byte) (*Device, error) {
	d := Device{
        Host: host,
        destinationType: deviceType,
        version: iscpVersion,
    }
	err := d.Connect()
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// NewReceiver - sugar for NewDevice with Receiver as device type and version 1
func NewReceiver(host string) (*Device, error) {
	return NewDevice(host, TypeReceiver, 0x01)
}

// Close connection
func (d *Device) Close() error {
	if d.conn != nil {
		return d.conn.Close()
	}
    d.conn = nil
	return nil
}


// Connect to an eISCP device by v4 IP address (not host name)
func (d *Device) Connect() error {
    r := net.TCPAddr{
        IP:   net.ParseIP(d.Host),
        Port: 60128,
    }

    // fmt.Printf("DialTCP: %+v\n", r)
    conn, err := net.DialTCP("tcp4", nil, &r)
    d.conn = conn
	if err != nil {
        fmt.Println(err.Error())
        return err
    }
	return nil
}

// Read, parse, and validate
func (d *Device) readResponse() (*Message, error) {
    raw := make([]byte, 1024)
    n, err := d.conn.Read(raw)
    if err != nil {
        fmt.Printf("Cannot read data from device: %s", err.Error())
        return nil, err
    }
    if n > 1024 {
        fmt.Println("result overran buffer, bailing")
        return nil, err
    }
    // fmt.Printf("raw response:\n%s\n(%d bytes read)\n", string(raw), n)

    var msg Message
    if err := msg.Parse(&raw); err != nil {
		return nil, err
	}
	return &msg, nil
}

// WriteCommand - write command with arg to remote connection
func (d *Device) writeCommand(command, arg string) error {
	if d.conn == nil {
		return fmt.Errorf("Not connected")
	}

	msg := Message{
        Destination: byte(d.destinationType),
        Version: d.version,
        ISCP: []byte(command + arg),
    }
	req := msg.BuildEISCP()
	// fmt.Printf("req: %+v %s\n", req, string(req))
	_, err := d.conn.Write(req)
	return err
}

// Set does a WriteCommand followed by a ReadResponse
func (d *Device) Set(command, arg string) (*Message, error) {
	err := d.writeCommand(command, arg)
	if err != nil {
		return nil, err
	}
    msg, err := d.readResponse()
	if err != nil {
		return nil, err
	}
    return msg, nil
}
