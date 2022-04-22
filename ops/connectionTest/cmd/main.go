package main

import (
	"bufio"
	"connectionTest/internal/repo"
	"io"
	"log"
	"strings"
)

func main() {
	var res string
	resPath := "conf/result.txt"
	filePath := "conf/ipWithNoRange.txt"
	readFile := repo.NewReader(filePath)
	ipChan := make(chan string, 1000)
	defer close(ipChan)
	// exitChan := make(chan bool, 10)
	// defer close(exitChan)

	file := readFile.ReadFile()
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		ipAndPort, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("read finished!")
				break
			} else {
				log.Println("read file Error: ", err)
			}
		}
		ipAndPort = strings.Replace(ipAndPort, "\t", ":", -1)
		ipAndPort = strings.Replace(ipAndPort, "\r\n", "", -1)
		// fmt.Println("ipAndPort is: ", ipAndPort)

		go repo.Connect(ipAndPort, ipChan)
	}
	for i := 0; i < 5462; i++ {
		str := <-ipChan
		res += str
	}
	// for {
	// 	_, ok := <-exitChan
	// 	if !ok {
	// 		break
	// 	}
	// }
	repo.WriteRes(resPath, res)
}
