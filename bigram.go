package cjk

import (
	"container/ring"
	"unicode/utf8"
)

type CJKBigramFilter struct {
	outputUnigram bool
}

func NewCJKBigramFilter(outputUnigram bool) *CJKBigramFilter {
	return &CJKBigramFilter{
		outputUnigram: outputUnigram,
	}
}

func (s *CJKBigramFilter) Filter(input TokenStream) TokenStream {
	r := ring.New(2)
	itemsInRing := 0
	pos := 1
	outputPos := 1

	rv := make(TokenStream, 0, len(input))

	for _, tokout := range input {
		if tokout.Type == Ideographic {
			runes := []rune(string(tokout.Term))
			sofar := 0
			for _, run := range runes {
				rlen := utf8.RuneLen(run)
				token := &Token{
					Term:     tokout.Term[sofar : sofar+rlen],
					Start:    tokout.Start + sofar,
					End:      tokout.Start + sofar + rlen,
					Position: pos,
					Type:     tokout.Type,
					KeyWord:  tokout.KeyWord,
				}
				pos++
				sofar += rlen
				if itemsInRing > 0 {
					// if items already buffered
					// check to see if this is aligned
					curr := r.Value.(*Token)
					if token.Start-curr.End != 0 {
						// not aligned flush
						flushToken := s.flush(r, &itemsInRing, outputPos)
						if flushToken != nil {
							outputPos++
							rv = append(rv, flushToken)
						}
					}
				}
				// now we can add this token to the buffer
				r = r.Next()
				r.Value = token
				if itemsInRing < 2 {
					itemsInRing++
				}
				builtUnigram := false
				if itemsInRing > 1 && s.outputUnigram {
					unigram := s.buildUnigram(r, &itemsInRing, outputPos)
					if unigram != nil {
						builtUnigram = true
						rv = append(rv, unigram)
					}
				}
				bigramToken := s.outputBigram(r, &itemsInRing, outputPos)
				if bigramToken != nil {
					rv = append(rv, bigramToken)
					outputPos++
				}

				// prev token should be removed if unigram was built
				if builtUnigram {
					itemsInRing--
				}
			}

		} else {
			// flush anything already buffered
			flushToken := s.flush(r, &itemsInRing, outputPos)
			if flushToken != nil {
				rv = append(rv, flushToken)
				outputPos++
			}
			// output this token as is
			tokout.Position = outputPos
			rv = append(rv, tokout)
			outputPos++
		}
	}

	// deal with possible trailing unigram
	if itemsInRing == 1 || s.outputUnigram {
		if itemsInRing == 2 {
			r = r.Next()
		}
		unigram := s.buildUnigram(r, &itemsInRing, outputPos)
		if unigram != nil {
			rv = append(rv, unigram)
		}
	}
	return rv
}

func (s *CJKBigramFilter) flush(r *ring.Ring, itemsInRing *int, pos int) *Token {
	var rv *Token
	if *itemsInRing == 1 {
		rv = s.buildUnigram(r, itemsInRing, pos)
	}
	r.Value = nil
	*itemsInRing = 0

	return rv
}

func (s *CJKBigramFilter) outputBigram(r *ring.Ring, itemsInRing *int, pos int) *Token {
	if *itemsInRing == 2 {
		thisShingleRing := r.Move(-1)
		shingledBytes := make([]byte, 0)

		// do first token
		prev := thisShingleRing.Value.(*Token)
		shingledBytes = append(shingledBytes, prev.Term...)

		// do second token
		thisShingleRing = thisShingleRing.Next()
		curr := thisShingleRing.Value.(*Token)
		shingledBytes = append(shingledBytes, curr.Term...)

		token := Token{
			Type:     Double,
			Term:     shingledBytes,
			Position: pos,
			Start:    prev.Start,
			End:      curr.End,
		}
		return &token
	}

	return nil
}

func (s *CJKBigramFilter) buildUnigram(r *ring.Ring, itemsInRing *int, pos int) *Token {
	switch *itemsInRing {
	case 2:
		thisShingleRing := r.Move(-1)
		// do first token
		prev := thisShingleRing.Value.(*Token)
		token := Token{
			Type:     Single,
			Term:     prev.Term,
			Position: pos,
			Start:    prev.Start,
			End:      prev.End,
		}
		return &token
	case 1:
		// do first token
		prev := r.Value.(*Token)
		token := Token{
			Type:     Single,
			Term:     prev.Term,
			Position: pos,
			Start:    prev.Start,
			End:      prev.End,
		}
		return &token
	}

	return nil
}
