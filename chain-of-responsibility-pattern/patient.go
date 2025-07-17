package chainofresponsibilitypattern

type Patient struct {
	name                    string
	isCompleteReception     bool
	isCompleteMeetDoctor    bool
	isCompleteGetMedication bool
	isCompletePayment       bool
}

func NewPatient(name string) *Patient {
	return &Patient{
		name:                    name,
		isCompleteReception:     false,
		isCompleteMeetDoctor:    false,
		isCompleteGetMedication: false,
		isCompletePayment:       false,
	}
}
