package main

import (
	"github.com/gin-gonic/gin"
	"go-sample/server"
	"log"
	"os"
	"path/filepath"
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
	ResetOnStart = false
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
	dbConn, err := server.OpenDB(getDBPath())
	if err != nil {
		cleanup()
		panic("Unable to create DB connection:" + err.Error())
	}
	//defer dbConn.Close()
	log.Println("Preparing database...")
	server.Migrate(dbConn)
	log.Println("Database prepared")
	// init db conn
	server.InitMovieRepository(dbConn)
	server.InitConfigRepository(dbConn)
	// api
	router := gin.Default()
	server.MovieRouterGroup(router)
	server.ConfigRouterGroup(router)

	// http play
	go server.HttpServer(":9001")
	//
	router.Run()
}

func cleanup() {
	log.Println("Cleaning up artifacts")
	os.RemoveAll(getAppDir())
}
