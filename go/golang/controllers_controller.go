package golang

const ControllerTemplate = `package controllers

import (
	"sync"

	{{pkgname}}_orm "{{PkgPathRoot}}/orm"

	"github.com/gin-gonic/gin"
)

// A Controller is the handler of all API REST calls matching the stack model
// It forwards API requests to the stack instance identified by the GONG_StackPath parameters in the request
// the stack instance is the BackRepo instance
type Controller struct {

	// Map_BackRepos is the map to the backRepo instance according to the stack instance path
	Map_BackRepos map[string]*{{pkgname}}_orm.BackRepoStruct
}

var instance *Controller
var once sync.Once

func Register(r *gin.Engine) {
	once.Do(func() {
		RegisterControllers(r)
	})
}

func GetController() *Controller {
	once.Do(func() {
		instance = &Controller{
			Map_BackRepos: make(map[string]*{{pkgname}}_orm.BackRepoStruct),
		}
		instance.Map_BackRepos[""] = &{{pkgname}}_orm.BackRepo
	})
	return instance
}
`
