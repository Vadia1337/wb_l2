package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

	Цепочка обязаностей является поведенческим паттерном. Она позволяет передавать запросы
	последовательно по цепочке обработчиков. Каждый последующий обработчик решает,
	может ли он обработать запрос сам и стоит ли передавать его дальше.

	Применимость:
	Когда программа содержит несколько объектов, способных обрабатывать тот или иной запрос,
	но заранее неизвестно какой запрос придёт и какой обработчик понадобится.
	Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	Когда набор объектов, способных обработать запрос, должен задаваться динамически.

	Плюсы:
	уменьшает зависимость между клиентом и обработчиком
	реализует принцип единственной обязанности
	реализует принцип открытости/закрытости

	Минусы:
	запрос может остаться без обработки
*/

type callCenterСlient struct { //Клиент колл центра
	levelQuestion int
	description   string
}

type supHandler interface {
	setNextHandler(supHandler)
	handle(*callCenterСlient)
}

type baseSupHandler struct {
	nextHandler supHandler
}

func (baseSH *baseSupHandler) setNextHandler(handler supHandler) {
	baseSH.nextHandler = handler
}

type juniorSupHandler struct {
	baseSupHandler
}

func (juniorSH *juniorSupHandler) handle(client *callCenterСlient) {
	if client.levelQuestion == 3 {
		fmt.Println("джун-работник обслужил клиента с вопросом: \n",
			client.description, "Уровень вопроса: ", client.levelQuestion)

		return
	}

	juniorSH.nextHandler.handle(client) // жун зовет на помощь кого постарше -)
}

type middleSupHandler struct {
	baseSupHandler
}

func (middleSH *middleSupHandler) handle(client *callCenterСlient) {
	if client.levelQuestion == 2 {
		fmt.Println("мидл-работник обслужил клиента с вопросом: \n",
			client.description, "Уровень вопроса: ", client.levelQuestion)

		return
	}

	middleSH.nextHandler.handle(client) // мидл зовет на помощь, по всей видимости у нашей конторки проблемы (
}

type directorSupHandler struct {
	baseSupHandler
}

func (directorSH *directorSupHandler) handle(client *callCenterСlient) {
	if client.levelQuestion == 1 {
		fmt.Println("директор обслужил клиента с вопросом: \n",
			client.description, "Уровень вопроса: ", client.levelQuestion)
	}
}

func main() {

	junior := &juniorSupHandler{}
	middle := &middleSupHandler{}
	director := &directorSupHandler{}

	junior.setNextHandler(middle)
	middle.setNextHandler(director)

	clients := []*callCenterСlient{
		{levelQuestion: 1, description: "Что у вас за контора такая? где мои деньги?"},
		{levelQuestion: 3, description: "Хочу просто поговорить..."},
		{levelQuestion: 2, description: "Хочу застраховать собаку"},

		{levelQuestion: 4, description: "Мой вопрос точно не останется без вашего внимания?"},
	}

	for _, v := range clients {
		junior.handle(v)
	}

}
