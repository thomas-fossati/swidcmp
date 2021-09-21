package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/veraison/swid"
)

var (
	saveCBORFile = true
)

// IN: a swid.xml file
// OUT: CoSWID compression factor
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s [SWID file]\n", os.Args[0])
	}

	xmlFile := os.Args[1]

	xml, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		log.Fatal(err)
	}

	var tag swid.SoftwareIdentity

	err = tag.FromXML(xml)
	if err != nil {
		log.Fatalf("reading SWID from %s: %v", xmlFile, err)
	}

	cbor, err := tag.ToCBOR()
	if err != nil {
		log.Fatal(err)
	}

	if saveCBORFile {
		cborFile := strings.TrimSuffix(xmlFile, filepath.Ext(xmlFile)) + ".cbor"

		err := ioutil.WriteFile(cborFile, cbor, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	delta := percentDelta(len(xml), len(cbor))

	log.Printf("Delta XML->CBOR: %f%% [%s]", delta, xmlFile)
}

func percentDelta(old, new int) float64 {
	diff := float64(new - old)
	return (diff / float64(old)) * 100
}
