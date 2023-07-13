package data

import (
	"github.com/gin-gonic/gin"

	gongrouter_fullstack "github.com/fullstack-lang/gongrouter/go/fullstack"
	gongtree_fullstack "github.com/fullstack-lang/gongtree/go/fullstack"
)

func Load(r *gin.Engine, stackPath string) {

	gongtreeStage := gongtree_fullstack.NewStackInstance(r, stackPath)
	_ = gongtreeStage
	gongrouterStage := gongrouter_fullstack.NewStackInstance(r, stackPath)
	_ = gongrouterStage
}
