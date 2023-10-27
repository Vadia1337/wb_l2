package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern


	Строитель является порождающим паттерном. Он позволяет создавать сложные объекты пошагово.
	Он даёт возможность использовать один и тот же код строительства для построения разных представлений объектов.

	Применимость:
	Когда нужно создавать разные представления какого-то объекта.
	Когда нужно собирать сложные составные объекты.

	Плюсы:
	позволяет создавать продукты пошагово
	позволяет использовать один и тот же код для создания различных продуктов
	изолирует сложный код сборки продукта от бизнес-логики

	Минусы:
	усложняет код засчёт доп. структур
	клиент привязывается к конкретным структурам строителей

*/

func main() {
	bulider := newHumanBuilder()
	humanVanya := bulider.setHeight(190).setWeight(100).setName("Ваня").Build()
	fmt.Println(humanVanya)
}

type Human struct {
	Height uint16
	Weight uint16
	Name   string
}

type HumanBuilder interface {
	setHeight(height uint16) HumanBuilder
	setWeight(weight uint16) HumanBuilder
	setName(name string) HumanBuilder
	Build() *Human
}

type commonHumanBuilder struct {
	*Human // можно задать собственную структуру у билдера
}

func (cB *commonHumanBuilder) setHeight(height uint16) HumanBuilder {
	cB.Height = height
	return cB
}

func (cB *commonHumanBuilder) setWeight(weight uint16) HumanBuilder {
	cB.Weight = weight
	return cB
}

func (cB *commonHumanBuilder) setName(name string) HumanBuilder {
	cB.Name = name
	return cB
}

func (cB *commonHumanBuilder) Build() *Human {
	return cB.Human
}

func newHumanBuilder() *commonHumanBuilder {
	return &commonHumanBuilder{
		&Human{},
	}
}
