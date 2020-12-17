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
	cubes := make([][]rune, len(lines))
	for i, line := range lines {
		cubes[i] = make([]rune, len(line))
		for j, r := range line {
			cubes[i][j] = r
		}
	}

	return cubes, err
}
