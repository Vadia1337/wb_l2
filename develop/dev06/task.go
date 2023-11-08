package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	Fields    string
	Delimiter string
	Separated bool

	NotSetFields    = errors.New("вы не выбрали колонки, которые необходимо вывести. флаг -f")
	IncorrectFields = errors.New("не корректно указаны параметры флага -f, корректно будет указать: '-f 1' '-f 1,2,3' '-f 1-5'")
	OutOfRange      = errors.New("колонки с таким номером(ами) не существует, попробуйте поменять значение флага -f")
)

func init() {
	flag.StringVar(&Fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&Delimiter, "d", "\t", "использовать другой разделитель (дефолтно: TAB)")
	flag.BoolVar(&Separated, "s", false, "только строки с разделителем")
}

func main() {
	flag.Parse()

	buf := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строку: ")
	for buf.Scan() { // ввожу строки последовательно, тк не сыскал способа ввести сразу много строк через консоль
		text := buf.Text()

		outStr, err := Cut(text)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}

		fmt.Println(outStr)

		fmt.Println("Введите строку или ctrl+c для выхода из утилиты: ")
	}
}

func Cut(text string) ([]string, error) {

	if Separated && !strings.Contains(text, Delimiter) { // если в строке нет Delimiter-a то просто пропускаем строку
		return nil, nil
	}

	DelimitedText := strings.Split(text, Delimiter)

	if Fields == "" {
		return nil, NotSetFields
	}

	outString := make([]string, 0, 2)

	switch {
	case strings.Contains(Fields, "-"): // если ввели диапазон ex: 1-5
		columns := strings.Split(Fields, "-")
		start, _ := strconv.Atoi(columns[0])
		end, _ := strconv.Atoi(columns[1])

		if checkOutOfRange(DelimitedText, end-1) {
			return nil, OutOfRange
		}

		outString = append(outString, DelimitedText[start-1:end]...)
	case strings.Contains(Fields, ","): // если ввели через , ex: 1,2,3
		columns := strings.Split(Fields, ",")
		for _, v := range columns {
			column, _ := strconv.Atoi(v)

			if checkOutOfRange(DelimitedText, column-1) {
				return nil, OutOfRange
			}

			outString = append(outString, DelimitedText[column-1])
		}
	default: // ввели одну цифру
		column, err := strconv.Atoi(Fields)
		if err != nil {
			return nil, IncorrectFields
		}

		if checkOutOfRange(DelimitedText, column-1) {
			return nil, OutOfRange
		}

		outString = append(outString, DelimitedText[column-1])
	}

	return outString, nil
}

func checkOutOfRange(slice []string, index int) bool {
	return !(len(slice) > index)
}
