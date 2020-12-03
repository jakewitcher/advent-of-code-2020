package input

import (
	"io/ioutil"
	"log"
	"strings"
)

func Extract() ([][]rune, error) {
	file, err := ioutil.ReadFile("internal/input/input.txt")
	if err != nil {
		log.Println("failed to read file")
		return nil, err
	}

	lines := strings.Split(string(file), "\r\n")
	matrix := make([][]rune, len(lines))

	for i, line := range lines {
		runes := make([]rune, len(line))
		for j, r := range line {
			runes[j] = r
		}
		matrix[i] = runes
	}

	return matrix, nil
}
