package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rsoaresgouveia/raspberry-bush/core/rest"
)

func main() {
	println("Starting MUX router")
	router := mux.NewRouter()

	router.HandleFunc("/toogle", rest.ToogleSignalInGPIO).Methods("POST")
	router.HandleFunc("/", rest.TestConnection).Methods("GET")
	router.HandleFunc("/ledRgb", rest.RGBcontroller).Methods("POST")

	log.Fatal(http.ListenAndServe(":7777", router))
}
