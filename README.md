# Go CJK Tokenizer

An independent CJK (Chinese, Japanese, Korean) tokenizer for Go, extracted from the [Bleve](https://github.com/blevesearch/bleve) search engine.

This library provides:

- **Width Normalization**: Full-width ASCII to half-width, half-width Katakana to full-width.
- **Lowercase Filter**: Unicode-aware lowercasing.
- **Bigram Filter**: Generates bigrams for CJK ideographs (Han, Hiragana, Katakana, Hangul).
- **Unicode Tokenizer**: A basic script-aware tokenizer to identify CJK blocks.

## Feature

- **Standalone**: No dependency on `bleve`.
- **Easy to use**: Simple API for analyzing text.
- **Efficient**: Minimal allocations and reused buffers where possible.

## Installation

```bash
go get github.com/bitxeno/go-cjk-tokenizer
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/bitxeno/go-cjk-tokenizer"
)

func main() {
    analyzer := cjk.NewAnalyzer()
    tokens := analyzer.Analyze([]byte("Hello 世界, こんにちは!"))

    for _, token := range tokens {
        fmt.Printf("Token: %s, Start: %d, End: %d, Type: %d\n",
            string(token.Term), token.Start, token.End, token.Type)
    }
}
```

## Credits

Based on the implementation in [Bleve](https://github.com/blevesearch/bleve/tree/master/analysis/lang/cjk).
Original code is Copyright (c) 2014 Couchbase, Inc. and licensed under Apache License 2.0.


## Other Tokenizer

* [sego](https://github.com/huichen/sego): Go Chinese Word Segmentation
* [gojieba](https://github.com/yanyiwu/gojieba): "Jieba" Chinese word segmentation in Golang version
* [simple](https://github.com/wangfenjin/simple): A SQLite3 fts5 tokenizer which supports Chinese and PinYin
* [language-tokenizer](https://github.com/savannstm/language-tokenizer): Rust Text tokenizer for more than 40 languages
