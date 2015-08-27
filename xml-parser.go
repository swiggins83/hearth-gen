package main

import (
	"encoding/xml"
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
	Value  string `xml:"value,attr"`
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

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func FetchCards() CardDefs {

	xmlFile, err := os.Open("cards.xml")
	CheckError(err)
	defer xmlFile.Close()

	data, _ := ioutil.ReadAll(xmlFile)

	var c CardDefs
	err = xml.Unmarshal([]byte(data), &c)
	CheckError(err)
	return c

}
