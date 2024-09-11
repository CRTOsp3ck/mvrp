// Code generated by MVRP Codegen Util. DO NOT EDIT.

package router

import (
	"log"
	"mvrp/util"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	httpSwagger "github.com/swaggo/http-swagger"
	"mvrp/http/handler/entity"
	"mvrp/http/handler/inventory"
	"mvrp/http/handler/invoice"
	"mvrp/http/handler/item"
	"mvrp/http/handler/sale"
	"mvrp/http/handler/enum"
)

func Init() {
	r := chi.NewRouter()
	c := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Cache-Control"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	r.Use(c)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	SetupRoutes(r)
	SetupSwagger(r)

	http.ListenAndServe(":6900", r)
}

func SetupRoutes(r chi.Router) {
	r.Get("/", getHome)
	r.Route("/v1", func(r chi.Router) {
		r.Route("/main", getMainRoutes())
		r.Route("/ext", getExtRoutes())
		r.Route("/enum", getEnumRoutes())
	})
}

func SetupSwagger(r chi.Router) {
	root, err := util.Util.FS.FindProjectRoot("go.mod")
	if err != nil {
		log.Fatalf("Error finding project root: %v\n", err)
	}
	filename := filepath.Join(root, "http", "router", "openapi.yaml")
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/openapi.yaml"),
	))
	r.Get("/swagger/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func getMainRoutes() func(chi.Router) {
    return func(r chi.Router) {
		r.Route("/entity", func(r chi.Router) {
			r.Route("/entity", func(r chi.Router) {
				r.Get("/", entity.ListEntity)
				r.Post("/", entity.CreateEntity)
				r.Post("/search", entity.SearchEntity)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(entity.EntityContext)
					r.Get("/", entity.GetEntity)
					r.Put("/", entity.UpdateEntity)
					r.Delete("/", entity.DeleteEntity)
				})
			})
			
		})
		
		r.Route("/inventory", func(r chi.Router) {
			r.Route("/goods_issue_note", func(r chi.Router) {
				r.Get("/", inventory.ListGoodsIssueNote)
				r.Post("/", inventory.CreateGoodsIssueNote)
				r.Post("/search", inventory.SearchGoodsIssueNote)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(inventory.GoodsIssueNoteContext)
					r.Get("/", inventory.GetGoodsIssueNote)
					r.Put("/", inventory.UpdateGoodsIssueNote)
					r.Delete("/", inventory.DeleteGoodsIssueNote)
				})
			})
			
			r.Route("/goods_issue_note_view", func(r chi.Router) {
				r.Get("/", inventory.ListGoodsIssueNoteView)
				r.Post("/search", inventory.SearchGoodsIssueNoteView)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(inventory.GoodsIssueNoteViewContext)
					r.Get("/", inventory.GetGoodsIssueNoteView)
				})
			})
			
			r.Route("/inventory", func(r chi.Router) {
				r.Get("/", inventory.ListInventory)
				r.Post("/", inventory.CreateInventory)
				r.Post("/search", inventory.SearchInventory)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(inventory.InventoryContext)
					r.Get("/", inventory.GetInventory)
					r.Put("/", inventory.UpdateInventory)
					r.Delete("/", inventory.DeleteInventory)
				})
			})
			
			r.Route("/inventory_view", func(r chi.Router) {
				r.Get("/", inventory.ListInventoryView)
				r.Post("/search", inventory.SearchInventoryView)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(inventory.InventoryViewContext)
					r.Get("/", inventory.GetInventoryView)
				})
			})
			
			r.Route("/inventory_transaction", func(r chi.Router) {
				r.Get("/", inventory.ListInventoryTransaction)
				r.Post("/", inventory.CreateInventoryTransaction)
				r.Post("/search", inventory.SearchInventoryTransaction)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(inventory.InventoryTransactionContext)
					r.Get("/", inventory.GetInventoryTransaction)
					r.Put("/", inventory.UpdateInventoryTransaction)
					r.Delete("/", inventory.DeleteInventoryTransaction)
				})
			})
			
			r.Route("/return_merchandise_authorization", func(r chi.Router) {
				r.Get("/", inventory.ListReturnMerchandiseAuthorization)
				r.Post("/", inventory.CreateReturnMerchandiseAuthorization)
				r.Post("/search", inventory.SearchReturnMerchandiseAuthorization)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(inventory.ReturnMerchandiseAuthorizationContext)
					r.Get("/", inventory.GetReturnMerchandiseAuthorization)
					r.Put("/", inventory.UpdateReturnMerchandiseAuthorization)
					r.Delete("/", inventory.DeleteReturnMerchandiseAuthorization)
				})
			})
			
			r.Route("/return_merchandise_authorization_view", func(r chi.Router) {
				r.Get("/", inventory.ListReturnMerchandiseAuthorizationView)
				r.Post("/search", inventory.SearchReturnMerchandiseAuthorizationView)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(inventory.ReturnMerchandiseAuthorizationViewContext)
					r.Get("/", inventory.GetReturnMerchandiseAuthorizationView)
				})
			})
			
			r.Route("/stock_count_sheet", func(r chi.Router) {
				r.Get("/", inventory.ListStockCountSheet)
				r.Post("/", inventory.CreateStockCountSheet)
				r.Post("/search", inventory.SearchStockCountSheet)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(inventory.StockCountSheetContext)
					r.Get("/", inventory.GetStockCountSheet)
					r.Put("/", inventory.UpdateStockCountSheet)
					r.Delete("/", inventory.DeleteStockCountSheet)
				})
			})
			
			r.Route("/stock_count_sheet_view", func(r chi.Router) {
				r.Get("/", inventory.ListStockCountSheetView)
				r.Post("/search", inventory.SearchStockCountSheetView)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(inventory.StockCountSheetViewContext)
					r.Get("/", inventory.GetStockCountSheetView)
				})
			})
			
		})
		
		r.Route("/invoice", func(r chi.Router) {
			r.Route("/credit_note", func(r chi.Router) {
				r.Get("/", invoice.ListCreditNote)
				r.Post("/", invoice.CreateCreditNote)
				r.Post("/search", invoice.SearchCreditNote)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(invoice.CreditNoteContext)
					r.Get("/", invoice.GetCreditNote)
					r.Put("/", invoice.UpdateCreditNote)
					r.Delete("/", invoice.DeleteCreditNote)
				})
			})
			
			r.Route("/debit_note", func(r chi.Router) {
				r.Get("/", invoice.ListDebitNote)
				r.Post("/", invoice.CreateDebitNote)
				r.Post("/search", invoice.SearchDebitNote)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(invoice.DebitNoteContext)
					r.Get("/", invoice.GetDebitNote)
					r.Put("/", invoice.UpdateDebitNote)
					r.Delete("/", invoice.DeleteDebitNote)
				})
			})
			
			r.Route("/invoice", func(r chi.Router) {
				r.Get("/", invoice.ListInvoice)
				r.Post("/", invoice.CreateInvoice)
				r.Post("/search", invoice.SearchInvoice)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(invoice.InvoiceContext)
					r.Get("/", invoice.GetInvoice)
					r.Put("/", invoice.UpdateInvoice)
					r.Delete("/", invoice.DeleteInvoice)
				})
			})
			
			r.Route("/payment_receipt", func(r chi.Router) {
				r.Get("/", invoice.ListPaymentReceipt)
				r.Post("/", invoice.CreatePaymentReceipt)
				r.Post("/search", invoice.SearchPaymentReceipt)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(invoice.PaymentReceiptContext)
					r.Get("/", invoice.GetPaymentReceipt)
					r.Put("/", invoice.UpdatePaymentReceipt)
					r.Delete("/", invoice.DeletePaymentReceipt)
				})
			})
			
		})
		
		r.Route("/item", func(r chi.Router) {
			r.Route("/item", func(r chi.Router) {
				r.Get("/", item.ListItem)
				r.Post("/", item.CreateItem)
				r.Post("/search", item.SearchItem)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(item.ItemContext)
					r.Get("/", item.GetItem)
					r.Put("/", item.UpdateItem)
					r.Delete("/", item.DeleteItem)
				})
			})
			
		})
		
		r.Route("/sale", func(r chi.Router) {
			r.Route("/delivery_note", func(r chi.Router) {
				r.Get("/", sale.ListDeliveryNote)
				r.Post("/", sale.CreateDeliveryNote)
				r.Post("/search", sale.SearchDeliveryNote)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(sale.DeliveryNoteContext)
					r.Get("/", sale.GetDeliveryNote)
					r.Put("/", sale.UpdateDeliveryNote)
					r.Delete("/", sale.DeleteDeliveryNote)
				})
			})
			
			r.Route("/delivery_note_view", func(r chi.Router) {
				r.Get("/", sale.ListDeliveryNoteView)
				r.Post("/search", sale.SearchDeliveryNoteView)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(sale.DeliveryNoteViewContext)
					r.Get("/", sale.GetDeliveryNoteView)
				})
			})
			
			r.Route("/goods_return_note", func(r chi.Router) {
				r.Get("/", sale.ListGoodsReturnNote)
				r.Post("/", sale.CreateGoodsReturnNote)
				r.Post("/search", sale.SearchGoodsReturnNote)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(sale.GoodsReturnNoteContext)
					r.Get("/", sale.GetGoodsReturnNote)
					r.Put("/", sale.UpdateGoodsReturnNote)
					r.Delete("/", sale.DeleteGoodsReturnNote)
				})
			})
			
			r.Route("/goods_return_note_view", func(r chi.Router) {
				r.Get("/", sale.ListGoodsReturnNoteView)
				r.Post("/search", sale.SearchGoodsReturnNoteView)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(sale.GoodsReturnNoteViewContext)
					r.Get("/", sale.GetGoodsReturnNoteView)
				})
			})
			
			r.Route("/order_confirmation", func(r chi.Router) {
				r.Get("/", sale.ListOrderConfirmation)
				r.Post("/", sale.CreateOrderConfirmation)
				r.Post("/search", sale.SearchOrderConfirmation)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(sale.OrderConfirmationContext)
					r.Get("/", sale.GetOrderConfirmation)
					r.Put("/", sale.UpdateOrderConfirmation)
					r.Delete("/", sale.DeleteOrderConfirmation)
				})
			})
			
			r.Route("/order_confirmation_view", func(r chi.Router) {
				r.Get("/", sale.ListOrderConfirmationView)
				r.Post("/search", sale.SearchOrderConfirmationView)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(sale.OrderConfirmationViewContext)
					r.Get("/", sale.GetOrderConfirmationView)
				})
			})
			
			r.Route("/sales_order", func(r chi.Router) {
				r.Get("/", sale.ListSalesOrder)
				r.Post("/", sale.CreateSalesOrder)
				r.Post("/search", sale.SearchSalesOrder)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(sale.SalesOrderContext)
					r.Get("/", sale.GetSalesOrder)
					r.Put("/", sale.UpdateSalesOrder)
					r.Delete("/", sale.DeleteSalesOrder)
				})
			})
			
			r.Route("/sales_order_view", func(r chi.Router) {
				r.Get("/", sale.ListSalesOrderView)
				r.Post("/search", sale.SearchSalesOrderView)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(sale.SalesOrderViewContext)
					r.Get("/", sale.GetSalesOrderView)
				})
			})
			
			r.Route("/sales_quotation", func(r chi.Router) {
				r.Get("/", sale.ListSalesQuotation)
				r.Post("/", sale.CreateSalesQuotation)
				r.Post("/search", sale.SearchSalesQuotation)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(sale.SalesQuotationContext)
					r.Get("/", sale.GetSalesQuotation)
					r.Put("/", sale.UpdateSalesQuotation)
					r.Delete("/", sale.DeleteSalesQuotation)
				})
			})
			
			r.Route("/sales_quotation_view", func(r chi.Router) {
				r.Get("/", sale.ListSalesQuotationView)
				r.Post("/search", sale.SearchSalesQuotationView)
				r.Route("/{id}", func(r chi.Router) {
					r.Use(sale.SalesQuotationViewContext)
					r.Get("/", sale.GetSalesQuotationView)
				})
			})
			
		})
		
    }
}

func getEnumRoutes() func(chi.Router) {
    return func(r chi.Router) {
        r.Get("/", enum.ListEnum)
    }
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to MVRP API"))
}