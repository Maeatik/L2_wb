package main

import (
	"fmt"
	"sort"
	"strings"
)

func createAnagram(arr []string) map[string][]string {
	//создадим мапу, в качестве ключа будет первое слово, для анаграмм,
	//в качестве параметров будут все анаграммы, включая ключ
	anagramMap := make(map[string][]string)

	//пробегаемся по словам из слайса
	for _, str := range arr {
		//так как регистр не важен, переводим все символы в нижний регистр
		strLower := strings.ToLower(str)

		//Добавление слова по ключу анаграммы. Если ключ нашелся, добавляем к нему слово
		if key, flag := searchKeyInAnagramMap(strLower, anagramMap); flag {
			anagramMap[key] = append(anagramMap[key], strLower)
		}
	}
	//проходим по получившейся мапе. Если в ней есть такие ключи,
	//к которому относятся только один элемент - этот ключ и элемент удаляются
	for key, val := range anagramMap {
		if len(val) == 1 {
			delete(anagramMap, key)
		} else {
			//сортируем элементы ключа по алфавитному порядку первых букв слов
			sort.Strings(anagramMap[key])
		}
	}
	//готовая мапа
	return anagramMap
}

func searchKeyInAnagramMap(str string, anagramMap map[string][]string) (string, bool) {
	//проходимся по ключам и их значениям
	for key, val := range anagramMap {
		//если введенное слово анаграмма ключа проверяем, есть ли уже это слово в мапе
		if isAnagram(str, key) {
			if elemInMap(str, val) {
				//если есть то возвращаем пустую строку и булево значение для проверки
				return "", false
			}
			//если такого слова нет, говорим об этом и возвращаем ключ
			return key, true
		}
	}
	//если такого ключа не нашлось, создаем ключ и добавляем этот элемент в мапу
	return str, true
}

//функция проверки на анаграмму
func isAnagram(str, anagram string) bool {
	//переводим проверяемое слово и слово-ключ из мапы в слайс рун
	strRunes := []rune(str)
	anagramRunes := []rune(anagram)
	//сортируем руны в порядке возрастания
	sort.Slice(strRunes, func(i, j int) bool {
		return strRunes[i] < strRunes[j]
	})
	sort.Slice(anagramRunes, func(i, j int) bool {
		return anagramRunes[i] < anagramRunes[j]
	})
	//если руны совпадают, значит слова являются анаграммами
	return string(strRunes) == string(anagramRunes)
}

//проверяем, есть ли слово в мапе
func elemInMap(str string, vals []string) bool {
	for _, val := range vals {
		if str == val {
			return true
		}
	}
	return false
}

func main() {
	//создаем слайс из слов для проверки на анаграммы
	arr := []string{"ПРивет", "тевИрп", "солнце", "солнце", "цеслон", "ценолс", "выываыва"}
	//создадим множество из анаграмм
	res := createAnagram(arr)
	fmt.Println(res)
}
