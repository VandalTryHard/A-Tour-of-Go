package main

import (
	"fmt"
	"log"

	"../calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(01)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetTitle("Hello Golang.")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event)

	fmt.Println(event.Date.Year())
	fmt.Println(event.Date.Month())
	fmt.Println(event.Date.Day())

	fmt.Println(event.Title())

}
