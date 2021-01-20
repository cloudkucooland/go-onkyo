package eiscp

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	Code  string
	Value string
}

// SetSource - Set Onkyo source channel by friendly name
func (d *Device) SetSource(source Source) (*Message, error) {
	return d.SetGetOne("SLI", string(source))
}

// SetSourceByCode - Set Onkyo source channel by code
func (d *Device) SetSourceByCode(code int) (*Message, error) {
	return d.SetGetOne("SLI", fmt.Sprintf("%02X", code))
}

// GetSource - Get Onkyo source channel. Use SourceToName to get readable name
func (d *Device) GetSource() (Source, error) {
	msg, err := d.SetGetOne("SLI", "QSTN")
	if err != nil {
		return "", err
	}
	return Source(msg.Response), err
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
	return msg.Response == "01", err
}

// SetVolume - set master volume in Onkyo receiver
func (d *Device) SetVolume(level uint8) (uint8, error) {
	msg, err := d.SetGetOne("MVL", strings.ToUpper(hex.EncodeToString([]byte{level})))
	if err != nil {
		return uint8(0), err
	}
	vol, err := strconv.ParseUint(msg.Response, 16, 8)
	return uint8(vol), err
}

// GetVolume - get master volume in Onkyo receiver
func (d *Device) GetVolume() (uint8, error) {
	msg, err := d.SetGetOne("MVL", "QSTN")
	if err != nil {
		return 0, err
	}
	vol, err := strconv.ParseUint(msg.Response, 16, 8)
	return uint8(vol), err
}

func (d *Device) GetMute() (bool, error) {
	msg, err := d.SetGetOne("AMT", "QSTN")
	if err != nil {
		return false, err
	}
	return msg.Response == "01", err
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
	return msg.Response == "01", err
}

func (d *Device) GetDetails() (*NRI, error) {
	msg, err := d.SetGetOne("NRI", "QSTN")
	if err != nil {
		return nil, err
	}
	var nri NRI
	if err := xml.Unmarshal([]byte(msg.Response), &nri); err != nil {
		return nil, err
	}
	return &nri, nil
}

