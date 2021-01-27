package main

import (
	"flag"
	"fmt"
	"github.com/cloudkucooland/go-eiscp"
	"strconv"
)

func main() {
	var command, value string
	host := flag.String("h", "192.168.1.152", "Onkyo host")
	// verbose := flag.Bool("v", false, "verbose")
	flag.Parse()

	args := flag.Args()
	argc := len(args)
	if argc == 0 {
		command = "unset"
	}
	if argc >= 1 {
		command = args[0]
	}
	if argc > 1 {
		value = args[1]
	}

	dev, err := eiscp.NewReceiver(*host, false)
	if err != nil {
		panic(err)
	}
	defer dev.Close()
	if value == "" {
		switch command {
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
			fmt.Println(src)
		case "network":
			ns, err := dev.GetNetworkStatus()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("network status: %+v\n", ns)
			nps, err := dev.GetNetworkPlayStatus()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("network play status: %+v\n", nps)
			resp, err := dev.GetNetworkInfo()
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
			fmt.Printf("network title info: %+v\n", nlt)
			list, err := dev.GetNetworkListInfo()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("list: %+v\n", list)
			nms, err := dev.GetNetworkMenuStatus()
			if err != nil {
				panic(err)
			}
			fmt.Printf("menu status: %+v\n", nms)
			// works with some stations, but not all -- figure out why
			nti, err := dev.GetNetworkTitleName()
			if err != nil {
				panic(err)
			}
			fmt.Printf("title: %s\n", nti)
		case "preset":
			p, _ := dev.GetPreset()
			fmt.Printf("preset: %s\n", p)
		case "temp":
			temp, _ := dev.GetTempData()
			fmt.Printf("temp: %d\n", temp)
		case "nms":
			nms, err := dev.GetNetworkMenuStatus()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%+v\n", nms)
		case "test":
			// nri, _:= dev.GetDetails()
			// fmt.Printf("details : %+v\n", nri)
			src, _ := dev.GetSource()
			fmt.Printf("source: %s\n", src)
			temp, _ := dev.GetTempData()
			fmt.Printf("temp: %d\n", temp)
			fwv, _ := dev.GetFirmwareVersion()
			fmt.Printf("firmware version: %s\n", fwv)
			dm, _ := dev.GetDisplayMode()
			fmt.Printf("display mode: %s\n", dm)
			ai, _ := dev.GetAudioInformation()
			fmt.Printf("audio information: %s\n", ai)
			dim, _ := dev.GetDimmer()
			fmt.Printf("dimmer : %s\n", dim)
			vi, _ := dev.GetVideoInformation()
			fmt.Printf("video information: %s\n", vi)
		case "listeningmode":
			s, err := dev.GetListeningMode()
			if err != nil {
				panic(err)
			}
			fmt.Printf("listening mode: %s\n", s)
		case "listeningmodes":
			for k, v := range eiscp.ListeningModes {
				fmt.Printf("%s: %s\n", k, v)
			}
		case "help":
			fmt.Println("get commands: test, nms, temp, preset, nowplaying, network, source, volume, power, details, listeningmode, listeningmodes")
		default:
			if len(command) != 3 {
				fmt.Println("usage: onkyo [command|CMD] [value]")
				return
			}
			mm, err := dev.SetGetAll(command, "QSTN")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			for _, v := range mm.Messages {
				fmt.Printf("reply: [%s] %s\n", v.Command, v.Response)
			}
		case "unset":
			fmt.Println("usage: onkyo [command] [value]")
		}
	} else {
		switch command {
		case "power":
			v, err := strconv.ParseBool(value)
			if err != nil {
				panic(err)
			}
			fmt.Println(dev.SetPower(v))
		case "volume":
			v, err := strconv.ParseInt(value, 10, 8)
			if err != nil {
				panic(err)
			}
			vol, err := dev.SetVolume(uint8(v))
			if err != nil {
				panic(err)
			}
			fmt.Printf("new volume: %d\n", vol)
		case "source":
			src, ok := eiscp.SourceByName[value]
			if !ok {
				panic("Unknown source")
			}
			fmt.Println(dev.SetSource(src))
		case "preset":
			msg, err := dev.SetPreset(value)
			if err != nil {
				panic(err)
			}
			fmt.Println(msg)
		case "netpreset": // hangs
			msg, err := dev.SetNetworkPreset(value)
			if err != nil {
				panic(err)
			}
			fmt.Println(msg)
		case "netsrc": // hangs
			err := dev.SetNetworkService(value)
			if err != nil {
				panic(err)
			}
			fmt.Println("set")
		case "nja": // turn on/off the network art -- saves bandwidth in my config
			s, err := strconv.ParseBool(value)
			if err != nil {
				panic(err)
			}
			state, err := dev.SetNetworkJacketArt(s)
			if err != nil {
				panic(err)
			}
			fmt.Println(state)
		case "select":
			i, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			err = dev.SelectNetworkListItem(i)
			if err != nil {
				panic(err)
			}
			fmt.Printf("selected: I%05d\n", i)
		case "listeningmode":
			s, err := dev.SetListeningMode(value)
			if err != nil {
				panic(err)
			}
			fmt.Printf("listening mode: %s\n", s)
		case "help":
			fmt.Println("set commands: select, nja, netsrc, netpreset, source, volume, power")
		default:
			mm, err := dev.SetGetAll(command, value)
			if err != nil {
				panic(err)
			}
			for _, v := range mm.Messages {
				fmt.Printf("reply: [%s] %s\n", v.Command, v.Response)
			}
		}
	}
}
