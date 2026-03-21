package models

type ConcreteType interface {
	GongstructIF
	GetAbstractElement() AbstractType
	SetAbstractElement(AbstractType)
}

type AssociationConcreteType interface {
	GongstructIF
	GetAbstractStartElement() AbstractType
	SetAbstractStartElement(AbstractType)
	GetAbstractEndElement() AbstractType
	SetAbstractEndElement(AbstractType)
}

type AssociationConcreteType2[SourceAT AbstractType, TargetAT AbstractType] interface {
	GongstructIF
	GetAbstractStartElement() SourceAT
	SetAbstractStartElement(SourceAT)
	GetAbstractEndElement() TargetAT
	SetAbstractEndElement(TargetAT)
}
