package input

import (
	"io/ioutil"
	"log"
	"strings"
)

func Extract() ([][]string, error) {
	file, err := ioutil.ReadFile("internal/input/input.txt")
	if err != nil {
		log.Println("failed to read file")
		return nil, err
	}

	input := strings.Split(string(file), "\r\n\r\n")
	groups := make([][]string, len(input))
	for i, group := range input {
		groups[i] = strings.Split(group, "\r\n")
	}
	return groups, nil
}
