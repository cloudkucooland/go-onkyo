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
	SrcDVD                  = "10"
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
}

// SourceToName - map source enum to channel name
var SourceToName = map[Source]string{
	SrcVCR:           "vcr",
	SrcCBL:           "cbl",
	SrcGame:          "game",
	SrcAux1:          "aux1",
	SrcAux2:          "aux2",
	SrcPC:            "pc",
	SrcDVD:           "dvd",
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
}
