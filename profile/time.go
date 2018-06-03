package profile

import (
	"encoding/xml"
	"fmt"
	"time"
)

type aiTime struct {
	time.Time
}

func (t *aiTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var v string
	d.DecodeElement(&v, &start)

	v = fmt.Sprintf("%sm", v)
	duration, err := time.ParseDuration(v)

	newTime := time.Now().UTC().Truncate(60 * 24 * time.Minute).Add(duration)

	*t = aiTime{newTime}

	return err
}

func (t *aiTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	e.EncodeToken(start)
	mins := fmt.Sprintf("%v", t.Hour()*60+t.Minute())
	minsBA := []byte(mins)
	e.EncodeToken(xml.CharData(minsBA))
	e.EncodeToken(start.End())

	return nil
}
