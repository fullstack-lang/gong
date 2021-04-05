package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/examples/laundromat/go/orm"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gong/examples/laundromat/go")
	{ // insertion point for registrations
		v1.GET("/v1/machines", GetMachines)
		v1.GET("/v1/machines/:id", GetMachine)
		v1.POST("/v1/machines", PostMachine)
		v1.PATCH("/v1/machines/:id", UpdateMachine)
		v1.PUT("/v1/machines/:id", UpdateMachine)
		v1.DELETE("/v1/machines/:id", DeleteMachine)

		v1.GET("/v1/simulations", GetSimulations)
		v1.GET("/v1/simulations/:id", GetSimulation)
		v1.POST("/v1/simulations", PostSimulation)
		v1.PATCH("/v1/simulations/:id", UpdateSimulation)
		v1.PUT("/v1/simulations/:id", UpdateSimulation)
		v1.DELETE("/v1/simulations/:id", DeleteSimulation)

		v1.GET("/v1/washers", GetWashers)
		v1.GET("/v1/washers/:id", GetWasher)
		v1.POST("/v1/washers", PostWasher)
		v1.PATCH("/v1/washers/:id", UpdateWasher)
		v1.PUT("/v1/washers/:id", UpdateWasher)
		v1.DELETE("/v1/washers/:id", DeleteWasher)

		v1.GET("/commitnb", GetLastCommitNb)
	}
}

// swagger:route GET /commitnb backrepo GetLastCommitNb
func GetLastCommitNb(c *gin.Context) {
	res := orm.GetLastCommitNb()

	c.JSON(http.StatusOK, res)
}
