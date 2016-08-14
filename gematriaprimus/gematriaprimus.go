package gematriaprimus

type Triple struct {
	Rune    rune
	Letters []string
	Value   uint
}

const (
	SPACE = '•'
)

var (
	Triples = []Triple{
		{'ᚠ', []string{"f"}, 2},
		{'ᚢ', []string{"u"}, 3},
		{'ᚦ', []string{"th"}, 5},
		{'ᚩ', []string{"o"}, 7},
		{'ᚱ', []string{"r"}, 11},
		{'ᚳ', []string{"c", "k"}, 13},
		{'ᚷ', []string{"g"}, 17},
		{'ᚹ', []string{"w"}, 19},
		{'ᚻ', []string{"h"}, 23},
		{'ᚾ', []string{"n"}, 29},
		{'ᛁ', []string{"i"}, 31},
		{'ᛂ', []string{"j"}, 37},
		{'ᛇ', []string{"eo"}, 41},
		{'ᛈ', []string{"p"}, 43},
		{'ᛉ', []string{"x"}, 47},
		{'ᛋ', []string{"s", "z"}, 53},
		{'ᛏ', []string{"t"}, 59},
		{'ᛒ', []string{"b"}, 61},
		{'ᛖ', []string{"e"}, 67},
		{'ᛗ', []string{"m"}, 71},
		{'ᛚ', []string{"l"}, 73},
		{'ᛝ', []string{"ng", "ing"}, 79},
		{'ᛟ', []string{"oe"}, 83},
		{'ᛞ', []string{"d"}, 89},
		{'ᚪ', []string{"a"}, 97},
		{'ᚫ', []string{"ae"}, 101},
		{'ᚣ', []string{"y"}, 103},
		{'ᛡ', []string{"ia", "io"}, 107},
		{'ᛠ', []string{"ea"}, 109},
	}
	Runes = func() []rune {
		p := make([]rune, len(Triples))
		for i, t := range Triples {
			p[i] = t.Rune
		}
		return p
	}()
	Values = func() []uint {
		p := make([]uint, len(Triples))
		for i, t := range Triples {
			p[i] = t.Value
		}
		return p
	}()
	Letters = func() []string {
		p := make([]string, len(Triples))
		for i, t := range Triples {
			p[i] = t.Letters[0]
		}
		return p
	}()
)

func IsRune(r rune) bool {
	return r == 'ᚠ' || r == 'ᚢ' || r == 'ᚦ' || r == 'ᚩ' ||
		r == 'ᚱ' || r == 'ᚳ' || r == 'ᚷ' || r == 'ᚹ' ||
		r == 'ᚻ' || r == 'ᚾ' || r == 'ᛁ' || r == 'ᛂ' ||
		r == 'ᛇ' || r == 'ᛈ' || r == 'ᛉ' || r == 'ᛋ' ||
		r == 'ᛏ' || r == 'ᛒ' || r == 'ᛖ' || r == 'ᛗ' ||
		r == 'ᛚ' || r == 'ᛝ' || r == 'ᛟ' || r == 'ᛞ' ||
		r == 'ᚪ' || r == 'ᚫ' || r == 'ᚣ' || r == 'ᛡ' ||
		r == 'ᛠ'
}

func RuneToValue(r rune) uint {
	for _, t := range Triples {
		if t.Rune == r {
			return t.Value
		}
	}
	return 0
}

func RuneToLetter(r rune) string {
	for _, t := range Triples {
		if t.Rune == r {
			return t.Letters[0]
		}
	}
	return ""
}

func RuneIndex(r rune) int {
	for i, t := range Triples {
		if t.Rune == r {
			return i
		}
	}
	return 0
}
