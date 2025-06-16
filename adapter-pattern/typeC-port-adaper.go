package adapterpattern

type TypeCChargerAdapter struct {
	lightningCharger *LightningCharger
}

func NewTypeCChargerAdapter(lightningCharger *LightningCharger) ChargeBattery {
	return &TypeCChargerAdapter{
		lightningCharger: lightningCharger,
	}
}

// implement the ChargeBattery interface
func (typeCChargerAdapter *TypeCChargerAdapter) Charge() error {
	// get tilPercent from third party charger
	tilPercent := 80
	// call the ProcessLightningCharge method
	return typeCChargerAdapter.lightningCharger.ProcessLightningCharge(tilPercent)
}
