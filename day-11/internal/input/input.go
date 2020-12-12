package input

import (
	"io/ioutil"
	"log"
	"strings"
)

func Extract(path string) ([][]rune, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("failed to read file")
		return nil, err
	}

	lines := strings.Split(string(file), "\r\n")
	input := make([][]rune, len(lines))
	for i, line := range lines {
		input[i] = []rune(line)
	}

	return input, nil
}
