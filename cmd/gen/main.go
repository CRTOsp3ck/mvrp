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
)

var migTables bool = true
var genModels bool = true
var genDTOs bool = true
var genRepos bool = true
var genHandlers bool = true
var genRouter bool = true

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
	}
	if genModels {
		err = model.Generate()
		if err != nil {
			logError(err)
			return
		}
	}
	if genDTOs {
		err = dto.Generate()
		if err != nil {
			logError(err)
			return
		}
	}
	if genRepos {
		err = repo.Generate()
		if err != nil {
			logError(err)
			return
		}
	}
	if genHandlers {
		err = handler.Generate()
		if err != nil {
			logError(err)
			return
		}
	}
	if genRouter {
		err = router.Generate()
		if err != nil {
			logError(err)
			return
		}
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
