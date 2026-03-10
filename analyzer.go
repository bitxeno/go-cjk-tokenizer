package cjk

type Analyzer struct {
	Tokenizer    Tokenizer
	TokenFilters []TokenFilter
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		Tokenizer: NewUnicodeTokenizer(),
		TokenFilters: []TokenFilter{
			NewCJKWidthFilter(),
			NewLowerCaseFilter(),
			NewCJKBigramFilter(false),
		},
	}
}

func (a *Analyzer) Analyze(text []byte) TokenStream {
	tokens := a.Tokenizer.Tokenize(text)
	for _, filter := range a.TokenFilters {
		tokens = filter.Filter(tokens)
	}
	return tokens
}
