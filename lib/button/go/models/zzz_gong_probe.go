// generated code - do not edit
package models

import (
	"time"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	form "github.com/fullstack-lang/gong/lib/table/go/models"
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
}
