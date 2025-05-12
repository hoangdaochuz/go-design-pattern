package notifier

import "fmt"

type MailNotifier struct {
}

func (*MailNotifier) SendNotification(message string) error {

	fmt.Println("Mail sent notification successfully")

	return nil
}
