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

	__Concern__00000000_ := (&models.Concern{Name: `Understand Environment & Mission`}).Stage(stage)
	__Concern__00000001_ := (&models.Concern{Name: `Map Stakeholders`}).Stage(stage)

	__ConcernOutputShape__00000000_ := (&models.ConcernOutputShape{Name: `Understand SoI to Gap Analysis`}).Stage(stage)
	__ConcernOutputShape__00000001_ := (&models.ConcernOutputShape{Name: `Understand SoI to OpsCon`}).Stage(stage)
	__ConcernOutputShape__00000002_ := (&models.ConcernOutputShape{Name: `Understand Environment & Mission to Mission Requirements`}).Stage(stage)
	__ConcernOutputShape__00000003_ := (&models.ConcernOutputShape{Name: `Map Stakeholders to Stakeholder Map`}).Stage(stage)

	__ConcernShape__00000000_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000001_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)

	__Deliverable__00000000_ := (&models.Deliverable{Name: `Stakeholder Map`}).Stage(stage)
	__Deliverable__00000001_ := (&models.Deliverable{Name: `OpsCon`}).Stage(stage)
	__Deliverable__00000003_ := (&models.Deliverable{Name: `Mission Requirements`}).Stage(stage)
	__Deliverable__00000004_ := (&models.Deliverable{Name: `Gap Analysis`}).Stage(stage)

	__DeliverableShape__00000000_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000001_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000003_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000004_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Stakeholder__00000000_ := (&models.Stakeholder{Name: `Mission Analyst`}).Stage(stage)
	__Stakeholder__00000001_ := (&models.Stakeholder{Name: `Enterprise Architect`}).Stage(stage)
	__Stakeholder__00000003_ := (&models.Stakeholder{Name: `Lead System Engineer`}).Stage(stage)
	__Stakeholder__00000004_ := (&models.Stakeholder{Name: `Executive Sponsor`}).Stage(stage)

	__StakeholderConcernShape__00000000_ := (&models.StakeholderConcernShape{Name: `Mission Analyst to Understand SoI`}).Stage(stage)
	__StakeholderConcernShape__00000001_ := (&models.StakeholderConcernShape{Name: `Enterprise Architect to Map Stakeholders`}).Stage(stage)
	__StakeholderConcernShape__00000002_ := (&models.StakeholderConcernShape{Name: `Lead System Engineer to Map Stakeholders`}).Stage(stage)

	__StakeholderShape__00000000_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000001_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000003_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000004_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Concern__00000000_.Name = `Understand Environment & Mission`
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

	__Concern__00000001_.Name = `Map Stakeholders`
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

	__ConcernOutputShape__00000000_.Name = `Understand SoI to Gap Analysis`
	__ConcernOutputShape__00000000_.StartRatio = 0.533597
	__ConcernOutputShape__00000000_.EndRatio = 0.680157
	__ConcernOutputShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000000_.CornerOffsetRatio = 1.326375
	__ConcernOutputShape__00000000_.IsHidden = false

	__ConcernOutputShape__00000001_.Name = `Understand SoI to OpsCon`
	__ConcernOutputShape__00000001_.StartRatio = 0.500000
	__ConcernOutputShape__00000001_.EndRatio = 0.327610
	__ConcernOutputShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000001_.CornerOffsetRatio = 1.232600
	__ConcernOutputShape__00000001_.IsHidden = false

	__ConcernOutputShape__00000002_.Name = `Understand Environment & Mission to Mission Requirements`
	__ConcernOutputShape__00000002_.StartRatio = 0.500000
	__ConcernOutputShape__00000002_.EndRatio = 0.484537
	__ConcernOutputShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000002_.CornerOffsetRatio = 1.223888
	__ConcernOutputShape__00000002_.IsHidden = false

	__ConcernOutputShape__00000003_.Name = `Map Stakeholders to Stakeholder Map`
	__ConcernOutputShape__00000003_.StartRatio = 0.693119
	__ConcernOutputShape__00000003_.EndRatio = 0.648312
	__ConcernOutputShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000003_.CornerOffsetRatio = 1.404601
	__ConcernOutputShape__00000003_.IsHidden = false

	__ConcernShape__00000000_.Name = `-Default Diagram`
	__ConcernShape__00000000_.IsExpanded = false
	__ConcernShape__00000000_.X = 440.510422
	__ConcernShape__00000000_.Y = 72.981540
	__ConcernShape__00000000_.Width = 250.000000
	__ConcernShape__00000000_.Height = 70.000000
	__ConcernShape__00000000_.IsHidden = false

	__ConcernShape__00000001_.Name = `-Default Diagram`
	__ConcernShape__00000001_.IsExpanded = false
	__ConcernShape__00000001_.X = 453.120458
	__ConcernShape__00000001_.Y = 347.814971
	__ConcernShape__00000001_.Width = 250.000000
	__ConcernShape__00000001_.Height = 70.000000
	__ConcernShape__00000001_.IsHidden = false

	__Deliverable__00000000_.Name = `Stakeholder Map`
	__Deliverable__00000000_.ComputedPrefix = `1`
	__Deliverable__00000000_.IsExpanded = false
	__Deliverable__00000000_.LayoutDirection = models.Vertical
	__Deliverable__00000000_.Description = ``
	__Deliverable__00000000_.IsProducersNodeExpanded = false
	__Deliverable__00000000_.IsConsumersNodeExpanded = false

	__Deliverable__00000001_.Name = `OpsCon`
	__Deliverable__00000001_.ComputedPrefix = `2`
	__Deliverable__00000001_.IsExpanded = false
	__Deliverable__00000001_.LayoutDirection = models.Vertical
	__Deliverable__00000001_.Description = ``
	__Deliverable__00000001_.IsProducersNodeExpanded = false
	__Deliverable__00000001_.IsConsumersNodeExpanded = false

	__Deliverable__00000003_.Name = `Mission Requirements`
	__Deliverable__00000003_.ComputedPrefix = `3`
	__Deliverable__00000003_.IsExpanded = false
	__Deliverable__00000003_.LayoutDirection = models.Vertical
	__Deliverable__00000003_.Description = ``
	__Deliverable__00000003_.IsProducersNodeExpanded = false
	__Deliverable__00000003_.IsConsumersNodeExpanded = false

	__Deliverable__00000004_.Name = `Gap Analysis`
	__Deliverable__00000004_.ComputedPrefix = `4`
	__Deliverable__00000004_.IsExpanded = false
	__Deliverable__00000004_.LayoutDirection = models.Vertical
	__Deliverable__00000004_.Description = ``
	__Deliverable__00000004_.IsProducersNodeExpanded = false
	__Deliverable__00000004_.IsConsumersNodeExpanded = false

	__DeliverableShape__00000000_.Name = `-Default Diagram`
	__DeliverableShape__00000000_.IsExpanded = false
	__DeliverableShape__00000000_.X = 841.752476
	__DeliverableShape__00000000_.Y = 343.951458
	__DeliverableShape__00000000_.Width = 250.000000
	__DeliverableShape__00000000_.Height = 70.000000
	__DeliverableShape__00000000_.IsHidden = false

	__DeliverableShape__00000001_.Name = `-Default Diagram`
	__DeliverableShape__00000001_.IsExpanded = false
	__DeliverableShape__00000001_.X = 848.660473
	__DeliverableShape__00000001_.Y = 158.400659
	__DeliverableShape__00000001_.Width = 250.000000
	__DeliverableShape__00000001_.Height = 70.000000
	__DeliverableShape__00000001_.IsHidden = false

	__DeliverableShape__00000003_.Name = `-Default Diagram`
	__DeliverableShape__00000003_.IsExpanded = false
	__DeliverableShape__00000003_.X = 846.482364
	__DeliverableShape__00000003_.Y = 246.415749
	__DeliverableShape__00000003_.Width = 250.000000
	__DeliverableShape__00000003_.Height = 70.000000
	__DeliverableShape__00000003_.IsHidden = false

	__DeliverableShape__00000004_.Name = `-Default Diagram`
	__DeliverableShape__00000004_.IsExpanded = false
	__DeliverableShape__00000004_.X = 844.884923
	__DeliverableShape__00000004_.Y = 65.722315
	__DeliverableShape__00000004_.Width = 250.000000
	__DeliverableShape__00000004_.Height = 70.000000
	__DeliverableShape__00000004_.IsHidden = false

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.LayoutDirection = models.Vertical
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 7017.436266
	__Diagram__00000000_.Height = 7000.000000
	__Diagram__00000000_.IsRequirementsNodeExpanded = false
	__Diagram__00000000_.IsConceptsNodeExpanded = false
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsConcernsNodeExpanded = true
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsStakeholdersNodeExpanded = true

	__Library__00000000_.Name = ``
	__Library__00000000_.IsRootLibrary = true
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.NbPixPerCharacter = 8.000000

	__Stakeholder__00000000_.Name = `Mission Analyst`
	__Stakeholder__00000000_.IDAirbus = ``
	__Stakeholder__00000000_.ComputedPrefix = `1`
	__Stakeholder__00000000_.IsExpanded = false
	__Stakeholder__00000000_.LayoutDirection = models.Vertical
	__Stakeholder__00000000_.Description = ``

	__Stakeholder__00000001_.Name = `Enterprise Architect`
	__Stakeholder__00000001_.IDAirbus = ``
	__Stakeholder__00000001_.ComputedPrefix = `2`
	__Stakeholder__00000001_.IsExpanded = false
	__Stakeholder__00000001_.LayoutDirection = models.Vertical
	__Stakeholder__00000001_.Description = ``

	__Stakeholder__00000003_.Name = `Lead System Engineer`
	__Stakeholder__00000003_.IDAirbus = ``
	__Stakeholder__00000003_.ComputedPrefix = `3`
	__Stakeholder__00000003_.IsExpanded = false
	__Stakeholder__00000003_.LayoutDirection = models.Vertical
	__Stakeholder__00000003_.Description = ``

	__Stakeholder__00000004_.Name = `Executive Sponsor`
	__Stakeholder__00000004_.IDAirbus = ``
	__Stakeholder__00000004_.ComputedPrefix = `4`
	__Stakeholder__00000004_.IsExpanded = false
	__Stakeholder__00000004_.LayoutDirection = models.Vertical
	__Stakeholder__00000004_.Description = ``

	__StakeholderConcernShape__00000000_.Name = `Mission Analyst to Understand SoI`
	__StakeholderConcernShape__00000000_.StartRatio = 0.500000
	__StakeholderConcernShape__00000000_.EndRatio = 0.500000
	__StakeholderConcernShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000000_.CornerOffsetRatio = 1.680000
	__StakeholderConcernShape__00000000_.IsHidden = false

	__StakeholderConcernShape__00000001_.Name = `Enterprise Architect to Map Stakeholders`
	__StakeholderConcernShape__00000001_.StartRatio = 0.500000
	__StakeholderConcernShape__00000001_.EndRatio = 0.500000
	__StakeholderConcernShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000001_.CornerOffsetRatio = 1.680000
	__StakeholderConcernShape__00000001_.IsHidden = false

	__StakeholderConcernShape__00000002_.Name = `Lead System Engineer to Map Stakeholders`
	__StakeholderConcernShape__00000002_.StartRatio = 0.500000
	__StakeholderConcernShape__00000002_.EndRatio = 0.500000
	__StakeholderConcernShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000002_.CornerOffsetRatio = 1.680000
	__StakeholderConcernShape__00000002_.IsHidden = false

	__StakeholderShape__00000000_.Name = `-Default Diagram`
	__StakeholderShape__00000000_.IsExpanded = false
	__StakeholderShape__00000000_.X = 46.436266
	__StakeholderShape__00000000_.Y = 69.775834
	__StakeholderShape__00000000_.Width = 250.000000
	__StakeholderShape__00000000_.Height = 70.000000
	__StakeholderShape__00000000_.IsHidden = false

	__StakeholderShape__00000001_.Name = `-Default Diagram`
	__StakeholderShape__00000001_.IsExpanded = false
	__StakeholderShape__00000001_.X = 40.483897
	__StakeholderShape__00000001_.Y = 291.208105
	__StakeholderShape__00000001_.Width = 250.000000
	__StakeholderShape__00000001_.Height = 70.000000
	__StakeholderShape__00000001_.IsHidden = false

	__StakeholderShape__00000003_.Name = `-Default Diagram`
	__StakeholderShape__00000003_.IsExpanded = false
	__StakeholderShape__00000003_.X = 38.609526
	__StakeholderShape__00000003_.Y = 381.285939
	__StakeholderShape__00000003_.Width = 250.000000
	__StakeholderShape__00000003_.Height = 70.000000
	__StakeholderShape__00000003_.IsHidden = false

	__StakeholderShape__00000004_.Name = `-Default Diagram`
	__StakeholderShape__00000004_.IsExpanded = false
	__StakeholderShape__00000004_.X = 45.616112
	__StakeholderShape__00000004_.Y = 541.956383
	__StakeholderShape__00000004_.Width = 250.000000
	__StakeholderShape__00000004_.Height = 70.000000
	__StakeholderShape__00000004_.IsHidden = false

	// insertion point for setup of pointers
	__Concern__00000000_.Outputs = append(__Concern__00000000_.Outputs, __Deliverable__00000004_)
	__Concern__00000000_.Outputs = append(__Concern__00000000_.Outputs, __Deliverable__00000001_)
	__Concern__00000000_.Outputs = append(__Concern__00000000_.Outputs, __Deliverable__00000003_)
	__Concern__00000001_.Outputs = append(__Concern__00000001_.Outputs, __Deliverable__00000000_)
	__ConcernOutputShape__00000000_.Task = __Concern__00000000_
	__ConcernOutputShape__00000000_.Deliverable = __Deliverable__00000004_
	__ConcernOutputShape__00000001_.Task = __Concern__00000000_
	__ConcernOutputShape__00000001_.Deliverable = __Deliverable__00000001_
	__ConcernOutputShape__00000002_.Task = __Concern__00000000_
	__ConcernOutputShape__00000002_.Deliverable = __Deliverable__00000003_
	__ConcernOutputShape__00000003_.Task = __Concern__00000001_
	__ConcernOutputShape__00000003_.Deliverable = __Deliverable__00000000_
	__ConcernShape__00000000_.Concern = __Concern__00000000_
	__ConcernShape__00000001_.Concern = __Concern__00000001_
	__DeliverableShape__00000000_.Deliverable = __Deliverable__00000000_
	__DeliverableShape__00000001_.Deliverable = __Deliverable__00000001_
	__DeliverableShape__00000003_.Deliverable = __Deliverable__00000003_
	__DeliverableShape__00000004_.Deliverable = __Deliverable__00000004_
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000000_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000001_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000003_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000004_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000000_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000001_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000000_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000001_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000002_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000003_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000000_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000001_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000003_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000004_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000000_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000001_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000002_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000000_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000001_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000003_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000004_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000000_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000001_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000000_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000001_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000003_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000004_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Stakeholder__00000000_.Concerns = append(__Stakeholder__00000000_.Concerns, __Concern__00000000_)
	__Stakeholder__00000001_.Concerns = append(__Stakeholder__00000001_.Concerns, __Concern__00000001_)
	__Stakeholder__00000003_.Concerns = append(__Stakeholder__00000003_.Concerns, __Concern__00000001_)
	__StakeholderConcernShape__00000000_.Stakeholder = __Stakeholder__00000000_
	__StakeholderConcernShape__00000000_.Concern = __Concern__00000000_
	__StakeholderConcernShape__00000001_.Stakeholder = __Stakeholder__00000001_
	__StakeholderConcernShape__00000001_.Concern = __Concern__00000001_
	__StakeholderConcernShape__00000002_.Stakeholder = __Stakeholder__00000003_
	__StakeholderConcernShape__00000002_.Concern = __Concern__00000001_
	__StakeholderShape__00000000_.Stakeholder = __Stakeholder__00000000_
	__StakeholderShape__00000001_.Stakeholder = __Stakeholder__00000001_
	__StakeholderShape__00000003_.Stakeholder = __Stakeholder__00000003_
	__StakeholderShape__00000004_.Stakeholder = __Stakeholder__00000004_
}
