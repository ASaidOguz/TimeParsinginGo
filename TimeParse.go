package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	Realtime := "Thu, 05/19/11, 10:47PM"

	timeTemplate := "Mon, 01/02/06, 03:04PM"

	t, err := time.Parse(timeTemplate, Realtime)

	if err != nil {
		log.Fatal("Error in parsing time:", err)
	}
	timeunix := time.Now()
	fmt.Println("Right now :", timeunix)
	fmt.Println("Time before parsing: ", Realtime)
	fmt.Println("----------------------")
	fmt.Println("Time after parsing: ", t)

}
