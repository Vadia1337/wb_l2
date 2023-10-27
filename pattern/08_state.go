package main

import (
	"fmt"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

	Состояние является поведенческим паттерном. Он позволяет объектам менять поведение в зависимости
	от своего состояния. Извне создаётся впечатление, что изменилась структура объекта.

	Применимость:
	Когда есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния.
	Когда код содержит множество условных операторов, которые выбирают поведение в зависимости от текущих полей структуры.

	Плюсы:
	избавляет от множества условных операторов машины состояний
	концентрирует в одном место код, связанныйм определённым состоянием
	упрощает код контекста

	Минусы:
	может усложнить код, если состояний мало или они редко меняются
*/

func main() {
	dr := &someDriver{}

	alc := &AlcoholState{
		description: "В алкогольном состоянии",
	}
	normal := &NormalState{
		description: "В нормальном состоянии",
	}

	alc.state(dr)
	fmt.Println("Водитель, в каком вы состоянии?", dr.getState())

	normal.state(dr)
	fmt.Println("Водитель, в каком вы состоянии?", dr.getState())

}

type driverState interface {
	state(*someDriver)
}

type someDriver struct {
	state driverState
}

func (dr *someDriver) getState() driverState {
	return dr.state
}

func (dr *someDriver) setState(state driverState) {
	dr.state = state
}

type AlcoholState struct {
	description string
}

func (state *AlcoholState) state(dr *someDriver) {
	dr.setState(state)
}

type NormalState struct {
	description string
}

func (state *NormalState) state(dr *someDriver) {
	dr.setState(state)
}
