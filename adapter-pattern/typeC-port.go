package adapterpattern

import "fmt"

type TypeCCharger struct {
}

func NewTypeCCharger() ChargeBattery {
	return &TypeCCharger{}
}

func (typeCPort *TypeCCharger) Charge() error {
	fmt.Println("Charging phone with typeC port")
	return nil
}
