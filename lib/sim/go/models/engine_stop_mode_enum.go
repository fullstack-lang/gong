package models

type EngineStopMode int

const (
	TEN_MINUTES EngineStopMode = iota
	STATE_CHANGED
)
