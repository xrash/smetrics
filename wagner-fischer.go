package smetrics

import (
	"math"
)

func WagnerFischer(a, b string, icost, dcost, scost int) int {
	lowerCost := int(math.Min(float64(icost), math.Min(float64(dcost), float64(scost))))

	// Allocate the array that will hold the last row.
	row1 := make([]int, len(a) + 1)
	row2 := make([]int, len(a) + 1)

	// Initialize the arrays.
	for i := 1; i <= len(a); i++ {
		row1[i] = i * lowerCost
	}

	for i := 1; i <= len(b); i++ {
		row2[0] = row1[0] + lowerCost

		for j := 1; j <= len(a); j++ {
			if a[j-1] == b[i-1] {
				row2[j] = row1[j-1]
			} else {
				insertion := float64(row2[j-1] + icost)
				deletion := float64(row1[j] + dcost)
				substitution := float64(row1[j-1] + scost)
				row2[j] = int(math.Min(deletion, math.Min(insertion, substitution)))
			}
		}

		for j := 0; j < len(row1); j++ {
			row1[j] = row2[j]
		}
	}

	return row2[len(row2) - 1]
}
