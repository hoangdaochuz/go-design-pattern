package adapterpattern

import "fmt"

type LightningCharger struct {
}

func NewLightningCharger() *LightningCharger {
	return &LightningCharger{}
}

func (lightNingCharger *LightningCharger) ProcessLightningCharge(tilPercent int) error {
	fmt.Println("Charging phone with lightning port")
	return nil
}
