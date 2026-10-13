package tests

import (
	"fmt"
	"testing"

	"github.com/xrash/smetrics"
)

func TestMetaphone(t *testing.T) {
	cases := []metaphonecase{
		// Empty and single character.
		{"", ""},
		{"A", "A"},
		{"a", "A"},
		{"B", "B"},

		// Non-letter input.
		{"123", ""},
		{"a1b", "AB"},

		// Common words.
		{"Donald", "TNLT"},
		{"Zach", "SX"},
		{"Campbel", "KMPBL"},
		{"David", "TFT"},
		{"Metaphone", "MTFN"},
		{"garbage", "KRBJ"},
		{"tech", "TX"},
		{"orange", "ORNJ"},
		{"science", "SNS"},
		{"dumb", "TM"},
		{"dodge", "TJ"},
		{"jack", "JK"},
		{"queen", "KN"},
		{"widow", "WT"},
		{"xerox", "SRKS"},
		{"vowel", "FWL"},
		{"zen", "SN"},
		{"tiamat", "XMT"},
		{"ratio", "RX"},
		{"the", "0"},
		{"acacia", "AKX"},
		{"bosch", "BSK"},
		{"chop", "XP"},
		{"hatch", "HX"},
		{"Incomprehensibility", "INKMPRHNSBLT"},

		// Initial character transforms.
		{"knight", "NT"},       // KN → N
		{"gnome", "NM"},        // GN → N
		{"pneumatic", "NMTK"},  // PN → N
		{"aesthetic", "ES0TK"}, // AE → E
		{"wrong", "RNK"},       // WR → R

		// Silent letters and digraphs.
		{"lamb", "LM"},  // silent B after M at end
		{"phone", "FN"}, // PH → F
		{"ship", "XP"},  // SH → X
		{"thick", "0K"}, // TH → 0 (theta)
		{"match", "MX"}, // TCH → silent T, CH → X
		{"edge", "EJ"},  // DGE → J

		// CH always maps to X per original algorithm.
		{"church", "XRX"},
		{"chess", "XS"},

		// Full output without truncation.
		{"crying", "KRYNK"},

		// GH silent at end and before consonant.
		{"weigh", "W"},
		{"ghost", "KST"},

		// WH → W at start.
		{"wh", "W"},
		{"whistle", "WSTL"},

		// GN silences G at any non-initial position (per original algorithm).
		{"signal", "SNL"},

		// Additional test cases
		{"programming", "PRKRMNK"},
		{"programmer", "PRKRMR"},
		{"Asterix", "ASTRKS"},
		{"Plaçe", "PL"}, // ç ignored, P L A E → P L
		{"Place", "PLS"},
		{"aebersold", "EBRSLT"}, // AE → E
		{"gnagy", "NJ"},         // GN → N, G before Y → J
		{"knuth", "N0"},         // KN → N, TH → 0
		{"pniewski", "NSK"},     // PN → N
		{"wright", "RT"},        // WR → R
		{"deng", "TNK"},         // D T, N N, G K
		{"xiaopeng", "XPNK"},    // S I A → X, P N K
		{"whalen", "WLN"},       // WH→W
		{"dumb", "TM"},          // B silent after M at end
		{"mccomb", "MKKM"},      // CC → K K, B silent
		{"cia", "X"},            // CIA → X
		{"ch", "X"},             // CH → X
		{"ci", "S"},             // CI → S
		{"ce", "S"},             // CE → S
		{"cy", "S"},             // CY → S
		{"sci", "S"},            // S I → S, C silent
		{"sce", "S"},            // S E → S, C silent
		{"scy", "S"},            // S Y → S, C silent
		{"sch", "SK"},           // SCH → SK
		{"dge", "J"},            // DGE → J
		{"dgy", "J"},            // DGY → J
		{"dgi", "J"},            // DGI → J
		{"gh", ""},              // GH silent at end
		{"gn", "N"},             // GN at start → N
		{"gned", "NT"},          // GNED → N T (G silent, NED → N E D → N T)
		{"gg", "K"},             // GG → K (not J)
		{"ck", "K"},             // CK → K
		{"ph", "F"},             // PH → F
		{"sh", "X"},             // SH → X
		{"sio", "X"},            // SIO → X
		{"sia", "X"},            // SIA → X
		{"tia", "X"},            // TIA → X
		{"tio", "X"},            // TIO → X
		{"th", "0"},             // TH → 0
		{"tch", "X"},            // T silent, C H → X
		{"v", "V"},              // Single letter
		{"w", "W"},              // Single letter
		{"wa", "W"},             // W followed by vowel
		{"x", "S"},              // Single letter, initial X → S
		{"y", "Y"},              // Single letter
		{"ya", "Y"},             // Y followed by vowel
		{"z", "Z"},              // Single letter

		// Additional edge cases for completeness
		{"scuba", "SKB"},      // SCUBA → S K U B A → S K B
		{"scissors", "SSRS"},  // SCISSORS → S C I S S O R S → S S R S
		{"queue", "K"},        // QUEUE → K U E U E → K
		{"rhythm", "R0M"},     // RHYTHM → R H Y T H M → R 0 M
		{"aaa", "A"},          // Duplicate vowels at start
		{"xyz", "SS"},         // X Y Z → S S
		{"123science", "SNS"}, // Non-letters ignored
		{"a1b2c", "ABK"},      // Mixed non-letters

		// Edge cases.
		{"aeiou", "E"},             // AE → E, then vowels dropped
		{"bcdfg", "BKTFK"},         // B K T F K
		{"hijklmnop", "HJKLMNP"},   // H J K L M N P
		{"qrstuvwxyz", "KRSTFKSS"}, // K R S T F K S S
	}

	for _, c := range cases {
		if r := smetrics.Metaphone(c.s); r != c.t {
			fmt.Println(r, "instead of", c.t, "for", c.s)
			t.Fail()
		}
	}
}

func BenchmarkMetaphone(b *testing.B) {
	testCases := []string{
		"programming",
		"metaphone",
		"incomprehensibility",
		"pneumatic",
		"knight",
		"wrong",
		"whistle",
		"signal",
		"asterix",
		"aebersold",
		"dengxiaopeng",
		"mccomb",
		"church",
		"ghost",
		"edge",
		"thick",
		"phone",
		"ship",
		"lamb",
		"zen",
		"vowel",
		"xerox",
		"queen",
		"jack",
		"dodge",
		"dumb",
		"science",
		"orange",
		"tech",
		"garbage",
		"metaphone",
		"david",
		"campbel",
		"zach",
		"donald",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testCases {
			smetrics.Metaphone(s)
		}
	}
}
