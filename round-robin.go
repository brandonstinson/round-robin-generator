package main

func createRound(people []string) round {
	var round round
	for i := 0; i < len(people); i++ {
		if i%2 == 0 {
			round = append(round, pair{people[i], people[len(people)-i-1]})
		}
	}
	return round
}

func createRoundRobin(people []string) roundRobin {
	var roundRobin roundRobin
	for i := 0; i < len(people)-1; i++ {
		val := people[1]
		people = append(people[:1], people[2:]...)
		people = append(people, val)
		roundRobin = append(roundRobin, createRound(people))
	}
	return roundRobin
}
