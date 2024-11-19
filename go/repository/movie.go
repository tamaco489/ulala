package repository

import (
	"github.com/miyabiii1210/ulala/go/model"
	"gorm.io/gorm"
)

type IMovieRepository interface {
	GetMovieCategories(movieTypes *[]model.GetMovieCategoriesResonse) error
	GetMovieListByTypeID(movies *[]model.GetMovieResponse, typeID uint32) error
	GetMovie(movie *model.GetMovieResponse, movieID uint32) error
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) IMovieRepository {
	return IMovieRepository(&movieRepository{db})
}

func (mr *movieRepository) GetMovieCategories(movieTypes *[]model.GetMovieCategoriesResonse) error {
	return mr.db.Table("movie_types").
		Find(&movieTypes).Error
}

func (mr *movieRepository) GetMovieListByTypeID(movies *[]model.GetMovieResponse, typeID uint32) error {
	return mr.db.Table("movies").
		Select("movies.movie_id, movies.title, movies.release_year, movies.description, movie_types.type_name, movie_formats.movie_format").
		Joins("JOIN movie_types ON movies.type_id = movie_types.type_id").
		Joins("JOIN movie_formats ON movies.format_id = movie_formats.format_id").
		Where("movies.type_id = ?", typeID).
		Find(&movies).Error
}

func (mr *movieRepository) GetMovie(movie *model.GetMovieResponse, movieID uint32) error {
	return mr.db.Table("movies").
		Select("movies.movie_id, movies.title, movies.release_year, movies.description, movie_types.type_name, movie_formats.movie_format").
		Joins("JOIN movie_types ON movies.type_id = movie_types.type_id").
		Joins("JOIN movie_formats ON movies.format_id = movie_formats.format_id").
		Where("movies.movie_id = ?", movieID).
		First(&movie).Error
}
