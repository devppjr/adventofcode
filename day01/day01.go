package day01

import (
	"sort"
	"strings"

	"github.com/adventofcode/internal"
)

func Part1(input string) int {
	elfSlice := strings.Split(input, "\n\n")
	mostCaloriesInElf := 0
	for _, elf := range elfSlice {
		sliceElf := internal.SliceOfStringToInt(strings.Split(elf, "\n"))
		elfCalories := internal.SumSliceOfInts(sliceElf)
		if elfCalories > mostCaloriesInElf {
			mostCaloriesInElf = elfCalories
		}
	}

	return mostCaloriesInElf
}

func Part2(input string) int {
	elfSlice := strings.Split(input, "\n\n")
	caloriesByElf := []int{}
	for _, elf := range elfSlice {
		sliceElf := internal.SliceOfStringToInt(strings.Split(elf, "\n"))
		caloriesByElf = append(caloriesByElf, internal.SumSliceOfInts(sliceElf))
	}

	sort.Ints(caloriesByElf)

	return internal.SumSliceOfInts(caloriesByElf[len(caloriesByElf)-3:])
}
