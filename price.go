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

// ConvertDollarsToCents - Converts a float64 currency to integers (in cents/pennies)
func ConvertDollarsToCents(price float64, err error) (int, error) {
	return int(price * 100), nil
}

// ConvertCentsToDollars - Converts a integers currency to float64
func ConvertCentsToDollars(price int, err error) (float64, error) {
	return float64(price * .01), nil
}
