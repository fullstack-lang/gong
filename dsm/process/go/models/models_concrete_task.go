package models

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"
)

const (
	ForkJoinBarThickness  = 8.0
	ForkJoinDefaultLength = 120.0
	ForkJoinMinLength     = 24.0
	DecisionDiamondSize   = 50.0
)

type TaskShape struct {
	Name string

	Task *Task

	IsExpanded bool

	// Used when Task.IsForkJoin is true
	IsVertical bool

	// Used when Task.IsForkJoin or Task.IsDecisionNode is true
	IsNameDisplayed bool

	RectShape
}

func (s *TaskShape) FillUpExtraShapeAppearanceForm(formStage *form.Stage, formGroup *form.FormGroup) {
	if s.Task == nil || (!s.Task.IsForkJoin && !s.Task.IsDecisionNode) {
		return
	}
	if s.Task.IsForkJoin {
		gongprobe.BasicFieldtoForm("IsVertical", s.IsVertical, formStage, formGroup, false, false, 0, false, 0, false)
	}
	gongprobe.BasicFieldtoForm("IsNameDisplayed", s.IsNameDisplayed, formStage, formGroup, false, false, 0, false, 0, false)
}

func (s *TaskShape) SaveExtraShapeAppearanceForm(formDiv *form.FormDiv) {
	if s.Task == nil || (!s.Task.IsForkJoin && !s.Task.IsDecisionNode) {
		return
	}
	switch formDiv.Name {
	case "IsVertical":
		if !s.Task.IsForkJoin {
			return
		}
		prevVertical := s.IsVertical
		gongprobe.FormDivBasicFieldToField(&s.IsVertical, formDiv)
		if prevVertical != s.IsVertical {
			centerX := s.X + s.Width/2.0
			centerY := s.Y + s.Height/2.0
			s.Width, s.Height = s.Height, s.Width
			s.X = centerX - s.Width/2.0
			s.Y = centerY - s.Height/2.0
		}
	case "IsNameDisplayed":
		gongprobe.FormDivBasicFieldToField(&s.IsNameDisplayed, formDiv)
	}
}

func (s *TaskShape) GetAbstractElement() AbstractType {
	if s.Task == nil {
		return nil // otherwise returns s.Task
	}
	return s.Task
}

func (s *TaskShape) SetAbstractElement(abstractElement AbstractType) {
	s.Task = abstractElement.(*Task)
}

var _ ConcreteType = (*TaskShape)(nil)
