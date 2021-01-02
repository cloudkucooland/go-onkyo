package eiscp

import (
	"encoding/hex"
	"strconv"
	"strings"
)

// SetSource - Set Onkyo source channel
func (d *Device) SetSource(source Source) (*Message, error) {
	return d.Set("SLI", string(source))
}

// GetSource - Get Onkyo source channel. Use SourceToName to get readable name
func (d *Device) GetSource() (Source, error) {
	msg, err := d.Set("SLI", "QSTN")
	if err != nil {
		return "", err
	}
	return Source(msg.Response), err
}

// SetPower - turn on/off Onkyo device
func (d *Device) SetPower(on bool) (*Message, error) {
	if on {
		return d.Set("PWR", "01")
	}
	return d.Set("PWR", "00")
}

// GetPower - get Onkyo power state
func (d *Device) GetPower() (bool, error) {
	msg, err := d.Set("PWR", "QSTN")
	if err != nil {
		return false, err
	}
	return msg.Response == "01", err
}

// SetVolume - set master volume in Onkyo receiver
func (d *Device) SetVolume(level uint8) (*Message, error) {
	return d.Set("MVL", strings.ToUpper(hex.EncodeToString([]byte{level})))
}

// GetVolume - get master volume in Onkyo receiver
func (d *Device) GetVolume() (uint8, error) {
	msg, err := d.Set("MVL", "QSTN")
	if err != nil {
		return 0, err
	}
	vol, err := strconv.ParseUint(msg.Response, 16, 8)
	return uint8(vol), err
}
