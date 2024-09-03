package main

import (
	"fmt"
	"mvrp/cmd/gen/dto"
	"mvrp/cmd/gen/handler"
	"mvrp/cmd/gen/migrate"
	"mvrp/cmd/gen/model"
	"mvrp/cmd/gen/repo"
	"mvrp/cmd/gen/router"
	"strings"
	"time"
)

var migTables bool = false
var genModels bool = false
var genDTOs bool = false
var genRepos bool = false
var genHandlers bool = false
var genRouter bool = false

func main() {
	fmt.Printf("\n%s\n", strings.Repeat("-", 50))
	fmt.Printf("\nZERP Codegen Util: Code generation started\n")
	fmt.Printf("\n%s\n", strings.Repeat("-", 50))

	var err error

	if migTables {
		err = migrate.Run()
		if err != nil {
			logError(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	if genModels {
		err = model.Generate()
		if err != nil {
			logError(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	if genDTOs {
		err = dto.Generate()
		if err != nil {
			logError(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	if genRepos {
		err = repo.Generate()
		if err != nil {
			logError(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	if genHandlers {
		err = handler.Generate()
		if err != nil {
			logError(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	if genRouter {
		err = router.Generate()
		if err != nil {
			logError(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("\n%s\n", strings.Repeat("-", 50))
	fmt.Printf("\nZERP Codegen Util: Code generation complete\n")
	fmt.Printf("\n%s\n", strings.Repeat("-", 50))
}

func logError(err error) {
	fmt.Printf("\n%s\n", strings.Repeat("-", 50))
	fmt.Printf("\nZERP Codegen Util: Code generation failed\n")
	fmt.Printf("\n%s\n", err.Error())
	fmt.Printf("\n%s\n", strings.Repeat("-", 50))
}
