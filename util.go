package cjk

import "unicode/utf8"

func buildTermFromRunes(runes []rune) []byte {
	buf := make([]byte, utf8.UTFMax*len(runes))
	offset := 0
	for _, r := range runes {
		offset += utf8.EncodeRune(buf[offset:], r)
	}
	return buf[:offset]
}

func deleteRune(runes []rune, i int) []rune {
	return append(runes[:i], runes[i+1:]...)
}
