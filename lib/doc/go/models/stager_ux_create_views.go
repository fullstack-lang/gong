package models

import (
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	splitlite "github.com/fullstack-lang/gong/lib/splitlite/go/models"
)

func (stager *Stager) sidebar() []*split.AsSplitArea {
	if !stager.embeddedDiagrams {

		return []*split.AsSplitArea{
			{
				Name:             "doc Tree",
				ShowNameInHeader: false,
				Size:             53,
				Tree: &split.Tree{
					StackName: stager.treeNavigationStage.GetName(),
				},
			},
			{
				Name:             "doc Tree",
				ShowNameInHeader: false,
				IsAny:            true,
				Tree: &split.Tree{
					StackName: stager.treeStage.GetName(),
				},
			},
			{
				Name: "temporary form stack",
				Size: 24,
				Form: &split.Form{
					StackName: stager.formStage.GetName(),
				},
			},
		}
	} else {
		return []*split.AsSplitArea{
			{
				Name:             "doc Tree",
				ShowNameInHeader: false,
				IsAny:            true,
				Tree: &split.Tree{
					StackName: stager.treeStage.GetName(),
				},
			},
		}
	}
}

func (stager *Stager) createViews(receivingAsSplitArea *split.AsSplitArea) {
	receivingAsSplitArea.AsSplit = &split.AsSplit{
		Name:                   "Root As Split for doc receiving area",
		Direction:              split.Horizontal,
		IsWithCustomGutterSize: true,
		GutterSize:             0,
		AsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "AsSplitArea 50% for Slit (Tree & Svg)",
				ShowNameInHeader: false,
				AsSplit: (&split.AsSplit{
					Direction:              split.Horizontal,
					IsWithCustomGutterSize: true,
					GutterSize:             1,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 25,
							AsSplit: &split.AsSplit{
								IsWithCustomGutterSize: true,
								GutterSize:             1,
								IsSizeInPixel:          true,
								Direction:              split.Vertical,
								AsSplitAreas:           stager.sidebar(),
							},
						},
						{
							Name:             "doc SVG",
							ShowNameInHeader: false,
							Size:             75,
							Svg: &split.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				}),
			},
		},
	}
}

func (stager *Stager) sidebarSplitlite() []*splitlite.AsSplitArea {
	if !stager.embeddedDiagrams {
		return []*splitlite.AsSplitArea{
			{
				Name:             "doc Tree",
				ShowNameInHeader: false,
				Size:             53,
				Tree: &splitlite.Tree{
					StackName: stager.treeNavigationStage.GetName(),
				},
			},
			{
				Name:             "doc Tree",
				ShowNameInHeader: false,
				IsAny:            true,
				Tree: &splitlite.Tree{
					StackName: stager.treeStage.GetName(),
				},
			},
			{
				Name: "temporary form stack",
				Size: 24,
				Form: &splitlite.Form{
					StackName: stager.formStage.GetName(),
				},
			},
		}
	} else {
		return []*splitlite.AsSplitArea{
			{
				Name:             "doc Tree",
				ShowNameInHeader: false,
				IsAny:            true,
				Tree: &splitlite.Tree{
					StackName: stager.treeStage.GetName(),
				},
			},
		}
	}
}

func (stager *Stager) createViewsSplitlite(receivingAsSplitArea *splitlite.AsSplitArea) {
	receivingAsSplitArea.AsSplit = &splitlite.AsSplit{
		Name:                   "Root As Split for doc receiving area",
		Direction:              splitlite.Horizontal,
		IsWithCustomGutterSize: true,
		GutterSize:             0,
		AsSplitAreas: []*splitlite.AsSplitArea{
			{
				Name:             "AsSplitArea 50% for Slit (Tree & Svg)",
				ShowNameInHeader: false,
				AsSplit: (&splitlite.AsSplit{
					Direction:              splitlite.Horizontal,
					IsWithCustomGutterSize: true,
					GutterSize:             1,
					AsSplitAreas: []*splitlite.AsSplitArea{
						{
							Size: 25,
							AsSplit: &splitlite.AsSplit{
								IsWithCustomGutterSize: true,
								GutterSize:             1,
								IsSizeInPixel:          true,
								Direction:              splitlite.Vertical,
								AsSplitAreas:           stager.sidebarSplitlite(),
							},
						},
						{
							Name:             "doc SVG",
							ShowNameInHeader: false,
							Size:             75,
							Svg: &splitlite.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				}),
			},
		},
	}
}
