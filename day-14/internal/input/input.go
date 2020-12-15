package input

import (
	"io/ioutil"
	"log"
	"strings"
)

func Extract(path string) ([]string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("failed to read file")
		return nil, err
	}

	lines := strings.Split(string(file), "\r\n")
	return lines, nil
}
