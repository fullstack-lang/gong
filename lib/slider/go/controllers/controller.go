// generated code - do not edit
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	slider_orm "github.com/fullstack-lang/gong/lib/slider/go/orm"
)

type H map[string]any

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writeJSON error: %v", err)
	}
}

// A Controller is the handler of all API REST calls matching the stack model
// It forwards API requests to the stack instance identified by the Name parameters in the request
// the stack instance is the BackRepo instance
type Controller struct {

	// Map_BackRepos is the map to the backRepo instance according to the stack instance path
	Map_BackRepos map[string]*slider_orm.BackRepoStruct

	listenerIndex int // Counter to track the number of listeners
}

var _controllerSingloton *Controller
var doRegisterOnce sync.Once

func Register(mux *http.ServeMux) {
	if mux == nil {
		return
	}
	doRegisterOnce.Do(func() {
		registerControllers(mux)
	})
}

var doControllerInitOnce sync.Once

func GetController() *Controller {
	doControllerInitOnce.Do(func() {
		_controllerSingloton = &Controller{
			Map_BackRepos: make(map[string]*slider_orm.BackRepoStruct),
		}
	})
	return _controllerSingloton
}

func (controller *Controller) AddBackRepo(backRepo *slider_orm.BackRepoStruct, stackPath string) {
	GetController().Map_BackRepos[stackPath] = backRepo
}
