package item

import (
	"mvrp/data/repo"
	"os"
	"strconv"
)

type ItemService struct {
	Repo *repo.RepoContainer
}

func NewItemService() *ItemService {
	return &ItemService{
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
