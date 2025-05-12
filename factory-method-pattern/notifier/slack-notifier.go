package notifier

import "fmt"

type SlackNotifier struct {
}

func (slack *SlackNotifier) SendNotification(message string) error {

	fmt.Println("Slack sent notification successfully")
	return nil
}
