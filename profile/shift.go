package profile

import (
	"time"

	"github.com/pkg/errors"
)

const timeFormat = "1504"

func (p *Profile) Shift(to string) error {
	now := time.Now()
	justTime, err := time.Parse(timeFormat, to)
	if err != nil {
		return errors.Errorf("Cannot parse shift time: %s", err)
	}
	t := time.Date(now.Year(), now.Month(), now.Day(), justTime.Hour(), justTime.Minute(), 0, 0, time.UTC)

	c := p.Colors

	firstPoint := c.Blue.Points[1].Time
	firstPoint = comparePoints(c.DeepRed, firstPoint)
	firstPoint = comparePoints(c.UltraViolet, firstPoint)
	firstPoint = comparePoints(c.Violet, firstPoint)
	firstPoint = comparePoints(c.CoolWhite, firstPoint)
	firstPoint = comparePoints(c.Green, firstPoint)
	firstPoint = comparePoints(c.Royal, firstPoint)

	shiftAmount := -firstPoint.Sub(t)

	c.DeepRed.Shift(shiftAmount)
	c.UltraViolet.Shift(shiftAmount)
	c.Violet.Shift(shiftAmount)
	c.CoolWhite.Shift(shiftAmount)
	c.Green.Shift(shiftAmount)
	c.Blue.Shift(shiftAmount)
	c.Royal.Shift(shiftAmount)

	return nil
}

func comparePoints(c color, currentPoint aiTime) aiTime {

	if len(c.Points) > 1 {
		if c.Points[1].Time.Unix() < currentPoint.Unix() {
			return c.Points[1].Time
		}
	}
	return currentPoint
}
