package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/adventofcode/day01"
)

func main() {
	// FIXME: make this parts automagic :)
	fmt.Println(day01.Part1(readInput("day01/input.txt")))
	fmt.Println(day01.Part2(readInput("day01/input.txt")))
}

func readInput(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
