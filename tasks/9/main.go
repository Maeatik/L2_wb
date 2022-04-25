package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var err error
	//если не указана хотя бы ссылка на сайт
	if len(os.Args) < 2 {
		log.Printf("usage: %s usl [us fname]")
	}

	//если указано 2 аргумента - ссылка и название файла, куда надо сохранить
	if len(os.Args) == 3 {
		err = wget(os.Args[1], os.Args[2])
	} else {
		//если указано 1 аргумент - ссылка на сайт для скачивания
		err = wget(os.Args[1], "")
	}
	//если произошла ошибка - грустим
	if err != nil {
		log.Printf("err: %+v\n", err)
		os.Exit(1)
	}
}

//Если ссылка не содержит в названии двоеточие - то скорее всего в ней нет протокола передачи, поэтому - добавляем
func wget(link string, output string) error {
	if !strings.Contains(link, ":") {
		link = "http://" + link
	}
	//создаем запрос по полученной ссылке
	request, err := http.NewRequest("GET", link, nil)

	if err != nil {
		return err
	}

	fname := ""
	//если было задано имя конечного файла - используем его

	if output != "" {
		fname = output
	}
	//создаем клиента
	client := &http.Client{}
	//клиент отправляет запрос, получаем ответ
	response, err := client.Do(request)

	if err != nil {
		return err
	}
	//отложено закрываем чтение
	defer response.Body.Close()

	//показываем статус запроса
	log.Printf("response status: %s\n", response.Status)

	//если не было указано имя файла - ставим его по умолчанию index.html
	if fname == "" {
		fname = "index.html"
	}

	var writer io.Writer
	var outputFile *os.File
	//открываем файл на запись, с установленными флагами и числом доступа
	outputFile, err = os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	writer = outputFile
	buffer := make([]byte, 2048)
	//сохраняем страницу
	for {
		//записываем скаченные данные в буфер пока в нем есть место, либо пока есть данные
		n, err := response.Body.Read(buffer)
		fmt.Println(n)

		if err != nil && err != io.EOF {
			return err
		}
		//если длина записанных данных в буфер = 0, то мы ничего не прочитали, значит страница скачена
		if n == 0 {
			break
		}
		//записываем из буфера полученные данные
		if _, err := writer.Write(buffer[:n]); err != nil {
			return err
		}
	}

	if err != nil {
		return err
	} else {
		//если все ОК, говорим, что скачали файл и указываем его название
		fmt.Printf("\n %v saved", fname)
	}
	//если мы что-то записывали в файл, закрываем его
	if outputFile != nil {
		err = outputFile.Close()
	}
	return err
}
