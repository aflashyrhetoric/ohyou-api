package main

import "strconv"

// Price - Representation of a cost of an item
type Price struct {
	Dollars int  `json:"dollar"`
	Cents   int8 `json:"cents"`
}

// ConvertToPrice - Converts int (cents) to price
func ConvertToPrice(price int, err error) (Price, error) {

	// price is received as pennies

	stringPrice := string(price)
	dollars, _ := strconv.Atoi(stringPrice[:len(stringPrice)-2])
	cents, _ := strconv.Atoi(stringPrice[len(stringPrice)-2:])

	return Price{Dollars: dollars, Cents: int8(cents)}, err
}
