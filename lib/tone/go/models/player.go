package models

import "fmt"

// Player is the singleton object to get the callbacks on play/paused
type Player struct {
	Name string

	Status Status

	// DI contains a dependency injection callback
	OnDI func(*Player) error
}

// InjectDependency is a helper method to set the DI callback
func (player *Player) InjectDependency(callback func(*Player) error) {
	player.OnDI = callback
}

// OnAfterUpdate is called after a Player update
func (player *Player) OnAfterUpdate(stage *Stage, stagedInstance, frontInstance *Player) {
	// Example of invoking the DI callback
	if player.OnDI != nil {
		player.Status = frontInstance.Status
		err := player.OnDI(player)
		if err != nil {
			fmt.Printf("Dependency injection failed: %v\n", err)
		}
	}
}
