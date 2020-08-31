package main

import (
	"fmt"
	"strings"
)

func buildShortRoundRobinStr(roundRobin roundRobin, intervalStr string) string {
	var roundRobinStr strings.Builder
	for i, round := range roundRobin {
		var pairs string
		for _, pair := range round {
			pairs = fmt.Sprintf("%s - %s/%s", pairs, pair[0], pair[1])
		}
		fmt.Fprintf(&roundRobinStr, "%s%v%s\n", intervalStr, i+1, pairs)
	}
	return roundRobinStr.String()
}

func buildLongRoundRobinStr(roundRobin roundRobin, intervalStr string) string {
	var roundRobinStr strings.Builder
	for i, round := range roundRobin {
		fmt.Fprintf(&roundRobinStr, "%s%v\n", intervalStr, i+1)
		for j, pair := range round {
			fmt.Fprintf(&roundRobinStr, "Pair %v - %s & %s\n", j+1, pair[0], pair[1])
		}
	}
	return roundRobinStr.String()
}

func buildIntervalString(interval string, short bool) string {
	var intervalStr string
	switch interval {
	case "d":
		if short == true {
			intervalStr = "D"
		} else {
			intervalStr = "Day "
		}
	case "w":
		if short == true {
			intervalStr = "W"
		} else {
			intervalStr = "Week "
		}
	case "m":
		if short == true {
			intervalStr = "M"
		} else {
			intervalStr = "Month "
		}
	default:
		if short == true {
			intervalStr = "R"
		} else {
			intervalStr = "Round "
		}
	}
	return intervalStr
}
