package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

func NewSphero(name, port string) Sphero {
	return &sphero_struct{Name: "Gundam", Port: port}
}

type Sphero interface {
	Start()
	Stop()
	SetRGB(r, g, b uint8)
	SetSpin(speed uint8, degree uint16)
	SetHeading(degree uint16)
}

type sphero_struct struct {
	Name       string
	Port       string
	device     *sphero.SpheroDriver
	connection *sphero.SpheroAdaptor
}

func (s *sphero_struct) Start() {
	gbot := gobot.NewGobot()

	adaptor := sphero.NewSpheroAdaptor("sphero", s.Port)
	driver := sphero.NewSpheroDriver(adaptor, "sphero")

	s.connection = adaptor
	s.device = driver

	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
	)

	gbot.AddRobot(robot)

	gbot.Start()
}

func (s *sphero_struct) Stop() {
	s.connection.Disconnect()
}

func (s *sphero_struct) SetRGB(r, g, b uint8) {
	s.device.SetRGB(r, g, b)
}

func (s *sphero_struct) SetSpin(speed uint8, degree uint16) {
	fmt.Printf("Roll by speed = %d, degree = %d \n", speed, degree)
	s.device.Roll(speed, degree)
}

func (s *sphero_struct) SetHeading(degree uint16) {
	if degree > 360 || degree < 0 {
		degree = 0
	}
	fmt.Printf("Set heading degree = %d \n", degree)
	s.device.SetHeading(degree)
}
