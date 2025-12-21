package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/dsme/project/go/models"
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

	__Classdiagram__00000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__00000000_Default_Product := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__00000001_Default_Project := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__00000002_Default_Task := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__00000000_SubProducts := (&models.LinkShape{}).Stage(stage)
	__LinkShape__00000001_SubTasks := (&models.LinkShape{}).Stage(stage)
	__LinkShape__00000002_RootProducts := (&models.LinkShape{}).Stage(stage)
	__LinkShape__00000003_RootTasks := (&models.LinkShape{}).Stage(stage)
	__LinkShape__00000004_Inputs := (&models.LinkShape{}).Stage(stage)
	__LinkShape__00000005_Outputs := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__00000000_Default.Name = `Default`
	__Classdiagram__00000000_Default.Description = ``
	__Classdiagram__00000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_Default.ShowNbInstances = true
	__Classdiagram__00000000_Default.ShowMultiplicity = true
	__Classdiagram__00000000_Default.ShowLinkNames = true
	__Classdiagram__00000000_Default.IsInRenameMode = false
	__Classdiagram__00000000_Default.IsExpanded = true
	__Classdiagram__00000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_Default.NodeGongStructNodeExpansion = `[false,true,false,true,false,true]`
	__Classdiagram__00000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.Name = `Diagram Package created the 2025-12-18T17:41:37Z`
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.Path = ``
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.GongModelPath = ``
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_Default_Product.Name = `Default-Product`
	__GongStructShape__00000000_Default_Product.X = 518.000000
	__GongStructShape__00000000_Default_Product.Y = 62.000000
	__GongStructShape__00000000_Default_Product.IdentifierMeta = ref_models.Product{}
	__GongStructShape__00000000_Default_Product.Width = 240.000000
	__GongStructShape__00000000_Default_Product.Height = 146.000000
	__GongStructShape__00000000_Default_Product.IsSelected = false

	__GongStructShape__00000001_Default_Project.Name = `Default-Project`
	__GongStructShape__00000001_Default_Project.X = 16.000000
	__GongStructShape__00000001_Default_Project.Y = 73.000000
	__GongStructShape__00000001_Default_Project.IdentifierMeta = ref_models.Project{}
	__GongStructShape__00000001_Default_Project.Width = 240.000000
	__GongStructShape__00000001_Default_Project.Height = 435.000000
	__GongStructShape__00000001_Default_Project.IsSelected = false

	__GongStructShape__00000002_Default_Task.Name = `Default-Task`
	__GongStructShape__00000002_Default_Task.X = 521.000000
	__GongStructShape__00000002_Default_Task.Y = 329.000000
	__GongStructShape__00000002_Default_Task.IdentifierMeta = ref_models.Task{}
	__GongStructShape__00000002_Default_Task.Width = 240.000000
	__GongStructShape__00000002_Default_Task.Height = 176.000000
	__GongStructShape__00000002_Default_Task.IsSelected = false

	__LinkShape__00000000_SubProducts.Name = `SubProducts`
	__LinkShape__00000000_SubProducts.IdentifierMeta = ref_models.Product{}.SubProducts
	__LinkShape__00000000_SubProducts.FieldTypeIdentifierMeta = ref_models.Product{}
	__LinkShape__00000000_SubProducts.FieldOffsetX = 0.000000
	__LinkShape__00000000_SubProducts.FieldOffsetY = 0.000000
	__LinkShape__00000000_SubProducts.TargetMultiplicity = models.MANY
	__LinkShape__00000000_SubProducts.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_SubProducts.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_SubProducts.SourceMultiplicity = models.MANY
	__LinkShape__00000000_SubProducts.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_SubProducts.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_SubProducts.X = 371.000000
	__LinkShape__00000000_SubProducts.Y = 51.500000
	__LinkShape__00000000_SubProducts.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_SubProducts.StartRatio = 0.216650
	__LinkShape__00000000_SubProducts.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_SubProducts.EndRatio = 0.730348
	__LinkShape__00000000_SubProducts.CornerOffsetRatio = 1.656893

	__LinkShape__00000001_SubTasks.Name = `SubTasks`
	__LinkShape__00000001_SubTasks.IdentifierMeta = ref_models.Task{}.SubTasks
	__LinkShape__00000001_SubTasks.FieldTypeIdentifierMeta = ref_models.Task{}
	__LinkShape__00000001_SubTasks.FieldOffsetX = 0.000000
	__LinkShape__00000001_SubTasks.FieldOffsetY = 0.000000
	__LinkShape__00000001_SubTasks.TargetMultiplicity = models.MANY
	__LinkShape__00000001_SubTasks.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_SubTasks.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_SubTasks.SourceMultiplicity = models.MANY
	__LinkShape__00000001_SubTasks.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_SubTasks.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_SubTasks.X = 884.000000
	__LinkShape__00000001_SubTasks.Y = 293.500000
	__LinkShape__00000001_SubTasks.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_SubTasks.StartRatio = 0.200720
	__LinkShape__00000001_SubTasks.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_SubTasks.EndRatio = 0.701738
	__LinkShape__00000001_SubTasks.CornerOffsetRatio = 1.590112

	__LinkShape__00000002_RootProducts.Name = `RootProducts`
	__LinkShape__00000002_RootProducts.IdentifierMeta = ref_models.Project{}.RootProducts
	__LinkShape__00000002_RootProducts.FieldTypeIdentifierMeta = ref_models.Product{}
	__LinkShape__00000002_RootProducts.FieldOffsetX = 0.000000
	__LinkShape__00000002_RootProducts.FieldOffsetY = 0.000000
	__LinkShape__00000002_RootProducts.TargetMultiplicity = models.MANY
	__LinkShape__00000002_RootProducts.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_RootProducts.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_RootProducts.SourceMultiplicity = models.MANY
	__LinkShape__00000002_RootProducts.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_RootProducts.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_RootProducts.X = 634.500000
	__LinkShape__00000002_RootProducts.Y = 170.000000
	__LinkShape__00000002_RootProducts.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_RootProducts.StartRatio = 0.159784
	__LinkShape__00000002_RootProducts.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_RootProducts.EndRatio = 0.500000
	__LinkShape__00000002_RootProducts.CornerOffsetRatio = 1.380000

	__LinkShape__00000003_RootTasks.Name = `RootTasks`
	__LinkShape__00000003_RootTasks.IdentifierMeta = ref_models.Project{}.RootTasks
	__LinkShape__00000003_RootTasks.FieldTypeIdentifierMeta = ref_models.Task{}
	__LinkShape__00000003_RootTasks.FieldOffsetX = 0.000000
	__LinkShape__00000003_RootTasks.FieldOffsetY = 0.000000
	__LinkShape__00000003_RootTasks.TargetMultiplicity = models.MANY
	__LinkShape__00000003_RootTasks.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_RootTasks.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_RootTasks.SourceMultiplicity = models.MANY
	__LinkShape__00000003_RootTasks.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_RootTasks.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_RootTasks.X = 637.500000
	__LinkShape__00000003_RootTasks.Y = 255.500000
	__LinkShape__00000003_RootTasks.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_RootTasks.StartRatio = 0.870128
	__LinkShape__00000003_RootTasks.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_RootTasks.EndRatio = 0.701738
	__LinkShape__00000003_RootTasks.CornerOffsetRatio = 1.380000

	__LinkShape__00000004_Inputs.Name = `Inputs`
	__LinkShape__00000004_Inputs.IdentifierMeta = ref_models.Task{}.Inputs
	__LinkShape__00000004_Inputs.FieldTypeIdentifierMeta = ref_models.Product{}
	__LinkShape__00000004_Inputs.FieldOffsetX = 0.000000
	__LinkShape__00000004_Inputs.FieldOffsetY = 0.000000
	__LinkShape__00000004_Inputs.TargetMultiplicity = models.MANY
	__LinkShape__00000004_Inputs.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_Inputs.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_Inputs.SourceMultiplicity = models.MANY
	__LinkShape__00000004_Inputs.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_Inputs.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_Inputs.X = 879.500000
	__LinkShape__00000004_Inputs.Y = 283.500000
	__LinkShape__00000004_Inputs.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000004_Inputs.StartRatio = 0.769279
	__LinkShape__00000004_Inputs.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000004_Inputs.EndRatio = 0.840112
	__LinkShape__00000004_Inputs.CornerOffsetRatio = -0.309626

	__LinkShape__00000005_Outputs.Name = `Outputs`
	__LinkShape__00000005_Outputs.IdentifierMeta = ref_models.Task{}.Outputs
	__LinkShape__00000005_Outputs.FieldTypeIdentifierMeta = ref_models.Product{}
	__LinkShape__00000005_Outputs.FieldOffsetX = 0.000000
	__LinkShape__00000005_Outputs.FieldOffsetY = 0.000000
	__LinkShape__00000005_Outputs.TargetMultiplicity = models.MANY
	__LinkShape__00000005_Outputs.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_Outputs.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_Outputs.SourceMultiplicity = models.MANY
	__LinkShape__00000005_Outputs.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_Outputs.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_Outputs.X = 879.500000
	__LinkShape__00000005_Outputs.Y = 283.500000
	__LinkShape__00000005_Outputs.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000005_Outputs.StartRatio = 0.177612
	__LinkShape__00000005_Outputs.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000005_Outputs.EndRatio = 0.215112
	__LinkShape__00000005_Outputs.CornerOffsetRatio = -0.213035

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__00000000_Default.GongStructShapes = append(__Classdiagram__00000000_Default.GongStructShapes, __GongStructShape__00000000_Default_Product)
	__Classdiagram__00000000_Default.GongStructShapes = append(__Classdiagram__00000000_Default.GongStructShapes, __GongStructShape__00000001_Default_Project)
	__Classdiagram__00000000_Default.GongStructShapes = append(__Classdiagram__00000000_Default.GongStructShapes, __GongStructShape__00000002_Default_Task)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.Classdiagrams = append(__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.Classdiagrams, __Classdiagram__00000000_Default)
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.SelectedClassdiagram = __Classdiagram__00000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__00000000_Default_Product.LinkShapes = append(__GongStructShape__00000000_Default_Product.LinkShapes, __LinkShape__00000000_SubProducts)
	__GongStructShape__00000001_Default_Project.LinkShapes = append(__GongStructShape__00000001_Default_Project.LinkShapes, __LinkShape__00000002_RootProducts)
	__GongStructShape__00000001_Default_Project.LinkShapes = append(__GongStructShape__00000001_Default_Project.LinkShapes, __LinkShape__00000003_RootTasks)
	__GongStructShape__00000002_Default_Task.LinkShapes = append(__GongStructShape__00000002_Default_Task.LinkShapes, __LinkShape__00000001_SubTasks)
	__GongStructShape__00000002_Default_Task.LinkShapes = append(__GongStructShape__00000002_Default_Task.LinkShapes, __LinkShape__00000004_Inputs)
	__GongStructShape__00000002_Default_Task.LinkShapes = append(__GongStructShape__00000002_Default_Task.LinkShapes, __LinkShape__00000005_Outputs)
	// setup of LinkShape instances pointers
}

