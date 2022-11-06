package fullstack

import (
	// test2 stack for model analysis

	test2_controllers "github.com/fullstack-lang/gong/test2/go/controllers"
	test2_orm "github.com/fullstack-lang/gong/test2/go/orm"
	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gong/test2/ng"
)

func Init(r *gin.Engine, filenames ...string) {

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	db_inMemory := test2_orm.SetupModels(false, filenames[0])

	// since test2sim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db_inMemory.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	test2_controllers.RegisterControllers(r)
}
