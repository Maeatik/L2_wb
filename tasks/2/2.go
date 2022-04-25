package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func newString(s string) (string, error) {
	//переводим строку в слайс рун
	runes := []rune(s)
	//узнаем количество символов
	runeLen := len(runes)
	//если есть строка и при этом первый ее символ число - кидаем ошибку
	if runeLen > 0 && unicode.IsDigit(runes[0]) {
		return "", errors.New("строка начианется с числа")
	}
	//создаем Билдер для эффективного создания новой строки
	resString := &strings.Builder{}
	var curSymbol rune
	//обрабатываем каждый символ
	for i := 0; i < runeLen; i++ {
		curRune := runes[i]
		//если текущий для проверки символ - число
		if unicode.IsDigit(curRune) {
			//если следующий символ - число - кидаем ошибку
			if i+1 < runeLen && unicode.IsDigit(runes[i+1]) {
				return "", errors.New("два числа не могут идти подряд")
			}
			//переводим руну в строку, а потом в инт
			count, err := strconv.Atoi(string(curRune))

			if err != nil {
				return "", nil
			}
			//увеличиваем количество символов
			getNewRunes(resString, curSymbol, count)
			continue
		}
		//если на вход дан слеш
		if curRune == '\\' {
			i++
			//если косая черта -последний символ строки - кидаем ошибку
			if i == runeLen {
				return "", errors.New("косая черта - последний символ строки")
			}
			//берем в текущую руну следующий элемент слайса после слеша
			curRune = runes[i]
		}
		//если после буквы стоит число - запоминаем руну
		if i+1 < runeLen && unicode.IsDigit(runes[i+1]) {
			curSymbol = curRune
			continue
		}
		//записываем в итоговую строку букву, после которой нет цифры, либо символ,
		//который стоит после слеша
		resString.WriteRune(curRune)

	}
	return resString.String(), nil

}

func getNewRunes(runes *strings.Builder, r rune, count int) {
	//записываем руну r count-раз
	for i := 0; i < count; i++ {
		runes.WriteRune(r)
	}
}

func main() {
	//задаем строку на обработку
	str, err := newString("ab3c\\4")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Err: %v\n", err)
	}
	fmt.Println(str)
}
