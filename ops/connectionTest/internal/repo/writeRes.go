package repo

import (
	"bufio"
	"log"
	"os"
)

func WriteRes(filePath string, res string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Println("open file failed, ", err)
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(res)
	writer.Flush()
	file.Close()
}
