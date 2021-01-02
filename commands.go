package eiscp

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
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
	msg, err := d.Set("MVL", strings.ToUpper(hex.EncodeToString([]byte{level})))
	return msg, err
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

func (d *Device) GetDetails() (*NRI, error) {
	msg, err := d.Set("NRI", "QSTN")
	if err != nil {
		return nil, err
	}
	var nri NRI
	if err := xml.Unmarshal([]byte(msg.Response), &nri); err != nil {
		return nil, err
	}
	fmt.Printf("%+v", nri)

	return &nri, nil
}

func (d *Device) GetDisplayMode() (string, error) {
	msg, err := d.Set("DIF", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetAudioInformation() (string, error) {
	msg, err := d.Set("IFA", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetDimmer() (string, error) {
	msg, err := d.Set("DIM", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetVideoInformation() (string, error) {
	msg, err := d.Set("IFV", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

// hangs
func (d *Device) GetFLInformation() (string, error) {
	msg, err := d.Set("FLD", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

var resolutions = map[string]string{
	"00": "through",
	"01": "auto",
	"02": "480p",
	"03": "720p",
	"04": "1080i",
	"05": "1080p",
	"06": "source",
	"07": "[1080p, 24fs]",
	"08": "4k-upscaling",
	"13": "1680x720p",
	"15": "2560x1080p",
}

// hangs
func (d *Device) GetMonitorResolution() (string, error) {
	msg, err := d.Set("RES", "QSTN")
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
	msg, err := d.Set("HOI", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

// hangs
func (d *Device) GetISF() (string, error) {
	msg, err := d.Set("ISF", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

var vwm = map[string]string{
	"00": "auto",
	"01": "4-3",
	"02": "full",
	"03": "zoom",
	"04": "wide zoom",
	"05": "smart zoom",
	"UP": "up",
}

// hangs
func (d *Device) GetWideVideoMode() (string, error) {
	msg, err := d.Set("VWM", "QSTN")
	if err != nil {
		return "", err
	}
	mode, ok := vwm[msg.Response]
	if !ok {
		mode = "unknown"
	}
	return mode, nil
}

var listeningmodes = map[string]string{
	"00":     "stereo",
	"01":     "direct",
	"02":     "surround",
	"03":     "[film, game-rpg]",
	"04":     "thx",
	"05":     "[action, game-action]",
	"06":     "[musical, game-rock]",
	"07":     "mono-movie",
	"08":     "orchestra",
	"09":     "unplugged",
	"0A":     "studio-mix",
	"0B":     "tv-logic",
	"0C":     "all-ch-stereo",
	"0D":     "theater-dimensional",
	"0E":     "enhanced-7, enhance, game-sports",
	"0F":     "mono",
	"11":     "pure-audio",
	"12":     "multiplex",
	"13":     "full-mono",
	"14":     "[dolby-virtual, surround-enhanced]",
	"15":     "dts-surround-sensation",
	"16":     "audyssey-dsx",
	"1F":     "whole house",
	"23":     "stage",
	"25":     "action",
	"26":     "music",
	"2E":     "sports",
	"40":     "straight-decode",
	"41":     "dolby-ex",
	"42":     "thx-cinema",
	"43":     "thx-surround-ex",
	"44":     "thx-music",
	"45":     "thx-games",
	"50":     "[thx-u2, s1, i, s-cinema, cinema2]",
	"51":     "[thx-musicmode, thx-u2, s2, i, s-music]",
	"52":     "[thx-games, thx-u2, s2, i, s-games]",
	"80":     "[plii, pliix-movie, dolby-atmos, dolby-surround]",
	"81":     "[plii, pliix-music]",
	"82":     "[neo-6-cinema, neo-x-cinema, dts-x, neural-x]",
	"83":     "[neo-6-music, neo-x-music]",
	"84":     "[plii, pliix-thx-cinema, dolby-surround-thx-cinema]",
	"85":     "[neo-6, neo-x-thx-cinema, dts-neural-x-thx-cinema]",
	"86":     "[plii, pliix-game]",
	"87":     "neural-surr",
	"88":     "[neural-thx, nexural-surround]",
	"89":     "[plii, pliix-thx-games, dolby-surround-thx-games]",
	"8A":     "[neo-6, neo-x-thx-games, dts-neural-x-thx-games]",
	"8B":     "[plii, pliix-thx-music, dolby-surround-thx-music]",
	"8C":     "[neo-6, neo-x-thx-music, dts-neural-x-thx-music]",
	"8D":     "neural-thx-cinema",
	"8E":     "neural-thx-music",
	"8F":     "neural-thx-games",
	"90":     "pliiz-height",
	"91":     "neo-6-cinema-dts-surround-sensation",
	"92":     "neo-6-music-dts-surround-sensation",
	"93":     "neural-digital-music",
	"94":     "pliiz-height-thx-cinema",
	"95":     "pliiz-height-thx-music",
	"96":     "pliiz-height-thx-games",
	"97":     "[pliiz-height-thx-u2, s2-cinema]",
	"98":     "[pliiz-height-thx-u2, s2-music]",
	"99":     "[pliiz-height-thx-u2, s2-games]",
	"9A":     "neo-x-game",
	"A0":     "[pliix, plii-movie-audyssey-dsx]",
	"A1":     "[pliix, plii-music-audyssey-dsx]",
	"A2":     "[pliix, plii-game-audyssey-dsx]",
	"A3":     "neo-6-cinema-audyssey-dsx",
	"A4":     "neo-6-music-audyssey-dsx",
	"A5":     "neural-surround-audyssey-dsx",
	"A6":     "neural-digital-music-audyssey-dsx",
	"A7":     "dolby-ex-audyssey-dsx",
	"FF":     "auto-surround",
	"MOVIE":  "movie",
	"MUSIC":  "music",
	"GAME":   "game",
	"THX":    "thx",
	"AUTO":   "auto",
	"SURR":   "surr",
	"STEREO": "stereo",
}

func (d *Device) GetListeningMode() (string, error) {
	msg, err := d.Set("LMD", "QSTN")
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
	msg, err := d.Set("NJA", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) SetNetworkJacketArt(s bool) (string, error) {
	state := "DIS"
	if s == true {
		state = "ENA"
	}
	msg, err := d.Set("NJA", state)
	if err != nil {
		return "", err
	}

	msg, err = d.Set("NJA", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

type NLT struct {
	ServiceType string // 2 - hex -- lookup table // 00 : DLNA, 01 : Favorite, 02 : vTuner, 03 : SiriusXM, 04 : Pandora, 05 : Rhapsody, 06 : Last.fm, 07 : Napster, 08 : Slacker, 09 : Mediafly, 0A : Spotify, 0B : AUPEO!, 0C : radiko, 0D : e-onkyo, 0E : TuneIn Radio, 0F : MP3tunes, 10 : Simfy, 11:Home Media, 12:Deezer, 13:iHeartRadio, F0 : USB Front, F1 : USB Rear, F2 : Internet Radio, F3 : NET, FF : None
	UIType      string // 1 - int // 0 : List, 1 : Menu, 2 : Playback, 3 : Popup, 4 : Keyboard, 5 : Menu
	LayerType   string // 1 - int // 0 : NET TOP, 1 : Service Top,DLNA/USB/iPod Top, 2 : under 2nd Layer
	CurrentPos  string // 4 - hex
	NumItems    string // 4 - hex
	NumLayers   string // 2 - hex
	Reserved    string // 2 - unused
	IconL       string // 2 -- hex -- lookup table // 00 : DLNA, 01 : Favorite, 02 : vTuner, 03 : SiriusXM, 04 : Pandora, 05 : Rhapsody, 06 : Last.fm, 07 : Napster, 08 : Slacker, 09 : Mediafly, 0A : Spotify, 0B : AUPEO!, 0C : radiko, 0D : e-onkyo, 0E : TuneIn Radio, 0F : MP3tunes, 10 : Simfy, 11:Home Media, 12:Deezer, 13:iHeartRadio, F0 : USB Front, F1 : USB Rear, F2 : Internet Radio, F3 : NET, FF : None
	IconR       string // 2 -- hex -- lookup table // 00 : DLNA, 01 : Favorite, 02 : vTuner, 03 : SiriusXM, 04 : Pandora, 05 : Rhapsody, 06 : Last.fm, 07 : Napster, 08 : Slacker, 09 : Mediafly, 0A : Spotify, 0B : AUPEO!, 0C : radiko, 0D : e-onkyo, 0E : TuneIn Radio, 0F : MP3tunes, 10 : Simfy, 11:Home Media, 12:Deezer, 13:iHeartRadio, FF : None
	Status      string // 2 -- hex -- lookup table // 00 : None, 01 : Connecting, 02 : Acquiring License, 03 : Buffering 04 : Cannot Play, 05 : Searching, 06 : Profile update, 07 : Operation disabled 08 : Server Start-up, 09 : Song rated as Favorite, 0A : Song banned from station, 0B : Authentication Failed, 0C : Spotify Paused(max 1 device), 0D : Track Not Available, 0E : Cannot Skip
	Title       string // the rest
}

func (d *Device) GetNetworkTitle() (*NLT, error) {
	msg, err := d.Set("NLT", "QSTN")
	if err != nil {
		return nil, err
	}
	var nlt NLT
	nlt.ServiceType = msg.Response[0:2]
	nlt.UIType = msg.Response[2:3]
	nlt.LayerType = msg.Response[3:4]
	nlt.CurrentPos = msg.Response[4:8]
	nlt.NumItems = msg.Response[8:12]
	nlt.NumLayers = msg.Response[12:14]
	nlt.IconL = msg.Response[16:18]
	nlt.IconR = msg.Response[18:20]
	nlt.Status = msg.Response[20:22]
	nlt.Title = msg.Response[22:len(msg.Response)]
	return &nlt, nil
}

type NLS struct {
	InfoType string // (A : ASCII letter, C : Cursor Info, U : Unicode letter)
	LineInfo string // (0-9 : 1st to 10th Line)
	Property string // varies based on context
	Line     string
}

func (d *Device) GetNetworkListInfo() (*NLS, error) {
	msg, err := d.Set("NLS", "QSTN")
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
	msg, err := d.Set("NLA", "L000100000000FF") // doesn't hang, but returns junk
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetFirmwareVersion() (string, error) {
	msg, err := d.Set("FWV", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetTempData() (string, error) {
	msg, err := d.Set("TPD", "QSTN")
	if err != nil {
		return "", err
	}
	vals := strings.Split(msg.Response, " ")
	return vals[2], nil
}

// AM/FM tuner preset
func (d *Device) GetPreset() (string, error) {
	msg, err := d.Set("PRS", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

// AM/FM tuner preset
func (d *Device) SetPreset(p string) (string, error) {
	msg, err := d.Set("PRS", p)
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) SetNetworkPreset(p string) (string, error) {
	// msg, err := d.Set("NPZ", p)
	msg, err := d.Set("NPR", p)
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetNetworkStatus() (string, error) {
	msg, err := d.Set("NDS", "QSTN")
	// msg, err := d.Set("NLA", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) GetNetworkPlayStatus() (string, error) {
	msg, err := d.Set("NST", "QSTN")
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

func (d *Device) SetNetworkService(s string) (string, error) {
	msg, err := d.Set("NSV", s)
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}

// this is what I'm after, being able to adjust the steam to which I'm listening by favorite number
func (d *Device) SetNetworkFavorite(s string) (string, error) {
	service := fmt.Sprintf("010%s", s)
	msg, err := d.Set("NSV", service)
	if err != nil {
		return "", err
	}
	return msg.Response, nil
}
