package main

import (
	"os"
	"testing"
)

func TestWget(t *testing.T) {
	TestDataSet := []struct {
		TestNumb int
		Addr     string
		OutErr   error
	}{
		{
			TestNumb: 1,
			Addr:     "",
			OutErr:   AddrNotSet,
		},
		{
			TestNumb: 2,
			Addr:     "edcjs",
			OutErr:   GetError,
		},
		{
			TestNumb: 3,
			Addr:     "https://ru.wikipedia.org/wiki/Go",
			OutErr:   nil,
		},
	}

	for _, v := range TestDataSet {

		os.Remove("index.html")

		err := wget(v.Addr)
		if err != v.OutErr {
			t.Error("Ошибки не совпали, ожидалось: ", v.OutErr, "получили: ", err, "набор данных №", v.TestNumb)
			continue
		}

		if err != nil { // пропускаем, дабы не доходить до теста файлов
			continue
		}

		_, err = os.Open("index.html")
		if err != nil {
			t.Error("С файлом что-то не то... набор данных №", v.TestNumb)
		}

	}
}
