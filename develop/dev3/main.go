package main

func Sort(inputData, exportData string, k int, n, r bool) {
	//создаем нового сортера, который знает откуда читать и куда записывать результат
	sorter1 := NewSorter(inputData, exportData)
	//читаем файл
	sorter1.readFile()
	//если переменная К не пустая, создаем матрицу по количству указанных столбцов и сортируем
	if k > 0 {
		sorter1.createMatrix()
		sorter1.pillarsSort(k, n, r)
		//если к=0, работаем без матрицы
	} else {
		sorter1.listSort(n, r)
	}
	//после обработки - создаем файл
	sorter1.createExportFile(k)
}
func main() {
	Sort("sort.txt", "324.txt", 2, false, false)
}
