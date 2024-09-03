package inventory

import (
	"mvrp/data/repo"
)

type InventoryService struct {
	Repo *repo.RepoContainer
}

func NewInventoryService() *InventoryService {
	return &InventoryService{
		Repo: repo.NewRepoContainer(),
	}
}
