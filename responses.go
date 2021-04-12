package eiscp

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// bundle these together, whenever we see an NLT, start over.
type menu struct {
	NLS []*NLS
	NLT *NLT
}

var Menu menu

func (r *Message) parseResponseValue() (interface{}, error) {
	switch r.Command {
	case "SLI":
		s, ok := SourceToName[Source(r.Response)]
		if !ok {
			s = "unknown"
		}
		return s, nil
	case "PWR":
		return r.Response == "01", nil
	case "MVL":
		vol, err := strconv.ParseUint(r.Response, 16, 8)
		if err != nil {
			return 0, err
		}
		return uint8(vol), nil
	case "AMT":
		return r.Response == "01", nil
	case "NRI":
		var nri NRI
		if err := xml.Unmarshal([]byte(r.Response), &nri); err != nil {
			ologger.Println(string(r.Response))
			return nil, err
		}
		return &nri, nil
	case "RES":
		res, ok := resolutions[r.Response]
		if !ok {
			res = "unknown"
		}
		return res, nil
	case "VWM":
		mode, ok := vwm[r.Response]
		if !ok {
			mode = "unknown"
		}
		return mode, nil
	case "LMD":
		mode, ok := ListeningModes[r.Response]
		if !ok {
			mode = r.Response
		}
		return mode, nil
	case "NJA":
		// this is useless for reading the cover art if sent, only for setting it to on/off
		if r.Response == "00" {
			return false, nil
		}
		return true, nil
	case "NLT":
		var nlt NLT
		nlt.ServiceType = NetSource(r.Response[0:2])
		nlt.UIType = r.Response[2:3]
		nlt.LayerType = r.Response[3:4]
		nlt.CurrentPos = r.Response[4:8]
		nlt.NumItems = r.Response[8:12]
		nlt.NumLayers = r.Response[12:14]
		nlt.IconL = NetSource(r.Response[16:18])
		nlt.IconR = NetSource(r.Response[18:20])
		nlt.Status = r.Response[20:22]
		nlt.Title = r.Response[22:len(r.Response)]

		// reset the menu
		Menu.NLS = nil
		Menu.NLT = &nlt
		return &nlt, nil
	case "NLS":
		var nls NLS
		nls.InfoType = r.Response[0:1]
		nls.LineInfo = r.Response[1:2]
		nls.Property = r.Response[2:3]
		nls.Line = r.Response[3:len(r.Response)]
		Menu.NLS = append(Menu.NLS, &nls)
		return &nls, nil
	case "TPD":
		// "F100C 38"
		sub := r.Response[6:8]
		if sub == "" || sub == " 0" {
			return uint8(38), nil
		}

		tempC, err := strconv.Atoi(sub)
		if err != nil {
			return uint8(38), err
		}
		return uint8(tempC), nil
	case "PRS":
		return r.Response, nil
	case "NDS":
		return parseNDS(r.Response)
	case "NST":
		return parseNST(r.Response)
	case "NMS":
		return parseNMS(r.Response)
	case "MOT":
		mot := false
		if r.Response == "01" {
			mot = true
		}
		return mot, nil
	case "RAS":
		ras := false
		if r.Response == "01" {
			ras = true
		}
		return ras, nil
	case "PCT":
		pct := false
		if r.Response == "01" {
			pct = true
		}
		return pct, nil
	case "DIM":
		switch r.Response {
		case "00":
			return "Bright", nil
		case "01":
			return "Medium", nil
		case "02":
			return "Dim", nil
		case "03":
			return "Off", nil
		case "08":
			return "Bright & LED-Off", nil
		default:
			return "unknown", nil
		}
	default:
		return r.Response, nil
	}
	// not reached
	return nil, nil
}

var DimmerState = map[string]string{
	"Bright":           "00",
	"Medium":           "01",
	"Dim":              "02",
	"Off":              "03",
	"Bright & LED-Off": "08",
	"unknown":          "ff",
}

func parseNDS(r string) (*NetworkStatus, error) {
	var ns NetworkStatus
	switch r[0:1] {
	case "-":
		ns.Source = "No Connection"
	case "E":
		ns.Source = "Ethernet"
	case "W":
		ns.Source = "Wireless"
	default:
		ns.Source = "Unknown"
	}

	switch r[1:2] {
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

	switch r[2:3] {
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

func parseNST(r string) (*NetworkPlayStatus, error) {
	var nps NetworkPlayStatus
	switch r[0:1] {
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

	switch r[1:2] {
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

	switch r[2:3] {
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

func parseNMS(r string) (*NetworkMenuStatus, error) {
	// Mxxxxx20e
	var nms NetworkMenuStatus
	if r[0:1] == "M" {
		nms.Menu = true
	}
	if r[1:3] == "F1" {
		nms.PositiveButtonIcon = true
	}
	if r[3:5] == "F2" {
		nms.NegativeButtonIcon = true
	}
	if r[5:6] == "S" {
		nms.SeekTime = true
	}
	switch r[6:7] {
	case "1":
		nms.ElapsedTimeMode = 1
	case "2":
		nms.ElapsedTimeMode = 2
	default:
		nms.ElapsedTimeMode = 0
	}
	nms.Service = r[7:]
	nms.ServiceName = NetSourceToName[NetSource(strings.ToUpper(nms.Service))]

	return &nms, nil
}
