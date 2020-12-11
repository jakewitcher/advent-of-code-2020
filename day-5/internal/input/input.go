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

	input := strings.Split(string(file), "\r\n")
	return input, nil
}
