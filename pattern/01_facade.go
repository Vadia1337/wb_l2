package main

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
фасад это
структурный шаблон проектирования, позволяющий скрыть сложность системы путём сведения всех возможных внешних вызовов
к одному объекту, делегирующему их соответствующим объектам системы.

Плюсы:
Изолирует клиентов от компонентов сложной подсистемы

Минусы:
Фасад легко может превратится в божественный объект.
Божественный объект (англ. God object) — антипаттерн объектно-ориентированного программирования,
описывающий объект, который хранит в себе «слишком много» или делает «слишком много».

*/

func main() {
	facade := newFarmFacade()
	facade.makeFood()
	facade.manageСhickenСoop()
}

type farmFacade struct {
	feedFactory *feedFactory
	chickenСoop *chickenСoop
}

func newFarmFacade() *farmFacade {
	return &farmFacade{
		feedFactory: &feedFactory{},
		chickenСoop: &chickenСoop{},
	}
}

func (f *farmFacade) makeFood() {
	f.feedFactory.fraction()
	f.feedFactory.mix()
	f.feedFactory.passThroughExtruder()
}

func (f *farmFacade) manageСhickenСoop() {
	f.chickenСoop.feedBirds()
	f.chickenСoop.collectEggs()
	f.chickenСoop.clean()
}

type feedFactory struct{}

func (f *feedFactory) fraction() {
	fmt.Println("Измельчили компоненты корма")
}

func (f *feedFactory) mix() {
	fmt.Println("Смешали компоненты корма")
}

func (f *feedFactory) passThroughExtruder() {
	fmt.Println("Пропустили через экструдер, корм готов!")
}

type chickenСoop struct{}

func (c *chickenСoop) feedBirds() {
	fmt.Println("Накормили птиц")
}

func (c *chickenСoop) collectEggs() {
	fmt.Println("Собрали яйца")
}

func (c *chickenСoop) clean() {
	fmt.Println("Убрались в курятнике")
}
