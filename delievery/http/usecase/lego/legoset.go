package lego

import (
	"context"
	models "legocy-go/pkg/lego/models"
	r "legocy-go/pkg/lego/repository"
)

type LegoSetUseCase struct {
	repo r.LegoSetRepository
}

func NewLegoSetUseCase(repo r.LegoSetRepository) LegoSetUseCase {
	return LegoSetUseCase{repo: repo}
}

func (u *LegoSetUseCase) ListLegoSets(c context.Context) ([]*models.LegoSet, error) {
	return u.repo.GetLegoSets(c)
}

func (u *LegoSetUseCase) LegoSetDetail(c context.Context, id int) (*models.LegoSet, error) {
	return u.repo.GetLegoSetByID(c, id)
}

func (u *LegoSetUseCase) LegoSetCreate(c context.Context, legoSet *models.LegoSetBasic) error {
	return u.repo.CreateLegoSet(c, legoSet)
}

func (u *LegoSetUseCase) LegoSetDelete(c context.Context, id int) error {
	return u.repo.DeleteLegoSet(c, id)
}
