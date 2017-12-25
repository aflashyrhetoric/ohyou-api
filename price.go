package main

import "strconv"

// Representation of a cost of an item
type Price struct {
	Dollars int  `json:"dollar"`
	Cents   int8 `json:"cents"`
}

func (p *Price) toPrice(price int) Price {

	// price is received as pennies

	stringPrice := string(price)

	dollars, _ := strconv.Atoi(stringPrice[:len(stringPrice)-2])
	cents, _ := strconv.Atoi(stringPrice[len(stringPrice)-2:])

	return Price{Dollars: dollars, Cents: cents}
}
