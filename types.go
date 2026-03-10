package cjk

type TokenType int

const (
	None TokenType = iota
	Alpha
	Numeric
	Ideographic // CJK characters
	Double      // Bigrams
	Single      // Unigrams
)

type Token struct {
	Term     []byte
	Start    int
	End      int
	Position int
	Type     TokenType
	KeyWord  bool
}

type TokenStream []*Token

type Tokenizer interface {
	Tokenize(text []byte) TokenStream
}

type TokenFilter interface {
	Filter(input TokenStream) TokenStream
}
