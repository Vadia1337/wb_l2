package main

import (
	"fmt"
	"slices"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	ArrayOfWords := &[]string{"нирвана", "нирвана", "столик", "корсет", "отсечка", "останки", "костер", "Равнина",
		"сеточка", "Стоечка", "Плотер", "рванина", "сектор", "листок", "остров", "пролет", "тесачок", "слиток",
		"остинка", "чесотка", "скотина", "трепло"}

	mapOfAnagrams := *getBunchesOfAnagrams(ArrayOfWords)
	for _, v := range mapOfAnagrams {
		fmt.Println(*v)
	}
}

func getBunchesOfAnagrams(ArrayOfWords *[]string) *map[string]*[]string {

	mapOfAnagrams := make(map[string]*[]string)
	sumRunesInWordAsMapKey := make(map[int32]string)

iterateArrayOfWords:
	for _, word := range *ArrayOfWords {
		wordToLower := strings.ToLower(word)

		var sumRunes int32
		for _, runeinWord := range wordToLower {
			sumRunes += runeinWord
		}

		mapKey, ok := sumRunesInWordAsMapKey[sumRunes]
		if !ok {
			sumRunesInWordAsMapKey[sumRunes] = wordToLower

			newSliceOfAnagrams := make([]string, 1, 2) // снижаем кол-во аллокаций, было &[]string{wordToLower}
			newSliceOfAnagrams[0] = wordToLower        //в данном случае можем указать капатиси т.к по условию слайс >=2
			mapOfAnagrams[wordToLower] = &newSliceOfAnagrams

			continue
		}

		// проверка на одинаковые слова, которые не являются аннограмами
		sliceOfAnagrams := *mapOfAnagrams[mapKey]
		for _, v := range sliceOfAnagrams {
			if v == wordToLower {

				continue iterateArrayOfWords
			}
		}

		*mapOfAnagrams[mapKey] = append(*mapOfAnagrams[mapKey], wordToLower)
	}

	//выкидываем одиночные слова, сортируем слайсы
	for i, v := range mapOfAnagrams {
		if len(*v) <= 1 {
			delete(mapOfAnagrams, i)
		}
		slices.Sort(*v)
	}

	return &mapOfAnagrams
}
