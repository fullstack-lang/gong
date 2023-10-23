// do not modify, generated file
package fullstack

import (
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/test2/go/models"
	"github.com/fullstack-lang/gong/test2/go/models/x"

	"github.com/fullstack-lang/gong/test2/go/orm"
	orm_x "github.com/fullstack-lang/gong/test2/go/orm/x"

	"github.com/fullstack-lang/gong/test2/go/controllers"
	controllers_x "github.com/fullstack-lang/gong/test2/go/controllers/x"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gong/test2/ng/projects"
)

type Stack struct {
	root *models.StageStruct
	x    *x.StageStruct

	ORM
}

type ORM struct {
	root *orm.BackRepoStruct
	x    *orm_x.BackRepoStruct
}

func (stack *Stack) Commit() {
	// insertion point for per struct back repo phase one commit
	stack.ORM.root.BackRepoA.CommitPhaseOne(stack.root)
	stack.ORM.root.BackRepoB.CommitPhaseOne(stack.root)

	stack.ORM.x.BackRepoA.CommitPhaseOne(stack.x)
	stack.ORM.x.BackRepoB.CommitPhaseOne(stack.x)

	// insertion point for per struct back repo phase two commit
	stack.ORM.root.BackRepoA.CommitPhaseTwo(stack.ORM.root)
	stack.ORM.root.BackRepoB.CommitPhaseTwo(stack.ORM.root)

	stack.ORM.x.BackRepoA.CommitPhaseTwo(stack.ORM.x)
	stack.ORM.x.BackRepoB.CommitPhaseTwo(stack.ORM.x)

	stack.ORM.root.IncrementCommitFromBackNb()
	stack.ORM.x.IncrementCommitFromBackNb()
}

func (stack *Stack) Checkout() {
	stack.ORM.root.BackRepoA.CheckoutPhaseOne()
	stack.ORM.root.BackRepoB.CheckoutPhaseOne()

	stack.ORM.x.BackRepoA.CheckoutPhaseOne()
	stack.ORM.x.BackRepoB.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	stack.ORM.root.BackRepoA.CheckoutPhaseTwo(stack.ORM.root)
	stack.ORM.root.BackRepoB.CheckoutPhaseTwo(stack.ORM.root)

	stack.ORM.x.BackRepoA.CheckoutPhaseTwo(stack.ORM.x)
	stack.ORM.x.BackRepoB.CheckoutPhaseTwo(stack.ORM.x)
}

// NewStackInstance creates a new stack instance from the Stack Model
// and returns the backRepo of the stack instance (you can get the stage from backRepo.GetStage()
//
// - the stackPath is the unique identifier of the stack
// - the optional parameter filenames is for the name of the database filename
// if filenames is omitted, the database is persisted in memory
func NewStack(
	r *gin.Engine,
	stackPath string,
	hasProbe bool,
	// filesnames is an optional parameter for the name of the database
	filenames ...string) (stack *Stack) {

	stack = new(Stack)

	stack.root = models.NewStage(stackPath)

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	stack.ORM.root = orm.NewBackRepo(stack.root, filenames[0])

	if stackPath != "" {
		controllers.GetController().AddBackRepo(stack.ORM.root, stackPath)
		controllers_x.GetController().AddBackRepo(stack.ORM.x, stackPath)
	}

	controllers.Register(r)

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.A](stack.root)
	models.SetOrchestratorOnAfterUpdate[models.B](stack.root)

	x.SetOrchestratorOnAfterUpdate[x.A](stack.x)
	x.SetOrchestratorOnAfterUpdate[x.B](stack.x)

	return
}
