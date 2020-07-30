package tests

import (
	"fmt"
	"github.com/xrash/smetrics"
	"testing"
)

func TestJaroWinkler(t *testing.T) {
	for _, c := range __jaro_winkler_cases {
		r := smetrics.JaroWinkler(c.a, c.b, 0.7, 4)
		result := fmt.Sprintf("%.3f", r)
		expected := fmt.Sprintf("%.3f", c.r)
		if result != expected {
			fmt.Println(c.a, c.b, result, "instead of", expected)
			t.Fail()
		}
	}
}
