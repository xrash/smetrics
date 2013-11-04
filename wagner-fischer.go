package smetrics

import (
	"math"
)

func WagnerFischer(a, b string, icost, dcost, scost int) int {
	lowerCost := int(math.Min(float64(icost), math.Min(float64(dcost), float64(scost))))

	// allocate enough memory for the matrix
	d := make([][]int, len(a)+1)
	for i, _ := range d {
		d[i] = make([]int, len(b)+1)
	}

	// initialize the values
	for i := 1; i <= len(a); i++ {
		d[i][0] = i * lowerCost
	}
	for i := 1; i <= len(b); i++ {
		d[0][i] = i * lowerCost
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				insertion := float64(d[i][j-1] + icost)
				deletion := float64(d[i-1][j] + dcost)
				substitution := float64(d[i-1][j-1] + scost)
				d[i][j] = int(math.Min(deletion, math.Min(insertion, substitution)))
			}
		}
	}

	return d[len(a)][len(b)]
}
