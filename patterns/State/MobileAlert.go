package main

// MobileAlert implements an alert depending on its state.
type MobileAlert struct {
	state State
}

// Alert returns a alert string
func (a *MobileAlert) Alert() string {
	return " " + a.state.Alert()
}

// SetState changes state
func (a *MobileAlert) SetState(state State) {
	a.state = state
}

func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}
