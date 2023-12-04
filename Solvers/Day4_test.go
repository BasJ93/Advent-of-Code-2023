package Solvers

import "testing"

func TestDay4_ChallengeOne(t *testing.T) {
	counter := 0
	lines := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
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

	uut := &Day4{}

	result := uut.ChallengeOne(ss)

	if result != "13" {
		t.Error()
	}
}

func TestDay4_ChallengeTwo(t *testing.T) {
	counter := 0
	lines := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
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

	uut := &Day4{}

	result := uut.ChallengeTwo(ss)

	if result != "30" {
		t.Error()
	}
}

var day4ParseTest = []struct {
	in  string
	out int64
}{
	{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8},
	{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
	{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
	{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
	{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
	{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
}

func TestDay4_ParseStringToGameState(t *testing.T) {
	for _, tt := range day4ParseTest {
		t.Run(tt.in, func(t *testing.T) {
			uut := &Day4{}

			result, err := uut.ParseCard(tt.in)

			if err != nil {
				t.Error(err)
			}

			if result != tt.out {
				t.Errorf("got %d expected %d", result, tt.out)
			}
		})
	}
}
