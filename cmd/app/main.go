package main

import (
	"fmt"
	"mvrp/data"
	"mvrp/domain"
	"mvrp/env"
	"mvrp/http"
	"strings"
)

func main() {
	initMessage()

	env.Init()

	data.Init()
	domain.Init()
	http.Init()
}

func initMessage() {
	fmt.Printf("\n%s\n", strings.Repeat("-", 80))
	fmt.Printf("\n%s\n", "MVERP API Started")
	fmt.Printf("\n%s\n\n", strings.Repeat("-", 80))
}
