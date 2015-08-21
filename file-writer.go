package main

import (
	"fmt"
)

func main() {

	cards := FetchCards()
	for _, entity := range cards.Entity {
		fmt.Println()
		for _, tag := range entity.Tag {
			if tag.Name == "CardTextInHand" {
				fmt.Printf("\t%s\n", tag.EnUS)
			}
		}
	}

}
