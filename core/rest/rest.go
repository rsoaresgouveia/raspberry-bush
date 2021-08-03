package rest

import (
	"fmt"
	"net/http"
	"strconv"

	data "github.com/rsoaresgouveia/raspberry-bush/entities"
	"github.com/stianeikeland/go-rpio"
)

func ToogleSignalInGPIO(w http.ResponseWriter, r *http.Request) {

	//Geting a single Query in the url
	gpio := r.URL.Query()["gpio"]

	i, err_str := strconv.Atoi(gpio[0])

	errorHandler(err_str)

	println("Opening gpio port connection")
	err_rpio := rpio.Open()

	errorHandler(err_rpio)

	println("Configuring gpio %b to output a signal", i)
	pin := rpio.Pin(i)

	pin.Output()

	if pin.Read() == 1 {
		println("Toogling signal to LOW")
	} else {
		println("Toogling signal to HIGH")
	}
	pin.Toggle()

	w.WriteHeader(http.StatusOK)

	println("Closing the gpio...")
	defer rpio.Close()

}

func TestConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	println("Testing connecton to the server on a raspberry pi")
}

func RGBcontroller(w http.ResponseWriter, r *http.Request) {
	freq := 64000

	println("Getting hex color from URL")
	query := r.URL.Query()["hex"]
	hex := data.HEX{query[0]}

	fmt.Printf("Hex color: %#s", hex.Value)
	println("Getting pins from URL")

	pinR, e1 := strconv.Atoi(r.URL.Query()["PinR"][0])
	pinG, e2 := strconv.Atoi(r.URL.Query()["PinG"][0])
	pinB, e3 := strconv.Atoi(r.URL.Query()["PinB"][0])

	errorHandler(e1)
	errorHandler(e2)
	errorHandler(e3)

	fmt.Printf("Pins are: r => %#d pin g => %#d b=> %#d\n", pinR, pinG, pinB)

	pinLayout := data.PinRGBlayout{PinR: pinR, PinG: pinG, PinB: pinB}

	red, eer1 := strconv.ParseInt(hex.Value[:2], 16, 32)
	green, ee2 := strconv.ParseInt(hex.Value[2:4], 16, 32)
	blue, ee3 := strconv.ParseInt(hex.Value[4:], 16, 32)

	errorHandler(eer1)
	errorHandler(ee2)
	errorHandler(ee3)

	rgb := data.Color{Red: red, Green: green, Blue: blue}

	err_rpio := rpio.Open()

	defer rpio.Close()

	errorHandler(err_rpio)

	gpioR := rpio.Pin(pinLayout.PinR)
	gpioG := rpio.Pin(pinLayout.PinG)
	gpioB := rpio.Pin(pinLayout.PinB)

	// pin := rpio.Pin(19)
	// pin.Mode(rpio.Pwm)
	// pin.Freq(64000)
	// pin.DutyCycle(0, 32)

	println("Creating pins frequency in PWM")

	println(rgb.Red)
	println(rgb.Green)
	println(rgb.Blue)
	println(float64(rgb.Red) / 255.0)
	println(uint32(rgb.Green) / 255)
	println(uint32(rgb.Blue) / 255)

	gpioR.Mode(rpio.Pwm)
	gpioR.Freq(freq)
	gpioR.DutyCycle(uint32(rgb.Red)/255, uint32(rgb.Red)/255)

	gpioG.Mode(rpio.Pwm)
	gpioG.Freq(freq)
	gpioG.DutyCycle(uint32(rgb.Green)/255, uint32(rgb.Green)/255)

	gpioB.Mode(rpio.Pwm)
	gpioB.Freq(freq)
	gpioB.DutyCycle(uint32(rgb.Blue)/255, uint32(rgb.Blue)/255)

	w.WriteHeader(http.StatusOK)

	println("Done")
}

func errorHandler(err error) {
	if err != nil {
		println(err.Error())
	}
}
