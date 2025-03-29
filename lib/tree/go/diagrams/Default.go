package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/tree/go/models"
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
	__Field__000001_Name := (&models.Field{}).Stage(stage)
	__Field__000002_IsExpanded := (&models.Field{}).Stage(stage)
	__Field__000003_HasCheckboxButton := (&models.Field{}).Stage(stage)
	__Field__000004_IsChecked := (&models.Field{}).Stage(stage)
	__Field__000005_Name := (&models.Field{}).Stage(stage)
	__Field__000006_Icon := (&models.Field{}).Stage(stage)
	__Field__000007_Name := (&models.Field{}).Stage(stage)
	__Field__000008_SVG := (&models.Field{}).Stage(stage)

	__GongStructShape__000000_Default_Tree := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Node := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Button := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_SVGIcon := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_RootNodes := (&models.Link{}).Stage(stage)
	__Link__000001_Children := (&models.Link{}).Stage(stage)
	__Link__000002_Buttons := (&models.Link{}).Stage(stage)
	__Link__000003_SVGIcon := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_Default_Tree := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_Node := (&models.Position{}).Stage(stage)
	__Position__000002_Pos_Default_Button := (&models.Position{}).Stage(stage)
	__Position__000003_Pos_Default_SVGIcon := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node := (&models.Vertice{}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node := (&models.Vertice{}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button := (&models.Vertice{}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Button_and_Default_SVGIcon := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Name.Name = `Name`

	//gong:ident [ref_models.Tree.Name] comment added to overcome the problem with the comment map association
	__Field__000000_Name.Identifier = `ref_models.Tree.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Tree`
	__Field__000000_Name.Fieldtypename = `string`

	__Field__000001_Name.Name = `Name`

	//gong:ident [ref_models.Node.Name] comment added to overcome the problem with the comment map association
	__Field__000001_Name.Identifier = `ref_models.Node.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Node`
	__Field__000001_Name.Fieldtypename = `string`

	__Field__000002_IsExpanded.Name = `IsExpanded`

	//gong:ident [ref_models.Node.IsExpanded] comment added to overcome the problem with the comment map association
	__Field__000002_IsExpanded.Identifier = `ref_models.Node.IsExpanded`
	__Field__000002_IsExpanded.FieldTypeAsString = ``
	__Field__000002_IsExpanded.Structname = `Node`
	__Field__000002_IsExpanded.Fieldtypename = `bool`

	__Field__000003_HasCheckboxButton.Name = `HasCheckboxButton`

	//gong:ident [ref_models.Node.HasCheckboxButton] comment added to overcome the problem with the comment map association
	__Field__000003_HasCheckboxButton.Identifier = `ref_models.Node.HasCheckboxButton`
	__Field__000003_HasCheckboxButton.FieldTypeAsString = ``
	__Field__000003_HasCheckboxButton.Structname = `Node`
	__Field__000003_HasCheckboxButton.Fieldtypename = `bool`

	__Field__000004_IsChecked.Name = `IsChecked`

	//gong:ident [ref_models.Node.IsChecked] comment added to overcome the problem with the comment map association
	__Field__000004_IsChecked.Identifier = `ref_models.Node.IsChecked`
	__Field__000004_IsChecked.FieldTypeAsString = ``
	__Field__000004_IsChecked.Structname = `Node`
	__Field__000004_IsChecked.Fieldtypename = `bool`

	__Field__000005_Name.Name = `Name`

	//gong:ident [ref_models.Button.Name] comment added to overcome the problem with the comment map association
	__Field__000005_Name.Identifier = `ref_models.Button.Name`
	__Field__000005_Name.FieldTypeAsString = ``
	__Field__000005_Name.Structname = `Button`
	__Field__000005_Name.Fieldtypename = `string`

	__Field__000006_Icon.Name = `Icon`

	//gong:ident [ref_models.Button.Icon] comment added to overcome the problem with the comment map association
	__Field__000006_Icon.Identifier = `ref_models.Button.Icon`
	__Field__000006_Icon.FieldTypeAsString = ``
	__Field__000006_Icon.Structname = `Button`
	__Field__000006_Icon.Fieldtypename = `string`

	__Field__000007_Name.Name = `Name`

	//gong:ident [ref_models.SVGIcon.Name] comment added to overcome the problem with the comment map association
	__Field__000007_Name.Identifier = `ref_models.SVGIcon.Name`
	__Field__000007_Name.FieldTypeAsString = ``
	__Field__000007_Name.Structname = `SVGIcon`
	__Field__000007_Name.Fieldtypename = `string`

	__Field__000008_SVG.Name = `SVG`

	//gong:ident [ref_models.SVGIcon.SVG] comment added to overcome the problem with the comment map association
	__Field__000008_SVG.Identifier = `ref_models.SVGIcon.SVG`
	__Field__000008_SVG.FieldTypeAsString = ``
	__Field__000008_SVG.Structname = `SVGIcon`
	__Field__000008_SVG.Fieldtypename = `string`

	__GongStructShape__000000_Default_Tree.Name = `Default-Tree`

	//gong:ident [ref_models.Tree] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Tree.Identifier = `ref_models.Tree`
	__GongStructShape__000000_Default_Tree.ShowNbInstances = false
	__GongStructShape__000000_Default_Tree.NbInstances = 0
	__GongStructShape__000000_Default_Tree.Width = 240.000000
	__GongStructShape__000000_Default_Tree.Height = 78.000000
	__GongStructShape__000000_Default_Tree.IsSelected = false

	__GongStructShape__000001_Default_Node.Name = `Default-Node`

	//gong:ident [ref_models.Node] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Node.Identifier = `ref_models.Node`
	__GongStructShape__000001_Default_Node.ShowNbInstances = false
	__GongStructShape__000001_Default_Node.NbInstances = 0
	__GongStructShape__000001_Default_Node.Width = 240.000000
	__GongStructShape__000001_Default_Node.Height = 123.000000
	__GongStructShape__000001_Default_Node.IsSelected = false

	__GongStructShape__000002_Default_Button.Name = `Default-Button`

	//gong:ident [ref_models.Button] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Button.Identifier = `ref_models.Button`
	__GongStructShape__000002_Default_Button.ShowNbInstances = false
	__GongStructShape__000002_Default_Button.NbInstances = 0
	__GongStructShape__000002_Default_Button.Width = 240.000000
	__GongStructShape__000002_Default_Button.Height = 93.000000
	__GongStructShape__000002_Default_Button.IsSelected = false

	__GongStructShape__000003_Default_SVGIcon.Name = `Default-SVGIcon`

	//gong:ident [ref_models.SVGIcon] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_SVGIcon.Identifier = `ref_models.SVGIcon`
	__GongStructShape__000003_Default_SVGIcon.ShowNbInstances = false
	__GongStructShape__000003_Default_SVGIcon.NbInstances = 0
	__GongStructShape__000003_Default_SVGIcon.Width = 240.000000
	__GongStructShape__000003_Default_SVGIcon.Height = 93.000000
	__GongStructShape__000003_Default_SVGIcon.IsSelected = false

	__Link__000000_RootNodes.Name = `RootNodes`

	//gong:ident [ref_models.Tree.RootNodes] comment added to overcome the problem with the comment map association
	__Link__000000_RootNodes.Identifier = `ref_models.Tree.RootNodes`

	//gong:ident [ref_models.Node] comment added to overcome the problem with the comment map association
	__Link__000000_RootNodes.Fieldtypename = `ref_models.Node`
	__Link__000000_RootNodes.FieldOffsetX = 0.000000
	__Link__000000_RootNodes.FieldOffsetY = 0.000000
	__Link__000000_RootNodes.TargetMultiplicity = models.MANY
	__Link__000000_RootNodes.TargetMultiplicityOffsetX = 0.000000
	__Link__000000_RootNodes.TargetMultiplicityOffsetY = 0.000000
	__Link__000000_RootNodes.SourceMultiplicity = models.MANY
	__Link__000000_RootNodes.SourceMultiplicityOffsetX = 0.000000
	__Link__000000_RootNodes.SourceMultiplicityOffsetY = 0.000000
	__Link__000000_RootNodes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_RootNodes.StartRatio = 0.500000
	__Link__000000_RootNodes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_RootNodes.EndRatio = 0.500000
	__Link__000000_RootNodes.CornerOffsetRatio = 1.380000

	__Link__000001_Children.Name = `Children`

	//gong:ident [ref_models.Node.Children] comment added to overcome the problem with the comment map association
	__Link__000001_Children.Identifier = `ref_models.Node.Children`

	//gong:ident [ref_models.Node] comment added to overcome the problem with the comment map association
	__Link__000001_Children.Fieldtypename = `ref_models.Node`
	__Link__000001_Children.FieldOffsetX = 0.000000
	__Link__000001_Children.FieldOffsetY = 0.000000
	__Link__000001_Children.TargetMultiplicity = models.MANY
	__Link__000001_Children.TargetMultiplicityOffsetX = 0.000000
	__Link__000001_Children.TargetMultiplicityOffsetY = 0.000000
	__Link__000001_Children.SourceMultiplicity = models.MANY
	__Link__000001_Children.SourceMultiplicityOffsetX = 0.000000
	__Link__000001_Children.SourceMultiplicityOffsetY = 0.000000
	__Link__000001_Children.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Children.StartRatio = 0.198584
	__Link__000001_Children.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Children.EndRatio = 0.751429
	__Link__000001_Children.CornerOffsetRatio = 1.380000

	__Link__000002_Buttons.Name = `Buttons`

	//gong:ident [ref_models.Node.Buttons] comment added to overcome the problem with the comment map association
	__Link__000002_Buttons.Identifier = `ref_models.Node.Buttons`

	//gong:ident [ref_models.Button] comment added to overcome the problem with the comment map association
	__Link__000002_Buttons.Fieldtypename = `ref_models.Button`
	__Link__000002_Buttons.FieldOffsetX = 0.000000
	__Link__000002_Buttons.FieldOffsetY = 0.000000
	__Link__000002_Buttons.TargetMultiplicity = models.MANY
	__Link__000002_Buttons.TargetMultiplicityOffsetX = 0.000000
	__Link__000002_Buttons.TargetMultiplicityOffsetY = 0.000000
	__Link__000002_Buttons.SourceMultiplicity = models.MANY
	__Link__000002_Buttons.SourceMultiplicityOffsetX = 0.000000
	__Link__000002_Buttons.SourceMultiplicityOffsetY = 0.000000
	__Link__000002_Buttons.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_Buttons.StartRatio = 0.238379
	__Link__000002_Buttons.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Buttons.EndRatio = 0.500000
	__Link__000002_Buttons.CornerOffsetRatio = 1.052242

	__Link__000003_SVGIcon.Name = `SVGIcon`

	//gong:ident [ref_models.Button.SVGIcon] comment added to overcome the problem with the comment map association
	__Link__000003_SVGIcon.Identifier = `ref_models.Button.SVGIcon`

	//gong:ident [ref_models.SVGIcon] comment added to overcome the problem with the comment map association
	__Link__000003_SVGIcon.Fieldtypename = `ref_models.SVGIcon`
	__Link__000003_SVGIcon.FieldOffsetX = 0.000000
	__Link__000003_SVGIcon.FieldOffsetY = 0.000000
	__Link__000003_SVGIcon.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_SVGIcon.TargetMultiplicityOffsetX = 0.000000
	__Link__000003_SVGIcon.TargetMultiplicityOffsetY = 0.000000
	__Link__000003_SVGIcon.SourceMultiplicity = models.MANY
	__Link__000003_SVGIcon.SourceMultiplicityOffsetX = 0.000000
	__Link__000003_SVGIcon.SourceMultiplicityOffsetY = 0.000000
	__Link__000003_SVGIcon.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_SVGIcon.StartRatio = 0.500000
	__Link__000003_SVGIcon.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_SVGIcon.EndRatio = 0.500000
	__Link__000003_SVGIcon.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_Tree.X = 43.000000
	__Position__000000_Pos_Default_Tree.Y = 23.000000
	__Position__000000_Pos_Default_Tree.Name = `Pos-Default-Tree`

	__Position__000001_Pos_Default_Node.X = 441.000000
	__Position__000001_Pos_Default_Node.Y = 6.000000
	__Position__000001_Pos_Default_Node.Name = `Pos-Default-Node`

	__Position__000002_Pos_Default_Button.X = 684.000000
	__Position__000002_Pos_Default_Button.Y = 180.000000
	__Position__000002_Pos_Default_Button.Name = `Pos-Default-Button`

	__Position__000003_Pos_Default_SVGIcon.X = 1063.000000
	__Position__000003_Pos_Default_SVGIcon.Y = 180.000000
	__Position__000003_Pos_Default_SVGIcon.Name = `Pos-Default-SVGIcon`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node.X = 447.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node.Y = 96.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node.Name = `Verticle in class diagram Default in middle between Default-Tree and Default-Node`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node.X = 446.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node.Y = 165.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node.Name = `Verticle in class diagram Default in middle between Default-Node and Default-Node`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button.X = 426.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button.Y = 135.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button.Name = `Verticle in class diagram Default in middle between Default-Node and Default-Button`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Button_and_Default_SVGIcon.X = 820.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Button_and_Default_SVGIcon.Y = 90.500000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Button_and_Default_SVGIcon.Name = `Verticle in class diagram Default in middle between Default-Button and Default-SVGIcon`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Tree)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Node)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Button)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_SVGIcon)
	// setup of Field instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Tree.Position = __Position__000000_Pos_Default_Tree
	__GongStructShape__000000_Default_Tree.Fields = append(__GongStructShape__000000_Default_Tree.Fields, __Field__000000_Name)
	__GongStructShape__000000_Default_Tree.Links = append(__GongStructShape__000000_Default_Tree.Links, __Link__000000_RootNodes)
	__GongStructShape__000001_Default_Node.Position = __Position__000001_Pos_Default_Node
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000001_Name)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000002_IsExpanded)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000003_HasCheckboxButton)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000004_IsChecked)
	__GongStructShape__000001_Default_Node.Links = append(__GongStructShape__000001_Default_Node.Links, __Link__000001_Children)
	__GongStructShape__000001_Default_Node.Links = append(__GongStructShape__000001_Default_Node.Links, __Link__000002_Buttons)
	__GongStructShape__000002_Default_Button.Position = __Position__000002_Pos_Default_Button
	__GongStructShape__000002_Default_Button.Fields = append(__GongStructShape__000002_Default_Button.Fields, __Field__000005_Name)
	__GongStructShape__000002_Default_Button.Fields = append(__GongStructShape__000002_Default_Button.Fields, __Field__000006_Icon)
	__GongStructShape__000002_Default_Button.Links = append(__GongStructShape__000002_Default_Button.Links, __Link__000003_SVGIcon)
	__GongStructShape__000003_Default_SVGIcon.Position = __Position__000003_Pos_Default_SVGIcon
	__GongStructShape__000003_Default_SVGIcon.Fields = append(__GongStructShape__000003_Default_SVGIcon.Fields, __Field__000007_Name)
	__GongStructShape__000003_Default_SVGIcon.Fields = append(__GongStructShape__000003_Default_SVGIcon.Fields, __Field__000008_SVG)
	// setup of Link instances pointers
	__Link__000000_RootNodes.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node
	__Link__000001_Children.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node
	__Link__000002_Buttons.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button
	__Link__000003_SVGIcon.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Button_and_Default_SVGIcon
	// setup of Position instances pointers
	// setup of Vertice instances pointers
}
