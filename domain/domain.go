package domain

import (
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/domain/service"
)

func Init() {
	/*
		Initialize the domain layer.
	*/

	// Initialize the DTOs.
	dto.Init()

	// Initialize the pre and post processes.
	proc.Init()

	// Initialize the services.
	service.Init()
}
