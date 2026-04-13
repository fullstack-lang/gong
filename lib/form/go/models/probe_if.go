// Generated code - do not edit
package models

import (
	"time"

	"github.com/fullstack-lang/gong/lib/table/go/models"
)

// for the split package, we define a ProbeIF interface
// otherwise, the split package cannot refer to the main gong package
type ProbeIF interface {
	Refresh()
	GetFormStage() *models.Stage
	FillUpFormFromGongstruct(instance any, formName string)
	AddNotification(date time.Time, message string)
	CommitNotificationTable()
	ResetNotifications()
	AddCommitNavigationNode(appendChildrenNodeFunc func(GongNodeIF))
	RefreshNavigationTree()
}

type GongNodeIF interface {
}
