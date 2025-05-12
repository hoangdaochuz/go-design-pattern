package notifier

type INotifier interface {
	SendNotification(string) error
}
