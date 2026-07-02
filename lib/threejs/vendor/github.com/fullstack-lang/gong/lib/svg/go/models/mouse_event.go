package models

type MouseEventKey string

const (
	MouseEventKeyShift MouseEventKey = "Shift"
	MouseEventKeyAlt   MouseEventKey = "Alt"
	MouseEventKeyMeta  MouseEventKey = "Meta"
)

// MouseEvent is added to Rect and Link to forward when the user interact with
// the object
type MouseEvent struct {
	MouseX, MouseY float64 // coordinates in the component coordinates
	MouseEventKey  MouseEventKey
}
