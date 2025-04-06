package models

import (
	"embed"

	"github.com/gin-gonic/gin"

	gong "github.com/fullstack-lang/gong/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	form "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gongdoc_load "github.com/fullstack-lang/gong/lib/doc/go/load"
	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	gongtree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"
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

type NamedStructInterface interface {
	GetName() string
	// GetFields() []Field
	// GetInstances() []Instance
	// NewInstance(stage ProbebStage) Instance
	// Unstage(stage ProbebStage, instance Instance)
}

type ProbebStage interface {
	GetName() string
	GetType() string
	GetProbeSplitStageName() string
	GetProbeTreeSidebarStageName() string
	GetProbeTableStageName() string
	GetProbeFormStageName() string
	GetMap_GongStructName_InstancesNb() map[string]int

	GetModelsEmbededDir() embed.FS
	GetDigramsEmbededDir() embed.FS
	GetNamedStructsNames() []string
	GetNamedStructNamesByOrder(namedStuctName string) []string
	// GetOrder(instance Instance)
}

// Called Probe2 for the moment because the legacy Probe name collision
type Probe2 struct {
	Name  string // not even sure we need a named struct (parametrization ?)
	stage *Stage

	stageOfInterest ProbebStage

	splitStage *split.Stage // split stage of the probe

	gongStage  *gong.Stage
	treeStage  *tree.Stage
	formStage  *form.Stage
	tableStage *form.Stage
}

func (probe *Probe2) GetSplitStage() (splitStage *split.Stage) {
	return probe.splitStage
}

const TableName = "Table"
const FormName = "Form"

func NewProbe2(r *gin.Engine, stage *Stage, stageOfInterest ProbebStage, embeddedDiagrams bool) (probe *Probe2) {

	probe = (&Probe2{
		Name:            stage.name,
		stage:           stage,
		stageOfInterest: stageOfInterest,
	}).Stage(stage)

	gongdoc_load.Load(
		"",
		probe.stageOfInterest.GetProbeSplitStageName(),
		probe.stageOfInterest.GetModelsEmbededDir(),
		probe.stageOfInterest.GetDigramsEmbededDir(),
		r,
		embeddedDiagrams,
		stageOfInterest.GetMap_GongStructName_InstancesNb())

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	probe.splitStage = split_stack.NewStack(r, stage.GetName(), "", "", "", true, false).Stage

	probe.gongStage, _ = gong_fullstack.NewStackInstance(r, stageOfInterest.GetName())
	gong.LoadEmbedded(probe.gongStage, stageOfInterest.GetModelsEmbededDir())

	// treeForSelectingDate that is on the sidebar
	probe.treeStage, _ = gongtree_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTreeSidebarStageName()+"2")

	// stage for main table
	probe.tableStage, _ = gongtable_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTableStageName()+"2")

	// stage for reusable form
	probe.formStage, _ = gongtable_fullstack.NewStackInstance(r, stageOfInterest.GetProbeFormStageName()+"2")

	name := stageOfInterest.GetName()
	_ = name

	(&split.View{
		Name:         "Probe of stage type: \"" + stageOfInterest.GetType() + "\", and name \": " + stageOfInterest.GetName() + "\"",
		ShowViewName: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name:             "editor",
				ShowNameInHeader: false,
				Size:             50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						(&split.AsSplitArea{
							Name:             "sidebar",
							ShowNameInHeader: false,
							Size:             20,
							Tree: (&split.Tree{
								Name:      "Sidebar",
								StackName: probe.treeStage.GetName(),
								TreeName:  SideBarTreeName,
							}).Stage(probe.splitStage),
						}).Stage(probe.splitStage),
						(&split.AsSplitArea{
							Name:             "table",
							ShowNameInHeader: false,
							Size:             50,
							Table: (&split.Table{
								Name:      "Table",
								StackName: probe.tableStage.GetName(),
								TableName: TableName,
							}).Stage(probe.splitStage),
						}).Stage(probe.splitStage),
						(&split.AsSplitArea{
							Name:             "form",
							ShowNameInHeader: false,
							Size:             30,
							Form: (&split.Form{
								Name:      "Form",
								StackName: probe.formStage.GetName(),
								FormName:  FormName,
							}).Stage(probe.splitStage),
						}).Stage(probe.splitStage),
					},
				}).Stage(probe.splitStage),
			}).Stage(probe.splitStage),
			(&split.AsSplitArea{
				Name:             "for doc stage",
				ShowNameInHeader: false,
				Size:             50,
				Doc: (&split.Doc{
					Name:      "Doc",
					StackName: stageOfInterest.GetProbeSplitStageName(),
				}).Stage(probe.splitStage),
			}).Stage(probe.splitStage),
		},
	}).Stage(probe.splitStage)

	probe.splitStage.Commit()
	probe.fillUpTree()

	return
}

func (probe *Probe2) Refresh() {
	probe.fillUpTree()
}
