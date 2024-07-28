package main

import (
	"fmt"
	"os"
	"regexp"
)

var digitRegexp = regexp.MustCompile(`\d+`)

func CopyDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	allDigits := digitRegexp.FindAll(b, -1)

	// Combine all matches into a single slice
	result := []byte{}

	for _, match := range allDigits {
		result = append(result, match...)
	}

	return result
}

func main() {
	filename := "../files/copyDigits.txt"
	digits := CopyDigits(filename)
	fmt.Println("Digits:", string(digits))
}
