package eiscp

import (
	"context"
	"strings"
	"time"

	"github.com/brutella/dnssd"
)

func Discover() (string, error) {
	discovered := ""
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	found := func(e dnssd.BrowseEntry) {
		if strings.Contains(e.Name, "@Onkyo") {
			// look through the list of IPs, pick something IPv4, IPV6 doesn't seem to work
			for _, ipa := range e.IPs {
				if ipa.To4() != nil {
					discovered = ipa.String()
					cancel()
					return
				}
			}
		}
	}

	if err := dnssd.LookupType(ctx, "_raop._tcp.local.", found, reject); err != nil {
		if err.Error() != "context canceled" {
			ologger.Printf("discovery: %v\n", err)
			return discovered, err
		}
	}
	return discovered, nil
}

func reject(e dnssd.BrowseEntry) {
	ologger.Printf("dnssd-lookup: %+v", e)
}
