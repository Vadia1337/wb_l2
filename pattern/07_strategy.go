package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

	Стратегия является поведенческим паттерном. Он определяет семейство схожих алгоритмов и помещает
	каждый в собственную структуру. После этого алгоритмы можно взаимозаменять по ходу программы.

	Применимость:
	Когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
	Когда похожие подструктуры отличают только некоторым поведением.
	Когда нужно скрыть детали реализации алгоритмов от другоих структур.
	Когда различне вариации алгоритмов реализованы в виде развесистого условного оператора.

	Плюсы:
	"горячая" замена алгоритмов "налету"
	изолирует код и данные алгоритмов от остального кода
	уход от наследования к делегированию
	реализует пирнцип открытости/закрытости

	Минусы:
	усложняет код засчёт доп. структур
	клиент должен знать разницу между стратегиями, чтобы выбрать подходящую
*/

func main() {
	dr := &driver{}

	aggressive := &aggressiveStrategy{}
	careful := &carefulStrategy{}

	dr.setStrategy(aggressive)
	dr.drive()

	dr.setStrategy(careful)
	dr.drive()
}

type driverStrategy interface {
	strategy()
}

type driver struct {
	strategy driverStrategy
}

func (d *driver) setStrategy(strategy driverStrategy) {
	d.strategy = strategy
}

func (d *driver) drive() {
	fmt.Println("Водитель начал движение")
	d.strategy.strategy()
}

type aggressiveStrategy struct{}

func (aggressiveStrategy) strategy() {
	fmt.Println("Водитель будет хасанить")
}

type carefulStrategy struct{}

func (carefulStrategy) strategy() {
	fmt.Println("Водитель будет ехать осторожно")
}
