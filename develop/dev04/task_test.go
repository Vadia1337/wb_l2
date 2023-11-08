package main

import (
	"reflect"
	"testing"
)

// проблема теста - нельзя сравнить две мапы с кол-вом ключей более 2, т.к мапа при итерации выводит рандомно
// но и такого набора данных нам для тестов будет достаточно
func TestBunchesOfAnagrams(t *testing.T) {
	testDataSet := []struct {
		testNumb    int
		Words       []string
		outAnagrams map[string]*[]string
	}{
		{
			testNumb:    1,
			Words:       []string{"нирвана", "нирвана"},
			outAnagrams: map[string]*[]string{},
		},
		{
			testNumb: 2,
			Words:    []string{"рванина", "равнина", "нирвана"}, // тестируем сортировку
			outAnagrams: map[string]*[]string{
				"рванина": &[]string{"нирвана", "равнина", "рванина"},
			},
		},
	}

	for _, v := range testDataSet {
		mapOfAnagrams := *getBunchesOfAnagrams(&v.Words)
		if len(mapOfAnagrams) != len(v.outAnagrams) {
			t.Error("Размеры входной и выходной карт, не совпали, набор данных №", v.testNumb)
		}

		for i, sliceInTestMap := range v.outAnagrams {

			// проверка ключей карты
			sliceInMap, ok := mapOfAnagrams[i]
			if !ok {
				t.Error("Ключи входной и выходной карт, не совпали набор данных №", v.testNumb)
			}

			// проверка значений карты
			if !reflect.DeepEqual(sliceInTestMap, sliceInMap) { // медленно, но верно :-)
				t.Error("Значения входной и выходной карт, не совпали набор данных №", v.testNumb)
			}
		}
	}
}

func BenchmarkBunchesOfAnagrams(b *testing.B) {
	ArrayOfWords := &[]string{"нирвана", "нирвана", "столик", "корсет", "отсечка", "останки", "костер", "Равнина",
		"сеточка", "Стоечка", "Плотер", "рванина", "сектор", "листок", "остров", "пролет", "тесачок", "слиток",
		"остинка", "чесотка", "скотина", "трепло"}

	for i := 0; i < b.N; i++ {
		getBunchesOfAnagrams(ArrayOfWords)
	}
}
