package models

import (
	"fmt"
	"math"
	"slices"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

func (stager *Stager) ux_slider() {
	stager.sliderStage.Reset()

	// Find active DiagramFlossEquation & its CompareAnalysis or System
	var activeCompareAnalysis *CompareAnalysis
	var activeSystem *System
	var activeDiagram *DiagramFlossEquation
	for d := range *stager.stage.GetInstancesSet[*DiagramFlossEquation]() {
		if d.IsChecked {
			activeDiagram = d
			activeCompareAnalysis = d.GetOwningCompareAnalysis()
			activeSystem = d.GetOwningSystem()
			if activeCompareAnalysis == nil && activeSystem == nil {
				for ca := range *stager.stage.GetInstancesSet[*CompareAnalysis]() {
					if slices.Contains(ca.DiagramFlossEquations, d) {
						activeCompareAnalysis = ca
					}
				}
				if activeCompareAnalysis == nil {
					for sys := range *stager.stage.GetInstancesSet[*System]() {
						if slices.Contains(sys.DiagramFlossEquations, d) {
							activeSystem = sys
						}
					}
				}
			}
			break
		}
	}

	if activeCompareAnalysis == nil && activeSystem == nil {
		stager.sliderStage.Commit()
		return
	}

	getScaleSliderBounds := func(val float64) (min, max, step float64) {
		min = 0.1
		max = 100.0
		if val > 100.0 {
			max = math.Max(100.0, math.Ceil(val*2.0/50.0)*50.0)
		}
		if val >= 50.0 {
			step = 1.0
		} else {
			step = 0.1
		}
		return min, max, step
	}

	formGroup := form.NewLiveFormGroup(stager.sliderStage, "Floss Controls")

	if activeDiagram != nil {
		form.StartAccordion(stager.sliderStage, formGroup, "Diagram", true)
		min, max, step := getScaleSliderBounds(activeDiagram.Scale)
		form.AppendSlider(
			stager.sliderStage,
			formGroup,
			"Scale",
			min,
			max,
			step,
			&activeDiagram.Scale,
			stager.OnAfterUpdateSliderElement,
		)
	}

	showSubsystems := activeDiagram == nil || activeDiagram.AreSubsystemsVisible

	// 1. Complexities category (preserves slice order and compounds subsystems)
	complexityMap := make(map[*Complexity]struct{})
	var complexities []*Complexity
	cSysMap := make(map[*Complexity]*System)

	if activeCompareAnalysis != nil {
		if activeCompareAnalysis.ToSystem != nil {
			var comps []*Complexity
			var sm map[*Complexity]*System
			if showSubsystems {
				comps, sm = activeCompareAnalysis.ToSystem.GetEffectiveComplexities()
			} else {
				comps = activeCompareAnalysis.ToSystem.Complexities
				sm = make(map[*Complexity]*System)
			}
			for _, c := range comps {
				if _, exists := complexityMap[c]; !exists {
					complexityMap[c] = struct{}{}
					complexities = append(complexities, c)
					cSysMap[c] = sm[c]
				}
			}
		}
		if activeCompareAnalysis.FromSystem != nil {
			var comps []*Complexity
			var sm map[*Complexity]*System
			if showSubsystems {
				comps, sm = activeCompareAnalysis.FromSystem.GetEffectiveComplexities()
			} else {
				comps = activeCompareAnalysis.FromSystem.Complexities
				sm = make(map[*Complexity]*System)
			}
			for _, c := range comps {
				if _, exists := complexityMap[c]; !exists {
					complexityMap[c] = struct{}{}
					complexities = append(complexities, c)
					cSysMap[c] = sm[c]
				}
			}
		}
	} else if activeSystem != nil {
		var comps []*Complexity
		var sm map[*Complexity]*System
		if showSubsystems {
			comps, sm = activeSystem.GetEffectiveComplexities()
		} else {
			comps = activeSystem.Complexities
			sm = make(map[*Complexity]*System)
		}
		for _, c := range comps {
			if _, exists := complexityMap[c]; !exists {
				complexityMap[c] = struct{}{}
				complexities = append(complexities, c)
				cSysMap[c] = sm[c]
			}
		}
	}

	isElementDisabled := func(sysOwner *System) bool {
		if showSubsystems {
			return false
		}
		if sysOwner != nil && sysOwner.AreCPEsCompoundedFromSubSystems && len(sysOwner.SubSystems) > 0 {
			return true
		}
		return false
	}

	getSliderBounds := func(val float64) (min, max, step float64) {
		max = 100.0
		if val > 100.0 {
			max = math.Ceil(val/50.0) * 50.0
		}
		return 0.0, max, 0.1
	}

	if len(complexities) > 0 {
		form.StartAccordion(stager.sliderStage, formGroup, "Complexities", true)
	}
	for _, c := range complexities {
		sysOwner := cSysMap[c]
		if sysOwner == nil && activeSystem != nil {
			sysOwner = activeSystem
		}
		label := c.Name
		if sysOwner != nil && activeSystem != nil && sysOwner != activeSystem {
			label = fmt.Sprintf("[%s] %s", sysOwner.Name, c.Name)
		}
		min, max, step := getSliderBounds(c.Strength)
		field := form.AppendSlider(
			stager.sliderStage,
			formGroup,
			label,
			min,
			max,
			step,
			&c.Strength,
			stager.OnAfterUpdateSliderElement,
		)
		if isElementDisabled(sysOwner) {
			field.IsDisabled = true
		}
	}

	// 2. Performances category (Mu + performance elements in slice order)
	form.StartAccordion(stager.sliderStage, formGroup, "Performances", true)
	if activeCompareAnalysis != nil {
		form.AppendSlider(
			stager.sliderStage,
			formGroup,
			"Mu (μ)",
			0.1,
			5.0,
			0.1,
			&activeCompareAnalysis.Mu,
			stager.OnAfterUpdateSliderElement,
		)
	}

	performanceMap := make(map[*Performance]struct{})
	var performances []*Performance
	pSysMap := make(map[*Performance]*System)

	if activeCompareAnalysis != nil {
		if activeCompareAnalysis.ToSystem != nil {
			var perfs []*Performance
			var sm map[*Performance]*System
			if showSubsystems {
				perfs, sm = activeCompareAnalysis.ToSystem.GetEffectivePerformances()
			} else {
				perfs = activeCompareAnalysis.ToSystem.Performances
				sm = make(map[*Performance]*System)
			}
			for _, p := range perfs {
				if _, exists := performanceMap[p]; !exists {
					performanceMap[p] = struct{}{}
					performances = append(performances, p)
					pSysMap[p] = sm[p]
				}
			}
		}
		if activeCompareAnalysis.FromSystem != nil {
			var perfs []*Performance
			var sm map[*Performance]*System
			if showSubsystems {
				perfs, sm = activeCompareAnalysis.FromSystem.GetEffectivePerformances()
			} else {
				perfs = activeCompareAnalysis.FromSystem.Performances
				sm = make(map[*Performance]*System)
			}
			for _, p := range perfs {
				if _, exists := performanceMap[p]; !exists {
					performanceMap[p] = struct{}{}
					performances = append(performances, p)
					pSysMap[p] = sm[p]
				}
			}
		}
	} else if activeSystem != nil {
		var perfs []*Performance
		var sm map[*Performance]*System
		if showSubsystems {
			perfs, sm = activeSystem.GetEffectivePerformances()
		} else {
			perfs = activeSystem.Performances
			sm = make(map[*Performance]*System)
		}
		for _, p := range perfs {
			if _, exists := performanceMap[p]; !exists {
				performanceMap[p] = struct{}{}
				performances = append(performances, p)
				pSysMap[p] = sm[p]
			}
		}
	}

	for _, p := range performances {
		sysOwner := pSysMap[p]
		if sysOwner == nil && activeSystem != nil {
			sysOwner = activeSystem
		}
		label := p.Name
		if sysOwner != nil && activeSystem != nil && sysOwner != activeSystem {
			label = fmt.Sprintf("[%s] %s", sysOwner.Name, p.Name)
		}
		min, max, step := getSliderBounds(p.Strength)
		field := form.AppendSlider(
			stager.sliderStage,
			formGroup,
			label,
			min,
			max,
			step,
			&p.Strength,
			stager.OnAfterUpdateSliderElement,
		)
		if isElementDisabled(sysOwner) {
			field.IsDisabled = true
		}
	}

	// 3. Efforts category (Epsilon + effort elements in slice order)
	form.StartAccordion(stager.sliderStage, formGroup, "Efforts", true)
	if activeCompareAnalysis != nil {
		form.AppendSlider(
			stager.sliderStage,
			formGroup,
			"Epsilon (ε)",
			0.0,
			5.0,
			0.1,
			&activeCompareAnalysis.Epsilon,
			stager.OnAfterUpdateSliderElement,
		)
	}

	effortMap := make(map[*Effort]struct{})
	var efforts []*Effort
	eSysMap := make(map[*Effort]*System)

	if activeCompareAnalysis != nil {
		if activeCompareAnalysis.ToSystem != nil {
			var effs []*Effort
			var sm map[*Effort]*System
			if showSubsystems {
				effs, sm = activeCompareAnalysis.ToSystem.GetEffectiveEfforts()
			} else {
				effs = activeCompareAnalysis.ToSystem.Efforts
				sm = make(map[*Effort]*System)
			}
			for _, e := range effs {
				if _, exists := effortMap[e]; !exists {
					effortMap[e] = struct{}{}
					efforts = append(efforts, e)
					eSysMap[e] = sm[e]
				}
			}
		}
		if activeCompareAnalysis.FromSystem != nil {
			var effs []*Effort
			var sm map[*Effort]*System
			if showSubsystems {
				effs, sm = activeCompareAnalysis.FromSystem.GetEffectiveEfforts()
			} else {
				effs = activeCompareAnalysis.FromSystem.Efforts
				sm = make(map[*Effort]*System)
			}
			for _, e := range effs {
				if _, exists := effortMap[e]; !exists {
					effortMap[e] = struct{}{}
					efforts = append(efforts, e)
					eSysMap[e] = sm[e]
				}
			}
		}
	} else if activeSystem != nil {
		var effs []*Effort
		var sm map[*Effort]*System
		if showSubsystems {
			effs, sm = activeSystem.GetEffectiveEfforts()
		} else {
			effs = activeSystem.Efforts
			sm = make(map[*Effort]*System)
		}
		for _, e := range effs {
			if _, exists := effortMap[e]; !exists {
				effortMap[e] = struct{}{}
				efforts = append(efforts, e)
				eSysMap[e] = sm[e]
			}
		}
	}

	for _, e := range efforts {
		sysOwner := eSysMap[e]
		if sysOwner == nil && activeSystem != nil {
			sysOwner = activeSystem
		}
		label := e.Name
		if sysOwner != nil && activeSystem != nil && sysOwner != activeSystem {
			label = fmt.Sprintf("[%s] %s", sysOwner.Name, e.Name)
		}
		min, max, step := getSliderBounds(e.Strength)
		field := form.AppendSlider(
			stager.sliderStage,
			formGroup,
			label,
			min,
			max,
			step,
			&e.Strength,
			stager.OnAfterUpdateSliderElement,
		)
		if isElementDisabled(sysOwner) {
			field.IsDisabled = true
		}
	}

	form.EndAccordion(stager.sliderStage, formGroup)

	stager.sliderStage.Commit()
}
