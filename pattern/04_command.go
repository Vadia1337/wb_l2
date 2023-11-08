package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

	Команда является поведенческим паттерном. Она превращает запросы в объекты,
	позоляя передавать их как аргументы при вызове методов, ставить запросы в очередь,
	логировать их, поддерживать отмену операций.

	Применимость:
	Когда нужно составить очередь из операций, выполнять их по расписанию или передавать их по сети.
	Когда нужна поддержка отмены.
	Когда нужно параметризовать объекты выполняемым действием.

	Плюсы:
	+ убирает прямую зависимость между объектами
	+ позволяет реализовать простую отмену, повтор операций, отложенный запуск команд
	+ позволет собирать сложные команды из простых
	+ реализует принцип открытости/закрытости

	Минусы:
	- усложняет код засчет доп. структур
*/

type Command interface {
	Execute()
}

type SimpleCommand struct {
	receiver *Receiver
	payload  string
}

func NewSimpleCommand(receiver *Receiver, payload string) *SimpleCommand {
	return &SimpleCommand{receiver: receiver, payload: payload}
}

func (c *SimpleCommand) Execute() {
	fmt.Printf("SimpleCommand: выполнение команды '%s'\n", c.payload)
	c.receiver.DoSomething(c.payload)
}

type Receiver struct{}

func (r *Receiver) DoSomething(payload string) {
	fmt.Printf("Получатель: работаю над: '%s'\n", payload)
}

type Invoker struct { //вызыватель
	history []Command
}

func (i *Invoker) StoreAndExecute(cmd Command) {
	i.history = append(i.history, cmd)
	cmd.Execute()
}

func main() {
	receiver := &Receiver{}
	cmd := NewSimpleCommand(receiver, "команда 1")
	inv := &Invoker{}
	inv.StoreAndExecute(cmd)

	cmd2 := NewSimpleCommand(receiver, "команда 2")
	inv.StoreAndExecute(cmd2)

	fmt.Println("Повторим что происходило:")
	for _, cmd := range inv.history {
		cmd.Execute()
	}
}
