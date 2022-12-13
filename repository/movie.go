package repository

import (
	"github.com/jinzhu/gorm"
	"go-sample/index"
)

// Movie is Gorm model of movie
type Movie struct {
	//gorm.Model
	index.Video
	Dir string `json:"dir"` //config path
}

func QMovies(db *gorm.DB, category string) (ret []Movie) {
	db.Model(&Movie{}).Where("category like ?", category+"%").Find(&ret)
	return
}
