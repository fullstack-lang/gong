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

	__Concept__00000000_ := (&models.Concept{Name: `State Machine`}).Stage(stage)
	__Concept__00000001_ := (&models.Concept{Name: `Transitons`}).Stage(stage)
	__Concept__00000002_ := (&models.Concept{Name: `Transition with message generation`}).Stage(stage)

	__ConceptShape__00000000_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000001_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000002_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)

	__Concern__00000000_ := (&models.Concern{Name: `develop CONOPS of Operationnal Messaging application`}).Stage(stage)
	__Concern__00000001_ := (&models.Concern{Name: `translate CONOPS for developers`}).Stage(stage)
	__Concern__00000002_ := (&models.Concern{Name: `Messages are legally bindings`}).Stage(stage)
	__Concern__00000003_ := (&models.Concern{Name: `Develop Tool`}).Stage(stage)
	__Concern__00000004_ := (&models.Concern{Name: `Reuse SysML Statemachines`}).Stage(stage)
	__Concern__00000005_ := (&models.Concern{Name: `Tailor Statemachines to add messages generation`}).Stage(stage)

	__ConcernCompositionShape__00000000_ := (&models.ConcernCompositionShape{Name: `share his CONOPS with stakeholders to Messages are legally bindings`}).Stage(stage)
	__ConcernCompositionShape__00000001_ := (&models.ConcernCompositionShape{Name: `Develop Tool to Reuse SysML Statemachines`}).Stage(stage)
	__ConcernCompositionShape__00000002_ := (&models.ConcernCompositionShape{Name: `Develop Tool to Tailor Statemachines to add messages generation`}).Stage(stage)

	__ConcernOutputShape__00000000_ := (&models.ConcernOutputShape{Name: `share his CONOPS with stakeholders to CONOPS`}).Stage(stage)
	__ConcernOutputShape__00000001_ := (&models.ConcernOutputShape{Name: `capture CONOPS for developpers to CONOPS`}).Stage(stage)
	__ConcernOutputShape__00000002_ := (&models.ConcernOutputShape{Name: `capture CONOPS for developpers to Structured CONOPS`}).Stage(stage)
	__ConcernOutputShape__00000003_ := (&models.ConcernOutputShape{Name: `Messages are legally bindings to Messages generated during the CONOPS`}).Stage(stage)
	__ConcernOutputShape__00000004_ := (&models.ConcernOutputShape{Name: `Develop Tool to MBSE Tool`}).Stage(stage)

	__ConcernShape__00000000_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000001_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000002_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000003_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000004_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000005_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)

	__ControlPointShape__00000000_ := (&models.ControlPointShape{Name: `Control Point Shape in Messages are legally bindings to Messages generated during the CONOPS 0`}).Stage(stage)
	__ControlPointShape__00000001_ := (&models.ControlPointShape{Name: `Control Point Shape in Messages are legally bindings to Messages generated during the CONOPS 1`}).Stage(stage)
	__ControlPointShape__00000002_ := (&models.ControlPointShape{Name: `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 0`}).Stage(stage)
	__ControlPointShape__00000003_ := (&models.ControlPointShape{Name: `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 1`}).Stage(stage)
	__ControlPointShape__00000004_ := (&models.ControlPointShape{Name: `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 2`}).Stage(stage)
	__ControlPointShape__00000005_ := (&models.ControlPointShape{Name: `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 2`}).Stage(stage)
	__ControlPointShape__00000006_ := (&models.ControlPointShape{Name: `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 2`}).Stage(stage)
	__ControlPointShape__00000007_ := (&models.ControlPointShape{Name: `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 3`}).Stage(stage)

	__Deliverable__00000000_ := (&models.Deliverable{Name: `CONOPS`}).Stage(stage)
	__Deliverable__00000001_ := (&models.Deliverable{Name: `Structured CONOPS`}).Stage(stage)
	__Deliverable__00000002_ := (&models.Deliverable{Name: `Scénarios`}).Stage(stage)
	__Deliverable__00000003_ := (&models.Deliverable{Name: `Structured Business Objects Lifecycle`}).Stage(stage)
	__Deliverable__00000004_ := (&models.Deliverable{Name: `Messages generated during the CONOPS`}).Stage(stage)
	__Deliverable__00000005_ := (&models.Deliverable{Name: `Messages generated during the lifecyle`}).Stage(stage)
	__Deliverable__00000006_ := (&models.Deliverable{Name: `MBSE Tool`}).Stage(stage)

	__DeliverableCompositionShape__00000000_ := (&models.DeliverableCompositionShape{Name: `Structured CONOPS to Scénarios`}).Stage(stage)
	__DeliverableCompositionShape__00000002_ := (&models.DeliverableCompositionShape{Name: `Structured CONOPS to Business Object`}).Stage(stage)
	__DeliverableCompositionShape__00000003_ := (&models.DeliverableCompositionShape{Name: `CONOPS to Messages generated during the CONOPS`}).Stage(stage)
	__DeliverableCompositionShape__00000004_ := (&models.DeliverableCompositionShape{Name: `Structured Business Objects Lifecycle to Messages generated during the lifecyle`}).Stage(stage)

	__DeliverableConceptShape__00000000_ := (&models.DeliverableConceptShape{Name: `Business Objects Lifecycle to State Machine`}).Stage(stage)
	__DeliverableConceptShape__00000001_ := (&models.DeliverableConceptShape{Name: `Messages generated during the lifecyle to Transitons`}).Stage(stage)
	__DeliverableConceptShape__00000002_ := (&models.DeliverableConceptShape{Name: `Messages generated during the lifecyle to Transition with message generation`}).Stage(stage)

	__DeliverableShape__00000000_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000001_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000002_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000003_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000004_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000005_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000006_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Capture of the state machine DSM concepts`}).Stage(stage)

	__Note__00000001_ := (&models.Note{Name: `The System of Interest is an operational Messaging Application `}).Stage(stage)

	__NoteShape__00000001_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)

	__NoteTaskShape__00000000_ := (&models.NoteTaskShape{Name: `The System of Interest is an operational Messaging Application  to Messages are legally bindings`}).Stage(stage)

	__Stakeholder__00000000_ := (&models.Stakeholder{Name: `End User`}).Stage(stage)
	__Stakeholder__00000001_ := (&models.Stakeholder{Name: `System Architect`}).Stage(stage)
	__Stakeholder__00000002_ := (&models.Stakeholder{Name: `Tool Developper`}).Stage(stage)
	__Stakeholder__00000003_ := (&models.Stakeholder{Name: `Model Developper`}).Stage(stage)

	__StakeholderConcernShape__00000000_ := (&models.StakeholderConcernShape{Name: `End User to share his CONOPS with stakeholders`}).Stage(stage)
	__StakeholderConcernShape__00000001_ := (&models.StakeholderConcernShape{Name: `System Architect to capture CONOPS for developpers`}).Stage(stage)
	__StakeholderConcernShape__00000002_ := (&models.StakeholderConcernShape{Name: `System Architect to share his CONOPS with stakeholders`}).Stage(stage)
	__StakeholderConcernShape__00000004_ := (&models.StakeholderConcernShape{Name: `Tool Developper to Reuse SysML Statemachines`}).Stage(stage)

	__StakeholderShape__00000000_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000001_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000002_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000003_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Concept__00000000_.Name = `State Machine`
	__Concept__00000000_.ComputedPrefix = ``
	__Concept__00000000_.IsExpanded = false
	__Concept__00000000_.LayoutDirection = models.Vertical

	__Concept__00000001_.Name = `Transitons`
	__Concept__00000001_.ComputedPrefix = ``
	__Concept__00000001_.IsExpanded = false
	__Concept__00000001_.LayoutDirection = models.Vertical

	__Concept__00000002_.Name = `Transition with message generation`
	__Concept__00000002_.ComputedPrefix = ``
	__Concept__00000002_.IsExpanded = false
	__Concept__00000002_.LayoutDirection = models.Vertical

	__ConceptShape__00000000_.Name = `-Default Diagram`
	__ConceptShape__00000000_.IsExpanded = false
	__ConceptShape__00000000_.X = 1185.236897
	__ConceptShape__00000000_.Y = 672.816417
	__ConceptShape__00000000_.Width = 250.000000
	__ConceptShape__00000000_.Height = 70.000000
	__ConceptShape__00000000_.IsHidden = false

	__ConceptShape__00000001_.Name = `-Default Diagram`
	__ConceptShape__00000001_.IsExpanded = false
	__ConceptShape__00000001_.X = 1195.929824
	__ConceptShape__00000001_.Y = 782.933474
	__ConceptShape__00000001_.Width = 250.000000
	__ConceptShape__00000001_.Height = 70.000000
	__ConceptShape__00000001_.IsHidden = false

	__ConceptShape__00000002_.Name = `-Default Diagram`
	__ConceptShape__00000002_.IsExpanded = false
	__ConceptShape__00000002_.X = 1187.506465
	__ConceptShape__00000002_.Y = 924.151534
	__ConceptShape__00000002_.Width = 250.000000
	__ConceptShape__00000002_.Height = 70.000000
	__ConceptShape__00000002_.IsHidden = false

	__Concern__00000000_.Name = `develop CONOPS of Operationnal Messaging application`
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

	__Concern__00000001_.Name = `translate CONOPS for developers`
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

	__Concern__00000002_.Name = `Messages are legally bindings`
	__Concern__00000002_.IDAirbus = ``
	__Concern__00000002_.Priority = ""
	__Concern__00000002_.ComputedPrefix = `1.1`
	__Concern__00000002_.IsExpanded = false
	__Concern__00000002_.LayoutDirection = models.Vertical
	__Concern__00000002_.Description = ``
	__Concern__00000002_.IsInputsNodeExpanded = false
	__Concern__00000002_.IsOutputsNodeExpanded = false
	__Concern__00000002_.IsWithCompletion = false
	__Concern__00000002_.Completion = ""

	__Concern__00000003_.Name = `Develop Tool`
	__Concern__00000003_.IDAirbus = ``
	__Concern__00000003_.Priority = ""
	__Concern__00000003_.ComputedPrefix = `3`
	__Concern__00000003_.IsExpanded = false
	__Concern__00000003_.LayoutDirection = models.Vertical
	__Concern__00000003_.Description = ``
	__Concern__00000003_.IsInputsNodeExpanded = false
	__Concern__00000003_.IsOutputsNodeExpanded = false
	__Concern__00000003_.IsWithCompletion = false
	__Concern__00000003_.Completion = ""

	__Concern__00000004_.Name = `Reuse SysML Statemachines`
	__Concern__00000004_.IDAirbus = ``
	__Concern__00000004_.Priority = ""
	__Concern__00000004_.ComputedPrefix = `3.1`
	__Concern__00000004_.IsExpanded = false
	__Concern__00000004_.LayoutDirection = models.Vertical
	__Concern__00000004_.Description = ``
	__Concern__00000004_.IsInputsNodeExpanded = false
	__Concern__00000004_.IsOutputsNodeExpanded = false
	__Concern__00000004_.IsWithCompletion = false
	__Concern__00000004_.Completion = ""

	__Concern__00000005_.Name = `Tailor Statemachines to add messages generation`
	__Concern__00000005_.IDAirbus = ``
	__Concern__00000005_.Priority = ""
	__Concern__00000005_.ComputedPrefix = `3.2`
	__Concern__00000005_.IsExpanded = false
	__Concern__00000005_.LayoutDirection = models.Vertical
	__Concern__00000005_.Description = ``
	__Concern__00000005_.IsInputsNodeExpanded = false
	__Concern__00000005_.IsOutputsNodeExpanded = false
	__Concern__00000005_.IsWithCompletion = false
	__Concern__00000005_.Completion = ""

	__ConcernCompositionShape__00000000_.Name = `share his CONOPS with stakeholders to Messages are legally bindings`
	__ConcernCompositionShape__00000000_.StartRatio = 0.500000
	__ConcernCompositionShape__00000000_.EndRatio = 0.500000
	__ConcernCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernCompositionShape__00000000_.CornerOffsetRatio = 1.532260
	__ConcernCompositionShape__00000000_.IsHidden = false

	__ConcernCompositionShape__00000001_.Name = `Develop Tool to Reuse SysML Statemachines`
	__ConcernCompositionShape__00000001_.StartRatio = 0.500000
	__ConcernCompositionShape__00000001_.EndRatio = 0.500000
	__ConcernCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernCompositionShape__00000001_.CornerOffsetRatio = 1.727902
	__ConcernCompositionShape__00000001_.IsHidden = false

	__ConcernCompositionShape__00000002_.Name = `Develop Tool to Tailor Statemachines to add messages generation`
	__ConcernCompositionShape__00000002_.StartRatio = 0.500000
	__ConcernCompositionShape__00000002_.EndRatio = 0.500000
	__ConcernCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernCompositionShape__00000002_.CornerOffsetRatio = 2.301688
	__ConcernCompositionShape__00000002_.IsHidden = false

	__ConcernOutputShape__00000000_.Name = `share his CONOPS with stakeholders to CONOPS`
	__ConcernOutputShape__00000000_.StartRatio = 0.500000
	__ConcernOutputShape__00000000_.EndRatio = 0.500000
	__ConcernOutputShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000000_.CornerOffsetRatio = 1.680000
	__ConcernOutputShape__00000000_.IsHidden = false

	__ConcernOutputShape__00000001_.Name = `capture CONOPS for developpers to CONOPS`
	__ConcernOutputShape__00000001_.StartRatio = 0.500000
	__ConcernOutputShape__00000001_.EndRatio = 0.500000
	__ConcernOutputShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000001_.CornerOffsetRatio = 1.680000
	__ConcernOutputShape__00000001_.IsHidden = false

	__ConcernOutputShape__00000002_.Name = `capture CONOPS for developpers to Structured CONOPS`
	__ConcernOutputShape__00000002_.StartRatio = 0.500000
	__ConcernOutputShape__00000002_.EndRatio = 0.500000
	__ConcernOutputShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000002_.CornerOffsetRatio = 1.680000
	__ConcernOutputShape__00000002_.IsHidden = false

	__ConcernOutputShape__00000003_.Name = `Messages are legally bindings to Messages generated during the CONOPS`
	__ConcernOutputShape__00000003_.StartRatio = 0.500000
	__ConcernOutputShape__00000003_.EndRatio = 0.500000
	__ConcernOutputShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000003_.CornerOffsetRatio = 1.680000
	__ConcernOutputShape__00000003_.IsHidden = false

	__ConcernOutputShape__00000004_.Name = `Develop Tool to MBSE Tool`
	__ConcernOutputShape__00000004_.StartRatio = 0.500000
	__ConcernOutputShape__00000004_.EndRatio = 0.500000
	__ConcernOutputShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ConcernOutputShape__00000004_.CornerOffsetRatio = 1.680000
	__ConcernOutputShape__00000004_.IsHidden = false

	__ConcernShape__00000000_.Name = `-Default Diagram`
	__ConcernShape__00000000_.IsExpanded = false
	__ConcernShape__00000000_.X = 403.378500
	__ConcernShape__00000000_.Y = 179.375107
	__ConcernShape__00000000_.Width = 250.000000
	__ConcernShape__00000000_.Height = 70.000000
	__ConcernShape__00000000_.IsHidden = false

	__ConcernShape__00000001_.Name = `-Default Diagram`
	__ConcernShape__00000001_.IsExpanded = false
	__ConcernShape__00000001_.X = 373.704161
	__ConcernShape__00000001_.Y = 348.421651
	__ConcernShape__00000001_.Width = 250.000000
	__ConcernShape__00000001_.Height = 70.000000
	__ConcernShape__00000001_.IsHidden = false

	__ConcernShape__00000002_.Name = `-Default Diagram`
	__ConcernShape__00000002_.IsExpanded = false
	__ConcernShape__00000002_.X = 758.084907
	__ConcernShape__00000002_.Y = 73.523131
	__ConcernShape__00000002_.Width = 250.000000
	__ConcernShape__00000002_.Height = 70.000000
	__ConcernShape__00000002_.IsHidden = false

	__ConcernShape__00000003_.Name = `-Default Diagram`
	__ConcernShape__00000003_.IsExpanded = false
	__ConcernShape__00000003_.X = 395.701814
	__ConcernShape__00000003_.Y = 949.792040
	__ConcernShape__00000003_.Width = 250.000000
	__ConcernShape__00000003_.Height = 70.000000
	__ConcernShape__00000003_.IsHidden = false

	__ConcernShape__00000004_.Name = `-Default Diagram`
	__ConcernShape__00000004_.IsExpanded = false
	__ConcernShape__00000004_.X = 291.560038
	__ConcernShape__00000004_.Y = 1121.698321
	__ConcernShape__00000004_.Width = 250.000000
	__ConcernShape__00000004_.Height = 70.000000
	__ConcernShape__00000004_.IsHidden = false

	__ConcernShape__00000005_.Name = `-Default Diagram`
	__ConcernShape__00000005_.IsExpanded = false
	__ConcernShape__00000005_.X = 599.816243
	__ConcernShape__00000005_.Y = 1131.028342
	__ConcernShape__00000005_.Width = 250.000000
	__ConcernShape__00000005_.Height = 70.000000
	__ConcernShape__00000005_.IsHidden = false

	__ControlPointShape__00000000_.Name = `Control Point Shape in Messages are legally bindings to Messages generated during the CONOPS 0`
	__ControlPointShape__00000000_.X_Relative = 0.362596
	__ControlPointShape__00000000_.Y_Relative = -1.934125
	__ControlPointShape__00000000_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000001_.Name = `Control Point Shape in Messages are legally bindings to Messages generated during the CONOPS 1`
	__ControlPointShape__00000001_.X_Relative = 0.552069
	__ControlPointShape__00000001_.Y_Relative = 0.547072
	__ControlPointShape__00000001_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000002_.Name = `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 0`
	__ControlPointShape__00000002_.X_Relative = 0.911759
	__ControlPointShape__00000002_.Y_Relative = 0.596052
	__ControlPointShape__00000002_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000003_.Name = `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 1`
	__ControlPointShape__00000003_.X_Relative = 1.207758
	__ControlPointShape__00000003_.Y_Relative = 1.121053
	__ControlPointShape__00000003_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000004_.Name = `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 2`
	__ControlPointShape__00000004_.X_Relative = 0.671760
	__ControlPointShape__00000004_.Y_Relative = 1.821054
	__ControlPointShape__00000004_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000005_.Name = `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 2`
	__ControlPointShape__00000005_.X_Relative = 0.691760
	__ControlPointShape__00000005_.Y_Relative = 1.871054
	__ControlPointShape__00000005_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000006_.Name = `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 2`
	__ControlPointShape__00000006_.X_Relative = 1.219758
	__ControlPointShape__00000006_.Y_Relative = 3.112717
	__ControlPointShape__00000006_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000007_.Name = `Control Point Shape in Modeling of the CONOPS translates a non structured CONOPS into a structure CONOPS to Messages generated during the lifecyle 3`
	__ControlPointShape__00000007_.X_Relative = 0.953029
	__ControlPointShape__00000007_.Y_Relative = 0.454706
	__ControlPointShape__00000007_.IsStartShapeTheClosestShape = false

	__Deliverable__00000000_.Name = `CONOPS`
	__Deliverable__00000000_.ComputedPrefix = `1`
	__Deliverable__00000000_.IsExpanded = false
	__Deliverable__00000000_.LayoutDirection = models.Vertical
	__Deliverable__00000000_.Description = ``
	__Deliverable__00000000_.IsProducersNodeExpanded = false
	__Deliverable__00000000_.IsConsumersNodeExpanded = false

	__Deliverable__00000001_.Name = `Structured CONOPS`
	__Deliverable__00000001_.ComputedPrefix = `2`
	__Deliverable__00000001_.IsExpanded = false
	__Deliverable__00000001_.LayoutDirection = models.Vertical
	__Deliverable__00000001_.Description = ``
	__Deliverable__00000001_.IsProducersNodeExpanded = false
	__Deliverable__00000001_.IsConsumersNodeExpanded = false

	__Deliverable__00000002_.Name = `Scénarios`
	__Deliverable__00000002_.ComputedPrefix = `2.1`
	__Deliverable__00000002_.IsExpanded = false
	__Deliverable__00000002_.LayoutDirection = models.Vertical
	__Deliverable__00000002_.Description = ``
	__Deliverable__00000002_.IsProducersNodeExpanded = false
	__Deliverable__00000002_.IsConsumersNodeExpanded = false

	__Deliverable__00000003_.Name = `Structured Business Objects Lifecycle`
	__Deliverable__00000003_.ComputedPrefix = `2.2`
	__Deliverable__00000003_.IsExpanded = false
	__Deliverable__00000003_.LayoutDirection = models.Vertical
	__Deliverable__00000003_.Description = ``
	__Deliverable__00000003_.IsProducersNodeExpanded = false
	__Deliverable__00000003_.IsConsumersNodeExpanded = false

	__Deliverable__00000004_.Name = `Messages generated during the CONOPS`
	__Deliverable__00000004_.ComputedPrefix = `1.1`
	__Deliverable__00000004_.IsExpanded = false
	__Deliverable__00000004_.LayoutDirection = models.Vertical
	__Deliverable__00000004_.Description = ``
	__Deliverable__00000004_.IsProducersNodeExpanded = false
	__Deliverable__00000004_.IsConsumersNodeExpanded = false

	__Deliverable__00000005_.Name = `Messages generated during the lifecyle`
	__Deliverable__00000005_.ComputedPrefix = `2.2.1`
	__Deliverable__00000005_.IsExpanded = false
	__Deliverable__00000005_.LayoutDirection = models.Vertical
	__Deliverable__00000005_.Description = ``
	__Deliverable__00000005_.IsProducersNodeExpanded = false
	__Deliverable__00000005_.IsConsumersNodeExpanded = false

	__Deliverable__00000006_.Name = `MBSE Tool`
	__Deliverable__00000006_.ComputedPrefix = `3`
	__Deliverable__00000006_.IsExpanded = false
	__Deliverable__00000006_.LayoutDirection = models.Vertical
	__Deliverable__00000006_.Description = ``
	__Deliverable__00000006_.IsProducersNodeExpanded = false
	__Deliverable__00000006_.IsConsumersNodeExpanded = false

	__DeliverableCompositionShape__00000000_.Name = `Structured CONOPS to Scénarios`
	__DeliverableCompositionShape__00000000_.StartRatio = 0.500000
	__DeliverableCompositionShape__00000000_.EndRatio = 0.500000
	__DeliverableCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__DeliverableCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__DeliverableCompositionShape__00000000_.CornerOffsetRatio = 1.977013
	__DeliverableCompositionShape__00000000_.IsHidden = false

	__DeliverableCompositionShape__00000002_.Name = `Structured CONOPS to Business Object`
	__DeliverableCompositionShape__00000002_.StartRatio = 0.500000
	__DeliverableCompositionShape__00000002_.EndRatio = 0.500000
	__DeliverableCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__DeliverableCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__DeliverableCompositionShape__00000002_.CornerOffsetRatio = 1.752221
	__DeliverableCompositionShape__00000002_.IsHidden = false

	__DeliverableCompositionShape__00000003_.Name = `CONOPS to Messages generated during the CONOPS`
	__DeliverableCompositionShape__00000003_.StartRatio = 0.500000
	__DeliverableCompositionShape__00000003_.EndRatio = 0.500000
	__DeliverableCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__DeliverableCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__DeliverableCompositionShape__00000003_.CornerOffsetRatio = 1.239833
	__DeliverableCompositionShape__00000003_.IsHidden = false

	__DeliverableCompositionShape__00000004_.Name = `Structured Business Objects Lifecycle to Messages generated during the lifecyle`
	__DeliverableCompositionShape__00000004_.StartRatio = 0.500000
	__DeliverableCompositionShape__00000004_.EndRatio = 0.500000
	__DeliverableCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__DeliverableCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__DeliverableCompositionShape__00000004_.CornerOffsetRatio = 1.681349
	__DeliverableCompositionShape__00000004_.IsHidden = false

	__DeliverableConceptShape__00000000_.Name = `Business Objects Lifecycle to State Machine`
	__DeliverableConceptShape__00000000_.StartRatio = 0.500000
	__DeliverableConceptShape__00000000_.EndRatio = 0.500000
	__DeliverableConceptShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000000_.CornerOffsetRatio = 1.680000
	__DeliverableConceptShape__00000000_.IsHidden = false

	__DeliverableConceptShape__00000001_.Name = `Messages generated during the lifecyle to Transitons`
	__DeliverableConceptShape__00000001_.StartRatio = 0.500000
	__DeliverableConceptShape__00000001_.EndRatio = 0.500000
	__DeliverableConceptShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000001_.CornerOffsetRatio = 1.680000
	__DeliverableConceptShape__00000001_.IsHidden = false

	__DeliverableConceptShape__00000002_.Name = `Messages generated during the lifecyle to Transition with message generation`
	__DeliverableConceptShape__00000002_.StartRatio = 0.500000
	__DeliverableConceptShape__00000002_.EndRatio = 0.500000
	__DeliverableConceptShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__DeliverableConceptShape__00000002_.CornerOffsetRatio = 1.680000
	__DeliverableConceptShape__00000002_.IsHidden = false

	__DeliverableShape__00000000_.Name = `-Default Diagram`
	__DeliverableShape__00000000_.IsExpanded = false
	__DeliverableShape__00000000_.X = 753.753767
	__DeliverableShape__00000000_.Y = 260.479157
	__DeliverableShape__00000000_.Width = 250.000000
	__DeliverableShape__00000000_.Height = 70.000000
	__DeliverableShape__00000000_.IsHidden = false

	__DeliverableShape__00000001_.Name = `-Default Diagram`
	__DeliverableShape__00000001_.IsExpanded = false
	__DeliverableShape__00000001_.X = 749.673110
	__DeliverableShape__00000001_.Y = 429.330291
	__DeliverableShape__00000001_.Width = 250.000000
	__DeliverableShape__00000001_.Height = 70.000000
	__DeliverableShape__00000001_.IsHidden = false

	__DeliverableShape__00000002_.Name = `-Default Diagram`
	__DeliverableShape__00000002_.IsExpanded = false
	__DeliverableShape__00000002_.X = 568.395240
	__DeliverableShape__00000002_.Y = 580.112103
	__DeliverableShape__00000002_.Width = 250.000000
	__DeliverableShape__00000002_.Height = 70.000000
	__DeliverableShape__00000002_.IsHidden = false

	__DeliverableShape__00000003_.Name = `-Default Diagram`
	__DeliverableShape__00000003_.IsExpanded = false
	__DeliverableShape__00000003_.X = 891.287537
	__DeliverableShape__00000003_.Y = 568.641290
	__DeliverableShape__00000003_.Width = 250.000000
	__DeliverableShape__00000003_.Height = 70.000000
	__DeliverableShape__00000003_.IsHidden = false

	__DeliverableShape__00000004_.Name = `-Default Diagram`
	__DeliverableShape__00000004_.IsExpanded = false
	__DeliverableShape__00000004_.X = 1156.517501
	__DeliverableShape__00000004_.Y = 346.424189
	__DeliverableShape__00000004_.Width = 250.000000
	__DeliverableShape__00000004_.Height = 70.000000
	__DeliverableShape__00000004_.IsHidden = false

	__DeliverableShape__00000005_.Name = `-Default Diagram`
	__DeliverableShape__00000005_.IsExpanded = false
	__DeliverableShape__00000005_.X = 849.200771
	__DeliverableShape__00000005_.Y = 797.503870
	__DeliverableShape__00000005_.Width = 250.000000
	__DeliverableShape__00000005_.Height = 70.000000
	__DeliverableShape__00000005_.IsHidden = false

	__DeliverableShape__00000006_.Name = `-Default Diagram`
	__DeliverableShape__00000006_.IsExpanded = false
	__DeliverableShape__00000006_.X = 757.077074
	__DeliverableShape__00000006_.Y = 959.044617
	__DeliverableShape__00000006_.Width = 250.000000
	__DeliverableShape__00000006_.Height = 70.000000
	__DeliverableShape__00000006_.IsHidden = false

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.LayoutDirection = models.Vertical
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 20800.000000
	__Diagram__00000000_.Height = 20800.000000
	__Diagram__00000000_.IsRequirementsNodeExpanded = false
	__Diagram__00000000_.IsConceptsNodeExpanded = true
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsConcernsNodeExpanded = true
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsStakeholdersNodeExpanded = false
	__Diagram__00000000_.IsDiagramsNodeExpanded = false

	__Library__00000000_.Name = `Capture of the state machine DSM concepts`
	__Library__00000000_.IsRootLibrary = true
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.NbPixPerCharacter = 8.000000

	__Note__00000001_.Name = `The System of Interest is an operational Messaging Application `
	__Note__00000001_.ComputedPrefix = `1`
	__Note__00000001_.IsExpanded = false
	__Note__00000001_.LayoutDirection = models.Vertical

	__NoteShape__00000001_.Name = `-Default Diagram`
	__NoteShape__00000001_.IsExpanded = false
	__NoteShape__00000001_.X = 367.079731
	__NoteShape__00000001_.Y = 7.955096
	__NoteShape__00000001_.Width = 250.000000
	__NoteShape__00000001_.Height = 70.000000
	__NoteShape__00000001_.IsHidden = false

	__NoteTaskShape__00000000_.Name = `The System of Interest is an operational Messaging Application  to Messages are legally bindings`
	__NoteTaskShape__00000000_.StartRatio = 0.500000
	__NoteTaskShape__00000000_.EndRatio = 0.500000
	__NoteTaskShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000000_.CornerOffsetRatio = 1.680000
	__NoteTaskShape__00000000_.IsHidden = false

	__Stakeholder__00000000_.Name = `End User`
	__Stakeholder__00000000_.IDAirbus = ``
	__Stakeholder__00000000_.ComputedPrefix = `1`
	__Stakeholder__00000000_.IsExpanded = false
	__Stakeholder__00000000_.LayoutDirection = models.Vertical
	__Stakeholder__00000000_.Description = ``

	__Stakeholder__00000001_.Name = `System Architect`
	__Stakeholder__00000001_.IDAirbus = ``
	__Stakeholder__00000001_.ComputedPrefix = `2`
	__Stakeholder__00000001_.IsExpanded = false
	__Stakeholder__00000001_.LayoutDirection = models.Vertical
	__Stakeholder__00000001_.Description = ``

	__Stakeholder__00000002_.Name = `Tool Developper`
	__Stakeholder__00000002_.IDAirbus = ``
	__Stakeholder__00000002_.ComputedPrefix = `3`
	__Stakeholder__00000002_.IsExpanded = false
	__Stakeholder__00000002_.LayoutDirection = models.Vertical
	__Stakeholder__00000002_.Description = ``

	__Stakeholder__00000003_.Name = `Model Developper`
	__Stakeholder__00000003_.IDAirbus = ``
	__Stakeholder__00000003_.ComputedPrefix = `4`
	__Stakeholder__00000003_.IsExpanded = false
	__Stakeholder__00000003_.LayoutDirection = models.Vertical
	__Stakeholder__00000003_.Description = ``

	__StakeholderConcernShape__00000000_.Name = `End User to share his CONOPS with stakeholders`
	__StakeholderConcernShape__00000000_.StartRatio = 0.500000
	__StakeholderConcernShape__00000000_.EndRatio = 0.500000
	__StakeholderConcernShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000000_.CornerOffsetRatio = 1.680000
	__StakeholderConcernShape__00000000_.IsHidden = false

	__StakeholderConcernShape__00000001_.Name = `System Architect to capture CONOPS for developpers`
	__StakeholderConcernShape__00000001_.StartRatio = 0.500000
	__StakeholderConcernShape__00000001_.EndRatio = 0.500000
	__StakeholderConcernShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000001_.CornerOffsetRatio = 1.680000
	__StakeholderConcernShape__00000001_.IsHidden = false

	__StakeholderConcernShape__00000002_.Name = `System Architect to share his CONOPS with stakeholders`
	__StakeholderConcernShape__00000002_.StartRatio = 0.500000
	__StakeholderConcernShape__00000002_.EndRatio = 0.500000
	__StakeholderConcernShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000002_.CornerOffsetRatio = 1.680000
	__StakeholderConcernShape__00000002_.IsHidden = false

	__StakeholderConcernShape__00000004_.Name = `Tool Developper to Reuse SysML Statemachines`
	__StakeholderConcernShape__00000004_.StartRatio = 0.500000
	__StakeholderConcernShape__00000004_.EndRatio = 0.500000
	__StakeholderConcernShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__StakeholderConcernShape__00000004_.CornerOffsetRatio = 1.680000
	__StakeholderConcernShape__00000004_.IsHidden = false

	__StakeholderShape__00000000_.Name = `-Default Diagram`
	__StakeholderShape__00000000_.IsExpanded = false
	__StakeholderShape__00000000_.X = 30.159868
	__StakeholderShape__00000000_.Y = 94.790253
	__StakeholderShape__00000000_.Width = 250.000000
	__StakeholderShape__00000000_.Height = 70.000000
	__StakeholderShape__00000000_.IsHidden = false

	__StakeholderShape__00000001_.Name = `-Default Diagram`
	__StakeholderShape__00000001_.IsExpanded = false
	__StakeholderShape__00000001_.X = 27.281448
	__StakeholderShape__00000001_.Y = 293.004818
	__StakeholderShape__00000001_.Width = 250.000000
	__StakeholderShape__00000001_.Height = 70.000000
	__StakeholderShape__00000001_.IsHidden = false

	__StakeholderShape__00000002_.Name = `-Default Diagram`
	__StakeholderShape__00000002_.IsExpanded = false
	__StakeholderShape__00000002_.X = 39.479303
	__StakeholderShape__00000002_.Y = 961.924090
	__StakeholderShape__00000002_.Width = 250.000000
	__StakeholderShape__00000002_.Height = 70.000000
	__StakeholderShape__00000002_.IsHidden = false

	__StakeholderShape__00000003_.Name = `-Default Diagram`
	__StakeholderShape__00000003_.IsExpanded = false
	__StakeholderShape__00000003_.X = 46.471908
	__StakeholderShape__00000003_.Y = 1270.288029
	__StakeholderShape__00000003_.Width = 250.000000
	__StakeholderShape__00000003_.Height = 70.000000
	__StakeholderShape__00000003_.IsHidden = false

	// insertion point for setup of pointers
	__ConceptShape__00000000_.Concept = __Concept__00000000_
	__ConceptShape__00000001_.Concept = __Concept__00000001_
	__ConceptShape__00000002_.Concept = __Concept__00000002_
	__Concern__00000000_.SubConcerns = append(__Concern__00000000_.SubConcerns, __Concern__00000002_)
	__Concern__00000000_.Outputs = append(__Concern__00000000_.Outputs, __Deliverable__00000000_)
	__Concern__00000001_.Outputs = append(__Concern__00000001_.Outputs, __Deliverable__00000000_)
	__Concern__00000001_.Outputs = append(__Concern__00000001_.Outputs, __Deliverable__00000001_)
	__Concern__00000002_.Outputs = append(__Concern__00000002_.Outputs, __Deliverable__00000004_)
	__Concern__00000003_.SubConcerns = append(__Concern__00000003_.SubConcerns, __Concern__00000004_)
	__Concern__00000003_.SubConcerns = append(__Concern__00000003_.SubConcerns, __Concern__00000005_)
	__Concern__00000003_.Outputs = append(__Concern__00000003_.Outputs, __Deliverable__00000006_)
	__ConcernCompositionShape__00000000_.Concern = __Concern__00000002_
	__ConcernCompositionShape__00000001_.Concern = __Concern__00000004_
	__ConcernCompositionShape__00000002_.Concern = __Concern__00000005_
	__ConcernOutputShape__00000000_.Concern = __Concern__00000000_
	__ConcernOutputShape__00000000_.Deliverable = __Deliverable__00000000_
	__ConcernOutputShape__00000001_.Concern = __Concern__00000001_
	__ConcernOutputShape__00000001_.Deliverable = __Deliverable__00000000_
	__ConcernOutputShape__00000002_.Concern = __Concern__00000001_
	__ConcernOutputShape__00000002_.Deliverable = __Deliverable__00000001_
	__ConcernOutputShape__00000003_.Concern = __Concern__00000002_
	__ConcernOutputShape__00000003_.Deliverable = __Deliverable__00000004_
	__ConcernOutputShape__00000003_.ControlPointShapes = append(__ConcernOutputShape__00000003_.ControlPointShapes, __ControlPointShape__00000000_)
	__ConcernOutputShape__00000003_.ControlPointShapes = append(__ConcernOutputShape__00000003_.ControlPointShapes, __ControlPointShape__00000001_)
	__ConcernOutputShape__00000004_.Concern = __Concern__00000003_
	__ConcernOutputShape__00000004_.Deliverable = __Deliverable__00000006_
	__ConcernShape__00000000_.Concern = __Concern__00000000_
	__ConcernShape__00000001_.Concern = __Concern__00000001_
	__ConcernShape__00000002_.Concern = __Concern__00000002_
	__ConcernShape__00000003_.Concern = __Concern__00000003_
	__ConcernShape__00000004_.Concern = __Concern__00000004_
	__ConcernShape__00000005_.Concern = __Concern__00000005_
	__Deliverable__00000000_.SubDeliverables = append(__Deliverable__00000000_.SubDeliverables, __Deliverable__00000004_)
	__Deliverable__00000001_.SubDeliverables = append(__Deliverable__00000001_.SubDeliverables, __Deliverable__00000002_)
	__Deliverable__00000001_.SubDeliverables = append(__Deliverable__00000001_.SubDeliverables, __Deliverable__00000003_)
	__Deliverable__00000003_.SubDeliverables = append(__Deliverable__00000003_.SubDeliverables, __Deliverable__00000005_)
	__Deliverable__00000003_.Concepts = append(__Deliverable__00000003_.Concepts, __Concept__00000000_)
	__Deliverable__00000005_.Concepts = append(__Deliverable__00000005_.Concepts, __Concept__00000001_)
	__Deliverable__00000005_.Concepts = append(__Deliverable__00000005_.Concepts, __Concept__00000002_)
	__DeliverableCompositionShape__00000000_.Deliverable = __Deliverable__00000002_
	__DeliverableCompositionShape__00000002_.Deliverable = __Deliverable__00000003_
	__DeliverableCompositionShape__00000003_.Deliverable = __Deliverable__00000004_
	__DeliverableCompositionShape__00000004_.Deliverable = __Deliverable__00000005_
	__DeliverableConceptShape__00000000_.Deliverable = __Deliverable__00000003_
	__DeliverableConceptShape__00000000_.Concept = __Concept__00000000_
	__DeliverableConceptShape__00000001_.Deliverable = __Deliverable__00000005_
	__DeliverableConceptShape__00000001_.Concept = __Concept__00000001_
	__DeliverableConceptShape__00000002_.Deliverable = __Deliverable__00000005_
	__DeliverableConceptShape__00000002_.Concept = __Concept__00000002_
	__DeliverableShape__00000000_.Deliverable = __Deliverable__00000000_
	__DeliverableShape__00000001_.Deliverable = __Deliverable__00000001_
	__DeliverableShape__00000002_.Deliverable = __Deliverable__00000002_
	__DeliverableShape__00000003_.Deliverable = __Deliverable__00000003_
	__DeliverableShape__00000004_.Deliverable = __Deliverable__00000004_
	__DeliverableShape__00000005_.Deliverable = __Deliverable__00000005_
	__DeliverableShape__00000006_.Deliverable = __Deliverable__00000006_
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000000_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000001_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000002_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000003_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000004_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000005_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000006_)
	__Diagram__00000000_.DeliverablesWhoseNodeIsExpanded = append(__Diagram__00000000_.DeliverablesWhoseNodeIsExpanded, __Deliverable__00000001_)
	__Diagram__00000000_.DeliverablesWhoseNodeIsExpanded = append(__Diagram__00000000_.DeliverablesWhoseNodeIsExpanded, __Deliverable__00000000_)
	__Diagram__00000000_.DeliverablesWhoseNodeIsExpanded = append(__Diagram__00000000_.DeliverablesWhoseNodeIsExpanded, __Deliverable__00000003_)
	__Diagram__00000000_.DeliverablesWhoseNodeIsExpanded = append(__Diagram__00000000_.DeliverablesWhoseNodeIsExpanded, __Deliverable__00000006_)
	__Diagram__00000000_.DeliverableComposition_Shapes = append(__Diagram__00000000_.DeliverableComposition_Shapes, __DeliverableCompositionShape__00000000_)
	__Diagram__00000000_.DeliverableComposition_Shapes = append(__Diagram__00000000_.DeliverableComposition_Shapes, __DeliverableCompositionShape__00000002_)
	__Diagram__00000000_.DeliverableComposition_Shapes = append(__Diagram__00000000_.DeliverableComposition_Shapes, __DeliverableCompositionShape__00000003_)
	__Diagram__00000000_.DeliverableComposition_Shapes = append(__Diagram__00000000_.DeliverableComposition_Shapes, __DeliverableCompositionShape__00000004_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000000_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000001_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000002_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000003_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000004_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000005_)
	__Diagram__00000000_.ConcernsWhoseNodeIsExpanded = append(__Diagram__00000000_.ConcernsWhoseNodeIsExpanded, __Concern__00000000_)
	__Diagram__00000000_.ConcernsWhoseNodeIsExpanded = append(__Diagram__00000000_.ConcernsWhoseNodeIsExpanded, __Concern__00000003_)
	__Diagram__00000000_.ConcernComposition_Shapes = append(__Diagram__00000000_.ConcernComposition_Shapes, __ConcernCompositionShape__00000000_)
	__Diagram__00000000_.ConcernComposition_Shapes = append(__Diagram__00000000_.ConcernComposition_Shapes, __ConcernCompositionShape__00000001_)
	__Diagram__00000000_.ConcernComposition_Shapes = append(__Diagram__00000000_.ConcernComposition_Shapes, __ConcernCompositionShape__00000002_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000000_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000001_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000002_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000003_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000004_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000001_)
	__Diagram__00000000_.NoteTaskShapes = append(__Diagram__00000000_.NoteTaskShapes, __NoteTaskShape__00000000_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000000_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000001_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000002_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000003_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000000_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000001_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000002_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000004_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000000_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000001_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000002_)
	__Diagram__00000000_.ConceptsWhoseNodeIsExpanded = append(__Diagram__00000000_.ConceptsWhoseNodeIsExpanded, __Concept__00000000_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000000_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000001_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000002_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000000_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000001_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000006_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000000_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000001_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000003_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000000_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000001_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000002_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000003_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000000_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000001_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000002_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000001_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Note__00000001_.Tasks = append(__Note__00000001_.Tasks, __Concern__00000002_)
	__NoteShape__00000001_.Note = __Note__00000001_
	__NoteTaskShape__00000000_.Note = __Note__00000001_
	__NoteTaskShape__00000000_.Task = __Concern__00000002_
	__Stakeholder__00000000_.Concerns = append(__Stakeholder__00000000_.Concerns, __Concern__00000000_)
	__Stakeholder__00000001_.Concerns = append(__Stakeholder__00000001_.Concerns, __Concern__00000001_)
	__Stakeholder__00000001_.Concerns = append(__Stakeholder__00000001_.Concerns, __Concern__00000000_)
	__Stakeholder__00000002_.Concerns = append(__Stakeholder__00000002_.Concerns, __Concern__00000004_)
	__StakeholderConcernShape__00000000_.Stakeholder = __Stakeholder__00000000_
	__StakeholderConcernShape__00000000_.Concern = __Concern__00000000_
	__StakeholderConcernShape__00000001_.Stakeholder = __Stakeholder__00000001_
	__StakeholderConcernShape__00000001_.Concern = __Concern__00000001_
	__StakeholderConcernShape__00000002_.Stakeholder = __Stakeholder__00000001_
	__StakeholderConcernShape__00000002_.Concern = __Concern__00000000_
	__StakeholderConcernShape__00000004_.Stakeholder = __Stakeholder__00000002_
	__StakeholderConcernShape__00000004_.Concern = __Concern__00000004_
	__StakeholderShape__00000000_.Stakeholder = __Stakeholder__00000000_
	__StakeholderShape__00000001_.Stakeholder = __Stakeholder__00000001_
	__StakeholderShape__00000002_.Stakeholder = __Stakeholder__00000002_
	__StakeholderShape__00000003_.Stakeholder = __Stakeholder__00000003_
}
