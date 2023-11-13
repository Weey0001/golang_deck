package main

import "fmt"

// create a new type of deck
// which is a slice of deck

type deck []string

func ( d deck ) print(){
  for i,card := range d {
		fmt.Println(i,card)
	}
}