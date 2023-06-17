package controllers

import (
	"sync"

	gongtree_orm "github.com/fullstack-lang/gongtree/go/orm"

	"github.com/gin-gonic/gin"
)

// A Controller is the handler of all API REST calls matching the stack model
// It forwards API requests to the stack instance identified by the GONG_StackPath parameters in the request
// the stack instance is the BackRepo instance
type Controller struct {

	// Map_BackRepos is the map to the backRepo instance according to the stack instance path
	Map_BackRepos map[string]*gongtree_orm.BackRepoStruct
}

var _controllerSingloton *Controller
var doRegisterOnce sync.Once

func Register(r *gin.Engine) {
	doRegisterOnce.Do(func() {
		registerControllers(r)
	})
}

var doControllerInitOnce sync.Once

func GetController() *Controller {
	doControllerInitOnce.Do(func() {
		_controllerSingloton = &Controller{
			Map_BackRepos: make(map[string]*gongtree_orm.BackRepoStruct),
		}
		_controllerSingloton.Map_BackRepos[""] = gongtree_orm.GetDefaultBackRepo()
	})
	return _controllerSingloton
}

func (controller *Controller) AddBackRepo(backRepo *gongtree_orm.BackRepoStruct, stackPath string) {
	GetController().Map_BackRepos[stackPath] = backRepo
}
