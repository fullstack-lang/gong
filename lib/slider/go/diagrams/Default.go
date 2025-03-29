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
	__Field__000001_Name := (&models.Field{}).Stage(stage)
	__Field__000002_IsFloat64 := (&models.Field{}).Stage(stage)
	__Field__000003_IsInt := (&models.Field{}).Stage(stage)
	__Field__000004_MinInt := (&models.Field{}).Stage(stage)
	__Field__000005_MaxInt := (&models.Field{}).Stage(stage)
	__Field__000006_StepInt := (&models.Field{}).Stage(stage)
	__Field__000007_ValueInt := (&models.Field{}).Stage(stage)
	__Field__000008_MinFloat64 := (&models.Field{}).Stage(stage)
	__Field__000009_MaxFloat64 := (&models.Field{}).Stage(stage)
	__Field__000010_StepFloat64 := (&models.Field{}).Stage(stage)
	__Field__000011_ValueFloat64 := (&models.Field{}).Stage(stage)
	__Field__000012_Name := (&models.Field{}).Stage(stage)
	__Field__000013_Percentage := (&models.Field{}).Stage(stage)
	__Field__000014_Name := (&models.Field{}).Stage(stage)
	__Field__000015_ValueBool := (&models.Field{}).Stage(stage)
	__Field__000016_LabelForTrue := (&models.Field{}).Stage(stage)
	__Field__000017_LabelForFalse := (&models.Field{}).Stage(stage)

	__GongStructShape__000000_Default_Layout := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Slider := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Group := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_Checkbox := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_Groups := (&models.Link{}).Stage(stage)
	__Link__000001_Sliders := (&models.Link{}).Stage(stage)
	__Link__000002_Checkboxes := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_Default_Layout := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_Slider := (&models.Position{}).Stage(stage)
	__Position__000002_Pos_Default_Group := (&models.Position{}).Stage(stage)
	__Position__000003_Pos_Default_Checkbox := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Layout_and_Default_Group := (&models.Vertice{}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Slider := (&models.Vertice{}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Checkbox := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Name.Name = `Name`

	//gong:ident [ref_models.Layout.Name] comment added to overcome the problem with the comment map association
	__Field__000000_Name.Identifier = `ref_models.Layout.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Layout`
	__Field__000000_Name.Fieldtypename = `string`

	__Field__000001_Name.Name = `Name`

	//gong:ident [ref_models.Slider.Name] comment added to overcome the problem with the comment map association
	__Field__000001_Name.Identifier = `ref_models.Slider.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Slider`
	__Field__000001_Name.Fieldtypename = `string`

	__Field__000002_IsFloat64.Name = `IsFloat64`

	//gong:ident [ref_models.Slider.IsFloat64] comment added to overcome the problem with the comment map association
	__Field__000002_IsFloat64.Identifier = `ref_models.Slider.IsFloat64`
	__Field__000002_IsFloat64.FieldTypeAsString = ``
	__Field__000002_IsFloat64.Structname = `Slider`
	__Field__000002_IsFloat64.Fieldtypename = `bool`

	__Field__000003_IsInt.Name = `IsInt`

	//gong:ident [ref_models.Slider.IsInt] comment added to overcome the problem with the comment map association
	__Field__000003_IsInt.Identifier = `ref_models.Slider.IsInt`
	__Field__000003_IsInt.FieldTypeAsString = ``
	__Field__000003_IsInt.Structname = `Slider`
	__Field__000003_IsInt.Fieldtypename = `bool`

	__Field__000004_MinInt.Name = `MinInt`

	//gong:ident [ref_models.Slider.MinInt] comment added to overcome the problem with the comment map association
	__Field__000004_MinInt.Identifier = `ref_models.Slider.MinInt`
	__Field__000004_MinInt.FieldTypeAsString = ``
	__Field__000004_MinInt.Structname = `Slider`
	__Field__000004_MinInt.Fieldtypename = `int`

	__Field__000005_MaxInt.Name = `MaxInt`

	//gong:ident [ref_models.Slider.MaxInt] comment added to overcome the problem with the comment map association
	__Field__000005_MaxInt.Identifier = `ref_models.Slider.MaxInt`
	__Field__000005_MaxInt.FieldTypeAsString = ``
	__Field__000005_MaxInt.Structname = `Slider`
	__Field__000005_MaxInt.Fieldtypename = `int`

	__Field__000006_StepInt.Name = `StepInt`

	//gong:ident [ref_models.Slider.StepInt] comment added to overcome the problem with the comment map association
	__Field__000006_StepInt.Identifier = `ref_models.Slider.StepInt`
	__Field__000006_StepInt.FieldTypeAsString = ``
	__Field__000006_StepInt.Structname = `Slider`
	__Field__000006_StepInt.Fieldtypename = `int`

	__Field__000007_ValueInt.Name = `ValueInt`

	//gong:ident [ref_models.Slider.ValueInt] comment added to overcome the problem with the comment map association
	__Field__000007_ValueInt.Identifier = `ref_models.Slider.ValueInt`
	__Field__000007_ValueInt.FieldTypeAsString = ``
	__Field__000007_ValueInt.Structname = `Slider`
	__Field__000007_ValueInt.Fieldtypename = `int`

	__Field__000008_MinFloat64.Name = `MinFloat64`

	//gong:ident [ref_models.Slider.MinFloat64] comment added to overcome the problem with the comment map association
	__Field__000008_MinFloat64.Identifier = `ref_models.Slider.MinFloat64`
	__Field__000008_MinFloat64.FieldTypeAsString = ``
	__Field__000008_MinFloat64.Structname = `Slider`
	__Field__000008_MinFloat64.Fieldtypename = `float64`

	__Field__000009_MaxFloat64.Name = `MaxFloat64`

	//gong:ident [ref_models.Slider.MaxFloat64] comment added to overcome the problem with the comment map association
	__Field__000009_MaxFloat64.Identifier = `ref_models.Slider.MaxFloat64`
	__Field__000009_MaxFloat64.FieldTypeAsString = ``
	__Field__000009_MaxFloat64.Structname = `Slider`
	__Field__000009_MaxFloat64.Fieldtypename = `float64`

	__Field__000010_StepFloat64.Name = `StepFloat64`

	//gong:ident [ref_models.Slider.StepFloat64] comment added to overcome the problem with the comment map association
	__Field__000010_StepFloat64.Identifier = `ref_models.Slider.StepFloat64`
	__Field__000010_StepFloat64.FieldTypeAsString = ``
	__Field__000010_StepFloat64.Structname = `Slider`
	__Field__000010_StepFloat64.Fieldtypename = `float64`

	__Field__000011_ValueFloat64.Name = `ValueFloat64`

	//gong:ident [ref_models.Slider.ValueFloat64] comment added to overcome the problem with the comment map association
	__Field__000011_ValueFloat64.Identifier = `ref_models.Slider.ValueFloat64`
	__Field__000011_ValueFloat64.FieldTypeAsString = ``
	__Field__000011_ValueFloat64.Structname = `Slider`
	__Field__000011_ValueFloat64.Fieldtypename = `float64`

	__Field__000012_Name.Name = `Name`

	//gong:ident [ref_models.Group.Name] comment added to overcome the problem with the comment map association
	__Field__000012_Name.Identifier = `ref_models.Group.Name`
	__Field__000012_Name.FieldTypeAsString = ``
	__Field__000012_Name.Structname = `Group`
	__Field__000012_Name.Fieldtypename = `string`

	__Field__000013_Percentage.Name = `Percentage`

	//gong:ident [ref_models.Group.Percentage] comment added to overcome the problem with the comment map association
	__Field__000013_Percentage.Identifier = `ref_models.Group.Percentage`
	__Field__000013_Percentage.FieldTypeAsString = ``
	__Field__000013_Percentage.Structname = `Group`
	__Field__000013_Percentage.Fieldtypename = `float64`

	__Field__000014_Name.Name = `Name`

	//gong:ident [ref_models.Checkbox.Name] comment added to overcome the problem with the comment map association
	__Field__000014_Name.Identifier = `ref_models.Checkbox.Name`
	__Field__000014_Name.FieldTypeAsString = ``
	__Field__000014_Name.Structname = `Checkbox`
	__Field__000014_Name.Fieldtypename = `string`

	__Field__000015_ValueBool.Name = `ValueBool`

	//gong:ident [ref_models.Checkbox.ValueBool] comment added to overcome the problem with the comment map association
	__Field__000015_ValueBool.Identifier = `ref_models.Checkbox.ValueBool`
	__Field__000015_ValueBool.FieldTypeAsString = ``
	__Field__000015_ValueBool.Structname = `Checkbox`
	__Field__000015_ValueBool.Fieldtypename = `bool`

	__Field__000016_LabelForTrue.Name = `LabelForTrue`

	//gong:ident [ref_models.Checkbox.LabelForTrue] comment added to overcome the problem with the comment map association
	__Field__000016_LabelForTrue.Identifier = `ref_models.Checkbox.LabelForTrue`
	__Field__000016_LabelForTrue.FieldTypeAsString = ``
	__Field__000016_LabelForTrue.Structname = `Checkbox`
	__Field__000016_LabelForTrue.Fieldtypename = `string`

	__Field__000017_LabelForFalse.Name = `LabelForFalse`

	//gong:ident [ref_models.Checkbox.LabelForFalse] comment added to overcome the problem with the comment map association
	__Field__000017_LabelForFalse.Identifier = `ref_models.Checkbox.LabelForFalse`
	__Field__000017_LabelForFalse.FieldTypeAsString = ``
	__Field__000017_LabelForFalse.Structname = `Checkbox`
	__Field__000017_LabelForFalse.Fieldtypename = `string`

	__GongStructShape__000000_Default_Layout.Name = `Default-Layout`

	//gong:ident [ref_models.Layout] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Layout.Identifier = `ref_models.Layout`
	__GongStructShape__000000_Default_Layout.ShowNbInstances = false
	__GongStructShape__000000_Default_Layout.NbInstances = 0
	__GongStructShape__000000_Default_Layout.Width = 240.000000
	__GongStructShape__000000_Default_Layout.Height = 78.000000
	__GongStructShape__000000_Default_Layout.IsSelected = false

	__GongStructShape__000001_Default_Slider.Name = `Default-Slider`

	//gong:ident [ref_models.Slider] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Slider.Identifier = `ref_models.Slider`
	__GongStructShape__000001_Default_Slider.ShowNbInstances = false
	__GongStructShape__000001_Default_Slider.NbInstances = 0
	__GongStructShape__000001_Default_Slider.Width = 240.000000
	__GongStructShape__000001_Default_Slider.Height = 228.000000
	__GongStructShape__000001_Default_Slider.IsSelected = false

	__GongStructShape__000002_Default_Group.Name = `Default-Group`

	//gong:ident [ref_models.Group] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Group.Identifier = `ref_models.Group`
	__GongStructShape__000002_Default_Group.ShowNbInstances = false
	__GongStructShape__000002_Default_Group.NbInstances = 0
	__GongStructShape__000002_Default_Group.Width = 240.000000
	__GongStructShape__000002_Default_Group.Height = 93.000000
	__GongStructShape__000002_Default_Group.IsSelected = false

	__GongStructShape__000003_Default_Checkbox.Name = `Default-Checkbox`

	//gong:ident [ref_models.Checkbox] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_Checkbox.Identifier = `ref_models.Checkbox`
	__GongStructShape__000003_Default_Checkbox.ShowNbInstances = false
	__GongStructShape__000003_Default_Checkbox.NbInstances = 0
	__GongStructShape__000003_Default_Checkbox.Width = 240.000000
	__GongStructShape__000003_Default_Checkbox.Height = 123.000000
	__GongStructShape__000003_Default_Checkbox.IsSelected = false

	__Link__000000_Groups.Name = `Groups`

	//gong:ident [ref_models.Layout.Groups] comment added to overcome the problem with the comment map association
	__Link__000000_Groups.Identifier = `ref_models.Layout.Groups`

	//gong:ident [ref_models.Group] comment added to overcome the problem with the comment map association
	__Link__000000_Groups.Fieldtypename = `ref_models.Group`
	__Link__000000_Groups.FieldOffsetX = 0.000000
	__Link__000000_Groups.FieldOffsetY = 0.000000
	__Link__000000_Groups.TargetMultiplicity = models.MANY
	__Link__000000_Groups.TargetMultiplicityOffsetX = 0.000000
	__Link__000000_Groups.TargetMultiplicityOffsetY = 0.000000
	__Link__000000_Groups.SourceMultiplicity = models.MANY
	__Link__000000_Groups.SourceMultiplicityOffsetX = 0.000000
	__Link__000000_Groups.SourceMultiplicityOffsetY = 0.000000
	__Link__000000_Groups.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Groups.StartRatio = 0.871688
	__Link__000000_Groups.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Groups.EndRatio = 0.813354
	__Link__000000_Groups.CornerOffsetRatio = 1.222431

	__Link__000001_Sliders.Name = `Sliders`

	//gong:ident [ref_models.Group.Sliders] comment added to overcome the problem with the comment map association
	__Link__000001_Sliders.Identifier = `ref_models.Group.Sliders`

	//gong:ident [ref_models.Slider] comment added to overcome the problem with the comment map association
	__Link__000001_Sliders.Fieldtypename = `ref_models.Slider`
	__Link__000001_Sliders.FieldOffsetX = 0.000000
	__Link__000001_Sliders.FieldOffsetY = 0.000000
	__Link__000001_Sliders.TargetMultiplicity = models.MANY
	__Link__000001_Sliders.TargetMultiplicityOffsetX = 0.000000
	__Link__000001_Sliders.TargetMultiplicityOffsetY = 0.000000
	__Link__000001_Sliders.SourceMultiplicity = models.MANY
	__Link__000001_Sliders.SourceMultiplicityOffsetX = 0.000000
	__Link__000001_Sliders.SourceMultiplicityOffsetY = 0.000000
	__Link__000001_Sliders.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Sliders.StartRatio = 0.500000
	__Link__000001_Sliders.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Sliders.EndRatio = 0.500000
	__Link__000001_Sliders.CornerOffsetRatio = 1.380000

	__Link__000002_Checkboxes.Name = `Checkboxes`

	//gong:ident [ref_models.Group.Checkboxes] comment added to overcome the problem with the comment map association
	__Link__000002_Checkboxes.Identifier = `ref_models.Group.Checkboxes`

	//gong:ident [ref_models.Checkbox] comment added to overcome the problem with the comment map association
	__Link__000002_Checkboxes.Fieldtypename = `ref_models.Checkbox`
	__Link__000002_Checkboxes.FieldOffsetX = 0.000000
	__Link__000002_Checkboxes.FieldOffsetY = 0.000000
	__Link__000002_Checkboxes.TargetMultiplicity = models.MANY
	__Link__000002_Checkboxes.TargetMultiplicityOffsetX = 0.000000
	__Link__000002_Checkboxes.TargetMultiplicityOffsetY = 0.000000
	__Link__000002_Checkboxes.SourceMultiplicity = models.MANY
	__Link__000002_Checkboxes.SourceMultiplicityOffsetX = 0.000000
	__Link__000002_Checkboxes.SourceMultiplicityOffsetY = 0.000000
	__Link__000002_Checkboxes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Checkboxes.StartRatio = 0.500000
	__Link__000002_Checkboxes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Checkboxes.EndRatio = 0.500000
	__Link__000002_Checkboxes.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_Layout.X = 48.000000
	__Position__000000_Pos_Default_Layout.Y = 44.000000
	__Position__000000_Pos_Default_Layout.Name = `Pos-Default-Layout`

	__Position__000001_Pos_Default_Slider.X = 867.000000
	__Position__000001_Pos_Default_Slider.Y = 24.000000
	__Position__000001_Pos_Default_Slider.Name = `Pos-Default-Slider`

	__Position__000002_Pos_Default_Group.X = 62.000000
	__Position__000002_Pos_Default_Group.Y = 228.000000
	__Position__000002_Pos_Default_Group.Name = `Pos-Default-Group`

	__Position__000003_Pos_Default_Checkbox.X = 866.000000
	__Position__000003_Pos_Default_Checkbox.Y = 300.000000
	__Position__000003_Pos_Default_Checkbox.Name = `Pos-Default-Checkbox`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Layout_and_Default_Group.X = 455.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Layout_and_Default_Group.Y = 127.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Layout_and_Default_Group.Name = `Verticle in class diagram Default in middle between Default-Layout and Default-Group`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Slider.X = 425.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Slider.Y = 139.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Slider.Name = `Verticle in class diagram Default in middle between Default-Group and Default-Slider`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Checkbox.X = 434.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Checkbox.Y = 117.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Checkbox.Name = `Verticle in class diagram Default in middle between Default-Group and Default-Checkbox`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Layout)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Slider)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Group)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Checkbox)
	// setup of Field instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Layout.Position = __Position__000000_Pos_Default_Layout
	__GongStructShape__000000_Default_Layout.Fields = append(__GongStructShape__000000_Default_Layout.Fields, __Field__000000_Name)
	__GongStructShape__000000_Default_Layout.Links = append(__GongStructShape__000000_Default_Layout.Links, __Link__000000_Groups)
	__GongStructShape__000001_Default_Slider.Position = __Position__000001_Pos_Default_Slider
	__GongStructShape__000001_Default_Slider.Fields = append(__GongStructShape__000001_Default_Slider.Fields, __Field__000001_Name)
	__GongStructShape__000001_Default_Slider.Fields = append(__GongStructShape__000001_Default_Slider.Fields, __Field__000002_IsFloat64)
	__GongStructShape__000001_Default_Slider.Fields = append(__GongStructShape__000001_Default_Slider.Fields, __Field__000003_IsInt)
	__GongStructShape__000001_Default_Slider.Fields = append(__GongStructShape__000001_Default_Slider.Fields, __Field__000004_MinInt)
	__GongStructShape__000001_Default_Slider.Fields = append(__GongStructShape__000001_Default_Slider.Fields, __Field__000005_MaxInt)
	__GongStructShape__000001_Default_Slider.Fields = append(__GongStructShape__000001_Default_Slider.Fields, __Field__000006_StepInt)
	__GongStructShape__000001_Default_Slider.Fields = append(__GongStructShape__000001_Default_Slider.Fields, __Field__000007_ValueInt)
	__GongStructShape__000001_Default_Slider.Fields = append(__GongStructShape__000001_Default_Slider.Fields, __Field__000008_MinFloat64)
	__GongStructShape__000001_Default_Slider.Fields = append(__GongStructShape__000001_Default_Slider.Fields, __Field__000009_MaxFloat64)
	__GongStructShape__000001_Default_Slider.Fields = append(__GongStructShape__000001_Default_Slider.Fields, __Field__000010_StepFloat64)
	__GongStructShape__000001_Default_Slider.Fields = append(__GongStructShape__000001_Default_Slider.Fields, __Field__000011_ValueFloat64)
	__GongStructShape__000002_Default_Group.Position = __Position__000002_Pos_Default_Group
	__GongStructShape__000002_Default_Group.Fields = append(__GongStructShape__000002_Default_Group.Fields, __Field__000012_Name)
	__GongStructShape__000002_Default_Group.Fields = append(__GongStructShape__000002_Default_Group.Fields, __Field__000013_Percentage)
	__GongStructShape__000002_Default_Group.Links = append(__GongStructShape__000002_Default_Group.Links, __Link__000001_Sliders)
	__GongStructShape__000002_Default_Group.Links = append(__GongStructShape__000002_Default_Group.Links, __Link__000002_Checkboxes)
	__GongStructShape__000003_Default_Checkbox.Position = __Position__000003_Pos_Default_Checkbox
	__GongStructShape__000003_Default_Checkbox.Fields = append(__GongStructShape__000003_Default_Checkbox.Fields, __Field__000014_Name)
	__GongStructShape__000003_Default_Checkbox.Fields = append(__GongStructShape__000003_Default_Checkbox.Fields, __Field__000015_ValueBool)
	__GongStructShape__000003_Default_Checkbox.Fields = append(__GongStructShape__000003_Default_Checkbox.Fields, __Field__000016_LabelForTrue)
	__GongStructShape__000003_Default_Checkbox.Fields = append(__GongStructShape__000003_Default_Checkbox.Fields, __Field__000017_LabelForFalse)
	// setup of Link instances pointers
	__Link__000000_Groups.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Layout_and_Default_Group
	__Link__000001_Sliders.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Slider
	__Link__000002_Checkboxes.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Group_and_Default_Checkbox
	// setup of Position instances pointers
	// setup of Vertice instances pointers
}
