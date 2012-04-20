package main

import "fmt"

type Range struct {
	Lo float64
	Hi float64
}

func (r *Range) has(x float64) bool {
	return x >= r.Lo
}

type Percentage float64

func (p Percentage) apply(amt float64) float64 {
	return float64(p) * amt
}

type Bracket struct {
	Range
	Percentage
}

func NewBracket(lo, hi, p float64) Bracket {
	return Bracket{Range{lo, hi}, Percentage(p)}
}

var brackets = []Bracket{
	// 2011 tax brackets.
	// http://en.wikipedia.org/wiki/Income_tax_in_the_United_States#Year_2011_income_brackets_and_tax_rates
	NewBracket(0, 8.5e3, 0.10),
	NewBracket(8.5e3+1, 34.5e3, 0.15),
	NewBracket(34.5e3+1, 83.6e3, 0.25),
	NewBracket(83.6e3+1, 174e3, 0.28),
	NewBracket(174e3+1, 379.15e3, 0.33),
	NewBracket(379.15e3, 1e10, 0.35),
}

func main() {
	// Sample income rates.
	income := []float64{9e3, 1e4, 15e3, 2e4, 3e4, 3e5, 3e6, 3e7, 3e9}

	for _, i := range income {
		var tax float64
		for _, b := range brackets {
			if b.Range.has(i) {
				if i < b.Hi {
					// This is the last bracket for this income.
					tax += b.Percentage.apply(i - b.Lo)
					break
				} else {
					tax += b.Percentage.apply(b.Hi - b.Lo)
				}
			}
		}
		fmt.Printf("Tax for %v is %v. Effective rate is %v\n", i, tax, tax/i)
	}
}
