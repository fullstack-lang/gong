// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type ActorState_WOP struct {
	// insertion point

	Name string

	Description string

	IsWithProbaility bool

	Probability ProbabilityEnum

	ComputedPrefix string

	IsExpanded bool
}

func (from *ActorState) CopyBasicFields(to *ActorState) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.IsWithProbaility = from.IsWithProbaility
	to.Probability = from.Probability
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type ActorStateShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ActorStateShape) CopyBasicFields(to *ActorStateShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type ActorStateTransition_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *ActorStateTransition) CopyBasicFields(to *ActorStateTransition) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type ActorStateTransitionShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ActorStateTransitionShape) CopyBasicFields(to *ActorStateTransitionShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Analysis_WOP struct {
	// insertion point

	Name string

	Description string

	IsScenariosNodeExpanded bool

	IsGroupUseNodeExpanded bool

	IsGeoObjectUseNodeExpanded bool

	IsMapUseNodeExpanded bool

	ComputedPrefix string

	IsExpanded bool
}

func (from *Analysis) CopyBasicFields(to *Analysis) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.IsScenariosNodeExpanded = from.IsScenariosNodeExpanded
	to.IsGroupUseNodeExpanded = from.IsGroupUseNodeExpanded
	to.IsGeoObjectUseNodeExpanded = from.IsGeoObjectUseNodeExpanded
	to.IsMapUseNodeExpanded = from.IsMapUseNodeExpanded
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type ControlPointShape_WOP struct {
	// insertion point

	Name string

	X_Relative float64

	Y_Relative float64

	IsStartShapeTheClosestShape bool
}

func (from *ControlPointShape) CopyBasicFields(to *ControlPointShape) {
	// insertion point
	to.Name = from.Name
	to.X_Relative = from.X_Relative
	to.Y_Relative = from.Y_Relative
	to.IsStartShapeTheClosestShape = from.IsStartShapeTheClosestShape
}

type Diagram_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	IsChecked bool

	IsShowPrefix bool

	Description string

	IsEvolutionDirectionsNodeExpanded bool

	IsActorStatesNodeExpanded bool

	IsParametersNodeExpanded bool

	IsParametersAggregatesNodeExpanded bool

	IsActorStateTransitionsNodeExpanded bool

	AxisOrign_X float64

	AxisOrign_Y float64

	VerticalAxis_Top_Y float64

	VerticalAxis_Bottom_Y float64

	VerticalAxis_StrokeWidth float64

	HorizontalAxis_Right_X float64

	Start time.Time

	End time.Time

	NumberOfYearsBetweenTicks int
}

func (from *Diagram) CopyBasicFields(to *Diagram) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsChecked = from.IsChecked
	to.IsShowPrefix = from.IsShowPrefix
	to.Description = from.Description
	to.IsEvolutionDirectionsNodeExpanded = from.IsEvolutionDirectionsNodeExpanded
	to.IsActorStatesNodeExpanded = from.IsActorStatesNodeExpanded
	to.IsParametersNodeExpanded = from.IsParametersNodeExpanded
	to.IsParametersAggregatesNodeExpanded = from.IsParametersAggregatesNodeExpanded
	to.IsActorStateTransitionsNodeExpanded = from.IsActorStateTransitionsNodeExpanded
	to.AxisOrign_X = from.AxisOrign_X
	to.AxisOrign_Y = from.AxisOrign_Y
	to.VerticalAxis_Top_Y = from.VerticalAxis_Top_Y
	to.VerticalAxis_Bottom_Y = from.VerticalAxis_Bottom_Y
	to.VerticalAxis_StrokeWidth = from.VerticalAxis_StrokeWidth
	to.HorizontalAxis_Right_X = from.HorizontalAxis_Right_X
	to.Start = from.Start
	to.End = from.End
	to.NumberOfYearsBetweenTicks = from.NumberOfYearsBetweenTicks
}

type Document_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Document) CopyBasicFields(to *Document) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type DocumentUse_WOP struct {
	// insertion point

	Name string
}

func (from *DocumentUse) CopyBasicFields(to *DocumentUse) {
	// insertion point
	to.Name = from.Name
}

type EvolutionDirection_WOP struct {
	// insertion point

	Name string

	Description string

	ComputedPrefix string

	IsExpanded bool
}

func (from *EvolutionDirection) CopyBasicFields(to *EvolutionDirection) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type EvolutionDirectionShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *EvolutionDirectionShape) CopyBasicFields(to *EvolutionDirectionShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Foo_WOP struct {
	// insertion point

	Name string
}

func (from *Foo) CopyBasicFields(to *Foo) {
	// insertion point
	to.Name = from.Name
}

type GeoObject_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *GeoObject) CopyBasicFields(to *GeoObject) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type GeoObjectUse_WOP struct {
	// insertion point

	Name string
}

