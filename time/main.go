package main

import (
	"fmt"
	"log"
	"time"
)

func useAfter() {
	fmt.Println("You have got two seconds to calculate 19 * 4")
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Time's up! The answer is 74")
			return
		}
	}
}

func useTick() {
	c := time.Tick(5 * time.Second)
	i := 0
	for t := range c {
		fmt.Printf("The time is now: %v\n", t)
		i++
		if i > 2 {
			return
		}

	}
}

func parseTime() {
	s := "2018-12-26T13:55:17+02:00"
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t)
	t2 := t.Add(20 * time.Second)
	fmt.Println("Add 20 sec:", t2)
	fmt.Println("is it after the initial time value:", t2.After(t))

	nextYearChristmas := time.Date(time.Now().Year()+1, 12, 24, 0, 0, 0, 0, time.UTC)
	fmt.Println("Number of days before next Christmas:", nextYearChristmas.Sub(time.Now()).Hours()/24)
}

func main() {
	fmt.Println("Now", time.Now())
	useAfter()
	useTick()
	parseTime()
}
