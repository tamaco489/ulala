package model

var MovieTypeList = map[uint32]string{
	1:  "action",
	2:  "anime",
	3:  "science_fiction",
	4:  "horror",
	5:  "comedy",
	6:  "romance",
	7:  "fantasy",
	8:  "sports",
	99: "other",
}

var MovieFormatList = map[uint32]string{
	1:  "mp4",
	2:  "avi",
	3:  "mov",
	4:  "wmv",
	5:  "flv",
	99: "other",
}

type MovieType struct {
	TypeID      uint32 `json:"type_id" gorm:"primaryKey"`
	TypeName    string `json:"type_name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MovieFormat struct {
	FormatID    uint32 `json:"format_id" gorm:"primaryKey"`
	MovieFormat string `json:"movie_format" gorm:"unique"`
}

type Movie struct {
	MovieID     uint32 `json:"movie_id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title"`
	ReleaseYear uint32 `json:"release_year"`
	Description string `json:"description"`
	TypeID      uint32 `json:"type_id"`
	FormatID    uint32 `json:"format_id"`
	ImageID     uint32 `json:"image_id" gorm:"unique"`
	ThumbnailID uint32 `json:"thumbnail_id" gorm:"unique"`
}

type GetMovieCategoriesResonse struct {
	TypeID      uint32 `json:"type_id" gorm:"primaryKey"`
	TypeName    string `json:"type_name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetMovieResponse struct {
	MovieID     uint32 `json:"movie_id"`
	Title       string `json:"title"`
	ReleaseYear uint32 `json:"release_year"`
	Description string `json:"description"`
	TypeName    string `json:"type_name"`
	MovieFormat string `json:"movie_format"`
}
