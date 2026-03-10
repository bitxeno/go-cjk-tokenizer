package cjk

import (
	"reflect"
	"testing"
)

func TestCJKAnalyzer(t *testing.T) {
	tests := []struct {
		input  []byte
		output TokenStream
	}{
		{
			input: []byte("こんにちは世界"),
			output: TokenStream{
				{
					Term:     []byte("こん"),
					Type:     Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				{
					Term:     []byte("んに"),
					Type:     Double,
					Position: 2,
					Start:    3,
					End:      9,
				},
				{
					Term:     []byte("にち"),
					Type:     Double,
					Position: 3,
					Start:    6,
					End:      12,
				},
				{
					Term:     []byte("ちは"),
					Type:     Double,
					Position: 4,
					Start:    9,
					End:      15,
				},
				{
					Term:     []byte("は世"),
					Type:     Double,
					Position: 5,
					Start:    12,
					End:      18,
				},
				{
					Term:     []byte("世界"),
					Type:     Double,
					Position: 6,
					Start:    15,
					End:      21,
				},
			},
		},
		{
			input: []byte("一二三四五六七八九十"),
			output: TokenStream{
				{
					Term:     []byte("一二"),
					Type:     Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				{
					Term:     []byte("二三"),
					Type:     Double,
					Position: 2,
					Start:    3,
					End:      9,
				},
				{
					Term:     []byte("三四"),
					Type:     Double,
					Position: 3,
					Start:    6,
					End:      12,
				},
				{
					Term:     []byte("四五"),
					Type:     Double,
					Position: 4,
					Start:    9,
					End:      15,
				},
				{
					Term:     []byte("五六"),
					Type:     Double,
					Position: 5,
					Start:    12,
					End:      18,
				},
				{
					Term:     []byte("六七"),
					Type:     Double,
					Position: 6,
					Start:    15,
					End:      21,
				},
				{
					Term:     []byte("七八"),
					Type:     Double,
					Position: 7,
					Start:    18,
					End:      24,
				},
				{
					Term:     []byte("八九"),
					Type:     Double,
					Position: 8,
					Start:    21,
					End:      27,
				},
				{
					Term:     []byte("九十"),
					Type:     Double,
					Position: 9,
					Start:    24,
					End:      30,
				},
			},
		},
		{
			input: []byte("一 二三四 五六七八九 十"),
			output: TokenStream{
				{
					Term:     []byte("一"),
					Type:     Single,
					Position: 1,
					Start:    0,
					End:      3,
				},
				{
					Term:     []byte("二三"),
					Type:     Double,
					Position: 2,
					Start:    4,
					End:      10,
				},
				{
					Term:     []byte("三四"),
					Type:     Double,
					Position: 3,
					Start:    7,
					End:      13,
				},
				{
					Term:     []byte("五六"),
					Type:     Double,
					Position: 4,
					Start:    14,
					End:      20,
				},
				{
					Term:     []byte("六七"),
					Type:     Double,
					Position: 5,
					Start:    17,
					End:      23,
				},
				{
					Term:     []byte("七八"),
					Type:     Double,
					Position: 6,
					Start:    20,
					End:      26,
				},
				{
					Term:     []byte("八九"),
					Type:     Double,
					Position: 7,
					Start:    23,
					End:      29,
				},
				{
					Term:     []byte("十"),
					Type:     Single,
					Position: 8,
					Start:    30,
					End:      33,
				},
			},
		},
		{
			input: []byte("abc defgh ijklmn opqrstu vwxy z"),
			output: TokenStream{
				{
					Term:     []byte("abc"),
					Type:     Alpha,
					Position: 1,
					Start:    0,
					End:      3,
				},
				{
					Term:     []byte("defgh"),
					Type:     Alpha,
					Position: 2,
					Start:    4,
					End:      9,
				},
				{
					Term:     []byte("ijklmn"),
					Type:     Alpha,
					Position: 3,
					Start:    10,
					End:      16,
				},
				{
					Term:     []byte("opqrstu"),
					Type:     Alpha,
					Position: 4,
					Start:    17,
					End:      24,
				},
				{
					Term:     []byte("vwxy"),
					Type:     Alpha,
					Position: 5,
					Start:    25,
					End:      29,
				},
				{
					Term:     []byte("z"),
					Type:     Alpha,
					Position: 6,
					Start:    30,
					End:      31,
				},
			},
		},
		{
			input: []byte("あい"),
			output: TokenStream{
				{
					Term:     []byte("あい"),
					Type:     Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
			},
		},
		{
			input: []byte("あい   "),
			output: TokenStream{
				{
					Term:     []byte("あい"),
					Type:     Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
			},
		},
		{
			input: []byte("test"),
			output: TokenStream{
				{
					Term:     []byte("test"),
					Type:     Alpha,
					Position: 1,
					Start:    0,
					End:      4,
				},
			},
		},
		{
			input: []byte("test   "),
			output: TokenStream{
				{
					Term:     []byte("test"),
					Type:     Alpha,
					Position: 1,
					Start:    0,
					End:      4,
				},
			},
		},
		{
			input: []byte("あいtest"),
			output: TokenStream{
				{
					Term:     []byte("あい"),
					Type:     Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				{
					Term:     []byte("test"),
					Type:     Alpha,
					Position: 2,
					Start:    6,
					End:      10,
				},
			},
		},
		{
			input: []byte("testあい    "),
			output: TokenStream{
				{
					Term:     []byte("test"),
					Type:     Alpha,
					Position: 1,
					Start:    0,
					End:      4,
				},
				{
					Term:     []byte("あい"),
					Type:     Double,
					Position: 2,
					Start:    4,
					End:      10,
				},
			},
		},
		{
			input: []byte("あいうえおabcかきくけこ"),
			output: TokenStream{
				{
					Term:     []byte("あい"),
					Type:     Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				{
					Term:     []byte("いう"),
					Type:     Double,
					Position: 2,
					Start:    3,
					End:      9,
				},
				{
					Term:     []byte("うえ"),
					Type:     Double,
					Position: 3,
					Start:    6,
					End:      12,
				},
				{
					Term:     []byte("えお"),
					Type:     Double,
					Position: 4,
					Start:    9,
					End:      15,
				},
				{
					Term:     []byte("abc"),
					Type:     Alpha,
					Position: 5,
					Start:    15,
					End:      18,
				},
				{
					Term:     []byte("かき"),
					Type:     Double,
					Position: 6,
					Start:    18,
					End:      24,
				},
				{
					Term:     []byte("きく"),
					Type:     Double,
					Position: 7,
					Start:    21,
					End:      27,
				},
				{
					Term:     []byte("くけ"),
					Type:     Double,
					Position: 8,
					Start:    24,
					End:      30,
				},
				{
					Term:     []byte("けこ"),
					Type:     Double,
					Position: 9,
					Start:    27,
					End:      33,
				},
			},
		},
		{
			input: []byte("あいうえおabんcかきくけ こ"),
			output: TokenStream{
				{
					Term:     []byte("あい"),
					Type:     Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				{
					Term:     []byte("いう"),
					Type:     Double,
					Position: 2,
					Start:    3,
					End:      9,
				},
				{
					Term:     []byte("うえ"),
					Type:     Double,
					Position: 3,
					Start:    6,
					End:      12,
				},
				{
					Term:     []byte("えお"),
					Type:     Double,
					Position: 4,
					Start:    9,
					End:      15,
				},
				{
					Term:     []byte("ab"),
					Type:     Alpha,
					Position: 5,
					Start:    15,
					End:      17,
				},
				{
					Term:     []byte("ん"),
					Type:     Single,
					Position: 6,
					Start:    17,
					End:      20,
				},
				{
					Term:     []byte("c"),
					Type:     Alpha,
					Position: 7,
					Start:    20,
					End:      21,
				},
				{
					Term:     []byte("かき"),
					Type:     Double,
					Position: 8,
					Start:    21,
					End:      27,
				},
				{
					Term:     []byte("きく"),
					Type:     Double,
					Position: 9,
					Start:    24,
					End:      30,
				},
				{
					Term:     []byte("くけ"),
					Type:     Double,
					Position: 10,
					Start:    27,
					End:      33,
				},
				{
					Term:     []byte("こ"),
					Type:     Single,
					Position: 11,
					Start:    34,
					End:      37,
				},
			},
		},
		{
			input: []byte("一 روبرت موير"),
			output: TokenStream{
				{
					Term:     []byte("一"),
					Type:     Single,
					Position: 1,
					Start:    0,
					End:      3,
				},
				{
					Term:     []byte("روبرت"),
					Type:     Alpha,
					Position: 2,
					Start:    4,
					End:      14,
				},
				{
					Term:     []byte("موير"),
					Type:     Alpha,
					Position: 3,
					Start:    15,
					End:      23,
				},
			},
		},
		{
			input: []byte("一 رُوبرت موير"),
			output: TokenStream{
				{
					Term:     []byte("一"),
					Type:     Single,
					Position: 1,
					Start:    0,
					End:      3,
				},
				{
					Term:     []byte("رُوبرت"),
					Type:     Alpha,
					Position: 2,
					Start:    4,
					End:      16,
				},
				{
					Term:     []byte("موير"),
					Type:     Alpha,
					Position: 3,
					Start:    17,
					End:      25,
				},
			},
		},
		{
			input: []byte("𩬅艱鍟䇹愯瀛"),
			output: TokenStream{
				{
					Term:     []byte("𩬅艱"),
					Type:     Double,
					Position: 1,
					Start:    0,
					End:      7,
				},
				{
					Term:     []byte("艱鍟"),
					Type:     Double,
					Position: 2,
					Start:    4,
					End:      10,
				},
				{
					Term:     []byte("鍟䇹"),
					Type:     Double,
					Position: 3,
					Start:    7,
					End:      13,
				},
				{
					Term:     []byte("䇹愯"),
					Type:     Double,
					Position: 4,
					Start:    10,
					End:      16,
				},
				{
					Term:     []byte("愯瀛"),
					Type:     Double,
					Position: 5,
					Start:    13,
					End:      19,
				},
			},
		},
		{
			input: []byte("一"),
			output: TokenStream{
				{
					Term:     []byte("一"),
					Type:     Single,
					Position: 1,
					Start:    0,
					End:      3,
				},
			},
		},
		{
			input: []byte("一丁丂"),
			output: TokenStream{
				{
					Term:     []byte("一丁"),
					Type:     Double,
					Position: 1,
					Start:    0,
					End:      6,
				},
				{
					Term:     []byte("丁丂"),
					Type:     Double,
					Position: 2,
					Start:    3,
					End:      9,
				},
			},
		},
	}

	analyzer := NewAnalyzer()
	for _, test := range tests {
		actual := analyzer.Analyze(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf("expected for %s:\n%#+v\ngot:\n%#+v", string(test.input), test.output, actual)
		}
	}
}

func BenchmarkCJKAnalyzer(b *testing.B) {
	analyzer := NewAnalyzer()
	for i := 0; i < b.N; i++ {
		analyzer.Analyze(bleveWikiArticleJapanese)
	}
}

var bleveWikiArticleJapanese = []byte(`加圧容器に貯蔵されている液体物質は、その時の気液平衡状態にあるが、火災により容器が加熱されていると容器内の液体は、その物質の大気圧のもとでの沸点より十分に高い温度まで加熱され、圧力も高くなる。この状態で容器が破裂すると容器内部の圧力は瞬間的に大気圧にまで低下する。
この時に容器内の平衡状態が破られ、液体は突沸し、気体になることで爆発現象を起こす。液化石油ガスなどでは、さらに拡散して空気と混ざったガスが自由空間蒸気雲爆発を起こす。液化石油ガスなどの常温常圧で気体になる物を高い圧力で液化して収納している容器、あるいは、そのような液体を輸送するためのパイプラインや配管などが火災などによって破壊されたときに起きる。
ブリーブという現象が明らかになったのは、フランス・リヨンの郊外にあるフェザンという町のフェザン製油所（ウニオン・ド・ゼネラル・ド・ペトロール）で大規模な爆発火災事故が発生したときだと言われている。
中身の液体が高温高圧の水である場合には「水蒸気爆発」と呼ばれる。`)
