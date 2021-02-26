package eiscp

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type Command struct {
	Code  string
	Value string
}

// info here on how to control TV/DVD over cec
// CDV & CTV commands, lots of args
// https://github.com/ouija/onkyo-eiscp/blob/master/eiscp/commands.py

// SetSource - Set Onkyo source channel by friendly name
func (d *Device) SetSource(source Source) (*Message, error) {
	return d.SetGetOne("SLI", string(source))
}

// SetSourceByCode - Set Onkyo source channel by code
func (d *Device) SetSourceByCode(code int) (*Message, error) {
	return d.SetGetOne("SLI", fmt.Sprintf("%02X", code))
}

// GetSource - Get Onkyo source channel. Use SourceToName to get readable name
func (d *Device) GetSource() (string, error) {
	msg, err := d.SetGetOne("SLI", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

// GetSourceByCode - Get Onkyo source channel. Use SourceToName to get readable name
func (d *Device) GetSourceByCode() (Source, error) {
	msg, err := d.SetGetOne("SLI", "QSTN")
	if err != nil {
		return "", err
	}
	return Source(msg.Response), nil
}

// SetPower - turn on/off Onkyo device
func (d *Device) SetPower(on bool) (*Message, error) {
	if on {
		return d.SetGetOne("PWR", "01")
	}
	return d.SetGetOne("PWR", "00")
}

// GetPower - get Onkyo power state
func (d *Device) GetPower() (bool, error) {
	msg, err := d.SetGetOne("PWR", "QSTN")
	if err != nil {
		return false, err
	}
	return msg.Parsed.(bool), err
}

// SetVolume - set master volume in Onkyo receiver
func (d *Device) SetVolume(level uint8) (uint8, error) {
	msg, err := d.SetGetOne("MVL", strings.ToUpper(hex.EncodeToString([]byte{level})))
	if err != nil {
		return uint8(0), err
	}
	return msg.Parsed.(uint8), err
}

// GetVolume - get master volume in Onkyo receiver
func (d *Device) GetVolume() (uint8, error) {
	msg, err := d.SetGetOne("MVL", "QSTN")
	if err != nil {
		return 0, err
	}
	return msg.Parsed.(uint8), err
}

func (d *Device) GetMute() (bool, error) {
	msg, err := d.SetGetOne("AMT", "QSTN")
	if err != nil {
		return false, err
	}
	return msg.Parsed.(bool), err
}

func (d *Device) SetMute(mute bool) (bool, error) {
	state := "00"
	if mute {
		state = "01"
	}
	msg, err := d.SetGetOne("AMT", state)
	if err != nil {
		return false, err
	}
	return msg.Parsed.(bool), err
}

func (d *Device) GetDetails() (*NRI, error) {
	msg, err := d.SetGetOne("NRI", "QSTN")
	if err != nil {
		return nil, err
	}
	return msg.Parsed.(*NRI), nil
}

func (d *Device) GetDisplayMode() (string, error) {
	msg, err := d.SetGetOne("DIF", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

func (d *Device) GetAudioInformation() (string, error) {
	msg, err := d.SetGetOne("IFA", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

func (d *Device) GetDimmer() (string, error) {
	msg, err := d.SetGetOne("DIM", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

func (d *Device) GetVideoInformation() (string, error) {
	msg, err := d.SetGetOne("IFV", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

// hangs
func (d *Device) GetFLInformation() (string, error) {
	msg, err := d.SetGetOne("FLD", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

func (d *Device) GetMonitorResolution() (string, error) {
	msg, err := d.SetGetOne("RES", "QSTN")
	if err != nil {
		return "unknown", err
	}
	return msg.Parsed.(string), nil
}

// hangs
func (d *Device) GetHDMIOut() (string, error) {
	msg, err := d.SetGetOne("HOI", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

// hangs
func (d *Device) GetISF() (string, error) {
	msg, err := d.SetGetOne("ISF", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

// hangs
func (d *Device) GetWideVideoMode() (string, error) {
	msg, err := d.SetGetOne("VWM", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

func (d *Device) GetListeningMode() (string, error) {
	msg, err := d.SetGetOne("LMD", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

func (d *Device) SetListeningMode(code string) (string, error) {
	if len(code) != 2 {
		for k, v := range ListeningModes {
			if v == code {
				code = k
				break
			}
		}
	}
	msg, err := d.SetGetOne("LMD", code)
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

func (d *Device) GetNetworkJacketArt() (string, error) {
	msg, err := d.SetGetOne("NJA", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

// this should return a bool...
func (d *Device) SetNetworkJacketArt(s bool) (bool, error) {
	state := "DIS"
	if s {
		state = "ENA"
	}
	err := d.SetOnly("NJA", state)
	if err != nil {
		return false, err
	}

	msg, err := d.SetGetOne("NJA", "QSTN")
	if err != nil {
		return false, err
	}
	if msg.Parsed == nil {
		return false, nil
	}
	return msg.Parsed.(bool), nil
}

func (d *Device) GetNetworkTitle() (*NLT, error) {
	msg, err := d.SetGetOne("NLT", "QSTN")
	if err != nil {
		return nil, err
	}
	return msg.Parsed.(*NLT), nil
}

func (d *Device) GetNetworkTitleName() (string, error) {
	msg, err := d.SetGetOne("NTI", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

func (d *Device) GetNetworkListInfo() (*NLS, error) {
	msg, err := d.SetGetOne("NLS", "QSTN")
	if err != nil {
		return nil, err
	}
	return msg.Parsed.(*NLS), nil
}

// hangs
func (d *Device) GetNetworkInfo() (string, error) {
	msg, err := d.SetGetOne("NLA", "L000100000000FF") // doesn't hang, but returns junk
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetFirmwareVersion() (string, error) {
	msg, err := d.SetGetOne("FWV", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

func (d *Device) GetTempData() (uint8, error) {
	msg, err := d.SetGetOne("TPD", "QSTN")
	if err != nil {
		return 0, err
	}
	return msg.Parsed.(uint8), err
}

// AM/FM tuner preset
func (d *Device) GetPreset() (string, error) {
	msg, err := d.SetGetOne("PRS", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

// AM/FM tuner preset
func (d *Device) SetPreset(p string) (string, error) {
	msg, err := d.SetGetOne("PRS", p)
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

func (d *Device) SetNetworkPreset(p string) (string, error) {
	// msg, err := d.SetGetOne("NPZ", p)
	msg, err := d.SetGetOne("NPR", p)
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), nil
}

func (d *Device) GetNetworkStatus() (*NetworkStatus, error) {
	msg, err := d.SetGetOne("NDS", "QSTN")
	if err != nil {
		return nil, err
	}
	return msg.Parsed.(*NetworkStatus), nil
}

func (d *Device) GetNetworkPlayStatus() (*NetworkPlayStatus, error) {
	msg, err := d.SetGetOne("NST", "QSTN")
	if err != nil {
		return nil, err
	}
	return msg.Parsed.(*NetworkPlayStatus), nil
}

// prs : e.g. Sxx or Pxx
// p -> Play Status: "S": STOP, "P": Play, "p": Pause, "F": FF, "R": FR, "E": EOF
// r Repeat Status: "-": Off, "R": All, "F": Folder, "1": Repeat 1, "x": disable
// s Shuffle Status: "-": Off, "S": All , "A": Album, "F": Folder, "x": disable
func (d *Device) SetNetworkPlayStatus(s string) (string, error) {
	msg, err := d.SetGetOne("NST", s)
	if err != nil {
		return "", err
	}
	return msg.Parsed.(string), err
}

func (d *Device) SetNetworkServiceTuneIn() error {
	return d.SetNetworkService(NetSrcTuneIn + "0")
}

func (d *Device) SetNetworkService(s string) error {
	err := d.SetOnly("NSV", s) // NSV hangs on reads
	return err
}

func (d *Device) SelectNetworkListItem(i int) error {
	line := fmt.Sprintf("I%05d", i)
	err := d.SetOnly("NLS", line)
	return err
}

func (d *Device) GetNetworkMenuStatus() (*NetworkMenuStatus, error) {
	msg, err := d.SetGetOne("NMS", "QSTN")
	if err != nil {
		return nil, err
	}
	return msg.Parsed.(*NetworkMenuStatus), nil
}
