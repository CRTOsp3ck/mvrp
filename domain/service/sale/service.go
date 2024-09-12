package sale

import (
	"mvrp/data/repo"
	"os"
	"strconv"
)

// default settings to be used in the service
var SyncCreatedSalesOrderFromSalesQuotation bool = false
var SyncCreatedDeliveryNoteFromSalesOrder bool = false

type SaleService struct {
	Repo *repo.RepoContainer
}

func NewSaleService() *SaleService {
	return &SaleService{
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
