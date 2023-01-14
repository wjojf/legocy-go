package v1

import (
	models "legocy-go/pkg/lego/models"
	r "legocy-go/pkg/lego/repository"

	"golang.org/x/net/context"
)

type LegoSeriesService struct {
	repo r.LegoSeriesRepository
}

func (s *LegoSeriesService) ListSeries(ctx context.Context) ([]*models.LegoSeries, error) {
	return s.repo.GetLegoSeriesList(ctx)
}

func (s *LegoSeriesService) DetailSeries(ctx context.Context, id int) (*models.LegoSeries, error) {
	return s.repo.GetLegoSeries(ctx, id)
}
