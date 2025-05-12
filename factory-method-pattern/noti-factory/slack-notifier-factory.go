package notifactory

import "git/go-design-pattern/factory-method-pattern/notifier"

type SlackNotifierFactory struct{}

func (slackNotifierFactory *SlackNotifierFactory) CreateNotifier() notifier.INotifier {
	return &notifier.SlackNotifier{}
}
