package day03

import (
	"strings"
)

var ALPHABET = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
}

func Part1(input string) (totalScore int) {
	rucksacks := strings.Split(strings.TrimSpace(input), "\n")
	for _, rucksack := range rucksacks {
		if len(rucksack)%2 != 0 {
			panic("rucksack should be even")
		}

		startRucksack := rucksack[:len(rucksack)/2]
		endRucksack := rucksack[len(rucksack)/2:]
		repeatedChar := ""
		for repeatedChar == "" {
			for _, char := range startRucksack {
				strChar := string(char)
				if strings.Contains(endRucksack, strChar) {
					repeatedChar = strChar
				}
			}
		}

		priority := ALPHABET[repeatedChar]
		if priority == 0 { // is uppercase. +26
			priority = ALPHABET[strings.ToLower(repeatedChar)] + 26
		}
		totalScore += priority
	}
	return
}
