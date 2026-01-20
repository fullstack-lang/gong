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
	// insertion point for declaration of instances to stage

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Abstract model`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2025-12-18T17:41:37Z`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-Product`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-Project`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Default-Task`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `Default-Note`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `SubProducts`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `SubTasks`}).Stage(stage)
	__LinkShape__00000002_ := (&models.LinkShape{Name: `RootProducts`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `RootTasks`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `Inputs`}).Stage(stage)
	__LinkShape__00000005_ := (&models.LinkShape{Name: `Outputs`}).Stage(stage)
	__LinkShape__00000006_ := (&models.LinkShape{Name: `Notes`}).Stage(stage)
	__LinkShape__00000007_ := (&models.LinkShape{Name: `Products`}).Stage(stage)
	__LinkShape__00000008_ := (&models.LinkShape{Name: `Tasks`}).Stage(stage)

	// insertion point for initialization of values

	__Classdiagram__00000000_.Name = `Abstract model`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = false
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[false,true,false,false,false,false,false,false,false]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2025-12-18T17:41:37Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_.Name = `Default-Product`
	__GongStructShape__00000000_.X = 518.000000
	__GongStructShape__00000000_.Y = 62.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.Product{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 210.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-Project`
	__GongStructShape__00000001_.X = 16.000000
	__GongStructShape__00000001_.Y = 73.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.Project{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 701.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Default-Task`
	__GongStructShape__00000002_.X = 521.000000
	__GongStructShape__00000002_.Y = 353.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.Task{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 211.000000
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `Default-Note`
	__GongStructShape__00000003_.X = 529.000000
	__GongStructShape__00000003_.Y = 636.000000
	__GongStructShape__00000003_.IdentifierMeta = ref_models.Note{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 105.000000
	__GongStructShape__00000003_.IsSelected = false

	__LinkShape__00000000_.Name = `SubProducts`
	__LinkShape__00000000_.IdentifierMeta = ref_models.Product{}.SubProducts
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.Product{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.MANY
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 371.000000
	__LinkShape__00000000_.Y = 51.500000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.StartRatio = 0.102427
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.378618
	__LinkShape__00000000_.CornerOffsetRatio = 1.656893

	__LinkShape__00000001_.Name = `SubTasks`
	__LinkShape__00000001_.IdentifierMeta = ref_models.Task{}.SubTasks
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.Task{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.MANY
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 884.000000
	__LinkShape__00000001_.Y = 293.500000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.200720
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.610851
	__LinkShape__00000001_.CornerOffsetRatio = 1.590112

	__LinkShape__00000002_.Name = `RootProducts`
	__LinkShape__00000002_.IdentifierMeta = ref_models.Project{}.RootProducts
	__LinkShape__00000002_.FieldTypeIdentifierMeta = ref_models.Product{}
	__LinkShape__00000002_.FieldOffsetX = 0.000000
	__LinkShape__00000002_.FieldOffsetY = 0.000000
	__LinkShape__00000002_.TargetMultiplicity = models.MANY
	__LinkShape__00000002_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.SourceMultiplicity = models.MANY
	__LinkShape__00000002_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.X = 634.500000
	__LinkShape__00000002_.Y = 170.000000
	__LinkShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.StartRatio = 0.076333
	__LinkShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.EndRatio = 0.278618
	__LinkShape__00000002_.CornerOffsetRatio = 1.380000

	__LinkShape__00000003_.Name = `RootTasks`
	__LinkShape__00000003_.IdentifierMeta = ref_models.Project{}.RootTasks
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.Task{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.MANY
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 637.500000
	__LinkShape__00000003_.Y = 255.500000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.StartRatio = 0.635535
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.EndRatio = 0.784406
	__LinkShape__00000003_.CornerOffsetRatio = 1.380000

	__LinkShape__00000004_.Name = `Inputs`
	__LinkShape__00000004_.IdentifierMeta = ref_models.Task{}.Inputs
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.Product{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.MANY
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 879.500000
	__LinkShape__00000004_.Y = 283.500000
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000004_.StartRatio = 0.769279
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000004_.EndRatio = 0.840112
	__LinkShape__00000004_.CornerOffsetRatio = -0.309626

	__LinkShape__00000005_.Name = `Outputs`
	__LinkShape__00000005_.IdentifierMeta = ref_models.Task{}.Outputs
	__LinkShape__00000005_.FieldTypeIdentifierMeta = ref_models.Product{}
	__LinkShape__00000005_.FieldOffsetX = 0.000000
	__LinkShape__00000005_.FieldOffsetY = 0.000000
	__LinkShape__00000005_.TargetMultiplicity = models.MANY
	__LinkShape__00000005_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.SourceMultiplicity = models.MANY
	__LinkShape__00000005_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.X = 879.500000
	__LinkShape__00000005_.Y = 283.500000
	__LinkShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000005_.StartRatio = 0.177612
	__LinkShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000005_.EndRatio = 0.215112
	__LinkShape__00000005_.CornerOffsetRatio = -0.213035

	__LinkShape__00000006_.Name = `Notes`
	__LinkShape__00000006_.IdentifierMeta = ref_models.Project{}.Notes
	__LinkShape__00000006_.FieldTypeIdentifierMeta = ref_models.Note{}
	__LinkShape__00000006_.FieldOffsetX = 0.000000
	__LinkShape__00000006_.FieldOffsetY = 0.000000
	__LinkShape__00000006_.TargetMultiplicity = models.MANY
	__LinkShape__00000006_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.SourceMultiplicity = models.MANY
	__LinkShape__00000006_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.X = 632.500000
	__LinkShape__00000006_.Y = 705.000000
	__LinkShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.StartRatio = 0.880898
	__LinkShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.EndRatio = 0.500000
	__LinkShape__00000006_.CornerOffsetRatio = 1.380000

	__LinkShape__00000007_.Name = `Products`
	__LinkShape__00000007_.IdentifierMeta = ref_models.Note{}.Products
	__LinkShape__00000007_.FieldTypeIdentifierMeta = ref_models.Product{}
	__LinkShape__00000007_.FieldOffsetX = 0.000000
	__LinkShape__00000007_.FieldOffsetY = 0.000000
	__LinkShape__00000007_.TargetMultiplicity = models.MANY
	__LinkShape__00000007_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.SourceMultiplicity = models.MANY
	__LinkShape__00000007_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.X = 883.500000
	__LinkShape__00000007_.Y = 380.500000
	__LinkShape__00000007_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.StartRatio = 0.700093
	__LinkShape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.EndRatio = 0.702427
	__LinkShape__00000007_.CornerOffsetRatio = 1.805729

	__LinkShape__00000008_.Name = `Tasks`
	__LinkShape__00000008_.IdentifierMeta = ref_models.Note{}.Tasks
	__LinkShape__00000008_.FieldTypeIdentifierMeta = ref_models.Task{}
	__LinkShape__00000008_.FieldOffsetX = 0.000000
	__LinkShape__00000008_.FieldOffsetY = 0.000000
	__LinkShape__00000008_.TargetMultiplicity = models.MANY
	__LinkShape__00000008_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.SourceMultiplicity = models.MANY
	__LinkShape__00000008_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.X = 885.000000
	__LinkShape__00000008_.Y = 514.000000
	__LinkShape__00000008_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.StartRatio = 0.328664
	__LinkShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.EndRatio = 0.872215
	__LinkShape__00000008_.CornerOffsetRatio = 1.614062

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000003_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000002_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000003_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000006_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000005_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000007_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000008_)
}
