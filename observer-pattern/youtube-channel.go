package observerpattern

import "slices"

type YoutubeChannel struct {
	subcribers []Subcriber
	name       string
}

func NewYoutubeChannel(name string) *YoutubeChannel {
	return &YoutubeChannel{
		name:       name,
		subcribers: make([]Subcriber, 0),
	}
}
func (y *YoutubeChannel) Register(subcriber Subcriber) {
	y.subcribers = append(y.subcribers, subcriber)
}

func (y *YoutubeChannel) Unregister(delSubcriber Subcriber) {
	y.subcribers = slices.DeleteFunc(y.subcribers, func(subcriber Subcriber) bool {
		return subcriber.GetId() == delSubcriber.GetId()
	})
}

func (y *YoutubeChannel) Notify(msg string) {
	for _, subcriber := range y.subcribers {
		subcriber.Update(msg)
	}
}

func (y *YoutubeChannel) GetChannelName() string {
	return y.name
}
