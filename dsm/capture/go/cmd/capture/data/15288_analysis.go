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

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Concept__00000000_ := (&models.Concept{Name: `Gap`}).Stage(stage)
	__Concept__00000001_ := (&models.Concept{Name: `Function`}).Stage(stage)
	__Concept__00000002_ := (&models.Concept{Name: `Operational Scenario`}).Stage(stage)
	__Concept__00000003_ := (&models.Concept{Name: `Mission Requirement`}).Stage(stage)
	__Concept__00000004_ := (&models.Concept{Name: `Stakeholder Concept`}).Stage(stage)
	__Concept__00000005_ := (&models.Concept{Name: `Problem`}).Stage(stage)
	__Concept__00000006_ := (&models.Concept{Name: `Opportunity`}).Stage(stage)
	__Concept__00000007_ := (&models.Concept{Name: `Solution Class`}).Stage(stage)
	__Concept__00000008_ := (&models.Concept{Name: `Risk`}).Stage(stage)
	__Concept__00000009_ := (&models.Concept{Name: `Benefit`}).Stage(stage)

	__ConceptShape__00000000_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000001_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000002_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000003_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000004_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000005_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000006_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000007_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000008_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)
	__ConceptShape__00000009_ := (&models.ConceptShape{Name: `-Default Diagram`}).Stage(stage)

	__Concern__00000000_ := (&models.Concern{Name: `Understand Environment & Mission`}).Stage(stage)
	__Concern__00000001_ := (&models.Concern{Name: `Map Stakeholders`}).Stage(stage)
	__Concern__00000002_ := (&models.Concern{Name: `Define Problem or Opportunity Space`}).Stage(stage)
	__Concern__00000003_ := (&models.Concern{Name: `Characterize Solution Space`}).Stage(stage)
	__Concern__00000004_ := (&models.Concern{Name: `Assess Business Case`}).Stage(stage)

	__ConcernOutputShape__00000000_ := (&models.ConcernOutputShape{Name: `Understand Environment & Mission to Gap Analysis`}).Stage(stage)
	__ConcernOutputShape__00000001_ := (&models.ConcernOutputShape{Name: `Understand Environment & Mission to OpsCon`}).Stage(stage)
	__ConcernOutputShape__00000002_ := (&models.ConcernOutputShape{Name: `Understand Environment & Mission to Mission Requirements`}).Stage(stage)
	__ConcernOutputShape__00000003_ := (&models.ConcernOutputShape{Name: `Define Problem or Opportunity Space to Problem or Opportunity Statement`}).Stage(stage)
	__ConcernOutputShape__00000004_ := (&models.ConcernOutputShape{Name: `Map Stakeholders to Stakeholder Map`}).Stage(stage)
	__ConcernOutputShape__00000005_ := (&models.ConcernOutputShape{Name: `Characterize Solution Space to Alternative Solution Classes`}).Stage(stage)
	__ConcernOutputShape__00000006_ := (&models.ConcernOutputShape{Name: `Assess Business Case to Business Case`}).Stage(stage)

	__ConcernShape__00000000_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000001_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000002_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000003_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)
	__ConcernShape__00000004_ := (&models.ConcernShape{Name: `-Default Diagram`}).Stage(stage)

	__ControlPointShape__00000000_ := (&models.ControlPointShape{Name: `Control Point Shape in Problem or Opportunity Statement to Opportunity 0`}).Stage(stage)
	__ControlPointShape__00000001_ := (&models.ControlPointShape{Name: `Control Point Shape in Problem or Opportunity Statement to Problem 0`}).Stage(stage)
	__ControlPointShape__00000002_ := (&models.ControlPointShape{Name: `Control Point Shape in Mission Analyst to Define Problem or Opportunity Space 0`}).Stage(stage)
	__ControlPointShape__00000003_ := (&models.ControlPointShape{Name: `Control Point Shape in Lead System Engineer to Map Stakeholders 0`}).Stage(stage)
	__ControlPointShape__00000004_ := (&models.ControlPointShape{Name: `Control Point Shape in Enterprise Architect to Characterize Solution Space 0`}).Stage(stage)
	__ControlPointShape__00000005_ := (&models.ControlPointShape{Name: `Control Point Shape in Problem or Opportunity Statement to Opportunity 0`}).Stage(stage)
	__ControlPointShape__00000006_ := (&models.ControlPointShape{Name: `Control Point Shape in Understand Environment & Mission to Mission Requirements 0`}).Stage(stage)
	__ControlPointShape__00000007_ := (&models.ControlPointShape{Name: `Control Point Shape in Understand Environment & Mission to Mission Requirements 1`}).Stage(stage)

	__Deliverable__00000000_ := (&models.Deliverable{Name: `Gap Analysis`}).Stage(stage)
	__Deliverable__00000001_ := (&models.Deliverable{Name: `OpsCon`}).Stage(stage)
	__Deliverable__00000002_ := (&models.Deliverable{Name: `Mission Requirements`}).Stage(stage)
	__Deliverable__00000003_ := (&models.Deliverable{Name: `Stakeholder Map`}).Stage(stage)
	__Deliverable__00000004_ := (&models.Deliverable{Name: `Problem or Opportunity Statement`}).Stage(stage)
	__Deliverable__00000005_ := (&models.Deliverable{Name: `Alternative Solution Classes`}).Stage(stage)
	__Deliverable__00000006_ := (&models.Deliverable{Name: `Business Case`}).Stage(stage)

	__DeliverableConceptShape__00000000_ := (&models.DeliverableConceptShape{Name: `Gap Analysis to Gap`}).Stage(stage)
	__DeliverableConceptShape__00000001_ := (&models.DeliverableConceptShape{Name: `Gap Analysis to Function`}).Stage(stage)
	__DeliverableConceptShape__00000002_ := (&models.DeliverableConceptShape{Name: `OpsCon to Operational Scenario`}).Stage(stage)
	__DeliverableConceptShape__00000003_ := (&models.DeliverableConceptShape{Name: `Mission Requirements to Mission Requirement`}).Stage(stage)
	__DeliverableConceptShape__00000004_ := (&models.DeliverableConceptShape{Name: `Problem or Opportunity Statement to Problem`}).Stage(stage)
	__DeliverableConceptShape__00000005_ := (&models.DeliverableConceptShape{Name: `Problem or Opportunity Statement to Opportunity`}).Stage(stage)
	__DeliverableConceptShape__00000006_ := (&models.DeliverableConceptShape{Name: `Stakeholder Map to Stakeholder Concept`}).Stage(stage)
	__DeliverableConceptShape__00000007_ := (&models.DeliverableConceptShape{Name: `Alternative Solution Classes to Solution Class`}).Stage(stage)
	__DeliverableConceptShape__00000008_ := (&models.DeliverableConceptShape{Name: `Business Case to Risk`}).Stage(stage)
	__DeliverableConceptShape__00000009_ := (&models.DeliverableConceptShape{Name: `Business Case to Benefit`}).Stage(stage)

	__DeliverableShape__00000000_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000001_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000002_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000003_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000004_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000005_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)
	__DeliverableShape__00000006_ := (&models.DeliverableShape{Name: `-Default Diagram`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Stakeholder__00000000_ := (&models.Stakeholder{Name: `Mission Analyst`}).Stage(stage)
	__Stakeholder__00000001_ := (&models.Stakeholder{Name: `Enterprise Architect`}).Stage(stage)
	__Stakeholder__00000002_ := (&models.Stakeholder{Name: `Lead System Engineer`}).Stage(stage)
	__Stakeholder__00000003_ := (&models.Stakeholder{Name: `Executive Sponsor`}).Stage(stage)
	__Stakeholder__00000004_ := (&models.Stakeholder{Name: `Business Analyst`}).Stage(stage)

	__StakeholderConcernShape__00000000_ := (&models.StakeholderConcernShape{Name: `Mission Analyst to Understand Environment & Mission`}).Stage(stage)
	__StakeholderConcernShape__00000001_ := (&models.StakeholderConcernShape{Name: `Mission Analyst to Define Problem or Opportunity Space`}).Stage(stage)
	__StakeholderConcernShape__00000002_ := (&models.StakeholderConcernShape{Name: `Business Analyst to Define Problem or Opportunity Space`}).Stage(stage)
	__StakeholderConcernShape__00000003_ := (&models.StakeholderConcernShape{Name: `Enterprise Architect to Map Stakeholders`}).Stage(stage)
	__StakeholderConcernShape__00000004_ := (&models.StakeholderConcernShape{Name: `Enterprise Architect to Characterize Solution Space`}).Stage(stage)
	__StakeholderConcernShape__00000005_ := (&models.StakeholderConcernShape{Name: `Lead System Engineer to Map Stakeholders`}).Stage(stage)
	__StakeholderConcernShape__00000006_ := (&models.StakeholderConcernShape{Name: `Lead System Engineer to Characterize Solution Space`}).Stage(stage)
	__StakeholderConcernShape__00000007_ := (&models.StakeholderConcernShape{Name: `Executive Sponsor to Assess Business Case`}).Stage(stage)

	__StakeholderShape__00000000_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000001_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000002_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000003_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)
	__StakeholderShape__00000004_ := (&models.StakeholderShape{Name: `-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Concept__00000000_.Name = `Gap`
	__Concept__00000000_.ComputedPrefix = ``
	__Concept__00000000_.IsExpanded = false

	__Concept__00000001_.Name = `Function`
	__Concept__00000001_.ComputedPrefix = ``
	__Concept__00000001_.IsExpanded = false

	__Concept__00000002_.Name = `Operational Scenario`
	__Concept__00000002_.ComputedPrefix = ``
	__Concept__00000002_.IsExpanded = false

	__Concept__00000003_.Name = `Mission Requirement`
	__Concept__00000003_.ComputedPrefix = ``
	__Concept__00000003_.IsExpanded = false

	__Concept__00000004_.Name = `Stakeholder Concept`
	__Concept__00000004_.ComputedPrefix = ``
	__Concept__00000004_.IsExpanded = false

	__Concept__00000005_.Name = `Problem`
	__Concept__00000005_.ComputedPrefix = ``
	__Concept__00000005_.IsExpanded = false

	__Concept__00000006_.Name = `Opportunity`
	__Concept__00000006_.ComputedPrefix = ``
	__Concept__00000006_.IsExpanded = false

	__Concept__00000007_.Name = `Solution Class`
	__Concept__00000007_.ComputedPrefix = ``
	__Concept__00000007_.IsExpanded = false

	__Concept__00000008_.Name = `Risk`
	__Concept__00000008_.ComputedPrefix = ``
	__Concept__00000008_.IsExpanded = false

	__Concept__00000009_.Name = `Benefit`
	__Concept__00000009_.ComputedPrefix = ``
	__Concept__00000009_.IsExpanded = false

	__ConceptShape__00000000_.Name = `-Default Diagram`
	__ConceptShape__00000000_.IsExpanded = false
	__ConceptShape__00000000_.X = 1250.000000
	__ConceptShape__00000000_.Y = 10.000000
	__ConceptShape__00000000_.Width = 250.000000
	__ConceptShape__00000000_.Height = 70.000000
	__ConceptShape__00000000_.IsHidden = false

	__ConceptShape__00000001_.Name = `-Default Diagram`
	__ConceptShape__00000001_.IsExpanded = false
	__ConceptShape__00000001_.X = 1250.000000
	__ConceptShape__00000001_.Y = 90.000000
	__ConceptShape__00000001_.Width = 250.000000
	__ConceptShape__00000001_.Height = 70.000000
	__ConceptShape__00000001_.IsHidden = false

	__ConceptShape__00000002_.Name = `-Default Diagram`
	__ConceptShape__00000002_.IsExpanded = false
	__ConceptShape__00000002_.X = 1250.000000
	__ConceptShape__00000002_.Y = 150.000000
	__ConceptShape__00000002_.Width = 250.000000
	__ConceptShape__00000002_.Height = 70.000000
	__ConceptShape__00000002_.IsHidden = false

	__ConceptShape__00000003_.Name = `-Default Diagram`
	__ConceptShape__00000003_.IsExpanded = false
	__ConceptShape__00000003_.X = 1248.000000
	__ConceptShape__00000003_.Y = 231.000000
	__ConceptShape__00000003_.Width = 250.000000
	__ConceptShape__00000003_.Height = 70.000000
	__ConceptShape__00000003_.IsHidden = false

	__ConceptShape__00000004_.Name = `-Default Diagram`
	__ConceptShape__00000004_.IsExpanded = false
	__ConceptShape__00000004_.X = 1248.000000
	__ConceptShape__00000004_.Y = 314.000000
	__ConceptShape__00000004_.Width = 250.000000
	__ConceptShape__00000004_.Height = 70.000000
	__ConceptShape__00000004_.IsHidden = false

	__ConceptShape__00000005_.Name = `-Default Diagram`
	__ConceptShape__00000005_.IsExpanded = false
	__ConceptShape__00000005_.X = 1250.000000
	__ConceptShape__00000005_.Y = 401.000000
	__ConceptShape__00000005_.Width = 250.000000
	__ConceptShape__00000005_.Height = 70.000000
	__ConceptShape__00000005_.IsHidden = false

	__ConceptShape__00000006_.Name = `-Default Diagram`
	__ConceptShape__00000006_.IsExpanded = false
	__ConceptShape__00000006_.X = 1250.000000
	__ConceptShape__00000006_.Y = 500.000000
	__ConceptShape__00000006_.Width = 250.000000
	__ConceptShape__00000006_.Height = 70.000000
	__ConceptShape__00000006_.IsHidden = false

	__ConceptShape__00000007_.Name = `-Default Diagram`
	__ConceptShape__00000007_.IsExpanded = false
	__ConceptShape__00000007_.X = 1250.000000
	__ConceptShape__00000007_.Y = 650.000000
	__ConceptShape__00000007_.Width = 250.000000
	__ConceptShape__00000007_.Height = 70.000000
	__ConceptShape__00000007_.IsHidden = false

	__ConceptShape__00000008_.Name = `-Default Diagram`
	__ConceptShape__00000008_.IsExpanded = false
	__ConceptShape__00000008_.X = 1250.000000
	__ConceptShape__00000008_.Y = 760.000000
	__ConceptShape__00000008_.Width = 250.000000
	__ConceptShape__00000008_.Height = 70.000000
	__ConceptShape__00000008_.IsHidden = false

	__ConceptShape__00000009_.Name = `-Default Diagram`
	__ConceptShape__00000009_.IsExpanded = false
	__ConceptShape__00000009_.X = 1250.000000
	__ConceptShape__00000009_.Y = 840.000000
	__ConceptShape__00000009_.Width = 250.000000
	__ConceptShape__00000009_.Height = 70.000000
	__ConceptShape__00000009_.IsHidden = false

	__Concern__00000000_.Name = `Understand Environment & Mission`
	__Concern__00000000_.IDAirbus = ``
	__Concern__00000000_.Priority = ""
	__Concern__00000000_.ComputedPrefix = `1`
	__Concern__00000000_.IsExpanded = false
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
	__Concern__00000001_.Description = ``
	__Concern__00000001_.IsInputsNodeExpanded = false
	__Concern__00000001_.IsOutputsNodeExpanded = false
	__Concern__00000001_.IsWithCompletion = false
	__Concern__00000001_.Completion = ""

	__Concern__00000002_.Name = `Define Problem or Opportunity Space`
	__Concern__00000002_.IDAirbus = ``
	__Concern__00000002_.Priority = ""
	__Concern__00000002_.ComputedPrefix = `3`
	__Concern__00000002_.IsExpanded = false
	__Concern__00000002_.Description = ``
	__Concern__00000002_.IsInputsNodeExpanded = false
	__Concern__00000002_.IsOutputsNodeExpanded = false
	__Concern__00000002_.IsWithCompletion = false
	__Concern__00000002_.Completion = ""

	__Concern__00000003_.Name = `Characterize Solution Space`
	__Concern__00000003_.IDAirbus = ``
	__Concern__00000003_.Priority = ""
	__Concern__00000003_.ComputedPrefix = `4`
	__Concern__00000003_.IsExpanded = false
	__Concern__00000003_.Description = ``
	__Concern__00000003_.IsInputsNodeExpanded = false
	__Concern__00000003_.IsOutputsNodeExpanded = false
	__Concern__00000003_.IsWithCompletion = false
	__Concern__00000003_.Completion = ""

	__Concern__00000004_.Name = `Assess Business Case`
	__Concern__00000004_.IDAirbus = ``
	__Concern__00000004_.Priority = ""
	__Concern__00000004_.ComputedPrefix = `5`
	__Concern__00000004_.IsExpanded = false
	__Concern__00000004_.Description = ``
	__Concern__00000004_.IsInputsNodeExpanded = false
	__Concern__00000004_.IsOutputsNodeExpanded = false
	__Concern__00000004_.IsWithCompletion = false
	__Concern__00000004_.Completion = ""

	__ConcernOutputShape__00000000_.Name = `Understand Environment & Mission to Gap Analysis`
	__ConcernOutputShape__00000000_.StartRatio = 0.500000
	__ConcernOutputShape__00000000_.EndRatio = 0.500000
	__ConcernOutputShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000000_.CornerOffsetRatio = 1.200000
	__ConcernOutputShape__00000000_.IsHidden = false

	__ConcernOutputShape__00000001_.Name = `Understand Environment & Mission to OpsCon`
	__ConcernOutputShape__00000001_.StartRatio = 0.500000
	__ConcernOutputShape__00000001_.EndRatio = 0.500000
	__ConcernOutputShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000001_.CornerOffsetRatio = 1.200000
	__ConcernOutputShape__00000001_.IsHidden = false

	__ConcernOutputShape__00000002_.Name = `Understand Environment & Mission to Mission Requirements`
	__ConcernOutputShape__00000002_.StartRatio = 0.500000
	__ConcernOutputShape__00000002_.EndRatio = 0.500000
	__ConcernOutputShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000002_.CornerOffsetRatio = 1.200000
	__ConcernOutputShape__00000002_.IsHidden = false

	__ConcernOutputShape__00000003_.Name = `Define Problem or Opportunity Space to Problem or Opportunity Statement`
	__ConcernOutputShape__00000003_.StartRatio = 0.500000
	__ConcernOutputShape__00000003_.EndRatio = 0.500000
	__ConcernOutputShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000003_.CornerOffsetRatio = 1.200000
	__ConcernOutputShape__00000003_.IsHidden = false

	__ConcernOutputShape__00000004_.Name = `Map Stakeholders to Stakeholder Map`
	__ConcernOutputShape__00000004_.StartRatio = 0.500000
	__ConcernOutputShape__00000004_.EndRatio = 0.500000
	__ConcernOutputShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000004_.CornerOffsetRatio = 1.200000
	__ConcernOutputShape__00000004_.IsHidden = false

	__ConcernOutputShape__00000005_.Name = `Characterize Solution Space to Alternative Solution Classes`
	__ConcernOutputShape__00000005_.StartRatio = 0.500000
	__ConcernOutputShape__00000005_.EndRatio = 0.500000
	__ConcernOutputShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000005_.CornerOffsetRatio = 1.200000
	__ConcernOutputShape__00000005_.IsHidden = false

	__ConcernOutputShape__00000006_.Name = `Assess Business Case to Business Case`
	__ConcernOutputShape__00000006_.StartRatio = 0.500000
	__ConcernOutputShape__00000006_.EndRatio = 0.500000
	__ConcernOutputShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ConcernOutputShape__00000006_.CornerOffsetRatio = 1.200000
	__ConcernOutputShape__00000006_.IsHidden = false

	__ConcernShape__00000000_.Name = `-Default Diagram`
	__ConcernShape__00000000_.IsExpanded = false
	__ConcernShape__00000000_.X = 446.000000
	__ConcernShape__00000000_.Y = 201.000000
	__ConcernShape__00000000_.Width = 250.000000
	__ConcernShape__00000000_.Height = 70.000000
	__ConcernShape__00000000_.IsHidden = false

	__ConcernShape__00000001_.Name = `-Default Diagram`
	__ConcernShape__00000001_.IsExpanded = false
	__ConcernShape__00000001_.X = 446.000000
	__ConcernShape__00000001_.Y = 351.000000
	__ConcernShape__00000001_.Width = 250.000000
	__ConcernShape__00000001_.Height = 70.000000
	__ConcernShape__00000001_.IsHidden = false

	__ConcernShape__00000002_.Name = `-Default Diagram`
	__ConcernShape__00000002_.IsExpanded = false
	__ConcernShape__00000002_.X = 450.000000
	__ConcernShape__00000002_.Y = 500.000000
	__ConcernShape__00000002_.Width = 250.000000
	__ConcernShape__00000002_.Height = 70.000000
	__ConcernShape__00000002_.IsHidden = false

	__ConcernShape__00000003_.Name = `-Default Diagram`
	__ConcernShape__00000003_.IsExpanded = false
	__ConcernShape__00000003_.X = 450.000000
	__ConcernShape__00000003_.Y = 650.000000
	__ConcernShape__00000003_.Width = 250.000000
	__ConcernShape__00000003_.Height = 70.000000
	__ConcernShape__00000003_.IsHidden = false

	__ConcernShape__00000004_.Name = `-Default Diagram`
	__ConcernShape__00000004_.IsExpanded = false
	__ConcernShape__00000004_.X = 450.000000
	__ConcernShape__00000004_.Y = 800.000000
	__ConcernShape__00000004_.Width = 250.000000
	__ConcernShape__00000004_.Height = 70.000000
	__ConcernShape__00000004_.IsHidden = false

	__ControlPointShape__00000000_.Name = `Control Point Shape in Problem or Opportunity Statement to Opportunity 0`
	__ControlPointShape__00000000_.X_Relative = 0.061833
	__ControlPointShape__00000000_.Y_Relative = 0.647617
	__ControlPointShape__00000000_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000001_.Name = `Control Point Shape in Problem or Opportunity Statement to Problem 0`
	__ControlPointShape__00000001_.X_Relative = -0.306167
	__ControlPointShape__00000001_.Y_Relative = 0.776189
	__ControlPointShape__00000001_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000002_.Name = `Control Point Shape in Mission Analyst to Define Problem or Opportunity Space 0`
	__ControlPointShape__00000002_.X_Relative = 0.141833
	__ControlPointShape__00000002_.Y_Relative = 0.361905
	__ControlPointShape__00000002_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000003_.Name = `Control Point Shape in Lead System Engineer to Map Stakeholders 0`
	__ControlPointShape__00000003_.X_Relative = 0.049833
	__ControlPointShape__00000003_.Y_Relative = 0.747616
	__ControlPointShape__00000003_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000004_.Name = `Control Point Shape in Enterprise Architect to Characterize Solution Space 0`
	__ControlPointShape__00000004_.X_Relative = 0.073833
	__ControlPointShape__00000004_.Y_Relative = 0.461904
	__ControlPointShape__00000004_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000005_.Name = `Control Point Shape in Problem or Opportunity Statement to Opportunity 0`
	__ControlPointShape__00000005_.X_Relative = 0.085833
	__ControlPointShape__00000005_.Y_Relative = 0.719046
	__ControlPointShape__00000005_.IsStartShapeTheClosestShape = false

	__ControlPointShape__00000006_.Name = `Control Point Shape in Understand Environment & Mission to Mission Requirements 0`
	__ControlPointShape__00000006_.X_Relative = 1.288957
	__ControlPointShape__00000006_.Y_Relative = 1.147615
	__ControlPointShape__00000006_.IsStartShapeTheClosestShape = true

	__ControlPointShape__00000007_.Name = `Control Point Shape in Understand Environment & Mission to Mission Requirements 1`
	__ControlPointShape__00000007_.X_Relative = 0.068958
	__ControlPointShape__00000007_.Y_Relative = 0.633332
	__ControlPointShape__00000007_.IsStartShapeTheClosestShape = false

	__Deliverable__00000000_.Name = `Gap Analysis`
	__Deliverable__00000000_.ComputedPrefix = `1`
	__Deliverable__00000000_.IsExpanded = false
	__Deliverable__00000000_.Description = ``
	__Deliverable__00000000_.IsProducersNodeExpanded = false
	__Deliverable__00000000_.IsConsumersNodeExpanded = false

	__Deliverable__00000001_.Name = `OpsCon`
	__Deliverable__00000001_.ComputedPrefix = `2`
	__Deliverable__00000001_.IsExpanded = false
	__Deliverable__00000001_.Description = ``
	__Deliverable__00000001_.IsProducersNodeExpanded = false
	__Deliverable__00000001_.IsConsumersNodeExpanded = false

	__Deliverable__00000002_.Name = `Mission Requirements`
	__Deliverable__00000002_.ComputedPrefix = `3`
	__Deliverable__00000002_.IsExpanded = false
	__Deliverable__00000002_.Description = ``
	__Deliverable__00000002_.IsProducersNodeExpanded = false
	__Deliverable__00000002_.IsConsumersNodeExpanded = false

	__Deliverable__00000003_.Name = `Stakeholder Map`
	__Deliverable__00000003_.ComputedPrefix = `4`
	__Deliverable__00000003_.IsExpanded = false
	__Deliverable__00000003_.Description = ``
	__Deliverable__00000003_.IsProducersNodeExpanded = false
	__Deliverable__00000003_.IsConsumersNodeExpanded = false

	__Deliverable__00000004_.Name = `Problem or Opportunity Statement`
	__Deliverable__00000004_.ComputedPrefix = `5`
	__Deliverable__00000004_.IsExpanded = false
	__Deliverable__00000004_.Description = ``
	__Deliverable__00000004_.IsProducersNodeExpanded = false
	__Deliverable__00000004_.IsConsumersNodeExpanded = false

	__Deliverable__00000005_.Name = `Alternative Solution Classes`
	__Deliverable__00000005_.ComputedPrefix = `6`
	__Deliverable__00000005_.IsExpanded = false
	__Deliverable__00000005_.Description = ``
	__Deliverable__00000005_.IsProducersNodeExpanded = false
	__Deliverable__00000005_.IsConsumersNodeExpanded = false

	__Deliverable__00000006_.Name = `Business Case`
	__Deliverable__00000006_.ComputedPrefix = `7`
	__Deliverable__00000006_.IsExpanded = false
	__Deliverable__00000006_.Description = ``
	__Deliverable__00000006_.IsProducersNodeExpanded = false
	__Deliverable__00000006_.IsConsumersNodeExpanded = false

	__DeliverableConceptShape__00000000_.Name = `Gap Analysis to Gap`
	__DeliverableConceptShape__00000000_.StartRatio = 0.000000
	__DeliverableConceptShape__00000000_.EndRatio = 0.000000
	__DeliverableConceptShape__00000000_.StartOrientation = ""
	__DeliverableConceptShape__00000000_.EndOrientation = ""
	__DeliverableConceptShape__00000000_.CornerOffsetRatio = 0.000000
	__DeliverableConceptShape__00000000_.IsHidden = false

	__DeliverableConceptShape__00000001_.Name = `Gap Analysis to Function`
	__DeliverableConceptShape__00000001_.StartRatio = 0.000000
	__DeliverableConceptShape__00000001_.EndRatio = 0.000000
	__DeliverableConceptShape__00000001_.StartOrientation = ""
	__DeliverableConceptShape__00000001_.EndOrientation = ""
	__DeliverableConceptShape__00000001_.CornerOffsetRatio = 0.000000
	__DeliverableConceptShape__00000001_.IsHidden = false

	__DeliverableConceptShape__00000002_.Name = `OpsCon to Operational Scenario`
	__DeliverableConceptShape__00000002_.StartRatio = 0.000000
	__DeliverableConceptShape__00000002_.EndRatio = 0.000000
	__DeliverableConceptShape__00000002_.StartOrientation = ""
	__DeliverableConceptShape__00000002_.EndOrientation = ""
	__DeliverableConceptShape__00000002_.CornerOffsetRatio = 0.000000
	__DeliverableConceptShape__00000002_.IsHidden = false

	__DeliverableConceptShape__00000003_.Name = `Mission Requirements to Mission Requirement`
	__DeliverableConceptShape__00000003_.StartRatio = 0.000000
	__DeliverableConceptShape__00000003_.EndRatio = 0.000000
	__DeliverableConceptShape__00000003_.StartOrientation = ""
	__DeliverableConceptShape__00000003_.EndOrientation = ""
	__DeliverableConceptShape__00000003_.CornerOffsetRatio = 0.000000
	__DeliverableConceptShape__00000003_.IsHidden = false

	__DeliverableConceptShape__00000004_.Name = `Problem or Opportunity Statement to Problem`
	__DeliverableConceptShape__00000004_.StartRatio = 0.000000
	__DeliverableConceptShape__00000004_.EndRatio = 0.000000
	__DeliverableConceptShape__00000004_.StartOrientation = ""
	__DeliverableConceptShape__00000004_.EndOrientation = ""
	__DeliverableConceptShape__00000004_.CornerOffsetRatio = 0.000000
	__DeliverableConceptShape__00000004_.IsHidden = false

	__DeliverableConceptShape__00000005_.Name = `Problem or Opportunity Statement to Opportunity`
	__DeliverableConceptShape__00000005_.StartRatio = 0.000000
	__DeliverableConceptShape__00000005_.EndRatio = 0.000000
	__DeliverableConceptShape__00000005_.StartOrientation = ""
	__DeliverableConceptShape__00000005_.EndOrientation = ""
	__DeliverableConceptShape__00000005_.CornerOffsetRatio = 0.000000
	__DeliverableConceptShape__00000005_.IsHidden = false

	__DeliverableConceptShape__00000006_.Name = `Stakeholder Map to Stakeholder Concept`
	__DeliverableConceptShape__00000006_.StartRatio = 0.000000
	__DeliverableConceptShape__00000006_.EndRatio = 0.000000
	__DeliverableConceptShape__00000006_.StartOrientation = ""
	__DeliverableConceptShape__00000006_.EndOrientation = ""
	__DeliverableConceptShape__00000006_.CornerOffsetRatio = 0.000000
	__DeliverableConceptShape__00000006_.IsHidden = false

	__DeliverableConceptShape__00000007_.Name = `Alternative Solution Classes to Solution Class`
	__DeliverableConceptShape__00000007_.StartRatio = 0.000000
	__DeliverableConceptShape__00000007_.EndRatio = 0.000000
	__DeliverableConceptShape__00000007_.StartOrientation = ""
	__DeliverableConceptShape__00000007_.EndOrientation = ""
	__DeliverableConceptShape__00000007_.CornerOffsetRatio = 0.000000
	__DeliverableConceptShape__00000007_.IsHidden = false

	__DeliverableConceptShape__00000008_.Name = `Business Case to Risk`
	__DeliverableConceptShape__00000008_.StartRatio = 0.000000
	__DeliverableConceptShape__00000008_.EndRatio = 0.000000
	__DeliverableConceptShape__00000008_.StartOrientation = ""
	__DeliverableConceptShape__00000008_.EndOrientation = ""
	__DeliverableConceptShape__00000008_.CornerOffsetRatio = 0.000000
	__DeliverableConceptShape__00000008_.IsHidden = false

	__DeliverableConceptShape__00000009_.Name = `Business Case to Benefit`
	__DeliverableConceptShape__00000009_.StartRatio = 0.000000
	__DeliverableConceptShape__00000009_.EndRatio = 0.000000
	__DeliverableConceptShape__00000009_.StartOrientation = ""
	__DeliverableConceptShape__00000009_.EndOrientation = ""
	__DeliverableConceptShape__00000009_.CornerOffsetRatio = 0.000000
	__DeliverableConceptShape__00000009_.IsHidden = false

	__DeliverableShape__00000000_.Name = `-Default Diagram`
	__DeliverableShape__00000000_.IsExpanded = false
	__DeliverableShape__00000000_.X = 850.000000
	__DeliverableShape__00000000_.Y = 50.000000
	__DeliverableShape__00000000_.Width = 250.000000
	__DeliverableShape__00000000_.Height = 70.000000
	__DeliverableShape__00000000_.IsHidden = false

	__DeliverableShape__00000001_.Name = `-Default Diagram`
	__DeliverableShape__00000001_.IsExpanded = false
	__DeliverableShape__00000001_.X = 850.000000
	__DeliverableShape__00000001_.Y = 150.000000
	__DeliverableShape__00000001_.Width = 250.000000
	__DeliverableShape__00000001_.Height = 70.000000
	__DeliverableShape__00000001_.IsHidden = false

	__DeliverableShape__00000002_.Name = `-Default Diagram`
	__DeliverableShape__00000002_.IsExpanded = false
	__DeliverableShape__00000002_.X = 850.000000
	__DeliverableShape__00000002_.Y = 250.000000
	__DeliverableShape__00000002_.Width = 250.000000
	__DeliverableShape__00000002_.Height = 70.000000
	__DeliverableShape__00000002_.IsHidden = false

	__DeliverableShape__00000003_.Name = `-Default Diagram`
	__DeliverableShape__00000003_.IsExpanded = false
	__DeliverableShape__00000003_.X = 850.000000
	__DeliverableShape__00000003_.Y = 350.000000
	__DeliverableShape__00000003_.Width = 250.000000
	__DeliverableShape__00000003_.Height = 70.000000
	__DeliverableShape__00000003_.IsHidden = false

	__DeliverableShape__00000004_.Name = `-Default Diagram`
	__DeliverableShape__00000004_.IsExpanded = false
	__DeliverableShape__00000004_.X = 850.000000
	__DeliverableShape__00000004_.Y = 500.000000
	__DeliverableShape__00000004_.Width = 250.000000
	__DeliverableShape__00000004_.Height = 70.000000
	__DeliverableShape__00000004_.IsHidden = false

	__DeliverableShape__00000005_.Name = `-Default Diagram`
	__DeliverableShape__00000005_.IsExpanded = false
	__DeliverableShape__00000005_.X = 850.000000
	__DeliverableShape__00000005_.Y = 650.000000
	__DeliverableShape__00000005_.Width = 250.000000
	__DeliverableShape__00000005_.Height = 70.000000
	__DeliverableShape__00000005_.IsHidden = false

	__DeliverableShape__00000006_.Name = `-Default Diagram`
	__DeliverableShape__00000006_.IsExpanded = false
	__DeliverableShape__00000006_.X = 850.000000
	__DeliverableShape__00000006_.Y = 800.000000
	__DeliverableShape__00000006_.Width = 250.000000
	__DeliverableShape__00000006_.Height = 70.000000
	__DeliverableShape__00000006_.IsHidden = false

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = false
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 14400.000000
	__Diagram__00000000_.Height = 14400.000000
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
	__Library__00000000_.NbPixPerCharacter = 8.000000

	__Stakeholder__00000000_.Name = `Mission Analyst`
	__Stakeholder__00000000_.IDAirbus = ``
	__Stakeholder__00000000_.ComputedPrefix = `1`
	__Stakeholder__00000000_.IsExpanded = false
	__Stakeholder__00000000_.Description = ``

	__Stakeholder__00000001_.Name = `Enterprise Architect`
	__Stakeholder__00000001_.IDAirbus = ``
	__Stakeholder__00000001_.ComputedPrefix = `2`
	__Stakeholder__00000001_.IsExpanded = false
	__Stakeholder__00000001_.Description = ``

	__Stakeholder__00000002_.Name = `Lead System Engineer`
	__Stakeholder__00000002_.IDAirbus = ``
	__Stakeholder__00000002_.ComputedPrefix = `3`
	__Stakeholder__00000002_.IsExpanded = false
	__Stakeholder__00000002_.Description = ``

	__Stakeholder__00000003_.Name = `Executive Sponsor`
	__Stakeholder__00000003_.IDAirbus = ``
	__Stakeholder__00000003_.ComputedPrefix = `4`
	__Stakeholder__00000003_.IsExpanded = false
	__Stakeholder__00000003_.Description = ``

	__Stakeholder__00000004_.Name = `Business Analyst`
	__Stakeholder__00000004_.IDAirbus = ``
	__Stakeholder__00000004_.ComputedPrefix = `5`
	__Stakeholder__00000004_.IsExpanded = false
	__Stakeholder__00000004_.Description = ``

	__StakeholderConcernShape__00000000_.Name = `Mission Analyst to Understand Environment & Mission`
	__StakeholderConcernShape__00000000_.StartRatio = 0.500000
	__StakeholderConcernShape__00000000_.EndRatio = 0.500000
	__StakeholderConcernShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000000_.CornerOffsetRatio = 1.200000
	__StakeholderConcernShape__00000000_.IsHidden = false

	__StakeholderConcernShape__00000001_.Name = `Mission Analyst to Define Problem or Opportunity Space`
	__StakeholderConcernShape__00000001_.StartRatio = 0.500000
	__StakeholderConcernShape__00000001_.EndRatio = 0.500000
	__StakeholderConcernShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000001_.CornerOffsetRatio = 1.200000
	__StakeholderConcernShape__00000001_.IsHidden = false

	__StakeholderConcernShape__00000002_.Name = `Business Analyst to Define Problem or Opportunity Space`
	__StakeholderConcernShape__00000002_.StartRatio = 0.500000
	__StakeholderConcernShape__00000002_.EndRatio = 0.500000
	__StakeholderConcernShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000002_.CornerOffsetRatio = 1.200000
	__StakeholderConcernShape__00000002_.IsHidden = false

	__StakeholderConcernShape__00000003_.Name = `Enterprise Architect to Map Stakeholders`
	__StakeholderConcernShape__00000003_.StartRatio = 0.500000
	__StakeholderConcernShape__00000003_.EndRatio = 0.500000
	__StakeholderConcernShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000003_.CornerOffsetRatio = 1.200000
	__StakeholderConcernShape__00000003_.IsHidden = false

	__StakeholderConcernShape__00000004_.Name = `Enterprise Architect to Characterize Solution Space`
	__StakeholderConcernShape__00000004_.StartRatio = 0.500000
	__StakeholderConcernShape__00000004_.EndRatio = 0.500000
	__StakeholderConcernShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000004_.CornerOffsetRatio = 1.300000
	__StakeholderConcernShape__00000004_.IsHidden = false

	__StakeholderConcernShape__00000005_.Name = `Lead System Engineer to Map Stakeholders`
	__StakeholderConcernShape__00000005_.StartRatio = 0.500000
	__StakeholderConcernShape__00000005_.EndRatio = 0.500000
	__StakeholderConcernShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000005_.CornerOffsetRatio = 1.400000
	__StakeholderConcernShape__00000005_.IsHidden = false

	__StakeholderConcernShape__00000006_.Name = `Lead System Engineer to Characterize Solution Space`
	__StakeholderConcernShape__00000006_.StartRatio = 0.500000
	__StakeholderConcernShape__00000006_.EndRatio = 0.500000
	__StakeholderConcernShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000006_.CornerOffsetRatio = 1.200000
	__StakeholderConcernShape__00000006_.IsHidden = false

	__StakeholderConcernShape__00000007_.Name = `Executive Sponsor to Assess Business Case`
	__StakeholderConcernShape__00000007_.StartRatio = 0.500000
	__StakeholderConcernShape__00000007_.EndRatio = 0.500000
	__StakeholderConcernShape__00000007_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__StakeholderConcernShape__00000007_.CornerOffsetRatio = 1.200000
	__StakeholderConcernShape__00000007_.IsHidden = false

	__StakeholderShape__00000000_.Name = `-Default Diagram`
	__StakeholderShape__00000000_.IsExpanded = false
	__StakeholderShape__00000000_.X = 50.000000
	__StakeholderShape__00000000_.Y = 200.000000
	__StakeholderShape__00000000_.Width = 250.000000
	__StakeholderShape__00000000_.Height = 70.000000
	__StakeholderShape__00000000_.IsHidden = false

	__StakeholderShape__00000001_.Name = `-Default Diagram`
	__StakeholderShape__00000001_.IsExpanded = false
	__StakeholderShape__00000001_.X = 50.000000
	__StakeholderShape__00000001_.Y = 350.000000
	__StakeholderShape__00000001_.Width = 250.000000
	__StakeholderShape__00000001_.Height = 70.000000
	__StakeholderShape__00000001_.IsHidden = false

	__StakeholderShape__00000002_.Name = `-Default Diagram`
	__StakeholderShape__00000002_.IsExpanded = false
	__StakeholderShape__00000002_.X = 50.000000
	__StakeholderShape__00000002_.Y = 500.000000
	__StakeholderShape__00000002_.Width = 250.000000
	__StakeholderShape__00000002_.Height = 70.000000
	__StakeholderShape__00000002_.IsHidden = false

	__StakeholderShape__00000003_.Name = `-Default Diagram`
	__StakeholderShape__00000003_.IsExpanded = false
	__StakeholderShape__00000003_.X = 50.000000
	__StakeholderShape__00000003_.Y = 650.000000
	__StakeholderShape__00000003_.Width = 250.000000
	__StakeholderShape__00000003_.Height = 70.000000
	__StakeholderShape__00000003_.IsHidden = false

	__StakeholderShape__00000004_.Name = `-Default Diagram`
	__StakeholderShape__00000004_.IsExpanded = false
	__StakeholderShape__00000004_.X = 50.000000
	__StakeholderShape__00000004_.Y = 800.000000
	__StakeholderShape__00000004_.Width = 250.000000
	__StakeholderShape__00000004_.Height = 70.000000
	__StakeholderShape__00000004_.IsHidden = false

	// insertion point for setup of pointers
	__ConceptShape__00000000_.Concept = __Concept__00000000_
	__ConceptShape__00000001_.Concept = __Concept__00000001_
	__ConceptShape__00000002_.Concept = __Concept__00000002_
	__ConceptShape__00000003_.Concept = __Concept__00000003_
	__ConceptShape__00000004_.Concept = __Concept__00000005_
	__ConceptShape__00000005_.Concept = __Concept__00000006_
	__ConceptShape__00000006_.Concept = __Concept__00000004_
	__ConceptShape__00000007_.Concept = __Concept__00000007_
	__ConceptShape__00000008_.Concept = __Concept__00000008_
	__ConceptShape__00000009_.Concept = __Concept__00000009_
	__Concern__00000000_.Outputs = append(__Concern__00000000_.Outputs, __Deliverable__00000000_)
	__Concern__00000000_.Outputs = append(__Concern__00000000_.Outputs, __Deliverable__00000001_)
	__Concern__00000000_.Outputs = append(__Concern__00000000_.Outputs, __Deliverable__00000002_)
	__Concern__00000001_.Outputs = append(__Concern__00000001_.Outputs, __Deliverable__00000003_)
	__Concern__00000002_.Outputs = append(__Concern__00000002_.Outputs, __Deliverable__00000004_)
	__Concern__00000003_.Outputs = append(__Concern__00000003_.Outputs, __Deliverable__00000005_)
	__Concern__00000004_.Outputs = append(__Concern__00000004_.Outputs, __Deliverable__00000006_)
	__ConcernOutputShape__00000000_.Concern = __Concern__00000000_
	__ConcernOutputShape__00000000_.Deliverable = __Deliverable__00000000_
	__ConcernOutputShape__00000001_.Concern = __Concern__00000000_
	__ConcernOutputShape__00000001_.Deliverable = __Deliverable__00000001_
	__ConcernOutputShape__00000002_.Concern = __Concern__00000000_
	__ConcernOutputShape__00000002_.Deliverable = __Deliverable__00000002_
	__ConcernOutputShape__00000002_.ControlPointShapes = append(__ConcernOutputShape__00000002_.ControlPointShapes, __ControlPointShape__00000007_)
	__ConcernOutputShape__00000003_.Concern = __Concern__00000002_
	__ConcernOutputShape__00000003_.Deliverable = __Deliverable__00000004_
	__ConcernOutputShape__00000004_.Concern = __Concern__00000001_
	__ConcernOutputShape__00000004_.Deliverable = __Deliverable__00000003_
	__ConcernOutputShape__00000005_.Concern = __Concern__00000003_
	__ConcernOutputShape__00000005_.Deliverable = __Deliverable__00000005_
	__ConcernOutputShape__00000006_.Concern = __Concern__00000004_
	__ConcernOutputShape__00000006_.Deliverable = __Deliverable__00000006_
	__ConcernShape__00000000_.Concern = __Concern__00000000_
	__ConcernShape__00000001_.Concern = __Concern__00000002_
	__ConcernShape__00000002_.Concern = __Concern__00000001_
	__ConcernShape__00000003_.Concern = __Concern__00000003_
	__ConcernShape__00000004_.Concern = __Concern__00000004_
	__Deliverable__00000000_.Concepts = append(__Deliverable__00000000_.Concepts, __Concept__00000000_)
	__Deliverable__00000000_.Concepts = append(__Deliverable__00000000_.Concepts, __Concept__00000001_)
	__Deliverable__00000001_.Concepts = append(__Deliverable__00000001_.Concepts, __Concept__00000002_)
	__Deliverable__00000002_.Concepts = append(__Deliverable__00000002_.Concepts, __Concept__00000003_)
	__Deliverable__00000003_.Concepts = append(__Deliverable__00000003_.Concepts, __Concept__00000004_)
	__Deliverable__00000004_.Concepts = append(__Deliverable__00000004_.Concepts, __Concept__00000005_)
	__Deliverable__00000004_.Concepts = append(__Deliverable__00000004_.Concepts, __Concept__00000006_)
	__Deliverable__00000005_.Concepts = append(__Deliverable__00000005_.Concepts, __Concept__00000007_)
	__Deliverable__00000006_.Concepts = append(__Deliverable__00000006_.Concepts, __Concept__00000008_)
	__Deliverable__00000006_.Concepts = append(__Deliverable__00000006_.Concepts, __Concept__00000009_)
	__DeliverableConceptShape__00000000_.Deliverable = __Deliverable__00000000_
	__DeliverableConceptShape__00000000_.Concept = __Concept__00000000_
	__DeliverableConceptShape__00000001_.Deliverable = __Deliverable__00000000_
	__DeliverableConceptShape__00000001_.Concept = __Concept__00000001_
	__DeliverableConceptShape__00000002_.Deliverable = __Deliverable__00000001_
	__DeliverableConceptShape__00000002_.Concept = __Concept__00000002_
	__DeliverableConceptShape__00000003_.Deliverable = __Deliverable__00000002_
	__DeliverableConceptShape__00000003_.Concept = __Concept__00000003_
	__DeliverableConceptShape__00000004_.Deliverable = __Deliverable__00000004_
	__DeliverableConceptShape__00000004_.Concept = __Concept__00000005_
	__DeliverableConceptShape__00000005_.Deliverable = __Deliverable__00000004_
	__DeliverableConceptShape__00000005_.Concept = __Concept__00000006_
	__DeliverableConceptShape__00000005_.ControlPointShapes = append(__DeliverableConceptShape__00000005_.ControlPointShapes, __ControlPointShape__00000005_)
	__DeliverableConceptShape__00000006_.Deliverable = __Deliverable__00000003_
	__DeliverableConceptShape__00000006_.Concept = __Concept__00000004_
	__DeliverableConceptShape__00000007_.Deliverable = __Deliverable__00000005_
	__DeliverableConceptShape__00000007_.Concept = __Concept__00000007_
	__DeliverableConceptShape__00000008_.Deliverable = __Deliverable__00000006_
	__DeliverableConceptShape__00000008_.Concept = __Concept__00000008_
	__DeliverableConceptShape__00000009_.Deliverable = __Deliverable__00000006_
	__DeliverableConceptShape__00000009_.Concept = __Concept__00000009_
	__DeliverableShape__00000000_.Deliverable = __Deliverable__00000000_
	__DeliverableShape__00000001_.Deliverable = __Deliverable__00000001_
	__DeliverableShape__00000002_.Deliverable = __Deliverable__00000002_
	__DeliverableShape__00000003_.Deliverable = __Deliverable__00000004_
	__DeliverableShape__00000004_.Deliverable = __Deliverable__00000003_
	__DeliverableShape__00000005_.Deliverable = __Deliverable__00000005_
	__DeliverableShape__00000006_.Deliverable = __Deliverable__00000006_
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000000_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000001_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000002_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000003_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000004_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000005_)
	__Diagram__00000000_.Deliverable_Shapes = append(__Diagram__00000000_.Deliverable_Shapes, __DeliverableShape__00000006_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000000_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000001_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000002_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000003_)
	__Diagram__00000000_.Concern_Shapes = append(__Diagram__00000000_.Concern_Shapes, __ConcernShape__00000004_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000000_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000001_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000002_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000003_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000004_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000005_)
	__Diagram__00000000_.ConcernOutputShapes = append(__Diagram__00000000_.ConcernOutputShapes, __ConcernOutputShape__00000006_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000000_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000001_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000002_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000003_)
	__Diagram__00000000_.Stakeholder_Shapes = append(__Diagram__00000000_.Stakeholder_Shapes, __StakeholderShape__00000004_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000000_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000001_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000002_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000003_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000004_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000005_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000006_)
	__Diagram__00000000_.StakeholderConcernShapes = append(__Diagram__00000000_.StakeholderConcernShapes, __StakeholderConcernShape__00000007_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000000_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000001_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000002_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000003_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000004_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000005_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000006_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000007_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000008_)
	__Diagram__00000000_.Concept_Shapes = append(__Diagram__00000000_.Concept_Shapes, __ConceptShape__00000009_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000000_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000001_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000002_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000003_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000004_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000005_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000006_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000007_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000008_)
	__Diagram__00000000_.DeliverableConceptShapes = append(__Diagram__00000000_.DeliverableConceptShapes, __DeliverableConceptShape__00000009_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000000_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000001_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000002_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000003_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000004_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000005_)
	__Library__00000000_.RootDeliverables = append(__Library__00000000_.RootDeliverables, __Deliverable__00000006_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000000_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000001_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000002_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000003_)
	__Library__00000000_.RootConcerns = append(__Library__00000000_.RootConcerns, __Concern__00000004_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000000_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000001_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000002_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000003_)
	__Library__00000000_.RootStakeholders = append(__Library__00000000_.RootStakeholders, __Stakeholder__00000004_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000000_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000001_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000002_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000003_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000004_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000005_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000006_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000007_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000008_)
	__Library__00000000_.RootConcepts = append(__Library__00000000_.RootConcepts, __Concept__00000009_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Stakeholder__00000000_.Concerns = append(__Stakeholder__00000000_.Concerns, __Concern__00000000_)
	__Stakeholder__00000000_.Concerns = append(__Stakeholder__00000000_.Concerns, __Concern__00000002_)
	__Stakeholder__00000001_.Concerns = append(__Stakeholder__00000001_.Concerns, __Concern__00000001_)
	__Stakeholder__00000001_.Concerns = append(__Stakeholder__00000001_.Concerns, __Concern__00000003_)
	__Stakeholder__00000002_.Concerns = append(__Stakeholder__00000002_.Concerns, __Concern__00000001_)
	__Stakeholder__00000002_.Concerns = append(__Stakeholder__00000002_.Concerns, __Concern__00000003_)
	__Stakeholder__00000003_.Concerns = append(__Stakeholder__00000003_.Concerns, __Concern__00000004_)
	__Stakeholder__00000004_.Concerns = append(__Stakeholder__00000004_.Concerns, __Concern__00000002_)
	__StakeholderConcernShape__00000000_.Stakeholder = __Stakeholder__00000000_
	__StakeholderConcernShape__00000000_.Concern = __Concern__00000000_
	__StakeholderConcernShape__00000001_.Stakeholder = __Stakeholder__00000000_
	__StakeholderConcernShape__00000001_.Concern = __Concern__00000002_
	__StakeholderConcernShape__00000001_.ControlPointShapes = append(__StakeholderConcernShape__00000001_.ControlPointShapes, __ControlPointShape__00000002_)
	__StakeholderConcernShape__00000002_.Stakeholder = __Stakeholder__00000004_
	__StakeholderConcernShape__00000002_.Concern = __Concern__00000002_
	__StakeholderConcernShape__00000003_.Stakeholder = __Stakeholder__00000001_
	__StakeholderConcernShape__00000003_.Concern = __Concern__00000001_
	__StakeholderConcernShape__00000004_.Stakeholder = __Stakeholder__00000001_
	__StakeholderConcernShape__00000004_.Concern = __Concern__00000003_
	__StakeholderConcernShape__00000004_.ControlPointShapes = append(__StakeholderConcernShape__00000004_.ControlPointShapes, __ControlPointShape__00000004_)
	__StakeholderConcernShape__00000005_.Stakeholder = __Stakeholder__00000002_
	__StakeholderConcernShape__00000005_.Concern = __Concern__00000001_
	__StakeholderConcernShape__00000005_.ControlPointShapes = append(__StakeholderConcernShape__00000005_.ControlPointShapes, __ControlPointShape__00000003_)
	__StakeholderConcernShape__00000006_.Stakeholder = __Stakeholder__00000002_
	__StakeholderConcernShape__00000006_.Concern = __Concern__00000003_
	__StakeholderConcernShape__00000007_.Stakeholder = __Stakeholder__00000003_
	__StakeholderConcernShape__00000007_.Concern = __Concern__00000004_
	__StakeholderShape__00000000_.Stakeholder = __Stakeholder__00000000_
	__StakeholderShape__00000001_.Stakeholder = __Stakeholder__00000004_
	__StakeholderShape__00000002_.Stakeholder = __Stakeholder__00000001_
	__StakeholderShape__00000003_.Stakeholder = __Stakeholder__00000002_
	__StakeholderShape__00000004_.Stakeholder = __Stakeholder__00000003_
}
