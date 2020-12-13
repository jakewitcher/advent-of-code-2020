package input

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Extract(path string) (int, []string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("failed to read file")
		return 0, nil, err
	}

	lines := strings.Split(string(file), "\r\n")
	timestamp, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, nil, err
	}

	buses := strings.Split(lines[1], ",")
	return timestamp, buses, nil
}
