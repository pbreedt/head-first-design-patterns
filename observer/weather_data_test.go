package observer

import (
	"fmt"
	"testing"
)

type MockObserver struct{}

func (t MockObserver) Update(temp float64, humidity float64, pressure float64) {
	fmt.Printf("MockObserver updated with temp:%f, humidity:%f, pressure:%f\n", temp, humidity, pressure)
}

func TestRegisterObserver(t *testing.T) {
	wd := NewWeatherData()
	obs := MockObserver{}
	obsbefore := len(wd.observers)
	wd.RegisterObserver(obs)
	obsafer := len(wd.observers)
	if obsafer != obsbefore+1 {
		t.Errorf("want %d, got %d", obsbefore+1, obsafer)
	}
}

func TestRemoveObserver(t *testing.T) {
	wd := NewWeatherData()
	obs := MockObserver{}
	obsbefore := len(wd.observers)
	wd.RegisterObserver(obs)
	obsaferadd := len(wd.observers)
	if obsaferadd != obsbefore+1 {
		t.Errorf("want %d, got %d", obsbefore+1, obsaferadd)
	}

	//remove
	wd.RemoveObserver(obs)
	obsaferrem := len(wd.observers)
	if obsaferrem != obsbefore {
		t.Errorf("want %d, got %d", obsbefore, obsaferrem)
	}
}

func TestWeatherData(t *testing.T) {
	wd := NewWeatherData()
	_ = NewCurrentConditionsDisplay(wd)
	_ = NewStatisticsDisplay(wd)

	wd.SetMeasurements(80, 65, 30.4)
	wd.SetMeasurements(82, 70, 29.2)
	wd.SetMeasurements(78, 90, 29.2)
}
