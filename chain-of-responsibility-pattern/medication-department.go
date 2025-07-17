package chainofresponsibilitypattern

import "fmt"

type MedicationDepartment struct {
	next Department
}

func (m *MedicationDepartment) SetNext(next Department) {
	m.next = next
}

func (m *MedicationDepartment) Execute(patient Patient) {
	if patient.isCompleteGetMedication {
		if m.next != nil {
			m.next.Execute(patient)
		}
		return
	}
	patient.isCompleteGetMedication = true
	fmt.Println("Completed get medication")
	if m.next != nil {
		m.next.Execute(patient)
	}
}