func (d *Device) GetDisplayMode() (string, error) {
	msg, err := d.SetGetOne("DIF", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetAudioInformation() (string, error) {
	msg, err := d.SetGetOne("IFA", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetDimmer() (string, error) {
	msg, err := d.SetGetOne("DIM", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetVideoInformation() (string, error) {
	msg, err := d.SetGetOne("IFV", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

// hangs
func (d *Device) GetFLInformation() (string, error) {
	msg, err := d.SetGetOne("FLD", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetMonitorResolution() (string, error) {
	msg, err := d.SetGetOne("RES", "QSTN")
	if err != nil {
		return "unknown", err
	}
	res, ok := resolutions[msg.Response]
	if !ok {
		res = "unknown"
	}
	return res, nil
}

// hangs
func (d *Device) GetHDMIOut() (string, error) {
	msg, err := d.SetGetOne("HOI", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

// hangs
func (d *Device) GetISF() (string, error) {
	msg, err := d.SetGetOne("ISF", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

// hangs
func (d *Device) GetWideVideoMode() (string, error) {
	msg, err := d.SetGetOne("VWM", "QSTN")
	if err != nil {
		return "", err
	}
	mode, ok := vwm[msg.Response]
	if !ok {
		mode = "unknown"
	}
	return mode, nil
}

func (d *Device) GetListeningMode() (string, error) {
	msg, err := d.SetGetOne("LMD", "QSTN")
	if err != nil {
		return "", err
	}
	mode, ok := listeningmodes[msg.Response]
	if !ok {
		mode = msg.Response
	}
	return mode, nil
}

func (d *Device) GetNetworkJacketArt() (string, error) {
	msg, err := d.SetGetOne("NJA", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) SetNetworkJacketArt(s bool) (string, error) {
	state := "DIS"
	if s {
		state = "ENA"
	}
	err := d.SetOnly("NJA", state)
	if err != nil {
		return "", err
	}

	msg, err := d.SetGetOne("NJA", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

type NLT struct {
	ServiceType NetSource
	UIType      string // 1 - int // 0 : List, 1 : Menu, 2 : Playback, 3 : Popup, 4 : Keyboard, 5 : Menu
	LayerType   string // 1 - int // 0 : NET TOP, 1 : Service Top,DLNA/USB/iPod Top, 2 : under 2nd Layer
	CurrentPos  string // 4 - hex
	NumItems    string // 4 - hex
	NumLayers   string // 2 - hex
	Reserved    string // 2 - unused
	IconL       NetSource
	IconR       NetSource
	Status      string // 2 -- hex -- lookup table // 00 : None, 01 : Connecting, 02 : Acquiring License, 03 : Buffering 04 : Cannot Play, 05 : Searching, 06 : Profile update, 07 : Operation disabled 08 : Server Start-up, 09 : Song rated as Favorite, 0A : Song banned from station, 0B : Authentication Failed, 0C : Spotify Paused(max 1 device), 0D : Track Not Available, 0E : Cannot Skip
	Title       string // the rest
}

func (d *Device) GetNetworkTitle() (*NLT, error) {
	msg, err := d.SetGetOne("NLT", "QSTN")
	if err != nil {
		return nil, err
	}
	var nlt NLT
	nlt.ServiceType = NetSource(msg.Response[0:2])
	nlt.UIType = msg.Response[2:3]
	nlt.LayerType = msg.Response[3:4]
	nlt.CurrentPos = msg.Response[4:8]
	nlt.NumItems = msg.Response[8:12]
	nlt.NumLayers = msg.Response[12:14]
	nlt.IconL = NetSource(msg.Response[16:18])
	nlt.IconR = NetSource(msg.Response[18:20])
	nlt.Status = msg.Response[20:22]
	nlt.Title = msg.Response[22:len(msg.Response)]
	return &nlt, nil
}

func (d *Device) GetNetworkTitleName() (string, error) {
	msg, err := d.SetGetOne("NTI", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

type NLS struct {
	InfoType string // (A : ASCII letter, C : Cursor Info, U : Unicode letter)
	LineInfo string // (0-9 : 1st to 10th Line)
	Property string // varies based on context
	Line     string
}

func (d *Device) GetNetworkListInfo() (*NLS, error) {
	msg, err := d.SetGetOne("NLS", "QSTN")
	if err != nil {
		return nil, err
	}
	var nls NLS
	nls.InfoType = msg.Response[0:1]
	nls.LineInfo = msg.Response[1:2]
	nls.Property = msg.Response[2:3]
	nls.Line = msg.Response[3:len(msg.Response)]
	return &nls, nil
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
	return msg.Response, nil
}

func (d *Device) GetTempData() (string, error) {
	msg, err := d.SetGetOne("TPD", "QSTN")
	if err != nil {
		return "", err
	}
	vals := strings.Split(msg.Response, " ")
	if len(vals) < 3 {
		return "", fmt.Errorf("did not get temp response")
	}
	return vals[2], nil
}

// AM/FM tuner preset
func (d *Device) GetPreset() (string, error) {
	msg, err := d.SetGetOne("PRS", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

// AM/FM tuner preset
func (d *Device) SetPreset(p string) (string, error) {
	msg, err := d.SetGetOne("PRS", p)
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) SetNetworkPreset(p string) (string, error) {
	// msg, err := d.SetGetOne("NPZ", p)
	msg, err := d.SetGetOne("NPR", p)
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

type NetworkStatus struct {
	Source string
	Front  string
	Rear   string
}

func (d *Device) GetNetworkStatus() (*NetworkStatus, error) {
	msg, err := d.SetGetOne("NDS", "QSTN")
	if err != nil {
		return nil, err
	}
	var ns NetworkStatus
	switch msg.Response[0:1] {
	case "-":
		ns.Source = "No Connection"
	case "E":
		ns.Source = "Ethernet"
	case "W":
		ns.Source = "Wireless"
	default:
		ns.Source = "Unknown"
	}

	switch msg.Response[1:2] {
	case "-":
		ns.Front = "No Device"
	case "i":
		ns.Front = "iPod"
	case "M":
		ns.Front = "Memory/NAS"
	case "W":
		ns.Front = "Wireless Adaptor"
	case "B":
		ns.Front = "Bluetooth Adaptor"
	case "x":
		ns.Front = "Disabled"
	default:
		ns.Front = "Unknown"
	}

	switch msg.Response[2:3] {
	case "-":
		ns.Rear = "no device"
	case "i":
		ns.Rear = "iPod"
	case "M":
		ns.Rear = "Memory/NAS"
	case "W":
		ns.Rear = "Wireless Adaptor"
	case "B":
		ns.Rear = "Bluetooth Adaptor"
	case "x":
		ns.Rear = "Disabled"
	default:
		ns.Rear = "Unknown"
	}

	return &ns, nil
}

type NetworkPlayStatus struct {
	State   string
	Repeat  string
	Shuffle string
}

func (d *Device) GetNetworkPlayStatus() (*NetworkPlayStatus, error) {
	msg, err := d.SetGetOne("NST", "QSTN")
	if err != nil {
		return nil, err
	}
	var nps NetworkPlayStatus
	switch msg.Response[0:1] {
	case "S":
		nps.State = "Stop"
	case "P":
		nps.State = "Play"
	case "p":
		nps.State = "Pause"
	case "F":
		nps.State = "Fast-Forward"
	case "R":
		nps.State = "Rewind"
	case "E":
		nps.State = "EOF"
	}

	switch msg.Response[1:2] {
	case "-":
		nps.Repeat = "Off"
	case "R":
		nps.Repeat = "All"
	case "F":
		nps.Repeat = "Folder"
	case "1":
		nps.Repeat = "One"
	case "x":
		nps.Repeat = "Disabled"
	default:
		nps.Repeat = "Unknown"
	}

	switch msg.Response[2:3] {
	case "-":
		nps.Shuffle = "Off"
	case "R":
		nps.Shuffle = "All"
	case "F":
		nps.Shuffle = "Folder"
	case "1":
		nps.Shuffle = "One"
	case "x":
		nps.Shuffle = "Disabled"
	default:
		nps.Shuffle = "Unknown"
	}
	return &nps, nil
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
	return msg.Response, err
}

func (d *Device) SetNetworkServiceTuneIn() error {
	return d.SetNetworkService(NetSrcTuneIn + "0")
}

func (d *Device) SetNetworkService(s string) error {
	err := d.SetOnly("NSV", s) // NSV hangs on reads
	return err
}

/* this is what I'm after, being able to adjust the steam to which I'm listening by favorite number
func (d *Device) SetNetworkFavorite(s string) (string, error) {
	service := fmt.Sprintf("010%s", s)
	msg, err := d.SetGetOne("NSV", service)
	if err != nil {
		return "", err
	}
	return msg.Response, nil
} */

func (d *Device) SelectNetworkListItem(i int) error {
	line := fmt.Sprintf("I%05d", i)
	err := d.SetOnly("NLS", line)
	return err
}

type NetworkMenuStatus struct {
	Menu               bool
	PositiveButtonIcon bool
	NegativeButtonIcon bool
	SeekTime           bool
	ElapsedTimeMode    int
	Service            string
	ServiceName        string
}

func (d *Device) GetNetworkMenuStatus() (*NetworkMenuStatus, error) {
	msg, err := d.SetGetOne("NMS", "QSTN")
	if err != nil {
		return nil, err
	}

	// Mxxxxx20e
	var nms NetworkMenuStatus
	if msg.Response[0:1] == "M" {
		nms.Menu = true
	}
	if msg.Response[1:3] == "F1" {
		nms.PositiveButtonIcon = true
	}
	if msg.Response[3:5] == "F2" {
		nms.NegativeButtonIcon = true
	}
	if msg.Response[5:6] == "S" {
		nms.SeekTime = true
	}
	switch msg.Response[6:7] {
	case "1":
		nms.ElapsedTimeMode = 1
	case "2":
		nms.ElapsedTimeMode = 2
	default:
		nms.ElapsedTimeMode = 0
	}
	nms.Service = msg.Response[7:]
	nms.ServiceName = NetSourceToName[NetSource(strings.ToUpper(nms.Service))]

	return &nms, nil
}