func (from *GeoObjectUse) CopyBasicFields(to *GeoObjectUse) {
	// insertion point
	to.Name = from.Name
}

type Group_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Group) CopyBasicFields(to *Group) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type GroupUse_WOP struct {
	// insertion point

	Name string
}

func (from *GroupUse) CopyBasicFields(to *GroupUse) {
	// insertion point
	to.Name = from.Name
}

type Library_WOP struct {
	// insertion point

	Name string

	Description string

	ComputedPrefix string

	IsExpanded bool

	IsRootLibrary bool

	IsAnalysesNodeExpanded bool

	IsSubLibrariesNodeExpanded bool

	NbPixPerCharacter float64

	LogoSVGFile string

	IsExpandedTmp bool
}

func (from *Library) CopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsRootLibrary = from.IsRootLibrary
	to.IsAnalysesNodeExpanded = from.IsAnalysesNodeExpanded
	to.IsSubLibrariesNodeExpanded = from.IsSubLibrariesNodeExpanded
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
	to.IsExpandedTmp = from.IsExpandedTmp
}

type MapObject_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *MapObject) CopyBasicFields(to *MapObject) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type MapObjectUse_WOP struct {
	// insertion point

	Name string
}

func (from *MapObjectUse) CopyBasicFields(to *MapObjectUse) {
	// insertion point
	to.Name = from.Name
}

type Parameter_WOP struct {
	// insertion point

	Name string

	Description string

	IsResponse bool

	Start time.Time

	End time.Time

	Force float64

	Tag string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Parameter) CopyBasicFields(to *Parameter) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.IsResponse = from.IsResponse
	to.Start = from.Start
	to.End = from.End
	to.Force = from.Force
	to.Tag = from.Tag
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type ParameterCategory_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *ParameterCategory) CopyBasicFields(to *ParameterCategory) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type ParameterCategoryUse_WOP struct {
	// insertion point

	Name string
}

func (from *ParameterCategoryUse) CopyBasicFields(to *ParameterCategoryUse) {
	// insertion point
	to.Name = from.Name
}

type ParameterShape_WOP struct {
	// insertion point

	Name string

	Direction DirectionType

	ShapeIsComputedFromModel bool

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ParameterShape) CopyBasicFields(to *ParameterShape) {
	// insertion point
	to.Name = from.Name
	to.Direction = from.Direction
	to.ShapeIsComputedFromModel = from.ShapeIsComputedFromModel
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type ParametersAggregate_WOP struct {
	// insertion point

	Name string

	Tag string

	Description string

	ComputedPrefix string

	IsExpanded bool
}

func (from *ParametersAggregate) CopyBasicFields(to *ParametersAggregate) {
	// insertion point
	to.Name = from.Name
	to.Tag = from.Tag
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type ParametersAggregateShape_WOP struct {
	// insertion point

	Name string

	Direction DirectionType

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ParametersAggregateShape) CopyBasicFields(to *ParametersAggregateShape) {
	// insertion point
	to.Name = from.Name
	to.Direction = from.Direction
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Position_WOP struct {
	// insertion point

	Name string

	Date time.Time

	Ordinate float64

	ComputedPrefix string

	IsExpanded bool
}

func (from *Position) CopyBasicFields(to *Position) {
	// insertion point
	to.Name = from.Name
	to.Date = from.Date
	to.Ordinate = from.Ordinate
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type Repository_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Repository) CopyBasicFields(to *Repository) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type Scenario_WOP struct {
	// insertion point

	Name string

	Description string

	IsDiagramsNodeExpanded bool

	IsActorStatesNodeExpanded bool

	IsActorStateTransitionsNodeExpanded bool

	IsEvolutionDirectionsNodeExpanded bool

	IsParametersNodeExpanded bool

	IsParametersAggretatesNodeExpanded bool

	ComputedPrefix string

	IsExpanded bool
}

func (from *Scenario) CopyBasicFields(to *Scenario) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.IsDiagramsNodeExpanded = from.IsDiagramsNodeExpanded
	to.IsActorStatesNodeExpanded = from.IsActorStatesNodeExpanded
	to.IsActorStateTransitionsNodeExpanded = from.IsActorStateTransitionsNodeExpanded
	to.IsEvolutionDirectionsNodeExpanded = from.IsEvolutionDirectionsNodeExpanded
	to.IsParametersNodeExpanded = from.IsParametersNodeExpanded
	to.IsParametersAggretatesNodeExpanded = from.IsParametersAggretatesNodeExpanded
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type User_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *User) CopyBasicFields(to *User) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type UserUse_WOP struct {
	// insertion point

	Name string
}

func (from *UserUse) CopyBasicFields(to *UserUse) {
	// insertion point
	to.Name = from.Name
}

type Workspace_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *Workspace) CopyBasicFields(to *Workspace) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

// end of insertion point
