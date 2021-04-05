package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/stacks/gongdoc/go/controllers"
	"github.com/fullstack-lang/gong/stacks/gongdoc/go/orm"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	apiFlag    = flag.Bool("api", false, "it true, use api controllers instead of default controllers")
)

func main() {

	log.SetPrefix("gongdoc: ")
	log.SetFlags(0)

	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("surplus arguments")
	}

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM
	db := orm.SetupModels(*logDBFlag, "./test.db")
	// mandatory, otherwise, bizarre errors occurs
	db.DB().SetMaxOpenConns(1)

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // a gin Context can have a map of variable that is set up at runtime
		c.Next()
	})

	controllers.RegisterControllers(r)

	r.Run()
}
