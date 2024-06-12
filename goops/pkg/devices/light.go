package devices

import (
	"fmt"
	"goops/pkg/utils"
)

func LightLogger() {
	utils.Writer("This is a light ack log")
}

type Light struct {
	brand     string // brand of the light
	isOn      bool   // status whether on or off
	usageTime int32  // usage time in hours
}

// func receiver name() (return-type) {} .... <syntax>

// implementing its own methods
func (d *Light) CheckOnStatus() bool {
	if d.isOn {
		utils.Writer("Light is on")
	} else {
		utils.Writer("Light is off")
	}

	return d.isOn
}

func (d *Light) Setter(brand string, isOn bool, usageTime int32) {
	d.brand = brand
	d.isOn = isOn
	d.usageTime = usageTime
}

// implementing the interface methods
// UsageTimeTracker() int32 // in hours
// TurnOff() string
// TurnOn() string
func (d *Light) UsageTimeTracker() int32 {
	// just increase usageTime hrs by 10, while prompting evrytime
	d.usageTime += 10
	var statusUpdt string = fmt.Sprintf("Light has been on for : %d hours", d.usageTime)
	utils.Writer(statusUpdt)
	return d.usageTime
}

func (d *Light) TurnOff() {
	if d.isOn {
		d.isOn = false
	}
}

func (d *Light) TurnOn() {
	if !d.isOn {
		d.isOn = true
	}
}
