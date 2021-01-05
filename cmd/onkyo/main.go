package main

import (
	"flag"
	"fmt"
	"github.com/cloudkucooland/go-eiscp"
	"strconv"
)

func main() {
	hostP := flag.String("host", "", "Onkyo host")
	command := flag.String("command", "", "Param name")
	value := flag.String("value", "", "Param value. Empty means only get")
	listSources := flag.Bool("list-source", false, "List source")
	flag.Parse()

	host := *hostP
	if host == "" {
		host = "192.168.1.152"
	}

	if *listSources {
		for k := range eiscp.SourceByName {
			fmt.Println(k)
		}
		return
	}
	dev, err := eiscp.NewReceiver(host)
	if err != nil {
		panic(err)
	}
	defer dev.Close()
	if *value == "" {
		switch *command {
		case "details":
			nri, err := dev.GetDetails()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("%+v\n", nri)
		case "power":
			resp, err := dev.GetPower()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(resp)
		case "volume":
			resp, err := dev.GetVolume()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("current volume: %d\n", resp)
		case "source":
			src, err := dev.GetSource()
			if err != nil {
				fmt.Println(err)
				return
			}
			if src == "" {
				fmt.Println("unknown")
				return
			}
			fmt.Println(eiscp.SourceToName[src])
		case "network":
			resp, err := dev.GetNetworkStatus()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("network status: %s\n", resp)
			resp, err = dev.GetNetworkPlayStatus()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("network play status: %s\n", resp)
			resp, err = dev.GetNetworkInfo()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("network info: %s\n", resp)
		case "nowplaying":
			nlt, err := dev.GetNetworkTitle()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("now playing: %+v\n", nlt)
			list, err := dev.GetNetworkListInfo()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("list: %+v\n", list)
		case "preset":
			p, _ := dev.GetPreset()
			fmt.Printf("preset: %s\n", p)
		case "temp":
			temp, _ := dev.GetTempData()
			fmt.Printf("temp: %s\n", temp)
		case "nms":
			nms, err := dev.GetNetworkMenuStatus()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%+v\n", nms)
		case "test":
			fwv, _ := dev.GetFirmwareVersion()
			fmt.Printf("firmware version: %s\n", fwv)
			dm, _ := dev.GetDisplayMode() // works
			fmt.Printf("display mode: %s\n", dm)
			ai, _ := dev.GetAudioInformation() // works
			fmt.Printf("audio information: %s\n", ai)
			dim, _ := dev.GetDimmer() // works
			fmt.Printf("dimmer : %s\n", dim)
			vi, _ := dev.GetVideoInformation() // works
			fmt.Printf("video information: %s\n", vi)
			// fl, _ := dev.GetFLInformation() // fails
			// fmt.Printf("fl: %s\n", fl)
			// mr, _ := dev.GetMonitorResolution() // fails
			// fmt.Printf("monitor resolution: %s\n", mr)
			// r, _ := dev.GetHDMIOut() // fails
			// fmt.Printf("hdmi out: %s\n", r)
			// s, _ := dev.GetISF() // fails
			// fmt.Printf("ISF: %s\n", s)
			// s, _ := dev.GetWideVideoMode() // fails
			// fmt.Printf("Wide Video Mode: %s\n", s)
			s, _ := dev.GetListeningMode() // works
			fmt.Printf("Listening Mode: %s\n", s)
		default:
			resp, err := dev.Set(*command, "QSTN")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("reply: %+v\n", resp)
		}
	} else {
		switch *command {
		case "power":
			v, err := strconv.ParseBool(*value)
			if err != nil {
				panic(err)
			}
			fmt.Println(dev.SetPower(v))
		case "volume":
			v, err := strconv.ParseInt(*value, 10, 8)
			if err != nil {
				panic(err)
			}
			fmt.Println(dev.SetVolume(uint8(v)))
		case "source":
			src, ok := eiscp.SourceByName[*value]
			if !ok {
				panic("Unknown source")
			}
			fmt.Println(dev.SetSource(src))
		case "preset":
			msg, err := dev.SetPreset(*value)
			if err != nil {
				panic(err)
			}
			fmt.Println(msg)
		case "netpreset": // hangs
			msg, err := dev.SetNetworkPreset(*value)
			if err != nil {
				panic(err)
			}
			fmt.Println(msg)
		case "netsrc": // hangs
			err := dev.SetNetworkService(*value)
			if err != nil {
				panic(err)
			}
			fmt.Println("set")
		/* case "netfav": // hangs
		msg, err := dev.SetNetworkFavorite(*value)
		if err != nil {
			panic(err)
		}
		fmt.Println(msg) */
		case "nja": // turn on/off the network art -- saves bandwidth in my config
			s, err := strconv.ParseBool(*value)
			if err != nil {
				panic(err)
			}
			state, err := dev.SetNetworkJacketArt(s)
			if err != nil {
				panic(err)
			}
			fmt.Println(state)
		default:
			msg, err := dev.Set(*command, *value)
			if err != nil {
				panic(err)
			}
			fmt.Printf("reply: %+v\n", msg)
		}
	}
}
