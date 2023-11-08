package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var (
	PrintAfter     int
	PrintBefore    int
	PrintContext   int
	PrintLineCount bool
	IgnoreCase     bool
	Invert         bool
	FixedStr       bool
	PrintLineNumb  bool

	CantFindFileError = errors.New("не могу найти файл, который вы указали")
	GetFileStatsError = errors.New("ошибка при получении информации о файле")
	ReadFileError     = errors.New("ошибка при чтении из файла")
)

func init() {
	flag.IntVar(&PrintAfter, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&PrintBefore, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&PrintContext, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&PrintLineCount, "c", false, "вывод кол-ва строк")
	flag.BoolVar(&IgnoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&Invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&FixedStr, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&PrintLineNumb, "n", false, "печатать номер строки")
}

func main() {
	flag.Parse()
	out, err := GrepStr(flag.Arg(0), flag.Args()[1:])
	if err != nil {
		log.Print(err)
	}

	for _, v := range out {
		fmt.Println(v)
	}
}

func GrepStr(pattern string, files []string) ([]string, error) {
	outLinesFiles := make([]string, 0, len(files))

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			return nil, CantFindFileError
		}
		defer file.Close()

		fileStats, err := file.Stat()
		if err != nil {
			return nil, GetFileStatsError
		}

		fileContens := make([]byte, fileStats.Size())

		_, err = file.Read(fileContens)
		if err != nil {
			return nil, ReadFileError
		}

		lines := strings.Split(string(fileContens), "\r\n")

		outLines := make([]string, 0, 2)

		for i, line := range lines {

			strNotContains := !strings.Contains(line, pattern)

			if IgnoreCase { // не зависим от регистра
				strNotContains = !strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			}

			if FixedStr {
				strNotContains = !(line == pattern) // подмена способа сравнения
			}

			if Invert {
				strNotContains = !strNotContains
			}

			if strNotContains {
				continue // к след файлу
			}

			if PrintLineNumb {
				line = strconv.Itoa(i) + ": " + line // медленная конкат..
			}

			if PrintContext > 0 {
				PrintBefore = PrintContext
				PrintAfter = PrintContext
			}

			if PrintBefore > 0 { // если попросят 100, когда строк 5 то выведет все что есть, и тихо завершится
				for lnsCount := PrintBefore; lnsCount > 0; lnsCount-- {
					if (i - lnsCount) < 0 {
						continue
					}

					outLines = append(outLines, lines[i-lnsCount])
				}
			}

			outLines = append(outLines, line)

			if PrintAfter > 0 { // если попросят 100, когда строк 5 то выведет все что есть, и тихо завершится
				for lnsCount := 1; lnsCount < PrintAfter+1; lnsCount++ {
					if (i + lnsCount) >= len(lines) {
						break
					}

					outLines = append(outLines, lines[i+lnsCount])
				}
			}

		}

		if PrintLineCount {
			outLinesFiles = append(outLinesFiles, strconv.Itoa(len(outLines)))
			continue // к след. файлу
		}

		outLines = append([]string{f + ": "}, outLines...) // название файла
		outLines = append(outLines, "--")                  // обозначаем конец файла

		outLinesFiles = append(outLinesFiles, outLines...)

	}

	return outLinesFiles, nil
}

//для тестов, оказалось больше аллокаций, и меньшая скорость, поэтому будем использовать с строками
/******************************************************************************/
func Grep(pattern string, files []string) error {

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			return CantFindFileError
		}
		defer file.Close()

		fileStats, err := file.Stat()
		if err != nil {
			return GetFileStatsError
		}

		fileContens := make([]byte, fileStats.Size())

		_, err = file.Read(fileContens)
		if err != nil {
			return ReadFileError
		}

		lines := bytes.Split(fileContens, []byte("\r\n"))

		for _, line := range lines {
			if bytes.Contains(line, []byte(pattern)) {
				//fmt.Println(string(line))
				_ = string(line)
			}
		}
	}

	return nil
}
