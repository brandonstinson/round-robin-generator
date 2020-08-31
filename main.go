package main

import (
	"flag"
	"fmt"
)

type pair []string
type round []pair
type roundRobin []round

func main() {
	short := flag.Bool("s", false, "provides a shorter response")
	interval := flag.String("i", "r", "interval - days [d], weeks, [w], months [m]")
	file := flag.Bool("f", false, "write to a file")
	flag.Parse()

	intervalStr := buildIntervalString(*interval, *short)

	people := flag.Args()

	if len(people)%2 != 0 {
		people = append(people, "bye")
	}

	roundRobin := createRoundRobin(people)

	var roundRobinStr string

	if *short == true {
		roundRobinStr = buildShortRoundRobinStr(roundRobin, intervalStr)
	} else {
		roundRobinStr = buildLongRoundRobinStr(roundRobin, intervalStr)
	}

	if *file == true {
		writeToFile("round-robin.txt", roundRobinStr)
	} else {
		fmt.Println(roundRobinStr)
	}
}
