package Solvers

import "testing"

func TestDay2_ChallengeOne(t *testing.T) {
	counter := 0
	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
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

	uut := &Day2{}

	result := uut.ChallengeOne(ss)

	if result != "8" {
		t.Error()
	}
}

func TestDay2_ChallengeTwo(t *testing.T) {
	counter := 0
	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
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

	uut := &Day2{}

	result := uut.ChallengeTwo(ss)

	if result != "2286" {
		t.Error()
	}
}

var day2ParseTest = []struct {
	in    string
	id    int64
	red   int64
	green int64
	blue  int64
}{
	{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 1, 4, 2, 6},
	{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 2, 1, 3, 4},
	{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 3, 20, 13, 6},
	{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 4, 14, 3, 15},
	{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 5, 6, 3, 2},
}

func TestDay2_ParseStringToGameState(t *testing.T) {
	for _, tt := range day2ParseTest {
		t.Run(tt.in, func(t *testing.T) {
			uut := &Day2{}

			result, err := uut.ParseStringToGameState(tt.in)

			if err != nil {
				t.Error(err)
			}

			if result.id != tt.id {
				t.Error("incorrect id")
			}

			if result.red != tt.red {
				t.Errorf("incorrect number of red. got %d expected %d", result.red, tt.red)
			}

			if result.green != tt.green {
				t.Errorf("incorrect number of green. got %d expected %d", result.green, tt.green)
			}

			if result.blue != tt.blue {
				t.Errorf("incorrect number of blue. got %d expected %d", result.blue, tt.blue)
			}
		})
	}
}
