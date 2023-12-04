package Solvers

import (
	"errors"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day4 struct{}

func (d Day4) ChallengeOne(input IStringScanner) string {
	var result int64

	for input.Scan() {
		cardPoints, err := d.ParseCard(input.Text())

		if err == nil {
			result += cardPoints
		}
	}

	return strconv.FormatInt(result, 10)
}

func (d Day4) ChallengeTwo(input IStringScanner) string {
	return ""
}

func (d Day4) ParseCard(input string) (int64, error) {
	var matches int64

	card := strings.Split(input, "|")
	if len(card) != 2 {
		return -1, errors.New("input split resulted in to many sections")
	}

	winningNumbers := strings.Split(card[0], ":")

	actualNumbers := strings.Split(winningNumbers[1], " ")
	cardNumbers := strings.Split(card[1], " ")

	for _, val := range cardNumbers {
		if val != "" {
			if slices.Contains(actualNumbers, val) {
				matches++
			}
		}
	}

	return int64(math.Pow(2, float64(matches-1))), nil
}
