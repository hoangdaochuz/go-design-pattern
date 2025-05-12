package notifactory

import "git/go-design-pattern/factory-method-pattern/notifier"

type MailNotifierFactory struct{}

func (mailNotiFactory *MailNotifierFactory) CreateNotifier() notifier.INotifier {
	return &notifier.MailNotifier{}
}
