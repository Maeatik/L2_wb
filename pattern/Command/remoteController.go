package main

type remoteController struct {
	command command
}

func (r *remoteController) pressButton() {
	r.command.execute()
}
