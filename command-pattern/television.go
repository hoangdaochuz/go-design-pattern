package commandpattern

import "fmt"

type Television struct {
	isDisplaying bool
}

func NewTV(isDisplaying bool) Device {
	return &Television{
		isDisplaying: isDisplaying,
	}
}

func (t *Television) On() {
	t.isDisplaying = true
	fmt.Println("TV is displaying")
}

func (t *Television) Off() {
	t.isDisplaying = false
	fmt.Println("TV is off")
}
