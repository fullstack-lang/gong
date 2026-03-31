// Generated code - do not edit
package models

import (
	"time"
)

// for the split package, we define a ProbeIF interface (that is never implemented by the split package)
// otherwise, the split package cannot refer to the main gong package
type ProbeIF interface {
	Refresh()
	GetDataEditor() *AsSplit
	GetDiagramEditor() *AsSplitArea
	FillUpFormFromGongstruct(instance any, formName string)
	AddNotification(date time.Time, message string)
	CommitNotificationTable()
	ResetNotifications()
	AddCommitNavigationNode(appendChildrenNodeFunc func(GongNodeIF))
	RefreshNavigationTree()
}

type GongNodeIF interface {
}
