package opencc

import (
	"testing"

	"github.com/griffinqiu/opencc"
)

func BenchmarkConvert_s2t(b *testing.B) {
	cc, _ := opencc.New("s2t")

	in := readFile("html-raw.txt")

	// 2.5 ms/op in Apple M1
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}

func BenchmarkConvert_t2s(b *testing.B) {
	cc, _ := opencc.New("t2s")

	in := readFile("html-s2t.txt")

	// 2.3 ms/op in Apple M1
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}

func BenchmarkConvert_s2tw(b *testing.B) {
	cc, _ := opencc.New("s2tw")

	in := readFile("html-raw.txt")

	// 3.8 ms/op in Apple M1
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}

func BenchmarkConvert_s2hk_finance(b *testing.B) {
	cc, _ := opencc.New("s2hk-finance")

	in := readFile("html-raw.txt")

	// 5 ms/op in Apple M1
	for n := 0; n < b.N; n++ {
		cc.Convert(in)
	}
}
