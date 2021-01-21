package main

import (
	"flag"
	"fmt"
	"github.com/cloudkucooland/go-eiscp"
	// "strconv"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	host := flag.String("host", "192.168.1.152", "Onkyo host")
	flag.Parse()
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	dev, err := eiscp.NewReceiver(*host, true)
	if err != nil {
		panic(err)
	}

	go func() {
		fmt.Println("starting local listener")
		for msg := range dev.Responses {
			fmt.Printf("local: %+v\n", msg)
		}
	}()

	time.Sleep(time.Second * 3)
	p, err := dev.GetPower()
	fmt.Printf("power: %t\n", p)
	deets, err := dev.GetDetails()
	fmt.Println(deets)

	// fmt.Println(<-sigs)
}
