package main

import (
	"go-sample/gorm-gen/query"
	"go-sample/server"
	"gorm.io/gen"
	"log"
	"os"
	"path/filepath"
)

func main() {
	movie, err := query.Movie.Where(query.Movie.Root.Eq("/Users/max/test")).First()
	log.Println(movie, err)
}
func main2() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// db
	var cacheDir, _ = os.UserCacheDir()
	dbpath := filepath.Join(filepath.Join(cacheDir, "go-media"), "media.db")
	gormdb, err := server.OpenDB(dbpath)
	if err != nil {
		log.Fatalln(err)
	}

	g.UseDB(gormdb) // reuse your gorm db
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(server.Movie{})
	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, server.Movie{})

	// Generate the code
	g.Execute()
}
