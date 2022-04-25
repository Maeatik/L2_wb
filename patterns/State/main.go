package main

import "fmt"

func main() {

	mobile := NewMobileAlert()

	result := mobile.Alert()
	result += mobile.Alert()

	mobile.SetState(&MobileAlertSong{})
	result += mobile.Alert()

	mobile.SetState(&MobileAlertVibration{})
	result += mobile.Alert()

	fmt.Println(result)
}
