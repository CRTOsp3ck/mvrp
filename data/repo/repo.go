package repo

import (
	"context"
	"database/sql"
	"mvrp/data/repo/base"
	"mvrp/data/repo/entity"
	"mvrp/data/repo/enum"
	"mvrp/data/repo/inventory"
	"mvrp/data/repo/invoice"
	"mvrp/data/repo/item"
	"mvrp/data/repo/purchase"
	"mvrp/data/repo/sale"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type RepoContainer struct {
	Enum      *enum.EnumRepository
	Base      *base.BaseRepository
	Entity    *entity.EntityRepository
	Inventory *inventory.InventoryRepository
	Invoice   *invoice.InvoiceRepository
	Item      *item.ItemRepository
	Purchase  *purchase.PurchaseRepository
	Sale      *sale.SaleRepository
}

func Init() {}

func NewRepoContainer() *RepoContainer {
	enumRepo := enum.NewEnumRepository()
	baseRepo := base.NewBaseRepository()
	entityRepo := entity.NewEntityRepository()
	inventoryRepo := inventory.NewInventoryRepository()
	invoiceRepo := invoice.NewInvoiceRepository()
	itemRepo := item.NewItemRepository()
	purchaseRepo := purchase.NewPurchaseRepository()
	saleRepo := sale.NewSaleRepository()

	repo := &RepoContainer{
		Enum:      &enumRepo,
		Base:      &baseRepo,
		Entity:    &entityRepo,
		Inventory: &inventoryRepo,
		Invoice:   &invoiceRepo,
		Item:      &itemRepo,
		Purchase:  &purchaseRepo,
		Sale:      &saleRepo,
	}

	return repo
}

func (rc *RepoContainer) Begin(ctx context.Context) (*sql.Tx, error) {
	return boil.BeginTx(ctx, nil)
}

type RepoTx struct {
	Tx *sql.Tx
}

func (rc *RepoContainer) BeginRepoTx(ctx context.Context) (*RepoTx, error) {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &RepoTx{Tx: tx}, nil
}
