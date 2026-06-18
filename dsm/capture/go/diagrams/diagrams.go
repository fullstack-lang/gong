package diagrams

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/dsm/capture/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `2. Abstract`}).Stage(stage)
	__Classdiagram__00000001_ := (&models.Classdiagram{Name: `1. Global View`}).Stage(stage)
	__Classdiagram__00000002_ := (&models.Classdiagram{Name: `3. Concrete`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2026-03-23T11:40:42Z`}).Stage(stage)

	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-Resource`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Abstract-Library`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `Abstract-Product`}).Stage(stage)
	__GongStructShape__00000004_ := (&models.GongStructShape{Name: `Abstract-Task`}).Stage(stage)
	__GongStructShape__00000006_ := (&models.GongStructShape{Name: `Global View-Concern`}).Stage(stage)
	__GongStructShape__00000007_ := (&models.GongStructShape{Name: `Global View-Stakeholders`}).Stage(stage)
	__GongStructShape__00000008_ := (&models.GongStructShape{Name: `Global View-Deliverable`}).Stage(stage)
	__GongStructShape__00000009_ := (&models.GongStructShape{Name: `3. Concrete-StakeholderShape`}).Stage(stage)
	__GongStructShape__00000010_ := (&models.GongStructShape{Name: `3. Concrete-Stakeholders`}).Stage(stage)
	__GongStructShape__00000011_ := (&models.GongStructShape{Name: `3. Concrete-Diagram`}).Stage(stage)
	__GongStructShape__00000012_ := (&models.GongStructShape{Name: `3. Concrete-ConcernShape`}).Stage(stage)
	__GongStructShape__00000013_ := (&models.GongStructShape{Name: `3. Concrete-Concern`}).Stage(stage)
	__GongStructShape__00000014_ := (&models.GongStructShape{Name: `3. Concrete-StakeholderCompositionShape`}).Stage(stage)
	__GongStructShape__00000015_ := (&models.GongStructShape{Name: `3. Concrete-StakeholderConcernShape`}).Stage(stage)
	__GongStructShape__00000016_ := (&models.GongStructShape{Name: `1. Global View-Concept`}).Stage(stage)
	__GongStructShape__00000017_ := (&models.GongStructShape{Name: `1. Global View-Requirement`}).Stage(stage)
	__GongStructShape__00000018_ := (&models.GongStructShape{Name: `1. Global View-Tool`}).Stage(stage)
	__GongStructShape__00000019_ := (&models.GongStructShape{Name: `1. Global View-SupportLevel`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `RootResources`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `RootDeliverables`}).Stage(stage)
	__LinkShape__00000002_ := (&models.LinkShape{Name: `RootConcerns`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `Concerns`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `Outputs`}).Stage(stage)
	__LinkShape__00000005_ := (&models.LinkShape{Name: `Inputs`}).Stage(stage)
	__LinkShape__00000006_ := (&models.LinkShape{Name: `Resource_Shapes`}).Stage(stage)
	__LinkShape__00000007_ := (&models.LinkShape{Name: `Stakeholder`}).Stage(stage)
	__LinkShape__00000008_ := (&models.LinkShape{Name: `Concern`}).Stage(stage)
	__LinkShape__00000009_ := (&models.LinkShape{Name: `Concern_Shapes`}).Stage(stage)
	__LinkShape__00000010_ := (&models.LinkShape{Name: `Stakeholder`}).Stage(stage)
	__LinkShape__00000011_ := (&models.LinkShape{Name: `Stakeholder`}).Stage(stage)
	__LinkShape__00000012_ := (&models.LinkShape{Name: `Concern`}).Stage(stage)
	__LinkShape__00000014_ := (&models.LinkShape{Name: `ResourceComposition_Shapes`}).Stage(stage)
	__LinkShape__00000015_ := (&models.LinkShape{Name: `StakeholderConcernShapes`}).Stage(stage)
	__LinkShape__00000017_ := (&models.LinkShape{Name: `Concepts`}).Stage(stage)
	__LinkShape__00000018_ := (&models.LinkShape{Name: `Concepts`}).Stage(stage)
	__LinkShape__00000019_ := (&models.LinkShape{Name: `Requirements`}).Stage(stage)
	__LinkShape__00000020_ := (&models.LinkShape{Name: `SupportLevels`}).Stage(stage)
	__LinkShape__00000022_ := (&models.LinkShape{Name: `Tool`}).Stage(stage)

	// insertion point for initialization of values

	__Classdiagram__00000000_.Name = `2. Abstract`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = false
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = false
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[false,false,false,false,false,false,false,true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__Classdiagram__00000001_.Name = `1. Global View`
	__Classdiagram__00000001_.Description = ``
	__Classdiagram__00000001_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000001_.ShowNbInstances = true
	__Classdiagram__00000001_.ShowMultiplicity = true
	__Classdiagram__00000001_.ShowLinkNames = true
	__Classdiagram__00000001_.IsInRenameMode = false
	__Classdiagram__00000001_.IsExpanded = false
	__Classdiagram__00000001_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000001_.NodeGongStructNodeExpansion = `[false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,false,true]`
	__Classdiagram__00000001_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000001_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000001_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000001_.NodeGongNoteNodeExpansion = ``

	__Classdiagram__00000002_.Name = `3. Concrete`
	__Classdiagram__00000002_.Description = ``
	__Classdiagram__00000002_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000002_.ShowNbInstances = false
	__Classdiagram__00000002_.ShowMultiplicity = false
	__Classdiagram__00000002_.ShowLinkNames = true
	__Classdiagram__00000002_.IsInRenameMode = false
	__Classdiagram__00000002_.IsExpanded = true
	__Classdiagram__00000002_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000002_.NodeGongStructNodeExpansion = `[false,false,false,false,false,true,false,true,false,false,false,false,false,false,false,false,false,true,true]`
	__Classdiagram__00000002_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000002_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000002_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000002_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2026-03-23T11:40:42Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000001_.Name = `Default-Resource`
	__GongStructShape__00000001_.X = 612.000000
	__GongStructShape__00000001_.Y = 58.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.Stakeholder{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 63.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Abstract-Library`
	__GongStructShape__00000002_.X = 102.000000
	__GongStructShape__00000002_.Y = 55.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.Library{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 63.000000
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `Abstract-Product`
	__GongStructShape__00000003_.X = 602.000000
	__GongStructShape__00000003_.Y = 395.000000
	__GongStructShape__00000003_.IdentifierMeta = ref_models.Deliverable{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 63.000000
	__GongStructShape__00000003_.IsSelected = false

	__GongStructShape__00000004_.Name = `Abstract-Task`
	__GongStructShape__00000004_.X = 612.000000
	__GongStructShape__00000004_.Y = 167.000000
	__GongStructShape__00000004_.IdentifierMeta = ref_models.Concern{}
	__GongStructShape__00000004_.Width = 240.000000
	__GongStructShape__00000004_.Height = 63.000000
	__GongStructShape__00000004_.IsSelected = false

	__GongStructShape__00000006_.Name = `Global View-Concern`
	__GongStructShape__00000006_.X = 180.000000
	__GongStructShape__00000006_.Y = 238.000000
	__GongStructShape__00000006_.IdentifierMeta = ref_models.Concern{}
	__GongStructShape__00000006_.Width = 240.000000
	__GongStructShape__00000006_.Height = 91.000000
	__GongStructShape__00000006_.IsSelected = false

	__GongStructShape__00000007_.Name = `Global View-Stakeholders`
	__GongStructShape__00000007_.X = 180.000000
	__GongStructShape__00000007_.Y = 39.000000
	__GongStructShape__00000007_.IdentifierMeta = ref_models.Stakeholder{}
	__GongStructShape__00000007_.Width = 240.000000
	__GongStructShape__00000007_.Height = 63.000000
	__GongStructShape__00000007_.IsSelected = false

	__GongStructShape__00000008_.Name = `Global View-Deliverable`
	__GongStructShape__00000008_.X = 631.000000
	__GongStructShape__00000008_.Y = 240.000000
	__GongStructShape__00000008_.IdentifierMeta = ref_models.Deliverable{}
	__GongStructShape__00000008_.Width = 240.000000
	__GongStructShape__00000008_.Height = 90.000000
	__GongStructShape__00000008_.IsSelected = false

	__GongStructShape__00000009_.Name = `3. Concrete-StakeholderShape`
	__GongStructShape__00000009_.X = 634.000000
	__GongStructShape__00000009_.Y = 171.000000
	__GongStructShape__00000009_.IdentifierMeta = ref_models.StakeholderShape{}
	__GongStructShape__00000009_.Width = 240.000000
	__GongStructShape__00000009_.Height = 63.000000
	__GongStructShape__00000009_.IsSelected = false

	__GongStructShape__00000010_.Name = `3. Concrete-Stakeholders`
	__GongStructShape__00000010_.X = 1048.000000
	__GongStructShape__00000010_.Y = 169.000000
	__GongStructShape__00000010_.IdentifierMeta = ref_models.Stakeholder{}
	__GongStructShape__00000010_.Width = 240.000000
	__GongStructShape__00000010_.Height = 63.000000
	__GongStructShape__00000010_.IsSelected = false

	__GongStructShape__00000011_.Name = `3. Concrete-Diagram`
	__GongStructShape__00000011_.X = 26.000000
	__GongStructShape__00000011_.Y = 40.000000
	__GongStructShape__00000011_.IdentifierMeta = ref_models.Diagram{}
	__GongStructShape__00000011_.Width = 240.000000
	__GongStructShape__00000011_.Height = 63.000000
	__GongStructShape__00000011_.IsSelected = false

	__GongStructShape__00000012_.Name = `3. Concrete-ConcernShape`
	__GongStructShape__00000012_.X = 646.000000
	__GongStructShape__00000012_.Y = 519.000000
	__GongStructShape__00000012_.IdentifierMeta = ref_models.ConcernShape{}
	__GongStructShape__00000012_.Width = 240.000000
	__GongStructShape__00000012_.Height = 63.000000
	__GongStructShape__00000012_.IsSelected = false

	__GongStructShape__00000013_.Name = `3. Concrete-Concern`
	__GongStructShape__00000013_.X = 1062.000000
	__GongStructShape__00000013_.Y = 519.000000
	__GongStructShape__00000013_.IdentifierMeta = ref_models.Concern{}
	__GongStructShape__00000013_.Width = 240.000000
	__GongStructShape__00000013_.Height = 63.000000
	__GongStructShape__00000013_.IsSelected = false

	__GongStructShape__00000014_.Name = `3. Concrete-StakeholderCompositionShape`
	__GongStructShape__00000014_.X = 819.000000
	__GongStructShape__00000014_.Y = 41.000000
	__GongStructShape__00000014_.IdentifierMeta = ref_models.StakeholderCompositionShape{}
	__GongStructShape__00000014_.Width = 240.000000
	__GongStructShape__00000014_.Height = 63.000000
	__GongStructShape__00000014_.IsSelected = false

	__GongStructShape__00000015_.Name = `3. Concrete-StakeholderConcernShape`
	__GongStructShape__00000015_.X = 841.000000
	__GongStructShape__00000015_.Y = 349.000000
	__GongStructShape__00000015_.IdentifierMeta = ref_models.StakeholderConcernShape{}
	__GongStructShape__00000015_.Width = 240.000000
	__GongStructShape__00000015_.Height = 63.000000
	__GongStructShape__00000015_.IsSelected = false

	__GongStructShape__00000016_.Name = `1. Global View-Concept`
	__GongStructShape__00000016_.X = 617.000000
	__GongStructShape__00000016_.Y = 452.999985
	__GongStructShape__00000016_.IdentifierMeta = ref_models.Concept{}
	__GongStructShape__00000016_.Width = 240.000000
	__GongStructShape__00000016_.Height = 63.000000
	__GongStructShape__00000016_.IsSelected = false

	__GongStructShape__00000017_.Name = `1. Global View-Requirement`
	__GongStructShape__00000017_.X = 180.000000
	__GongStructShape__00000017_.Y = 448.999985
	__GongStructShape__00000017_.IdentifierMeta = ref_models.Requirement{}
	__GongStructShape__00000017_.Width = 240.000000
	__GongStructShape__00000017_.Height = 63.000000
	__GongStructShape__00000017_.IsSelected = false

	__GongStructShape__00000018_.Name = `1. Global View-Tool`
	__GongStructShape__00000018_.X = 187.000000
	__GongStructShape__00000018_.Y = 724.999985
	__GongStructShape__00000018_.IdentifierMeta = ref_models.Tool{}
	__GongStructShape__00000018_.Width = 240.000000
	__GongStructShape__00000018_.Height = 63.000000
	__GongStructShape__00000018_.IsSelected = false

	__GongStructShape__00000019_.Name = `1. Global View-SupportLevel`
	__GongStructShape__00000019_.X = 178.000000
	__GongStructShape__00000019_.Y = 586.000000
	__GongStructShape__00000019_.IdentifierMeta = ref_models.SupportLevel{}
	__GongStructShape__00000019_.Width = 240.000000
	__GongStructShape__00000019_.Height = 63.000000
	__GongStructShape__00000019_.IsSelected = false

	__LinkShape__00000000_.Name = `RootResources`
	__LinkShape__00000000_.IdentifierMeta = ref_models.Library{}.RootStakeholders
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.Stakeholder{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.MANY
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 485.000000
	__LinkShape__00000000_.Y = 214.000000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.StartRatio = 0.500000
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.500000
	__LinkShape__00000000_.CornerOffsetRatio = 1.380000

	__LinkShape__00000001_.Name = `RootDeliverables`
	__LinkShape__00000001_.IdentifierMeta = ref_models.Library{}.RootDeliverables
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.Deliverable{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.MANY
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 709.500000
	__LinkShape__00000001_.Y = 89.500000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.500000
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.500000
	__LinkShape__00000001_.CornerOffsetRatio = 1.380000

	__LinkShape__00000002_.Name = `RootConcerns`
	__LinkShape__00000002_.IdentifierMeta = ref_models.Library{}.RootConcerns
	__LinkShape__00000002_.FieldTypeIdentifierMeta = ref_models.Concern{}
	__LinkShape__00000002_.FieldOffsetX = 0.000000
	__LinkShape__00000002_.FieldOffsetY = 0.000000
	__LinkShape__00000002_.TargetMultiplicity = models.MANY
	__LinkShape__00000002_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.SourceMultiplicity = models.MANY
	__LinkShape__00000002_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.X = 709.500000
	__LinkShape__00000002_.Y = 244.500000
	__LinkShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.StartRatio = 0.500000
	__LinkShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.EndRatio = 0.500000
	__LinkShape__00000002_.CornerOffsetRatio = 1.380000

	__LinkShape__00000003_.Name = `Concerns`
	__LinkShape__00000003_.IdentifierMeta = ref_models.Stakeholder{}.Concerns
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.Concern{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.MANY
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 540.000000
	__LinkShape__00000003_.Y = 170.000000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000003_.StartRatio = 0.696908
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000003_.EndRatio = 0.696908
	__LinkShape__00000003_.CornerOffsetRatio = 1.591518

	__LinkShape__00000004_.Name = `Outputs`
	__LinkShape__00000004_.IdentifierMeta = ref_models.Concern{}.Outputs
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.Deliverable{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.MANY
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 746.000000
	__LinkShape__00000004_.Y = 270.000000
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.StartRatio = 0.178819
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.EndRatio = 0.226438
	__LinkShape__00000004_.CornerOffsetRatio = 1.380000

	__LinkShape__00000005_.Name = `Inputs`
	__LinkShape__00000005_.IdentifierMeta = ref_models.Concern{}.Inputs
	__LinkShape__00000005_.FieldTypeIdentifierMeta = ref_models.Deliverable{}
	__LinkShape__00000005_.FieldOffsetX = 0.000000
	__LinkShape__00000005_.FieldOffsetY = 0.000000
	__LinkShape__00000005_.TargetMultiplicity = models.MANY
	__LinkShape__00000005_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.SourceMultiplicity = models.MANY
	__LinkShape__00000005_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.X = 746.000000
	__LinkShape__00000005_.Y = 270.000000
	__LinkShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.StartRatio = 0.877232
	__LinkShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.EndRatio = 0.877232
	__LinkShape__00000005_.CornerOffsetRatio = 1.380000

	__LinkShape__00000006_.Name = `Resource_Shapes`
	__LinkShape__00000006_.IdentifierMeta = ref_models.Diagram{}.Stakeholder_Shapes
	__LinkShape__00000006_.FieldTypeIdentifierMeta = ref_models.StakeholderShape{}
	__LinkShape__00000006_.FieldOffsetX = 0.000000
	__LinkShape__00000006_.FieldOffsetY = 0.000000
	__LinkShape__00000006_.TargetMultiplicity = models.MANY
	__LinkShape__00000006_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.SourceMultiplicity = models.MANY
	__LinkShape__00000006_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.X = 626.500000
	__LinkShape__00000006_.Y = 133.000000
	__LinkShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.StartRatio = 0.500000
	__LinkShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.EndRatio = 0.500000
	__LinkShape__00000006_.CornerOffsetRatio = 1.380000

	__LinkShape__00000007_.Name = `Stakeholder`
	__LinkShape__00000007_.IdentifierMeta = ref_models.StakeholderShape{}.Stakeholder
	__LinkShape__00000007_.FieldTypeIdentifierMeta = ref_models.Stakeholder{}
	__LinkShape__00000007_.FieldOffsetX = 0.000000
	__LinkShape__00000007_.FieldOffsetY = 0.000000
	__LinkShape__00000007_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000007_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.SourceMultiplicity = models.MANY
	__LinkShape__00000007_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.X = 1084.000000
	__LinkShape__00000007_.Y = 201.500000
	__LinkShape__00000007_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.StartRatio = 0.607515
	__LinkShape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.EndRatio = 0.671007
	__LinkShape__00000007_.CornerOffsetRatio = 1.392741

	__LinkShape__00000008_.Name = `Concern`
	__LinkShape__00000008_.IdentifierMeta = ref_models.ConcernShape{}.Concern
	__LinkShape__00000008_.FieldTypeIdentifierMeta = ref_models.Concern{}
	__LinkShape__00000008_.FieldOffsetX = 0.000000
	__LinkShape__00000008_.FieldOffsetY = 0.000000
	__LinkShape__00000008_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000008_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.SourceMultiplicity = models.MANY
	__LinkShape__00000008_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.X = 1149.500000
	__LinkShape__00000008_.Y = 386.000000
	__LinkShape__00000008_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.StartRatio = 0.500000
	__LinkShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.EndRatio = 0.500000
	__LinkShape__00000008_.CornerOffsetRatio = 1.380000

	__LinkShape__00000009_.Name = `Concern_Shapes`
	__LinkShape__00000009_.IdentifierMeta = ref_models.Diagram{}.Concern_Shapes
	__LinkShape__00000009_.FieldTypeIdentifierMeta = ref_models.ConcernShape{}
	__LinkShape__00000009_.FieldOffsetX = 0.000000
	__LinkShape__00000009_.FieldOffsetY = 0.000000
	__LinkShape__00000009_.TargetMultiplicity = models.MANY
	__LinkShape__00000009_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.SourceMultiplicity = models.MANY
	__LinkShape__00000009_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.X = 687.000000
	__LinkShape__00000009_.Y = 224.000000
	__LinkShape__00000009_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.StartRatio = 0.500000
	__LinkShape__00000009_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.EndRatio = 0.500000
	__LinkShape__00000009_.CornerOffsetRatio = 1.380000

	__LinkShape__00000010_.Name = `Stakeholder`
	__LinkShape__00000010_.IdentifierMeta = ref_models.StakeholderCompositionShape{}.Stakeholder
	__LinkShape__00000010_.FieldTypeIdentifierMeta = ref_models.Stakeholder{}
	__LinkShape__00000010_.FieldOffsetX = 0.000000
	__LinkShape__00000010_.FieldOffsetY = 0.000000
	__LinkShape__00000010_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000010_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.SourceMultiplicity = models.MANY
	__LinkShape__00000010_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.X = 1136.500000
	__LinkShape__00000010_.Y = 133.500000
	__LinkShape__00000010_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000010_.StartRatio = 0.500000
	__LinkShape__00000010_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000010_.EndRatio = 0.463574
	__LinkShape__00000010_.CornerOffsetRatio = 2.020213

	__LinkShape__00000011_.Name = `Stakeholder`
	__LinkShape__00000011_.IdentifierMeta = ref_models.StakeholderConcernShape{}.Stakeholder
	__LinkShape__00000011_.FieldTypeIdentifierMeta = ref_models.Stakeholder{}
	__LinkShape__00000011_.FieldOffsetX = 0.000000
	__LinkShape__00000011_.FieldOffsetY = 0.000000
	__LinkShape__00000011_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000011_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.SourceMultiplicity = models.MANY
	__LinkShape__00000011_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.X = 1143.000000
	__LinkShape__00000011_.Y = 288.500000
	__LinkShape__00000011_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000011_.StartRatio = 0.115451
	__LinkShape__00000011_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000011_.EndRatio = 0.534408
	__LinkShape__00000011_.CornerOffsetRatio = -0.408358

	__LinkShape__00000012_.Name = `Concern`
	__LinkShape__00000012_.IdentifierMeta = ref_models.StakeholderConcernShape{}.Concern
	__LinkShape__00000012_.FieldTypeIdentifierMeta = ref_models.Concern{}
	__LinkShape__00000012_.FieldOffsetX = 0.000000
	__LinkShape__00000012_.FieldOffsetY = 0.000000
	__LinkShape__00000012_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000012_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.SourceMultiplicity = models.MANY
	__LinkShape__00000012_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.X = 1150.000000
	__LinkShape__00000012_.Y = 463.500000
	__LinkShape__00000012_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000012_.StartRatio = 0.500000
	__LinkShape__00000012_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000012_.EndRatio = 0.463574
	__LinkShape__00000012_.CornerOffsetRatio = 2.528150

	__LinkShape__00000014_.Name = `ResourceComposition_Shapes`
	__LinkShape__00000014_.IdentifierMeta = ref_models.Diagram{}.ResourceComposition_Shapes
	__LinkShape__00000014_.FieldTypeIdentifierMeta = ref_models.StakeholderCompositionShape{}
	__LinkShape__00000014_.FieldOffsetX = 0.000000
	__LinkShape__00000014_.FieldOffsetY = 0.000000
	__LinkShape__00000014_.TargetMultiplicity = models.MANY
	__LinkShape__00000014_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000014_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000014_.SourceMultiplicity = models.MANY
	__LinkShape__00000014_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000014_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000014_.X = 679.000000
	__LinkShape__00000014_.Y = 65.000000
	__LinkShape__00000014_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000014_.StartRatio = 0.500000
	__LinkShape__00000014_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000014_.EndRatio = 0.500000
	__LinkShape__00000014_.CornerOffsetRatio = 1.380000

	__LinkShape__00000015_.Name = `StakeholderConcernShapes`
	__LinkShape__00000015_.IdentifierMeta = ref_models.Diagram{}.StakeholderConcernShapes
	__LinkShape__00000015_.FieldTypeIdentifierMeta = ref_models.StakeholderConcernShape{}
	__LinkShape__00000015_.FieldOffsetX = 0.000000
	__LinkShape__00000015_.FieldOffsetY = 0.000000
	__LinkShape__00000015_.TargetMultiplicity = models.MANY
	__LinkShape__00000015_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000015_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000015_.SourceMultiplicity = models.MANY
	__LinkShape__00000015_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000015_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000015_.X = 685.500000
	__LinkShape__00000015_.Y = 220.000000
	__LinkShape__00000015_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000015_.StartRatio = 0.500000
	__LinkShape__00000015_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000015_.EndRatio = 0.500000
	__LinkShape__00000015_.CornerOffsetRatio = 1.380000

	__LinkShape__00000017_.Name = `Concepts`
	__LinkShape__00000017_.IdentifierMeta = ref_models.Deliverable{}.Concepts
	__LinkShape__00000017_.FieldTypeIdentifierMeta = ref_models.Concept{}
	__LinkShape__00000017_.FieldOffsetX = 0.000000
	__LinkShape__00000017_.FieldOffsetY = 0.000000
	__LinkShape__00000017_.TargetMultiplicity = models.MANY
	__LinkShape__00000017_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000017_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000017_.SourceMultiplicity = models.MANY
	__LinkShape__00000017_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000017_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000017_.X = 984.000000
	__LinkShape__00000017_.Y = 391.499992
	__LinkShape__00000017_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000017_.StartRatio = 0.459462
	__LinkShape__00000017_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000017_.EndRatio = 0.559462
	__LinkShape__00000017_.CornerOffsetRatio = 1.633565

	__LinkShape__00000018_.Name = `Concepts`
	__LinkShape__00000018_.IdentifierMeta = ref_models.Requirement{}.Concepts
	__LinkShape__00000018_.FieldTypeIdentifierMeta = ref_models.Concept{}
	__LinkShape__00000018_.FieldOffsetX = 0.000000
	__LinkShape__00000018_.FieldOffsetY = 0.000000
	__LinkShape__00000018_.TargetMultiplicity = models.MANY
	__LinkShape__00000018_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000018_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000018_.SourceMultiplicity = models.MANY
	__LinkShape__00000018_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000018_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000018_.X = 758.500000
	__LinkShape__00000018_.Y = 482.499985
	__LinkShape__00000018_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000018_.StartRatio = 0.603175
	__LinkShape__00000018_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000018_.EndRatio = 0.500000
	__LinkShape__00000018_.CornerOffsetRatio = 1.380000

	__LinkShape__00000019_.Name = `Requirements`
	__LinkShape__00000019_.IdentifierMeta = ref_models.Concern{}.Requirements
	__LinkShape__00000019_.FieldTypeIdentifierMeta = ref_models.Requirement{}
	__LinkShape__00000019_.FieldOffsetX = 0.000000
	__LinkShape__00000019_.FieldOffsetY = 0.000000
	__LinkShape__00000019_.TargetMultiplicity = models.MANY
	__LinkShape__00000019_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000019_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000019_.SourceMultiplicity = models.MANY
	__LinkShape__00000019_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000019_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000019_.X = 540.000000
	__LinkShape__00000019_.Y = 388.999992
	__LinkShape__00000019_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000019_.StartRatio = 0.738628
	__LinkShape__00000019_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000019_.EndRatio = 0.738628
	__LinkShape__00000019_.CornerOffsetRatio = 1.373626

	__LinkShape__00000020_.Name = `SupportLevels`
	__LinkShape__00000020_.IdentifierMeta = ref_models.Requirement{}.SupportLevels
	__LinkShape__00000020_.FieldTypeIdentifierMeta = ref_models.SupportLevel{}
	__LinkShape__00000020_.FieldOffsetX = 0.000000
	__LinkShape__00000020_.FieldOffsetY = 0.000000
	__LinkShape__00000020_.TargetMultiplicity = models.MANY
	__LinkShape__00000020_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000020_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000020_.SourceMultiplicity = models.MANY
	__LinkShape__00000020_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000020_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000020_.X = 539.000000
	__LinkShape__00000020_.Y = 548.999992
	__LinkShape__00000020_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000020_.StartRatio = 0.546962
	__LinkShape__00000020_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000020_.EndRatio = 0.571962
	__LinkShape__00000020_.CornerOffsetRatio = 1.603175

	__LinkShape__00000022_.Name = `Tool`
	__LinkShape__00000022_.IdentifierMeta = ref_models.SupportLevel{}.Tool
	__LinkShape__00000022_.FieldTypeIdentifierMeta = ref_models.Tool{}
	__LinkShape__00000022_.FieldOffsetX = 0.000000
	__LinkShape__00000022_.FieldOffsetY = 0.000000
	__LinkShape__00000022_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000022_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000022_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000022_.SourceMultiplicity = models.MANY
	__LinkShape__00000022_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000022_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000022_.X = 542.500000
	__LinkShape__00000022_.Y = 686.999992
	__LinkShape__00000022_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000022_.StartRatio = 0.530295
	__LinkShape__00000022_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000022_.EndRatio = 0.455295
	__LinkShape__00000022_.CornerOffsetRatio = 1.539683

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000003_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000004_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000006_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000007_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000008_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000016_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000017_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000018_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000019_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000009_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000010_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000011_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000012_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000013_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000014_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000015_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000001_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000002_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000002_
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000002_)
	__GongStructShape__00000006_.LinkShapes = append(__GongStructShape__00000006_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000006_.LinkShapes = append(__GongStructShape__00000006_.LinkShapes, __LinkShape__00000005_)
	__GongStructShape__00000006_.LinkShapes = append(__GongStructShape__00000006_.LinkShapes, __LinkShape__00000019_)
	__GongStructShape__00000007_.LinkShapes = append(__GongStructShape__00000007_.LinkShapes, __LinkShape__00000003_)
	__GongStructShape__00000008_.LinkShapes = append(__GongStructShape__00000008_.LinkShapes, __LinkShape__00000017_)
	__GongStructShape__00000009_.LinkShapes = append(__GongStructShape__00000009_.LinkShapes, __LinkShape__00000007_)
	__GongStructShape__00000011_.LinkShapes = append(__GongStructShape__00000011_.LinkShapes, __LinkShape__00000006_)
	__GongStructShape__00000011_.LinkShapes = append(__GongStructShape__00000011_.LinkShapes, __LinkShape__00000009_)
	__GongStructShape__00000011_.LinkShapes = append(__GongStructShape__00000011_.LinkShapes, __LinkShape__00000014_)
	__GongStructShape__00000011_.LinkShapes = append(__GongStructShape__00000011_.LinkShapes, __LinkShape__00000015_)
	__GongStructShape__00000012_.LinkShapes = append(__GongStructShape__00000012_.LinkShapes, __LinkShape__00000008_)
	__GongStructShape__00000014_.LinkShapes = append(__GongStructShape__00000014_.LinkShapes, __LinkShape__00000010_)
	__GongStructShape__00000015_.LinkShapes = append(__GongStructShape__00000015_.LinkShapes, __LinkShape__00000011_)
	__GongStructShape__00000015_.LinkShapes = append(__GongStructShape__00000015_.LinkShapes, __LinkShape__00000012_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000018_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000020_)
	__GongStructShape__00000019_.LinkShapes = append(__GongStructShape__00000019_.LinkShapes, __LinkShape__00000022_)
}
