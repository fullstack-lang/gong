package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/doc/go/models"
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

	__Field__000000_Name := (&models.Field{}).Stage(stage)
	__Field__000001_Label := (&models.Field{}).Stage(stage)
	__Field__000002_Icon := (&models.Field{}).Stage(stage)
	__Field__000003_Name := (&models.Field{}).Stage(stage)
	__Field__000004_Percentage := (&models.Field{}).Stage(stage)
	__Field__000005_Name := (&models.Field{}).Stage(stage)

	__GongStructShape__000000_Default_Button := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Group := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Layout := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_Buttons := (&models.Link{}).Stage(stage)
	__Link__000001_Groups := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_Default_Button := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_Group := (&models.Position{}).Stage(stage)
	__Position__000002_Pos_Default_Layout := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Button := (&models.Vertice{}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Layout_and_Default_Group := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Name.Name = `Name`

	//gong:ident [ref_models.Button.Name] comment added to overcome the problem with the comment map association
	__Field__000000_Name.Identifier = `ref_models.Button.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Button`
	__Field__000000_Name.Fieldtypename = `string`

	__Field__000001_Label.Name = `Label`

	//gong:ident [ref_models.Button.Label] comment added to overcome the problem with the comment map association
	__Field__000001_Label.Identifier = `ref_models.Button.Label`
	__Field__000001_Label.FieldTypeAsString = ``
	__Field__000001_Label.Structname = `Button`
	__Field__000001_Label.Fieldtypename = `string`

	__Field__000002_Icon.Name = `Icon`

	//gong:ident [ref_models.Button.Icon] comment added to overcome the problem with the comment map association
	__Field__000002_Icon.Identifier = `ref_models.Button.Icon`
	__Field__000002_Icon.FieldTypeAsString = ``
	__Field__000002_Icon.Structname = `Button`
	__Field__000002_Icon.Fieldtypename = `string`

	__Field__000003_Name.Name = `Name`

	//gong:ident [ref_models.Group.Name] comment added to overcome the problem with the comment map association
	__Field__000003_Name.Identifier = `ref_models.Group.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `Group`
	__Field__000003_Name.Fieldtypename = `string`

	__Field__000004_Percentage.Name = `Percentage`

	//gong:ident [ref_models.Group.Percentage] comment added to overcome the problem with the comment map association
	__Field__000004_Percentage.Identifier = `ref_models.Group.Percentage`
	__Field__000004_Percentage.FieldTypeAsString = ``
	__Field__000004_Percentage.Structname = `Group`
	__Field__000004_Percentage.Fieldtypename = `float64`

	__Field__000005_Name.Name = `Name`

	//gong:ident [ref_models.Layout.Name] comment added to overcome the problem with the comment map association
	__Field__000005_Name.Identifier = `ref_models.Layout.Name`
	__Field__000005_Name.FieldTypeAsString = ``
	__Field__000005_Name.Structname = `Layout`
	__Field__000005_Name.Fieldtypename = `string`

	__GongStructShape__000000_Default_Button.Name = `Default-Button`

	//gong:ident [ref_models.Button] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Button.Identifier = `ref_models.Button`
	__GongStructShape__000000_Default_Button.ShowNbInstances = false
	__GongStructShape__000000_Default_Button.NbInstances = 0
	__GongStructShape__000000_Default_Button.Width = 240.000000
	__GongStructShape__000000_Default_Button.Height = 108.000000
	__GongStructShape__000000_Default_Button.IsSelected = false

	__GongStructShape__000001_Default_Group.Name = `Default-Group`

	//gong:ident [ref_models.Group] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Group.Identifier = `ref_models.Group`
	__GongStructShape__000001_Default_Group.ShowNbInstances = false
	__GongStructShape__000001_Default_Group.NbInstances = 0
	__GongStructShape__000001_Default_Group.Width = 240.000000
	__GongStructShape__000001_Default_Group.Height = 93.000000
	__GongStructShape__000001_Default_Group.IsSelected = false

	__GongStructShape__000002_Default_Layout.Name = `Default-Layout`

	//gong:ident [ref_models.Layout] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Layout.Identifier = `ref_models.Layout`
	__GongStructShape__000002_Default_Layout.ShowNbInstances = false
	__GongStructShape__000002_Default_Layout.NbInstances = 0
	__GongStructShape__000002_Default_Layout.Width = 240.000000
	__GongStructShape__000002_Default_Layout.Height = 78.000000
	__GongStructShape__000002_Default_Layout.IsSelected = false

	__Link__000000_Buttons.Name = `Buttons`

	//gong:ident [ref_models.Group.Buttons] comment added to overcome the problem with the comment map association
	__Link__000000_Buttons.Identifier = `ref_models.Group.Buttons`

	//gong:ident [ref_models.Button] comment added to overcome the problem with the comment map association
	__Link__000000_Buttons.Fieldtypename = `ref_models.Button`
	__Link__000000_Buttons.FieldOffsetX = 0.000000
	__Link__000000_Buttons.FieldOffsetY = 0.000000
	__Link__000000_Buttons.TargetMultiplicity = models.MANY
	__Link__000000_Buttons.TargetMultiplicityOffsetX = 0.000000
	__Link__000000_Buttons.TargetMultiplicityOffsetY = 0.000000
	__Link__000000_Buttons.SourceMultiplicity = models.MANY
	__Link__000000_Buttons.SourceMultiplicityOffsetX = 0.000000
	__Link__000000_Buttons.SourceMultiplicityOffsetY = 0.000000
	__Link__000000_Buttons.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Buttons.StartRatio = 0.500000
	__Link__000000_Buttons.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Buttons.EndRatio = 0.500000
	__Link__000000_Buttons.CornerOffsetRatio = 1.380000

	__Link__000001_Groups.Name = `Groups`

	//gong:ident [ref_models.Layout.Groups] comment added to overcome the problem with the comment map association
	__Link__000001_Groups.Identifier = `ref_models.Layout.Groups`

	//gong:ident [ref_models.Group] comment added to overcome the problem with the comment map association
	__Link__000001_Groups.Fieldtypename = `ref_models.Group`
	__Link__000001_Groups.FieldOffsetX = 0.000000
	__Link__000001_Groups.FieldOffsetY = 0.000000
	__Link__000001_Groups.TargetMultiplicity = models.MANY
	__Link__000001_Groups.TargetMultiplicityOffsetX = 0.000000
	__Link__000001_Groups.TargetMultiplicityOffsetY = 0.000000
	__Link__000001_Groups.SourceMultiplicity = models.MANY
	__Link__000001_Groups.SourceMultiplicityOffsetX = 0.000000
	__Link__000001_Groups.SourceMultiplicityOffsetY = 0.000000
	__Link__000001_Groups.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Groups.StartRatio = 0.500000
	__Link__000001_Groups.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Groups.EndRatio = 0.500000
	__Link__000001_Groups.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_Button.X = 989.000000
	__Position__000000_Pos_Default_Button.Y = 30.000000
	__Position__000000_Pos_Default_Button.Name = `Pos-Default-Button`

	__Position__000001_Pos_Default_Group.X = 540.000000
	__Position__000001_Pos_Default_Group.Y = 26.000000
	__Position__000001_Pos_Default_Group.Name = `Pos-Default-Group`

	__Position__000002_Pos_Default_Layout.X = 47.000000
	__Position__000002_Pos_Default_Layout.Y = 26.000000
	__Position__000002_Pos_Default_Layout.Name = `Pos-Default-Layout`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Button.X = 586.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Button.Y = 98.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Button.Name = `Verticle in class diagram Default in middle between Default-Group and Default-Button`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Layout_and_Default_Group.X = 617.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Layout_and_Default_Group.Y = 86.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Layout_and_Default_Group.Name = `Verticle in class diagram Default in middle between Default-Layout and Default-Group`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Button)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Group)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Layout)
	// setup of Field instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Button.Position = __Position__000000_Pos_Default_Button
	__GongStructShape__000000_Default_Button.Fields = append(__GongStructShape__000000_Default_Button.Fields, __Field__000000_Name)
	__GongStructShape__000000_Default_Button.Fields = append(__GongStructShape__000000_Default_Button.Fields, __Field__000001_Label)
	__GongStructShape__000000_Default_Button.Fields = append(__GongStructShape__000000_Default_Button.Fields, __Field__000002_Icon)
	__GongStructShape__000001_Default_Group.Position = __Position__000001_Pos_Default_Group
	__GongStructShape__000001_Default_Group.Fields = append(__GongStructShape__000001_Default_Group.Fields, __Field__000003_Name)
	__GongStructShape__000001_Default_Group.Fields = append(__GongStructShape__000001_Default_Group.Fields, __Field__000004_Percentage)
	__GongStructShape__000001_Default_Group.Links = append(__GongStructShape__000001_Default_Group.Links, __Link__000000_Buttons)
	__GongStructShape__000002_Default_Layout.Position = __Position__000002_Pos_Default_Layout
	__GongStructShape__000002_Default_Layout.Fields = append(__GongStructShape__000002_Default_Layout.Fields, __Field__000005_Name)
	__GongStructShape__000002_Default_Layout.Links = append(__GongStructShape__000002_Default_Layout.Links, __Link__000001_Groups)
	// setup of Link instances pointers
	__Link__000000_Buttons.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Button
	__Link__000001_Groups.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Layout_and_Default_Group
	// setup of Position instances pointers
	// setup of Vertice instances pointers
}
