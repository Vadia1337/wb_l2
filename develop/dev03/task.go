package main

import (
	"bufio"
	"bytes"
	flag "flag"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	SelectColumnFlag int
	SortByInt        bool
	ReverseSort      bool
	UniqueValues     bool
)

func init() {
	flag.IntVar(&SelectColumnFlag, "k", 1, "selecting a column to sort, for example: -k 1 - sorting by first column")
	flag.BoolVar(&SortByInt, "n", false, "sort by int")
	flag.BoolVar(&ReverseSort, "r", false, "reverse sort")
	flag.BoolVar(&UniqueValues, "u", false, "leaves only unique values")
}

func main() {
	flag.Parse()
	SortSmallFile(flag.Arg(0))
}

func SortSmallFile(fileName string) {
	// валидация аргументов

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Не могу найти файл, который вы указали")
		os.Exit(1)
	}
	defer file.Close()

	fileStats, err := file.Stat()
	if err != nil {
		fmt.Println("Ошибка при получении информации о файле")
		os.Exit(1)
	}

	strLines := make(map[string][]byte)  // размер = кол-во строк в файле
	strLinesKeys := make([]string, 0, 2) // размер = кол-во строк в файле

	fileContens := make([]byte, fileStats.Size())

	_, err = file.Read(fileContens)
	if err != nil {
		fmt.Println("Ошибка при чтении из файла")
		os.Exit(1)
	}

	lines := bytes.Split(fileContens, []byte("\r\n"))

	for i, v := range lines {
		lineKey := string(bytes.Split(v, []byte(" "))[SelectColumnFlag-1])

		_, ok := strLines[lineKey]
		if ok {
			if UniqueValues {
				continue
			}

			if !SortByInt {
				lineKey += fmt.Sprintf("%f", float32(i)/100000)
			}

		}

		strLines[lineKey] = v
		strLinesKeys = append(strLinesKeys, lineKey)
	}

	fmt.Println(strLinesKeys)

	if SortByInt { // решил замудрить, чтобы поиспользовать что-то новое :-)

		type sortStruct struct {
			strSlice string
			intSlice float64
		}

		var sortStructs []sortStruct

		for i, s := range strLinesKeys {
			strToInt, err := strconv.ParseFloat(s, 32)
			if err != nil {
				fmt.Println("Столбец не содержит числа, попробуйте выбрать другой столбец для сортировки по числу")
				os.Exit(1)
			}
			strToInt += float64(i) / 100000
			strLinesKeys[i] = fmt.Sprintf("%f", strToInt)
			sortStructs = append(sortStructs, sortStruct{
				strSlice: s,
				intSlice: strToInt,
			})
		}

		sort.Slice(sortStructs, func(i, j int) bool {
			return sortStructs[i].intSlice < sortStructs[j].intSlice
		})

		for i, v := range sortStructs {
			strLinesKeys[i] = v.strSlice
		}

	} else {
		sort.Strings(strLinesKeys)
	}

	if ReverseSort {
		slices.Reverse(strLinesKeys)
	}

	open, err := os.Create("out.txt")
	if err != nil {
		fmt.Println("Не получилось открыть файл для записи")
		os.Exit(1)
	}
	defer open.Close()

	var buffer bytes.Buffer
	for _, v := range strLinesKeys {
		buffer.Write(strLines[v])
		buffer.Write([]byte("\r\n"))
	}

	_, err = open.Write(buffer.Bytes())
	if err != nil {
		fmt.Println("Ошибка при записи в файл")
		os.Exit(1)
	}

}

/******************************************/
// реализовал для тестов
func Sort(fileName string) {

	// валидация аргументов

	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.Stat()
	if err != nil {
		return
	}

	strLines := make(map[string]string)  // размер = кол-во строк в файле
	strLinesKeys := make([]string, 0, 2) // размер = кол-во строк в файле

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		lineKey := strings.Split(line, " ")[SelectColumnFlag-1]
		strLines[lineKey] = line
		strLinesKeys = append(strLinesKeys, lineKey)
	}

	sort.Strings(strLinesKeys)

	open, err := os.Create("out.txt")
	if err != nil {
		return
	}
	defer open.Close()

	for _, v := range strLinesKeys {
		line := strLines[v]

		bs := make([]byte, len(line)+len("\n"))
		bl := 0
		bl += copy(bs[bl:], line)
		bl += copy(bs[bl:], "\n")

		_, err = open.Write(bs)
		if err != nil {
			return
		}

	}
}
