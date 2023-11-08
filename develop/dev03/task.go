package main

import (
	"bufio"
	"bytes"
	"errors"
	flag "flag"
	"log"
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

	SortStrByNFlagError      = errors.New("вы пытаетесь отсортировать строку с использованием флага -n, который сортирует числа")
	CantOpenFileToWriteError = errors.New("не получилось открыть файл для записи")
	WriteToFileError         = errors.New("ошибка при записи в файл")
	CantFindFileError        = errors.New("не могу найти файл, который вы указали")
	GetFileStatsError        = errors.New("ошибка при получении информации о файле")
	ReadFileError            = errors.New("ошибка при чтении из файла")
)

func init() {
	flag.IntVar(&SelectColumnFlag, "k", 1, "selecting a column to sort, for example: -k 1 - sorting by first column")
	flag.BoolVar(&SortByInt, "n", false, "sort by int")
	flag.BoolVar(&ReverseSort, "r", false, "reverse sort")
	flag.BoolVar(&UniqueValues, "u", false, "leaves only unique values")
}

func main() {
	flag.Parse()
	err := Sort(flag.Arg(0))
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

// было решено сгрузить весь файл в озу, что в случае с большими файлами, черевато переполнением памяти.

// другим решением, обеспечивающем работу с файлами большого размера, могло стать построчное чтение файла, или чтение
// партиями, и тут же сохранение в файлы по определенному признаку, к примеру для int - в 1 файле может храниться
// отсортированнный диапазон значений в который могут добавляться новые, в случае с строками дело обстоит сложнее..
// далее такие файлы нужно склеить друг с другом и получить итоговый отсортированный файл, но такая реализация выглядит
// очень сложной, больше похожей на разработку собственной БД ;-)
func Sort(fileName string) error {

	file, err := os.Open(fileName)
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

	if SortByInt {
		err = sortByInt(lines)
		if err != nil {
			return err
		}
	} else {
		err = sortByStr(lines)
		if err != nil {
			return err
		}
	}

	return nil
}

func sortByStr(lines [][]byte) error {
	strLines := make(map[string][]byte)
	strLinesKeys := make([]string, 0, 2)

	for i, v := range lines {
		lineKey := string(bytes.Split(v, []byte(" "))[SelectColumnFlag-1])

		_, ok := strLines[lineKey]
		if ok {
			if UniqueValues {
				continue
			}

			lineKey += strconv.Itoa(i)
		}

		strLines[lineKey] = v
		strLinesKeys = append(strLinesKeys, lineKey)
	}

	sort.Strings(strLinesKeys)

	if ReverseSort {
		slices.Reverse(strLinesKeys)
	}

	open, err := os.Create("out.txt")
	if err != nil {
		return CantOpenFileToWriteError
	}
	defer open.Close()

	var buffer bytes.Buffer
	for _, v := range strLinesKeys {
		buffer.Write(strLines[v])
		buffer.Write([]byte("\r\n"))
	}

	_, err = open.Write(buffer.Bytes())
	if err != nil {
		return WriteToFileError
	}
	return nil
}

func sortByInt(lines [][]byte) error {
	strLines := make(map[float64][]byte)
	strLinesKeys := make([]float64, 0, 2)

	for i, v := range lines {
		strToFloat := string(bytes.Split(v, []byte(" "))[SelectColumnFlag-1])
		lineKey, err := strconv.ParseFloat(strToFloat, 64)
		if err != nil {
			return SortStrByNFlagError
		}

		_, ok := strLines[lineKey]
		if ok {
			if UniqueValues {
				continue
			}

			lineKey += float64(i) / 100000
		}

		strLines[lineKey] = v
		strLinesKeys = append(strLinesKeys, lineKey)
	}

	sort.Float64s(strLinesKeys)

	if ReverseSort {
		slices.Reverse(strLinesKeys)
	}

	open, err := os.Create("out.txt")
	if err != nil {
		return CantOpenFileToWriteError
	}
	defer open.Close()

	var buffer bytes.Buffer
	for _, v := range strLinesKeys {
		buffer.Write(strLines[v])
		buffer.Write([]byte("\r\n"))
	}

	_, err = open.Write(buffer.Bytes())
	if err != nil {
		return WriteToFileError
	}
	return nil
}

/******************************************/
// реализовал для тестов
func SortOld(fileName string) {

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
