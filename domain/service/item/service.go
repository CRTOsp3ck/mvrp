package item

import (
	"mvrp/data/repo"
)

type ItemService struct {
	Repo *repo.RepoContainer
}

func NewItemService() *ItemService {
	return &ItemService{
		Repo: repo.NewRepoContainer(),
	}
}
