package fullstack

const FullstackNewStackInstanceTemplate = //
FullstackNewStackInstanceTemplatePart1 +
	FullstackNewStackInstanceTemplatePart2 +
	FullstackNewStackInstanceTemplatePart3

const FullstackNewStackInstanceTemplateLevel1 = //
FullstackNewStackInstanceTemplatePart1 +
	FullstackNewStackInstanceTemplatePart2Level1 +
	FullstackNewStackInstanceTemplatePart3

const FullstackNewStackInstanceTemplatePart1 = `// do not modify, generated file
package fullstack

import (
	"{{PkgPathRoot}}/controllers"
	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code`

const FullstackNewStackInstanceTemplatePart2 = `
	_ "{{PkgPathAboveRoot}}/{{NgWorkspaceName}}"
)`

const FullstackNewStackInstanceTemplatePart2Level1 = `
	// This is a level 1 gong application, no need to import the angular code
	// therefore, the following line that is necessary in level 2 applications, is commented
	// _ "{{PkgPathAboveRoot}}/{{NgWorkspaceName}}"
)`

const FullstackNewStackInstanceTemplatePart3 = `

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
	stage *models.Stage,
	backRepo *orm.BackRepoStruct) {

	stage = models.NewStage(stackPath)

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo = orm.NewBackRepo(stage, filenames[0])

	controllers.GetController().AddBackRepo(backRepo, stackPath)

	controllers.Register(r)

	// add orchestration
	// insertion point{{` + string(rune(ModelGongNewStackInstanceSet)) + `}}

	return
}
`

type ModelGongNewStackInstanceStructInsertionId int

const (
	ModelGongNewStackInstanceSet ModelGongNewStackInstanceStructInsertionId = iota
)

var ModelGongNewStackInstanceStructSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(ModelGongNewStackInstanceSet)): `
	models.SetOrchestratorOnAfterUpdate[models.{{Structname}}](stage)`,
}
