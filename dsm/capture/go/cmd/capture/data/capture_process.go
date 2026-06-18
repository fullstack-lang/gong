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

	__Concept__00000000_ := (&models.Concept{Name: `Concept`}).Stage(stage)
	__Concept__00000001_ := (&models.Concept{Name: `Deliverable`}).Stage(stage)
	__Concept__00000002_ := (&models.Concept{Name: `Stakeholder`}).Stage(stage)
	__Concept__00000003_ := (&models.Concept{Name: `Concern`}).Stage(stage)

	__ConceptShape__00000000_ := (&models.ConceptShape{Name: `C1-Default Diagram`}).Stage(stage)
	__ConceptShape__00000001_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000002_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000003_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)

	__Concern__00000000_ := (&models.Concern{Name: `Edit M0 & generates deliverables`}).Stage(stage)
	__Concern__00000001_ := (&models.Concern{Name: `Consumes SE deliverables`}).Stage(stage)
	__Concern__00000002_ := (&models.Concern{Name: `Edit M1 & develop tools to generate deliverables`}).Stage(stage)
	__Concern__00000003_ := (&models.Concern{Name: `Edit SEMP`}).Stage(stage)

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
	__Deliverable__00000003_ := (&models.Deliverable{Name: `System Engineering Management Plan`}).Stage(stage)

	__DeliverableConceptShape__00000004_ := (&models.DeliverableConceptShape{Name: `Requirements for SE deliverables + concepts to C1`}).Stage(stage)
	__DeliverableConceptShape__00000005_ := (&models.DeliverableConceptShape{Name: `Requirements for SE deliverables + concepts to Deliverable`}).Stage(stage)
	__DeliverableConceptShape__00000006_ := (&models.DeliverableConceptShape{Name: `Requirements for SE deliverables + concepts to Stakeholder`}).Stage(stage)
	__DeliverableConceptShape__00000007_ := (&models.DeliverableConceptShape{Name: `Requirements for SE deliverables + concepts to Concern`}).Stage(stage)

	__DeliverableShape__00000000_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000002_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000003_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `With SysML V2, the M1 editor is supposed to edit SysML v2 files for the defintion of the M1. The pain point with SysML v2 is:
- the high complexity required  for mastering SysML v2 syntax
- the lack of standard for the graphical views.`}).Stage(stage)
	__Note__00000001_ := (&models.Note{Name: `System Engineering Management Plan`}).Stage(stage)

	__NoteDeliverableShape__00000000_ := (&models.NoteDeliverableShape{Name: `M0 is based on M1 to Tools to edit M0`}).Stage(stage)
	__NoteDeliverableShape__00000001_ := (&models.NoteDeliverableShape{Name: `System Engineering Management Plan to SEMP`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)
	__NoteShape__00000001_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)

	__NoteTaskShape__00000002_ := (&models.NoteTaskShape{Name: `M0 is based on M1 to Edit M1 & develop tools to generate deliverables`}).Stage(stage)
	__NoteTaskShape__00000003_ := (&models.NoteTaskShape{Name: `System Engineering Management Plan to Edit SEMP`}).Stage(stage)

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

	__Concept__00000000_.Name = `Concept`
	__Concept__00000000_.ComputedPrefix = ``
	__Concept__00000000_.IsExpanded = false
	__Concept__00000000_.LayoutDirection = models.Vertical

	__Concept__00000001_.Name = `Deliverable`
	__Concept__00000001_.ComputedPrefix = ``
	__Concept__00000001_.IsExpanded = false
	__Concept__00000001_.LayoutDirection = models.Vertical

	__Concept__00000002_.Name = `Stakeholder`
	__Concept__00000002_.ComputedPrefix = ``
	__Concept__00000002_.IsExpanded = false
	__Concept__00000002_.LayoutDirection = models.Vertical

	__Concept__00000003_.Name = `Concern`
	__Concept__00000003_.ComputedPrefix = ``
	__Concept__00000003_.IsExpanded = false
	__Concept__00000003_.LayoutDirection = models.Vertical

	__ConceptShape__00000000_.Name = `C1-Default Diagram`
	__ConceptShape__00000000_.IsExpanded = false
	__ConceptShape__00000000_.X = 908.146635
	__ConceptShape__00000000_.Y = 300.050745
	__ConceptShape__00000000_.Width = 250.000000
	__ConceptShape__00000000_.Height = 70.000000
	__ConceptShape__00000000_.IsHidden = false

	__ConceptShape__00000001_.Name = `-Default Diagram`
	__ConceptShape__00000001_.IsExpanded = false
	__ConceptShape__00000001_.X = 908.268579
	__ConceptShape__00000001_.Y = 205.344947
	__ConceptShape__00000001_.Width = 250.000000
	__ConceptShape__00000001_.Height = 70.000000
	__ConceptShape__00000001_.IsHidden = false

	__ConceptShape__00000002_.Name = `-Default Diagram`
	__ConceptShape__00000002_.IsExpanded = false
	__ConceptShape__00000002_.X = 908.674336
	__ConceptShape__00000002_.Y = 25.778790
	__ConceptShape__00000002_.Width = 250.000000
	__ConceptShape__00000002_.Height = 70.000000
	__ConceptShape__00000002_.IsHidden = false

	__ConceptShape__00000003_.Name = `-Default Diagram`
	__ConceptShape__00000003_.IsExpanded = false
	__ConceptShape__00000003_.X = 911.783111
	__ConceptShape__00000003_.Y = 117.790156
	__ConceptShape__00000003_.Width = 250.000000
	__ConceptShape__00000003_.Height = 70.000000
	__ConceptShape__00000003_.IsHidden = false

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

	__Concern__00000003_.Name = `Edit SEMP`
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
	__ConcernInputShape__00000002_.StartRatio = 0.224448
	__ConcernInputShape__00000002_.EndRatio = 0.500000
	__ConcernInputShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernInputShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernInputShape__00000002_.CornerOffsetRatio = 1.335412
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
	__ConcernOutputShape__00000002_.StartRatio = 0.387436
	__ConcernOutputShape__00000002_.EndRatio = 0.500000
	__ConcernOutputShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000002_.CornerOffsetRatio = 1.397049
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
	__ConcernShape__00000002_.X = 489.507589
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

	__Deliverable__00000003_.Name = `System Engineering Management Plan`
	__Deliverable__00000003_.ComputedPrefix = `3`
	__Deliverable__00000003_.IsExpanded = false
	__Deliverable__00000003_.LayoutDirection = models.Vertical
	__Deliverable__00000003_.Description = ``
	__Deliverable__00000003_.IsProducersNodeExpanded = false
	__Deliverable__00000003_.IsConsumersNodeExpanded = false

	__DeliverableConceptShape__00000004_.Name = `Requirements for SE deliverables + concepts to C1`
	__DeliverableConceptShape__00000004_.StartRatio = 0.500000
	__DeliverableConceptShape__00000004_.EndRatio = 0.500000
	__DeliverableConceptShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000004_.CornerOffsetRatio = 1.680000
	__DeliverableConceptShape__00000004_.IsHidden = false

	__DeliverableConceptShape__00000005_.Name = `Requirements for SE deliverables + concepts to Deliverable`
	__DeliverableConceptShape__00000005_.StartRatio = 0.500000
	__DeliverableConceptShape__00000005_.EndRatio = 0.500000
	__DeliverableConceptShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000005_.CornerOffsetRatio = 1.680000
	__DeliverableConceptShape__00000005_.IsHidden = false

	__DeliverableConceptShape__00000006_.Name = `Requirements for SE deliverables + concepts to Stakeholder`
	__DeliverableConceptShape__00000006_.StartRatio = 0.500000
	__DeliverableConceptShape__00000006_.EndRatio = 0.500000
	__DeliverableConceptShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000006_.CornerOffsetRatio = 1.680000
	__DeliverableConceptShape__00000006_.IsHidden = false

	__DeliverableConceptShape__00000007_.Name = `Requirements for SE deliverables + concepts to Concern`
	__DeliverableConceptShape__00000007_.StartRatio = 0.500000
	__DeliverableConceptShape__00000007_.EndRatio = 0.500000
	__DeliverableConceptShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000007_.CornerOffsetRatio = 1.680000
	__DeliverableConceptShape__00000007_.IsHidden = false

	__DeliverableShape__00000000_.Name = `-Default Diagram`
	__DeliverableShape__00000000_.IsExpanded = false
	__DeliverableShape__00000000_.X = 490.705042
	__DeliverableShape__00000000_.Y = 745.311484
	__DeliverableShape__00000000_.Width = 250.000000
	__DeliverableShape__00000000_.Height = 70.000000
	__DeliverableShape__00000000_.IsHidden = false

	__DeliverableShape__00000002_.Name = `-Default Diagram`
	__DeliverableShape__00000002_.IsExpanded = false
	__DeliverableShape__00000002_.X = 489.953677
	__DeliverableShape__00000002_.Y = 444.394466
	__DeliverableShape__00000002_.Width = 250.000000
	__DeliverableShape__00000002_.Height = 70.000000
	__DeliverableShape__00000002_.IsHidden = false

	__DeliverableShape__00000003_.Name = `-Default Diagram`
	__DeliverableShape__00000003_.IsExpanded = false
	__DeliverableShape__00000003_.X = 489.919156
	__DeliverableShape__00000003_.Y = 171.854495
	__DeliverableShape__00000003_.Width = 250.000000
	__DeliverableShape__00000003_.Height = 70.000000
	__DeliverableShape__00000003_.IsHidden = false

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.LayoutDirection = models.Vertical
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 22400.000000
	__Diagram__00000000_.Height = 22400.000000
	__Diagram__00000000_.IsRequirementsNodeExpanded = true
	__Diagram__00000000_.IsConceptsNodeExpanded = true
	__Diagram__00000000_.IsPBSNodeExpanded = true
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

	__Note__00000001_.Name = `System Engineering Management Plan`
	__Note__00000001_.ComputedPrefix = `2`
	__Note__00000001_.IsExpanded = false
	__Note__00000001_.LayoutDirection = models.Vertical

	__NoteDeliverableShape__00000000_.Name = `M0 is based on M1 to Tools to edit M0`
	__NoteDeliverableShape__00000000_.StartRatio = 0.500000
	__NoteDeliverableShape__00000000_.EndRatio = 0.500000
	__NoteDeliverableShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteDeliverableShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteDeliverableShape__00000000_.CornerOffsetRatio = 1.680000
	__NoteDeliverableShape__00000000_.IsHidden = false

	__NoteDeliverableShape__00000001_.Name = `System Engineering Management Plan to SEMP`
	__NoteDeliverableShape__00000001_.StartRatio = 0.500000
	__NoteDeliverableShape__00000001_.EndRatio = 0.500000
	__NoteDeliverableShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteDeliverableShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteDeliverableShape__00000001_.CornerOffsetRatio = 1.680000
	__NoteDeliverableShape__00000001_.IsHidden = false

	__NoteShape__00000000_.Name = `-Default Diagram`
	__NoteShape__00000000_.IsExpanded = false
	__NoteShape__00000000_.X = 818.586218
	__NoteShape__00000000_.Y = 408.916605
	__NoteShape__00000000_.Width = 466.000000
	__NoteShape__00000000_.Height = 222.000000
	__NoteShape__00000000_.IsHidden = false

	__NoteShape__00000001_.Name = `-Default Diagram`
	__NoteShape__00000001_.IsExpanded = false
	__NoteShape__00000001_.X = 74.858056
	__NoteShape__00000001_.Y = 173.233449
	__NoteShape__00000001_.Width = 250.000000
	__NoteShape__00000001_.Height = 70.000000
	__NoteShape__00000001_.IsHidden = true

	__NoteTaskShape__00000002_.Name = `M0 is based on M1 to Edit M1 & develop tools to generate deliverables`
	__NoteTaskShape__00000002_.StartRatio = 0.500000
	__NoteTaskShape__00000002_.EndRatio = 0.500000
	__NoteTaskShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000002_.CornerOffsetRatio = 1.680000
	__NoteTaskShape__00000002_.IsHidden = false

	__NoteTaskShape__00000003_.Name = `System Engineering Management Plan to Edit SEMP`
	__NoteTaskShape__00000003_.StartRatio = 0.500000
	__NoteTaskShape__00000003_.EndRatio = 0.500000
	__NoteTaskShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000003_.CornerOffsetRatio = 1.680000
	__NoteTaskShape__00000003_.IsHidden = false

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
	__StakeholderShape__00000000_.Y = 596.758204
	__StakeholderShape__00000000_.Width = 250.000000
	__StakeholderShape__00000000_.Height = 70.000000
	__StakeholderShape__00000000_.IsHidden = false

	__StakeholderShape__00000001_.Name = `-Default Diagram`
	__StakeholderShape__00000001_.IsExpanded = false
	__StakeholderShape__00000001_.X = 62.967288
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
	__ConceptShape__00000001_.Concept = __Concept__00000001_
	__ConceptShape__00000002_.Concept = __Concept__00000002_
	__ConceptShape__00000003_.Concept = __Concept__00000003_
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
	__ConcernOutputShape__00000000_.Concern = __Concern__00000000_
	__ConcernOutputShape__00000000_.Deliverable = __Deliverable__00000000_
	__ConcernOutputShape__00000001_.Concern = __Concern__00000002_
	__ConcernOutputShape__00000001_.Deliverable = __Deliverable__00000002_
	__ConcernOutputShape__00000002_.Concern = __Concern__00000003_
	__ConcernOutputShape__00000002_.Deliverable = __Deliverable__00000003_
	__ConcernShape__00000000_.Concern = __Concern__00000000_
	__ConcernShape__00000001_.Concern = __Concern__00000001_
	__ConcernShape__00000002_.Concern = __Concern__00000002_
	__ConcernShape__00000003_.Concern = __Concern__00000003_
	__Deliverable__00000003_.Concepts = append(__Deliverable__00000003_.Concepts, __Concept__00000000_)
	__Deliverable__00000003_.Concepts = append(__Deliverable__00000003_.Concepts, __Concept__00000001_)
	__Deliverable__00000003_.Concepts = append(__Deliverable__00000003_.Concepts, __Concept__00000002_)
	__Deliverable__00000003_.Concepts = append(__Deliverable__00000003_.Concepts, __Concept__00000003_)
	__DeliverableConceptShape__00000004_.Deliverable = __Deliverable__00000003_
	__DeliverableConceptShape__00000004_.Concept = __Concept__00000000_
	__DeliverableConceptShape__00000005_.Deliverable = __Deliverable__00000003_
	__DeliverableConceptShape__00000005_.Concept = __Concept__00000001_
	__DeliverableConceptShape__00000006_.Deliverable = __Deliverable__00000003_
	__DeliverableConceptShape__00000006_.Concept = __Concept__00000002_
	__DeliverableConceptShape__00000007_.Deliverable = __Deliverable__00000003_
	__DeliverableConceptShape__00000007_.Concept = __Concept__00000003_
	__DeliverableShape__00000000_.Deliverable = __Deliverable__00000000_
	__DeliverableShape__00000002_.Deliverable = __Deliverable__00000002_
	__DeliverableShape__00000003_.Deliverable = __Deliverable__00000003_
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000000_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000002_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000003_)
	__Diagram__00000000_.DeliverablesWhoseNodeIsExpanded = append(__Diagram__00000000_.DeliverablesWhoseNodeIsExpanded, __Deliverable__00000000_)
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
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000001_)
	__Diagram__00000000_.NoteDeliverableShapes = append(__Diagram__00000000_.NoteDeliverableShapes, __NoteDeliverableShape__00000000_)
	__Diagram__00000000_.NoteDeliverableShapes = append(__Diagram__00000000_.NoteDeliverableShapes, __NoteDeliverableShape__00000001_)
	__Diagram__00000000_.NoteTaskShapes = append(__Diagram__00000000_.NoteTaskShapes, __NoteTaskShape__00000002_)
	__Diagram__00000000_.NoteTaskShapes = append(__Diagram__00000000_.NoteTaskShapes, __NoteTaskShape__00000003_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000000_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000001_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000002_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000003_)
	__Diagram__00000000_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000000_.ResourcesWhoseNodeIsExpanded, __Stakeholder__00000002_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000000_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000003_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000005_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000006_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000000_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000001_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000002_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000003_)
	__Diagram__00000000_.ConceptsWhoseNodeIsExpanded = append(__Diagram__00000000_.ConceptsWhoseNodeIsExpanded, __Concept__00000000_)
	__Diagram__00000000_.ConceptsWhoseDeliverablesNodeIsExpanded = append(__Diagram__00000000_.ConceptsWhoseDeliverablesNodeIsExpanded, __Concept__00000000_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000004_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000005_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000006_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000007_)
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
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000000_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000001_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000002_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000003_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000000_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000001_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Note__00000000_.Deliverables = append(__Note__00000000_.Deliverables, __Deliverable__00000002_)
	__Note__00000000_.Tasks = append(__Note__00000000_.Tasks, __Concern__00000002_)
	__Note__00000001_.Deliverables = append(__Note__00000001_.Deliverables, __Deliverable__00000003_)
	__Note__00000001_.Tasks = append(__Note__00000001_.Tasks, __Concern__00000003_)
	__NoteDeliverableShape__00000000_.Note = __Note__00000000_
	__NoteDeliverableShape__00000000_.Deliverable = __Deliverable__00000002_
	__NoteDeliverableShape__00000001_.Note = __Note__00000001_
	__NoteDeliverableShape__00000001_.Deliverable = __Deliverable__00000003_
	__NoteShape__00000000_.Note = __Note__00000000_
	__NoteShape__00000001_.Note = __Note__00000001_
	__NoteTaskShape__00000002_.Note = __Note__00000000_
	__NoteTaskShape__00000002_.Task = __Concern__00000002_
	__NoteTaskShape__00000003_.Note = __Note__00000001_
	__NoteTaskShape__00000003_.Task = __Concern__00000003_
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
