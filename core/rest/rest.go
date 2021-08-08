package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	data "github.com/rsoaresgouveia/raspberry-bush/entities"
	"github.com/stianeikeland/go-rpio"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
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
	RGBLinker := data.RGBLinker{}

	err := json.NewDecoder(r.Body).Decode(&RGBLinker)

	errorHandler(err)

	fmt.Printf("Pins are: r => %d pin g => %d b=> %d\n", RGBLinker.PinRGBlayout.PinR, RGBLinker.PinRGBlayout.PinG, RGBLinker.PinRGBlayout.PinB)

	err_rpio := rpio.Open()

	defer rpio.Close()

	errorHandler(err_rpio)

	gpioR := rpio.Pin(RGBLinker.PinRGBlayout.PinR)
	gpioG := rpio.Pin(RGBLinker.PinRGBlayout.PinG)
	gpioB := rpio.Pin(RGBLinker.PinRGBlayout.PinB)

	println("Creating pins frequency in PWM")

	// println(RGBLinker.RGB.Red)
	// println(RGBLinker.RGB.Green)
	// println(RGBLinker.RGB.Blue)
	// println(uint32(RGBLinker.RGB.Red) / uint32(RGBLinker.Cycle.PinRcycle))
	// println(uint32(RGBLinker.RGB.Green) / uint32(RGBLinker.Cycle.PinGcycle))
	// println(uint32(RGBLinker.RGB.Blue) / uint32(RGBLinker.Cycle.PinBcycle))

	gpioR.Mode(rpio.Pwm)
	gpioR.DutyCycle(uint32(RGBLinker.RGB.Red), uint32(RGBLinker.Cycle.PinRcycle))
	gpioR.Freq(RGBLinker.Freq)

	gpioG.Mode(rpio.Pwm)
	gpioR.DutyCycle(uint32(RGBLinker.RGB.Green), uint32(RGBLinker.Cycle.PinGcycle))
	gpioG.Freq(RGBLinker.Freq)

	gpioB.Mode(rpio.Pwm)
	gpioR.DutyCycle(uint32(RGBLinker.RGB.Blue), uint32(RGBLinker.Cycle.PinBcycle))
	gpioB.Freq(RGBLinker.Freq)

	w.WriteHeader(http.StatusOK)

	println("Done")
}

func errorHandler(err error) {
	if err != nil {
		println(err.Error())
	}
}

func GobotTest(w http.ResponseWriter, r *http.Request) {
	RGBLinker := data.RGBLinker{}

	err := json.NewDecoder(r.Body).Decode(&RGBLinker)

	errorHandler(err)
	driver := raspi.NewAdaptor()
	rgbDriver := gpio.NewRgbLedDriver(driver, string(RGBLinker.RGB.Red), string(RGBLinker.RGB.Green), string(RGBLinker.RGB.Blue))

	rgbDriver.Connection().Connect()

	if rgbDriver.State() != false {
		rgbDriver.On()
	} else {
		rgbDriver.Off()
	}
	w.WriteHeader(http.StatusOK)

}
