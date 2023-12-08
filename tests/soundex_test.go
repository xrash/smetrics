package tests

import (
	"fmt"
	"github.com/xrash/smetrics"
	"testing"
)

//func BenchmarkSoundex(b *testing.B) {
//	fd, err := os.Open("testdata/english_names.txt")
//	if err != nil {
//		b.Fatal(err)
//	}
//	defer fd.Close()
//
//	scanner := bufio.NewScanner(fd)
//	for scanner.Scan() {
//		s := smetrics.Soundex(scanner.Text())
//
//		// Trivially use the result to avoid compiler optimizations.
//		if len(s) != 4 {
//			b.Fail()
//		}
//	}
//}

func TestSoundex(t *testing.T) {
	cases := []soundexcase{
		{"Euler", "E460"},
		{"Ellery", "E460"},
		{"Gauss", "G200"},
		{"Ghosh", "G200"},
		{"Hilbert", "H416"},
		{"Heilbrohn", "H416"},
		{"Knuth", "K530"},
		{"Kant", "K530"},
		{"Lloyd", "L300"},
		{"Ladd", "L300"},
		{"Lukasiewicz", "L222"},
		{"Lissjous", "L222"},
		{"Ravi", "R100"},
		{"Ravee", "R100"},
	}

	for _, c := range cases {
		if r := smetrics.Soundex(c.s); r != c.t {
			fmt.Println(r, "instead of", c.t, "for", c.s)
			t.Fail()
		}
	}
}
