package models

import (
	"log"
	"time"
)

// --- from actor_state_shape.go ---
// ActorStateShape is abstract syntax (this is part of a scenario/diagram)
//
// It might be considered oncrete syntax (here is the way it is represented in the diagram)
// but it is not. The position is part of the definition of the actor state in a scenario
//
//	(not like UML where the position of a shape has no effect on the code)
type ActorStateShape struct {
	Name string

	//gong:type
	ActorState *ActorState

	RectShape
}

func (actorStateShape *ActorStateShape) GetImplName() string {
	return actorStateShape.ActorState.Name
}

func (actorStateShape *ActorStateShape) GetModel() *ActorState {
	return actorStateShape.ActorState
}

// this getter function is used for casting the ActorStateShape into a Gongstruct
func (s *ActorStateShape) GetShape() (r *ActorStateShape) {
	r = s
	return
}

func (s *ActorStateShape) GetAbstractElement() AbstractType {
	if s.ActorState == nil {
		return nil
	}
	return s.ActorState
}

func (s *ActorStateShape) SetAbstractElement(abstractElement AbstractType) {
	s.ActorState = abstractElement.(*ActorState)
}

// --- from actor_state_transition_shape.go ---
type ActorStateTransitionShape struct {
	Name string

	//gong:type
	ActorStateTransition *ActorStateTransition
	Start                *ActorStateShape
	End                  *ActorStateShape

	RectShape

	ControlPointShapes []*ControlPointShape
}

func (s *ActorStateTransitionShape) GetAbstractElement() AbstractType {
	if s.ActorStateTransition == nil {
		return nil
	}
	return s.ActorStateTransition
}

func (s *ActorStateTransitionShape) SetAbstractElement(abstractElement AbstractType) {
	s.ActorStateTransition = abstractElement.(*ActorStateTransition)
}

// --- from control_point_shape.go ---
type ControlPointShape struct {
	Name string

	// Control Points are defined relative to their
	// closest rect between the start rect and the end rect
	X_Relative, Y_Relative      float64
	IsStartShapeTheClosestShape bool
}

// --- from evolution_direction_shape.go ---
type EvolutionDirectionShape struct {
	Name string

	//gong:type
	EvolutionDirection *EvolutionDirection

	RectShape
}

func (evolutionDirectionShape *EvolutionDirectionShape) GetImplName() string {
	return evolutionDirectionShape.EvolutionDirection.Name
}

func (evolutionDirectionShape *EvolutionDirectionShape) GetModel() *EvolutionDirection {
	return evolutionDirectionShape.EvolutionDirection
}

// this getter function is used for casting the GenericRectShape into a Gongstruct
func (s *EvolutionDirectionShape) GetShape() (r *EvolutionDirectionShape) {
	r = s
	return
}

func (s *EvolutionDirectionShape) GetAbstractElement() AbstractType {
	if s.EvolutionDirection == nil {
		return nil
	}
	return s.EvolutionDirection
}

func (s *EvolutionDirectionShape) SetAbstractElement(abstractElement AbstractType) {
	s.EvolutionDirection = abstractElement.(*EvolutionDirection)
}

// --- from parameter_shape.go ---
type ParameterShape struct {
	Name string

	//gong:type
	Parameter *Parameter

	// The shape might have arrowing showing the
	// influence on state actors
	Direction DirectionType

	// ShapeIsComputedFromModel means the shape is dependant on the
	// [models.Parameter.Start] and [models.Parameter.End]
	ShapeIsComputedFromModel bool

	RectShape
}

func (parameterShape *ParameterShape) GetImplName() string {
	return parameterShape.Parameter.Name
}

func (parameterShape *ParameterShape) GetModel() *Parameter {
	return parameterShape.Parameter
}

// this getter function is used for casting the GenericRectShape into a Gongstruct
func (s *ParameterShape) GetShape() (r *ParameterShape) {
	r = s
	return
}

