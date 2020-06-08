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
	ledOne   LED
	ledTwo   LED
	ledThree LED

	currentState int
}

// NewRunningLED returns a warningLED
func NewRunningLED() *RunningLED {

	var r RunningLED

	r.ledOne.l = machine.Pin(12)
	r.ledTwo.l = machine.Pin(11)
	r.ledThree.l = machine.Pin(10)

	r.ledOne.l.Configure(machine.PinConfig{Mode: machine.PinOutput})
	r.ledTwo.l.Configure(machine.PinConfig{Mode: machine.PinOutput})
	r.ledThree.l.Configure(machine.PinConfig{Mode: machine.PinOutput})

	r.ledOne.state = []bool{true, true, true, true, false, false, false, true, false}
	r.ledTwo.state = []bool{true, false, true, false, true, true, false, true, false}
	r.ledThree.state = []bool{true, false, false, true, true, false, true, false, false}

	r.currentState = 0

	return &r

}

// TurnOff will set all the LED to low
func (r *RunningLED) TurnOff() {
	r.ledOne.l.Low()
	r.ledTwo.l.Low()
	r.ledThree.l.Low()
}

// Blink will change the state of the warning LED's
func (r *RunningLED) Blink() {

	r.ledOne.blink(r.currentState)
	r.ledTwo.blink(r.currentState)
	r.ledThree.blink(r.currentState)

	r.currentState++
	if r.currentState > 8 {
		r.currentState = 0
	}

}
