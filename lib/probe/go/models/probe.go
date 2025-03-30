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
	GetName()
	GetType()
	GetNamedStructs() []NamedStruct
	GetOrder(instance Instance)
}

// Called Probe2 for the moment because the legacy Probe name collision
type Probe2 struct {
	Name       string // not even sure we need a named struct (parametrization ?)
	stage      *Stage
	splitStage *split.Stage // root
}

func NewProbe2(r *gin.Engine, stage *Stage) (probe *Probe2) {

	probe = new(Probe2)

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	probe.splitStage = split_stack.NewStack(r, stage.GetName(), "", "", "", true, false).Stage

	(&split.View{
		Name: "Probe 2",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
			}).Stage(probe.splitStage),
			(&split.AsSplitArea{
				Size: 50,
			}).Stage(probe.splitStage),
		},
	}).Stage(probe.splitStage)

	probe.splitStage.Commit()

	return
}
