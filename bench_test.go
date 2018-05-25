package bench

import (
	"fmt"
	"testing"
)

// BenchmarkSliceAddressing compares the hypothetical performance of an
// image.Gray14 using []uint16 instead of the current []byte.
func BenchmarkSliceAddressing(b *testing.B) {
	b.Run("Uint16", func(b *testing.B) {
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
	})
	b.Run("Uint8Pair", func(b *testing.B) {
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
	})
}

// BenchmarkFunction measures the different costs for function calls.
func BenchmarkFunction(b *testing.B) {
	b.Run("Normal", func(b *testing.B) {
		b.ReportAllocs()
		t := 0
		for i := 0; i < b.N; i++ {
			t = multiple2(i)
		}
		fmt.Sprintln(t)
	})
	b.Run("NormalNotInline", func(b *testing.B) {
		b.ReportAllocs()
		t := 0
		for i := 0; i < b.N; i++ {
			t = multiple2NoInline(i)
		}
		fmt.Sprintln(t)
	})
	b.Run("ClosureArgument", func(b *testing.B) {
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
	})
	b.Run("Closure", func(b *testing.B) {
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
	})
	b.Run("FunctionPointer", func(b *testing.B) {
		b.ReportAllocs()
		t := 0
		f := getFunc()
		for i := 0; i < b.N; i++ {
			t = f(i)
		}
		fmt.Sprintln(t)
	})
	b.Run("FunctionPointerRepeated", func(b *testing.B) {
		b.ReportAllocs()
		t := 0
		for i := 0; i < b.N; i++ {
			f := getFunc()
			t = f(i)
		}
		fmt.Sprintln(t)
	})
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

// BenchmarkDivision measures the division performance of various native types.
// This is because ARM instruction set doesn't declare integer division and
// Xeon has surprising performance characteristics.
func BenchmarkDivision(b *testing.B) {
	b.Run("Int32", func(b *testing.B) {
		t := int32(1<<31 - 1)
		d := int32(b.N)
		for i := 0; i < b.N; i++ {
			t /= d
		}
		fmt.Sprintln(t)
	})
	b.Run("Int64", func(b *testing.B) {
		t := int64(1<<63 - 1)
		d := int64(b.N)
		for i := 0; i < b.N; i++ {
			t /= d
		}
		fmt.Sprintln(t)
	})
	b.Run("Float32", func(b *testing.B) {
		t := float32(1e38)
		d := float32(b.N)
		for i := 0; i < b.N; i++ {
			t /= d
		}
		fmt.Sprintln(t)
	})
	b.Run("Float64", func(b *testing.B) {
		t := float64(1e308)
		d := float64(b.N)
		for i := 0; i < b.N; i++ {
			t /= d
		}
		fmt.Sprintln(t)
	})
	// This can't be faster than an int64 division, right? Right?
	b.Run("Int64ConvertedtoFloat64", func(b *testing.B) {
		t := int64(1<<63 - 1)
		d := int64(b.N)
		for i := 0; i < b.N; i++ {
			t = int64(float64(t) / float64(d))
		}
		fmt.Sprintln(t)
	})
}

func BenchmarkShift(b *testing.B) {
	b.Run("Int32", func(b *testing.B) {
		t := int32(1<<31 - 1)
		for i := 0; i < b.N; i++ {
			t >>= 1
		}
		fmt.Sprintln(t)
	})
	b.Run("Int64", func(b *testing.B) {
		t := int64(1<<63 - 1)
		for i := 0; i < b.N; i++ {
			t >>= 1
		}
		fmt.Sprintln(t)
	})
}
