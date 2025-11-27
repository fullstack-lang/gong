package models

type Layout struct {
	Name string

	Groups       []*Group
	GroupToogles []*GroupToogle
}

type Group struct {
	Name       string
	Percentage float64
	Buttons    []*Button
	NbColumns  int
}

type Button struct {
	Name string

	Label string

	Icon string

	IsDisabled bool

	Color MatButtonPaletteType

	MatButtonType MatButtonType

	MatButtonAppearance MatButtonAppearance

	Proxy ButtonProxyInterface
}

type GroupToogle struct {
	Name          string
	Percentage    float64
	ButtonToggles []*ButtonToggle

	IsSingleSelector bool
}

type ButtonToggle struct {
	Name string

	Label string

	Icon string

	IsDisabled bool

	Proxy ButtonProxyInterface
}

/*
text	Default appearance. Text buttons are used for the lowest priority actions, especially when presenting multiple options.
filled	High-emphasis buttons used for final or unblocking actions in a flow, such as saving or confirming.
tonal	Medium-emphasis buttons often used for final or unblocking actions in a flow, but with less visual emphasis than a filled button.
outlined	Medium-emphasis buttons often used for actions that need attention but aren't the primary action.
elevated	Medium-emphasis buttons often used when a button requires visual separation from a patterned background.
*/
type MatButtonAppearance string

const (
	MatButtonAppearanceText     MatButtonAppearance = "text"
	MatButtonAppearanceFilled   MatButtonAppearance = "filled"
	MatButtonAppearanceElevated MatButtonAppearance = "elevated"
	MatButtonAppearanceOutlined MatButtonAppearance = "outlined"
	MatButtonAppearanceTonal    MatButtonAppearance = "tonal"
)

/*
matButton	Rectangular button that can contain text and icons
matIconButton	Smaller, circular button, meant to contain an icon and no text
matFab	Rectangular button w/ elevation and rounded corners, meant to contain an icon. Can be extended to a rectangle to also fit a label
matMiniFab	Smaller variant of matFab
*/
type MatButtonType string

const (
	MatButtonTypeBasic       MatButtonType = "matButton"
	MatButtonTypeIcon        MatButtonType = "matIconButton"
	MatButtonTypeFab         MatButtonType = "matFab"
	MatButtonTypeMiniFab     MatButtonType = "matMiniFab"
	MatButtonTypeExtendedFab MatButtonType = "matFab extended"
)

type MatButtonPaletteType string

const (
	MatButtonPaletteTypePrimary MatButtonPaletteType = "primary"
	MatButtonPaletteTypeWarn    MatButtonPaletteType = "warn"
	MatButtonPaletteTypeAccent  MatButtonPaletteType = "accent"
)

type ButtonProxyInterface interface {
	Updated()
}

func (button *Button) OnAfterUpdate(
	stage *Stage,
	stageButton, frontButton *Button) {

	if button.Proxy != nil {
		button.Proxy.Updated()
	}
}

func (buttonToggle *ButtonToggle) OnAfterUpdate(
	stage *Stage,
	stageButton, frontButton *Button) {

	if buttonToggle.Proxy != nil {
		buttonToggle.Proxy.Updated()
	}
}
