package svg2gantt

import (
	"log"

	gonggantt_models "github.com/fullstack-lang/gong/lib/gantt/go/models"
	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type BarImpl struct {
	Bar            *gonggantt_models.Bar
	GanttToRender  *gonggantt_models.Gantt
	GongganttStage *gonggantt_models.StageStruct
}

func (barImpl *BarImpl) RectUpdated(updatedRect *gongsvg_models.Rect) {

	var someethingChanged bool
	newStart := XtoDate(
		barImpl.GanttToRender.XLeftLanes,
		updatedRect.X,
		barImpl.GanttToRender.XRightMargin,
		barImpl.GanttToRender.ComputedStart,
		barImpl.GanttToRender.ComputedEnd)

	newEnd := XtoDate(
		barImpl.GanttToRender.XLeftLanes,
		updatedRect.X+updatedRect.Width,
		barImpl.GanttToRender.XRightMargin,
		barImpl.GanttToRender.ComputedStart,
		barImpl.GanttToRender.ComputedEnd)

	diffStart := barImpl.Bar.Start.Sub(newStart)
	diffEnd := barImpl.Bar.End.Sub(newEnd)

	someethingChanged = diffStart.Abs() > 0 || diffEnd.Abs() > 0

	barImpl.Bar.Start = newStart
	barImpl.Bar.End = newEnd

	if someethingChanged {
		log.Println(diffStart, diffEnd)
		barImpl.Bar.Commit(barImpl.GongganttStage)
		barImpl.GongganttStage.Commit()
	}
}
