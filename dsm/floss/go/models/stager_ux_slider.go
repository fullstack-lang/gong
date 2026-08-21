package models

import (
	"fmt"

	m "github.com/fullstack-lang/gong/lib/slider/go/models"
)

func (stager *Stager) ux_slider() {
	stager.sliderStage.Reset()

	layout := new(m.Layout).Stage(stager.sliderStage)

	// 3 separate groups for the 3 categories
	groupComplexities := new(m.Group).Stage(stager.sliderStage)
	groupComplexities.Name = "Complexities"
	groupComplexities.Percentage = 33.3
	layout.Groups = append(layout.Groups, groupComplexities)

	groupPerformances := new(m.Group).Stage(stager.sliderStage)
	groupPerformances.Name = "Performances"
	groupPerformances.Percentage = 33.3
	layout.Groups = append(layout.Groups, groupPerformances)

	groupEfforts := new(m.Group).Stage(stager.sliderStage)
	groupEfforts.Name = "Efforts"
	groupEfforts.Percentage = 33.4
	layout.Groups = append(layout.Groups, groupEfforts)

	// Find active DiagramFlossEquation & its CompareAnalysis or System
	var activeCompareAnalysis *CompareAnalysis
	var activeSystem *System
	var activeDiagram *DiagramFlossEquation
	for d := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
		if d.IsChecked {
			activeDiagram = d
			activeCompareAnalysis = d.GetOwningCompareAnalysis()
			activeSystem = d.GetOwningSystem()
			if activeCompareAnalysis == nil && activeSystem == nil {
				for ca := range *GetGongstructInstancesSet[CompareAnalysis](stager.stage) {
					for _, eqD := range ca.DiagramFlossEquations {
						if eqD == d {
							activeCompareAnalysis = ca
							break
						}
					}
				}
				if activeCompareAnalysis == nil {
					for sys := range *GetGongstructInstancesSet[System](stager.stage) {
						for _, eqD := range sys.DiagramFlossEquations {
							if eqD == d {
								activeSystem = sys
								break
							}
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
		if sysOwner != nil && sysOwner.AreCPEsCompoundedFromSubSystems && len(sysOwner.SubSystemes) > 0 {
			return true
		}
		return false
	}

	getSliderBounds := func(val float64) (min, max, step float64) {
		if val <= 2.0 {
			return 0.0, 2.0, 0.01
		}
		return 0.0, 100.0, 0.5
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
		slider := m.NewSlider(
			stager,
			label,
			min,
			max,
			step,
			&c.Strength,
		)
		if isElementDisabled(sysOwner) {
			slider.IsDisabled = true
		}
		groupComplexities.Sliders = append(groupComplexities.Sliders, slider)
	}

	// 2. Performances category (Alpha + performance elements in slice order)
	if activeCompareAnalysis != nil {
		groupPerformances.Sliders = append(
			groupPerformances.Sliders,
			m.NewSlider(
				stager,
				"Alpha (α)",
				0.1,
				5.0,
				0.01,
				&activeCompareAnalysis.Alpha,
			),
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
		slider := m.NewSlider(
			stager,
			label,
			min,
			max,
			step,
			&p.Strength,
		)
		if isElementDisabled(sysOwner) {
			slider.IsDisabled = true
		}
		groupPerformances.Sliders = append(groupPerformances.Sliders, slider)
	}

	// 3. Efforts category (Beta + effort elements in slice order)
	if activeCompareAnalysis != nil {
		groupEfforts.Sliders = append(
			groupEfforts.Sliders,
			m.NewSlider(
				stager,
				"Beta (β)",
				0.0,
				5.0,
				0.01,
				&activeCompareAnalysis.Beta,
			),
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
		slider := m.NewSlider(
			stager,
			label,
			min,
			max,
			step,
			&e.Strength,
		)
		if isElementDisabled(sysOwner) {
			slider.IsDisabled = true
		}
		groupEfforts.Sliders = append(groupEfforts.Sliders, slider)
	}

	stager.sliderStage.Commit()
}
