package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Config struct {
	Path string `gorm:"unique"`
}

type ConfigRepository struct {
	db *gorm.DB
}

var configRepo = &ConfigRepository{}

func InitConfigRepository(db *gorm.DB) {
	configRepo.db = db
}

func (rep ConfigRepository) add(path string) {
	if len(path) < 3 {
		return
	}
	rep.db.Create(&Config{Path: path})
}

func (rep ConfigRepository) remove(path string) {
	rep.db.Where("path = ?", path).Delete(&Config{})
}

func (rep ConfigRepository) query(path string) (ret []Config) {
	rep.db.Model(&Config{}).Where("path = ?", path).Find(&ret)
	return
}

func (rep ConfigRepository) list() (ret []Config) {
	rep.db.Model(&Config{}).Find(&ret)
	return
}

func ConfigRouterGroup(router *gin.Engine) {
	group := router.Group("/config") //.Use(middleware.HeaderHandler()).Use(middleware.PermissionHandler())
	// my demo
	group.GET("/add", addConfig)
	group.GET("/remove", removeConfig)
	group.GET("/list", listConfig)
}

func addConfig(c *gin.Context) {
	path := c.Query("path")
	configRepo.add(path)
	c.String(http.StatusOK, "success")
}

func removeConfig(c *gin.Context) {
	path := c.Query("path")
	configRepo.remove(path)
	movieRepo.remove(path)
	c.String(http.StatusOK, "success")
}

func listConfig(c *gin.Context) {
	c.JSON(http.StatusOK, configRepo.list())
}
