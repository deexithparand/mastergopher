package main

import (
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
}
