package data

import (
	"mvrp/data/db"
	"mvrp/data/model"
	"mvrp/data/repo"
)

func Init() {
	/*
		Initialize the data layer.
	*/

	// Initialize the database.
	db.Init()

	// Initialize the model.
	model.Init()

	// Initialize the repository.
	repo.Init()
}
