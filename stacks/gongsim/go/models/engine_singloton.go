package models

// EngineSingloton is the default engine
var EngineSingloton = NewEngineSingloton()

// creates and register The EngineSingloton
func NewEngineSingloton() (engine *Engine) {
	engine = new(Engine)
	engine.Name = "Simulation Engine"
	engine.Stage()
	return
}
