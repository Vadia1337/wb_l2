package main

import (
	"fmt"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern


	Фабричный метод является порождающим паттерном. Он определяет обищй инерфейс для создания объектов в
	суперструктуре,позволяя подструктурам изменять типа создаваемых объектов.

	Применимость:
	Когда заранее неизвестны типа и зависимости объектов, с которыми должен работать код.
	Когда нужно предоставить возможность пользователям расширять части фреймворка или пакета.
	Когда нужно экономить системные ресурсы, повторно используя уже созданные объекты.


	Плюсы
	избавляет структуру от привязки к конкретным структурам продуктов
	выделеляет код производства продуктов в одно место, упрощая поддержку кода
	упрощает добавление новых продуктов в программу
	реализует принцип открытости/закрытости

	Минусы
	может привести к созданию больших параллельных иерархий структур, так как для
	каждой структуры продукта нужно создать свою подструктуру создателя.
*/

const (
	DellMonitorType = "dell"
	LGMonitorType   = "lg"
	DEXPMonitorType = "dexp"
)

func main() {

	dell := newMonitor("dell")
	dell.sayAbout()

	lg := newMonitor("lg")
	lg.sayAbout()

	dexp := newMonitor("dexp")
	dexp.sayAbout()
}

type monitor interface {
	sayAbout()
}

func newMonitor(typeName string) monitor {
	switch typeName {
	case DellMonitorType:
		return newDellMonitor()
	case LGMonitorType:
		return newLGMonitor()
	case DEXPMonitorType:
		return newDEXPMonitor()
	default:
		fmt.Println("Такой тип не знаю")
		return nil
	}
}

type dellMonitor struct {
	diagonal    string
	resolution  string
	description string
}

func newDellMonitor() *dellMonitor {
	return &dellMonitor{
		diagonal:    "27",
		resolution:  "1920x1080",
		description: "Широкий монитор с плохим качеством",
	}
}

func (dell *dellMonitor) sayAbout() {
	fmt.Println(dell)
}

type lgMonitor struct {
	diagonal    string
	resolution  string
	description string
}

func newLGMonitor() *lgMonitor {
	return &lgMonitor{
		diagonal:    "23",
		resolution:  "2440x1440",
		description: "Средний монитор с хорошим качеством",
	}
}

func (lg *lgMonitor) sayAbout() {
	fmt.Println(lg)
}

type DEXPMonitor struct {
	diagonal    string
	resolution  string
	description string
}

func newDEXPMonitor() *DEXPMonitor {
	return &DEXPMonitor{
		diagonal:    "21",
		resolution:  "1440x820",
		description: "Плохой монитор",
	}
}

func (dexp *DEXPMonitor) sayAbout() {
	fmt.Println(dexp)
}
