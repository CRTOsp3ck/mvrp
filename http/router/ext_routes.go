package router

import (
	"mvrp/http/handler/inventory"

	"github.com/go-chi/chi/v5"
)

func getExtRoutes() func(chi.Router) {
	return func(r chi.Router) {
		r.Route("/inventory", func(r chi.Router) {

			// Inventory routes
			r.Route("/inventory", func(r chi.Router) {
				r.Get("/exists_by_item_id/{id}", inventory.GetInventoryExistsByItemID)
			})

			// Inventory Transaction routes
			r.Route("/inventory_transaction", func(r chi.Router) {
				r.Post("/search_all", inventory.SearchAllInventoryTransaction)
			})
		})
	}
}
