package main

import (
	"github.com/gin-gonic/gin"
	"go-sample/index"
	"go-sample/repository"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	vers    *bool
	help    *bool
	conf    *string
	testing *string
)

const (
	appDirName   = "go-media"
	dbFileName   = "media.db"
	ResetOnStart = true
)

var cacheDir, _ = os.UserCacheDir()

func getAppDir() string {
	return filepath.Join(cacheDir, appDirName)
}

func getDBPath() string {
	return filepath.Join(getAppDir(), dbFileName)
}

// function init run before main
func init() {
	// create working app dir
	appDirPath := getAppDir()
	if _, err := os.Stat(appDirPath); os.IsNotExist(err) {
		err = os.MkdirAll(appDirPath, 0777)
		if err != nil {
			panic("Unable to create App Dir on " + appDirPath)
		}
		log.Println("Created App dir at", appDirPath)
	}

	// create database if not exist
	dbPath := getDBPath()
	_, err := os.Stat(dbPath)
	if ResetOnStart {
		log.Println("Obsolete DB detected, removing...")
		if err = os.RemoveAll(dbPath); err != nil {
			panic("Unable removing obsolete DB")
		}
	}
	if ResetOnStart || os.IsNotExist(err) {
		_, err = os.Create(dbPath)
		if err != nil {
			log.Println("Unable to init db file", err)
			os.Exit(1)
		}
		log.Println("DB initialized at", dbPath)
	}

	//vers = flag.Bool("v", false, "display the version.")
	//help = flag.Bool("h", false, "print this help.")
	//conf = flag.String("f", "", "specify configuration file.")
	//testing = flag.String("t", "", "test configuration.")
	//flag.Parse()
	//
	//fmt.Println(*vers, *help, *conf, *testing)
}

func main() {
	dbConn, err := repository.OpenDB(getDBPath())
	if err != nil {
		cleanup()
		panic("Unable to create DB connection:" + err.Error())
	}

	defer dbConn.Close()

	log.Println("Preparing database...")
	repository.Migrate(dbConn)
	log.Println("Database prepared")

	//test.Pi()
	baseDir := "/Users/max/test"
	start := time.Now().UnixNano()
	files := index.Scan(baseDir)
	for _, f := range files {
		video := index.VideoInfo(f, baseDir)
		if video == nil {
			continue
		}
		dbConn.Create(&repository.Movie{Video: *video, Dir: baseDir})
		log.Println(video)
	}
	end := time.Now().UnixNano()

	log.Printf("%d, %d", (end-start)/1000000, len(files))

	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/movie/list", func(c *gin.Context) {
		cate := c.Query("cate")
		movies := repository.QMovies(dbConn, cate)
		c.JSON(http.StatusOK, movies)
	})

	router.Run()
}

func cleanup() {
	log.Println("Cleaning up artifacts")
	os.RemoveAll(getAppDir())
}
