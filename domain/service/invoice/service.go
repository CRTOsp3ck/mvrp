package invoice

import (
	"mvrp/data/repo"
)

type InvoiceService struct {
	Repo *repo.RepoContainer
}

func NewInvoiceService() *InvoiceService {
	return &InvoiceService{
		Repo: repo.NewRepoContainer(),
	}
}
