package main

import (
	"Advent_of_Code_2023/Solvers"
	"bytes"
	"fmt"
	"log"
)
import "os"
import "bufio"

func main() {
	//fmt.Println("Hello AoC2023")

	dayNumber := 4

	buf1 := bytes.Buffer{}
	buf2 := bytes.Buffer{}

	file, err := os.Open(fmt.Sprintf("Input/day%d.txt", dayNumber))

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	initialScanner := bufio.NewScanner(file)

	for initialScanner.Scan() {
		line := initialScanner.Text()
		fmt.Fprintln(&buf1, line) // save for later
		fmt.Fprintln(&buf2, line) // save for later
	}

	scanner1 := bufio.NewScanner(&buf1)
	scanner2 := bufio.NewScanner(&buf2)

	day := Solvers.CreateSolver(dayNumber)

	result1 := day.ChallengeOne(scanner1)
	result2 := day.ChallengeTwo(scanner2)

	fmt.Printf("Result 1: %s\n", result1)
	fmt.Printf("Result 2: %s\n", result2)
}
