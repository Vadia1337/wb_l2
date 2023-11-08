package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"testing"
)

func Test_Sort(t *testing.T) {

	testDataSet := []struct {
		TestNum          int
		SelectColumnFlag int
		SortByInt        bool
		ReverseSort      bool
		UniqueValues     bool
		FileName         string
		ReferenceFile    string
		OutErr           error
	}{
		{
			TestNum:          1,
			SelectColumnFlag: 1,
			SortByInt:        true,
			ReverseSort:      true,
			UniqueValues:     true,
			FileName:         "text.txt",
			ReferenceFile:    "refFiles/OutTest1.txt",
			OutErr:           nil,
		},
		{
			TestNum:          2,
			SelectColumnFlag: 2,
			SortByInt:        true,
			ReverseSort:      true,
			UniqueValues:     true,
			FileName:         "text.txt",
			ReferenceFile:    "refFiles/OutTest1.txt",
			OutErr:           nil,
		},
		{
			TestNum:          3,
			SelectColumnFlag: 3,
			SortByInt:        true,
			ReverseSort:      true,
			UniqueValues:     true,
			FileName:         "text.txt",
			ReferenceFile:    "refFiles/OutTest1.txt",
			OutErr:           SortStrByNFlagError,
		},
		{
			TestNum:          4,
			SelectColumnFlag: 4,
			SortByInt:        true,
			ReverseSort:      true,
			UniqueValues:     true,
			FileName:         "text.txt",
			ReferenceFile:    "refFiles/OutTest1.txt",
			OutErr:           SortStrByNFlagError,
		},
		{
			TestNum:          5,
			SelectColumnFlag: 3,
			SortByInt:        false,
			ReverseSort:      true,
			UniqueValues:     true,
			FileName:         "text.txt",
			ReferenceFile:    "refFiles/OutTest2.txt",
			OutErr:           nil,
		},
		{
			TestNum:          6,
			SelectColumnFlag: 4,
			SortByInt:        false,
			ReverseSort:      true,
			UniqueValues:     true,
			FileName:         "text.txt",
			ReferenceFile:    "refFiles/OutTest3.txt",
			OutErr:           nil,
		},
		{
			TestNum:          7,
			SelectColumnFlag: 3,
			SortByInt:        false,
			ReverseSort:      false,
			UniqueValues:     false,
			FileName:         "text.txt",
			ReferenceFile:    "refFiles/OutTest4.txt",
			OutErr:           nil,
		},
		{
			TestNum:          8,
			SelectColumnFlag: 2,
			SortByInt:        false,
			ReverseSort:      false,
			UniqueValues:     false,
			FileName:         "text.txt",
			ReferenceFile:    "refFiles/OutTest5.txt",
			OutErr:           nil,
		},
		{
			TestNum:          9,
			SelectColumnFlag: 1,
			SortByInt:        false,
			ReverseSort:      false,
			UniqueValues:     false,
			FileName:         "text.txt",
			ReferenceFile:    "refFiles/OutTest5.txt",
			OutErr:           nil,
		},
		{
			TestNum:          10,
			SelectColumnFlag: 1,
			SortByInt:        false,
			ReverseSort:      false,
			UniqueValues:     false,
			FileName:         "",
			ReferenceFile:    "refFiles/OutTest5.txt",
			OutErr:           CantFindFileError,
		},
	}

	for _, v := range testDataSet {
		SelectColumnFlag = v.SelectColumnFlag
		SortByInt = v.SortByInt
		ReverseSort = v.ReverseSort
		UniqueValues = v.UniqueValues

		err := Sort(v.FileName)
		if err != v.OutErr {
			t.Error("Ожидалось, что ф-я вернет ошибку:", v.OutErr, "Но вернула:", err,
				"тест проводился на наборе №", v.TestNum)
		}

		if FileMD5("out.txt") != FileMD5(v.ReferenceFile) {
			t.Error("Сортировка создала файл, несоответствующий эталонному файлу, тест проводился на наборе №",
				v.TestNum, "Хэши файлов не совпали")
		}
	}

}

// FileMD5 создает md5-хеш из содержимого нашего файла.
func FileMD5(path string) string {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func BenchmarkSort(b *testing.B) {
	SelectColumnFlag = 2
	for n := 0; n < b.N; n++ {
		SortOld("text.txt")
	}
}

func BenchmarkSortSmallFile(b *testing.B) {
	SelectColumnFlag = 2
	for n := 0; n < b.N; n++ {
		Sort("text.txt")
	}
}

func BenchmarkSortIntSmallFile(b *testing.B) {
	SelectColumnFlag = 1
	SortByInt = true
	for n := 0; n < b.N; n++ {
		Sort("text.txt")
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
