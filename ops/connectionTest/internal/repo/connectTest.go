package repo

import (
	"log"
	"net"
	"time"
)

func Connect(address string, res chan string) {
	// fmt.Println("prepared connect ", address)
	// _, err := net.DialTimeout("tcp", address, time.Second*5)
	dialer := &net.Dialer{
		Timeout: time.Millisecond * 300,
		// Deadline:  time.UnixMilli(int64(time.Millisecond * 300)),
		KeepAlive: time.Millisecond,
	}
	log.Println("prepared connect ", address)
	_, err := dialer.Dial("tcp", address)
	if err != nil {
		// log.Println(err)
		result := address + "\t" + "false" + "\n"
		res <- result
	} else {
		result := address + "\t" + "true" + "\n"
		res <- result
	}
}
