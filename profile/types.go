package profile

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Profile struct {
	XMLName xml.Name `xml:"ramp"`
	Header  header   `xml:"header"`
	Colors  colors   `xml:"colors"`
}

type header struct {
	Version  string `xml:"version"`
	Checksum string `xml:"checksum"`
}

type colors struct {
	XMLName     xml.Name `xml:"colors"`
	DeepRed     color    `xml:"deep_red"`
	UltraViolet color    `xml:"uv"`
	Violet      color    `xml:"violet"`
	CoolWhite   color    `xml:"cool_white"`
	Green       color    `xml:"green"`
	Blue        color    `xml:"blue"`
	Royal       color    `xml:"royal"`
}

type color struct {
	Points []Point `xml:"point"`
}

type Point struct {
	Intensity uint64 `xml:"intensity"`
	Time      aiTime `xml:"time"`
}

type aiTime struct {
	time.Duration
}

func (t *aiTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var v string
	d.DecodeElement(&v, &start)

	v = fmt.Sprintf("%sm", v)
	pT, err := time.ParseDuration(v)
	*t = aiTime{pT}

	return err
}

func (t *aiTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	e.EncodeToken(start)
	mins := fmt.Sprintf("%v", t.Minutes())
	minsBA := []byte(mins)
	e.EncodeToken(xml.CharData(minsBA))
	e.EncodeToken(start.End())

	return nil
}
