package entity

import (
	"mvrp/data/repo"
)

type EntityService struct {
	Repo *repo.RepoContainer
}

func NewEntityService() *EntityService {
	return &EntityService{
		Repo: repo.NewRepoContainer(),
	}
}
