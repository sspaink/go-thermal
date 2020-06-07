package main

import (
	"machine"
)

// LED defines a led with state
type LED struct {
	l     machine.Pin
	state []bool
}

func (led *LED) blink(i int) {
	if led.state[i] {
		led.l.High()
	} else {
		led.l.Low()
	}
}

// RunningLED defines the sequence leds for the three front lights
type RunningLED struct {
	onLed    machine.Pin
	ledOne   LED
	ledTwo   LED
	ledThree LED

	currentState int
}

// NewRunningLED returns a warningLED
func NewRunningLED() *RunningLED {

	var w RunningLED

	w.onLed = machine.Pin(13)
	w.ledOne.l = machine.Pin(12)
	w.ledTwo.l = machine.Pin(11)
	w.ledThree.l = machine.Pin(10)

	w.onLed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	w.ledOne.l.Configure(machine.PinConfig{Mode: machine.PinOutput})
	w.ledTwo.l.Configure(machine.PinConfig{Mode: machine.PinOutput})
	w.ledThree.l.Configure(machine.PinConfig{Mode: machine.PinOutput})

	w.ledOne.state = []bool{true, true, true, true, false, false, false, true, false}
	w.ledTwo.state = []bool{true, false, true, false, true, true, false, true, false}
	w.ledThree.state = []bool{true, false, false, true, true, false, true, false, false}

	w.currentState = 0

	return &w

}

func (w *RunningLED) blink(rollerSwitch machine.Pin) {

	if !rollerSwitch.Get() {
		w.onLed.Low()
		w.ledOne.l.Low()
		w.ledTwo.l.Low()
		w.ledThree.l.Low()
		return
	}

	w.ledOne.blink(w.currentState)
	w.ledTwo.blink(w.currentState)
	w.ledThree.blink(w.currentState)

	w.currentState++
	if w.currentState > 8 {
		w.currentState = 0
	}

}
