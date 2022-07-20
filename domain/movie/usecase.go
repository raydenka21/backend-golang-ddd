package movie

import (
	"backend-app/domain/movie/model"
	"backend-app/domain/movie/postgres"
	"context"
	"fmt"
)

type UseCase interface {
	List(context context.Context) ([]model.Movie, error)
	Save(context context.Context,movie model.Movie) (model.Movie, error)
	FindById(context context.Context, id int)  (model.Movie, error)
	Update(context context.Context, movie model.Movie)  (model.Movie, error)
	DeleteById(context context.Context, id int) error

}

type useCase struct {
	repo postgres.Repository
}

// NewUseCase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewUseCase(repo postgres.Repository) UseCase {
	fmt.Println("usecase")
	return &useCase{
		repo: repo,
	}
}

func (us *useCase) Save(context context.Context,newMovie model.Movie) (res model.Movie, err error) {
	res, err = us.repo.Save(context,newMovie)
	return newMovie, err
}

func (us *useCase) FindById(context context.Context,id int) (res model.Movie, err error) {
	res, err = us.repo.FindById(context,id)
	return res, err
}

func (us *useCase) List(context context.Context) (res []model.Movie, err error) {
	res, err = us.repo.List(context)
	return res, err
}

func (us *useCase) Update(context context.Context, editedMovie model.Movie) (res model.Movie, err error) {
	res, err = us.repo.Update(context, editedMovie)
	return res, err
}

func (us *useCase) DeleteById(context context.Context,id int) (err error) {
	err = us.repo.DeleteById(context,id)
	return err
}



