package repository

import (
	"context"
	models "legocy-go/pkg/lego/models"
)

type LegoSeriesRepository interface {
	CreateLegoSeries(c context.Context, s *models.LegoSeries) error
	GetLegoSeriesList(c context.Context) ([]*models.LegoSeries, error)
	GetLegoSeries(c context.Context, id int) (*models.LegoSeries, error)
	DeleteLegoSeries(c context.Context, id int) error
}
