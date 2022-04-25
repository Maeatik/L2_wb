package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//Структура со всеми флагами
type flags struct {
	After      int  // -A
	Before     int  // -B
	Context    int  // -C
	Count      bool // -c
	IgnoreCase bool // -i
	Invert     bool // -v
	Fixed      bool // -F
	LineNum    bool // -n
}

type checking interface {
	check(s string) bool
}

type regexpChecker struct {
	R *regexp.Regexp
}

func newRegexpChecker(r *regexp.Regexp) *regexpChecker {
	return &regexpChecker{R: r}
}

func (r *regexpChecker) check(str string) bool {
	return r.R.MatchString(str)
}

type equalChecker struct {
	str string
}

func newEqualChecker(str string) *equalChecker {
	return &equalChecker{str: str}
}

func (e *equalChecker) check(str string) bool {
	return strings.Contains(str, e.str)
}

func grep(reader io.Reader, str string, flag *flags) (string, error) {

	var ch checking
	//если был указан флаг на игнорирование регистра переводим строку для проверки в нижний регистр
	if flag.IgnoreCase {
		str = strings.ToLower(str)
	}
	//если был указан флаг на точное совпадение
	if flag.Fixed {
		//создаем чекера
		ch = newEqualChecker(str)
		//если такого флага нет
	} else {
		//анализируем строку и переводим в regexp, для сопоставления текста
		rx, err := regexp.Compile(str)
		if err != nil {
			return "", err
		}
		//создаем чекера
		ch = newRegexpChecker(rx)
	}
	//создаем переменные для флагов
	numBefore := flag.Context
	numAfter := flag.Context
	//если есть флаг для поиска строк до совпадения, записываем его значение в эту переменную
	if flag.Before > 0 {
		numBefore = flag.Before
	}
	//если есть флаг для поиска строк после совпадения, записываем его значение в эту переменную
	if flag.After > 0 {
		numAfter = flag.After
	}
	//создаем мапу, в качестве ключей будут числа, в качестве значений - були.
	//в нее будем записывать строки, которые прошли проверки
	strIndex := make(map[int]bool)
	var id []int
	var allStr []string
	var counter, i int

	//читаем файл
	sc := bufio.NewScanner(reader)
	for sc.Scan() {
		//читаем строку
		text := sc.Text()
		//добавляем строку в слайс со всеми строками
		allStr = append(allStr, text)
		//если указан флаг на игнорирование регистра - переводим строку в нижний регистр
		if flag.IgnoreCase {
			text = strings.ToLower(text)
		}
		//если строки совпали
		if ch.check(text) {
			//увеличиваем счетчик совпадений
			counter = counter + 1
			//если есть флаг для поиска строк до совпадения
			if numBefore > 0 {
				//проверяем строки до совпадения
				for j := i - numBefore; j < i; j++ {
					//если эта строка не является совпадение запоминаем его индекс
					if _, ok := strIndex[j]; !ok {
						id = append(id, j)
						strIndex[j] = false
					}
				}
			}
			//если есть флаг для поиска строк после совпадения
			if numAfter > 0 {
				//проверяем строки после совпадения
				for j := i + numAfter; j > i; j-- {
					//если эта строка не является совпадение запоминаем его индекс
					if _, ok := strIndex[j]; !ok {
						id = append(id, j)
						strIndex[j] = false
					}
				}
			}
			//если флагов на поиск строк до и после совпадения нет просто
			//добавляем совпадение и ставим значение true, в качестве показателя, что совпадение нашлось
			if _, ok := strIndex[i]; !ok {
				id = append(id, i)
			}
			strIndex[i] = true
		}
		//увеличиваем счетчик для перехода к следующей строке
		i = i + 1
	}
	//если был указан флаг на показание совпадений, возвращаем значение совпадений
	if flag.Count {
		return fmt.Sprintf("%v совпадений", counter), nil
	}
	//сортируем айдишники найденных совпадений
	sort.Ints(id)

	var result []string

	index := ""
	//если не был указан флаг, который исключает
	if !flag.Invert {
		match := ""
		//проходим по всем айдишникам
		for _, j := range id {
			//если айди относится к строке файла
			if j > -1 && j < len(allStr) {
				//если по айди нашлось совпадение, пишем +
				if strIndex[j] {
					match = "+ "
					//если у айди значение false - ничего не добавляем
				} else {
					match = " "
				}
				//если был указан флаг на печать номера строки совпадения
				if flag.LineNum {
					//в переменную индекса добавляем номер строки
					index = strconv.Itoa(j+1) + " "
				}
				//в качестве результата пишем все подходящие с учетом флагов строки и указываем их индексы
				result = append(result, fmt.Sprintf("%v%v%v", index, match, allStr[j]))
			}
		}
		//если был указан флаг, который исключает
	} else {
		for ind, j := range allStr {
			//если строка не нашлась - она проходит условие
			if _, ok := strIndex[ind]; !ok {
				//если был указан флаг на номер строки
				if flag.LineNum {
					//в переменную индекса добавляем номер строки
					index = strconv.Itoa(ind+1) + " "
				}
				//в качестве результата пишем все не совпавшие строки и указываем их индексы
				result = append(result, fmt.Sprintf("%v%v", index, j))
			}
		}
	}
	return strings.Join(result, "\n"), nil
}

func main() {
	//объявим возможные флаги
	flagsP := &flags{}
	flag.IntVar(&flagsP.After, "A", 0, "количество линей после совпадения")
	flag.IntVar(&flagsP.Before, "B", 0, "количество линей до совпадения")
	flag.IntVar(&flagsP.Context, "C", 0, "количество линей рядом с совпадением")
	flag.BoolVar(&flagsP.Count, "c", false, "количество совпадений")
	flag.BoolVar(&flagsP.IgnoreCase, "i", false, "игнорирование регистра")
	flag.BoolVar(&flagsP.Invert, "v", false, "исключает совпадение")
	flag.BoolVar(&flagsP.Fixed, "F", false, "что то делает")
	flag.BoolVar(&flagsP.LineNum, "n", false, "выводит количество линий")
	flag.Parse()
	//пример запуска кода
	//go run ./main.go -c  "abc" test.txt
	//если количество аргументов при запуске программы меньше двух(строка и файл для проверки) - кидаем ошибку
	if flag.NArg() < 2 {
		fmt.Println("err")
		return
	}
	//открываем файл, который был задан при запуске программы
	reader, err := os.Open(flag.Arg(1))
	//если не открылся - кидаем ошибку
	if err != nil {
		fmt.Println("grep err:", err.Error())
		return
	}
	//считываем строку для работы "grep'а"
	str := flag.Arg(0)
	fmt.Println(str)
	//запускаем grep, передаем файл, строку и все флаги
	res, err := grep(reader, str, flagsP)
	if err != nil {
		fmt.Println("grep err:", err.Error())
		return
	}
	fmt.Println(res)
}
