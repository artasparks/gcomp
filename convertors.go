package gcomp

import (
	"fmt"
	"strconv"
)

// Converts a string to a float and panics if the conversion fails.
func Stoi(s string) int64 {
	iv, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Error parsing float: %v", err))
	}
	return iv
}

func Stof(s string) float64 {
	ft, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Sprintf("Error parsing float: %v", err))
	}
	return ft
}
