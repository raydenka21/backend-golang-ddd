package postgres

import (
	"backend-app/domain/movie/model"
	"context"
)

func (r *repository) List(ctx context.Context) (res []model.Movie, err error) {
	result := r.db.Find(&res)
	return res, result.Error
}

func (r *repository) FindById(ctx context.Context,id int) (res model.Movie, err error) {
	result := r.db.Find(&res,id)
	return res, result.Error
}




