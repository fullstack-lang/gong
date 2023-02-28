package controllers

import (
	"sync"

	"github.com/fullstack-lang/gong/test/go/orm"

	"github.com/gin-gonic/gin"
)

// A Controller is the handler of all API REST calls matching the stack model
// It forwards API requests to the stack instance identified by the GONG_StackPath parameters in the request
// the stack instance is the BackRepo instance
type Controller struct {

	// Map_BackRepos is the map to the backRepo instance according to the stack instance path
	Map_BackRepos map[string]*orm.BackRepoStruct
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
			Map_BackRepos: make(map[string]*orm.BackRepoStruct),
		}
	})
	return instance
}
