package program

import "day-14/internal/program"

var parseMaskTestCases = []struct {
	input    string
	expected []program.Bit
}{
	{
		input: "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		expected: []program.Bit{
			{Value: 1, Shift: 6},
			{Value: 0, Shift: 1},
		},
	},
	{
		input: "mask = 000XX11XXX010101101100111110X110X0XX",
		expected: []program.Bit{
			{Value: 0, Shift: 35},
			{Value: 0, Shift: 34},
			{Value: 0, Shift: 33},
			{Value: 1, Shift: 30},
			{Value: 1, Shift: 29},
			{Value: 0, Shift: 25},
			{Value: 1, Shift: 24},
			{Value: 0, Shift: 23},
			{Value: 1, Shift: 22},
			{Value: 0, Shift: 21},
			{Value: 1, Shift: 20},
			{Value: 1, Shift: 19},
			{Value: 0, Shift: 18},
			{Value: 1, Shift: 17},
			{Value: 1, Shift: 16},
			{Value: 0, Shift: 15},
			{Value: 0, Shift: 14},
			{Value: 1, Shift: 13},
			{Value: 1, Shift: 12},
			{Value: 1, Shift: 11},
			{Value: 1, Shift: 10},
			{Value: 1, Shift: 9},
			{Value: 0, Shift: 8},
			{Value: 1, Shift: 6},
			{Value: 1, Shift: 5},
			{Value: 0, Shift: 4},
			{Value: 0, Shift: 2},
		},
	},
}

var parseMemWriteTestCases = []struct {
	input    string
	expected program.MemWrite
}{
	{
		input: "mem[8] = 11",
		expected: program.MemWrite{
			Value: 11,
			Index: 8,
		},
	},
	{
		input: "mem[63416] = 3606",
		expected: program.MemWrite{
			Value: 3606,
			Index: 63416,
		},
	},
}

var decoderApplyTestCases = []struct {
	input    int
	expected int
	mask     program.Mask
}{
	{
		input:    11,
		expected: 73,
		mask: program.Mask{
			Bits: []program.Bit{
				{Value: 1, Shift: 6},
				{Value: 0, Shift: 1},
			},
		},
	},
	{
		input:    101,
		expected: 101,
		mask: program.Mask{
			Bits: []program.Bit{
				{Value: 1, Shift: 6},
				{Value: 0, Shift: 1},
			},
		},
	},
	{
		input:    0,
		expected: 64,
		mask: program.Mask{
			Bits: []program.Bit{
				{Value: 1, Shift: 6},
				{Value: 0, Shift: 1},
			},
		},
	},
}

var runInitializationVersionTwoTestCases = []struct {
	mask     program.Mask
	memWrite program.MemWrite
	expected program.Memory
}{
	{
		mask: program.Mask{
			Bits: []program.Bit{
				{Value: 0, Shift: 35},
				{Value: 0, Shift: 34},
				{Value: 0, Shift: 33},
				{Value: 0, Shift: 32},
				{Value: 0, Shift: 31},
				{Value: 0, Shift: 30},
				{Value: 0, Shift: 29},
				{Value: 0, Shift: 28},
				{Value: 0, Shift: 27},
				{Value: 0, Shift: 26},
				{Value: 0, Shift: 25},
				{Value: 0, Shift: 24},
				{Value: 0, Shift: 23},
				{Value: 0, Shift: 22},
				{Value: 0, Shift: 21},
				{Value: 0, Shift: 20},
				{Value: 0, Shift: 19},
				{Value: 0, Shift: 18},
				{Value: 0, Shift: 17},
				{Value: 0, Shift: 16},
				{Value: 0, Shift: 15},
				{Value: 0, Shift: 14},
				{Value: 0, Shift: 13},
				{Value: 0, Shift: 12},
				{Value: 0, Shift: 11},
				{Value: 0, Shift: 10},
				{Value: 0, Shift: 9},
				{Value: 0, Shift: 8},
				{Value: 0, Shift: 7},
				{Value: 0, Shift: 6},
				{Value: 1, Shift: 4},
				{Value: 1, Shift: 3},
				{Value: 0, Shift: 2},
				{Value: 0, Shift: 1},
			},
		},
		memWrite: program.MemWrite{
			Value: 100,
			Index: 42,
		},
		expected: map[int]int{
			26: 100,
			27: 100,
			58: 100,
			59: 100,
		},
	},
	{
		mask: program.Mask{
			Bits: []program.Bit{
				{Value: 0, Shift: 35},
				{Value: 0, Shift: 34},
				{Value: 0, Shift: 33},
				{Value: 0, Shift: 32},
				{Value: 0, Shift: 31},
				{Value: 0, Shift: 30},
				{Value: 0, Shift: 29},
				{Value: 0, Shift: 28},
				{Value: 0, Shift: 27},
				{Value: 0, Shift: 26},
				{Value: 0, Shift: 25},
				{Value: 0, Shift: 24},
				{Value: 0, Shift: 23},
				{Value: 0, Shift: 22},
				{Value: 0, Shift: 21},
				{Value: 0, Shift: 20},
				{Value: 0, Shift: 19},
				{Value: 0, Shift: 18},
				{Value: 0, Shift: 17},
				{Value: 0, Shift: 16},
				{Value: 0, Shift: 15},
				{Value: 0, Shift: 14},
				{Value: 0, Shift: 13},
				{Value: 0, Shift: 12},
				{Value: 0, Shift: 11},
				{Value: 0, Shift: 10},
				{Value: 0, Shift: 9},
				{Value: 0, Shift: 8},
				{Value: 0, Shift: 7},
				{Value: 0, Shift: 6},
				{Value: 0, Shift: 5},
				{Value: 1, Shift: 4},
				{Value: 0, Shift: 2},
			},
		},
		memWrite: program.MemWrite{
			Value: 1,
			Index: 26,
		},
		expected: map[int]int{
			16: 1,
			17: 1,
			18: 1,
			19: 1,
			24: 1,
			25: 1,
			26: 1,
			27: 1,
		},
	},
}
