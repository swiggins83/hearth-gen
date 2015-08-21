package main

import (
	"io/ioutil"
	"strings"
)

func main() {

	cards := FetchCards()
	var allCardTexts string
	for _, entity := range cards.Entity {
		for _, tag := range entity.Tag {
			if tag.Name == "CardTextInHand" {
				s := strings.Replace(tag.EnUS, "<b>", "", 100)
				s = strings.Replace(s, "</b>", "", 100)
				s = strings.Replace(s, "<i>", "", 100)
				s = strings.Replace(s, "</i>", "", 100)
				if !strings.Contains(s, "\n") {
					s += "\n"
				}
				allCardTexts += s
			}
		}
	}
	err := ioutil.WriteFile("cards.txt", []byte(allCardTexts), 0644)
	CheckError(err)

}
