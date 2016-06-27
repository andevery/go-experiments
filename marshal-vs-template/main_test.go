package main

import (
	"testing"
)

// go test --bench=.
// PASS
// BenchmarkTemplateRender1-4            300000        5241 ns/op
// BenchmarkTemplateRender10-4            30000       39102 ns/op
// BenchmarkTemplateRender100-4            5000      367930 ns/op
// BenchmarkTemplateRender1000-4            500     3608926 ns/op
// BenchmarkMarshalRender1-4            1000000        1143 ns/op
// BenchmarkMarshalRender10-4            300000        5046 ns/op
// BenchmarkMarshalRender100-4            30000       44743 ns/op
// BenchmarkMarshalRender1000-4            3000      442827 ns/op
// BenchmarkQuickTemplateRender1-4      2000000         677 ns/op
// BenchmarkQuickTemplateRender10-4      500000        3697 ns/op
// BenchmarkQuickTemplateRender100-4      50000       34420 ns/op
// BenchmarkQuickTemplateRender1000-4      5000      335233 ns/op
// ok    github.com/andevery/go-experiments/marshal-vs-template  20.908s
//
// sysctl -n machdep.cpu.brand_string
// Intel(R) Core(TM) i5-4260U CPU @ 1.40GHz

func TestEqualRenderedData(t *testing.T) {
	a := NewApp(2)
	mData := a.MarshalRender()
	tData := a.TemplateRender()
	qData := a.TemplateRender()
	if string(mData) == string(tData) && string(mData) == string(qData) {
		return
	}
	t.Errorf("Different data")
}

func render(b *testing.B, t string, i int) {
	a := NewApp(i)
	var f func() []byte
	switch t {
	case "m":
		f = a.MarshalRender
	case "t":
		f = a.TemplateRender
	case "q":
		f = a.QuickTemplateRender
	}
	for n := 0; n < b.N; n++ {
		_ = f()
	}
}

func BenchmarkTemplateRender1(b *testing.B) {
	render(b, "t", 1)
}

func BenchmarkTemplateRender10(b *testing.B) {
	render(b, "t", 10)
}

func BenchmarkTemplateRender100(b *testing.B) {
	render(b, "t", 100)
}

func BenchmarkTemplateRender1000(b *testing.B) {
	render(b, "t", 1000)
}

func BenchmarkMarshalRender1(b *testing.B) {
	render(b, "m", 1)
}

func BenchmarkMarshalRender10(b *testing.B) {
	render(b, "m", 10)
}

func BenchmarkMarshalRender100(b *testing.B) {
	render(b, "m", 100)
}

func BenchmarkMarshalRender1000(b *testing.B) {
	render(b, "m", 1000)
}

func BenchmarkQuickTemplateRender1(b *testing.B) {
	render(b, "q", 1)
}

func BenchmarkQuickTemplateRender10(b *testing.B) {
	render(b, "q", 10)
}

func BenchmarkQuickTemplateRender100(b *testing.B) {
	render(b, "q", 100)
}

func BenchmarkQuickTemplateRender1000(b *testing.B) {
	render(b, "q", 1000)
}
