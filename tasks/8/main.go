package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mitchellh/go-ps"
	"log"
	"os"
	"strings"
)

func getPath() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("shell: %v$ ", wd)
}

func main() {
	//создаем сканер и получаем путь текущей директории
	sc := bufio.NewScanner(os.Stdin)
	getPath()
	//считываем команды из консоли
	for sc.Scan() {
		text := sc.Text()
		if text == "exit" {
			return
		}
		//Проверяем является ли то, что ввели в консоль - командой. Если да, то выполняем
		parsePipeline(os.Stdout, text)
		//заново проверяем путь текущей директории
		getPath()
	}

}

func parsePipeline(stdout *os.File, str string) {
	//проверяем строку на команды, которые указаны через |
	commands := strings.Split(str, "|")
	var result string
	var err error
	//выполняем указанные команды(если они есть)
	for _, c := range commands {
		result, err = parseCommand(result, c)
		if err != nil {
			fmt.Fprintln(stdout, err)
		}
	}
	//если команда существует и она выполнилась - показываем результат
	if result != "" {
		fmt.Fprintln(stdout, result)
	}
}

func parseCommand(input string, c string) (string, error) {
	str := strings.Fields(c)
	//если команд нет - кидаем ошибку
	if len(str) == 0 {
		return "", errors.New("пустая команда")
	}
	if input != "{" {
		str = append(str, input)
	}
	//выполняем первую указанную команду
	switch str[0] {
	case "cd":
		//если параметра не 3(2) - сам cd и директория, куда надо перейти - кидаем ошибку
		if len(str) != 3 {
			fmt.Println(str, len(str))
			return "", errors.New("cd: должен быть 1 параметр")
		}
		//переходим в указанную директорию
		if err := os.Chdir(str[1]); err != nil {
			return "", err
		}
	//если указан дополнительный параметр - кидаем ошибку
	case "pwd":
		if len(str) != 2 {
			return "", errors.New("pwd: неиспользованный параметр")
		}
		//находим текущую директорию и возвращаем ее
		result, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return result, nil
	//если параметра не 3(2) - сам echo и сообщение, для повтора - кидаем ошибку
	case "echo":
		if len(str) != 3 {
			return "", errors.New("echo: должен быть 1 параметр")
		}
		return str[1], nil

	case "ps":
		//получаем текущие процессы
		processes, err := ps.Processes()
		if err != nil {
			return "", err
		}

		var strForProc strings.Builder
		strForProc.WriteString("\tPID\tCMD\n")
		//пишем процессы
		for _, proc := range processes {
			strForProc.WriteString(fmt.Sprintf("\t%v\t%v\n", proc.Pid(), proc.Executable()))
		}
		return strForProc.String(), nil
	}

	return "", nil
}
