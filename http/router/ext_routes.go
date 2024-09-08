package router

import (
	"mvrp/http/handler/inventory"

	"github.com/go-chi/chi/v5"
)

func getExtRoutes() func(chi.Router) {
	return func(r chi.Router) {
		r.Route("/inventory", func(r chi.Router) {
			r.Route("/inventory", func(r chi.Router) {
				r.Get("/exists_by_item_id/{id}", inventory.GetInventoryExistsByItemID)
			})
		})
	}
}
