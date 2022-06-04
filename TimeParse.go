package main

import (
	"fmt"
	"time"
)

func main() {
	Realtime := "Thu, 05/19/11, 10:47"

	timeTemplate := "Mon, 01/02/06, 03:04"

	t, _ := time.Parse(timeTemplate, Realtime)

	fmt.Println("Time before parsing: ", Realtime)
	fmt.Println("----------------------")
	fmt.Println("Time after parsing: ", t)

}
