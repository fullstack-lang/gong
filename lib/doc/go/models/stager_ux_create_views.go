package models

import split "github.com/fullstack-lang/gong/lib/split/go/models"

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
				Form2: &split.Form2{
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
