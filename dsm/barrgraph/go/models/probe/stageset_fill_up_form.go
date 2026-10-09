// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
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
	case *models.ArtefactType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ArtefactTypeShapes {
				if src.ArtefactType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ArtefactTypeShape", "ArtefactType", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Influences {
				if src.SourceArtefactType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Influence", "SourceArtefactType", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Influences {
				if src.TargetArtefactType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Influence", "TargetArtefactType", refNames, formGroup, probe.formStage)
		}
	case *models.ArtefactTypeShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ArtefactType", inst.ArtefactType, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ArtefactType](), probe.formStage)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ArtefactTypeShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ArtefactTypeShapes", refNames, formGroup, probe.formStage)
		}
	case *models.Artist:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsDead", inst.IsDead, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DateOfDeath", inst.DateOfDeath, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Place", inst.Place, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Place](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ArtistShapes {
				if src.Artist == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ArtistShape", "Artist", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Influences {
				if src.SourceArtist == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Influence", "SourceArtist", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Influences {
				if src.TargetArtist == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Influence", "TargetArtist", refNames, formGroup, probe.formStage)
		}
	case *models.ArtistShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Artist", inst.Artist, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Artist](), probe.formStage)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ImagePng_X", inst.ImagePng_X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ImagePng_Y", inst.ImagePng_Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ImagePng_Width", inst.ImagePng_Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ImagePng_Height", inst.ImagePng_Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ImagePng_X_Offset", inst.ImagePng_X_Offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ImagePng_Y_Offset", inst.ImagePng_Y_Offset, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ImagePng_RectAnchorType", inst.ImagePng_RectAnchorType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ImagePngBase64Content", inst.ImagePngBase64Content, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ArtistShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ArtistShapes", refNames, formGroup, probe.formStage)
		}
	case *models.ControlPointShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X_Relative", inst.X_Relative, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y_Relative", inst.Y_Relative, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsStartShapeTheClosestShape", inst.IsStartShapeTheClosestShape, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.InfluenceShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.InfluenceShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}
	case *models.Desk:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SelectedDiagram", inst.SelectedDiagram, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Diagram](), probe.formStage)
	case *models.Diagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)

		{
			// Slice of pointers: MovementShapes
			div := (&form.FormDiv{Name: "MovementShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MovementShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MovementShapes",
				Label: "MovementShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ArtefactTypeShapes
			div := (&form.FormDiv{Name: "ArtefactTypeShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ArtefactTypeShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ArtefactTypeShapes",
				Label: "ArtefactTypeShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ArtistShapes
			div := (&form.FormDiv{Name: "ArtistShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ArtistShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ArtistShapes",
				Label: "ArtistShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: InfluenceShapes
			div := (&form.FormDiv{Name: "InfluenceShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.InfluenceShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "InfluenceShapes",
				Label: "InfluenceShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsEditable", inst.IsEditable, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsNodeExpanded", inst.IsNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsMovementCategoryNodeExpanded", inst.IsMovementCategoryNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsArtefactTypeCategoryNodeExpanded", inst.IsArtefactTypeCategoryNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsArtistCategoryNodeExpanded", inst.IsArtistCategoryNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsInfluenceCategoryNodeExpanded", inst.IsInfluenceCategoryNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsMovementCategoryHidden", inst.IsMovementCategoryHidden, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsArtefactTypeCategoryHidden", inst.IsArtefactTypeCategoryHidden, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsArtistCategoryHidden", inst.IsArtistCategoryHidden, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsInfluenceCategoryHidden", inst.IsInfluenceCategoryHidden, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StartDate", inst.StartDate, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndDate", inst.EndDate, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NbYearsForIntervals", inst.NbYearsForIntervals, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("XMargin", inst.XMargin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("YMargin", inst.YMargin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NextVerticalDateXMargin", inst.NextVerticalDateXMargin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RedColorCode", inst.RedColorCode, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BackgroundGreyColorCode", inst.BackgroundGreyColorCode, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("GrayColorCode", inst.GrayColorCode, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BottomBoxYOffset", inst.BottomBoxYOffset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BottomBoxWidth", inst.BottomBoxWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BottomBoxHeigth", inst.BottomBoxHeigth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BottomBoxFontSize", inst.BottomBoxFontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BottomBoxFontWeigth", inst.BottomBoxFontWeigth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BottomBoxFontFamily", inst.BottomBoxFontFamily, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BottomBoxLetterSpacing", inst.BottomBoxLetterSpacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BottomBoxLetterColorCode", inst.BottomBoxLetterColorCode, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("MovementRectAnchorType", inst.MovementRectAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("MovementTextAnchorType", inst.MovementTextAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("MovementDominantBaselineType", inst.MovementDominantBaselineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("MovementFontSize", inst.MovementFontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MajorMovementFontSize", inst.MajorMovementFontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinorMovementFontSize", inst.MinorMovementFontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MovementFontWeigth", inst.MovementFontWeigth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MovementFontFamily", inst.MovementFontFamily, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MovementLetterSpacing", inst.MovementLetterSpacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AbstractMovementFontSize", inst.AbstractMovementFontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("AbstractMovementRectAnchorType", inst.AbstractMovementRectAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("AbstractMovementTextAnchorType", inst.AbstractMovementTextAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("AbstractDominantBaselineType", inst.AbstractDominantBaselineType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("MovementDateRectAnchorType", inst.MovementDateRectAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("MovementDateTextAnchorType", inst.MovementDateTextAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("MovementDateTextDominantBaselineType", inst.MovementDateTextDominantBaselineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("MovementDateAndPlacesFontSize", inst.MovementDateAndPlacesFontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MovementDateAndPlacesFontWeigth", inst.MovementDateAndPlacesFontWeigth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MovementDateAndPlacesFontFamily", inst.MovementDateAndPlacesFontFamily, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MovementDateAndPlacesLetterSpacing", inst.MovementDateAndPlacesLetterSpacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MovementBelowArcY_Offset", inst.MovementBelowArcY_Offset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MovementBelowArcY_OffsetPerPlace", inst.MovementBelowArcY_OffsetPerPlace, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("MovementPlacesRectAnchorType", inst.MovementPlacesRectAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("MovementPlacesTextAnchorType", inst.MovementPlacesTextAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("MovementPlacesDominantBaselineType", inst.MovementPlacesDominantBaselineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ArtefactTypeFontSize", inst.ArtefactTypeFontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArtefactTypeFontWeigth", inst.ArtefactTypeFontWeigth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArtefactTypeFontFamily", inst.ArtefactTypeFontFamily, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArtefactTypeLetterSpacing", inst.ArtefactTypeLetterSpacing, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ArtefactTypeRectAnchorType", inst.ArtefactTypeRectAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ArtefactDominantBaselineType", inst.ArtefactDominantBaselineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ArtefactTypeStrokeWidth", inst.ArtefactTypeStrokeWidth, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ArtistRectAnchorType", inst.ArtistRectAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ArtistTextAnchorType", inst.ArtistTextAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ArtistDominantBaselineType", inst.ArtistDominantBaselineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ArtistFontSize", inst.ArtistFontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MajorArtistFontSize", inst.MajorArtistFontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinorArtistFontSize", inst.MinorArtistFontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArtistFontWeigth", inst.ArtistFontWeigth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArtistFontFamily", inst.ArtistFontFamily, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArtistLetterSpacing", inst.ArtistLetterSpacing, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ArtistDateRectAnchorType", inst.ArtistDateRectAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ArtistDateTextAnchorType", inst.ArtistDateTextAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ArtistDateDominantBaselineType", inst.ArtistDateDominantBaselineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ArtistDateAndPlacesFontSize", inst.ArtistDateAndPlacesFontSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArtistDateAndPlacesFontWeigth", inst.ArtistDateAndPlacesFontWeigth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArtistDateAndPlacesFontFamily", inst.ArtistDateAndPlacesFontFamily, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArtistDateAndPlacesLetterSpacing", inst.ArtistDateAndPlacesLetterSpacing, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ArtistPlacesRectAnchorType", inst.ArtistPlacesRectAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ArtistPlacesTextAnchorType", inst.ArtistPlacesTextAnchorType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ArtistPlacesDominantBaselineType", inst.ArtistPlacesDominantBaselineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("InfluenceArrowSize", inst.InfluenceArrowSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("InfluenceArrowStartOffset", inst.InfluenceArrowStartOffset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("InfluenceArrowEndOffset", inst.InfluenceArrowEndOffset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("InfluenceCornerRadius", inst.InfluenceCornerRadius, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("InfluenceDashedLinePattern", inst.InfluenceDashedLinePattern, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Desks {
				if src.SelectedDiagram == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Desk", "SelectedDiagram", refNames, formGroup, probe.formStage)
		}
	case *models.Influence:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SourceMovement", inst.SourceMovement, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Movement](), probe.formStage)
		StageSetAssociationFieldToForm("SourceArtefactType", inst.SourceArtefactType, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ArtefactType](), probe.formStage)
		StageSetAssociationFieldToForm("SourceArtist", inst.SourceArtist, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Artist](), probe.formStage)
		StageSetAssociationFieldToForm("TargetMovement", inst.TargetMovement, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Movement](), probe.formStage)
		StageSetAssociationFieldToForm("TargetArtefactType", inst.TargetArtefactType, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ArtefactType](), probe.formStage)
		StageSetAssociationFieldToForm("TargetArtist", inst.TargetArtist, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Artist](), probe.formStage)
		StageSetBasicFieldtoForm("IsHypothtical", inst.IsHypothtical, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.InfluenceShapes {
				if src.Influence == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.InfluenceShape", "Influence", refNames, formGroup, probe.formStage)
		}
	case *models.InfluenceShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Influence", inst.Influence, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Influence](), probe.formStage)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.InfluenceShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "InfluenceShapes", refNames, formGroup, probe.formStage)
		}
	case *models.Library:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsRootLibrary", inst.IsRootLibrary, probe.formStage, formGroup)

		{
			// Slice of pointers: SubLibraries
			div := (&form.FormDiv{Name: "SubLibraries"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubLibraries {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubLibraries",
				Label: "SubLibraries",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsSubLibrariesNodeExpanded", inst.IsSubLibrariesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: SubLibrariesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "SubLibrariesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubLibrariesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubLibrariesWhoseNodeIsExpanded",
				Label: "SubLibrariesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("NbPixPerCharacter", inst.NbPixPerCharacter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LogoSVGFile", inst.LogoSVGFile, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpandedTmp", inst.IsExpandedTmp, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.SubLibraries {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "SubLibraries", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.SubLibrariesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "SubLibrariesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}
	case *models.Movement:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Date", inst.Date, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HideDate", inst.HideDate, probe.formStage, formGroup)

		{
			// Slice of pointers: Places
			div := (&form.FormDiv{Name: "Places"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Places {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Places",
				Label: "Places",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("HasTaxonomicFilter", inst.HasTaxonomicFilter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TaxonomicFilter", inst.TaxonomicFilter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsFeatured", inst.IsFeatured, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FeaturePrefix", inst.FeaturePrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsMajor", inst.IsMajor, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsMinor", inst.IsMinor, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AdditionnalName", inst.AdditionnalName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Influences {
				if src.SourceMovement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Influence", "SourceMovement", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Influences {
				if src.TargetMovement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Influence", "TargetMovement", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.MovementShapes {
				if src.Movement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.MovementShape", "Movement", refNames, formGroup, probe.formStage)
		}
	case *models.MovementShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Movement", inst.Movement, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Movement](), probe.formStage)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.MovementShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "MovementShapes", refNames, formGroup, probe.formStage)
		}
	case *models.Place:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Artists {
				if src.Place == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Artist", "Place", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Movements {
				for _, target := range src.Places {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Movement", "Places", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
