package Solvers

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Day2 struct{}

type GameState struct {
	id    int64
	red   int64
	green int64
	blue  int64
}

func (d Day2) ChallengeOne(input IStringScanner) string {
	var maxRed, maxGreen, maxBlue, result int64
	maxRed = 12
	maxGreen = 13
	maxBlue = 14

	for input.Scan() {
		state, err := d.ParseStringToGameState(input.Text())

		if err == nil {
			if state.red <= maxRed && state.green <= maxGreen && state.blue <= maxBlue {
				result += state.id
			}
		}
	}

	return strconv.FormatInt(result, 10)
}

func (d Day2) ChallengeTwo(input IStringScanner) string {
	var result int64

	for input.Scan() {
		state, err := d.ParseStringToGameState(input.Text())

		if err == nil {
			result += state.Power()
		}
	}

	return strconv.FormatInt(result, 10)

}

func (d Day2) ParseStringToGameState(input string) (GameState, error) {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	state := GameState{}

	r, _ := regexp.Compile("(\\d+) red")
	g, _ := regexp.Compile("(\\d+) green")
	b, _ := regexp.Compile("(\\d+) blue")

	split1 := strings.Split(input, ":")
	if len(split1) > 2 {
		return state, errors.New("multiple : detected")
	}

	idSplit := strings.Split(split1[0], " ")
	if len(idSplit) > 2 {
		return state, errors.New("id section contains multiple spaces")
	}

	state.id, _ = strconv.ParseInt(idSplit[1], 10, 64)

	redValues := r.FindAllString(input, -1)
	greenValues := g.FindAllString(input, -1)
	blueValues := b.FindAllString(input, -1)

	for _, s := range redValues {
		valSplit := strings.Split(s, " ")
		if len(valSplit) == 2 {
			tempVal, _ := strconv.ParseInt(valSplit[0], 10, 64)
			if tempVal > state.red {
				state.red = tempVal
			}
		}
	}

	for _, s := range greenValues {
		valSplit := strings.Split(s, " ")
		if len(valSplit) == 2 {
			tempVal, _ := strconv.ParseInt(valSplit[0], 10, 64)
			if tempVal > state.green {
				state.green = tempVal
			}
		}
	}

	for _, s := range blueValues {
		valSplit := strings.Split(s, " ")
		if len(valSplit) == 2 {
			tempVal, _ := strconv.ParseInt(valSplit[0], 10, 64)
			if tempVal > state.blue {
				state.blue = tempVal
			}
		}
	}

	return state, nil
}

func (g GameState) Power() int64 {
	return g.red * g.blue * g.green
}
