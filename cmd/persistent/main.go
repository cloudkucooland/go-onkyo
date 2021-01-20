package main

import (
	"flag"
	"fmt"
	"github.com/cloudkucooland/go-eiscp"
	// "strconv"
)

func main() {
	hostP := flag.String("host", "", "Onkyo host")
	flag.Parse()

	host := *hostP
	if host == "" {
		host = "192.168.1.152"
	}

	dev, err := eiscp.NewReceiver(host, true)
	if err != nil {
		panic(err)
	}

	for msg := range dev.Responses {
		fmt.Printf("%+v\n", msg)
	}
}
