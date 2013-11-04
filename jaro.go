package smetrics

import (
	"math"
)

func Jaro(a, b string) float64 {
	matchRange := int(math.Floor(math.Max(float64(len(a)), float64(len(b))) / float64(2))) - 1
	matchRange = int(math.Max(0, float64(matchRange - 1)))
	var matches, halfs float64
	transposed := make([]bool, len(b))

	for i := 0; i < len(a); i++ {
		start := int(math.Max(0, float64(i-matchRange)))
		end := int(math.Min(float64(len(b)-1), float64(i+matchRange)))

		for j := start; j <= end; j++ {
			if transposed[j] {
				continue
			}

			if a[i] == b[j] {
				if i != j {
					halfs++
				}
				matches++
				transposed[j] = true
				break
			}
		}
	}

	return ((matches/float64(len(a))) + (matches/float64(len(b))) + ((matches-math.Floor(float64(halfs/2))))/matches) / float64(3)
}
