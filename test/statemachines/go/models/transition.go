package models

import (
	"time"

	buttons "github.com/fullstack-lang/gong/lib/button/go/models"
)

// Transition decribes authorized between states
type Transition struct {
	//gong:width 600
	Name string

	//gong:width 600
	Start *State

	//gong:width 600
	End *State

	// For Role Base Access Control
	RolesWithPermissions []*Role

	GeneratedMessages []*MessageType

	// Diagram dans lequel la transition est présente
	Diagrams []*Diagram
}

func (transition *Transition) OnAfterUpdate(stage *Stage, stagedTransition, frontTransition *Transition) {

	transition.performTransition(stage)
}

// Generic button creation function
func NewButtonTranstion(
	target buttons.Target,
	name string,
	icon string,
	label string,
	transition *Transition,
	stage *Stage,
) *buttons.Button {
	button := new(buttons.Button).Stage(target.GetButtonsStage())
	button.Name = name
	button.Icon = icon
	button.Label = label

	proxy := NewButtonTransitionProxy(
		button,
		target,
		transition,
		stage,
	)

	button.Proxy = proxy

	return button
}

// NewButtonProxy creates a new proxy for a button
func NewButtonTransitionProxy(
	button *buttons.Button,
	target buttons.Target,
	transition *Transition,
	stage *Stage,
) *ButtonTransitionProxy {
	proxy := new(ButtonTransitionProxy)
	proxy.Button = button
	proxy.Target = target
	proxy.transition = transition
	proxy.stage = stage

	return proxy
}

// ButtonProxy is a generic proxy for both int and float64
type ButtonTransitionProxy struct {
	Button     *buttons.Button
	Target     buttons.Target
	transition *Transition
	stage      *Stage
}

// Updated handles updating values when the button changes
func (proxy *ButtonTransitionProxy) Updated() {

	proxy.transition.performTransition(proxy.stage)
}

func (transition *Transition) performTransition(stage *Stage) {

	objectSet := *GetGongstructInstancesSet[Object](stage)
	for object := range objectSet {
		if object.IsSelected {

			object.State = transition.End

			var messageTypeGenere *MessageType
			for _, messageTypeGenere_ := range transition.GeneratedMessages {
				messageTypeGenere = messageTypeGenere_
			}

			if messageTypeGenere != nil {
				message := new(Message).Stage(stage)

				message.Name = time.Now().Format("15:04:05") + " " + transition.Name + " -> " + messageTypeGenere.Name
				message.MessageType = messageTypeGenere
				message.OriginTransition = transition
				object.Messages = append(object.Messages, message)

				// set new message to selected
				for message_ := range *GetGongstructInstancesSet[Message](stage) {
					message_.IsSelected = false
				}
				message.IsSelected = true
			}

			object.Proxy.stager.stage.Commit()
		}
	}
}
