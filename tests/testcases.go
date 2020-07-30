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

	// Equal strings.
	{a: "", b: "", r: 1.000},
	{a: "A", b: "A", r: 1.000},
	{a: "AA", b: "AA", r: 1.000},
	{a: "AAA", b: "AAA", r: 1.000},
	{a: "AAAA", b: "AAAA", r: 1.000},
	{a: "AAAAA", b: "AAAAA", r: 1.000},
	{a: "AAAAAA", b: "AAAAAA", r: 1.000},
	{
		a: "Legend of the Galactic Heroes",
		b: "Legend of the Galactic Heroes",
		r: 1.000,
	},
	{
		a: "Home is the place where, when you have to go there, they have to take you in.",
		b: "Home is the place where, when you have to go there, they have to take you in.",
		r: 1.000,
	},
	{
		a: "Pedro de Alcântara João Carlos Leopoldo Salvador Bibiano Francisco Xavier de Paula Leocádio Miguel Gabriel Rafael Gonzaga de Habsburgo-Lorena e Bragança",
		b: "Pedro de Alcântara João Carlos Leopoldo Salvador Bibiano Francisco Xavier de Paula Leocádio Miguel Gabriel Rafael Gonzaga de Habsburgo-Lorena e Bragança",
		r: 1.000,
	},
	{
		a: "Et tu, Brute",
		b: "Et tu, Brute",
		r: 1.000,
	},

	// Completely different strings.
	{a: "", b: "A", r: 0.000},
	{a: "", b: "AA", r: 0.000},
	{a: "", b: "AAA", r: 0.000},
	{a: "", b: "AAAA", r: 0.000},
	{a: "", b: "AAAAA", r: 0.000},
	{a: "A", b: "", r: 0.000},
	{a: "AA", b: "", r: 0.000},
	{a: "AAA", b: "", r: 0.000},
	{a: "AAAA", b: "", r: 0.000},
	{a: "AAAAA", b: "", r: 0.000},
	{a: "A", b: "B", r: 0.000},
	{a: "AA", b: "BB", r: 0.000},
	{a: "AAA", b: "BBB", r: 0.000},
	{a: "AAAA", b: "BBBB", r: 0.000},
	{a: "AAAAa", b: "BBBBB", r: 0.000},
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
	{a: "w", b: "w", r: 1.000},

	// Equal strings.
	{a: "", b: "", r: 1.000},
	{a: "A", b: "A", r: 1.000},
	{a: "AA", b: "AA", r: 1.000},
	{a: "AAA", b: "AAA", r: 1.000},
	{a: "AAAA", b: "AAAA", r: 1.000},
	{a: "AAAAA", b: "AAAAA", r: 1.000},
	{a: "AAAAAA", b: "AAAAAA", r: 1.000},
	{
		a: "Legend of the Galactic Heroes",
		b: "Legend of the Galactic Heroes",
		r: 1.000,
	},
	{
		a: "Home is the place where, when you have to go there, they have to take you in.",
		b: "Home is the place where, when you have to go there, they have to take you in.",
		r: 1.000,
	},
	{
		a: "Pedro de Alcântara João Carlos Leopoldo Salvador Bibiano Francisco Xavier de Paula Leocádio Miguel Gabriel Rafael Gonzaga de Habsburgo-Lorena e Bragança",
		b: "Pedro de Alcântara João Carlos Leopoldo Salvador Bibiano Francisco Xavier de Paula Leocádio Miguel Gabriel Rafael Gonzaga de Habsburgo-Lorena e Bragança",
		r: 1.000,
	},
	{
		a: "Et tu, Brute",
		b: "Et tu, Brute",
		r: 1.000,
	},

	// Completely different strings.
	{a: "", b: "A", r: 0.000},
	{a: "", b: "AA", r: 0.000},
	{a: "", b: "AAA", r: 0.000},
	{a: "", b: "AAAA", r: 0.000},
	{a: "", b: "AAAAA", r: 0.000},
	{a: "A", b: "", r: 0.000},
	{a: "AA", b: "", r: 0.000},
	{a: "AAA", b: "", r: 0.000},
	{a: "AAAA", b: "", r: 0.000},
	{a: "AAAAA", b: "", r: 0.000},
	{a: "A", b: "B", r: 0.000},
	{a: "AA", b: "BB", r: 0.000},
	{a: "AAA", b: "BBB", r: 0.000},
	{a: "AAAA", b: "BBBB", r: 0.000},
	{a: "AAAAa", b: "BBBBB", r: 0.000},
}
