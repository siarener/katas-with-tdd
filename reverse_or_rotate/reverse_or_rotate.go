package kata

import (
	"strings"
)

func chunkString(input string, size int) []string {
	endOfString := len(input)
	var chunks []string
	for i := 0; i < endOfString; i += size {
		endPosition := i + size
		if endPosition > endOfString {
			endPosition = endOfString
		}
		chunks = append(chunks, input[i:endPosition])
	}
	return chunks
}

func rotateStringToLeft(input string) string {
	return input[1:] + input[:1]
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sumOfCubes(input string) int {
	var sumOfCubes int
	for _, i := range input {
		// get int representation of digit from an ASCII value
		number := int(i - '0')
		sumOfCubes += number * number * number
	}
	return sumOfCubes
}

func isSumOfCubesOfDigitsDivisibleByTwo(input string) bool {
	return sumOfCubes(input)%2 == 0
}

func Revrot(s string, n int) string {
	var revisedChunks []string
	chunks := chunkString(s, n)

	for _, chunk := range chunks {
		if len(chunk) == n {
			if !isSumOfCubesOfDigitsDivisibleByTwo(chunk) {
				revisedChunks = append(revisedChunks, rotateStringToLeft(chunk))
			} else {
				revisedChunks = append(revisedChunks, reverseString(chunk))
			}
		}
	}
	return strings.Join(revisedChunks, "")
}

// revrot("123456987654", 6) --> "234561876549"
