package program

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func RunInitialization(input []string, decoder Decoder) (int, error) {
	p := Program{
		Decoder: decoder,
		Memory:  make(Memory),
	}

	for _, instruction := range input {
		if strings.HasPrefix(instruction, "mask") {
			mask, err := ParseMask(instruction)
			if err != nil {
				return 0, err
			}

			p.UpdateMask(mask)
			continue
		}

		if strings.HasPrefix(instruction, "mem") {
			memWrite, err := ParseMemWrite(instruction)
			if err != nil {
				return 0, err
			}

			p.Write(memWrite)
		}
	}

	return p.Sum(), nil
}

type MemWrite struct {
	Value int
	Index int
}

type Memory map[int]int

func (m Memory) Sum() int {
	var sum int
	for _, n := range m {
		sum += n
	}

	return sum
}

type Program struct {
	Decoder
	Memory
}

func (p *Program) Write(mw *MemWrite) {
	p.Decode(p.Memory, mw)
}

func ParseMask(input string) (*Mask, error) {
	input = strings.TrimPrefix(input, "mask = ")
	bits := make([]Bit, 0)
	length := len(input)
	for i, r := range input {
		if r == 'X' {
			continue
		}

		value, err := strconv.Atoi(string(r))
		if err != nil {
			return nil, err
		}

		shift := length - (i + 1)
		bit := Bit{Value: value, Shift: shift}
		bits = append(bits, bit)
	}

	return &Mask{Bits: bits}, nil
}

func ParseMemWrite(input string) (*MemWrite, error) {
	regex, err := regexp.Compile("\\d+")
	if err != nil {
		log.Fatal(err.Error())
	}

	nums := regex.FindAll([]byte(input), 2)
	if nums == nil || len(nums) != 2 {
		return nil, fmt.Errorf("invalid memory set string: %s", input)
	}

	index, err := strconv.Atoi(string(nums[0]))
	if err != nil {
		return nil, err
	}

	value, err := strconv.Atoi(string(nums[1]))
	if err != nil {
		return nil, err
	}

	return &MemWrite{Value: value, Index: index}, nil
}
