package input

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Extract(path string) ([]int, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("failed to read file")
		return nil, err
	}

	lines := strings.Split(string(file), "\r\n")
	input := make([]int, len(lines))

	for i, line := range lines {
		n, err := strconv.Atoi(line)

		if err != nil {
			log.Printf("failed to parse %v as int\n", line)
			return nil, err
		}

		input[i] = n
	}

	return input, nil
}
