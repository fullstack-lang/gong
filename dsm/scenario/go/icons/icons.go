package icons

import (
	_ "embed"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

//go:embed direction_evolution.svg
var direction_evolution string
var DirectionEvolutionIcon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "direction_evolution",
	SVG:  direction_evolution,
})

//go:embed diagram_time_evol_FILL0_wght400_GRAD0_opsz24.svg
var diagram_time_evol_FILL0_wght400_GRAD0_opsz24 string
var DiagramTimeEvolutionIcon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "time_evolution_icon",
	SVG:  diagram_time_evol_FILL0_wght400_GRAD0_opsz24,
})

//go:embed scenario.svg
var scenario string
var ScenarioIcon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "scenario_icon",
	SVG:  scenario,
})

//go:embed ActorStateTransitionIcon.svg
var ActorStateTransitionIcon_ string
var ActorStateTransitionIcon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "transition_icon",
	SVG:  ActorStateTransitionIcon_,
})

//go:embed ActorStateIcon.svg
var ActorStateIcon_ string
var Actor_stateIcon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "actor_state_icon",
	SVG:  ActorStateIcon_,
})

//go:embed "parameter push higher.svg"
var parameter_push_higher string
var Parameter_push_hightIcon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "parameter_push_higher_icon",
	SVG:  parameter_push_higher,
})

//go:embed "arrow_up_sided_to_right.svg"
var arrow_up_sided_to_right string
var Parameter_left_Arrow_Up_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "arrow_up_sided_to_right_icon",
	SVG:  arrow_up_sided_to_right,
})

//go:embed "arrow_up_sided_to_left.svg"
var arrow_up_sided_to_left string
var Parameter_right_Arrow_Up_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "arrow_up_sided_to_left_icon",
	SVG:  arrow_up_sided_to_left,
})

//go:embed "arrow_down_sided_to_right.svg"
var arrow_down_sided_to_right string
var Parameter_left_Arrow_Down_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "arrow_down_sided_to_right_icon",
	SVG:  arrow_down_sided_to_right,
})

//go:embed "arrow_down_sided_to_left.svg"
var arrow_down_sided_to_left string
var Parameter_right_Arrow_Down_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "arrow_down_sided_to_left_icon",
	SVG:  arrow_down_sided_to_left,
})

//go:embed "arrow_up_and_arrow_down.svg"
var arrow_up_and_arrow_down string
var Arrow_up_and_arrow_downIcon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "arrow_up_and_arrow_down",
	SVG:  arrow_up_and_arrow_down,
})

//go:embed "flip_direction.svg"
var flip_direction string
var Flip_directionIcon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "flip_direction",
	SVG:  flip_direction,
})

//go:embed "percent_000.svg"
var percent_000 string
var Percent_000_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "percent_000",
	SVG:  percent_000,
})

//go:embed "percent_025.svg"
var percent_025 string
var Percent_025_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "percent_025",
	SVG:  percent_025,
})

//go:embed "percent_050.svg"
var percent_050 string
var Percent_050_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "percent_050",
	SVG:  percent_050,
})

//go:embed "percent_075.svg"
var percent_075 string
var Percent_075_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "percent_075",
	SVG:  percent_075,
})

//go:embed "percent_100.svg"
var percent_100 string
var Percent_100_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "percent_100",
	SVG:  percent_100,
})
