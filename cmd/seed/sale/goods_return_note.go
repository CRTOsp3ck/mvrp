package sale

import (
	"context"
	"mvrp/domain/dto"
	saleService "mvrp/domain/service/sale"
)

func SeedGoodsReturnNote(count int) error {
	// define services
	saleSvc := saleService.NewSaleService()

	// get all sales order with delivery notes
	searchDnDto := dto.SearchDeliveryNoteDTO{
		ShippingStatus: "shipped",
	}
	srDnReq := saleSvc.NewSearchDeliveryNoteViewRequest(context.Background(), searchDnDto)
	srDnResp, err := saleSvc.SearchDeliveryNoteView(srDnReq)
	if err != nil {
		return err
	}
	_ = srDnResp
	_ = searchDnDto

	return nil
}
