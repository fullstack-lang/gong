package golang

const FullstackInitTemplate = `package fullstack

import (
	// {{pkgname}} stack for model analysis

	{{pkgname}}_controllers "{{PkgPathRoot}}/controllers"
	{{pkgname}}_models "{{PkgPathRoot}}/models"
	{{pkgname}}_orm "{{PkgPathRoot}}/orm"
	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "{{PkgPathAboveRoot}}/ng/projects"
)

func Init(r *gin.Engine, filenames ...string) {

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	db_inMemory := {{pkgname}}_orm.SetupModels(&{{pkgname}}_models.Stage, false, filenames[0])

	// since {{pkgname}}sim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db_inMemory.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	{{pkgname}}_controllers.RegisterControllers(r)
}
`
