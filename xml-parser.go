package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type CardDefs struct {
	XMLName xml.Name `xml:"CardDefs"`
	Entity  []Entity `xml:"Entity"`
}

type Entity struct {
	CardID      string `xml:"CardID,attr"`
	Version     string `xml:"version,attr"`
	MasterPower string `xml:"MasterPower"`
	Tag         []Tag  `xml:"Tag"`
	Power       Power  `xml:"Power"`
}

type Tag struct {
	EnumID string `xml:"enumID,attr"`
	Name   string `xml:"name,attr"`
	Type   string `xml:"type,attr"`
	EnUS   string `xml:"enUS"`
}

type Power struct {
	Definition      string          `xml:"definition,attr"`
	PlayRequirement PlayRequirement `xml:"PlayRequirement"`
}

type PlayRequirement struct {
	Param string `xml:"param,attr"`
	ReqID string `xml:"reqID,attr"`
}

func main() {

	xmlFile, err := os.Open("cards.xml")
	if err != nil {
		return
	}
	defer xmlFile.Close()

	data, _ := ioutil.ReadAll(xmlFile)

	var c CardDefs
	err = xml.Unmarshal([]byte(data), &c)
	if err != nil {
		return
	}
	fmt.Println(c.Entity[0].Tag[0].EnUS)
	//for _, entity := range episodiv.Entity {
	//	fmt.Printf("\t%s\n", entity)
	//}

}
