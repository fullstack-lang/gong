package gantt2svg

import (
	gonggantt_models "github.com/fullstack-lang/gong/lib/gantt/go/models"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// type of the singloton for interception of gantt commit in order to generate
// the svg
type GanttSVGMapper struct {
	mapBar_Rect     map[*gonggantt_models.Bar]*gongsvg_models.Rect
	mapRect4Bar_Bar map[*gongsvg_models.Rect]*gonggantt_models.Bar

	ganttToRender *gonggantt_models.Gantt

	// to overwrite the gantt definition file
	GanttOuputFile string
}
