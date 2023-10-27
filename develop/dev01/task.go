package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.

! модуль конечно нужно называть по типу https://github.com/имяЮзера/названиеМодуля, но так как грузить
 не собираемся, то можно просто назвать)
*/

const timeServer = "0.beevik-ntp.pool.ntp.org"

func main() {
	time, err := ntp.Time(timeServer)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error. Go module NTPtime.main: ", err.Error())
		os.Exit(1)
		return
	}

	fmt.Println(time)
}
