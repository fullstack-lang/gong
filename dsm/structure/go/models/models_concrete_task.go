package models

type TaskShape struct {
	Name string

	Task *Task

	IsExpanded bool

	RectShape
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
