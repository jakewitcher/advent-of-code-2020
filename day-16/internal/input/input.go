package input

import (
	"io/ioutil"
	"log"
	"strings"
)

func Extract(path string) (fields []string, yourTicket string, nearbyTickets []string, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("failed to read file")
		return nil, "", nil, err
	}

	sections := strings.Split(string(file), "\r\n\r\n")
	subsections := make([][]string, len(sections))
	for i, section := range sections {
		subsections[i] = strings.Split(section, "\r\n")
	}

	fields = subsections[0]
	yourTicket = subsections[1][1]
	nearbyTickets = subsections[2][1:]

	return fields, yourTicket, nearbyTickets, nil
}
