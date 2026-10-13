package smetrics

// Metaphone computes the Metaphone phonetic encoding of the given string.
// It is a phonetic algorithm that produces a code representing the English
// pronunciation of a word. Words that sound similar should produce the same code.
// The algorithm was published by Lawrence Philips in "Hanging on the Metaphone",
// Computer Language, Vol. 7, No. 12, December 1990.
//
// Unlike some implementations that truncate to 4 or 6 characters, this function
// returns the full encoding. Callers can truncate the result if needed.
func Metaphone(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Filter to letters only and convert to uppercase.
	chars := filterLetters(s)
	if len(chars) == 0 {
		return ""
	}

	// Transform first characters.
	chars = transformFirst(chars)
	n := len(chars)
	if n == 1 {
		return string(chars)
	}

	result := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		c := chars[i]

		// Skip duplicate consecutive characters (except C).
		if c != 'C' && i > 0 && chars[i-1] == c {
			continue
		}

		switch c {
		case 'A', 'E', 'I', 'O', 'U':
			// Vowels are only kept at the beginning.
			if i == 0 {
				result = append(result, c)
			}

		case 'B':
			// B is silent in final MB.
			if !(i == n-1 && i > 0 && chars[i-1] == 'M') {
				result = append(result, 'B')
			}

		case 'C':
			if i > 0 && chars[i-1] == 'S' && i+1 < n && isFrontvByte(chars[i+1]) {
				// SCI, SCE, SCY — C is silent.
				continue
			}
			if i > 0 && chars[i-1] == 'S' && i+1 < n && chars[i+1] == 'H' {
				// SCH → SK
				result = append(result, 'K')
			} else if (i+1 < n && chars[i+1] == 'H') || (i+2 < n && chars[i+1] == 'I' && chars[i+2] == 'A') {
				// CIA or CH → X
				result = append(result, 'X')
			} else if i+1 < n && isFrontvByte(chars[i+1]) {
				// CI, CE, CY → S
				result = append(result, 'S')
			} else {
				result = append(result, 'K')
			}

		case 'D':
			if i+2 < n && chars[i+1] == 'G' && isFrontvByte(chars[i+2]) {
				// DGE, DGI, DGY → J
				result = append(result, 'J')
				i += 2
			} else {
				result = append(result, 'T')
			}

		case 'F':
			result = append(result, 'F')

		case 'G':
			// GH at end or before consonant → silent.
			if i+1 < n && chars[i+1] == 'H' {
				if i+1 == n-1 || !isVowelByte(chars[i+2]) {
					continue
				}
			}
			// GN not at start → silent G.
			if i > 0 && i+1 < n && chars[i+1] == 'N' {
				continue
			}
			if i+1 < n && isFrontvByte(chars[i+1]) {
				result = append(result, 'J')
			} else {
				result = append(result, 'K')
			}

		case 'H':
			// H is silent at end, after varson chars (C, S, P, T, G), or
			// before a consonant. Digraph handlers (SH, PH, TH, CH) don't
			// advance past the H; instead this rule silences it because
			// the leading consonant is always in the varson set.
			if i == n-1 {
				continue
			}
			if i > 0 && isVarsonByte(chars[i-1]) {
				continue
			}
			if isVowelByte(chars[i+1]) {
				result = append(result, 'H')
			}

		case 'J':
			result = append(result, 'J')

		case 'K':
			// K after C is silent.
			if i > 0 && chars[i-1] == 'C' {
				continue
			}
			result = append(result, 'K')

		case 'L':
			result = append(result, 'L')

		case 'M':
			result = append(result, 'M')

		case 'N':
			result = append(result, 'N')

		case 'P':
			if i+1 < n && chars[i+1] == 'H' {
				result = append(result, 'F')
			} else {
				result = append(result, 'P')
			}

		case 'Q':
			result = append(result, 'K')

		case 'R':
			result = append(result, 'R')

		case 'S':
			if i+1 < n && chars[i+1] == 'H' {
				result = append(result, 'X')
			} else if i+2 < n && chars[i+1] == 'I' && (chars[i+2] == 'O' || chars[i+2] == 'A') {
				result = append(result, 'X')
			} else {
				result = append(result, 'S')
			}

		case 'T':
			if i+2 < n && chars[i+1] == 'I' && (chars[i+2] == 'A' || chars[i+2] == 'O') {
				result = append(result, 'X')
			} else if i+2 < n && chars[i+1] == 'C' && chars[i+2] == 'H' {
				// TCH — T is silent.
				continue
			} else if i+1 < n && chars[i+1] == 'H' {
				result = append(result, '0') // theta
			} else {
				result = append(result, 'T')
			}

		case 'V':
			result = append(result, 'F')

		case 'W':
			if i+1 < n && isVowelByte(chars[i+1]) {
				result = append(result, 'W')
			}

		case 'X':
			result = append(result, 'K')
			result = append(result, 'S')

		case 'Y':
			if i+1 < n && isVowelByte(chars[i+1]) {
				result = append(result, 'Y')
			}

		case 'Z':
			result = append(result, 'S')
		}
	}

	return string(result)
}

// filterLetters removes non-letter characters and converts to uppercase.
func filterLetters(s string) []byte {
	count := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			count++
		}
	}
	out := make([]byte, count)
	j := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			out[j] = c - 32
			j++
		} else if c >= 'A' && c <= 'Z' {
			out[j] = c
			j++
		}
	}
	return out
}

// transformFirst applies first-character transformations per the original algorithm:
// KN, GN, PN → drop first letter; AE → drop A; WR → drop W; WH → drop H; X → S.
func transformFirst(s []byte) []byte {
	if len(s) == 1 {
		if s[0] == 'X' {
			s[0] = 'S'
		}
		return s
	}
	if len(s) < 2 {
		return s
	}
	switch s[0] {
	case 'K', 'G', 'P':
		if s[1] == 'N' {
			return s[1:]
		}
	case 'A':
		if s[1] == 'E' {
			return s[1:]
		}
	case 'W':
		if s[1] == 'R' {
			return s[1:]
		}
		if s[1] == 'H' {
			// WH → W: drop the H.
			copy(s[1:], s[2:])
			return s[:len(s)-1]
		}
	case 'X':
		s[0] = 'S'
	}
	return s
}

func isVowelByte(c byte) bool {
	return c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func isFrontvByte(c byte) bool {
	return c == 'E' || c == 'I' || c == 'Y'
}

func isVarsonByte(c byte) bool {
	return c == 'C' || c == 'S' || c == 'P' || c == 'T' || c == 'G'
}
