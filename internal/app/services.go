package app

import (
	auth "legocy-go/delievery/http/usecase/auth"
	"legocy-go/delievery/http/usecase/lego"
	"legocy-go/delievery/http/usecase/marketplace"
)

func (a *App) GetUserService() auth.UserUseCase {
	return auth.NewUserUsecase(a.GetUserRepo())
}

func (a *App) GetUserImagesService() auth.UserImageUseCase {
	return auth.NewUserImageUseCase(a.GetUserImagesRepo())
}

func (a *App) GetLegoSeriesService() lego.LegoSeriesService {
	return lego.NewLegoSeriesService(a.GetLegoSeriesRepo())
}

func (a *App) GetLegoSetService() lego.LegoSetUseCase {
	return lego.NewLegoSetUseCase(a.GetLegoSetRepo())
}

func (a *App) GetLocationService() marketplace.LocationUseCase {
	return marketplace.NewLocationUseCase(a.GetLocationRepo())
}

func (a *App) GetCurrencyService() marketplace.CurrencyUseCase {
	return marketplace.NewCurrencyUseCase(a.GetCurrencyRepo())
}

func (a *App) GetMarketItemService() marketplace.MarketItemService {
	return marketplace.NewMarketItemSerivce(a.GetMarketItemRepo())
}
