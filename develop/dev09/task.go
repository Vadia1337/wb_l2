package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	AddrNotSet  = errors.New("вы не указали адрес сайта")
	GetError    = errors.New("ошибка при отправке get запроса")
	CreateError = errors.New("ошибка при создании index.html файла")
	CopyError   = errors.New("ошибка при копировании сайта в файл")
)

func main() {
	flag.Parse()
	addr := flag.Arg(0)

	err := wget(addr)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func wget(addr string) error {

	if addr == "" {
		return AddrNotSet
	}

	fmt.Println("Ожидайте завершения ...")

	resp, err := http.Get(addr)
	if err != nil {
		return GetError
	}

	defer resp.Body.Close()

	file, err := os.Create("index.html")
	if err != nil {
		return CreateError
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return CopyError
	}

	fmt.Println("Сайт успешно скачался, ищите его в папке index.html")

	return nil
}
