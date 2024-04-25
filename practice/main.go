package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// fmt.Println("Day", dow)

	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "Sunday"
		// fallthrough
	case 2:
		result = "Monday"
		// fallthrough
	default:
		result = "Other"
	}
	fmt.Println(result)
}
