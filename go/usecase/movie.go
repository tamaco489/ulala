package usecase

import (
	"github.com/miyabiii1210/ulala/go/model"
	"github.com/miyabiii1210/ulala/go/repository"
)

type IMovieUsecase interface {
	GetMovieCategories() ([]model.GetMovieCategoriesResonse, error)
	GetMovieListByTypeID(typeID uint32) ([]model.GetMovieResponse, error)
	GetMovie(movieID uint32) (model.GetMovieResponse, error)
}

type movieUsecase struct {
	mr repository.IMovieRepository
}

func NewMovieUsecase(mr repository.IMovieRepository) IMovieUsecase {
	return &movieUsecase{mr}
}

func (mu *movieUsecase) GetMovieCategories() ([]model.GetMovieCategoriesResonse, error) {
	movieTypes := []model.GetMovieCategoriesResonse{}
	if err := mu.mr.GetMovieCategories(&movieTypes); err != nil {
		return []model.GetMovieCategoriesResonse{}, err
	}

	res := []model.GetMovieCategoriesResonse{}
	for _, movieType := range movieTypes {
		movieType := model.GetMovieCategoriesResonse{
			TypeID:      movieType.TypeID,
			TypeName:    movieType.TypeName,
			Title:       movieType.Title,
			Description: movieType.Description,
		}
		res = append(res, movieType)
	}

	return res, nil
}

func (mu *movieUsecase) GetMovieListByTypeID(typeID uint32) ([]model.GetMovieResponse, error) {
	movies := []model.GetMovieResponse{}
	if err := mu.mr.GetMovieListByTypeID(&movies, typeID); err != nil {
		return []model.GetMovieResponse{}, err
	}

	res := []model.GetMovieResponse{}
	for _, movie := range movies {
		movie := model.GetMovieResponse{
			MovieID:     movie.MovieID,
			Title:       movie.Title,
			ReleaseYear: movie.ReleaseYear,
			Description: movie.Description,
			TypeName:    movie.TypeName,
			MovieFormat: movie.MovieFormat,
		}
		res = append(res, movie)
	}

	return res, nil
}

func (mu *movieUsecase) GetMovie(movieID uint32) (model.GetMovieResponse, error) {
	movie := model.GetMovieResponse{}
	if err := mu.mr.GetMovie(&movie, movieID); err != nil {
		return model.GetMovieResponse{}, err
	}

	res := model.GetMovieResponse{
		MovieID:     movie.MovieID,
		Title:       movie.Title,
		ReleaseYear: movie.ReleaseYear,
		Description: movie.Description,
		TypeName:    movie.TypeName,
		MovieFormat: movie.MovieFormat,
	}

	return res, nil
}
