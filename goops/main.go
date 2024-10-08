package main

import (
	"fmt"
	"goops/pkg/devices"
	"goops/pkg/utils"
)

func main() {
	utils.Writer("smart-home-systems project started ... ")
	// just for checking the package setup
	// devices.LightLogger()
	// devices.CameraLogger()
	// devices.ThermostatLogger()

	// 1. encapsulation
	// ... having private fields/methods, accessible via public methods
	// ... light.go
	// var bulb devices.Light

	// bulb.Setter("philips", false, 0)
	// bulb.CheckOnStatus()

	// 2. composition
	// ... implemented using camera.go
	// var camera devices.
	// var cam devices.Camera
	// err := cam.Setter("Canon", false, 0, 50, "Cisco")
	// if err != nil {
	// 	utils.Writer(err)
	// } else {
	// 	fmt.Println("Camera values set successfully!")
	// }

	// utils.Writer(cam.ConnectNic())
	// utils.Writer(cam)

	// 3.polymorphism
	// proper usage of interface, where multiple struct types can use its methods
	var seccam devices.Camera
	err := seccam.Setter("Nikon", false, 0, 50, "Verizon")
	if err != nil {
		utils.Writer(err)
	} else {
		fmt.Println("Camera values set successfully!")
	}
	seccam.ConnectDevice()

	var tms devices.Thermostat
	err = tms.Setter("Honeywell", "CISCO")
	if err != nil {
		utils.Writer(err)
	} else {
		fmt.Println("thermostat values set successfully!")
	}
	tms.ConnectDevice()

	// 4.abstraction
	// hiding the complex information, using functions from packages is one example of it

}
