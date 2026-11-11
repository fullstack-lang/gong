package models

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"
)

// A ProtocolShape is the concrete shape for a Protocol carried by a DataFlow.
type ProtocolShape struct {
	Name string

	Protocol *Protocol
	DataFlow *DataFlow

	RectShape
}

func (s *ProtocolShape) GetAbstractElement() AbstractType {
	if s.Protocol == nil {
		return nil
	}
	return s.Protocol
}

func (s *ProtocolShape) SetAbstractElement(abstractElement AbstractType) {
	s.Protocol = abstractElement.(*Protocol)
}

func (s *ProtocolShape) FillUpExtraShapeAppearanceForm(formStage *form.Stage, formGroup *form.FormGroup) {
	gongprobe.BasicFieldtoForm("Width", s.Width, formStage, formGroup, false, false, 0, false, 0, false)
}

func (s *ProtocolShape) SaveExtraShapeAppearanceForm(formDiv *form.FormDiv) {
	switch formDiv.Name {
	case "Width":
		gongprobe.FormDivBasicFieldToField(&s.Width, formDiv)
	}
}

var _ ConcreteType = (*ProtocolShape)(nil)
var _ RectShapeInterface = (*ProtocolShape)(nil)

// A ProtocolLayerShape is the concrete shape for a ProtocolLayer of a DataFlow's Protocol.
type ProtocolLayerShape struct {
	Name string

	ProtocolLayer *ProtocolLayer
	DataFlow      *DataFlow

	RectShape
}

func (s *ProtocolLayerShape) GetAbstractElement() AbstractType {
	if s.ProtocolLayer == nil {
		return nil
	}
	return s.ProtocolLayer
}

func (s *ProtocolLayerShape) SetAbstractElement(abstractElement AbstractType) {
	s.ProtocolLayer = abstractElement.(*ProtocolLayer)
}

func (s *ProtocolLayerShape) FillUpExtraShapeAppearanceForm(formStage *form.Stage, formGroup *form.FormGroup) {
	gongprobe.BasicFieldtoForm("Width", s.Width, formStage, formGroup, false, false, 0, false, 0, false)
}

func (s *ProtocolLayerShape) SaveExtraShapeAppearanceForm(formDiv *form.FormDiv) {
	switch formDiv.Name {
	case "Width":
		gongprobe.FormDivBasicFieldToField(&s.Width, formDiv)
	}
}

var _ ConcreteType = (*ProtocolLayerShape)(nil)
var _ RectShapeInterface = (*ProtocolLayerShape)(nil)
