package Solvers

import (
	"fmt"
	"reflect"
)

func CreateSolver(day int) ISolver {
	var solver ISolver

	solverName := fmt.Sprintf("Day%d", day)

	typename := reflect.TypeOf(solverName)

	if typename.Kind() != reflect.String {
		fmt.Println("Invalid day selected")
		return nil
	}

	days := map[string]reflect.Type{
		"Day1": reflect.TypeOf(Day1{}),
		"Day2": reflect.TypeOf(Day2{}),
		"Day3": reflect.TypeOf(Day3{}),
		"Day4": reflect.TypeOf(Day4{}),
		/*"Day5": reflect.TypeOf(Day5{}),
		"Day6": reflect.TypeOf(Day6{}),
		"Day7": reflect.TypeOf(Day7{}),*/
	}

	if st, ok := days[solverName]; ok {
		instance := reflect.New(st).Elem()

		solver = instance.Interface().(ISolver)
	} else {
		fmt.Println("Unknown solver type")
		return nil
	}

	return solver
}
