package models

type Message struct {
	Name       string
	IsSelected bool

	MessageType *MessageType

	OriginTransition *Transition
}
