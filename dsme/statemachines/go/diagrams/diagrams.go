package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/dsme/statemachines/go/models"
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

	__AttributeShape__00000000_ := (&models.AttributeShape{Name: `IsDecisionNode`}).Stage(stage)
	__AttributeShape__00000001_ := (&models.AttributeShape{Name: `IsFictif`}).Stage(stage)
	__AttributeShape__00000002_ := (&models.AttributeShape{Name: `IsEndState`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `abstract syntax`}).Stage(stage)
	__Classdiagram__00000001_ := (&models.Classdiagram{Name: `concrete syntax`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2025-11-29T13:01:53Z`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-Architecture`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-MessageType`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Default-Role`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `Default-State`}).Stage(stage)
	__GongStructShape__00000004_ := (&models.GongStructShape{Name: `Default-StateMachine`}).Stage(stage)
	__GongStructShape__00000005_ := (&models.GongStructShape{Name: `Default-Transition`}).Stage(stage)
	__GongStructShape__00000006_ := (&models.GongStructShape{Name: `concrete syntax-Diagram`}).Stage(stage)
	__GongStructShape__00000007_ := (&models.GongStructShape{Name: `concrete syntax-StateMachine`}).Stage(stage)
	__GongStructShape__00000008_ := (&models.GongStructShape{Name: `concrete syntax-StateShape`}).Stage(stage)
	__GongStructShape__00000009_ := (&models.GongStructShape{Name: `concrete syntax-Transition_Shape`}).Stage(stage)
	__GongStructShape__00000010_ := (&models.GongStructShape{Name: `concrete syntax-Transition`}).Stage(stage)
	__GongStructShape__00000011_ := (&models.GongStructShape{Name: `concrete syntax-State`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `Start`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `End`}).Stage(stage)
	__LinkShape__00000002_ := (&models.LinkShape{Name: `RolesWithPermissions`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `GeneratedMessages`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `States`}).Stage(stage)
	__LinkShape__00000005_ := (&models.LinkShape{Name: `InitialState`}).Stage(stage)
	__LinkShape__00000006_ := (&models.LinkShape{Name: `SubStates`}).Stage(stage)
	__LinkShape__00000007_ := (&models.LinkShape{Name: `RolesWithSamePermissions`}).Stage(stage)
	__LinkShape__00000008_ := (&models.LinkShape{Name: `StateMachines`}).Stage(stage)
	__LinkShape__00000009_ := (&models.LinkShape{Name: `Roles`}).Stage(stage)
	__LinkShape__00000010_ := (&models.LinkShape{Name: `Diagrams`}).Stage(stage)
	__LinkShape__00000011_ := (&models.LinkShape{Name: `State_Shapes`}).Stage(stage)
	__LinkShape__00000012_ := (&models.LinkShape{Name: `Transition_Shapes`}).Stage(stage)
	__LinkShape__00000013_ := (&models.LinkShape{Name: `State`}).Stage(stage)
	__LinkShape__00000014_ := (&models.LinkShape{Name: `Transition`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `IsDecisionNode`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.State{}.IsDecisionNode
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `State`
	__AttributeShape__00000000_.Fieldtypename = `bool`

	__AttributeShape__00000001_.Name = `IsFictif`
	__AttributeShape__00000001_.IdentifierMeta = ref_models.State{}.IsFictif
	__AttributeShape__00000001_.FieldTypeAsString = ``
	__AttributeShape__00000001_.Structname = `State`
	__AttributeShape__00000001_.Fieldtypename = `bool`

	__AttributeShape__00000002_.Name = `IsEndState`
	__AttributeShape__00000002_.IdentifierMeta = ref_models.State{}.IsEndState
	__AttributeShape__00000002_.FieldTypeAsString = ``
	__AttributeShape__00000002_.Structname = `State`
	__AttributeShape__00000002_.Fieldtypename = `bool`

	__Classdiagram__00000000_.Name = `abstract syntax`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = false
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = false
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[true,false,false,true,true,false,true,true,true,false,false,false]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__Classdiagram__00000001_.Name = `concrete syntax`
	__Classdiagram__00000001_.Description = ``
	__Classdiagram__00000001_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000001_.ShowNbInstances = true
	__Classdiagram__00000001_.ShowMultiplicity = true
	__Classdiagram__00000001_.ShowLinkNames = true
	__Classdiagram__00000001_.IsInRenameMode = false
	__Classdiagram__00000001_.IsExpanded = true
	__Classdiagram__00000001_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000001_.NodeGongStructNodeExpansion = `[false,true,false,false,false,false,false,false,true,false,false,false]`
	__Classdiagram__00000001_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000001_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000001_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000001_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2025-11-29T13:01:53Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_.Name = `Default-Architecture`
	__GongStructShape__00000000_.X = 91.000000
	__GongStructShape__00000000_.Y = 45.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.Architecture{}
	__GongStructShape__00000000_.Width = 1476.000000
	__GongStructShape__00000000_.Height = 70.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-MessageType`
	__GongStructShape__00000001_.X = 1269.000000
	__GongStructShape__00000001_.Y = 600.000015
	__GongStructShape__00000001_.IdentifierMeta = ref_models.MessageType{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 63.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Default-Role`
	__GongStructShape__00000002_.X = 1262.000000
	__GongStructShape__00000002_.Y = 247.999985
	__GongStructShape__00000002_.IdentifierMeta = ref_models.Role{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 178.000000
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `Default-State`
	__GongStructShape__00000003_.X = 379.000000
	__GongStructShape__00000003_.Y = 375.999985
	__GongStructShape__00000003_.IdentifierMeta = ref_models.State{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 123.000000
	__GongStructShape__00000003_.IsSelected = false

	__GongStructShape__00000004_.Name = `Default-StateMachine`
	__GongStructShape__00000004_.X = 85.000000
	__GongStructShape__00000004_.Y = 199.000000
	__GongStructShape__00000004_.IdentifierMeta = ref_models.StateMachine{}
	__GongStructShape__00000004_.Width = 240.000000
	__GongStructShape__00000004_.Height = 141.999985
	__GongStructShape__00000004_.IsSelected = false

	__GongStructShape__00000005_.Name = `Default-Transition`
	__GongStructShape__00000005_.X = 810.000000
	__GongStructShape__00000005_.Y = 250.999985
	__GongStructShape__00000005_.IdentifierMeta = ref_models.Transition{}
	__GongStructShape__00000005_.Width = 240.000000
	__GongStructShape__00000005_.Height = 315.000031
	__GongStructShape__00000005_.IsSelected = false

	__GongStructShape__00000006_.Name = `concrete syntax-Diagram`
	__GongStructShape__00000006_.X = 98.000000
	__GongStructShape__00000006_.Y = 274.999985
	__GongStructShape__00000006_.IdentifierMeta = ref_models.Diagram{}
	__GongStructShape__00000006_.Width = 240.000000
	__GongStructShape__00000006_.Height = 63.000000
	__GongStructShape__00000006_.IsSelected = false

	__GongStructShape__00000007_.Name = `concrete syntax-StateMachine`
	__GongStructShape__00000007_.X = 93.000000
	__GongStructShape__00000007_.Y = 106.000000
	__GongStructShape__00000007_.IdentifierMeta = ref_models.StateMachine{}
	__GongStructShape__00000007_.Width = 240.000000
	__GongStructShape__00000007_.Height = 63.000000
	__GongStructShape__00000007_.IsSelected = false

	__GongStructShape__00000008_.Name = `concrete syntax-StateShape`
	__GongStructShape__00000008_.X = 656.000000
	__GongStructShape__00000008_.Y = 110.000000
	__GongStructShape__00000008_.IdentifierMeta = ref_models.StateShape{}
	__GongStructShape__00000008_.Width = 240.000000
	__GongStructShape__00000008_.Height = 63.000000
	__GongStructShape__00000008_.IsSelected = false

	__GongStructShape__00000009_.Name = `concrete syntax-Transition_Shape`
	__GongStructShape__00000009_.X = 672.000000
	__GongStructShape__00000009_.Y = 274.999985
	__GongStructShape__00000009_.IdentifierMeta = ref_models.Transition_Shape{}
	__GongStructShape__00000009_.Width = 240.000000
	__GongStructShape__00000009_.Height = 63.000000
	__GongStructShape__00000009_.IsSelected = false

	__GongStructShape__00000010_.Name = `concrete syntax-Transition`
	__GongStructShape__00000010_.X = 1072.000000
	__GongStructShape__00000010_.Y = 272.999985
	__GongStructShape__00000010_.IdentifierMeta = ref_models.Transition{}
	__GongStructShape__00000010_.Width = 240.000000
	__GongStructShape__00000010_.Height = 63.000000
	__GongStructShape__00000010_.IsSelected = false

	__GongStructShape__00000011_.Name = `concrete syntax-State`
	__GongStructShape__00000011_.X = 1063.000000
	__GongStructShape__00000011_.Y = 111.000000
	__GongStructShape__00000011_.IdentifierMeta = ref_models.State{}
	__GongStructShape__00000011_.Width = 240.000000
	__GongStructShape__00000011_.Height = 63.000000
	__GongStructShape__00000011_.IsSelected = false

	__LinkShape__00000000_.Name = `Start`
	__LinkShape__00000000_.IdentifierMeta = ref_models.Transition{}.Start
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 406.000000
	__LinkShape__00000000_.Y = 76.500000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.StartRatio = 0.479387
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.284609
	__LinkShape__00000000_.CornerOffsetRatio = -0.331120

	__LinkShape__00000001_.Name = `End`
	__LinkShape__00000001_.IdentifierMeta = ref_models.Transition{}.End
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 406.000000
	__LinkShape__00000001_.Y = 76.500000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.707958
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.731764
	__LinkShape__00000001_.CornerOffsetRatio = -0.414453

	__LinkShape__00000002_.Name = `RolesWithPermissions`
	__LinkShape__00000002_.IdentifierMeta = ref_models.Transition{}.RolesWithPermissions
	__LinkShape__00000002_.FieldTypeIdentifierMeta = ref_models.Role{}
	__LinkShape__00000002_.FieldOffsetX = 0.000000
	__LinkShape__00000002_.FieldOffsetY = 0.000000
	__LinkShape__00000002_.TargetMultiplicity = models.MANY
	__LinkShape__00000002_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.SourceMultiplicity = models.MANY
	__LinkShape__00000002_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.X = 396.000000
	__LinkShape__00000002_.Y = 87.500000
	__LinkShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.StartRatio = 0.149228
	__LinkShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.EndRatio = 0.258466
	__LinkShape__00000002_.CornerOffsetRatio = 1.380000

	__LinkShape__00000003_.Name = `GeneratedMessages`
	__LinkShape__00000003_.IdentifierMeta = ref_models.Transition{}.GeneratedMessages
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.MessageType{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.MANY
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 384.000000
	__LinkShape__00000003_.Y = 51.000000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000003_.StartRatio = 0.618880
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.EndRatio = 0.500000
	__LinkShape__00000003_.CornerOffsetRatio = 1.047729

	__LinkShape__00000004_.Name = `States`
	__LinkShape__00000004_.IdentifierMeta = ref_models.StateMachine{}.States
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.MANY
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 838.500000
	__LinkShape__00000004_.Y = 199.499992
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.StartRatio = 0.232443
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000004_.EndRatio = 0.748047
	__LinkShape__00000004_.CornerOffsetRatio = 2.254078

	__LinkShape__00000005_.Name = `InitialState`
	__LinkShape__00000005_.IdentifierMeta = ref_models.StateMachine{}.InitialState
	__LinkShape__00000005_.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__00000005_.FieldOffsetX = 0.000000
	__LinkShape__00000005_.FieldOffsetY = 0.000000
	__LinkShape__00000005_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000005_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.SourceMultiplicity = models.MANY
	__LinkShape__00000005_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.X = 838.500000
	__LinkShape__00000005_.Y = 199.499992
	__LinkShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.StartRatio = 0.795824
	__LinkShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000005_.EndRatio = 0.323047
	__LinkShape__00000005_.CornerOffsetRatio = 2.301698

	__LinkShape__00000006_.Name = `SubStates`
	__LinkShape__00000006_.IdentifierMeta = ref_models.State{}.SubStates
	__LinkShape__00000006_.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__00000006_.FieldOffsetX = 0.000000
	__LinkShape__00000006_.FieldOffsetY = 0.000000
	__LinkShape__00000006_.TargetMultiplicity = models.MANY
	__LinkShape__00000006_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.SourceMultiplicity = models.MANY
	__LinkShape__00000006_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.X = 1244.000000
	__LinkShape__00000006_.Y = 444.499985
	__LinkShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000006_.StartRatio = 0.227214
	__LinkShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000006_.EndRatio = 0.793880
	__LinkShape__00000006_.CornerOffsetRatio = 1.674853

	__LinkShape__00000007_.Name = `RolesWithSamePermissions`
	__LinkShape__00000007_.IdentifierMeta = ref_models.Role{}.RolesWithSamePermissions
	__LinkShape__00000007_.FieldTypeIdentifierMeta = ref_models.Role{}
	__LinkShape__00000007_.FieldOffsetX = 0.000000
	__LinkShape__00000007_.FieldOffsetY = 0.000000
	__LinkShape__00000007_.TargetMultiplicity = models.MANY
	__LinkShape__00000007_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.SourceMultiplicity = models.MANY
	__LinkShape__00000007_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.X = 443.000000
	__LinkShape__00000007_.Y = 404.499985
	__LinkShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000007_.StartRatio = 0.760547
	__LinkShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000007_.EndRatio = 0.468880
	__LinkShape__00000007_.CornerOffsetRatio = 1.427005

	__LinkShape__00000008_.Name = `StateMachines`
	__LinkShape__00000008_.IdentifierMeta = ref_models.Architecture{}.StateMachines
	__LinkShape__00000008_.FieldTypeIdentifierMeta = ref_models.StateMachine{}
	__LinkShape__00000008_.FieldOffsetX = 0.000000
	__LinkShape__00000008_.FieldOffsetY = 0.000000
	__LinkShape__00000008_.TargetMultiplicity = models.MANY
	__LinkShape__00000008_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.SourceMultiplicity = models.MANY
	__LinkShape__00000008_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.X = 829.500000
	__LinkShape__00000008_.Y = 197.499992
	__LinkShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000008_.StartRatio = 0.146024
	__LinkShape__00000008_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000008_.EndRatio = 0.985547
	__LinkShape__00000008_.CornerOffsetRatio = 1.442956

	__LinkShape__00000009_.Name = `Roles`
	__LinkShape__00000009_.IdentifierMeta = ref_models.Architecture{}.Roles
	__LinkShape__00000009_.FieldTypeIdentifierMeta = ref_models.Role{}
	__LinkShape__00000009_.FieldOffsetX = 0.000000
	__LinkShape__00000009_.FieldOffsetY = 0.000000
	__LinkShape__00000009_.TargetMultiplicity = models.MANY
	__LinkShape__00000009_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.SourceMultiplicity = models.MANY
	__LinkShape__00000009_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.X = 471.500000
	__LinkShape__00000009_.Y = 254.499992
	__LinkShape__00000009_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000009_.StartRatio = 0.927189
	__LinkShape__00000009_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000009_.EndRatio = 0.843880
	__LinkShape__00000009_.CornerOffsetRatio = 1.700099

	__LinkShape__00000010_.Name = `Diagrams`
	__LinkShape__00000010_.IdentifierMeta = ref_models.StateMachine{}.Diagrams
	__LinkShape__00000010_.FieldTypeIdentifierMeta = ref_models.Diagram{}
	__LinkShape__00000010_.FieldOffsetX = 0.000000
	__LinkShape__00000010_.FieldOffsetY = 0.000000
	__LinkShape__00000010_.TargetMultiplicity = models.MANY
	__LinkShape__00000010_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.SourceMultiplicity = models.MANY
	__LinkShape__00000010_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.X = 443.000000
	__LinkShape__00000010_.Y = 65.500000
	__LinkShape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000010_.StartRatio = 0.437630
	__LinkShape__00000010_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000010_.EndRatio = 0.387630
	__LinkShape__00000010_.CornerOffsetRatio = 1.571539

	__LinkShape__00000011_.Name = `State_Shapes`
	__LinkShape__00000011_.IdentifierMeta = ref_models.Diagram{}.State_Shapes
	__LinkShape__00000011_.FieldTypeIdentifierMeta = ref_models.StateShape{}
	__LinkShape__00000011_.FieldOffsetX = 0.000000
	__LinkShape__00000011_.FieldOffsetY = 0.000000
	__LinkShape__00000011_.TargetMultiplicity = models.MANY
	__LinkShape__00000011_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.SourceMultiplicity = models.MANY
	__LinkShape__00000011_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.X = 655.000000
	__LinkShape__00000011_.Y = 223.499992
	__LinkShape__00000011_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000011_.StartRatio = 0.500000
	__LinkShape__00000011_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000011_.EndRatio = 0.500000
	__LinkShape__00000011_.CornerOffsetRatio = 1.380000

	__LinkShape__00000012_.Name = `Transition_Shapes`
	__LinkShape__00000012_.IdentifierMeta = ref_models.Diagram{}.Transition_Shapes
	__LinkShape__00000012_.FieldTypeIdentifierMeta = ref_models.Transition_Shape{}
	__LinkShape__00000012_.FieldOffsetX = 0.000000
	__LinkShape__00000012_.FieldOffsetY = 0.000000
	__LinkShape__00000012_.TargetMultiplicity = models.MANY
	__LinkShape__00000012_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.SourceMultiplicity = models.MANY
	__LinkShape__00000012_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.X = 658.000000
	__LinkShape__00000012_.Y = 303.999985
	__LinkShape__00000012_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000012_.StartRatio = 0.500000
	__LinkShape__00000012_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000012_.EndRatio = 0.500000
	__LinkShape__00000012_.CornerOffsetRatio = 1.380000

	__LinkShape__00000013_.Name = `State`
	__LinkShape__00000013_.IdentifierMeta = ref_models.StateShape{}.State
	__LinkShape__00000013_.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__00000013_.FieldOffsetX = 0.000000
	__LinkShape__00000013_.FieldOffsetY = 0.000000
	__LinkShape__00000013_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000013_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000013_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000013_.SourceMultiplicity = models.MANY
	__LinkShape__00000013_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000013_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000013_.X = 1224.000000
	__LinkShape__00000013_.Y = 221.999992
	__LinkShape__00000013_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000013_.StartRatio = 0.500000
	__LinkShape__00000013_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000013_.EndRatio = 0.500000
	__LinkShape__00000013_.CornerOffsetRatio = 1.380000

	__LinkShape__00000014_.Name = `Transition`
	__LinkShape__00000014_.IdentifierMeta = ref_models.Transition_Shape{}.Transition
	__LinkShape__00000014_.FieldTypeIdentifierMeta = ref_models.Transition{}
	__LinkShape__00000014_.FieldOffsetX = 0.000000
	__LinkShape__00000014_.FieldOffsetY = 0.000000
	__LinkShape__00000014_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000014_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000014_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000014_.SourceMultiplicity = models.MANY
	__LinkShape__00000014_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000014_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000014_.X = 1228.500000
	__LinkShape__00000014_.Y = 223.999992
	__LinkShape__00000014_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000014_.StartRatio = 0.500000
	__LinkShape__00000014_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000014_.EndRatio = 0.500000
	__LinkShape__00000014_.CornerOffsetRatio = 1.380000

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000003_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000004_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000005_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000006_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000007_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000008_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000009_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000010_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000011_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000001_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000008_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000009_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000007_)
	__GongStructShape__00000003_.AttributeShapes = append(__GongStructShape__00000003_.AttributeShapes, __AttributeShape__00000000_)
	__GongStructShape__00000003_.AttributeShapes = append(__GongStructShape__00000003_.AttributeShapes, __AttributeShape__00000001_)
	__GongStructShape__00000003_.AttributeShapes = append(__GongStructShape__00000003_.AttributeShapes, __AttributeShape__00000002_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000006_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000005_)
	__GongStructShape__00000005_.LinkShapes = append(__GongStructShape__00000005_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000005_.LinkShapes = append(__GongStructShape__00000005_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000005_.LinkShapes = append(__GongStructShape__00000005_.LinkShapes, __LinkShape__00000002_)
	__GongStructShape__00000005_.LinkShapes = append(__GongStructShape__00000005_.LinkShapes, __LinkShape__00000003_)
	__GongStructShape__00000006_.LinkShapes = append(__GongStructShape__00000006_.LinkShapes, __LinkShape__00000011_)
	__GongStructShape__00000006_.LinkShapes = append(__GongStructShape__00000006_.LinkShapes, __LinkShape__00000012_)
	__GongStructShape__00000007_.LinkShapes = append(__GongStructShape__00000007_.LinkShapes, __LinkShape__00000010_)
	__GongStructShape__00000008_.LinkShapes = append(__GongStructShape__00000008_.LinkShapes, __LinkShape__00000013_)
	__GongStructShape__00000009_.LinkShapes = append(__GongStructShape__00000009_.LinkShapes, __LinkShape__00000014_)
}
