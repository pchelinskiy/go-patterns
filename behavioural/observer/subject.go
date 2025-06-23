package observer

type Subject interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	notifyAll()
}
