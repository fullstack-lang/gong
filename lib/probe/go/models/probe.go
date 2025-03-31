package models

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

type FieldType int

const (
	FieldTypeInt FieldType = iota
	FieldTypeFloat64
	FieldTypeBool
	FieldTypeTime
	FieldTypeDuration
	FieldTypePointer
	FieldTypeSliceOfPointer
)

type FieldValue interface {
	GetType() FieldType
	GetValue()
}

type Instance interface {
	GetName()
	GetFieldValues() []FieldValue
	GetFieldValue(field Field) []FieldValue
	SetFieldValue(field Field, fieldValue FieldValue)
}

type Field interface {
	GetName()
}

type NamedStruct interface {
	GetName()
	GetFields() []Field
	GetInstances() []Instance
	NewInstance(stage ProbebStage) Instance
	Unstage(stage ProbebStage, instance Instance)
}

type ProbebStage interface {
	GetName() string
	GetType() string
	GetProbeSplitStageName() string
	// GetNamedStructs() []NamedStruct
	// GetOrder(instance Instance)
}

// Called Probe2 for the moment because the legacy Probe name collision
type Probe2 struct {
	Name        string // not even sure we need a named struct (parametrization ?)
	stage       *Stage
	splitStage  *split.Stage // split stage of the probe
	probedStage ProbebStage
}

func (probe *Probe2) GetSplitStage() (splitStage *split.Stage) {
	return probe.splitStage
}

func NewProbe2(r *gin.Engine, stage *Stage, probedStage ProbebStage) (probe *Probe2) {

	probe = (&Probe2{
		Name:        stage.name,
		stage:       stage,
		probedStage: probedStage,
	}).Stage(stage)

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	probe.splitStage = split_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage

	name := probedStage.GetName()
	_ = name

	(&split.View{
		Name:         "Probe of stage type: \"" + probedStage.GetType() + "\", and name \": " + probedStage.GetName() + "\"",
		ShowViewName: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name:             "editor",
				ShowNameInHeader: true,
				Size:             50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						(&split.AsSplitArea{
							Name:             "sidebar",
							ShowNameInHeader: true,
							Size:             20,
						}).Stage(probe.splitStage),
						(&split.AsSplitArea{
							Name:             "table",
							ShowNameInHeader: true,
							Size:             50,
						}).Stage(probe.splitStage),
						(&split.AsSplitArea{
							Name:             "form",
							ShowNameInHeader: true,
							Size:             30,
						}).Stage(probe.splitStage),
					},
				}).Stage(probe.splitStage),
			}).Stage(probe.splitStage),
			(&split.AsSplitArea{
				Name:             "for doc stage",
				ShowNameInHeader: true,
				Size:             50,
				Doc: (&split.Doc{
					Name:      "Doc",
					StackName: probedStage.GetProbeSplitStageName(),
				}).Stage(probe.splitStage),
			}).Stage(probe.splitStage),
		},
	}).Stage(probe.splitStage)

	probe.splitStage.Commit()

	return
}
