package main

// MobileAlertVibration implements vibration alert
type MobileAlertVibration struct {
}

// Alert returns a alert string
func (a *MobileAlertVibration) Alert() string {
	return "Брр-брр-брр"
}
