package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

	Посетитель является поведенческим паттерном. Он позволяет создавать новые
	операции, не меняя структуры объектов, над которыми эти операции могут выполняться.

	Применимость:
	Когда нужно выполнить операцию над всеми элементами сложной структуры объектов-
	посетитель позволяет применять одну и ту же операцию к объектам различных структур.

	плюсы:
	упрощает добавление новых операций над всей связанной структурой объектов
	объединяет родственные операции в одной структуре
	может накапливать состояние при обходе структуры компонентов

	минусы:
	применение неоправдано, если иерархия компонентов часто меняется
	может привести к нарушению инкапсуляции компонентов
*/

func main() {
	sdek := &sdekDelivery{
		price:       500,
		coefficient: .9,
	}

	pochta := &pochtaRossiiDelivery{
		price:       300,
		coefficient: .6,
	}

	boxberry := &boxberryDelivery{
		price:       250,
		coefficient: .7,
	}

	commonCalc := commonCalcVisitor{}
	commonCalc.visitSdek(sdek)
	commonCalc.visitPochtaRossii(pochta)
	commonCalc.visitBoxberry(boxberry)

	calcWithMyFormula := calcWithMyFormulaVisitor{}
	calcWithMyFormula.visitSdek(sdek)
	calcWithMyFormula.visitPochtaRossii(pochta)
	calcWithMyFormula.visitBoxberry(boxberry)

}

type visitor interface {
	visitSdek(sdek *sdekDelivery)
	visitPochtaRossii(pochta *pochtaRossiiDelivery)
	visitBoxberry(boxberry *boxberryDelivery)
}

type calcWithMyFormulaVisitor struct{} // у визиторов могут быть свои уникальные структуры.

func (calc *calcWithMyFormulaVisitor) visitSdek(sdek *sdekDelivery) {
	deliveryPrice := (2*sdek.price)*sdek.coefficient + 500
	fmt.Println("Цена доставки в сдеке по моей ХИТРОЙ формуле: ", deliveryPrice)
}

func (calc *calcWithMyFormulaVisitor) visitPochtaRossii(pochta *pochtaRossiiDelivery) {
	deliveryPrice := (2*pochta.price)*pochta.coefficient + 500
	fmt.Println("Цена доставки в почте России по моей ХИТРОЙ формуле: ", deliveryPrice)
}

func (calc *calcWithMyFormulaVisitor) visitBoxberry(boxberry *boxberryDelivery) {
	deliveryPrice := (2*boxberry.price)*boxberry.coefficient + 500
	fmt.Println("Цена доставки в боксбери по моей ХИТРОЙ формуле: ", deliveryPrice)
}

type commonCalcVisitor struct{}

func (calc *commonCalcVisitor) visitSdek(sdek *sdekDelivery) {
	deliveryPrice := sdek.price * sdek.coefficient
	fmt.Println("Цена доставки в сдеке по обычной формуле: ", deliveryPrice)
}

func (calc *commonCalcVisitor) visitPochtaRossii(pochta *pochtaRossiiDelivery) {
	deliveryPrice := pochta.price * pochta.coefficient
	fmt.Println("Цена доставки в почте России по обычной формуле: ", deliveryPrice)
}

func (calc *commonCalcVisitor) visitBoxberry(boxberry *boxberryDelivery) {
	deliveryPrice := boxberry.price * boxberry.coefficient
	fmt.Println("Цена доставки в боксбери по обычной формуле: ", deliveryPrice)
}

type sdekDelivery struct {
	price       float64
	coefficient float64
}

func (sdek *sdekDelivery) accept(v visitor) {
	v.visitSdek(sdek)
}

type pochtaRossiiDelivery struct {
	price       float64
	coefficient float64
}

func (pochta *pochtaRossiiDelivery) accept(v visitor) {
	v.visitPochtaRossii(pochta)
}

type boxberryDelivery struct {
	price       float64
	coefficient float64
}

func (boxberry *boxberryDelivery) accept(v visitor) {
	v.visitBoxberry(boxberry)
}
