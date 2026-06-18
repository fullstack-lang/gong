package models

type ShapeWithDirection interface {
	GetDirection() DirectionType
	SetDirection(direction DirectionType)
}
