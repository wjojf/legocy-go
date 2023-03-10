package marketplace

import (
	"legocy-go/delievery/http/resources/auth"
	"legocy-go/delievery/http/resources/lego"
	models "legocy-go/pkg/marketplace/models"
)

type MarketItemRequest struct {
	LegoSetID  int     `json:"lego_set_id"`
	Price      float32 `json:"price"`
	CurrencyID int     `json:"currency_id"`
	LocationID int     `json:"location_id"`
}

func (r *MarketItemRequest) ToMarketItemBasic(sellerID int) *models.MarketItemBasic {
	return &models.MarketItemBasic{
		LegoSetID:  r.LegoSetID,
		SellerID:   sellerID,
		Price:      r.Price,
		CurrencyID: r.CurrencyID,
		LocationID: r.LocationID,
	}
}

type MarketItemResponse struct {
	ID       int                     `json:"id"`
	Price    float32                 `json:"price"`
	Currency CurrencyResponse        `json:"currency"`
	Location LocationResponse        `json:"location"`
	LegoSet  lego.LegoSetResponse    `json:"lego_set"`
	Seller   auth.UserDetailResponse `json:"seller"`
}

func GetMarketItemResponse(m *models.MarketItem) MarketItemResponse {
	return MarketItemResponse{
		ID:       m.ID,
		Price:    m.Price,
		Currency: GetCurrencyResponse(&m.Currency),
		Location: GetLocationResponse(&m.Location),
		LegoSet:  lego.GetLegoSetResponse(&m.LegoSet),
		Seller:   auth.GetUserDetailResponse(&m.Seller),
	}
}
