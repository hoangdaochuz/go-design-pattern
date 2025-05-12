package notifactory

import "git/go-design-pattern/factory-method-pattern/notifier"

type INotifierFactory interface {
	CreateNotifier() notifier.INotifier
}
