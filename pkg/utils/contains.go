package utils

// @summary: Check if a string slice contains a string.
// @param: slice []string - The slice to search.
// @param: s string - The string to search for.
// @return: bool - True if the string is found in the slice, false otherwise.
// @example: StringContains([]string{"a", "b", "c"}, "a") -> true.
// @example: StringContains([]string{"a", "b", "c"}, "d") -> false.
func StringContains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
