package profile

import "encoding/xml"

type Profile struct {
	XMLName xml.Name `xml:"ramp"`
	Header  Header   `xml:"header"`
	Colors  Colors   `xml:"colors"`
}

type Header struct {
	Version  string `xml:"version"`
	Checksum string `xml:"checksum"`
}

type Colors struct {
	XMLName     xml.Name `xml:"colors"`
	DeepRed     Color    `xml:"deep_red"`
	UltraViolet Color    `xml:"uv"`
	Violet      Color    `xml:"violet"`
	CoolWhite   Color    `xml:"cool_white"`
	Green       Color    `xml:"green"`
	Blue        Color    `xml:"blue"`
	Royal       Color    `xml:"royal"`
}

type Color struct {
	Points []Point `xml:"point"`
}

type Point struct {
	Intensity uint64 `xml:"intensity"`
	Time      uint64 `xml:"time"`
}
