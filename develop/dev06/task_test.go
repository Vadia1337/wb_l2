package main

import (
	"slices"
	"strconv"
	"testing"
)

func TestCut(t *testing.T) {

	TestDataSet := []struct {
		TestNumb  int
		Fields    string
		Delimiter string
		Separated bool
		InText    string
		OutText   []string
		OutErr    error
	}{
		{
			TestNumb:  1,
			Fields:    "",
			Delimiter: "\t",
			Separated: false,
			InText:    "Валерий:001\tВиктор:002",
			OutText:   []string{},
			OutErr:    NotSetFields,
		},
		{
			TestNumb:  2,
			Fields:    "s",
			Delimiter: "\t",
			Separated: false,
			InText:    "Валерий:001\tВиктор:002",
			OutText:   []string{},
			OutErr:    IncorrectFields,
		},
		{
			TestNumb:  3,
			Fields:    "1-5",
			Delimiter: "\t",
			Separated: false,
			InText:    "Валерий:001\tВиктор:002",
			OutText:   []string{},
			OutErr:    OutOfRange,
		},
		{
			TestNumb:  4,
			Fields:    "1-3",
			Delimiter: ":",
			Separated: false,
			InText:    "Валерий:001\tВиктор:002",
			OutText: []string{
				"Валерий",
				"001\tВиктор",
				"002",
			},
			OutErr: nil,
		},
		{
			TestNumb:  5,
			Fields:    "1,3",
			Delimiter: ":",
			Separated: false,
			InText:    "Валерий:001\tВиктор:002",
			OutText: []string{
				"Валерий",
				"002",
			},
			OutErr: nil,
		},
		{
			TestNumb:  6,
			Fields:    "3",
			Delimiter: ":",
			Separated: false,
			InText:    "Валерий:001\tВиктор:002",
			OutText: []string{
				"002",
			},
			OutErr: nil,
		},
		{
			TestNumb:  7,
			Fields:    "1",
			Delimiter: "*",
			Separated: false,
			InText:    "Валерий:001\tВиктор:002",
			OutText: []string{
				"Валерий:001\tВиктор:002",
			},
			OutErr: nil,
		},
		{
			TestNumb:  8,
			Fields:    "1",
			Delimiter: "*",
			Separated: true,
			InText:    "Валерий:001\tВиктор:002",
			OutText:   []string{},
			OutErr:    nil,
		},
	}

	for _, v := range TestDataSet {
		t.Run(strconv.Itoa(v.TestNumb), func(t *testing.T) {})

		Fields = v.Fields
		Delimiter = v.Delimiter
		Separated = v.Separated

		OutText, err := Cut(v.InText)
		if err != v.OutErr {
			t.Error("Ошибки не совпали, ожидалось: ", v.OutErr, "получили: ", err, "набор данных №", v.TestNumb)
			continue
		}

		if len(OutText) != len(v.OutText) {
			t.Error("размеры слайсов, не совпали, ожидалось: ", v.OutText, "размер: ", len(v.OutText),
				"получили: ", OutText, "размер: ", len(OutText), "набор данных №", v.TestNumb)
		}

		if !slices.Equal(OutText, v.OutText) {
			t.Error("Слайсы, не совпали, ожидалось: ", v.OutText, "получили: ", OutText,
				"набор данных №", v.TestNumb)
		}
	}
}
