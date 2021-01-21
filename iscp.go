package eiscp

import (
	"bytes"
	"encoding/binary"
	// "fmt"
)

// Message eISCP
type Message struct {
	Version     byte   // always the same
	Destination byte   // always the same
	headerSize  uint32 // always the same
	dataSize    uint32 // does this need to be here now? once parsed it is never touched again
	raw         []byte // used for sending
	Command     string // verify you've got the right Command
	Response    string // the response value
	Valid       bool   // if the packet was able to be parsed
}

type MultiMessage struct {
	Messages []*Message
}

// Parse raw message from network into an eISCP message
func (msg *Message) Parse(rawP *[]byte) {
	raw := *rawP
	if string(raw[:4]) != "ISCP" {
		// return fmt.Errorf("this is not an EISCP message: %s", string(*rawP))
		msg.Valid = false
		return
	}
	msg.headerSize = binary.BigEndian.Uint32(raw[4:8])
	if msg.headerSize != 16 {
		// return fmt.Errorf("invalid header size")
		msg.Valid = false
		return
	}

	msg.dataSize = binary.BigEndian.Uint32(raw[8:12])
	msg.Version = raw[12]
	if msg.Version != 1 {
		msg.Valid = false
		return
	}

	msg.Command = string(raw[18:21])
	msg.Response = string(raw[21 : 16+msg.dataSize-3])
	msg.Valid = true
}

// BuildISCP - Build ISCP message
func (msg *Message) BuildISCP() []byte {
	buffer := bytes.Buffer{}
	buffer.WriteRune('!')             // Start character
	buffer.WriteByte(msg.Destination) // Receiver
	buffer.Write(msg.raw)
	buffer.Write([]byte{0x0D})
	return buffer.Bytes()
}

// BuildEISCP - Build ISCP message into ethernet frame
func (msg *Message) BuildEISCP() []byte {
	iscp := msg.BuildISCP()
	sizebuf := make([]byte, 4)
	buffer := bytes.Buffer{}
	buffer.WriteString("ISCP")
	buffer.Write([]byte{0, 0, 0, 0x10}) // Header size

	binary.BigEndian.PutUint32(sizebuf, uint32(len(iscp)))
	buffer.Write(sizebuf)         // Data size
	buffer.WriteByte(msg.Version) // Version
	buffer.Write([]byte{0, 0, 0}) // Reserved

	buffer.Write(iscp) //Data
	return buffer.Bytes()
}
