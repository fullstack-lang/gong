package models

// the states of the engine drivers
type EngineDriverState int

const (
	COMMIT_AGENT_STATES EngineDriverState = iota
	CHECKOUT_AGENT_STATES
	FIRE_ONE_EVENT
	SLEEP_100_MS
	RESET_SIMULATION
	UNKOWN
)
