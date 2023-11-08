package main

import (
	"slices"
	"strconv"
	"testing"
)

func TestGrepStr(t *testing.T) {
	TestDataSet := []struct {
		TestNumb       int
		PrintAfter     int
		PrintBefore    int
		PrintContext   int
		PrintLineCount bool
		IgnoreCase     bool
		Invert         bool
		FixedStr       bool
		PrintLineNumb  bool
		Pattern        string
		fileName       []string
		OutText        []string
	}{
		{
			TestNumb:       1,
			PrintAfter:     0,
			PrintBefore:    0,
			PrintContext:   0,
			PrintLineCount: false,
			IgnoreCase:     false,
			Invert:         false,
			FixedStr:       false,
			PrintLineNumb:  false,
			Pattern:        "тебе",
			fileName:       []string{"text.txt"},
			OutText: []string{
				"text.txt: ",
				"Бывало, милые тебе —",
				"--",
			},
		},
		{
			TestNumb:       2,
			PrintAfter:     0,
			PrintBefore:    0,
			PrintContext:   0,
			PrintLineCount: true,
			IgnoreCase:     false,
			Invert:         false,
			FixedStr:       false,
			PrintLineNumb:  false,
			Pattern:        "тебе",
			fileName:       []string{"text.txt"},
			OutText: []string{
				"1",
			},
		},
		{
			TestNumb:       3,
			PrintAfter:     0,
			PrintBefore:    0,
			PrintContext:   0,
			PrintLineCount: false,
			IgnoreCase:     true,
			Invert:         false,
			FixedStr:       false,
			PrintLineNumb:  true,
			Pattern:        "тебе",
			fileName:       []string{"text.txt"},
			OutText: []string{
				"text.txt: ",
				"0: Тебе — но голос музы тёмной",
				"9: Бывало, милые тебе —",
				"--",
			},
		},
		{
			TestNumb:       4,
			PrintAfter:     0,
			PrintBefore:    0,
			PrintContext:   0,
			PrintLineCount: false,
			IgnoreCase:     false,
			Invert:         true,
			FixedStr:       true,
			PrintLineNumb:  true,
			Pattern:        "Тебе — но голос музы тёмной",
			fileName:       []string{"text.txt"},
			OutText: []string{
				"text.txt: ",
				"1: Коснется ль уха твоего?",
				"2: Поймешь ли ты душою скромной",
				"3: Стремленье сердца моего?",
				"4: Иль посвящение поэта,",
				"5: Как некогда его любовь,",
				"6: Перед тобою без ответа",
				"7: Пройдет, непризнанное вновь?",
				"8: Узнай, по крайней мере, звуки,",
				"9: Бывало, милые тебе —",
				"10: И думай, что во дни разлуки,",
				"11: В моей изменчивой судьбе,",
				"--",
			},
		},
		{
			TestNumb:       5,
			PrintAfter:     2,
			PrintBefore:    2,
			PrintContext:   0,
			PrintLineCount: false,
			IgnoreCase:     true,
			Invert:         false,
			FixedStr:       false,
			PrintLineNumb:  true,
			Pattern:        "Тебе",
			fileName:       []string{"text.txt"},
			OutText: []string{
				"text.txt: ",
				"0: Тебе — но голос музы тёмной",
				"Коснется ль уха твоего?",
				"Поймешь ли ты душою скромной",
				"Пройдет, непризнанное вновь?",
				"Узнай, по крайней мере, звуки,",
				"9: Бывало, милые тебе —",
				"И думай, что во дни разлуки,",
				"В моей изменчивой судьбе,",
				"--",
			},
		},
		{
			TestNumb:       6,
			PrintAfter:     0,
			PrintBefore:    0,
			PrintContext:   2,
			PrintLineCount: false,
			IgnoreCase:     false,
			Invert:         false,
			FixedStr:       false,
			PrintLineNumb:  false,
			Pattern:        "тебе",
			fileName:       []string{"text.txt"},
			OutText: []string{
				"text.txt: ",
				"Пройдет, непризнанное вновь?",
				"Узнай, по крайней мере, звуки,",
				"Бывало, милые тебе —",
				"И думай, что во дни разлуки,",
				"В моей изменчивой судьбе,",
				"--",
			},
		},
		{
			TestNumb:       7,
			PrintAfter:     0,
			PrintBefore:    0,
			PrintContext:   2,
			PrintLineCount: true,
			IgnoreCase:     false,
			Invert:         false,
			FixedStr:       false,
			PrintLineNumb:  false,
			Pattern:        "тебе",
			fileName:       []string{"text.txt"},
			OutText: []string{
				"5",
			},
		},
		{
			TestNumb:       8,
			PrintAfter:     0,
			PrintBefore:    0,
			PrintContext:   0,
			PrintLineCount: false,
			IgnoreCase:     true,
			Invert:         false,
			FixedStr:       false,
			PrintLineNumb:  true,
			Pattern:        "тебе",
			fileName:       []string{"text.txt", "text1.txt"},
			OutText: []string{
				"text.txt: ",
				"0: Тебе — но голос музы тёмной",
				"9: Бывало, милые тебе —",
				"--",
				"text1.txt: ",
				"0: Тебе — но голос музы тёмной",
				"9: Бывало, милые тебе —",
				"--",
			},
		},
	}

	for _, v := range TestDataSet {
		t.Run(strconv.Itoa(v.TestNumb), func(t *testing.T) {})

		PrintAfter = v.PrintAfter
		PrintBefore = v.PrintBefore
		PrintContext = v.PrintContext
		PrintLineCount = v.PrintLineCount
		IgnoreCase = v.IgnoreCase
		Invert = v.Invert
		FixedStr = v.FixedStr
		PrintLineNumb = v.PrintLineNumb

		str, err := GrepStr(v.Pattern, v.fileName)
		if err != nil {
			return
		}

		if len(str) != len(v.OutText) {
			t.Error("размеры слайсов, не совпали, ожидалось: ", v.OutText, "размер: ", len(v.OutText),
				"получили: ", str, "размер: ", len(str), "набор данных №", v.TestNumb)
		}

		if !slices.Equal(str, v.OutText) {
			t.Error("Слайсы, не совпали, ожидалось: ", v.OutText, "получили: ", str,
				"набор данных №", v.TestNumb)
		}
	}
}

func BenchmarkGrep(b *testing.B) {
	PrintAfter = 0
	PrintBefore = 0
	PrintContext = 0
	PrintLineCount = false
	IgnoreCase = false
	Invert = false
	FixedStr = false
	PrintLineNumb = false

	for i := 0; i < b.N; i++ {
		Grep("и", []string{"text.txt"})
	}
}

func BenchmarkGrepStr(b *testing.B) {
	PrintAfter = 0
	PrintBefore = 0
	PrintContext = 0
	PrintLineCount = false
	IgnoreCase = false
	Invert = false
	FixedStr = false
	PrintLineNumb = true

	for i := 0; i < b.N; i++ {
		GrepStr("и", []string{"text.txt"})
	}
}
