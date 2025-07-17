package chainofresponsibilitypattern

import "fmt"

type ReceptionDepartment struct {
	next Department
}

func (r *ReceptionDepartment) SetNext(next Department) {
	r.next = next
}

func (r *ReceptionDepartment) Execute(patient Patient) {
	if patient.isCompleteReception {
		if r.next != nil {
			r.next.Execute(patient)
		}
		return
	}
	patient.isCompleteReception = true
	fmt.Println("Completed reception")
	if r.next != nil {
		r.next.Execute(patient)
	}
}
