package smetrics

import (
	"math"
)

func Ukkonen(a, b string, icost, dcost, scost int) int {
	lowerCost := int(math.Min(float64(icost), math.Min(float64(dcost), float64(scost))))
	infinite := math.MaxInt32/2
	var r []int

	if tmp := b; len(a) > len(b) {
		b = a
		a = tmp
	}

	var k, kprime, p int

	t := (len(b) - len(a) + 1) * lowerCost

	for {
		if (t/lowerCost) < (len(b) - len(a)) {
			continue
		}

		p = int(math.Floor(0.5*((float64(t)/float64(lowerCost)) - float64(len(b) - len(a)))))

		k = -p
		kprime = k

		rowlength := (len(b) - len(a)) + (2*p)

		r = make([]int, rowlength + 2)

		for i := 0; i < rowlength + 2; i++ {
			r[i] = infinite
		}

		for i := 0; i <= len(a); i++ {
			for j := 0; j <= rowlength; j++ {
				if i == j+k && i == 0 {
					r[j] = 0
				} else {
					var ins int
					if (j-1 < 0) {
						ins = infinite
					} else {
						ins = r[j-1] + icost
					}

					del := r[j+1] + dcost
					sub := r[j] + scost

					if i-1 < 0 || i-1 >= len(a) || j+k-1 >= len(b) || j+k-1 < 0 {
						sub = infinite
					} else if a[i-1] == b[j+k-1] {
						sub = r[j]
					}

					r[j] = int(math.Min(float64(ins), math.Min(float64(del), float64(sub))))
				}
			}
			k++
		}

		if r[(len(b) - len(a)) + (2 * p) + kprime] <= t {
			break;
		} else {
			t *= 2
		}
	}

	return r[(len(b) - len(a)) + (2 * p) + kprime]
}
