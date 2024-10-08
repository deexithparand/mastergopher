package devices

import (
	"fmt"
	"goops/pkg/utils"
)

func CameraLogger() {
	utils.Writer("This is a security camera ack log")
}

type Camera struct {
	SmartDevices // struct embedding device.go

	brand     string // brand of the light
	isOn      bool   // status whether on or off
	usageTime int32  // usage time in hours
	lensSize  int32  // in mm
}

// implementing an interface to promote polymorphism
func (c *Camera) ConnectDevice() {
	utils.Writer(fmt.Sprintf("Camera logs : connected to %s device via nic of brand %s", c.brand, c.NIC))
}

// setter to follow encapsulation oops
func (c *Camera) Setter(
	brand string,
	isOn bool,
	usageTime int32,
	lensSize int32,
	nicName string,
) error {
	c.brand = brand
	c.isOn = isOn
	c.usageTime = usageTime
	c.lensSize = lensSize

	// setter also calls the createNic
	err := c.CreateNic(nicName)
	if err != nil {
		return err
	}

	return nil
}

// implementating the interface smart devices later
