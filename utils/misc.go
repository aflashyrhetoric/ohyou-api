package utils

// ArrayContainsInt ...Checks to see if list contains a
func ArrayContainsInt(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
