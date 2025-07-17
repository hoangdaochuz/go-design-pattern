package chainofresponsibilitypattern

import "fmt"

type DoctorDepartment struct {
	next Department
}

func (d *DoctorDepartment) SetNext(next Department) {
	d.next = next
}

func (d *DoctorDepartment) Execute(patient Patient) {
	if patient.isCompleteMeetDoctor {
		if d.next != nil {
			d.next.Execute(patient)
		}
		return
	}
	patient.isCompleteMeetDoctor = true
	fmt.Println("Completed meet doctor")
	if d.next != nil {
		d.next.Execute(patient)
	}
}
