package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/cld/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Category1__000000_Intensit_de_la_Concurrence := (&models.Category1{}).Stage(stage)
	__Category1__000001_Niveau_des_prix := (&models.Category1{}).Stage(stage)
	__Category1__000002_Hauteur_des_barri_res_l_entr_e := (&models.Category1{}).Stage(stage)

	__Category1Shape__000000_ := (&models.Category1Shape{}).Stage(stage)
	__Category1Shape__000001_ := (&models.Category1Shape{}).Stage(stage)
	__Category1Shape__000002_ := (&models.Category1Shape{}).Stage(stage)

	__ControlPointShape__000000_Control_Point_Shape_in_Intensit_de_la_Concurrence_to_Niveau_des_prix_0 := (&models.ControlPointShape{}).Stage(stage)
	__ControlPointShape__000001_Control_Point_Shape_in_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence_0 := (&models.ControlPointShape{}).Stage(stage)

	__Desk__000000_Desk := (&models.Desk{}).Stage(stage)

	__Diagram__000000_Default := (&models.Diagram{}).Stage(stage)

	__Influence__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix := (&models.Influence{}).Stage(stage)
	__Influence__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence := (&models.Influence{}).Stage(stage)

	__InfluenceShape__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix := (&models.InfluenceShape{}).Stage(stage)
	__InfluenceShape__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence := (&models.InfluenceShape{}).Stage(stage)

	// Setup of values

	__Category1__000000_Intensit_de_la_Concurrence.Name = `Intensité de la Concurrence`

	__Category1__000001_Niveau_des_prix.Name = `Niveau des prix`

	__Category1__000002_Hauteur_des_barri_res_l_entr_e.Name = `Hauteur des barrières à l'entrée`

	__Category1Shape__000000_.Name = ``
	__Category1Shape__000000_.X = 242.000000
	__Category1Shape__000000_.Y = 245.999985
	__Category1Shape__000000_.Width = 240.000000
	__Category1Shape__000000_.Height = 80.000000

	__Category1Shape__000001_.Name = ``
	__Category1Shape__000001_.X = 118.999938
	__Category1Shape__000001_.Y = 449.000016
	__Category1Shape__000001_.Width = 240.000000
	__Category1Shape__000001_.Height = 80.000000

	__Category1Shape__000002_.Name = ``
	__Category1Shape__000002_.X = 435.000000
	__Category1Shape__000002_.Y = 56.000000
	__Category1Shape__000002_.Width = 240.000000
	__Category1Shape__000002_.Height = 80.000000

	__ControlPointShape__000000_Control_Point_Shape_in_Intensit_de_la_Concurrence_to_Niveau_des_prix_0.Name = `Control Point Shape in Intensité de la Concurrence to Niveau des prix 0`
	__ControlPointShape__000000_Control_Point_Shape_in_Intensit_de_la_Concurrence_to_Niveau_des_prix_0.X_Relative = 0.465104
	__ControlPointShape__000000_Control_Point_Shape_in_Intensit_de_la_Concurrence_to_Niveau_des_prix_0.Y_Relative = -0.774911
	__ControlPointShape__000000_Control_Point_Shape_in_Intensit_de_la_Concurrence_to_Niveau_des_prix_0.IsStartShapeTheClosestShape = false

	__ControlPointShape__000001_Control_Point_Shape_in_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence_0.Name = `Control Point Shape in Hauteur des barrières à l'entrée to Intensité de la Concurrence 0`
	__ControlPointShape__000001_Control_Point_Shape_in_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence_0.X_Relative = 0.273437
	__ControlPointShape__000001_Control_Point_Shape_in_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence_0.Y_Relative = 1.587583
	__ControlPointShape__000001_Control_Point_Shape_in_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence_0.IsStartShapeTheClosestShape = true

	__Desk__000000_Desk.Name = `Desk`

	__Diagram__000000_Default.Name = `Default`
	__Diagram__000000_Default.IsEditable = false
	__Diagram__000000_Default.IsNodeExpanded = true
	__Diagram__000000_Default.IsCategory1NodeExpanded = true
	__Diagram__000000_Default.IsCategory2NodeExpanded = false
	__Diagram__000000_Default.IsCategory3NodeExpanded = false
	__Diagram__000000_Default.IsInfluenceCategoryNodeExpanded = true
	__Diagram__000000_Default.IsCategory1Shown = true
	__Diagram__000000_Default.IsCategory2Shown = false
	__Diagram__000000_Default.IsCategory3Shown = false
	__Diagram__000000_Default.IsInfluenceCategoryShown = true
	__Diagram__000000_Default.XMargin = 0.000000
	__Diagram__000000_Default.YMargin = 0.000000
	__Diagram__000000_Default.Height = 2000.000000
	__Diagram__000000_Default.Width = 2000.000000
	__Diagram__000000_Default.RedColorCode = `salmon `
	__Diagram__000000_Default.BackgroundGreyColorCode = `white`
	__Diagram__000000_Default.GrayColorCode = `gray`
	__Diagram__000000_Default.Category1RectAnchorType = models.RECT_RIGHT
	__Diagram__000000_Default.Category1TextAnchorType = models.TEXT_ANCHOR_END
	__Diagram__000000_Default.Category1FontSize = ``
	__Diagram__000000_Default.Category1FontWeigth = ``
	__Diagram__000000_Default.Category1FontFamily = ``
	__Diagram__000000_Default.Category1LetterSpacing = ``
	__Diagram__000000_Default.Category2TypeFontSize = ``
	__Diagram__000000_Default.Category2TypeFontWeigth = ``
	__Diagram__000000_Default.Category2TypeFontFamily = ``
	__Diagram__000000_Default.Category2TypeLetterSpacing = ``
	__Diagram__000000_Default.Category2StrokeWidth = 0.000000
	__Diagram__000000_Default.Category3FontSize = ``
	__Diagram__000000_Default.Category3FontWeigth = ``
	__Diagram__000000_Default.Category3FontFamily = ``
	__Diagram__000000_Default.Category3LetterSpacing = ``
	__Diagram__000000_Default.InfluenceStrokeWidth = 1.000000
	__Diagram__000000_Default.InfluenceArrowSize = 10.000000
	__Diagram__000000_Default.InfluenceArrowStartOffset = 0.000000
	__Diagram__000000_Default.InfluenceArrowEndOffset = 0.000000
	__Diagram__000000_Default.InfluenceCornerRadius = 19.000000
	__Diagram__000000_Default.InfluenceDashedLinePattern = ``

	__Influence__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix.Name = `Intensité de la Concurrence to Niveau des prix`
	__Influence__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix.IsHypothtical = false
	__Influence__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix.TextAtEndOfArrow = `+`

	__Influence__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence.Name = `Hauteur des barrières à l'entrée to Intensité de la Concurrence`
	__Influence__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence.IsHypothtical = false
	__Influence__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence.TextAtEndOfArrow = ``

	__InfluenceShape__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix.Name = `Intensité de la Concurrence to Niveau des prix`

	__InfluenceShape__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence.Name = `Hauteur des barrières à l'entrée to Intensité de la Concurrence`

	// Setup of pointers
	// setup of Category1 instances pointers
	// setup of Category1Shape instances pointers
	__Category1Shape__000000_.Category1 = __Category1__000000_Intensit_de_la_Concurrence
	__Category1Shape__000001_.Category1 = __Category1__000001_Niveau_des_prix
	__Category1Shape__000002_.Category1 = __Category1__000002_Hauteur_des_barri_res_l_entr_e
	// setup of ControlPointShape instances pointers
	// setup of Desk instances pointers
	__Desk__000000_Desk.SelectedDiagram = __Diagram__000000_Default
	// setup of Diagram instances pointers
	__Diagram__000000_Default.Category1Shapes = append(__Diagram__000000_Default.Category1Shapes, __Category1Shape__000000_)
	__Diagram__000000_Default.Category1Shapes = append(__Diagram__000000_Default.Category1Shapes, __Category1Shape__000001_)
	__Diagram__000000_Default.Category1Shapes = append(__Diagram__000000_Default.Category1Shapes, __Category1Shape__000002_)
	__Diagram__000000_Default.InfluenceShapes = append(__Diagram__000000_Default.InfluenceShapes, __InfluenceShape__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix)
	__Diagram__000000_Default.InfluenceShapes = append(__Diagram__000000_Default.InfluenceShapes, __InfluenceShape__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence)
	// setup of Influence instances pointers
	__Influence__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix.SourceCategory1 = __Category1__000000_Intensit_de_la_Concurrence
	__Influence__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix.TargetCategory1 = __Category1__000001_Niveau_des_prix
	__Influence__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence.SourceCategory1 = __Category1__000002_Hauteur_des_barri_res_l_entr_e
	__Influence__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence.TargetCategory1 = __Category1__000000_Intensit_de_la_Concurrence
	// setup of InfluenceShape instances pointers
	__InfluenceShape__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix.Influence = __Influence__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix
	__InfluenceShape__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix.ControlPointShapes = append(__InfluenceShape__000000_Intensit_de_la_Concurrence_to_Niveau_des_prix.ControlPointShapes, __ControlPointShape__000000_Control_Point_Shape_in_Intensit_de_la_Concurrence_to_Niveau_des_prix_0)
	__InfluenceShape__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence.Influence = __Influence__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence
	__InfluenceShape__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence.ControlPointShapes = append(__InfluenceShape__000001_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence.ControlPointShapes, __ControlPointShape__000001_Control_Point_Shape_in_Hauteur_des_barri_res_l_entr_e_to_Intensit_de_la_Concurrence_0)
}

