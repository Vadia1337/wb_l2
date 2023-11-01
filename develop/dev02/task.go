package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var IncorrectStringError = errors.New("expected incorrect string")

func main() {
	str := "a4bc2d5e"

	outStr, err := UnpackingString(str)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(outStr)

}

func UnpackingString(inputStr string) (string, error) {

	outStr := make([]byte, 0, len(inputStr))

	lastRune := ""
	for _, v := range inputStr {
		str := string(v)
		strRepeats, err := strconv.Atoi(str) // тут много аллокаций ((
		if err == nil {

			if lastRune == "" {
				return "", IncorrectStringError
			}

			for i := 0; i < strRepeats-1; i++ {
				outStr = append(outStr, lastRune...)
			}

			continue
		}

		lastRune = str
		outStr = append(outStr, lastRune...)
	}

	return string(outStr), nil
}

/****************************************************************************************************************/

// реализовал для сравнения по времени и памяти
func unpackingString2(inputStr string) (string, error) {

	var buffer bytes.Buffer
	lastRune := ""
	for _, v := range inputStr {
		str := string(v)

		strRepeats, err := strconv.Atoi(str)
		if err == nil {

			if lastRune == "" {
				return "", IncorrectStringError
			}

			for i := 0; i < strRepeats-1; i++ {
				buffer.WriteString(lastRune)
			}

			continue
		}

		lastRune = str
		buffer.WriteString(lastRune)
	}

	return buffer.String(), nil
}
