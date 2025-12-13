package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/slider/go/models"
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

	const __write__local_time = "2025-06-18 06:48:53.429151 CEST"
	const __write__utc_time__ = "2025-06-18 04:48:53.429151 UTC"

	const __commitId__ = "0000000050"

	// Declaration of instances to stage

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_IsFloat64 := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_IsInt := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000003_MinInt := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000004_MaxInt := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000005_StepInt := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000006_ValueInt := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000007_MinFloat64 := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000008_MaxFloat64 := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000009_StepFloat64 := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000010_ValueFloat64 := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000011_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000012_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000013_ValueBool := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000014_LabelForTrue := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000015_LabelForFalse := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000016_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000017_Percentage := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_47_38Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Slider := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Layout := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Group := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_Checkbox := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Groups := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_Sliders := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_Checkboxes := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.Slider{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Slider`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_IsFloat64.Name = `IsFloat64`
	__AttributeShape__000001_IsFloat64.IdentifierMeta = ref_models.Slider{}.IsFloat64
	__AttributeShape__000001_IsFloat64.FieldTypeAsString = ``
	__AttributeShape__000001_IsFloat64.Structname = `Slider`
	__AttributeShape__000001_IsFloat64.Fieldtypename = `bool`

	__AttributeShape__000002_IsInt.Name = `IsInt`
	__AttributeShape__000002_IsInt.IdentifierMeta = ref_models.Slider{}.IsInt
	__AttributeShape__000002_IsInt.FieldTypeAsString = ``
	__AttributeShape__000002_IsInt.Structname = `Slider`
	__AttributeShape__000002_IsInt.Fieldtypename = `bool`

	__AttributeShape__000003_MinInt.Name = `MinInt`
	__AttributeShape__000003_MinInt.IdentifierMeta = ref_models.Slider{}.MinInt
	__AttributeShape__000003_MinInt.FieldTypeAsString = ``
	__AttributeShape__000003_MinInt.Structname = `Slider`
	__AttributeShape__000003_MinInt.Fieldtypename = `int`

	__AttributeShape__000004_MaxInt.Name = `MaxInt`
	__AttributeShape__000004_MaxInt.IdentifierMeta = ref_models.Slider{}.MaxInt
	__AttributeShape__000004_MaxInt.FieldTypeAsString = ``
	__AttributeShape__000004_MaxInt.Structname = `Slider`
	__AttributeShape__000004_MaxInt.Fieldtypename = `int`

	__AttributeShape__000005_StepInt.Name = `StepInt`
	__AttributeShape__000005_StepInt.IdentifierMeta = ref_models.Slider{}.StepInt
	__AttributeShape__000005_StepInt.FieldTypeAsString = ``
	__AttributeShape__000005_StepInt.Structname = `Slider`
	__AttributeShape__000005_StepInt.Fieldtypename = `int`

	__AttributeShape__000006_ValueInt.Name = `ValueInt`
	__AttributeShape__000006_ValueInt.IdentifierMeta = ref_models.Slider{}.ValueInt
	__AttributeShape__000006_ValueInt.FieldTypeAsString = ``
	__AttributeShape__000006_ValueInt.Structname = `Slider`
	__AttributeShape__000006_ValueInt.Fieldtypename = `int`

	__AttributeShape__000007_MinFloat64.Name = `MinFloat64`
	__AttributeShape__000007_MinFloat64.IdentifierMeta = ref_models.Slider{}.MinFloat64
	__AttributeShape__000007_MinFloat64.FieldTypeAsString = ``
	__AttributeShape__000007_MinFloat64.Structname = `Slider`
	__AttributeShape__000007_MinFloat64.Fieldtypename = `float64`

	__AttributeShape__000008_MaxFloat64.Name = `MaxFloat64`
	__AttributeShape__000008_MaxFloat64.IdentifierMeta = ref_models.Slider{}.MaxFloat64
	__AttributeShape__000008_MaxFloat64.FieldTypeAsString = ``
	__AttributeShape__000008_MaxFloat64.Structname = `Slider`
	__AttributeShape__000008_MaxFloat64.Fieldtypename = `float64`

	__AttributeShape__000009_StepFloat64.Name = `StepFloat64`
	__AttributeShape__000009_StepFloat64.IdentifierMeta = ref_models.Slider{}.StepFloat64
	__AttributeShape__000009_StepFloat64.FieldTypeAsString = ``
	__AttributeShape__000009_StepFloat64.Structname = `Slider`
	__AttributeShape__000009_StepFloat64.Fieldtypename = `float64`

	__AttributeShape__000010_ValueFloat64.Name = `ValueFloat64`
	__AttributeShape__000010_ValueFloat64.IdentifierMeta = ref_models.Slider{}.ValueFloat64
	__AttributeShape__000010_ValueFloat64.FieldTypeAsString = ``
	__AttributeShape__000010_ValueFloat64.Structname = `Slider`
	__AttributeShape__000010_ValueFloat64.Fieldtypename = `float64`

	__AttributeShape__000011_Name.Name = `Name`
	__AttributeShape__000011_Name.IdentifierMeta = ref_models.Layout{}.Name
	__AttributeShape__000011_Name.FieldTypeAsString = ``
	__AttributeShape__000011_Name.Structname = `Layout`
	__AttributeShape__000011_Name.Fieldtypename = `string`

	__AttributeShape__000012_Name.Name = `Name`
	__AttributeShape__000012_Name.IdentifierMeta = ref_models.Checkbox{}.Name
	__AttributeShape__000012_Name.FieldTypeAsString = ``
	__AttributeShape__000012_Name.Structname = `Checkbox`
	__AttributeShape__000012_Name.Fieldtypename = `string`

	__AttributeShape__000013_ValueBool.Name = `ValueBool`
	__AttributeShape__000013_ValueBool.IdentifierMeta = ref_models.Checkbox{}.ValueBool
	__AttributeShape__000013_ValueBool.FieldTypeAsString = ``
	__AttributeShape__000013_ValueBool.Structname = `Checkbox`
	__AttributeShape__000013_ValueBool.Fieldtypename = `bool`

	__AttributeShape__000014_LabelForTrue.Name = `LabelForTrue`
	__AttributeShape__000014_LabelForTrue.IdentifierMeta = ref_models.Checkbox{}.LabelForTrue
	__AttributeShape__000014_LabelForTrue.FieldTypeAsString = ``
	__AttributeShape__000014_LabelForTrue.Structname = `Checkbox`
	__AttributeShape__000014_LabelForTrue.Fieldtypename = `string`

	__AttributeShape__000015_LabelForFalse.Name = `LabelForFalse`
	__AttributeShape__000015_LabelForFalse.IdentifierMeta = ref_models.Checkbox{}.LabelForFalse
	__AttributeShape__000015_LabelForFalse.FieldTypeAsString = ``
	__AttributeShape__000015_LabelForFalse.Structname = `Checkbox`
	__AttributeShape__000015_LabelForFalse.Fieldtypename = `string`

	__AttributeShape__000016_Name.Name = `Name`
	__AttributeShape__000016_Name.IdentifierMeta = ref_models.Group{}.Name
	__AttributeShape__000016_Name.FieldTypeAsString = ``
	__AttributeShape__000016_Name.Structname = `Group`
	__AttributeShape__000016_Name.Fieldtypename = `string`

	__AttributeShape__000017_Percentage.Name = `Percentage`
	__AttributeShape__000017_Percentage.IdentifierMeta = ref_models.Group{}.Percentage
	__AttributeShape__000017_Percentage.FieldTypeAsString = ``
	__AttributeShape__000017_Percentage.Structname = `Group`
	__AttributeShape__000017_Percentage.Fieldtypename = `float64`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[true,true,true,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_47_38Z.Name = `Diagram Package created the 2025-06-18T04:47:38Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_47_38Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_47_38Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_47_38Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Slider.Name = `Default-Slider`
	__GongStructShape__000000_Default_Slider.X = 919.000000
	__GongStructShape__000000_Default_Slider.Y = 257.000000
	__GongStructShape__000000_Default_Slider.IdentifierMeta = ref_models.Slider{}

	__GongStructShape__000000_Default_Slider.Width = 240.000000
	__GongStructShape__000000_Default_Slider.Height = 283.000000
	__GongStructShape__000000_Default_Slider.IsSelected = false

	__GongStructShape__000001_Default_Layout.Name = `Default-Layout`
	__GongStructShape__000001_Default_Layout.X = 43.000000
	__GongStructShape__000001_Default_Layout.Y = 99.000000
	__GongStructShape__000001_Default_Layout.IdentifierMeta = ref_models.Layout{}

	__GongStructShape__000001_Default_Layout.Width = 240.000000
	__GongStructShape__000001_Default_Layout.Height = 83.000000
	__GongStructShape__000001_Default_Layout.IsSelected = false

	__GongStructShape__000002_Default_Group.Name = `Default-Group`
	__GongStructShape__000002_Default_Group.X = 411.000000
	__GongStructShape__000002_Default_Group.Y = 97.000000
	__GongStructShape__000002_Default_Group.IdentifierMeta = ref_models.Group{}

	__GongStructShape__000002_Default_Group.Width = 240.000000
	__GongStructShape__000002_Default_Group.Height = 103.000000
	__GongStructShape__000002_Default_Group.IsSelected = false

	__GongStructShape__000003_Default_Checkbox.Name = `Default-Checkbox`
	__GongStructShape__000003_Default_Checkbox.X = 916.000000
	__GongStructShape__000003_Default_Checkbox.Y = 82.000000
	__GongStructShape__000003_Default_Checkbox.IdentifierMeta = ref_models.Checkbox{}

	__GongStructShape__000003_Default_Checkbox.Width = 240.000000
	__GongStructShape__000003_Default_Checkbox.Height = 143.000000
	__GongStructShape__000003_Default_Checkbox.IsSelected = false

	__LinkShape__000000_Groups.Name = `Groups`
	__LinkShape__000000_Groups.IdentifierMeta = ref_models.Layout{}.Groups
	__LinkShape__000000_Groups.FieldTypeIdentifierMeta = ref_models.Group{}
	__LinkShape__000000_Groups.FieldOffsetX = 0.000000
	__LinkShape__000000_Groups.FieldOffsetY = 0.000000
	__LinkShape__000000_Groups.TargetMultiplicity = models.MANY
	__LinkShape__000000_Groups.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Groups.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Groups.SourceMultiplicity = models.MANY
	__LinkShape__000000_Groups.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Groups.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Groups.X = 587.000000
	__LinkShape__000000_Groups.Y = 139.500000
	__LinkShape__000000_Groups.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Groups.StartRatio = 0.500000
	__LinkShape__000000_Groups.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Groups.EndRatio = 0.500000
	__LinkShape__000000_Groups.CornerOffsetRatio = 1.380000

	__LinkShape__000001_Sliders.Name = `Sliders`
	__LinkShape__000001_Sliders.IdentifierMeta = ref_models.Group{}.Sliders
	__LinkShape__000001_Sliders.FieldTypeIdentifierMeta = ref_models.Slider{}
	__LinkShape__000001_Sliders.FieldOffsetX = 0.000000
	__LinkShape__000001_Sliders.FieldOffsetY = 0.000000
	__LinkShape__000001_Sliders.TargetMultiplicity = models.MANY
	__LinkShape__000001_Sliders.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Sliders.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Sliders.SourceMultiplicity = models.MANY
	__LinkShape__000001_Sliders.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Sliders.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Sliders.X = 1003.000000
	__LinkShape__000001_Sliders.Y = 237.500000
	__LinkShape__000001_Sliders.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Sliders.StartRatio = 0.500000
	__LinkShape__000001_Sliders.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Sliders.EndRatio = 0.500000
	__LinkShape__000001_Sliders.CornerOffsetRatio = 1.380000

	__LinkShape__000002_Checkboxes.Name = `Checkboxes`
	__LinkShape__000002_Checkboxes.IdentifierMeta = ref_models.Group{}.Checkboxes
	__LinkShape__000002_Checkboxes.FieldTypeIdentifierMeta = ref_models.Checkbox{}
	__LinkShape__000002_Checkboxes.FieldOffsetX = 0.000000
	__LinkShape__000002_Checkboxes.FieldOffsetY = 0.000000
	__LinkShape__000002_Checkboxes.TargetMultiplicity = models.MANY
	__LinkShape__000002_Checkboxes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Checkboxes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Checkboxes.SourceMultiplicity = models.MANY
	__LinkShape__000002_Checkboxes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Checkboxes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Checkboxes.X = 1023.500000
	__LinkShape__000002_Checkboxes.Y = 141.000000
	__LinkShape__000002_Checkboxes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Checkboxes.StartRatio = 0.500000
	__LinkShape__000002_Checkboxes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Checkboxes.EndRatio = 0.500000
	__LinkShape__000002_Checkboxes.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Slider)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Layout)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Group)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Checkbox)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_47_38Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_47_38Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_47_38Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Slider.AttributeShapes = append(__GongStructShape__000000_Default_Slider.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_Slider.AttributeShapes = append(__GongStructShape__000000_Default_Slider.AttributeShapes, __AttributeShape__000001_IsFloat64)
	__GongStructShape__000000_Default_Slider.AttributeShapes = append(__GongStructShape__000000_Default_Slider.AttributeShapes, __AttributeShape__000002_IsInt)
	__GongStructShape__000000_Default_Slider.AttributeShapes = append(__GongStructShape__000000_Default_Slider.AttributeShapes, __AttributeShape__000003_MinInt)
	__GongStructShape__000000_Default_Slider.AttributeShapes = append(__GongStructShape__000000_Default_Slider.AttributeShapes, __AttributeShape__000004_MaxInt)
	__GongStructShape__000000_Default_Slider.AttributeShapes = append(__GongStructShape__000000_Default_Slider.AttributeShapes, __AttributeShape__000005_StepInt)
	__GongStructShape__000000_Default_Slider.AttributeShapes = append(__GongStructShape__000000_Default_Slider.AttributeShapes, __AttributeShape__000006_ValueInt)
	__GongStructShape__000000_Default_Slider.AttributeShapes = append(__GongStructShape__000000_Default_Slider.AttributeShapes, __AttributeShape__000007_MinFloat64)
	__GongStructShape__000000_Default_Slider.AttributeShapes = append(__GongStructShape__000000_Default_Slider.AttributeShapes, __AttributeShape__000008_MaxFloat64)
	__GongStructShape__000000_Default_Slider.AttributeShapes = append(__GongStructShape__000000_Default_Slider.AttributeShapes, __AttributeShape__000009_StepFloat64)
	__GongStructShape__000000_Default_Slider.AttributeShapes = append(__GongStructShape__000000_Default_Slider.AttributeShapes, __AttributeShape__000010_ValueFloat64)
	__GongStructShape__000001_Default_Layout.AttributeShapes = append(__GongStructShape__000001_Default_Layout.AttributeShapes, __AttributeShape__000011_Name)
	__GongStructShape__000001_Default_Layout.LinkShapes = append(__GongStructShape__000001_Default_Layout.LinkShapes, __LinkShape__000000_Groups)
	__GongStructShape__000002_Default_Group.AttributeShapes = append(__GongStructShape__000002_Default_Group.AttributeShapes, __AttributeShape__000016_Name)
	__GongStructShape__000002_Default_Group.AttributeShapes = append(__GongStructShape__000002_Default_Group.AttributeShapes, __AttributeShape__000017_Percentage)
	__GongStructShape__000002_Default_Group.LinkShapes = append(__GongStructShape__000002_Default_Group.LinkShapes, __LinkShape__000001_Sliders)
	__GongStructShape__000002_Default_Group.LinkShapes = append(__GongStructShape__000002_Default_Group.LinkShapes, __LinkShape__000002_Checkboxes)
	__GongStructShape__000003_Default_Checkbox.AttributeShapes = append(__GongStructShape__000003_Default_Checkbox.AttributeShapes, __AttributeShape__000012_Name)
	__GongStructShape__000003_Default_Checkbox.AttributeShapes = append(__GongStructShape__000003_Default_Checkbox.AttributeShapes, __AttributeShape__000013_ValueBool)
	__GongStructShape__000003_Default_Checkbox.AttributeShapes = append(__GongStructShape__000003_Default_Checkbox.AttributeShapes, __AttributeShape__000014_LabelForTrue)
	__GongStructShape__000003_Default_Checkbox.AttributeShapes = append(__GongStructShape__000003_Default_Checkbox.AttributeShapes, __AttributeShape__000015_LabelForFalse)
	// setup of LinkShape instances pointers
}
