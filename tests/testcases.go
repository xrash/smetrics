package tests

type levenshteincase struct {
	s     string
	t     string
	icost int
	dcost int
	scost int
	r     int
}

type soundexcase struct {
	s string
	t string
}

type hammingcase struct {
	a    string
	b    string
	diff int
}

type jarocase struct {
	a string
	b string
	r float64
}

var __jaro_cases = []*jarocase{
	{a: "SHACKLEFORD", b: "SHACKELFORD", r: 0.970},
	{a: "DUNNINGHAM", b: "CUNNIGHAM", r: 0.896},
	{a: "NICHLESON", b: "NICHULSON", r: 0.926},
	{a: "JONES", b: "JOHNSON", r: 0.790},
	{a: "MASSEY", b: "MASSIE", r: 0.889},
	{a: "ABROMS", b: "ABRAMS", r: 0.889},
	{a: "HARDIN", b: "MARTINEZ", r: 0.722},
	{a: "ITMAN", b: "SMITH", r: 0.467},
	{a: "JERALDINE", b: "GERALDINE", r: 0.926},
	{a: "MARHTA", b: "MARTHA", r: 0.944},
	{a: "MICHELLE", b: "MICHAEL", r: 0.869},
	{a: "JULIES", b: "JULIUS", r: 0.889},
	{a: "TANYA", b: "TONYA", r: 0.867},
	{a: "DWAYNE", b: "DUANE", r: 0.822},
	{a: "SEAN", b: "SUSAN", r: 0.783},
	{a: "JON", b: "JOHN", r: 0.917},
//	{a: "JON", b: "JAN", r: 0.000},
	{a: "BROOKHAVEN", b: "BRROKHAVEN", r: 0.933},
	{a: "BROOK HALLOW", b: "BROOK HLLW", r: 0.944},
	{a: "DECATUR", b: "DECATIR", r: 0.905},
	{a: "FITZRUREITER", b: "FITZENREITER", r: 0.856},
	{a: "HIGBEE", b: "HIGHEE", r: 0.889},
	{a: "HIGBEE", b: "HIGVEE", r: 0.889},
	{a: "LACURA", b: "LOCURA", r: 0.889},
	{a: "IOWA", b: "IONA", r: 0.833},
//	{a: "1ST", b: "IST", r: 0.000},
}

var __jaro_winkler_cases = []*jarocase{
	{a: "SHACKLEFORD", b: "SHACKELFORD", r: 0.982},
	{a: "DUNNINGHAM", b: "CUNNIGHAM", r: 0.896},
	{a: "NICHLESON", b: "NICHULSON", r: 0.956},
	{a: "JONES", b: "JOHNSON", r: 0.832},
	{a: "MASSEY", b: "MASSIE", r: 0.933},
	{a: "ABROMS", b: "ABRAMS", r: 0.922},
	{a: "HARDIN", b: "MARTINEZ", r: 0.722},
	{a: "ITMAN", b: "SMITH", r: 0.467},
	{a: "JERALDINE", b: "GERALDINE", r: 0.926},
	{a: "MARHTA", b: "MARTHA", r: 0.961},
	{a: "MICHELLE", b: "MICHAEL", r: 0.921},
	{a: "JULIES", b: "JULIUS", r: 0.933},
	{a: "TANYA", b: "TONYA", r: 0.880},
	{a: "DWAYNE", b: "DUANE", r: 0.840},
	{a: "SEAN", b: "SUSAN", r: 0.805},
	{a: "JON", b: "JOHN", r: 0.933},
//	{a: "JON", b: "JAN", r: 0.000},
	{a: "BROOKHAVEN", b: "BRROKHAVEN", r: 0.947},
	{a: "BROOK HALLOW", b: "BROOK HLLW", r: 0.967},
	{a: "DECATUR", b: "DECATIR", r: 0.943},
	{a: "FITZRUREITER", b: "FITZENREITER", r: 0.913},
	{a: "HIGBEE", b: "HIGHEE", r: 0.922},
	{a: "HIGBEE", b: "HIGVEE", r: 0.922},
	{a: "LACURA", b: "LOCURA", r: 0.900},
	{a: "IOWA", b: "IONA", r: 0.867},
//	{a: "1ST", b: "IST", r: 0.000},
}
