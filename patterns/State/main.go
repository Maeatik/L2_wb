package main

import "fmt"

func main() {
	//Создаем телефон. По умолчанию находится в состоянии вибрации
	mobile := NewMobileAlert()

	//записываем в результат действие состояния, в котором находится телефон
	result := mobile.Alert()
	result += mobile.Alert()

	//меняем состояние
	mobile.SetState(&MobileAlertSong{})
	result += mobile.Alert()

	mobile.SetState(&MobileAlertVibration{})
	result += mobile.Alert()

	//показываем результат записи действия состояний и их смены
	fmt.Println(result)
}
