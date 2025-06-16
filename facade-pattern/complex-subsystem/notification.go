package complexsubsystem

import "fmt"

type Notification struct {
}

func NewNotification() *Notification {
	return &Notification{}
}

func (notification *Notification) SendNotification(message string, receiver string) error {
	fmt.Printf("Sending notification to %s: %s\n", receiver, message)
	return nil
}
