package profile

import (
	"encoding/xml"
	"fmt"
)

func (p *Profile) UpdateChecksum() error {
	xmlColors, err := xml.Marshal(p.Colors)
	if err != nil {
		return fmt.Errorf("Cannot marshal xml colors: ", err)
	}

	var l int32
	for _, char := range xmlColors {
		l = ((l << 5) - l) + int32(char)
	}

	if l < 0 {
		l = -(l + 1)
	}

	p.Header.Checksum = fmt.Sprintf("%v", l)

	return nil
}
