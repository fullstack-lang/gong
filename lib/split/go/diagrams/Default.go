package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/split/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__GongStructShape__000000_Default_AsSplit := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_AsSplitArea := (&models.GongStructShape{}).Stage(stage)

	__Position__000000_Pos_Default_AsSplit := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_AsSplitArea := (&models.Position{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__GongStructShape__000000_Default_AsSplit.Name = `Default-AsSplit`

	//gong:ident [ref_models.AsSplit] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_AsSplit.Identifier = `ref_models.AsSplit`
	__GongStructShape__000000_Default_AsSplit.ShowNbInstances = false
	__GongStructShape__000000_Default_AsSplit.NbInstances = 0
	__GongStructShape__000000_Default_AsSplit.Width = 240.000000
	__GongStructShape__000000_Default_AsSplit.Height = 63.000000
	__GongStructShape__000000_Default_AsSplit.IsSelected = false

	__GongStructShape__000001_Default_AsSplitArea.Name = `Default-AsSplitArea`

	//gong:ident [ref_models.AsSplitArea] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_AsSplitArea.Identifier = `ref_models.AsSplitArea`
	__GongStructShape__000001_Default_AsSplitArea.ShowNbInstances = false
	__GongStructShape__000001_Default_AsSplitArea.NbInstances = 0
	__GongStructShape__000001_Default_AsSplitArea.Width = 240.000000
	__GongStructShape__000001_Default_AsSplitArea.Height = 63.000000
	__GongStructShape__000001_Default_AsSplitArea.IsSelected = false

	__Position__000000_Pos_Default_AsSplit.X = 334.000000
	__Position__000000_Pos_Default_AsSplit.Y = 91.000000
	__Position__000000_Pos_Default_AsSplit.Name = `Pos-Default-AsSplit`

	__Position__000001_Pos_Default_AsSplitArea.X = 50.000000
	__Position__000001_Pos_Default_AsSplitArea.Y = 5.000000
	__Position__000001_Pos_Default_AsSplitArea.Name = `Pos-Default-AsSplitArea`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_AsSplit)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_AsSplitArea)
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_AsSplit.Position = __Position__000000_Pos_Default_AsSplit
	__GongStructShape__000001_Default_AsSplitArea.Position = __Position__000001_Pos_Default_AsSplitArea
	// setup of Position instances pointers
}
