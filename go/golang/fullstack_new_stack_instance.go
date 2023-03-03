package golang

const FullstackNewStackInstanceTemplate = `package fullstack

import (
	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/orm"

	"github.com/gin-gonic/gin"
)

// NewStackInstance creates a new stack instance from the Stack Model
// and returns the backRepo of the stack instance (you can get the stage from backRepo.GetStage()
//
// - the stackPath is the unique identifier of the stack
// - the optional parameter filenames is for the name of the database filename
// if filenames is omitted, the database is persisted in memory
func NewStackInstance(
	r *gin.Engine,
	stackPath string,
	// filesnames is an optional parameter for the name of the database
	filenames ...string) (
	stage *models.StageStruct,
	backRepo *orm.BackRepoStruct) {

	Init(r, filenames...)

	// temporary
	stage = &models.Stage
	backRepo = &orm.BackRepo

	return

}
`
