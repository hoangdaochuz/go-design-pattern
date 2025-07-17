package chainofresponsibilitypattern

import "fmt"

type PaymentDepartment struct {
	next Department
}

func (p *PaymentDepartment) SetNext(next Department) {
	p.next = next
}

func (p *PaymentDepartment) Execute(patient Patient) {
	if patient.isCompletePayment {
		if p.next != nil {
			p.next.Execute(patient)
		}
	}
	patient.isCompletePayment = true
	fmt.Println("Completed payment")
	if p.next != nil {
		p.next.Execute(patient)
	}
}
