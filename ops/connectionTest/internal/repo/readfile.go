package repo

import (
	"log"
	"os"
)

type ReadFile struct {
	filePath string
}

func NewReader(filePath string) *ReadFile {
	return &ReadFile{
		filePath: filePath,
	}
}

func (r ReadFile) ReadFile() *os.File {
	file, err := os.OpenFile(r.filePath, os.O_RDONLY, 0777)
	if err != nil {
		log.Println("Error: ", err)
	}
	return file
}
