package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svgTaskInputLink(
	startRect *svg.Rect,
	endRect *svg.Rect,
	InputShape *LinkShape,
	subTask *Task,
	layer *svg.Layer,
	isDashed bool,
) {
	if startRect == nil || endRect == nil {
		return
	}

	link := new(svg.Link)

	link.Name = startRect.Name + " to " + endRect.Name
	link.Stroke = svg.Black.ToString()
	link.StrokeWidth = 1.5
	link.StrokeOpacity = 1

	link.Type = svg.LINK_TYPE_FLOATING_ORTHOGONAL

	link.Start = startRect
	link.StartOrientation = svg.OrientationType(InputShape.StartOrientation)
	link.StartRatio = InputShape.StartRatio

	link.End = endRect
	link.EndOrientation = svg.OrientationType(InputShape.EndOrientation)
	link.EndRatio = InputShape.EndRatio
	link.HasEndArrow = true
	link.EndArrowSize = 10

	link.CornerOffsetRatio = InputShape.CornerOffsetRatio

	// specific
	link.CornerRadius = 5
	if isDashed {
		link.StrokeDashArray = "5 5"
	}

	// For task inputs
	link.StrokeDashArray = "5 5"

	link.Impl = NewTaskInput_LinkShapeProxy(
		InputShape,
		subTask,
		stager,
	)

	layer.Links = append(layer.Links, link)
}

func NewTaskInput_LinkShapeProxy(
	taskShapeInputInterface LinkShapeInterface,
	subTask *Task,
	stager *Stager,
) (proxy *TaskInput_LinkShapeProxy) {
	proxy = new(TaskInput_LinkShapeProxy)
	proxy.linkShapeInterface = taskShapeInputInterface
	proxy.subTask = subTask
	proxy.stager = stager
	return
}

type TaskInput_LinkShapeProxy struct {
	linkShapeInterface LinkShapeInterface
	subTask            *Task
	stager             *Stager
}

// LinkUpdated implements models.LinkImplInterface.
func (p *TaskInput_LinkShapeProxy) LinkUpdated(updatedLink *svg.Link) {
	diff := p.linkShapeInterface.GetStartRatio() != updatedLink.StartRatio ||
		p.linkShapeInterface.GetEndRatio() != updatedLink.EndRatio ||
		p.linkShapeInterface.GetStartOrientation() != OrientationType(updatedLink.StartOrientation) ||
		p.linkShapeInterface.GetEndOrientation() != OrientationType(updatedLink.EndOrientation) ||
		p.linkShapeInterface.GetCornerOffsetRatio() != updatedLink.CornerOffsetRatio

	p.linkShapeInterface.SetStartRatio(updatedLink.StartRatio)
	p.linkShapeInterface.SetEndRatio(updatedLink.EndRatio)
	p.linkShapeInterface.SetCornerOffsetRatio(updatedLink.CornerOffsetRatio)
	p.linkShapeInterface.SetStartOrientation(OrientationType(updatedLink.StartOrientation))
	p.linkShapeInterface.SetEndOrientation(OrientationType(updatedLink.EndOrientation))

	if !diff {
		p.stager.stage.CommitWithSuspendedCallbacks()
		p.stager.probeForm.FillUpFormFromGongstruct(p.subTask, GetPointerToGongstructName[*Task]())
	} else {
		p.stager.stage.Commit()
	}
}
