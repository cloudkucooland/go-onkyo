package eiscp

import (
	"encoding/xml"
)

type NRI struct {
	XMLName xml.Name `xml:"response"`
	Device  struct {
		ID               string `xml:"id,attr"`
		Brand            string `xml:"brand"`
		Category         string `xml:"category"`
		Year             string `xml:"year"`
		Model            string `xml:"model"`
		Destination      string `xml:"destination"`
		ProductID        string `xml:"productid"`
		DeviceSerial     string `xml:"deviceserial"`
		MacAddress       string `xml:"macaddress"`
		ModelIconURL     string `xml:"modeliconurl"`
		FriendlyName     string `xml:"friendlyname"`
		FirmwareVersion  string `xml:"firmwareversion"`
		EcosystemVersion string `xml:"ecosystemversion"`
		NetServiceList   struct {
			NetService []struct {
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
		ZoneList struct {
			Zone []struct {
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
		SelectorList struct {
			Selector []struct {
				ID     string `xml:"id,attr"`
				Value  string `xml:"value,attr"`
				Name   string `xml:"name,attr"`
				Zone   string `xml:"zone,attr"`
				Iconid string `xml:"iconid,attr"`
			} `xml:"selector"`
		} `xml:"selectorlist"`
		PresetList struct {
			Preset []struct {
				ID   string `xml:"id,attr"`
				Band string `xml:"band,attr"`
				Freq string `xml:"freq,attr"`
				Name string `xml:"name,attr"`
			} `xml:"preset"`
		} `xml:"presetlist"`
		ControlList struct {
			Control []struct {
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
		FunctionList struct {
			Function []struct {
				ID    string `xml:"id,attr"`
				Value string `xml:"value,attr"`
			} `xml:"function"`
		} `xml:"functionlist"`
		Tuners struct {
			Tuner []struct {
				Band string `xml:"band,attr"`
				Min  string `xml:"min,attr"`
				Max  string `xml:"max,attr"`
				Step string `xml:"step,attr"`
			} `xml:"tuner"`
		} `xml:"tuners"`
	} `xml:"device"`
}
