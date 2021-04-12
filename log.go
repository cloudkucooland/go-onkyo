package eiscp

import (
	"log"
)

var ologger onkyologger = log.Default()

type onkyologger interface {
	Println(...interface{})
	Printf(string, ...interface{})
}

func SetLogger(l onkyologger) {
	ologger = l
}