func (s *ParameterShape) GetAbstractElement() AbstractType {
	if s.Parameter == nil {
		return nil
	}
	return s.Parameter
}

func (s *ParameterShape) SetAbstractElement(abstractElement AbstractType) {
	s.Parameter = abstractElement.(*Parameter)
}

func (s *ParameterShape) GetDirection() DirectionType {
	return s.Direction
}

func (s *ParameterShape) SetDirection(direction DirectionType) {
	s.Direction = direction
}

func (parameterShape *ParameterShape) ComputeModelFromShape(stage *Stage) {
	// log.Println("ParamterShape", "ComputeModelFromShape")

	if !parameterShape.ShapeIsComputedFromModel {
		return
	}

	// get diagram
	workspace := GetWorkspace(stage)
	diagram := workspace.SelectedDiagram

	parameterShape.Parameter.Start = XToTime(diagram, parameterShape.X)

	parameterShape.Parameter.End = XToTime(diagram, parameterShape.X+parameterShape.Width)

	// log.Println("yRelative:", yRelative)
	// log.Println("relativeDurationY:", relativeDurationY)
	// log.Println("relativeDurationY:")
	// printDuration(relativeDurationY)
	// log.Println("parameterShape.Parameter.End", parameterShape.Parameter.End)

	diagramHeigth := diagram.VerticalAxis_Bottom_Y - diagram.VerticalAxis_Top_Y

	parameterShape.Parameter.Force = parameterShape.Height / diagramHeigth

	switch parameterShape.Direction {
	case DIRECTION_DOWN:
		parameterShape.Parameter.Force = -parameterShape.Parameter.Force
	case "":
		parameterShape.Parameter.Force = 0
	}

	// log.Println("ParamterShape", "ComputeModelFromShape", s.Parameter.Start)
}

func (parameterShape *ParameterShape) ComputeShapeFromModel(stage *Stage) {
	// log.Println("ParamterShape", "ComputeModelFromShape")

	if !parameterShape.ShapeIsComputedFromModel {
		return
	}

	// get diagram
	workspace := GetWorkspace(stage)
	diagram := workspace.SelectedDiagram

	parameter := parameterShape.Parameter
	if parameter.Start == parameter.End {
		return
	}

	parameterShape.X = TimeToX(diagram, parameter.Start)
	parameterShape.Width = TimeToX(diagram, parameter.End) - parameterShape.X

	// log.Println("ParamterShape", "ComputeShapeFromModel", parameterShape.X, parameterShape.Width)
}

func XToTime(diagram *Diagram, x float64) (date time.Time) {

	// compute relative x to diagram
	diagramWidth := diagram.HorizontalAxis_Right_X - diagram.AxisOrign_X
	diagramDuration := diagram.End.Sub(diagram.Start)

	// log.Println("diagramWidth:", diagramWidth)
	// log.Println("diagramHeigth:", diagramHeigth)
	// log.Println("diagramDuration:", diagramDuration)
	// log.Println("diagram start", diagram.Start)
	// log.Println("diagram end", diagram.End)

	// printDuration(diagramDuration)

	xRelative := (x - diagram.AxisOrign_X) / diagramWidth
	relativeDurationX := time.Duration(float64(diagramDuration) * xRelative)
	date = diagram.Start.Add(relativeDurationX)

	return
}

func TimeToX(diagram *Diagram, date time.Time) (x float64) {

	diagramDuration := diagram.End.Sub(diagram.Start)
	diagramWidth := diagram.HorizontalAxis_Right_X - diagram.AxisOrign_X

	ratioOnDiagram := float64(date.Sub(diagram.Start)) / float64(diagramDuration)

	x = diagram.AxisOrign_X + ratioOnDiagram*diagramWidth

	return
}

