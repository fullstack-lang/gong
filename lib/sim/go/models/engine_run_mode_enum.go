package models

// the simulation can run with a relative speed to realtime or full speed
type EngineRunMode int

const (
	RELATIVE_SPEED EngineRunMode = iota
	FULL_SPEED
)
