package profile

import (
	"encoding/xml"
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
