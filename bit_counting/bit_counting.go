/* https://www.codewars.com/kata/526571aae218b8ee490006f4/train/go

Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.

Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case */

package kata

import (
	"fmt"
	"strconv"
)

func CountBits(num uint) (int, error) {
	bitRepresentation := fmt.Sprintf("%b", num)
	var count int

	for _, digit := range string(bitRepresentation) {
		number, err := strconv.Atoi(string(digit))
		if err != nil {
			return 0, fmt.Errorf("error converting character to integer: %w", err)
		}
		count += int(number)
	}
	return count, nil
}
