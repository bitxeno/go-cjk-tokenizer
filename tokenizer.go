package cjk

import (
	"unicode"
	"unicode/utf8"
)

type UnicodeTokenizer struct{}

func NewUnicodeTokenizer() *UnicodeTokenizer {
	return &UnicodeTokenizer{}
}

func (t *UnicodeTokenizer) Tokenize(input []byte) TokenStream {
	rv := make(TokenStream, 0)
	pos := 1
	start := 0

	for start < len(input) {
		r, size := utf8.DecodeRune(input[start:])
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			start += size
			continue
		}

		end := start + size
		isIdeo := isIdeographic(r)
		
		// Find the end of the current token based on whether it's Ideographic or not
		for end < len(input) {
			rNext, sizeNext := utf8.DecodeRune(input[end:])
			if unicode.IsSpace(rNext) || unicode.IsPunct(rNext) {
				break
			}
			if isIdeo != isIdeographic(rNext) {
				break
			}
			end += sizeNext
		}

		typ := Alpha
		if isIdeo {
			typ = Ideographic
		} else if unicode.IsNumber(r) {
			typ = Numeric
		}

		token := &Token{
			Term:     input[start:end],
			Start:    start,
			End:      end,
			Position: pos,
			Type:     typ,
		}
		rv = append(rv, token)
		pos++
		start = end
	}

	return rv
}

func isIdeographic(r rune) bool {
	if unicode.Is(unicode.Han, r) ||
		unicode.Is(unicode.Hiragana, r) ||
		unicode.Is(unicode.Katakana, r) ||
		unicode.Is(unicode.Hangul, r) ||
		unicode.Is(unicode.Ideographic, r) {
		return true
	}
	return false
}
