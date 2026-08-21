package models

import (
	"fmt"
	"slices"

	m "github.com/fullstack-lang/gong/lib/slider/go/models"
)

func (stager *Stager) ux_slider() {
	stager.sliderStage.Reset()

	layout := new(m.Layout).Stage(stager.sliderStage)

	// Single unified panel for all sliders
	group := new(m.Group).Stage(stager.sliderStage)
	group.Name = "FLOSS Parameters & Element Weights"
	group.Percentage = 100
	layout.Groups = append(layout.Groups, group)

	// Find active DiagramFlossEquation & its CompareAnalysis
	var activeCompareAnalysis *CompareAnalysis
	for d := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
		if d.IsChecked {
			activeCompareAnalysis = d.GetOwningCompareAnalysis()
			if activeCompareAnalysis == nil {
				for ca := range *GetGongstructInstancesSet[CompareAnalysis](stager.stage) {
					for _, eqD := range ca.DiagramFlossEquations {
						if eqD == d {
							activeCompareAnalysis = ca
							break
						}
					}
				}
			}
			break
		}
	}

	if activeCompareAnalysis == nil {
		stager.sliderStage.Commit()
		return
	}

	// 1. Alpha slider
	group.Sliders = append(
		group.Sliders,
		m.NewSlider(
			stager,
			"Alpha (α)",
			0.1,
			5.0,
			0.1,
			&activeCompareAnalysis.Alpha,
		),
	)

	// Collect unique Complexities in the active diagram (from FromSystem and ToSystem)
	complexityMap := make(map[*Complexity]struct{})
	var complexities []*Complexity
	if activeCompareAnalysis.FromSystem != nil {
		for _, c := range activeCompareAnalysis.FromSystem.Complexities {
			if _, exists := complexityMap[c]; !exists {
				complexityMap[c] = struct{}{}
				complexities = append(complexities, c)
			}
		}
	}
	if activeCompareAnalysis.ToSystem != nil {
		for _, c := range activeCompareAnalysis.ToSystem.Complexities {
			if _, exists := complexityMap[c]; !exists {
				complexityMap[c] = struct{}{}
				complexities = append(complexities, c)
			}
		}
	}
	slices.SortFunc(complexities, func(a, b *Complexity) int {
		if a.Name < b.Name {
			return -1
		} else if a.Name > b.Name {
			return 1
		}
		return 0
	})

	for _, c := range complexities {
		group.Sliders = append(
			group.Sliders,
			m.NewSlider(
				stager,
				fmt.Sprintf("Complexity: %s", c.Name),
				0.0,
				100.0,
				0.5,
				&c.Strength,
			),
		)
	}

	// Collect unique Performances in the active diagram (from FromSystem and ToSystem)
	performanceMap := make(map[*Performance]struct{})
	var performances []*Performance
	if activeCompareAnalysis.FromSystem != nil {
		for _, p := range activeCompareAnalysis.FromSystem.Performances {
			if _, exists := performanceMap[p]; !exists {
				performanceMap[p] = struct{}{}
				performances = append(performances, p)
			}
		}
	}
	if activeCompareAnalysis.ToSystem != nil {
		for _, p := range activeCompareAnalysis.ToSystem.Performances {
			if _, exists := performanceMap[p]; !exists {
				performanceMap[p] = struct{}{}
				performances = append(performances, p)
			}
		}
	}
	slices.SortFunc(performances, func(a, b *Performance) int {
		if a.Name < b.Name {
			return -1
		} else if a.Name > b.Name {
			return 1
		}
		return 0
	})

	for _, p := range performances {
		group.Sliders = append(
			group.Sliders,
			m.NewSlider(
				stager,
				fmt.Sprintf("Performance: %s", p.Name),
				0.0,
				100.0,
				0.5,
				&p.Strength,
			),
		)
	}

	// Collect unique Efforts in the active diagram (from FromSystem and ToSystem)
	effortMap := make(map[*Effort]struct{})
	var efforts []*Effort
	if activeCompareAnalysis.FromSystem != nil {
		for _, e := range activeCompareAnalysis.FromSystem.Efforts {
			if _, exists := effortMap[e]; !exists {
				effortMap[e] = struct{}{}
				efforts = append(efforts, e)
			}
		}
	}
	if activeCompareAnalysis.ToSystem != nil {
		for _, e := range activeCompareAnalysis.ToSystem.Efforts {
			if _, exists := effortMap[e]; !exists {
				effortMap[e] = struct{}{}
				efforts = append(efforts, e)
			}
		}
	}
	slices.SortFunc(efforts, func(a, b *Effort) int {
		if a.Name < b.Name {
			return -1
		} else if a.Name > b.Name {
			return 1
		}
		return 0
	})

	for _, e := range efforts {
		group.Sliders = append(
			group.Sliders,
			m.NewSlider(
				stager,
				fmt.Sprintf("Effort: %s", e.Name),
				0.0,
				100.0,
				0.5,
				&e.Strength,
			),
		)
	}

	stager.sliderStage.Commit()
}
