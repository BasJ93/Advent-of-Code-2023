package Solvers

import (
	"testing"
)

var challengeOneTests = []struct {
	in  string
	out string
}{
	{"1abc2", "12"},
	{"pqr3stu8vwx", "38"},
	{"a1b2c3d4e5f", "15"},
	{"treb7uchet", "77"},
}

func TestDay1_GetCalibrationValueFromString(t *testing.T) {
	for _, tt := range challengeOneTests {
		t.Run(tt.in, func(t *testing.T) {
			uut := &Day1{}

			result := uut.GetCalibrationValueFromString(tt.in)

			if result != tt.out {
				t.Error()
			}
		})
	}
}

func TestDay1_ChallengeOne(t *testing.T) {
	counter := 0
	lines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	ss := &MockScanner{
		ScanFunc: func() bool {
			if counter < len(lines) {
				counter++
				return true
			}
			return false
		},
		TextFunc: func() string {
			return lines[counter-1]
		},
	}

	uut := &Day1{}

	result := uut.ChallengeOne(ss)

	if result != "142" {
		t.Error()
	}
}

var challengeTwoTests = []struct {
	in  string
	out string
}{
	{"two1nine", "29"},
	{"eightwothree", "83"},
	{"abcone2threexyz", "13"},
	{"xtwone3four", "24"},
	{"4nineeightseven2", "42"},
	{"zoneight234", "14"},
	{"7pqrstsixteen", "76"},
}

func TestDay1_GetCalibrationValueChallenge2FromString(t *testing.T) {
	for _, tt := range challengeTwoTests {
		t.Run(tt.in, func(t *testing.T) {
			uut := &Day1{}

			result := uut.GetCalibrationValueChallenge2FromString(tt.in)

			if result != tt.out {
				t.Error()
			}
		})
	}
}
