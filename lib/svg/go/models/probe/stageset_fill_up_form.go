// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
	case *models.Animate:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AttributeName", inst.AttributeName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Values", inst.Values, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("From", inst.From, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("To", inst.To, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dur", inst.Dur, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RepeatCount", inst.RepeatCount, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Circles {
				for _, target := range src.Animations {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Circle", "Animations", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Ellipses {
				for _, target := range src.Animates {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Ellipse", "Animates", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Lines {
				for _, target := range src.Animates {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Line", "Animates", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.LinkAnchoredTexts {
				for _, target := range src.Animates {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.LinkAnchoredText", "Animates", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Paths {
				for _, target := range src.Animates {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Path", "Animates", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Polygones {
				for _, target := range src.Animates {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Polygone", "Animates", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Polylines {
				for _, target := range src.Animates {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Polyline", "Animates", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rects {
				for _, target := range src.Animations {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rect", "Animations", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.RectAnchoredTexts {
				for _, target := range src.Animates {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RectAnchoredText", "Animates", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Texts {
				for _, target := range src.Animates {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Text", "Animates", refNames, formGroup, probe.formStage)
		}
	case *models.Circle:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CX", inst.CX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CY", inst.CY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Radius", inst.Radius, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			// Slice of pointers: Animations
			div := (&form.FormDiv{Name: "Animations"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Animations {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Animations",
				Label: "Animations",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layers {
				for _, target := range src.Circles {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layer", "Circles", refNames, formGroup, probe.formStage)
		}
	case *models.Condition:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rects {
				for _, target := range src.HoveringTrigger {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rect", "HoveringTrigger", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rects {
				for _, target := range src.DisplayConditions {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rect", "DisplayConditions", refNames, formGroup, probe.formStage)
		}
	case *models.ControlPoint:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X_Relative", inst.X_Relative, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y_Relative", inst.Y_Relative, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ClosestRect", inst.ClosestRect, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Links {
				for _, target := range src.ControlPoints {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Link", "ControlPoints", refNames, formGroup, probe.formStage)
		}
	case *models.Ellipse:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CX", inst.CX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CY", inst.CY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RX", inst.RX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RY", inst.RY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			// Slice of pointers: Animates
			div := (&form.FormDiv{Name: "Animates"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Animates {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Animates",
				Label: "Animates",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layers {
				for _, target := range src.Ellipses {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layer", "Ellipses", refNames, formGroup, probe.formStage)
		}
	case *models.FileToDownload:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Base64EncodedContent", inst.Base64EncodedContent, probe.formStage, formGroup)
	case *models.Layer:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Rects
			div := (&form.FormDiv{Name: "Rects"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Rects {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Rects",
				Label: "Rects",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Texts
			div := (&form.FormDiv{Name: "Texts"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Texts {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Texts",
				Label: "Texts",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Circles
			div := (&form.FormDiv{Name: "Circles"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Circles {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Circles",
				Label: "Circles",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Lines
			div := (&form.FormDiv{Name: "Lines"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Lines {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Lines",
				Label: "Lines",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Ellipses
			div := (&form.FormDiv{Name: "Ellipses"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Ellipses {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Ellipses",
				Label: "Ellipses",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Polylines
			div := (&form.FormDiv{Name: "Polylines"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Polylines {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Polylines",
				Label: "Polylines",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Polygones
			div := (&form.FormDiv{Name: "Polygones"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Polygones {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Polygones",
				Label: "Polygones",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Paths
			div := (&form.FormDiv{Name: "Paths"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Paths {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Paths",
				Label: "Paths",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Links
			div := (&form.FormDiv{Name: "Links"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Links {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Links",
				Label: "Links",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RectLinkLinks
			div := (&form.FormDiv{Name: "RectLinkLinks"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RectLinkLinks {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RectLinkLinks",
				Label: "RectLinkLinks",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.SVGs {
				for _, target := range src.Layers {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SVG", "Layers", refNames, formGroup, probe.formStage)
		}
	case *models.Line:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X1", inst.X1, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y1", inst.Y1, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X2", inst.X2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y2", inst.Y2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			// Slice of pointers: Animates
			div := (&form.FormDiv{Name: "Animates"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Animates {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Animates",
				Label: "Animates",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("MouseClickX", inst.MouseClickX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MouseClickY", inst.MouseClickY, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layers {
				for _, target := range src.Lines {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layer", "Lines", refNames, formGroup, probe.formStage)
		}
	case *models.Link:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("IsBezierCurve", inst.IsBezierCurve, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Start", inst.Start, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), probe.formStage)
		StageSetEnumStringFieldToForm("StartAnchorType", inst.StartAnchorType, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("End", inst.End, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), probe.formStage)
		StageSetEnumStringFieldToForm("EndAnchorType", inst.EndAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CornerRadius", inst.CornerRadius, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasEndArrow", inst.HasEndArrow, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndArrowSize", inst.EndArrowSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndArrowOffset", inst.EndArrowOffset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasStartArrow", inst.HasStartArrow, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StartArrowSize", inst.StartArrowSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StartArrowOffset", inst.StartArrowOffset, probe.formStage, formGroup)

		{
			// Slice of pointers: TextAtArrowStart
			div := (&form.FormDiv{Name: "TextAtArrowStart"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TextAtArrowStart {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TextAtArrowStart",
				Label: "TextAtArrowStart",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TextAtArrowEnd
			div := (&form.FormDiv{Name: "TextAtArrowEnd"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TextAtArrowEnd {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TextAtArrowEnd",
				Label: "TextAtArrowEnd",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TextAtCorner
			div := (&form.FormDiv{Name: "TextAtCorner"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TextAtCorner {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TextAtCorner",
				Label: "TextAtCorner",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: PathAtArrowStart
			div := (&form.FormDiv{Name: "PathAtArrowStart"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PathAtArrowStart {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PathAtArrowStart",
				Label: "PathAtArrowStart",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: PathAtArrowEnd
			div := (&form.FormDiv{Name: "PathAtArrowEnd"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PathAtArrowEnd {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PathAtArrowEnd",
				Label: "PathAtArrowEnd",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: PathAtCorner
			div := (&form.FormDiv{Name: "PathAtCorner"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PathAtCorner {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PathAtCorner",
				Label: "PathAtCorner",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ControlPoints
			div := (&form.FormDiv{Name: "ControlPoints"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPoints {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPoints",
				Label: "ControlPoints",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MouseX", inst.MouseX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MouseY", inst.MouseY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("MouseEventKey", inst.MouseEventKey, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layers {
				for _, target := range src.Links {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layer", "Links", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.RectLinkLinks {
				if src.End == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RectLinkLink", "End", refNames, formGroup, probe.formStage)
		}
	case *models.LinkAnchoredPath:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Definition", inst.Definition, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X_Offset", inst.X_Offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y_Offset", inst.Y_Offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ScalePropotionnally", inst.ScalePropotionnally, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AppliedScaling", inst.AppliedScaling, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Links {
				for _, target := range src.PathAtArrowStart {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Link", "PathAtArrowStart", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Links {
				for _, target := range src.PathAtArrowEnd {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Link", "PathAtArrowEnd", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Links {
				for _, target := range src.PathAtCorner {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Link", "PathAtCorner", refNames, formGroup, probe.formStage)
		}
	case *models.LinkAnchoredText:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AutomaticLayout", inst.AutomaticLayout, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("LinkAnchorType", inst.LinkAnchorType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("X_Offset", inst.X_Offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y_Offset", inst.Y_Offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontWeight", inst.FontWeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontStyle", inst.FontStyle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LetterSpacing", inst.LetterSpacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("WhiteSpace", inst.WhiteSpace, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			// Slice of pointers: Animates
			div := (&form.FormDiv{Name: "Animates"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Animates {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Animates",
				Label: "Animates",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Links {
				for _, target := range src.TextAtArrowStart {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Link", "TextAtArrowStart", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Links {
				for _, target := range src.TextAtArrowEnd {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Link", "TextAtArrowEnd", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Links {
				for _, target := range src.TextAtCorner {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Link", "TextAtCorner", refNames, formGroup, probe.formStage)
		}
	case *models.Path:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Definition", inst.Definition, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			// Slice of pointers: Animates
			div := (&form.FormDiv{Name: "Animates"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Animates {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Animates",
				Label: "Animates",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layers {
				for _, target := range src.Paths {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layer", "Paths", refNames, formGroup, probe.formStage)
		}
	case *models.Point:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
	case *models.Polygone:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Points", inst.Points, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			// Slice of pointers: Animates
			div := (&form.FormDiv{Name: "Animates"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Animates {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Animates",
				Label: "Animates",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layers {
				for _, target := range src.Polygones {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layer", "Polygones", refNames, formGroup, probe.formStage)
		}
	case *models.Polyline:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Points", inst.Points, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			// Slice of pointers: Animates
			div := (&form.FormDiv{Name: "Animates"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Animates {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Animates",
				Label: "Animates",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layers {
				for _, target := range src.Polylines {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layer", "Polylines", refNames, formGroup, probe.formStage)
		}
	case *models.Rect:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RX", inst.RX, probe.formStage, formGroup)

		{
			// Slice of pointers: Peers
			div := (&form.FormDiv{Name: "Peers"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Peers {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Peers",
				Label: "Peers",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("EnclosingRect", inst.EnclosingRect, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), probe.formStage)

		{
			// Slice of pointers: Obstacles
			div := (&form.FormDiv{Name: "Obstacles"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Obstacles {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Obstacles",
				Label: "Obstacles",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("AnchoredTo", inst.AnchoredTo, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			// Slice of pointers: HoveringTrigger
			div := (&form.FormDiv{Name: "HoveringTrigger"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.HoveringTrigger {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "HoveringTrigger",
				Label: "HoveringTrigger",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DisplayConditions
			div := (&form.FormDiv{Name: "DisplayConditions"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DisplayConditions {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DisplayConditions",
				Label: "DisplayConditions",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Animations
			div := (&form.FormDiv{Name: "Animations"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Animations {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Animations",
				Label: "Animations",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsSelectable", inst.IsSelectable, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsSelected", inst.IsSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CanHaveLeftHandle", inst.CanHaveLeftHandle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasLeftHandle", inst.HasLeftHandle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CanHaveRightHandle", inst.CanHaveRightHandle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasRightHandle", inst.HasRightHandle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CanHaveTopHandle", inst.CanHaveTopHandle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasTopHandle", inst.HasTopHandle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsScalingProportionally", inst.IsScalingProportionally, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CanHaveBottomHandle", inst.CanHaveBottomHandle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasBottomHandle", inst.HasBottomHandle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CanMoveHorizontaly", inst.CanMoveHorizontaly, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CanMoveVerticaly", inst.CanMoveVerticaly, probe.formStage, formGroup)

		{
			// Slice of pointers: RectAnchoredTexts
			div := (&form.FormDiv{Name: "RectAnchoredTexts"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RectAnchoredTexts {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RectAnchoredTexts",
				Label: "RectAnchoredTexts",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RectAnchoredRects
			div := (&form.FormDiv{Name: "RectAnchoredRects"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RectAnchoredRects {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RectAnchoredRects",
				Label: "RectAnchoredRects",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RectAnchoredPaths
			div := (&form.FormDiv{Name: "RectAnchoredPaths"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RectAnchoredPaths {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RectAnchoredPaths",
				Label: "RectAnchoredPaths",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RectAnchoredPngImages
			div := (&form.FormDiv{Name: "RectAnchoredPngImages"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RectAnchoredPngImages {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RectAnchoredPngImages",
				Label: "RectAnchoredPngImages",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ChangeColorWhenHovered", inst.ChangeColorWhenHovered, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ColorWhenHovered", inst.ColorWhenHovered, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OriginalColor", inst.OriginalColor, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacityWhenHovered", inst.FillOpacityWhenHovered, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OriginalFillOpacity", inst.OriginalFillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasToolTip", inst.HasToolTip, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ToolTipText", inst.ToolTipText, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ToolTipPosition", inst.ToolTipPosition, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("MouseX", inst.MouseX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MouseY", inst.MouseY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("MouseEventKey", inst.MouseEventKey, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("URLPath", inst.URLPath, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("URLTarget", inst.URLTarget, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ControlPoints {
				if src.ClosestRect == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ControlPoint", "ClosestRect", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layers {
				for _, target := range src.Rects {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layer", "Rects", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Links {
				if src.Start == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Link", "Start", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Links {
				if src.End == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Link", "End", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rects {
				for _, target := range src.Peers {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rect", "Peers", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rects {
				if src.EnclosingRect == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rect", "EnclosingRect", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rects {
				for _, target := range src.Obstacles {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rect", "Obstacles", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rects {
				if src.AnchoredTo == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rect", "AnchoredTo", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.RectLinkLinks {
				if src.Start == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RectLinkLink", "Start", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.SVGs {
				if src.StartRect == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SVG", "StartRect", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.SVGs {
				if src.EndRect == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SVG", "EndRect", refNames, formGroup, probe.formStage)
		}
	case *models.RectAnchoredPath:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Definition", inst.Definition, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X_Offset", inst.X_Offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y_Offset", inst.Y_Offset, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("RectAnchorType", inst.RectAnchorType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ScalePropotionnally", inst.ScalePropotionnally, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AppliedScaling", inst.AppliedScaling, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rects {
				for _, target := range src.RectAnchoredPaths {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rect", "RectAnchoredPaths", refNames, formGroup, probe.formStage)
		}
	case *models.RectAnchoredPngImage:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RX", inst.RX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X_Offset", inst.X_Offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y_Offset", inst.Y_Offset, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("RectAnchorType", inst.RectAnchorType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Base64Content", inst.Base64Content, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rects {
				for _, target := range src.RectAnchoredPngImages {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rect", "RectAnchoredPngImages", refNames, formGroup, probe.formStage)
		}
	case *models.RectAnchoredRect:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RX", inst.RX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X_Offset", inst.X_Offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y_Offset", inst.Y_Offset, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("RectAnchorType", inst.RectAnchorType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("WidthFollowRect", inst.WidthFollowRect, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HeightFollowRect", inst.HeightFollowRect, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasToolTip", inst.HasToolTip, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ToolTipText", inst.ToolTipText, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rects {
				for _, target := range src.RectAnchoredRects {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rect", "RectAnchoredRects", refNames, formGroup, probe.formStage)
		}
	case *models.RectAnchoredText:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontWeight", inst.FontWeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontStyle", inst.FontStyle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LetterSpacing", inst.LetterSpacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("WhiteSpace", inst.WhiteSpace, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("X_Offset", inst.X_Offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y_Offset", inst.Y_Offset, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("RectAnchorType", inst.RectAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("TextAnchorType", inst.TextAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("DominantBaseline", inst.DominantBaseline, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("WritingMode", inst.WritingMode, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			// Slice of pointers: Animates
			div := (&form.FormDiv{Name: "Animates"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Animates {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Animates",
				Label: "Animates",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("URLPath", inst.URLPath, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("URLTarget", inst.URLTarget, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rects {
				for _, target := range src.RectAnchoredTexts {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rect", "RectAnchoredTexts", refNames, formGroup, probe.formStage)
		}
	case *models.RectLinkLink:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Start", inst.Start, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), probe.formStage)
		StageSetAssociationFieldToForm("End", inst.End, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Link](), probe.formStage)
		StageSetBasicFieldtoForm("TargetAnchorPosition", inst.TargetAnchorPosition, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layers {
				for _, target := range src.RectLinkLinks {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layer", "RectLinkLinks", refNames, formGroup, probe.formStage)
		}
	case *models.SVG:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Layers
			div := (&form.FormDiv{Name: "Layers"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Layers {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Layers",
				Label: "Layers",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetEnumStringFieldToForm("DrawingState", inst.DrawingState, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("StartRect", inst.StartRect, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), probe.formStage)
		StageSetAssociationFieldToForm("EndRect", inst.EndRect, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rect](), probe.formStage)
		StageSetBasicFieldtoForm("IsEditable", inst.IsEditable, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsSVGFrontEndFileGenerated", inst.IsSVGFrontEndFileGenerated, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsSVGBackEndFileGenerated", inst.IsSVGBackEndFileGenerated, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultDirectoryForGeneratedImages", inst.DefaultDirectoryForGeneratedImages, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsControlBannerHidden", inst.IsControlBannerHidden, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PanX", inst.PanX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PanY", inst.PanY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Zoom", inst.Zoom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OverrideWidth", inst.OverrideWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OverriddenWidth", inst.OverriddenWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OverrideHeight", inst.OverrideHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OverriddenHeight", inst.OverriddenHeight, probe.formStage, formGroup)
	case *models.SvgText:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
	case *models.Text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FillOpacity", inst.FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Stroke", inst.Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeOpacity", inst.StrokeOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeWidth", inst.StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArray", inst.StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StrokeDashArrayWhenSelected", inst.StrokeDashArrayWhenSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transform", inst.Transform, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontWeight", inst.FontWeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontStyle", inst.FontStyle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LetterSpacing", inst.LetterSpacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("WhiteSpace", inst.WhiteSpace, formGroup, probe.formStage)

		{
			// Slice of pointers: Animates
			div := (&form.FormDiv{Name: "Animates"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Animates {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Animates",
				Label: "Animates",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layers {
				for _, target := range src.Texts {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layer", "Texts", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
