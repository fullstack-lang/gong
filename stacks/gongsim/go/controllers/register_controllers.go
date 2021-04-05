package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/stacks/gongsim/go/orm"
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
	v1 := r.Group("/api/github.com/fullstack-lang/gong/stacks/gongsim/go")
	{ // insertion point for registrations
		v1.GET("/v1/engines", GetEngines)
		v1.GET("/v1/engines/:id", GetEngine)
		v1.POST("/v1/engines", PostEngine)
		v1.PATCH("/v1/engines/:id", UpdateEngine)
		v1.PUT("/v1/engines/:id", UpdateEngine)
		v1.DELETE("/v1/engines/:id", DeleteEngine)

		v1.GET("/v1/events", GetEvents)
		v1.GET("/v1/events/:id", GetEvent)
		v1.POST("/v1/events", PostEvent)
		v1.PATCH("/v1/events/:id", UpdateEvent)
		v1.PUT("/v1/events/:id", UpdateEvent)
		v1.DELETE("/v1/events/:id", DeleteEvent)

		v1.GET("/v1/gongsimcommands", GetGongsimCommands)
		v1.GET("/v1/gongsimcommands/:id", GetGongsimCommand)
		v1.POST("/v1/gongsimcommands", PostGongsimCommand)
		v1.PATCH("/v1/gongsimcommands/:id", UpdateGongsimCommand)
		v1.PUT("/v1/gongsimcommands/:id", UpdateGongsimCommand)
		v1.DELETE("/v1/gongsimcommands/:id", DeleteGongsimCommand)

		v1.GET("/v1/gongsimstatuss", GetGongsimStatuss)
		v1.GET("/v1/gongsimstatuss/:id", GetGongsimStatus)
		v1.POST("/v1/gongsimstatuss", PostGongsimStatus)
		v1.PATCH("/v1/gongsimstatuss/:id", UpdateGongsimStatus)
		v1.PUT("/v1/gongsimstatuss/:id", UpdateGongsimStatus)
		v1.DELETE("/v1/gongsimstatuss/:id", DeleteGongsimStatus)

		v1.GET("/v1/updatestates", GetUpdateStates)
		v1.GET("/v1/updatestates/:id", GetUpdateState)
		v1.POST("/v1/updatestates", PostUpdateState)
		v1.PATCH("/v1/updatestates/:id", UpdateUpdateState)
		v1.PUT("/v1/updatestates/:id", UpdateUpdateState)
		v1.DELETE("/v1/updatestates/:id", DeleteUpdateState)

		v1.GET("/commitnb", GetLastCommitNb)
	}
}

// swagger:route GET /commitnb backrepo GetLastCommitNb
func GetLastCommitNb(c *gin.Context) {
	res := orm.GetLastCommitNb()

	c.JSON(http.StatusOK, res)
}
