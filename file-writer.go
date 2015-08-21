package main

import (
	"io/ioutil"
)

func main() {

	cards := FetchCards()
	var allCardTexts string
	for _, entity := range cards.Entity {
		for _, tag := range entity.Tag {
			if tag.Name == "CardTextInHand" {
				allCardTexts += tag.EnUS
			}
		}
	}
	err := ioutil.WriteFile("cards.txt", []byte(allCardTexts), 0644)
	CheckError(err)

}
