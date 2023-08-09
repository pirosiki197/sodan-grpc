package util

import "strconv"

// ConvertToInt converts string to int.
// If the number is not valid, return 0.
func ConvertToInt(number string) int {
	value, err := strconv.Atoi(number)
	if err != nil {
		return 0
	}
	return value
}

// ConvertToUint converts string to uint.
// Check the number is empty before calling this function.
func ConvertToUint(number string) uint {
	return uint(ConvertToInt(number))
}
