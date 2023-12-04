package Solvers

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Day1 struct{}

func (d Day1) GetCalibrationValueFromString(input string) string {
	var sb strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) {
			sb.WriteRune(char)
			break
		}
	}

	for _, char := range Reverse(input) {
		if unicode.IsDigit(char) {
			sb.WriteRune(char)
			break
		}
	}

	return sb.String()
}

func (d Day1) GetCalibrationValueChallenge2FromString(input string) string {
	writtenDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	// Come up with a way to use a sliding widow to check for a written version of a digit, or the actual digit we run across

	replaced := input

	for target, newData := range writtenDigits {
		replaced = strings.ReplaceAll(replaced, target, newData)
	}

	return d.GetCalibrationValueFromString(replaced)
}

func (d Day1) ChallengeOne(input IStringScanner) string {
	var total int64

	for input.Scan() {
		calValue := d.GetCalibrationValueFromString(input.Text())

		intVal, _ := strconv.ParseInt(calValue, 10, 64)
		total += intVal
	}

	return strconv.FormatInt(total, 10)
}

func (d Day1) ChallengeTwo(input IStringScanner) string {
	return ""
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
