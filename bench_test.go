package bench

import (
	"fmt"
	"testing"
)

// These two benchmarks are to compare the hypothetical performance of an
// image.Gray14 using []uint16 instead of the current []byte.

func BenchmarkSliceUint16(b *testing.B) {
	pix := make([]uint16, b.N)
	b.ResetTimer()
	n := uint16(0)
	for i := 0; i < b.N; i++ {
		v := pix[i]
		n += v
	}
	if n != 0 {
		b.Fatal("Oops")
	}
}

func BenchmarkSliceUint8Pairs(b *testing.B) {
	pix := make([]uint8, b.N*2)
	b.ResetTimer()
	n := uint16(0)
	for i := 0; i < b.N; i++ {
		v := uint16(2*pix[i+0])<<8 | uint16(pix[2*i+1])
		n += v
	}
	if n != 0 {
		b.Fatal("Oops")
	}
}

// These 3 benchmarks test the different costs for function calls.

func BenchmarkAnonymousFunction(b *testing.B) {
	b.ReportAllocs()
	t := 0
	for i := 0; i < b.N; i++ {
		f := func(j int) (r int) {
			r = j * 2
			return
		}
		t = f(i)
	}
	fmt.Sprintln(t)
}

func BenchmarkClosure(b *testing.B) {
	b.ReportAllocs()
	t := 0
	for i := 0; i < b.N; i++ {
		f := func() (r int) {
			r = i * 2
			return
		}
		t = f()
	}
	fmt.Sprintln(t)
}

func BenchmarkNormalFunction(b *testing.B) {
	b.ReportAllocs()
	t := 0
	for i := 0; i < b.N; i++ {
		t = multiple2(i)
	}
	fmt.Sprintln(t)
}

func BenchmarkNormalFunctionNotInline(b *testing.B) {
	b.ReportAllocs()
	t := 0
	for i := 0; i < b.N; i++ {
		t = multiple2NoInline(i)
	}
	fmt.Sprintln(t)
}

func BenchmarkReturnedFunction(b *testing.B) {
	b.ReportAllocs()
	t := 0
	f := getFunc()
	for i := 0; i < b.N; i++ {
		t = f(i)
	}
	fmt.Sprintln(t)
}

func BenchmarkReturnedFunctionRepeated(b *testing.B) {
	b.ReportAllocs()
	t := 0
	for i := 0; i < b.N; i++ {
		f := getFunc()
		t = f(i)
	}
	fmt.Sprintln(t)
}

func multiple2(i int) int {
	return i * 2
}

//go:noinline
func multiple2NoInline(i int) int {
	return i * 2
}

func getFunc() func(int) int {
	return func(i int) int {
		return i * 2
	}
}

// These benchmarks tests the int64 and float64 division performance. This is
// because ARM instruction set doesn't declare integer division.

func BenchmarkDivisionInt64(b *testing.B) {
	t := int64(12345678901234)
	d := int64(b.N)
	for i := 0; i < b.N; i++ {
		t /= d
	}
	fmt.Sprintln(t)
}

func BenchmarkDivisionFloat64(b *testing.B) {
	t := float64(1e100)
	d := float64(b.N)
	for i := 0; i < b.N; i++ {
		t /= d
	}
	fmt.Sprintln(t)
}
