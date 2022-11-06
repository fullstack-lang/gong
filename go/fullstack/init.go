package fullstack

import (
	// gong stack for model analysis

	gong_controllers "github.com/fullstack-lang/gong/go/controllers"
	gong_orm "github.com/fullstack-lang/gong/go/orm"
	"github.com/gin-gonic/gin"

	// insertion point for gong front end import
	_ "github.com/fullstack-lang/gong/ng"
)

func Init(r *gin.Engine) {

	db_inMemory := gong_orm.SetupModels(false, ":memory:")

	// since gongsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db_inMemory.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	gong_controllers.RegisterControllers(r)
}
