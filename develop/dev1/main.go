package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func PrintCurrentTime() {
	//получаем данные о текущем времени с NTP сервера
	time, err := ntp.Time("ntp5.stratum2.ru")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %v\n", err)
		os.Exit(1)
	}
	//Форматировано выводим текущее время
	fmt.Println(time.Format("Сейчас: 15:04:05 MST"))
}
func main() {
	PrintCurrentTime()
}
