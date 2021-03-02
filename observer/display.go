package observer

import "fmt"

// DisplayElement interface, added but not used
type DisplayElement interface {
	Display()
}

// CurrentConditionsDisplay
// ------------------------
type CurrentConditionsDisplay struct {
	subject  Subject
	temp     float64
	humidity float64
}

func NewCurrentConditionsDisplay(subj Subject) *CurrentConditionsDisplay {
	ccd := CurrentConditionsDisplay{subject: subj}
	ccd.subject.RegisterObserver(&ccd)
	return &ccd
}

func (ccd *CurrentConditionsDisplay) Update(temp float64, humidity float64, pressure float64) {
	ccd.temp = temp
	ccd.humidity = humidity
	ccd.Display()
}

func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current Conditions: %.1f degrees F with %.1f%% humidity\n", ccd.temp, ccd.humidity)
}

// StatisticsDisplay
// -----------------
type StatisticsDisplay struct {
	subject Subject
	// don't wanna keep all the stats - just calc using all values in Display
	temps []float64
}

func NewStatisticsDisplay(subj Subject) *StatisticsDisplay {
	sd := StatisticsDisplay{subject: subj}
	sd.temps = make([]float64, 0)
	sd.subject.RegisterObserver(&sd)
	return &sd
}

func (sd *StatisticsDisplay) Update(temp float64, humidity float64, pressure float64) {
	sd.temps = append(sd.temps, temp)
	sd.Display()
}

func (sd *StatisticsDisplay) Display() {
	min := 100000000.0
	max := -100000000.0

	tot := 0.0
	for _, tmp := range sd.temps {
		tot += tmp
		if tmp < min {
			min = tmp
		}
		if tmp > max {
			max = tmp
		}
	}
	avg := tot / float64(len(sd.temps))
	fmt.Printf("Stats: Avg/Max/Min: %.1f / %.1f / %.1f\n", avg, max, min)
}
