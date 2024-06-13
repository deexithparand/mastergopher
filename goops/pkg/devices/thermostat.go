package devices

import (
	"fmt"
	"goops/pkg/utils"
)

func ThermostatLogger() {
	utils.Writer("This is a thermostat ack log")
}

// struct embedding
type Thermostat struct {
	SmartDevices
	brand string
}

func (t *Thermostat) ConnectDevice() {
	utils.Writer(fmt.Sprintf("thermostat logs : connected to %s device via nic of brand %s", t.brand, t.NIC))
}

// setter for thermostat
func (t *Thermostat) Setter(thermostatBrand string, nicbrand string) error {
	t.brand = thermostatBrand

	// setter also calls the createNic
	err := t.CreateNic(nicbrand)
	if err != nil {
		return err
	}

	return nil
}
