package devices

import (
	"errors"
	"fmt"
)

// generalizes and lists the common features based on category

// common features of smart device
type DeviceActions interface {
	UsageTimeTracker() int32 // in hours
	TurnOff()
	TurnOn()
}

// interface to connect to any network device
type DeviceConfig interface {
	ConnectDevice()
}

// now implement these features in different smart devices in different modules

// To display struct composition
// also can check is functions can be reuses
type SmartDevices struct {
	NIC string
}

// setter
func (s *SmartDevices) CreateNic(brand string) error {

	if brand == "" {
		return errors.New("Invalid Brand name, its empty")
	}

	s.NIC = brand

	return nil
}

func (s *SmartDevices) ConnectNic() string {
	return fmt.Sprintf("Connected to %s nic card", s.NIC)
}
