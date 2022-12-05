package internal

import "strconv"

// SliceOfStringToInt transform slice of string into slice of int
// this code ignore errors
func SliceOfStringToInt(s []string) (intSlice []int) {
	for _, v := range s {
		intValue, _ := strconv.Atoi(v)
		intSlice = append(intSlice, intValue)
	}
	return
}

// SumSliceOfInts sum all items in slice of items
func SumSliceOfInts(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}
