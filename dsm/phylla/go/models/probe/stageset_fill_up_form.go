// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/stool"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/clock"
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
	case *models.Angle0Shape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.Angle0Shape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeVase3DDiagram", "Angle0Shape", refNames, formGroup, probe.formStage)
		}
	case *models.BottomCurvePlane1Shape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.BottomCurvePlane1Shape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeVase3DDiagram", "BottomCurvePlane1Shape", refNames, formGroup, probe.formStage)
		}
	case *models.BottomCurvePlane2Shape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.BottomCurvePlane2Shape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeVase3DDiagram", "BottomCurvePlane2Shape", refNames, formGroup, probe.formStage)
		}
	case *models.Circumference3DShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.Circumference3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Plant3DDiagram", "Circumference3DShape", refNames, formGroup, probe.formStage)
		}
	case *models.Clock2DDiagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Zoom", inst.Zoom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenAxesShape", inst.IsHiddenAxesShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Clock2DDiagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "Clock2DDiagrams", refNames, formGroup, probe.formStage)
		}
	case *models.Clock3DDiagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenClockTopCurveShape", inst.IsHiddenClockTopCurveShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ClockTopCurveShape", inst.ClockTopCurveShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ClockTopCurveShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenTorus3DShape", inst.IsHiddenTorus3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Torus3DShape", inst.Torus3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Torus3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenSampledPoints3DShape", inst.IsHiddenSampledPoints3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SampledPoints3DShape", inst.SampledPoints3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SampledPoints3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenTiledFloor3DShape", inst.IsHiddenTiledFloor3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TiledFloor3DShape", inst.TiledFloor3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TiledFloor3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("Rendered3DShape", inst.Rendered3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rendered3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Clock3DDiagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "Clock3DDiagrams", refNames, formGroup, probe.formStage)
		}
	case *models.CutLine3DShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.CutLine3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Plant3DDiagram", "CutLine3DShape", refNames, formGroup, probe.formStage)
		}
	case *models.Leaves3DShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.Leaves3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Plant3DDiagram", "Leaves3DShape", refNames, formGroup, probe.formStage)
		}
	case *models.Library:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Plants
			div := (&form.FormDiv{Name: "Plants"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Plants {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Plants",
				Label: "Plants",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

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
		StageSetBasicFieldtoForm("NbPixPerCharacter", inst.NbPixPerCharacter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LogoSVGFile", inst.LogoSVGFile, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsRootLibrary", inst.IsRootLibrary, probe.formStage, formGroup)

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
	case *models.OriginalPoints3DShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.OriginalPoints3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeVase3DDiagram", "OriginalPoints3DShape", refNames, formGroup, probe.formStage)
		}
	case *models.ParastichyMCurves3DShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.ParastichyMCurves3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Plant3DDiagram", "ParastichyMCurves3DShape", refNames, formGroup, probe.formStage)
		}
	case *models.ParastichyNCurves3DShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.ParastichyNCurves3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Plant3DDiagram", "ParastichyNCurves3DShape", refNames, formGroup, probe.formStage)
		}
	case *models.Plant2DDiagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OriginX", inst.OriginX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OriginY", inst.OriginY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Zoom", inst.Zoom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsRhombusNodesExpanded", inst.IsRhombusNodesExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsArcNodesExpanded", inst.IsArcNodesExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenAxesShape", inst.IsHiddenAxesShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenReferenceRhombus", inst.IsHiddenReferenceRhombus, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenPlantCircumferenceShape", inst.IsHiddenPlantCircumferenceShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenGridPathShape", inst.IsHiddenGridPathShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenRhombusGridShape", inst.IsHiddenRhombusGridShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenExplanationTextShape", inst.IsHiddenExplanationTextShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenRotatedReferenceRhombus", inst.IsHiddenRotatedReferenceRhombus, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenRotatedPlantCircumferenceShape", inst.IsHiddenRotatedPlantCircumferenceShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenRotatedGridPathShape", inst.IsHiddenRotatedGridPathShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenRotatedRhombusGridShape", inst.IsHiddenRotatedRhombusGridShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenGrowthPathRhombusGridShape", inst.IsHiddenGrowthPathRhombusGridShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenGrowthVectorShape", inst.IsHiddenGrowthVectorShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenPerpendicularVectorGrid", inst.IsHiddenPerpendicularVectorGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenBaseVectorShapeGrid", inst.IsHiddenBaseVectorShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenArcNormalVectorShapeGrid", inst.IsHiddenArcNormalVectorShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStartArcShapeGrid", inst.IsHiddenStartArcShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenMidArcVectorShapeGrid", inst.IsHiddenMidArcVectorShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenEndArcShapeGrid", inst.IsHiddenEndArcShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenGrowthCurve2D", inst.IsHiddenGrowthCurve2D, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStackOfGrowthCurve2DByGrowthVector", inst.IsHiddenStackOfGrowthCurve2DByGrowthVector, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Plant2DDiagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "Plant2DDiagrams", refNames, formGroup, probe.formStage)
		}
	case *models.Plant3DDiagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStemCylinder3DShape", inst.IsHiddenStemCylinder3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("StemCylinder3DShape", inst.StemCylinder3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StemCylinder3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenParastichyNCurves3DShape", inst.IsHiddenParastichyNCurves3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ParastichyNCurves3DShape", inst.ParastichyNCurves3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ParastichyNCurves3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenParastichyMCurves3DShape", inst.IsHiddenParastichyMCurves3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ParastichyMCurves3DShape", inst.ParastichyMCurves3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ParastichyMCurves3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenCutLine3DShape", inst.IsHiddenCutLine3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("CutLine3DShape", inst.CutLine3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.CutLine3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenCircumference3DShape", inst.IsHiddenCircumference3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Circumference3DShape", inst.Circumference3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Circumference3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenTiledFloor3DShape", inst.IsHiddenTiledFloor3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TiledFloor3DShape", inst.TiledFloor3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TiledFloor3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenLeaves3DShape", inst.IsHiddenLeaves3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Leaves3DShape", inst.Leaves3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Leaves3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("Rendered3DShape", inst.Rendered3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rendered3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Plant3DDiagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "Plant3DDiagrams", refNames, formGroup, probe.formStage)
		}
	case *models.PlantAbstract:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("N", inst.N, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("M", inst.M, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackHeight", inst.StackHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RhombusInsideAngle", inst.RhombusInsideAngle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RhombusSideLength", inst.RhombusSideLength, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PlantType", inst.PlantType, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("TubeVaseAbstract", inst.TubeVaseAbstract, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TubeVaseAbstract](), probe.formStage)
		StageSetAssociationFieldToForm("StoolAbstract", inst.StoolAbstract, formGroup, probe.stageSet.StoolStage.GetInstancesSet[*stool.StoolAbstract](), probe.formStage)
		StageSetAssociationFieldToForm("ClockAbstract", inst.ClockAbstract, formGroup, probe.stageSet.ClockStage.GetInstancesSet[*clock.ClockAbstract](), probe.formStage)
		StageSetAssociationFieldToForm("MusicAbstract", inst.MusicAbstract, formGroup, probe.stageSet.MusicStage.GetInstancesSet[*music.MusicAbstract](), probe.formStage)
		StageSetEnumStringFieldToForm("CurrentView", inst.CurrentView, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsSelected", inst.IsSelected, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsPlant2DDiagramsNodeExpanded", inst.IsPlant2DDiagramsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Plant2DDiagrams
			div := (&form.FormDiv{Name: "Plant2DDiagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Plant2DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Plant2DDiagrams",
				Label: "Plant2DDiagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsPlant3DDiagramsNodeExpanded", inst.IsPlant3DDiagramsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Plant3DDiagrams
			div := (&form.FormDiv{Name: "Plant3DDiagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Plant3DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Plant3DDiagrams",
				Label: "Plant3DDiagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsVase2DDiagramsNodeExpanded", inst.IsVase2DDiagramsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Vase2DDiagrams
			div := (&form.FormDiv{Name: "Vase2DDiagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Vase2DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Vase2DDiagrams",
				Label: "Vase2DDiagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsTubeVase3DDiagramsNodeExpanded", inst.IsTubeVase3DDiagramsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: TubeVase3DDiagrams
			div := (&form.FormDiv{Name: "TubeVase3DDiagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TubeVase3DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TubeVase3DDiagrams",
				Label: "TubeVase3DDiagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsStool2DDiagramsNodeExpanded", inst.IsStool2DDiagramsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Stool2DDiagrams
			div := (&form.FormDiv{Name: "Stool2DDiagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Stool2DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Stool2DDiagrams",
				Label: "Stool2DDiagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsStool3DDiagramsNodeExpanded", inst.IsStool3DDiagramsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Stool3DDiagrams
			div := (&form.FormDiv{Name: "Stool3DDiagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Stool3DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Stool3DDiagrams",
				Label: "Stool3DDiagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsClock2DDiagramsNodeExpanded", inst.IsClock2DDiagramsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Clock2DDiagrams
			div := (&form.FormDiv{Name: "Clock2DDiagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Clock2DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Clock2DDiagrams",
				Label: "Clock2DDiagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsClock3DDiagramsNodeExpanded", inst.IsClock3DDiagramsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Clock3DDiagrams
			div := (&form.FormDiv{Name: "Clock3DDiagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Clock3DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Clock3DDiagrams",
				Label: "Clock3DDiagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("AxesShape", inst.AxesShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.AxesShape](), probe.formStage)
		StageSetAssociationFieldToForm("RhombusStuff", inst.RhombusStuff, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.RhombusStuff](), probe.formStage)
		StageSetAssociationFieldToForm("GrowthVectorShape", inst.GrowthVectorShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.GrowthVectorShape](), probe.formStage)
		StageSetAssociationFieldToForm("PerpendicularVectorGrid", inst.PerpendicularVectorGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PerpendicularVectorGrid](), probe.formStage)
		StageSetAssociationFieldToForm("BaseVectorShapeGrid", inst.BaseVectorShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.BaseVectorShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("ArcNormalVectorShapeGrid", inst.ArcNormalVectorShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ArcNormalVectorShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("StartArcShapeGrid", inst.StartArcShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StartArcShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("MidArcVectorShapeGrid", inst.MidArcVectorShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.MidArcVectorShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("EndArcShapeGrid", inst.EndArcShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.EndArcShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("GrowthCurve2D", inst.GrowthCurve2D, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.GrowthCurve2D](), probe.formStage)
		StageSetAssociationFieldToForm("StackOfGrowthCurve2DByGrowthVector", inst.StackOfGrowthCurve2DByGrowthVector, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StackOfGrowthCurve2DByGrowthVector](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.Plants {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "Plants", refNames, formGroup, probe.formStage)
		}
	case *models.Rendered3DShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ViewX", inst.ViewX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ViewY", inst.ViewY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ViewZ", inst.ViewZ, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TargetX", inst.TargetX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TargetY", inst.TargetY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TargetZ", inst.TargetZ, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Fov", inst.Fov, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Clock3DDiagrams {
				if src.Rendered3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Clock3DDiagram", "Rendered3DShape", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.Rendered3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Plant3DDiagram", "Rendered3DShape", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Stool3DDiagrams {
				if src.Rendered3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Stool3DDiagram", "Rendered3DShape", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.Rendered3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeVase3DDiagram", "Rendered3DShape", refNames, formGroup, probe.formStage)
		}
	case *models.SampledPoints3DShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Clock3DDiagrams {
				if src.SampledPoints3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Clock3DDiagram", "SampledPoints3DShape", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Stool3DDiagrams {
				if src.SampledPoints3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Stool3DDiagram", "SampledPoints3DShape", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.SampledPoints3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeVase3DDiagram", "SampledPoints3DShape", refNames, formGroup, probe.formStage)
		}
	case *models.StackOfRotatedVaseTrapezeRingsShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.StackOfRotatedVaseTrapezeRingsShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeVase3DDiagram", "StackOfRotatedVaseTrapezeRingsShape", refNames, formGroup, probe.formStage)
		}
	case *models.StackOfVaseTrapezeRingsShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.StackOfVaseTrapezeRingsShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeVase3DDiagram", "StackOfVaseTrapezeRingsShape", refNames, formGroup, probe.formStage)
		}
	case *models.StemCylinder3DShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transparency", inst.Transparency, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.StemCylinder3DShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Plant3DDiagram", "StemCylinder3DShape", refNames, formGroup, probe.formStage)
		}
	case *models.Stool2DDiagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Zoom", inst.Zoom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenAxesShape", inst.IsHiddenAxesShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Stool2DDiagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "Stool2DDiagrams", refNames, formGroup, probe.formStage)
		}
	case *models.Stool3DDiagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenSeatTopCurveShape", inst.IsHiddenSeatTopCurveShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SeatTopCurveShape", inst.SeatTopCurveShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SeatTopCurveShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenRotatedSeatTopCurveShape", inst.IsHiddenRotatedSeatTopCurveShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RotatedSeatTopCurveShape", inst.RotatedSeatTopCurveShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyRotatedSeatTopCurveShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenSeatBottomCurveShape", inst.IsHiddenSeatBottomCurveShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SeatBottomCurveShape", inst.SeatBottomCurveShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SeatBottomCurveShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenRotatedSeatBottomCurveShape", inst.IsHiddenRotatedSeatBottomCurveShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RotatedSeatBottomCurveShape", inst.RotatedSeatBottomCurveShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyRotatedSeatBottomCurveShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenTorus3DShape", inst.IsHiddenTorus3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Torus3DShape", inst.Torus3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Torus3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenRotatedTorusShape", inst.IsHiddenRotatedTorusShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RotatedTorusShape", inst.RotatedTorusShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyRotatedTorusShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenSampledPoints3DShape", inst.IsHiddenSampledPoints3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SampledPoints3DShape", inst.SampledPoints3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SampledPoints3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenRotatedSampledPoints3DShape", inst.IsHiddenRotatedSampledPoints3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RotatedSampledPoints3DShape", inst.RotatedSampledPoints3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.RotatedSampledPoints3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenEyeSampledPoints3DShape", inst.IsHiddenEyeSampledPoints3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("EyeSampledPoints3DShape", inst.EyeSampledPoints3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.EyeSampledPoints3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenEyeCornersSampledPoints3DShape", inst.IsHiddenEyeCornersSampledPoints3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("EyeCornersSampledPoints3DShape", inst.EyeCornersSampledPoints3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.EyeCornersSampledPoints3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenEye3DShape", inst.IsHiddenEye3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Eye3DShape", inst.Eye3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Eye3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenEyeSeatBottomCurveShape", inst.IsHiddenEyeSeatBottomCurveShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("EyeSeatBottomCurveShape", inst.EyeSeatBottomCurveShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.EyeSeatBottomCurveShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenEyeStoolBottomCurveShape", inst.IsHiddenEyeStoolBottomCurveShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("EyeStoolBottomCurveShape", inst.EyeStoolBottomCurveShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.EyeStoolBottomCurveShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenSeat3DShape", inst.IsHiddenSeat3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Seat3DShape", inst.Seat3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Seat3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenEyeVolume3DShape", inst.IsHiddenEyeVolume3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("EyeVolume3DShape", inst.EyeVolume3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.EyeVolume3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenSeatAndLegs3DShape", inst.IsHiddenSeatAndLegs3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SeatAndLegs3DShape", inst.SeatAndLegs3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SeatAndLegs3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenRotatedSeatAndLegs3DShape", inst.IsHiddenRotatedSeatAndLegs3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RotatedSeatAndLegs3DShape", inst.RotatedSeatAndLegs3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.RotatedSeatAndLegs3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsHiddenTiledFloor3DShape", inst.IsHiddenTiledFloor3DShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TiledFloor3DShape", inst.TiledFloor3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TiledFloor3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("Rendered3DShape", inst.Rendered3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rendered3DShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Stool3DDiagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "Stool3DDiagrams", refNames, formGroup, probe.formStage)
		}
	case *models.TopCurvePlane1Shape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.TopCurvePlane1Shape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeVase3DDiagram", "TopCurvePlane1Shape", refNames, formGroup, probe.formStage)
		}
	case *models.TopCurvePlane2Shape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.TopCurvePlane2Shape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeVase3DDiagram", "TopCurvePlane2Shape", refNames, formGroup, probe.formStage)
		}
	case *models.TubeVase3DDiagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon", inst.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTorusStackShape", inst.IsHiddenTorusStackShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenVerticalTorusStackShape", inst.IsHiddenVerticalTorusStackShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenPartiallyRotatedTorusShape", inst.IsHiddenPartiallyRotatedTorusShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStackOfPartiallyRotatedTorusShape", inst.IsHiddenStackOfPartiallyRotatedTorusShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenPointsAndLines3DShape", inst.IsHiddenPointsAndLines3DShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenKeyHole3DShape", inst.IsHiddenKeyHole3DShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenKey3DShape", inst.IsHiddenKey3DShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenVolumeKey3DShape", inst.IsHiddenVolumeKey3DShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTorusEdge3DShape", inst.IsHiddenTorusEdge3DShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenSampledPoints3DShape", inst.IsHiddenSampledPoints3DShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenOriginalPoints3DShape", inst.IsHiddenOriginalPoints3DShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenAngle0Shape", inst.IsHiddenAngle0Shape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTiledFloor3DShape", inst.IsHiddenTiledFloor3DShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTopCurvePlane1Shape", inst.IsHiddenTopCurvePlane1Shape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenBottomCurvePlane1Shape", inst.IsHiddenBottomCurvePlane1Shape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTopCurvePlane2Shape", inst.IsHiddenTopCurvePlane2Shape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenBottomCurvePlane2Shape", inst.IsHiddenBottomCurvePlane2Shape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenVaseTrapezeRingShape", inst.IsHiddenVaseTrapezeRingShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStackOfVaseTrapezeRingsShape", inst.IsHiddenStackOfVaseTrapezeRingsShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStackOfRotatedVaseTrapezeRingsShape", inst.IsHiddenStackOfRotatedVaseTrapezeRingsShape, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Rendered3DShape", inst.Rendered3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rendered3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("TorusStackShape", inst.TorusStackShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TorusStackShape](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalTorusStackShape", inst.VerticalTorusStackShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.VerticalTorusStackShape](), probe.formStage)
		StageSetAssociationFieldToForm("PartiallyRotatedTorusShape", inst.PartiallyRotatedTorusShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyRotatedTorusShape](), probe.formStage)
		StageSetAssociationFieldToForm("StackOfPartiallyRotatedTorusShape", inst.StackOfPartiallyRotatedTorusShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StackOfPartiallyRotatedTorusShape](), probe.formStage)
		StageSetAssociationFieldToForm("PointsAndLines3DShape", inst.PointsAndLines3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PointsAndLines3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("SampledPoints3DShape", inst.SampledPoints3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SampledPoints3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("OriginalPoints3DShape", inst.OriginalPoints3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.OriginalPoints3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("Angle0Shape", inst.Angle0Shape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Angle0Shape](), probe.formStage)
		StageSetAssociationFieldToForm("KeyHole3DShape", inst.KeyHole3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.KeyHole3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("Key3DShape", inst.Key3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Key3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("VolumeKey3DShape", inst.VolumeKey3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.VolumeKey3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("TorusEdge3DShape", inst.TorusEdge3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TorusEdge3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("TiledFloor3DShape", inst.TiledFloor3DShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TiledFloor3DShape](), probe.formStage)
		StageSetAssociationFieldToForm("TopCurvePlane1Shape", inst.TopCurvePlane1Shape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TopCurvePlane1Shape](), probe.formStage)
		StageSetAssociationFieldToForm("BottomCurvePlane1Shape", inst.BottomCurvePlane1Shape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.BottomCurvePlane1Shape](), probe.formStage)
		StageSetAssociationFieldToForm("TopCurvePlane2Shape", inst.TopCurvePlane2Shape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TopCurvePlane2Shape](), probe.formStage)
		StageSetAssociationFieldToForm("BottomCurvePlane2Shape", inst.BottomCurvePlane2Shape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.BottomCurvePlane2Shape](), probe.formStage)
		StageSetAssociationFieldToForm("VaseTrapezeRingShape", inst.VaseTrapezeRingShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.VaseTrapezeRingShape](), probe.formStage)
		StageSetAssociationFieldToForm("StackOfVaseTrapezeRingsShape", inst.StackOfVaseTrapezeRingsShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StackOfVaseTrapezeRingsShape](), probe.formStage)
		StageSetAssociationFieldToForm("StackOfRotatedVaseTrapezeRingsShape", inst.StackOfRotatedVaseTrapezeRingsShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StackOfRotatedVaseTrapezeRingsShape](), probe.formStage)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.TubeVase3DDiagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "TubeVase3DDiagrams", refNames, formGroup, probe.formStage)
		}
	case *models.TubeVaseAbstract:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Z_Ribbon", inst.Z_Ribbon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RibbonVerticalScale", inst.RibbonVerticalScale, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Plane1Height", inst.Plane1Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Plane2Height", inst.Plane2Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ProjectionAngle", inst.ProjectionAngle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeVerticalThickness", inst.RelativeVerticalThickness, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeRadialThickness", inst.RelativeRadialThickness, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeCuttedStackFloorHeight", inst.RelativeCuttedStackFloorHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeRotatedTorusSeparation", inst.RelativeRotatedTorusSeparation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RotationRatio", inst.RotationRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RadialRepetitions", inst.RadialRepetitions, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transparency", inst.Transparency, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasAlternatingRingColors", inst.HasAlternatingRingColors, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeTrajectoryOffsetX", inst.RelativeTrajectoryOffsetX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeTrajectoryOffsetY", inst.RelativeTrajectoryOffsetY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NbStepP1P2", inst.NbStepP1P2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ChosenStep", inst.ChosenStep, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeHorizontalRingsHeight", inst.RelativeHorizontalRingsHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OffsetKeyX", inst.OffsetKeyX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OffsetKeyY", inst.OffsetKeyY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HeightKey", inst.HeightKey, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("WidthKey", inst.WidthKey, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeKeySize", inst.RelativeKeySize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MovieNbFrames", inst.MovieNbFrames, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PerpendicularVectorGridHalfway", inst.PerpendicularVectorGridHalfway, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PerpendicularVectorGridHalfway](), probe.formStage)
		StageSetAssociationFieldToForm("TopStartArcShapeGrid", inst.TopStartArcShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TopStartArcShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("TopEndArcShapeGrid", inst.TopEndArcShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TopEndArcShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("ShiftedBottomTopStartArcShapeGrid", inst.ShiftedBottomTopStartArcShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ShiftedBottomTopStartArcShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("TopMidArcVectorShapeGrid", inst.TopMidArcVectorShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TopMidArcVectorShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("StartHalfwayArcShapeGrid", inst.StartHalfwayArcShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StartHalfwayArcShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("TopStartHalfwayArcShapeGrid", inst.TopStartHalfwayArcShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TopStartHalfwayArcShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("EndHalfwayArcShapeGrid", inst.EndHalfwayArcShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.EndHalfwayArcShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("TopEndHalfwayArcShapeGrid", inst.TopEndHalfwayArcShapeGrid, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TopEndHalfwayArcShapeGrid](), probe.formStage)
		StageSetAssociationFieldToForm("StackOfRotatedGrowthCurve2D", inst.StackOfRotatedGrowthCurve2D, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StackOfRotatedGrowthCurve2D](), probe.formStage)
		StageSetAssociationFieldToForm("TopStackOfRotatedGrowthCurve2D", inst.TopStackOfRotatedGrowthCurve2D, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TopStackOfRotatedGrowthCurve2D](), probe.formStage)
		StageSetAssociationFieldToForm("TopGrowthCurve2D", inst.TopGrowthCurve2D, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TopGrowthCurve2D](), probe.formStage)
		StageSetAssociationFieldToForm("StackOfGrowthCurve2D", inst.StackOfGrowthCurve2D, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StackOfGrowthCurve2D](), probe.formStage)
		StageSetAssociationFieldToForm("TopStackOfGrowthCurve2D", inst.TopStackOfGrowthCurve2D, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TopStackOfGrowthCurve2D](), probe.formStage)
		StageSetAssociationFieldToForm("StackOfGrowthCurve2DRibbon", inst.StackOfGrowthCurve2DRibbon, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StackOfGrowthCurve2DRibbon](), probe.formStage)
		StageSetAssociationFieldToForm("StackOfRotatedGrowthCurve2DRibbon", inst.StackOfRotatedGrowthCurve2DRibbon, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StackOfRotatedGrowthCurve2DRibbon](), probe.formStage)
		StageSetAssociationFieldToForm("GrowthCurve2DRibbon", inst.GrowthCurve2DRibbon, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.GrowthCurve2DRibbon](), probe.formStage)
		StageSetAssociationFieldToForm("ShiftedRightGrowthCurve2DRibbon", inst.ShiftedRightGrowthCurve2DRibbon, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ShiftedRightGrowthCurve2DRibbon](), probe.formStage)
		StageSetAssociationFieldToForm("ShiftedLeftGrowthCurve2DRibbon", inst.ShiftedLeftGrowthCurve2DRibbon, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ShiftedLeftGrowthCurve2DRibbon](), probe.formStage)
		StageSetAssociationFieldToForm("PartiallyGrowthCurve2DRibbon", inst.PartiallyGrowthCurve2DRibbon, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyGrowthCurve2DRibbon](), probe.formStage)
		StageSetAssociationFieldToForm("ShiftedLeftPartiallyGrowthCurve2DRibbon", inst.ShiftedLeftPartiallyGrowthCurve2DRibbon, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ShiftedLeftPartiallyGrowthCurve2DRibbon](), probe.formStage)
		StageSetAssociationFieldToForm("PartiallyGrowthCurve2DTrajectory", inst.PartiallyGrowthCurve2DTrajectory, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyGrowthCurve2DTrajectory](), probe.formStage)
		StageSetAssociationFieldToForm("PartiallyGrowthCurve2DTrajectoryP1P2", inst.PartiallyGrowthCurve2DTrajectoryP1P2, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PartiallyGrowthCurve2DTrajectoryP1P2](), probe.formStage)
		StageSetAssociationFieldToForm("PxShape", inst.PxShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PxShape](), probe.formStage)
		StageSetAssociationFieldToForm("ChosenP1P2PairShape", inst.ChosenP1P2PairShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ChosenP1P2PairShape](), probe.formStage)
		StageSetAssociationFieldToForm("KeyHoleShape", inst.KeyHoleShape, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.KeyHoleShape](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				if src.TubeVaseAbstract == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "TubeVaseAbstract", refNames, formGroup, probe.formStage)
		}
	case *models.Vase2DDiagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Zoom", inst.Zoom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsVaseArcNodesExpanded", inst.IsVaseArcNodesExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsVaseClampingNodesExpanded", inst.IsVaseClampingNodesExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenAxesShape", inst.IsHiddenAxesShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenBottomStartArcShapeGrid", inst.IsHiddenBottomStartArcShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenBottomEndArcShapeGrid", inst.IsHiddenBottomEndArcShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenBottomStackOfGrowthCurve", inst.IsHiddenBottomStackOfGrowthCurve, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenShiftedLeftStackOfGrowthCurve", inst.IsHiddenShiftedLeftStackOfGrowthCurve, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenShiftedLeftStackOfNormalVector", inst.IsHiddenShiftedLeftStackOfNormalVector, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenPerpendicularVectorGridHalfway", inst.IsHiddenPerpendicularVectorGridHalfway, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTopStartArcShapeGrid", inst.IsHiddenTopStartArcShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenShiftedBottomTopStartArcShapeGrid", inst.IsHiddenShiftedBottomTopStartArcShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTopMidArcVectorShapeGrid", inst.IsHiddenTopMidArcVectorShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStartHalfwayArcShapeGrid", inst.IsHiddenStartHalfwayArcShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTopStartHalfwayArcShapeGrid", inst.IsHiddenTopStartHalfwayArcShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenEndHalfwayArcShapeGrid", inst.IsHiddenEndHalfwayArcShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTopEndHalfwayArcShapeGrid", inst.IsHiddenTopEndHalfwayArcShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTopEndArcShapeGrid", inst.IsHiddenTopEndArcShapeGrid, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStackOfGrowthCurve", inst.IsHiddenStackOfGrowthCurve, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTopStackOfGrowthCurve", inst.IsHiddenTopStackOfGrowthCurve, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTopGrowthCurve2D", inst.IsHiddenTopGrowthCurve2D, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStackOfGrowthCurve2D", inst.IsHiddenStackOfGrowthCurve2D, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenTopStackOfGrowthCurve2D", inst.IsHiddenTopStackOfGrowthCurve2D, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenGrowthCurve2DRibbon", inst.IsHiddenGrowthCurve2DRibbon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenShiftedRightGrowthCurve2DRibbon", inst.IsHiddenShiftedRightGrowthCurve2DRibbon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenShiftedLeftGrowthCurve2DRibbon", inst.IsHiddenShiftedLeftGrowthCurve2DRibbon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStackOfGrowthCurve2DRibbon", inst.IsHiddenStackOfGrowthCurve2DRibbon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenStackOfRotatedGrowthCurve2DRibbon", inst.IsHiddenStackOfRotatedGrowthCurve2DRibbon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenPartiallyGrowthCurve2DRibbon", inst.IsHiddenPartiallyGrowthCurve2DRibbon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon", inst.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenPartiallyGrowthCurve2DTrajectory", inst.IsHiddenPartiallyGrowthCurve2DTrajectory, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2", inst.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenPxShape", inst.IsHiddenPxShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenChosenP1P2PairShape", inst.IsHiddenChosenP1P2PairShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHiddenKeyHoleShape", inst.IsHiddenKeyHoleShape, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Vase2DDiagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "Vase2DDiagrams", refNames, formGroup, probe.formStage)
		}
	case *models.VaseTrapezeRingShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.VaseTrapezeRingShape == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TubeVase3DDiagram", "VaseTrapezeRingShape", refNames, formGroup, probe.formStage)
		}
	case *stool.StoolAbstract:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RadialRepetitions", inst.RadialRepetitions, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transparency", inst.Transparency, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeTubeDiameter", inst.RelativeTubeDiameter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeHeight3DTorus", inst.RelativeHeight3DTorus, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StoolTorusVerticalScale", inst.StoolTorusVerticalScale, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeHeight", inst.RelativeHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeSeatThickness", inst.RelativeSeatThickness, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ProjectionAngle", inst.ProjectionAngle, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeEyeSeparationCriteria", inst.RelativeEyeSeparationCriteria, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeEyeCornerControlVectorStrength", inst.RelativeEyeCornerControlVectorStrength, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				if src.StoolAbstract == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "StoolAbstract", refNames, formGroup, probe.formStage)
		}
	case *music.MusicAbstract:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PitchHeight", inst.PitchHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NbOfBeatsInTheme", inst.NbOfBeatsInTheme, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BeatsPerSecond", inst.BeatsPerSecond, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FirstVoiceShiftX", inst.FirstVoiceShiftX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FirstVoiceShiftY", inst.FirstVoiceShiftY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PitchDifference", inst.PitchDifference, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Level", inst.Level, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ActualBeatsTemporalShift", inst.ActualBeatsTemporalShift, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsMinor", inst.IsMinor, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ThemeBinaryEncoding", inst.ThemeBinaryEncoding, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BezierControlLengthRatio", inst.BezierControlLengthRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NbPitchLines", inst.NbPitchLines, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NbBeatLines", inst.NbBeatLines, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OriginX", inst.OriginX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OriginY", inst.OriginY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ScoreScale", inst.ScoreScale, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowFirstVoice", inst.ShowFirstVoice, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowFirstVoiceShiftRight", inst.ShowFirstVoiceShiftRight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowSecondVoice", inst.ShowSecondVoice, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowSecondVoiceShiftRight", inst.ShowSecondVoiceShiftRight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowFirstVoiceNotes", inst.ShowFirstVoiceNotes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowFirstVoiceNotesShiftRight", inst.ShowFirstVoiceNotesShiftRight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowSecondVoiceNotes", inst.ShowSecondVoiceNotes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowSecondVoiceNotesShiftRight", inst.ShowSecondVoiceNotesShiftRight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsComposerNodeExpanded", inst.IsComposerNodeExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				if src.MusicAbstract == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "MusicAbstract", refNames, formGroup, probe.formStage)
		}
	case *clock.ClockAbstract:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RadialRepetitions", inst.RadialRepetitions, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Transparency", inst.Transparency, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeTubeDiameter", inst.RelativeTubeDiameter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeHeight3DTorus", inst.RelativeHeight3DTorus, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ClockTorusVerticalScale", inst.ClockTorusVerticalScale, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeHeight", inst.RelativeHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ProjectionAngle", inst.ProjectionAngle, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				if src.ClockAbstract == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PlantAbstract", "ClockAbstract", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
