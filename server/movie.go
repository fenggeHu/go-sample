package server

import (
	"github.com/gin-gonic/gin"
	"go-sample/index"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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
	gorm.Model
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
				if o.DeletedAt.Valid {
					o.DeletedAt = gorm.DeletedAt{}
					news = append(news, o)
				}
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

func (rep *MovieRepository) queryById(id uint) (ret Movie) {
	rep.db.First(&ret, id)
	return
}

func (rep *MovieRepository) queryByCategory(category string) (ret []Movie) {
	rep.db.Model(&Movie{}).Where("category like ?", category+"%").Find(&ret)
	return
}

func (rep *MovieRepository) queryByRoot(root string) (ret []Movie) {
	rep.db.Unscoped().Model(&Movie{}).Where("root = ?", root).Find(&ret)
	return
}

func (rep *MovieRepository) batchInsert(movie []Movie) {
	rep.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Model(&Movie{}).CreateInBatches(movie, 100)
}

func (rep *MovieRepository) batchDelete(paths []string) {
	rep.db.Where("path in (?)", paths).Delete(&Movie{})
}

func (rep *MovieRepository) remove(root string) (ret []string) {
	rep.db.Delete(&Movie{}, "root = ?", root)
	return
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
	var configs []Config
	if len(root) < 3 {
		configs = configRepo.list()
	} else {
		configs = configRepo.query(root)
	}

	for _, v := range configs {
		movieRepo.reset(v.Path)
	}
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

func HttpServer(addr string) {
	srv := http.Server{
		Addr: addr,
	}

	http.HandleFunc("/movie/play", streamHandler)

	srv.ListenAndServe()
}
func streamHandler(w http.ResponseWriter, r *http.Request) {
	// 首次请求拿不到BasicAuth
	//username, password, ok := r.BasicAuth()
	//if !ok {
	//	sendErrorResponse(w, http.StatusUnauthorized, "Basic Authorization failed")
	//	return
	//}
	//if username != "max" || password != "hi12345" {
	//	sendErrorResponse(w, http.StatusUnauthorized, "User/Password Error")
	//	return
	//}

	id := r.URL.Query().Get("id")
	uid, _ := strconv.Atoi(id)
	movie := movieRepo.queryById(uint(uid))
	if len(movie.Path) == 0 {
		sendErrorResponse(w, http.StatusNotFound, "Basic Authorization failed")
		return
	}
	video, err := os.Open(movie.Path)
	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.Header().Set("Content-Type", "video/"+movie.Type[1:])
	http.ServeContent(w, r, movie.Name, time.Now(), video)

	//defer video.Close()
}

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}
