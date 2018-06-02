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
)

func main() {
	kingpin.Parse()

	xmlString, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatalf("Cannot read open %s: %s", filename, err)
	}

	doc := profile.Profile{}
	err = xml.Unmarshal(xmlString, &doc)
	if err != nil {
		log.Fatal("Cannot convert xml: ", err)
	}

	err = doc.UpdateChecksum()
	if err != nil {
		log.Fatal("Error generating checksum: ", err)
	}

	xmlOut, err := xml.MarshalIndent(doc, "", "\t")
	if err != nil {
		log.Fatal("Cannot convert json back into xml: ", err)
	}

	fmt.Printf("%s\n%s\n", xmlHeader, xmlOut)

}
