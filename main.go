package main

import (
	"machine"
	"time"
)

// ThermalDetonator defines the components needed for the prop
type ThermalDetonator struct {
	runningLED   *RunningLED
	dfMiniPlayer *Dfminiplayer
	rollerSwitch machine.Pin
	prevSwitch   bool
}

// NewThermalDetonator setups the thermal detonator prop
func NewThermalDetonator() *ThermalDetonator {
	var thermDet ThermalDetonator

	thermDet.runningLED = NewRunningLED()

	thermDet.dfMiniPlayer = NewDfminiplayer()
	thermDet.dfMiniPlayer.Volume(20)

	thermDet.rollerSwitch = machine.Pin(7)
	thermDet.rollerSwitch.Configure(machine.PinConfig{Mode: machine.PinInput})

	thermDet.prevSwitch = thermDet.rollerSwitch.Get()

	return &thermDet
}

func (thermDet *ThermalDetonator) loop() {
	for {
		if thermDet.prevSwitch != thermDet.rollerSwitch.Get() {

			thermDet.dfMiniPlayer.Pause()
			thermDet.dfMiniPlayer.Play(3)

			thermDet.prevSwitch = thermDet.rollerSwitch.Get()
		} else {
			if thermDet.dfMiniPlayer.busy.Get() && thermDet.rollerSwitch.Get() {
				thermDet.dfMiniPlayer.Play(2)
			}
		}

		thermDet.runningLED.blink(thermDet.rollerSwitch)

		time.Sleep(time.Second)
	}
}

func main() {
	thermDet := NewThermalDetonator()
	thermDet.loop()
}
