package entity

import (
	"mvrp/data/repo"
	"os"
	"strconv"
)

type EntityService struct {
	Repo *repo.RepoContainer
}

func NewEntityService() *EntityService {
	return &EntityService{
		Repo: repo.NewRepoContainer(),
	}
}

var loopInterval int

func Init() {
	err := loadServiceLoopInterval()
	if err != nil {
		panic(err)
	}
	_ = loopInterval
}

func loadServiceLoopInterval() error {
	v := os.Getenv("SERVICE_LOOP_INTERVAL_MS")
	conv, err := strconv.Atoi(v)
	if err != nil {
		return err
	}
	loopInterval = conv
	return nil
}
