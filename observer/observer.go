package observer

type Observer interface {
	Update(temp float64, humidity float64, pressure float64)
}

// Subject interface used to decouple from displays
type Subject interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers()
}
