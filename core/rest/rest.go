package rest

import (
	"net/http"
	"strconv"

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

func errorHandler(err error) {
	if err != nil {
		println(err.Error())
	}
}
