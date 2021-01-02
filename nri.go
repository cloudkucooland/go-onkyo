package eiscp

import (
	"encoding/xml"
)

type NRI struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Status  string   `xml:"status,attr"`
	Device  struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"id,attr"`
		Brand            string `xml:"brand"`
		Category         string `xml:"category"`
		Year             string `xml:"year"`
		Model            string `xml:"model"`
		Destination      string `xml:"destination"`
		Productid        string `xml:"productid"`
		Deviceserial     string `xml:"deviceserial"`
		Macaddress       string `xml:"macaddress"`
		Modeliconurl     string `xml:"modeliconurl"`
		Friendlyname     string `xml:"friendlyname"`
		Firmwareversion  string `xml:"firmwareversion"`
		Ecosystemversion string `xml:"ecosystemversion"`
		Netservicelist   struct {
			Text       string `xml:",chardata"`
			Count      string `xml:"count,attr"`
			Netservice []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Value    string `xml:"value,attr"`
				Name     string `xml:"name,attr"`
				Account  string `xml:"account,attr"`
				Password string `xml:"password,attr"`
				Zone     string `xml:"zone,attr"`
				Enable   string `xml:"enable,attr"`
				Addqueue string `xml:"addqueue,attr"`
				Sort     string `xml:"sort,attr"`
			} `xml:"netservice"`
		} `xml:"netservicelist"`
		Zonelist struct {
			Text  string `xml:",chardata"`
			Count string `xml:"count,attr"`
			Zone  []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Value    string `xml:"value,attr"`
				Name     string `xml:"name,attr"`
				Volmax   string `xml:"volmax,attr"`
				Volstep  string `xml:"volstep,attr"`
				Src      string `xml:"src,attr"`
				Dst      string `xml:"dst,attr"`
				Lrselect string `xml:"lrselect,attr"`
			} `xml:"zone"`
		} `xml:"zonelist"`
		Selectorlist struct {
			Text     string `xml:",chardata"`
			Count    string `xml:"count,attr"`
			Selector []struct {
				Text   string `xml:",chardata"`
				ID     string `xml:"id,attr"`
				Value  string `xml:"value,attr"`
				Name   string `xml:"name,attr"`
				Zone   string `xml:"zone,attr"`
				Iconid string `xml:"iconid,attr"`
			} `xml:"selector"`
		} `xml:"selectorlist"`
		Presetlist struct {
			Text   string `xml:",chardata"`
			Count  string `xml:"count,attr"`
			Preset []struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Band string `xml:"band,attr"`
				Freq string `xml:"freq,attr"`
				Name string `xml:"name,attr"`
			} `xml:"preset"`
		} `xml:"presetlist"`
		Controllist struct {
			Text    string `xml:",chardata"`
			Count   string `xml:"count,attr"`
			Control []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Value    string `xml:"value,attr"`
				Zone     string `xml:"zone,attr"`
				Min      string `xml:"min,attr"`
				Max      string `xml:"max,attr"`
				Step     string `xml:"step,attr"`
				Code     string `xml:"code,attr"`
				Position string `xml:"position,attr"`
			} `xml:"control"`
		} `xml:"controllist"`
		Functionlist struct {
			Text     string `xml:",chardata"`
			Count    string `xml:"count,attr"`
			Function []struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Value string `xml:"value,attr"`
			} `xml:"function"`
		} `xml:"functionlist"`
		Tuners struct {
			Text  string `xml:",chardata"`
			Count string `xml:"count,attr"`
			Tuner []struct {
				Text string `xml:",chardata"`
				Band string `xml:"band,attr"`
				Min  string `xml:"min,attr"`
				Max  string `xml:"max,attr"`
				Step string `xml:"step,attr"`
			} `xml:"tuner"`
		} `xml:"tuners"`
	} `xml:"device"`
}
