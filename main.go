package main

import (
	"machine"
	"time"
)

// ThermalDetonator defines the components needed for the prop
type ThermalDetonator struct {
	onLED        machine.Pin
	tiltSensor   machine.Pin
	runningLED   *RunningLED
	dfMiniPlayer *Dfminiplayer
	rollerSwitch machine.Pin
	prevSwitch   bool
}

// NewThermalDetonator setups the thermal detonator prop
func NewThermalDetonator() *ThermalDetonator {
	var thermDet ThermalDetonator

	thermDet.tiltSensor = machine.Pin(9)
	thermDet.tiltSensor.Configure(machine.PinConfig{Mode: machine.PinInput})

	thermDet.runningLED = NewRunningLED()

	thermDet.dfMiniPlayer = NewDfminiplayer()
	thermDet.dfMiniPlayer.Volume(30)

	thermDet.rollerSwitch = machine.Pin(7)
	thermDet.rollerSwitch.Configure(machine.PinConfig{Mode: machine.PinInput})

	thermDet.prevSwitch = thermDet.rollerSwitch.Get()

	thermDet.onLED = machine.Pin(13)
	thermDet.onLED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return &thermDet
}

func (thermDet *ThermalDetonator) turnOffLights() {
	thermDet.onLED.Low()
	thermDet.runningLED.TurnOff()
}

func (thermDet *ThermalDetonator) loop() {

	exploded := false
	started := thermDet.rollerSwitch.Get()

	for {

		// Check if the rollerswitch has been pressed
		if thermDet.prevSwitch != thermDet.rollerSwitch.Get() {

			exploded = false

			thermDet.dfMiniPlayer.Pause()
			thermDet.dfMiniPlayer.Play(3)

			thermDet.prevSwitch = thermDet.rollerSwitch.Get()

			if thermDet.rollerSwitch.Get() {
				started = true
			} else {
				thermDet.turnOffLights()
				started = false
			}
		} else {

			if !exploded && started && thermDet.rollerSwitch.Get() {
				if thermDet.dfMiniPlayer.busy.Get() {
					thermDet.dfMiniPlayer.Play(2)
				}

				if thermDet.tiltSensor.Get() {
					thermDet.dfMiniPlayer.Pause()
					thermDet.dfMiniPlayer.Play(1)
					thermDet.runningLED.TurnOff()
					thermDet.runningLED.Exploding()
					exploded = true
					thermDet.prevSwitch = thermDet.rollerSwitch.Get()
					continue
				}
			}

		}

		if started && thermDet.rollerSwitch.Get() {
			if !exploded {
				thermDet.onLED.High()
				thermDet.runningLED.Blink()
				time.Sleep(time.Second)
			} else if exploded {
				if thermDet.dfMiniPlayer.busy.Get() {
					thermDet.turnOffLights()
				} else {
					thermDet.onLED.Set(!thermDet.onLED.Get())
					thermDet.runningLED.Exploding()
					time.Sleep(time.Second / 2)
				}

			}
		}

	}
}

func main() {
	thermDet := NewThermalDetonator()
	thermDet.loop()
}
