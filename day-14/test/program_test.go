package program

import (
	"day-14/internal/input"
	"day-14/internal/program"
	"fmt"
	"testing"
)

func TestParseMask(t *testing.T) {
	for _, test := range parseMaskTestCases {
		actual, err := program.ParseMask(test.input)
		if err != nil {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if len(actual.Bits) != len(test.expected) {
			t.Fatalf("expected number of bits %d, actual %d", len(test.expected), len(actual.Bits))
		}

		for i, bit := range actual.Bits {
			if bit != test.expected[i] {
				t.Fatalf("expected %v, actual %v", test.expected[i], bit)
			}
		}
	}
}

func TestParseMemWrite(t *testing.T) {
	for _, test := range parseMemWriteTestCases {
		actual, err := program.ParseMemWrite(test.input)
		if err != nil {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if *actual != test.expected {
			t.Fatalf("expected %v, actual %v", test.expected, *actual)
		}
	}
}

func TestMaskApply(t *testing.T) {
	for _, test := range decoderApplyTestCases {
		decoder := program.VersionOneDecoder{Mask: &test.mask}
		if actual := decoder.Apply(test.input); actual != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func TestRunInitializationVersionTwoDecoder(t *testing.T) {
	for _, test := range runInitializationVersionTwoTestCases {
		decoder := program.VersionTwoDecoder{Mask: &test.mask}
		memory := make(map[int]int)
		decoder.Decode(memory, &test.memWrite)

		if len(memory) != len(test.expected) {
			t.Fatalf("expected map of length %d, actual %d", len(test.expected), len(memory))
		}

		for k, v := range test.expected {
			if n, ok := memory[k]; !ok || n != v {
				var msg string
				if !ok {
					msg = " but key was missing"
				} else {
					msg = fmt.Sprintf(", actual %d", n)
				}
				t.Fatalf("expected key %d with value %d%s", k, v, msg)
			}
		}
	}
}

func BenchmarkRunInitializationVersionOneDecoder(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = program.RunInitialization(n, &program.VersionOneDecoder{})
	}
}

func BenchmarkRunInitializationVersionTwoDecoder(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = program.RunInitialization(n, &program.VersionTwoDecoder{})
	}
}
