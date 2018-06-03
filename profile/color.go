package profile

import (
	"log"
	"time"
)

type color struct {
	Points []Point `xml:"point"`
}

type Point struct {
	Intensity uint64 `xml:"intensity"`
	Time      aiTime `xml:"time"`
}

func (c *color) Shift(shiftAmount time.Duration) {
	// Find the first non-zero point
	for i, _ := range c.Points {
		if i == 0 {
			// Skip the first point
			continue
		}
		p := &c.Points[i]
		newTime := p.Time.Add(shiftAmount)
		//log.Printf("New point is at %v, to is %v, value is %v; shift Duration is %v\n", newTime, to, p.Intensity, shiftAmount)
		if newTime.YearDay() != p.Time.YearDay() {
			if p.Intensity == 0 {
				// Just allign the zeros to the end of the day
				newTime = newTime.Truncate(24 * 60 * time.Minute).Add(-time.Second)
			} else {
				log.Fatalf("Shifting point at %v by %v went off the end of the day: %v\nCouldn't truncate it because intensity is %v", p.Time, shiftAmount, newTime, p.Intensity)
			}
		}
		p.Time = aiTime{newTime}
	}
}
