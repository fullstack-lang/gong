package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Concept__00000000_ := (&models.Concept{Name: `C1`}).Stage(stage)

	__ConceptShape__00000000_ := (&models.ConceptShape{Name: `C1-Default Diagram`}).Stage(stage)

	__Concern__00000000_ := (&models.Concern{Name: `Edit M0 & generates deliverables`}).Stage(stage)
	__Concern__00000001_ := (&models.Concern{Name: `Consumes SE deliverables`}).Stage(stage)
	__Concern__00000002_ := (&models.Concern{Name: `Edit M1 & develop tools to generate deliverables`}).Stage(stage)
	__Concern__00000003_ := (&models.Concern{Name: `Capture stakeholder map, concerns and SE deliverables`}).Stage(stage)

	__ConcernInputShape__00000000_ := (&models.ConcernInputShape{Name: `Consumes views to Views of the architecture`}).Stage(stage)
	__ConcernInputShape__00000001_ := (&models.ConcernInputShape{Name: `Edit M0 to Tools to edit M0 based on M1`}).Stage(stage)
	__ConcernInputShape__00000002_ := (&models.ConcernInputShape{Name: `Edit M1 & develop tools to edit M0 to Needed views`}).Stage(stage)

	__ConcernOutputShape__00000000_ := (&models.ConcernOutputShape{Name: `Edit views to Views of the architecture`}).Stage(stage)
	__ConcernOutputShape__00000001_ := (&models.ConcernOutputShape{Name: `Edit M1 to Tools to edit M0 based on M1`}).Stage(stage)
	__ConcernOutputShape__00000002_ := (&models.ConcernOutputShape{Name: `Edit SEMP that defines needed views (M0) for stakeholders to Needed views`}).Stage(stage)

	__ConcernShape__00000000_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000001_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000002_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000003_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)

	__Deliverable__00000000_ := (&models.Deliverable{Name: `SE Deliverables`}).Stage(stage)
	__Deliverable__00000002_ := (&models.Deliverable{Name: `Tools to edit M0 & generates deliverables`}).Stage(stage)
	__Deliverable__00000003_ := (&models.Deliverable{Name: `Requirements for SE deliverables + concepts`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `With SysML V2, the M1 editor is supposed to edit SysML v2 files for the defintion of the M1. The pain point with SysML v2 is:
- the high complexity required  for mastering SysML v2 syntax
- the lack of standard for the graphical views.`}).Stage(stage)

	__NoteProductShape__00000000_ := (&models.NoteProductShape{Name: `M0 is based on M1 to Tools to edit M0`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)

	__NoteTaskShape__00000002_ := (&models.NoteTaskShape{Name: `M0 is based on M1 to Edit M1 & develop tools to generate deliverables`}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000002_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000003_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)

	__Requirement__00000000_ := (&models.Requirement{Name: `R1`}).Stage(stage)

	__RequirementShape__00000000_ := (&models.RequirementShape{Name: `R1-Default Diagram`}).Stage(stage)

	__Stakeholder__00000000_ := (&models.Stakeholder{Name: `M0 editor`}).Stage(stage)
	__Stakeholder__00000001_ := (&models.Stakeholder{Name: `M1 editor`}).Stage(stage)
	__Stakeholder__00000002_ := (&models.Stakeholder{Name: `Stakeholder`}).Stage(stage)
	__Stakeholder__00000003_ := (&models.Stakeholder{Name: `Chief Engineer`}).Stage(stage)

	__StakeholderConcernShape__00000000_ := (&models.StakeholderConcernShape{Name: `M0 editor to Edit views`}).Stage(stage)
	__StakeholderConcernShape__00000003_ := (&models.StakeholderConcernShape{Name: `Stakeholder to Consumes views`}).Stage(stage)
	__StakeholderConcernShape__00000005_ := (&models.StakeholderConcernShape{Name: `M1 editor to Edit M1`}).Stage(stage)
	__StakeholderConcernShape__00000006_ := (&models.StakeholderConcernShape{Name: `Chief Engineer to Define needed views for stakeholders`}).Stage(stage)

	__StakeholderShape__00000000_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000001_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000002_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000003_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Concept__00000000_.Name = `C1`
	__Concept__00000000_.ComputedPrefix = ``
	__Concept__00000000_.IsExpanded = false
	__Concept__00000000_.LayoutDirection = models.Vertical

	__ConceptShape__00000000_.Name = `C1-Default Diagram`
	__ConceptShape__00000000_.IsExpanded = false
	__ConceptShape__00000000_.X = 879.146635
	__ConceptShape__00000000_.Y = 629.050761
	__ConceptShape__00000000_.Width = 250.000000
	__ConceptShape__00000000_.Height = 70.000000
	__ConceptShape__00000000_.IsHidden = false

	__Concern__00000000_.Name = `Edit M0 & generates deliverables`
	__Concern__00000000_.IDAirbus = ``
	__Concern__00000000_.Priority = ""
	__Concern__00000000_.ComputedPrefix = `1`
	__Concern__00000000_.IsExpanded = false
	__Concern__00000000_.LayoutDirection = models.Vertical
	__Concern__00000000_.Description = ``
	__Concern__00000000_.IsInputsNodeExpanded = false
	__Concern__00000000_.IsOutputsNodeExpanded = false
	__Concern__00000000_.IsWithCompletion = false
	__Concern__00000000_.Completion = ""

	__Concern__00000001_.Name = `Consumes SE deliverables`
	__Concern__00000001_.IDAirbus = ``
	__Concern__00000001_.Priority = ""
	__Concern__00000001_.ComputedPrefix = `2`
	__Concern__00000001_.IsExpanded = false
	__Concern__00000001_.LayoutDirection = models.Vertical
	__Concern__00000001_.Description = ``
	__Concern__00000001_.IsInputsNodeExpanded = false
	__Concern__00000001_.IsOutputsNodeExpanded = false
	__Concern__00000001_.IsWithCompletion = false
	__Concern__00000001_.Completion = ""

	__Concern__00000002_.Name = `Edit M1 & develop tools to generate deliverables`
	__Concern__00000002_.IDAirbus = ``
	__Concern__00000002_.Priority = ""
	__Concern__00000002_.ComputedPrefix = `3`
	__Concern__00000002_.IsExpanded = false
	__Concern__00000002_.LayoutDirection = models.Vertical
	__Concern__00000002_.Description = ``
	__Concern__00000002_.IsInputsNodeExpanded = false
	__Concern__00000002_.IsOutputsNodeExpanded = false
	__Concern__00000002_.IsWithCompletion = false
	__Concern__00000002_.Completion = ""

	__Concern__00000003_.Name = `Capture stakeholder map, concerns and SE deliverables`
	__Concern__00000003_.IDAirbus = ``
	__Concern__00000003_.Priority = ""
	__Concern__00000003_.ComputedPrefix = `4`
	__Concern__00000003_.IsExpanded = false
	__Concern__00000003_.LayoutDirection = models.Vertical
	__Concern__00000003_.Description = ``
	__Concern__00000003_.IsInputsNodeExpanded = false
	__Concern__00000003_.IsOutputsNodeExpanded = false
	__Concern__00000003_.IsWithCompletion = false
	__Concern__00000003_.Completion = ""

	__ConcernInputShape__00000000_.Name = `Consumes views to Views of the architecture`
	__ConcernInputShape__00000000_.StartRatio = 0.500000
	__ConcernInputShape__00000000_.EndRatio = 0.500000
	__ConcernInputShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernInputShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernInputShape__00000000_.CornerOffsetRatio = 1.680000
	__ConcernInputShape__00000000_.IsHidden = false

	__ConcernInputShape__00000001_.Name = `Edit M0 to Tools to edit M0 based on M1`
	__ConcernInputShape__00000001_.StartRatio = 0.500000
	__ConcernInputShape__00000001_.EndRatio = 0.500000
	__ConcernInputShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernInputShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernInputShape__00000001_.CornerOffsetRatio = 1.680000
	__ConcernInputShape__00000001_.IsHidden = false

	__ConcernInputShape__00000002_.Name = `Edit M1 & develop tools to edit M0 to Needed views`
	__ConcernInputShape__00000002_.StartRatio = 0.500000
	__ConcernInputShape__00000002_.EndRatio = 0.500000
	__ConcernInputShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernInputShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernInputShape__00000002_.CornerOffsetRatio = 1.680000
	__ConcernInputShape__00000002_.IsHidden = false

	__ConcernOutputShape__00000000_.Name = `Edit views to Views of the architecture`
	__ConcernOutputShape__00000000_.StartRatio = 0.500000
	__ConcernOutputShape__00000000_.EndRatio = 0.500000
	__ConcernOutputShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000000_.CornerOffsetRatio = 1.680000
	__ConcernOutputShape__00000000_.IsHidden = false

	__ConcernOutputShape__00000001_.Name = `Edit M1 to Tools to edit M0 based on M1`
	__ConcernOutputShape__00000001_.StartRatio = 0.500000
	__ConcernOutputShape__00000001_.EndRatio = 0.500000
	__ConcernOutputShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000001_.CornerOffsetRatio = 1.680000
	__ConcernOutputShape__00000001_.IsHidden = false

	__ConcernOutputShape__00000002_.Name = `Edit SEMP that defines needed views (M0) for stakeholders to Needed views`
	__ConcernOutputShape__00000002_.StartRatio = 0.500000
	__ConcernOutputShape__00000002_.EndRatio = 0.500000
	__ConcernOutputShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000002_.CornerOffsetRatio = 1.680000
	__ConcernOutputShape__00000002_.IsHidden = false

	__ConcernShape__00000000_.Name = `-Default Diagram`
	__ConcernShape__00000000_.IsExpanded = false
	__ConcernShape__00000000_.X = 489.172253
	__ConcernShape__00000000_.Y = 600.542541
	__ConcernShape__00000000_.Width = 250.000000
	__ConcernShape__00000000_.Height = 70.000000
	__ConcernShape__00000000_.IsHidden = false

	__ConcernShape__00000001_.Name = `-Default Diagram`
	__ConcernShape__00000001_.IsExpanded = false
	__ConcernShape__00000001_.X = 491.626829
	__ConcernShape__00000001_.Y = 899.843004
	__ConcernShape__00000001_.Width = 250.000000
	__ConcernShape__00000001_.Height = 70.000000
	__ConcernShape__00000001_.IsHidden = false

	__ConcernShape__00000002_.Name = `-Default Diagram`
	__ConcernShape__00000002_.IsExpanded = false
	__ConcernShape__00000002_.X = 495.507589
	__ConcernShape__00000002_.Y = 298.618265
	__ConcernShape__00000002_.Width = 250.000000
	__ConcernShape__00000002_.Height = 70.000000
	__ConcernShape__00000002_.IsHidden = false

	__ConcernShape__00000003_.Name = `-Default Diagram`
	__ConcernShape__00000003_.IsExpanded = false
	__ConcernShape__00000003_.X = 492.172173
	__ConcernShape__00000003_.Y = 46.539908
	__ConcernShape__00000003_.Width = 250.000000
	__ConcernShape__00000003_.Height = 70.000000
	__ConcernShape__00000003_.IsHidden = false

	__Deliverable__00000000_.Name = `SE Deliverables`
	__Deliverable__00000000_.ComputedPrefix = `1`
	__Deliverable__00000000_.IsExpanded = false
	__Deliverable__00000000_.LayoutDirection = models.Vertical
	__Deliverable__00000000_.Description = ``
	__Deliverable__00000000_.IsProducersNodeExpanded = false
	__Deliverable__00000000_.IsConsumersNodeExpanded = false

	__Deliverable__00000002_.Name = `Tools to edit M0 & generates deliverables`
	__Deliverable__00000002_.ComputedPrefix = `2`
	__Deliverable__00000002_.IsExpanded = false
	__Deliverable__00000002_.LayoutDirection = models.Vertical
	__Deliverable__00000002_.Description = ``
	__Deliverable__00000002_.IsProducersNodeExpanded = false
	__Deliverable__00000002_.IsConsumersNodeExpanded = false

	__Deliverable__00000003_.Name = `Requirements for SE deliverables + concepts`
	__Deliverable__00000003_.ComputedPrefix = `3`
	__Deliverable__00000003_.IsExpanded = false
	__Deliverable__00000003_.LayoutDirection = models.Vertical
	__Deliverable__00000003_.Description = ``
	__Deliverable__00000003_.IsProducersNodeExpanded = false
	__Deliverable__00000003_.IsConsumersNodeExpanded = false

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.LayoutDirection = models.Vertical
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 13900.000000
	__Diagram__00000000_.Height = 13900.000000
	__Diagram__00000000_.IsRequirementsNodeExpanded = true
	__Diagram__00000000_.IsConceptsNodeExpanded = true
	__Diagram__00000000_.IsPBSNodeExpanded = false
	__Diagram__00000000_.IsConcernsNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsStakeholdersNodeExpanded = false

	__Library__00000000_.Name = ``
	__Library__00000000_.IsRootLibrary = true
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.NbPixPerCharacter = 8.000000

	__Note__00000000_.Name = `With SysML V2, the M1 editor is supposed to edit SysML v2 files for the defintion of the M1. The pain point with SysML v2 is:
- the high complexity required  for mastering SysML v2 syntax
- the lack of standard for the graphical views.`
	__Note__00000000_.ComputedPrefix = `1`
	__Note__00000000_.IsExpanded = false
	__Note__00000000_.LayoutDirection = models.Vertical

	__NoteProductShape__00000000_.Name = `M0 is based on M1 to Tools to edit M0`
	__NoteProductShape__00000000_.StartRatio = 0.500000
	__NoteProductShape__00000000_.EndRatio = 0.500000
	__NoteProductShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000000_.IsHidden = false

	__NoteShape__00000000_.Name = `-Default Diagram`
	__NoteShape__00000000_.IsExpanded = false
	__NoteShape__00000000_.X = 821.586218
	__NoteShape__00000000_.Y = 287.916605
	__NoteShape__00000000_.Width = 398.000000
	__NoteShape__00000000_.Height = 222.000000
	__NoteShape__00000000_.IsHidden = false

	__NoteTaskShape__00000002_.Name = `M0 is based on M1 to Edit M1 & develop tools to generate deliverables`
	__NoteTaskShape__00000002_.StartRatio = 0.500000
	__NoteTaskShape__00000002_.EndRatio = 0.500000
	__NoteTaskShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000002_.CornerOffsetRatio = 1.680000
	__NoteTaskShape__00000002_.IsHidden = false

	__ProductShape__00000000_.Name = `-Default Diagram`
	__ProductShape__00000000_.IsExpanded = false
	__ProductShape__00000000_.X = 490.705042
	__ProductShape__00000000_.Y = 745.311484
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.IsHidden = false

	__ProductShape__00000002_.Name = `-Default Diagram`
	__ProductShape__00000002_.IsExpanded = false
	__ProductShape__00000002_.X = 489.953677
	__ProductShape__00000002_.Y = 444.394466
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000
	__ProductShape__00000002_.IsHidden = false

	__ProductShape__00000003_.Name = `-Default Diagram`
	__ProductShape__00000003_.IsExpanded = false
	__ProductShape__00000003_.X = 497.919156
	__ProductShape__00000003_.Y = 164.854495
	__ProductShape__00000003_.Width = 250.000000
	__ProductShape__00000003_.Height = 70.000000
	__ProductShape__00000003_.IsHidden = false

	__Requirement__00000000_.Name = `R1`
	__Requirement__00000000_.ComputedPrefix = ``
	__Requirement__00000000_.IsExpanded = false
	__Requirement__00000000_.LayoutDirection = models.Vertical

	__RequirementShape__00000000_.Name = `R1-Default Diagram`
	__RequirementShape__00000000_.IsExpanded = false
	__RequirementShape__00000000_.X = 868.311542
	__RequirementShape__00000000_.Y = 767.444900
	__RequirementShape__00000000_.Width = 250.000000
	__RequirementShape__00000000_.Height = 70.000000
	__RequirementShape__00000000_.IsHidden = false

	__Stakeholder__00000000_.Name = `M0 editor`
	__Stakeholder__00000000_.IDAirbus = ``
	__Stakeholder__00000000_.ComputedPrefix = `1`
	__Stakeholder__00000000_.IsExpanded = false
	__Stakeholder__00000000_.LayoutDirection = models.Vertical
	__Stakeholder__00000000_.Description = ``

	__Stakeholder__00000001_.Name = `M1 editor`
	__Stakeholder__00000001_.IDAirbus = ``
	__Stakeholder__00000001_.ComputedPrefix = `2`
	__Stakeholder__00000001_.IsExpanded = false
	__Stakeholder__00000001_.LayoutDirection = models.Vertical
	__Stakeholder__00000001_.Description = ``

	__Stakeholder__00000002_.Name = `Stakeholder`
	__Stakeholder__00000002_.IDAirbus = ``
	__Stakeholder__00000002_.ComputedPrefix = `3`
	__Stakeholder__00000002_.IsExpanded = false
	__Stakeholder__00000002_.LayoutDirection = models.Vertical
	__Stakeholder__00000002_.Description = ``

	__Stakeholder__00000003_.Name = `Chief Engineer`
	__Stakeholder__00000003_.IDAirbus = ``
	__Stakeholder__00000003_.ComputedPrefix = `4`
	__Stakeholder__00000003_.IsExpanded = false
	__Stakeholder__00000003_.LayoutDirection = models.Vertical
	__Stakeholder__00000003_.Description = ``

	__StakeholderConcernShape__00000000_.Name = `M0 editor to Edit views`
	__StakeholderConcernShape__00000000_.StartRatio = 0.500000
	__StakeholderConcernShape__00000000_.EndRatio = 0.500000
	__StakeholderConcernShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000000_.CornerOffsetRatio = 1.680000
	__StakeholderConcernShape__00000000_.IsHidden = false

	__StakeholderConcernShape__00000003_.Name = `Stakeholder to Consumes views`
	__StakeholderConcernShape__00000003_.StartRatio = 0.500000
	__StakeholderConcernShape__00000003_.EndRatio = 0.500000
	__StakeholderConcernShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000003_.CornerOffsetRatio = 1.680000
	__StakeholderConcernShape__00000003_.IsHidden = false

	__StakeholderConcernShape__00000005_.Name = `M1 editor to Edit M1`
	__StakeholderConcernShape__00000005_.StartRatio = 0.500000
	__StakeholderConcernShape__00000005_.EndRatio = 0.500000
	__StakeholderConcernShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000005_.CornerOffsetRatio = 1.680000
	__StakeholderConcernShape__00000005_.IsHidden = false

	__StakeholderConcernShape__00000006_.Name = `Chief Engineer to Define needed views for stakeholders`
	__StakeholderConcernShape__00000006_.StartRatio = 0.500000
	__StakeholderConcernShape__00000006_.EndRatio = 0.500000
	__StakeholderConcernShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000006_.CornerOffsetRatio = 1.680000
	__StakeholderConcernShape__00000006_.IsHidden = false

	__StakeholderShape__00000000_.Name = `-Default Diagram`
	__StakeholderShape__00000000_.IsExpanded = false
	__StakeholderShape__00000000_.X = 67.788415
	__StakeholderShape__00000000_.Y = 600.758204
	__StakeholderShape__00000000_.Width = 250.000000
	__StakeholderShape__00000000_.Height = 70.000000
	__StakeholderShape__00000000_.IsHidden = false

	__StakeholderShape__00000001_.Name = `-Default Diagram`
	__StakeholderShape__00000001_.IsExpanded = false
	__StakeholderShape__00000001_.X = 72.967288
	__StakeholderShape__00000001_.Y = 296.167318
	__StakeholderShape__00000001_.Width = 250.000000
	__StakeholderShape__00000001_.Height = 70.000000
	__StakeholderShape__00000001_.IsHidden = false

	__StakeholderShape__00000002_.Name = `-Default Diagram`
	__StakeholderShape__00000002_.IsExpanded = false
	__StakeholderShape__00000002_.X = 67.305489
	__StakeholderShape__00000002_.Y = 895.366676
	__StakeholderShape__00000002_.Width = 250.000000
	__StakeholderShape__00000002_.Height = 70.000000
	__StakeholderShape__00000002_.IsHidden = false

	__StakeholderShape__00000003_.Name = `-Default Diagram`
	__StakeholderShape__00000003_.IsExpanded = false
	__StakeholderShape__00000003_.X = 68.482242
	__StakeholderShape__00000003_.Y = 44.830683
	__StakeholderShape__00000003_.Width = 250.000000
	__StakeholderShape__00000003_.Height = 70.000000
	__StakeholderShape__00000003_.IsHidden = false

	// insertion point for setup of pointers
	__ConceptShape__00000000_.Concept = __Concept__00000000_
	__Concern__00000000_.Inputs = append(__Concern__00000000_.Inputs, __Deliverable__00000002_)
	__Concern__00000000_.Outputs = append(__Concern__00000000_.Outputs, __Deliverable__00000000_)
	__Concern__00000001_.Inputs = append(__Concern__00000001_.Inputs, __Deliverable__00000000_)
	__Concern__00000002_.Inputs = append(__Concern__00000002_.Inputs, __Deliverable__00000003_)
	__Concern__00000002_.Outputs = append(__Concern__00000002_.Outputs, __Deliverable__00000002_)
	__Concern__00000003_.Outputs = append(__Concern__00000003_.Outputs, __Deliverable__00000003_)
	__ConcernInputShape__00000000_.Deliverable = __Deliverable__00000000_
	__ConcernInputShape__00000000_.Concern = __Concern__00000001_
	__ConcernInputShape__00000001_.Deliverable = __Deliverable__00000002_
	__ConcernInputShape__00000001_.Concern = __Concern__00000000_
	__ConcernInputShape__00000002_.Deliverable = __Deliverable__00000003_
	__ConcernInputShape__00000002_.Concern = __Concern__00000002_
	__ConcernOutputShape__00000000_.Task = __Concern__00000000_
	__ConcernOutputShape__00000000_.Product = __Deliverable__00000000_
	__ConcernOutputShape__00000001_.Task = __Concern__00000002_
	__ConcernOutputShape__00000001_.Product = __Deliverable__00000002_
	__ConcernOutputShape__00000002_.Task = __Concern__00000003_
	__ConcernOutputShape__00000002_.Product = __Deliverable__00000003_
	__ConcernShape__00000000_.Concern = __Concern__00000000_
	__ConcernShape__00000001_.Concern = __Concern__00000001_
	__ConcernShape__00000002_.Concern = __Concern__00000002_
	__ConcernShape__00000003_.Concern = __Concern__00000003_
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000002_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000003_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000000_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000001_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000002_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000003_)
	__Diagram__00000000_.ConcernInputShapes = append(__Diagram__00000000_.ConcernInputShapes, __ConcernInputShape__00000000_)
	__Diagram__00000000_.ConcernInputShapes = append(__Diagram__00000000_.ConcernInputShapes, __ConcernInputShape__00000001_)
	__Diagram__00000000_.ConcernInputShapes = append(__Diagram__00000000_.ConcernInputShapes, __ConcernInputShape__00000002_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000000_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000001_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000002_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Diagram__00000000_.NoteProductShapes = append(__Diagram__00000000_.NoteProductShapes, __NoteProductShape__00000000_)
	__Diagram__00000000_.NoteTaskShapes = append(__Diagram__00000000_.NoteTaskShapes, __NoteTaskShape__00000002_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000000_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000001_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000002_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000003_)
	__Diagram__00000000_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000000_.ResourcesWhoseNodeIsExpanded, __Stakeholder__00000002_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000000_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000003_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000005_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000006_)
	__Diagram__00000000_.Requirement_Shapes = append(__Diagram__00000000_.Requirement_Shapes, __RequirementShape__00000000_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000000_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000000_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000002_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000003_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000000_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000001_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000002_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000003_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000000_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000001_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000002_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000003_)
	__Library__00000000_.RootRequirements = append(__Library__00000000_.RootRequirements, __Requirement__00000000_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000000_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Note__00000000_.Products = append(__Note__00000000_.Products, __Deliverable__00000002_)
	__Note__00000000_.Tasks = append(__Note__00000000_.Tasks, __Concern__00000002_)
	__NoteProductShape__00000000_.Note = __Note__00000000_
	__NoteProductShape__00000000_.Product = __Deliverable__00000002_
	__NoteShape__00000000_.Note = __Note__00000000_
	__NoteTaskShape__00000002_.Note = __Note__00000000_
	__NoteTaskShape__00000002_.Task = __Concern__00000002_
	__ProductShape__00000000_.Product = __Deliverable__00000000_
	__ProductShape__00000002_.Product = __Deliverable__00000002_
	__ProductShape__00000003_.Product = __Deliverable__00000003_
	__RequirementShape__00000000_.Requirement = __Requirement__00000000_
	__Stakeholder__00000000_.Concerns = append(__Stakeholder__00000000_.Concerns, __Concern__00000000_)
	__Stakeholder__00000001_.Concerns = append(__Stakeholder__00000001_.Concerns, __Concern__00000002_)
	__Stakeholder__00000002_.Concerns = append(__Stakeholder__00000002_.Concerns, __Concern__00000001_)
	__Stakeholder__00000003_.Concerns = append(__Stakeholder__00000003_.Concerns, __Concern__00000003_)
	__StakeholderConcernShape__00000000_.Stakeholder = __Stakeholder__00000000_
	__StakeholderConcernShape__00000000_.Concern = __Concern__00000000_
	__StakeholderConcernShape__00000003_.Stakeholder = __Stakeholder__00000002_
	__StakeholderConcernShape__00000003_.Concern = __Concern__00000001_
	__StakeholderConcernShape__00000005_.Stakeholder = __Stakeholder__00000001_
	__StakeholderConcernShape__00000005_.Concern = __Concern__00000002_
	__StakeholderConcernShape__00000006_.Stakeholder = __Stakeholder__00000003_
	__StakeholderConcernShape__00000006_.Concern = __Concern__00000003_
	__StakeholderShape__00000000_.Stakeholder = __Stakeholder__00000000_
	__StakeholderShape__00000001_.Stakeholder = __Stakeholder__00000001_
	__StakeholderShape__00000002_.Stakeholder = __Stakeholder__00000002_
	__StakeholderShape__00000003_.Stakeholder = __Stakeholder__00000003_
}
