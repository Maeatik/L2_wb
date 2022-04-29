package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Sorter создаем структуру с файлами и слайсами для сортировок
type Sorter struct {
	inputDataFile  string
	outputDataFile string
	linesForSort   []string
	pillarsForSort [][]string
}

// NewSorter конструтор структуры Sorter
func NewSorter(inputDataFile, outputDataFile string) Sorter {
	return Sorter{
		inputDataFile:  inputDataFile,
		outputDataFile: outputDataFile,
		linesForSort:   make([]string, 0),
		pillarsForSort: make([][]string, 0),
	}
}

//Читаем строки из файла, поданного для обработки
func (s *Sorter) readFile() {
	f, err := os.Open(s.inputDataFile)
	if err != nil {
		log.Fatalf("Ошибка в открытии файла. Ошибка %v\n", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		s.linesForSort = append(s.linesForSort, sc.Text())
	}
}

//Создаем файл, куда запишется результат сортировок
func (s Sorter) createExportFile(k int) {
	f, err := os.Create(s.outputDataFile)
	if err != nil {
		log.Fatalf("Ошибка в создании файла. Ошибка %v\n", err)
	}

	defer f.Close()
	//если к > 0 берем данные из матрицы для сортировки
	if k > 0 {
		for _, str := range s.pillarsForSort {
			f.WriteString(strings.Join(str, " ") + "\n")
		}
		//если к = 0 берем данные из слайса для сортировки
	} else {
		for _, str := range s.linesForSort {
			f.WriteString(str + "\n")
		}
	}
}

func strToInt(str string) (int, bool) {
	digit, err := strconv.Atoi(str)
	if err != nil {
		return 0, false
	}
	return digit, true
}

//Линейная сортировка, без столбцов
func (s *Sorter) listSort(num, reverse bool) {
	//проверяем флаги - -r - reverse, -n  num
	//Если задано сразу два флага
	if reverse && num {
		//используем анонимную функцию для сравнения элементов с целью найти меньший
		sort.Slice(s.linesForSort, func(i, j int) bool {
			//берем два элемента на сортировку
			s_i := s.linesForSort[i]
			s_j := s.linesForSort[j]
			//пытаемся перевести их в числа. Если элемент - строка, ловим это в flag
			digit_i, flag_i := strToInt(s_i)
			digit_j, flag_j := strToInt(s_j)
			//если два элемента числа - проверяем, больше ли первое число, чем второе
			if flag_j && flag_i {
				return digit_i > digit_j
			}
			//если первый элемент - число, а второй - нет - возвращаем ДА
			if flag_i && !flag_j {
				return true
			}
			//Если наоборот - НЕТ
			if !flag_i && flag_j {
				return false
			}

			//если оба элемента - символы - проверяем, меньше ли первый элемент, чем второй
			return digit_i < digit_j
		})
	}
	//Если идет обратная сортровка
	if reverse && !num {
		//используем анонимную функцию
		sort.Slice(s.linesForSort, func(i, j int) bool {
			//проверяем, больше ли первый элемент, чем второй элемент
			return s.linesForSort[i] > s.linesForSort[j]
		})
	}
	//если идет обычная сортировка по числам
	if !reverse && num {
		//тоже самое что в случае с обратной сортировкой по числам, но меняются знаки на
		//противоположные и true и false меняются местами
		sort.Slice(s.linesForSort, func(i, j int) bool {
			s_i := s.linesForSort[i]
			s_j := s.linesForSort[j]

			digit_i, flag_i := strToInt(s_i)
			digit_j, flag_j := strToInt(s_j)

			if flag_j && flag_i {
				return digit_i < digit_j
			}
			if flag_i && !flag_j {
				return false
			}
			if !flag_i && flag_j {
				return true
			}
			return digit_i > digit_j
		})
	}
	//если надо просто отсортировать элементы файла - сортируем
	if !reverse && !num {
		sort.Strings(s.linesForSort)
	}
}

//если был задан флаг по количеству столбцов - создаем матрицу
func (s *Sorter) createMatrix() {
	//берем каждую строку из слайса для обычной сортировки
	for _, val := range s.linesForSort {
		//разделяем строку по пробелам и добавляем в матрицу
		s.pillarsForSort = append(s.pillarsForSort, strings.Split(val, " "))
	}
}

//функция сортировки в матрице
func (s *Sorter) pillarsSort(k int, num, reverse bool) {
	//обратная числовая сорировка
	if reverse && num {
		//анонимная функция
		sort.Slice(s.pillarsForSort, func(i, j int) bool {
			i_len := len(s.pillarsForSort[i])
			j_len := len(s.pillarsForSort[j])
			var flag_i, flag_j bool
			var digit_i, digit_j int
			//если в строках слов больше или равное количество чем указанных столбцов,
			//переводим последнее слово строки в число
			if i_len >= k {
				digit_i, flag_i = strToInt(s.pillarsForSort[i][k-1])
			}
			if j_len >= k {
				digit_j, flag_j = strToInt(s.pillarsForSort[j][k-1])
			}
			//если оба слова оказались числами проверяем, больше ли первое число чем второе
			if flag_j && flag_i {
				return digit_i > digit_j
			}
			//если первое слово - число, а второе нет - возвращаем ДА
			if flag_i && !flag_j {
				return true
			}
			//если первое слово - не число, а второе число - возвращаем НЕТ
			if !flag_i && flag_j {
				return false
			}
			//если оба слова не число, смотрим какая строка меньше
			return strings.Join(s.pillarsForSort[i], " ") < strings.Join(s.pillarsForSort[j], " ")
		})
	}
	//обратная не числовая сортировка
	if reverse && !num {
		sort.Slice(s.pillarsForSort, func(i, j int) bool {
			//получаем длину двух элементов из матрицы
			i_len := len(s.pillarsForSort[i])
			j_len := len(s.pillarsForSort[j])

			//если количество слов в строке больше или равно заданному количеству столбцов
			//проверяем последние элементы из строк, на то меньше ли первая строка чем вторая
			if i_len >= k && j_len >= k {
				return s.pillarsForSort[i][k-1] < s.pillarsForSort[j][k-1]
			}
			//если в первой строке слов меньше чем указанное количество столбцов, а во второй - больше или равно - возвращаем - ДА
			if i_len < k && j_len >= k {
				return true
			}
			//если в обеих строках слов меньше, чем заданное количество столбцов
			if i_len < k && j_len < k {
				//объединяем строки и сравниваем что меньше
				return strings.Join(s.pillarsForSort[i], " ") > strings.Join(s.pillarsForSort[j], " ")
			}
			// если в первой строке слов больше чем столбцов, а во второй меньше возвращаем НЕТ
			return false
		})
	}
	//не обратная сортировка чисел
	if !reverse && num {
		//анонимная функция
		sort.Slice(s.pillarsForSort, func(i, j int) bool {
			//длины двух строк
			i_len := len(s.pillarsForSort[i])
			j_len := len(s.pillarsForSort[j])
			var flag_i, flag_j bool
			var digit_i, digit_j int
			//если в первой или второй строках слов больше или равно количеству заданных столбцов
			//переводим в них
			if i_len >= k {
				digit_i, flag_i = strToInt(s.pillarsForSort[i][k-1])
			}
			if j_len >= k {
				digit_j, flag_j = strToInt(s.pillarsForSort[j][k-1])
			}
			//если оба слова - числа
			//знаки меняются на обратные от обратной сортировки чисел, то же самое с return'ом булевых выражений
			if flag_j && flag_i {
				return digit_i < digit_j
			}
			if flag_i && !flag_j {
				return false
			}
			if !flag_i && flag_j {
				return true
			}

			return strings.Join(s.pillarsForSort[i], " ") > strings.Join(s.pillarsForSort[j], " ")
		})
	}
	//обычная сортировка слов по столбам
	if !reverse && !num {
		sort.Slice(s.pillarsForSort, func(i, j int) bool {
			i_len := len(s.pillarsForSort[i])
			j_len := len(s.pillarsForSort[j])
			//если в обеих строках слов больше, чем указанное число столбов проверяем какая строка больше
			if i_len >= k && j_len >= k {
				fmt.Println(s.pillarsForSort[i][k-1])
				fmt.Println(s.pillarsForSort[j][k-1])
				return s.pillarsForSort[i][k-1] < s.pillarsForSort[j][k-1]
			}
			//если в первой строке слов меньше столбцов, а во второй больше возвращаем ДА
			if i_len < k && j_len >= k {
				return true
			}
			//объединяем строки и сравниваем что меньше
			if i_len < k && j_len < k {
				return strings.Join(s.pillarsForSort[i], " ") < strings.Join(s.pillarsForSort[j], " ")
			}
			//если в первой строке слов больше столбцов, а во второй меньше возвращаем НЕТ
			return false
		})
	}
}
