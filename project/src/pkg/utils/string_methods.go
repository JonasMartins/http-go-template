package utils

// given a string term and a array
// of strings, it will return true
// if the arrays contains the string,
// false otherwise
func StringContains(str *string, arr *[]string) bool {
	for _, c := range *arr {
		if *str == c {
			return true
		}
	}
	return false
}
