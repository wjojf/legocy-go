package postgres

import (
	models "legocy-go/pkg/lego/models"
)

type LegoSeriesPostgres struct {
	Model
	Name string `gorm:"unique"`
}

func FromLegoSeriesBasic(s *models.LegoSeriesBasic) *LegoSeriesPostgres {
	return &LegoSeriesPostgres{
		Name: s.Name,
	}
}

func FromLegoSeries(s *models.LegoSeries) *LegoSeriesPostgres {
	return &LegoSeriesPostgres{
		Name: s.Name,
	}
}

func (s *LegoSeriesPostgres) ToLegoSeries() *models.LegoSeries {
	return &models.LegoSeries{
		ID:   int(s.ID),
		Name: s.Name,
	}
}
