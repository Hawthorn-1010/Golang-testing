package lib

import (
	"testing"
	"unicode/utf8"
)

type AddArray struct {
	result int
	num1   int
	num2   int
}

// 功能测试，判断是否正确
func TestAdd(t *testing.T) {
	testData := [3]AddArray{
		{2, 1, 1},
		{5, 2, 3},
		{4, 2, 2},
	}
	for _, v := range testData {
		if v.result != Add(v.num1, v.num2) {
			t.Errorf("Add (%d, %d) != %d \n", v.num1, v.num2, v.result)
		}
	}
}

func BenchmarkSingleAdd(b *testing.B) {
	// 运行性能测试时，会根据测试函数的运行情况自动调整 b.N 的值
	// 让测试函数的总运行时间保持在一个合理的范围内，以便更好地评估程序的性能
	for i := 0; i < b.N; i++ {
		num1 := 100
		num2 := 200
		Add(num1, num2)
	}
}

func BenchmarkParallelAdd(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			num1 := 100
			num2 := 200
			Add(num1, num2)
		}
	})
}

func FuzzReverse(f *testing.F) {
	testcases := []string{
		"Hello World",
		" ",
		"12@3234f",
	}
	for _, s := range testcases {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %s, after: %s", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Invalid string %s", rev)
		}
	})
}
