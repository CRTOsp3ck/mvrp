package enum

import (
	"mvrp/data/repo"
)

type EnumService struct {
	Repo *repo.RepoContainer
}

func NewEnumService() *EnumService {
	return &EnumService{
		Repo: repo.NewRepoContainer(),
	}
}
