// FUNCTION StringToIntSlice(str):

// Create an empty list called result

// FOR each character r in str DO
// Convert r to its integer value
// Append this integer to result
// END FOR

// RETURN result

// END FUNCTION

package piscine

func StringToIntSlice(str string) []int {
	var result []int
	for _, r := range str {
		result = append(result, int(r))
	}
	return result
}
