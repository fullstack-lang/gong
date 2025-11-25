package models

import (
	"time"
)

type Object struct {
	Name       string
	State      *State
	IsSelected bool
	Rank       int

	DOF time.Time

	Proxy *Object_Tree_Proxy

	Messages []*Message
}
