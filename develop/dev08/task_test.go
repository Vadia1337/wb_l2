package main

import (
	"fmt"
	"testing"
)

func TestShell(t *testing.T) {
	testDataSet := []struct {
		TestNumb  int
		InputLine string
		Output    string
	}{
		{
			TestNumb:  1,
			InputLine: "echo ss-ll",
			Output:    "ss-ll",
		},
		{
			TestNumb:  2,
			InputLine: "echo ss-ll | grep ss",
			Output:    "ss-ll",
		},
		{
			TestNumb:  3,
			InputLine: "echo ss-ll | grep gg",
			Output:    "",
		},
		{
			TestNumb:  4,
			InputLine: "echo ss ll ff | grep ss",
			Output:    "ss ll ff",
		},
	}
	// было лень делать буффер для тестирования, куда бы писал stdout )))
	// + не придумал как & оттетсировать (только с time sleep)
	fmt.Println("Проверка в полуручном режиме")
	for _, v := range testDataSet {
		fmt.Println("Тест №", v.TestNumb, "Ожидается: ", v.Output, "Получили: ")
		Shell(v.InputLine)
	}
}
