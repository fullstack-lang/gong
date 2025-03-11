package gantt2svg

import (
	"fmt"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// appends 5 lines to the SVG
func generate_arrow(
	gongsvgStage *gongsvg_models.StageStruct,
	svg *gongsvg_models.Layer, X1, X2, Y1, Y2, horyzontalLenght,
	arrowTipLength float64,
	optionnalColor,
	optionnalStroke,
	name string) {

	polyline := new(gongsvg_models.Polyline).Stage(gongsvgStage)
	polyline.Name = name
	svg.Polylines = append(svg.Polylines, polyline)

	polyline.Points = fmt.Sprintf(
		"%d,%d    %d,%d    %d,%d    %d,%d    %d,%d    %d,%d    %d,%d    %d,%d    %d,%d",

		// point 1
		int64(X1), int64(Y1),

		// point 2
		int64(X1+horyzontalLenght), int64(Y1),

		// point 3
		int64(X1+horyzontalLenght), int64((Y1+Y2)/2),

		// point 4
		int64(X2-horyzontalLenght), int64((Y1+Y2)/2),

		// point 5
		int64(X2-horyzontalLenght), int64(Y2),

		// point 6
		int64(X2), int64(Y2),

		// point 7
		int64(X2-arrowTipLength), int64(Y2+arrowTipLength),

		// point 8
		int64(X2), int64(Y2),

		// point 9
		int64(X2-arrowTipLength), int64(Y2-arrowTipLength),
	)
	polyline.Stroke = "blue"
	if optionnalStroke != "" {
		polyline.Stroke = optionnalStroke
	}

	polyline.StrokeWidth = 0.5
	polyline.StrokeDashArray = ""
}
