package cards

import "math"

type Determinator struct {
	originalValue int64
	currentValue  int64
	tickValue     int64
}

func NewDeterminator(seedValue int64) *Determinator {
	dt := Determinator{}
	dt.Seed(seedValue)
	return &dt
}

func (det *Determinator) Int63() int64 {
	return det.currentValue
}

func (det *Determinator) Seed(s int64) {
	det.originalValue = s
	det.currentValue = s
	det.tickValue = 1
}

func (det *Determinator) Tick() {
	if det.tickValue > int64(math.Sqrt(math.MaxInt64)) {
		det.tickValue = 1
	}
	if det.currentValue == math.MaxInt64 {
		det.tickValue++
		det.currentValue = det.tickValue
	}
}
