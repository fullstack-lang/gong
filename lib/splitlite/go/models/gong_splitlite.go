package models

import (
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

type GongProbeIF interface {
	Refresh()
	GetFormStage() *form.Stage
	GetDataEditor() *AsSplit
	GetDiagramEditor() *AsSplitArea
	FillUpFormFromGongstruct(instance any, formName string)
	AddNotification(date time.Time, message string)
	CommitNotificationTable()
	ResetNotifications()
	SetMaxElementsNbPerGongStructNode(nb int)
	GetMaxElementsNbPerGongStructNode() int
	AddCommitNavigationNode(appendChildrenNodeFunc func(GongNodeIF))
	SetCommitMode(bool)
	RefreshNavigationTree()

	GetProbeLoadStageName() string

	UpdateSliceOfPointersCallback(instance any, fieldName string, slicePtr any)
	SetUpdateSliceOfPointersCallback(cb func(instance any, fieldName string, slicePtr any))
}

type ProbeIF = GongProbeIF

type GongNodeIF interface {
}
