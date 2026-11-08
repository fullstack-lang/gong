package models

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"
)

// A DataShape is the concrete for a Data that is carried by a DataFlow.
type DataShape struct {
	Name string

	Data     *Data
	DataFlow *DataFlow

	RectShape
}

func (s *DataShape) GetAbstractElement() AbstractType {
	if s.Data == nil {
		return nil // otherwise returns s.Data
	}
	return s.Data
}

func (s *DataShape) SetAbstractElement(abstractElement AbstractType) {
	s.Data = abstractElement.(*Data)
}

func (s *DataShape) FillUpExtraShapeAppearanceForm(formStage *form.Stage, formGroup *form.FormGroup) {
	gongprobe.BasicFieldtoForm("Width", s.Width, formStage, formGroup, false, false, 0, false, 0, false)
}

func (s *DataShape) SaveExtraShapeAppearanceForm(formDiv *form.FormDiv) {
	switch formDiv.Name {
	case "Width":
		gongprobe.FormDivBasicFieldToField(&s.Width, formDiv)
	}
}

var _ ConcreteType = (*DataShape)(nil)
var _ RectShapeInterface = (*DataShape)(nil)

