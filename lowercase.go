package cjk

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

type LowerCaseFilter struct{}

func NewLowerCaseFilter() *LowerCaseFilter {
	return &LowerCaseFilter{}
}

func (f *LowerCaseFilter) Filter(input TokenStream) TokenStream {
	for _, token := range input {
		token.Term = toLower(token.Term)
	}
	return input
}

func toLower(s []byte) []byte {
	j := 0
	b := make([]byte, len(s))
	copy(b, s)
	for i := 0; i < len(b); {
		wid := 1
		r := rune(b[i])
		if r >= utf8.RuneSelf {
			r, wid = utf8.DecodeRune(b[i:])
		}

		l := unicode.ToLower(r)

		if l == r {
			i += wid
			j += wid
			continue
		}

		// Handles the Unicode edge-case where the last
		// rune in a word on the greek Σ needs to be converted
		// differently.
		if l == 'σ' && i+2 == len(b) {
			l = 'ς'
		}

		lwid := utf8.RuneLen(l)
		if lwid > wid {
			// utf-8 encoded replacement is wider
			// fallback to bytes.ToLower
			return bytes.ToLower(s)
		} else {
			utf8.EncodeRune(b[j:], l)
		}
		i += wid
		j += lwid
	}
	return b[:j]
}
