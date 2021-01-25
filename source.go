package eiscp

// Source name of input channel
type Source string

// Sample sources
const (
	SrcVCR           Source = "00"
	SrcCBL                  = "01"
	SrcGame                 = "02"
	SrcAux1                 = "03"
	SrcAux2                 = "04"
	SrcPC                   = "05"
	SrcVideo7               = "06"
	SrcDVD                  = "10"
	SrcStrm                 = "11"
	SrcTV                   = "12"
	SrcTape                 = "20"
	SrcTape2                = "21"
	SrcPhono                = "22"
	SrcCD                   = "23"
	SrcFM                   = "24"
	SrcAM                   = "25"
	SrcTuner                = "26"
	SrcDLNA2                = "27"
	SrcInternetRadio        = "28"
	SrcUsbFront             = "29"
	SrcUsbRear              = "2A"
	SrcNetwork              = "2B"
	SrcUSBToggle            = "2C"
	SrcAirplay              = "2D"
	SrcBluetooth            = "2E"
	SrcDAC                  = "2F"
	SrcMultiChan            = "30"
	SrcXM                   = "31"
	SrcSirius               = "32"
	SrcDAB                  = "33"
	SrcUniversal            = "40"
	SrcLine                 = "41"
	SrcLine2                = "42"
	SrcOptical              = "44"
	SrcCoax                 = "45"
	SrcHDMI5                = "55"
	SrcHDMI6                = "56"
	SrcHDMI7                = "57"
)

// SourceByName - map channel name to source enum const
var SourceByName = map[string]Source{
	"vcr":            SrcVCR,
	"cbl":            SrcCBL,
	"game":           SrcGame,
	"aux1":           SrcAux1,
	"aux2":           SrcAux2,
	"pc":             SrcPC,
	"dvd":            SrcDVD,
	"phono":          SrcPhono,
	"cd":             SrcCD,
	"fm":             SrcFM,
	"am":             SrcAM,
	"tuner":          SrcTuner,
	"dlna2":          SrcDLNA2,
	"internet-radio": SrcInternetRadio,
	"usb-front":      SrcUsbFront,
	"usb-rear":       SrcUsbRear,
	"network":        SrcNetwork,
	"video7":         SrcVideo7,
	"strm-box":       SrcStrm,
	"tape":           SrcTape,
	"tape2":          SrcTape2,
	"tv":             SrcTV,
	"bluetooth":      SrcBluetooth,
	"dac":            SrcDAC,
	"airplay":        SrcAirplay,
	"usb-toggle":     SrcUSBToggle,
	"line":           SrcLine,
	"line2":          SrcLine2,
	"universal":      SrcUniversal,
	"optical":        SrcOptical,
	"coax":           SrcCoax,
	"multi-ch":       SrcMultiChan,
	"xm":             SrcXM,
	"sirius":         SrcSirius,
	"dab":            SrcDAB,
	"hdmi-5":         SrcHDMI5,
	"hdmi-6":         SrcHDMI6,
	"hdmi-7":         SrcHDMI7,
}

// SourceToName - map source enum to channel name
// XXX finish this
var SourceToName = map[Source]string{
	SrcVCR:           "vcr",
	SrcCBL:           "cbl",
	SrcGame:          "game",
	SrcAux1:          "aux1",
	SrcAux2:          "aux2",
	SrcPC:            "pc",
	SrcVideo7:        "video7",
	SrcDVD:           "dvd",
	SrcStrm:          "strm-box",
	SrcTV:            "TV",
	SrcTape:          "tape",
	SrcTape2:         "tape1",
	SrcPhono:         "phono",
	SrcCD:            "cd",
	SrcFM:            "fm",
	SrcAM:            "am",
	SrcTuner:         "tuner",
	SrcDLNA2:         "dlna2",
	SrcInternetRadio: "internet-radio",
	SrcUsbFront:      "usb-front",
	SrcUsbRear:       "usb-rear",
	SrcNetwork:       "network",
	SrcUSBToggle:     "usb-toggle",
	SrcAirplay:       "airplay",
	SrcBluetooth:     "bluetooth",
	SrcDAC:           "dac",
	SrcMultiChan:     "multi-ch",
	SrcXM:            "xm",
	SrcSirius:        "sirius",
	SrcDAB:           "dab",
	SrcUniversal:     "universal",
	SrcLine:          "line",
	SrcLine2:         "line2",
	SrcOptical:       "optical",
	SrcCoax:          "coax",
	SrcHDMI5:         "hdmi-5",
	SrcHDMI6:         "hdmi-6",
	SrcHDMI7:         "hdmi-7",
}
