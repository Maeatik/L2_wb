package main

func main() {
	//создаем объект, которым будем управлять
	projector := &projector{}

	//создаем команды и говорим, каким устройством они будут управлять
	onCommand := &onCommand{device: projector}

	offCommand := &offCommand{device: projector}

	//создаем пульт и задаем ему команду
	Remote := &remoteController{command: onCommand}

	//жмякаем на кнопку
	Remote.pressButton()

	Remote = &remoteController{command: offCommand}

	Remote.pressButton()
	Remote.pressButton()
}
