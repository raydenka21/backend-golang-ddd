package postgres

import (
	"backend-app/domain/movie/model"
	"context"
	"fmt"
	"gorm.io/gorm"
)

// Repository Interface
type Repository interface {
	List(ctx context.Context) ([]model.Movie, error)
	Save(context context.Context,movie model.Movie) (model.Movie, error)
	FindById(context context.Context,id int) (model.Movie, error)
	Update(context context.Context, movie model.Movie)  (model.Movie, error)
	DeleteById(context context.Context,id int) error

}

// NewRepository will implement MovieRepository Interface
func NewRepository(db *gorm.DB) Repository {
	fmt.Println("repo")
	return &repository{db}
}

type repository struct {
	db *gorm.DB
}