// OrdinateToY translates a -1 ; 1 to a Y coordinate.
// If ordinate is -1, y is diagram.VerticalAxis_Bottom_Y
// If ordinate is 1, y is diagram.VerticalAxis_Top_Y
func OrdinateToY(diagram *Diagram, ordinate float64) (y float64) {

	diagramHeight := diagram.VerticalAxis_Bottom_Y - diagram.VerticalAxis_Top_Y

	// Map the ordinate (-1 to 1) to the Y coordinate
	// If ordinate is -1, return Bottom_Y
	// If ordinate is 1, return Top_Y
	// For intermediate values, interpolate linearly
	y = diagram.VerticalAxis_Bottom_Y - ((ordinate+1)/2)*diagramHeight

	return
}

func printDuration(d time.Duration) {
	years := int(d.Truncate(time.Hour*24*365).Hours() / (24 * 365))
	months := int(d.Truncate(time.Hour*24*30).Hours()/(24*30)) % 12
	days := int(d.Truncate(time.Hour*24).Hours()/24) % 30
	hours := int(d.Truncate(time.Hour).Hours()) % 24
	minutes := int(d.Truncate(time.Minute).Minutes()) % 60

	if years > 0 {
		log.Printf("%d year%s ", years, pluralize(years))
	}
	if months > 0 {
		log.Printf("%d month%s ", months, pluralize(months))
	}
	if days > 0 {
		log.Printf("%d day%s ", days, pluralize(days))
	}
	if hours > 0 {
		log.Printf("%d hour%s ", hours, pluralize(hours))
	}
	if minutes > 0 {
		log.Printf("%d minute%s", minutes, pluralize(minutes))
	}
	log.Println()
}

func pluralize(count int) string {
	if count != 1 {
		return "s"
	}
	return ""
}

// --- from parameters_aggregate_shape.go ---
type ParametersAggregateShape struct {
	Name string

	//gong:type
	ScenarioParameter *ParametersAggregate

	// The shape might have arrowing showing the
	// influence on state actors
	Direction DirectionType

	RectShape
}

func (scenarioParameterShape *ParametersAggregateShape) GetImplName() string {
	return scenarioParameterShape.ScenarioParameter.Name
}

func (scenarioParameterShape *ParametersAggregateShape) GetModel() *ParametersAggregate {
	return scenarioParameterShape.ScenarioParameter
}

// this getter function is used for casting the GenericRectShape into a Gongstruct
func (s *ParametersAggregateShape) GetShape() (r *ParametersAggregateShape) {
	r = s
	return
}

func (s *ParametersAggregateShape) GetAbstractElement() AbstractType {
	if s.ScenarioParameter == nil {
		return nil
	}
	return s.ScenarioParameter
}

func (s *ParametersAggregateShape) SetAbstractElement(abstractElement AbstractType) {
	s.ScenarioParameter = abstractElement.(*ParametersAggregate)
}

func (s *ParametersAggregateShape) GetDirection() DirectionType {
	return s.Direction
}

func (s *ParametersAggregateShape) SetDirection(direction DirectionType) {
	s.Direction = direction
}

