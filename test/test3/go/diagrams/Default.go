package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/test/test3/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__Field__000000_Name := (&models.Field{}).Stage(stage)

	__GongStructShape__000000_Default_A := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_B := (&models.GongStructShape{}).Stage(stage)

	__Position__000000_Pos_Default_A := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_B := (&models.Position{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Name.Name = `Name`

	//gong:ident [ref_models.A.Name] comment added to overcome the problem with the comment map association
	__Field__000000_Name.Identifier = `ref_models.A.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `A`
	__Field__000000_Name.Fieldtypename = `string`

	__GongStructShape__000000_Default_A.Name = `Default-A`

	//gong:ident [ref_models.A] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_A.Identifier = `ref_models.A`
	__GongStructShape__000000_Default_A.ShowNbInstances = false
	__GongStructShape__000000_Default_A.NbInstances = 0
	__GongStructShape__000000_Default_A.Width = 240.000000
	__GongStructShape__000000_Default_A.Height = 78.000000
	__GongStructShape__000000_Default_A.IsSelected = false

	__GongStructShape__000001_Default_B.Name = `Default-B`

	//gong:ident [ref_models.B] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_B.Identifier = `ref_models.B`
	__GongStructShape__000001_Default_B.ShowNbInstances = false
	__GongStructShape__000001_Default_B.NbInstances = 0
	__GongStructShape__000001_Default_B.Width = 240.000000
	__GongStructShape__000001_Default_B.Height = 63.000000
	__GongStructShape__000001_Default_B.IsSelected = false

	__Position__000000_Pos_Default_A.X = 485.000000
	__Position__000000_Pos_Default_A.Y = 71.000000
	__Position__000000_Pos_Default_A.Name = `Pos-Default-A`

	__Position__000001_Pos_Default_B.X = 16.000000
	__Position__000001_Pos_Default_B.Y = 91.000000
	__Position__000001_Pos_Default_B.Name = `Pos-Default-B`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_A)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_B)
	__GongStructShape__000000_Default_A.Position = __Position__000000_Pos_Default_A
	__GongStructShape__000000_Default_A.Fields = append(__GongStructShape__000000_Default_A.Fields, __Field__000000_Name)
	__GongStructShape__000001_Default_B.Position = __Position__000001_Pos_Default_B
}
