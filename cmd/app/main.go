package main

import (
	"fmt"
	"mvrp/data"
	"mvrp/domain"
	"mvrp/http"
	"strings"
)

func main() {
	initMessage()
	data.Init()
	domain.Init()
	http.Init()
}

func initMessage() {
	fmt.Printf("\n%s\n", strings.Repeat("-", 80))
	fmt.Printf("\n%s\n", "MVERP API Started")
	fmt.Printf("\n%s\n\n", strings.Repeat("-", 80))
}