// --- from diagram.go ---
type Diagram struct {
	Name string

	// mandatory for --dsm
	AbstractTypeFields

	// if checked, will be displayed
	IsChecked bool
	// end of mandatory for --dsm

	IsShowPrefix       bool
	DiagramListElement AbstractType

	//gong:text, magic code to have the form editor have a text area instead of an input
	Description string

	EvolutionDirectionShapes               []*EvolutionDirectionShape
	EvolutionDirectionsWhoseNodeIsExpanded []*EvolutionDirection
	IsEvolutionDirectionsNodeExpanded      bool

	ActorStateShapes               []*ActorStateShape
	ActorStatesWhoseNodeIsExpanded []*ActorState
	IsActorStatesNodeExpanded      bool

	ParameterShapes               []*ParameterShape
	ParametersWhoseNodeIsExpanded []*Parameter
	IsParametersNodeExpanded      bool

	ScenarioParameterShapes                 []*ParametersAggregateShape
	ParametersAggregatesWhoseNodeIsExpanded []*ParametersAggregate
	IsParametersAggregatesNodeExpanded      bool

	ActorStateTransitionShapes               []*ActorStateTransitionShape
	ActorStateTransitionsWhoseNodeIsExpanded []*ActorStateTransition
	IsActorStateTransitionsNodeExpanded      bool

	// origin of axis
	AxisOrign_X float64 // left of the arrow
	AxisOrign_Y float64

	// position of the vertical axis
	// in SVG, vertical coordinate go from 0 (top of the screen) to +infinite (bottom of the screen)
	VerticalAxis_Top_Y       float64
	VerticalAxis_Bottom_Y    float64
	VerticalAxis_StrokeWidth float64

	// HoryzontalAxis_Right_X is the X coordinate of the arrow at the end of the horizontal axis
	HorizontalAxis_Right_X float64

	// abscisse dates
	Start time.Time
	End   time.Time

	// NumberOfYearsBetweenTicks is when too many years would clutter the display
	// if 0 years is equivalent to 1
	NumberOfYearsBetweenTicks int

	// IsInDrawMode indicates the the drawing can be edited (in development mode)
	// or not (in production mode)
	// IsInDrawMode is a transciant value. It can only be persisted as a false value
	// therefore, it is not persisted at all
	// gong:ignore
	IsInDrawMode bool
	LibraryAbstractFields
}

func (d *Diagram) GetYearsWithJanuaryFirstBetween() (years []time.Time) {
	for year := d.Start.Year(); year <= d.End.Year(); year++ {
		janFirst := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		if janFirst.After(d.Start) && janFirst.Before(d.End) {
			years = append(years, janFirst)
		}
	}
	return
}

// mandatory for --dsm
func (d *Diagram) GetIsChecked() bool {
	return d.IsChecked
}

func (d *Diagram) SetIsChecked(v bool) {
	d.IsChecked = v
}

func (d *Diagram) GetIsShowPrefix() bool {
	return d.IsShowPrefix
}

func (d *Diagram) SetIsShowPrefix(v bool) {
	d.IsShowPrefix = v
}

func (d *Diagram) GetDefaultBoxWidth() float64 {
	return 200.0 // Default width
}

func (d *Diagram) GetDefaultBoxHeigth() float64 {
	return 50.0 // Default height
}

func (d *Diagram) IsEditable() bool {
	return true
}

func (d *Diagram) GetDiagramListElement() AbstractType {
	return d.DiagramListElement
}

func (d *Diagram) SetDiagramListElement(abstractElement AbstractType) {
	d.DiagramListElement = abstractElement
}

// --- from generic_rect_shape.go ---
// GenericRectShape
type GenericRectShape[ShapeType, ModelType Gongstruct] interface {
	*EvolutionDirectionShape | *ActorStateShape | *ParameterShape | *ParametersAggregateShape

	// https://github.com/golang/go/issues/63940
	// waiting to avoid the below code

	GetImplName() string
	GetModel() *ModelType
	GetShape() *ShapeType

	GetName() string
	GetX() float64
	GetY() float64
	GetWidth() float64
	GetHeight() float64
	GetFillColor() string
	GetFillOpacity() float64
	GetStroke() string
	GetStrokeWidth() float64
	GetRX() float64

	SetName(name string)
	SetX(x float64)
	SetY(y float64)
	SetWidth(width float64)
	SetHeight(Height float64)
	SetFillColor(fillColor string)
	SetFillOpacity(fillOpacity float64)
	SetStroke(stroke string)
	SetStrokeWidth(strokeWidth float64)
	SetRX(rx float64)

	// ComputeModelFromShape allows, for some shape
	// to compute the model fields from the shape
	ComputeModelFromShape(stage *Stage)
	ComputeShapeFromModel(stage *Stage)
}
