package postgres

import (
	"backend-app/domain/movie/model"
	"context"
)

func (r *repository) Save(ctx context.Context, movie model.Movie) (res model.Movie, err error) {
	result := r.db.Save(&movie)
	return movie, result.Error
}

func (r *repository) Update(ctx context.Context, movie model.Movie) (res model.Movie, err error) {
	result := r.db.Save(&movie)
	return movie, result.Error
}

func (r *repository) DeleteById(ctx context.Context,id int) (err error) {
	result := r.db.Delete(&model.Movie{ID: id})
	return result.Error
}


