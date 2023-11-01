package main

import (
	"strings"
	"testing"
)

func Test_unpackingString(t *testing.T) {

	TestDataset := []struct {
		str        string
		expectStr  string
		testNumber int
		err        error
	}{
		{
			"a4b4c7f", "aaaabbbbcccccccf", 0, nil,
		},
		{
			"Я4я4Ы7", "ЯЯЯЯяяяяыыыыыыы", 1, nil,
		},
		{
			"ва3дю2с4ч", "вааадююссссч", 2, nil,
		},
		{
			"abcd", "abcd", 3, nil,
		},
		{
			"ccddffmm", "ccddffmm", 4, nil,
		},
		{
			"a9f9g9e9", "aaaaaaaaafffffffffgggggggggeeeeeeeee", 5, nil,
		},
		{
			"45", "", 6, IncorrectStringError,
		},
		{
			"4b5c", "", 7, IncorrectStringError,
		},
	}

	for _, v := range TestDataset {
		outStr, err := UnpackingString(v.str)
		if err != v.err {
			t.Error("Неожиданная ошибка, Ожидали :", v.err, "Получили:", err)
		}

		if !strings.EqualFold(outStr, v.expectStr) {
			t.Error("Функция выдала некорректный результат, ожидалось: ", v.expectStr, "Получили: ", outStr,
				"Ошибка произошла на наборе данных №", v.testNumber)
		}
	}

}

func BenchmarkUnpackingString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnpackingString("s2b4k8")
	}
}

func BenchmarkUnpackingString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unpackingString2("s2b4k8")
	}
}
