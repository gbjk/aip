package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/gbjk/aip/profile"

	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	xmlHeader = `<?xml version='1.0' encoding='UTF-8' standalone='yes' ?>` + "\n"
)

var (
	filename = kingpin.Flag("file", "File to parse").Short('f').Required().String()
	shift    = kingpin.Command("shift", "Shift time")
	shiftTo  = shift.Arg("to", "24 hour Time to start at").Required().String()
)

func main() {
	command := kingpin.Parse()

	xmlString, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatalf("Cannot read open %s: %s", *filename, err)
	}

	p := profile.Profile{}
	err = xml.Unmarshal(xmlString, &p)
	if err != nil {
		log.Fatal("Cannot convert xml: ", err)
	}

	switch command {
	case "shift":
		err := p.Shift(*shiftTo)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = p.UpdateChecksum()
	if err != nil {
		log.Fatal("Error generating checksum: ", err)
	}

	xmlOut, err := xml.MarshalIndent(p, "", "\t")
	if err != nil {
		log.Fatal("Cannot convert json back into xml: ", err)
	}

	fmt.Printf("%s\n%s\n", xmlHeader, xmlOut)

}
