package main

func main() {
	projector := &projector{}

	onCommand := &onCommand{device: projector}

	offCommand := &offCommand{device: projector}

	onRemote := &remoteController{command: onCommand}

	onRemote.pressButton()

	offRemote := &remoteController{command: offCommand}

	offRemote.pressButton()
	offRemote.pressButton()
}
