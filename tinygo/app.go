package main

import (
	"strconv"

	. "wasm-wifi/arduino"
	serial "wasm-wifi/serial"
	wifi "wasm-wifi/wifi"
)

var ledPin = GetPinLED()
var ledState = false
var lastMillis uint
var deviceID = ""

const (
	blinkInterval = 1000
	ssid          = "YOUR_SSID"
	password      = "YOUR_PASSWORD"
)

func setup() {
	PinMode(ledPin, OUTPUT)

	DigitalWrite(ledPin, LOW)
	lastMillis = Millis()
	serial.Println("Hello from TinyGo 🐹")
	deviceID = GetChipID()
}

func connect() {
	if wifi.Status() == wifi.WL_CONNECTED {
		return
	}

	wifi.Connect(ssid, password)
	serial.Println("Connecting")
	attempts := 0
	for wifi.Status() != wifi.WL_CONNECTED {
		Delay(500)
		serial.Print(".")
		attempts++
		if attempts >= 10 {
			serial.Println("Failed to connect!")
			return
		}
	}
	serial.Println("Connected!")
	serial.Println(wifi.LocalIp())
}

func loop() {
	connect()
	currentMillis := Millis()
	if (currentMillis - lastMillis) >= blinkInterval {
		connected := wifi.Status() == wifi.WL_CONNECTED
		localIP := wifi.LocalIp()

		serial.Print("[")
		serial.Print(strconv.FormatUint(uint64(currentMillis), 10))
		serial.Print("] [" + deviceID)
		serial.Print("] [connected : ")
		serial.Print(strconv.FormatBool(connected))
		serial.Print("] [" + localIP + "] ")
		serial.Println("TinyGo 🐹")

		ledState = !ledState
		if ledState {
			DigitalWrite(ledPin, HIGH)
		} else {
			DigitalWrite(ledPin, LOW)
		}

		lastMillis = Millis()
	}
	Delay(10)
}

/*
 * Entry point
 */
func main() {
	setup()
	for {
		loop()
	}
}
