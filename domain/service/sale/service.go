package sale

import (
	"mvrp/data/repo"
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
