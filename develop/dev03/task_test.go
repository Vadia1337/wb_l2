package main

import (
	"bytes"
	"testing"
)

func Test_Sort(t *testing.T) {
	SelectColumnFlag = 2
	Sort("text.txt")
}

func BenchmarkSort(b *testing.B) {
	SelectColumnFlag = 2
	for n := 0; n < b.N; n++ {
		Sort("text.txt")
	}
}

func BenchmarkSortSmallFile(b *testing.B) {
	SelectColumnFlag = 2
	for n := 0; n < b.N; n++ {
		SortSmallFile("text.txt")
	}
}

func BenchmarkSortIntSmallFile(b *testing.B) {
	SelectColumnFlag = 1
	SortByInt = true
	for n := 0; n < b.N; n++ {
		SortSmallFile("text.txt")
	}
}

//тестировал способы конкатенации
/***********************************************/
const testString = "test"

func BenchmarkConcat(b *testing.B) {
	b.ResetTimer()
	var str string
	for n := 0; n < b.N; n++ {
		str += testString
	}
	b.StopTimer()
}

func BenchmarkBufferConcat(b *testing.B) {
	b.ResetTimer()
	var buffer bytes.Buffer

	for n := 0; n < b.N; n++ {
		buffer.WriteString(testString)
	}
	_ = buffer.String()
	b.StopTimer()
}

func BenchmarkCopyConcat(b *testing.B) {
	b.ResetTimer()

	bs := make([]byte, b.N)
	bl := 0

	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], testString)
	}

	_ = string(bs)
	b.StopTimer()
}
