package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

//структура флагов
type flags struct {
	column     int
	anotherSep string
	separator  bool
}

//
func isFlag(flag_str string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flag_str {
			found = true
		}
	})
	return found
}

//проверяем был ли введен обязательный флаг f и его аргумент
func input(flagP *flags) ([]string, error) {
	if !isFlag("f") {
		return nil, fmt.Errorf("%s: а где твой аругмент F -_-", os.Args[0])
	} else if flagP.column <= 0 {
		return nil, fmt.Errorf("%n: ну указал ты F, а где число то?", os.Args[0])
	}
	//считываем строку для CUT'а
	str := make([]string, 0)
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println(1)
	for sc.Scan() {
		text := sc.Text()
		if text == "" {
			break
		}
		str = append(str, sc.Text())
	}

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("error of reading: %s", err)
	}
	return str, nil

}

func cut(flagP *flags, str []string) ([]string, error) {
	//
	strSlice := make([][]string, 0)
	// проверяем каждый элемент введенного массива
	for _, elem := range str {
		//если указан флаг на разделитель
		if flagP.separator {
			//проверяем, был ли задан флаг, на смену разделителя и смотрим какой разделитель поставили. Делим строку по нему
			if strings.Contains(elem, flagP.anotherSep) {
				strSlice = append(strSlice, strings.Split(elem, flagP.anotherSep))
			}
			//если флаг не указан режем строку по пробелам
		} else {
			strSlice = append(strSlice, strings.Split(elem, " "))
		}
	}
	//создаем слайс с разрезанной строкой, задаем четкие размеры - по длине слайса strSlice
	result := make([]string, 0, len(strSlice))
	for _, elem := range strSlice {
		//если указанный флаг столбца больше длины строки в столбце - к результату добавляем пустую строку
		if flagP.column > len(elem) {
			result = append(result, "")
		} else {
			//если такой столбец существует - добавляем слово к ответу
			result = append(result, elem[flagP.column-1])
		}
	}
	return result, nil
}
func main() {
	//пример запуска кода
	// go run ./6.go -f=5 -s=true
	//объявим возможные флаги
	flagP := &flags{}
	flag.IntVar(&flagP.column, "f", 0, "количество столбцов")
	flag.StringVar(&flagP.anotherSep, "d", "\t", "определение нового разделителя")
	flag.BoolVar(&flagP.separator, "s", false, "строки только с разделителем")
	flag.Parse()

	str, err := input(flagP)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: [%s]\n", err)
		os.Exit(1)
	}
	//режем строку с учетом указанных флагов
	res, err := cut(flagP, str)

	//если что то не окей - кидаем ошибку
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: [%s]\n", err)
		os.Exit(1)
	}
	// выводим порезанную строку через стдаут
	//(мне было очень лень переводить на английский(а это написать было не лень))
	for _, elem := range res {
		fmt.Fprintln(os.Stdout, elem)
	}

}
