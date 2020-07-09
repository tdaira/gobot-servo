package main

import (
	"fmt"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"net/http"
)

type ServoController struct {
	adaptor *raspi.Adaptor
	servo   *gpio.ServoDriver
}

func NewServoController() *ServoController {
	adaptor := raspi.NewAdaptor()
	servo := gpio.NewServoDriver(adaptor, "12")
	return &ServoController{
		adaptor: adaptor,
		servo:   servo,
	}
}

func (s *ServoController) Open(w http.ResponseWriter, r *http.Request) {
	fmt.Println("open")
	err := s.servo.Move(30)
	if err != nil {
		fmt.Printf("servo move error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (s *ServoController) Close(w http.ResponseWriter, r *http.Request) {
	fmt.Println("close")
	err := s.servo.Move(45)
	if err != nil {
		fmt.Printf("servo move error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	servoController := NewServoController()
	http.HandleFunc("/open", servoController.Open)
	http.HandleFunc("/close", servoController.Close)
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("initialization err: %s\n", err)
}
