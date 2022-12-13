package repository

import (
	"github.com/jinzhu/gorm"
	"go-sample/index"
)

// Movie is Gorm model of movie
type Movie struct {
	gorm.Model
	index.Video
}

// Subtitle is Gorm model of subtitle
type Subtitle struct {
	gorm.Model
	index.Info
}
