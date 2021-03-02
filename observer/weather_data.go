package observer

import "fmt"

type WeatherData struct {
	observers []Observer
	temp      float64
	humidity  float64
	pressure  float64
}

func NewWeatherData() *WeatherData {
	wd := WeatherData{}
	wd.observers = make([]Observer, 0)
	return &wd
}

func (wd *WeatherData) RegisterObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) NotifyObservers() {
	for _, obs := range wd.observers {
		obs.Update(wd.temp, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) RemoveObserver(o Observer) {
	for idx, obs := range wd.observers {
		if obs == o {
			fmt.Printf("Remove %v from observers", o)
			wd.observers = remove(wd.observers, idx)
		}
	}
}

func remove(s []Observer, i int) []Observer {
	return append(s[:i], s[i+1:]...)
}

func (wd *WeatherData) MeasurementsChanged() {
	wd.NotifyObservers()
}

func (wd *WeatherData) SetMeasurements(temp float64, humidity float64, pressure float64) {
	wd.temp = temp
	wd.humidity = humidity
	wd.pressure = pressure
	wd.MeasurementsChanged()
}
