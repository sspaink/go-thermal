package main

import (
	"machine"
	"time"
)

// Dfminiplayer defines the components for playing mp3
// Reference: https://wiki.dfrobot.com/DFPlayer_Mini_SKU_DFR0299
type Dfminiplayer struct {
	busy           machine.Pin
	hardwareSerial machine.UART
	send           []byte
}

const (
	// CMD inidicates the specific function such as play, pause, etc.
	CMD = 3
	// PARAMETER is extra information for the CMD
	PARAMETER = 6
)

// NewDfminiplayer sets up new mp3 player
func NewDfminiplayer() *Dfminiplayer {

	var df Dfminiplayer

	df.hardwareSerial = machine.UART0
	df.hardwareSerial.Configure(machine.UARTConfig{BaudRate: 9600, TX: 1, RX: 0})

	df.busy = machine.Pin(2)
	df.busy.Configure(machine.PinConfig{Mode: machine.PinInput})

	df.resetSend()

	df.Reset()
	time.Sleep(time.Millisecond * 1500)

	return &df
}

func (df *Dfminiplayer) resetSend() {
	df.send = []byte{0x7E, 0xFF, 6, 0, 0, 0, 0, 0xEF}
}

// Reset will reset the module
func (df *Dfminiplayer) Reset() {
	df.send[CMD] = 0x0C
	df.hardwareSerial.Write(df.send)
	df.resetSend()
}

// Volume will set the volume between 0-30
func (df *Dfminiplayer) Volume(v byte) {
	if v > 30 {
		v = 30
	}
	// sending = []byte{0x7E, 0xFF, 6, 6, 1, 0, 0x0A, 0xFE, 0xEA, 0xEF}
	df.send[CMD] = 6
	df.send[PARAMETER] = v
	df.hardwareSerial.Write(df.send)
	df.resetSend()
}

// Play will tell dfminiplayer to play the file number provided
func (df *Dfminiplayer) Play(file byte) {
	df.send[CMD] = 3
	df.send[PARAMETER] = file
	df.hardwareSerial.Write(df.send)
	df.resetSend()
}

// Pause will tell dfminiplayer to pause
func (df *Dfminiplayer) Pause() {
	df.send[CMD] = 0x0E
	df.hardwareSerial.Write(df.send)
	df.resetSend()
}
