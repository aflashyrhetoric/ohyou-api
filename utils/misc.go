package utils

import "math/rand"

// ArrayContainsInt ...Checks to see if list contains a
func ArrayContainsInt(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// RandomBool returns TRUE or FALSE randomly
func RandomBool() bool {
	return ((rand.Intn(2)+1)%2 == 0)
}
