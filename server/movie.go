package server

import (
	"github.com/gin-gonic/gin"
	"go-sample/index"
	"gorm.io/gorm"
	"net/http"
)

type MovieRepository struct {
	db *gorm.DB
}

var movieRepo = &MovieRepository{}

func InitMovieRepository(db *gorm.DB) {
	movieRepo.db = db
}

// Movie is Gorm model of movie
type Movie struct {
	//gorm.Model
	index.Video
}

// 重新扫描
func (rep *MovieRepository) reset(root string) {
	// TODO 判断root in config
	olds := rep.queryByRoot(root)
	locals := searchMovies(root)
	var news []Movie
	for _, n := range locals {
		isNew := true
		for _, o := range olds {
			if o.Path == n.Path {
				isNew = false
			}
		}
		if isNew {
			news = append(news, n)
		}
	}
	var losts []string
	for _, o := range olds {
		isLost := true
		for _, n := range locals {
			if n.Path == o.Path {
				isLost = false
			}
		}
		if isLost {
			losts = append(losts, o.Path)
		}
	}
	if len(losts) > 0 {
		movieRepo.batchDelete(losts)
	}
	if len(news) > 0 {
		movieRepo.batchInsert(news)
	}
}

func searchMovies(root string) (ret []Movie) {
	files := index.Scan(root)
	for _, f := range files {
		video := index.VideoInfo(f, root)
		if video == nil {
			continue
		}
		ret = append(ret, Movie{Video: *video})
	}
	return
}

func (rep *MovieRepository) queryByCategory(category string) (ret []Movie) {
	rep.db.Model(&Movie{}).Where("category like ?", category+"%").Find(&ret)
	return
}

func (rep *MovieRepository) queryByRoot(root string) (ret []Movie) {
	rep.db.Model(&Movie{}).Where("root = ?", root).Find(&ret)
	return
}

func (rep *MovieRepository) batchInsert(movie []Movie) {
	rep.db.Model(&Movie{}).CreateInBatches(movie, 100)
}

func (rep *MovieRepository) batchDelete(paths []string) {
	rep.db.Where("path in (?)", paths).Delete(&Movie{})
}

func (rep *MovieRepository) listRoot() (ret []string) {
	rep.db.Model(&Movie{}).Distinct().Pluck("root", &ret)
	return
}
func (rep *MovieRepository) list() (ret []Movie) {
	rep.db.Model(&Movie{}).Find(&ret)
	return
}

func MovieRouterGroup(router *gin.Engine) {
	group := router.Group("/movie") //.Use(middleware.HeaderHandler()).Use(middleware.PermissionHandler())
	// my demo
	group.GET("/root/list", listRoot)
	group.GET("/root", qMovieByRoot)
	group.GET("/cate", qMovieByCategory)
	group.GET("/index", indexMovies)
	group.GET("/list", list)
}

func list(c *gin.Context) {
	c.JSON(http.StatusOK, movieRepo.list())
}
func listRoot(c *gin.Context) {
	c.JSON(http.StatusOK, movieRepo.listRoot())
}

func indexMovies(c *gin.Context) {
	root := c.Query("root")
	//root := "/Users/max/test"
	movieRepo.reset(root)
	c.String(http.StatusOK, "success")
}
func qMovieByCategory(c *gin.Context) {
	cate := c.Query("cate")
	movies := movieRepo.queryByCategory(cate)
	c.JSON(http.StatusOK, movies)
}
func qMovieByRoot(c *gin.Context) {
	root := c.Query("root")
	movies := movieRepo.queryByRoot(root)
	c.JSON(http.StatusOK, movies)
}
