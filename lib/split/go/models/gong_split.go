// Generated code - do not edit
package models

import (
	"time"
)

// for the split package, we define a ProbeIF interface
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
}

type GongNodeIF interface {
}
