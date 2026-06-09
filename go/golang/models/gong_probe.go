package models

const ModelGongProbeFileTemplate = `// generated code - do not edit
package models

import (
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

type ProbeIF interface {
	Refresh()
	GetFormStage() *form.Stage
	GetDataEditor() *split.AsSplit
	GetDiagramEditor() *split.AsSplitArea
	FillUpFormFromGongstruct(instance any, formName string)
	AddNotification(date time.Time, message string)
	CommitNotificationTable()
	ResetNotifications()
	SetMaxElementsNbPerGongStructNode(nb int)
	GetMaxElementsNbPerGongStructNode() int
	AddCommitNavigationNode(appendChildrenNodeFunc func(GongNodeIF))
	SetCommitMode(bool)
	RefreshNavigationTree() // to be called in delta mode, when the navigation tree shall be refreshed to navigate commits

	// UpdateSliceOfPointersCallback is called after a SliceOfPointers field is updated in the probe
	UpdateSliceOfPointersCallback(instance any, fieldName string, slicePtr any)
	SetUpdateSliceOfPointersCallback(cb func(instance any, fieldName string, slicePtr any))
}

type GongNodeIF interface {
}
`
