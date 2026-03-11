package models

import split "github.com/fullstack-lang/gong/lib/split/go/models"

func (stager *Stager) createViews(receivingAsSplitArea *split.AsSplitArea) {
	receivingAsSplitArea.AsSplit = &split.AsSplit{
		Name:      "Root As Split for doc receiving area",
		Direction: split.Horizontal,
		AsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "AsSplitArea 50% for Slit (Tree & Svg)",
				ShowNameInHeader: false,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 25,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Name:             "doc Tree",
										ShowNameInHeader: false,
										Size:             66,
										Tree: &split.Tree{
											StackName: stager.treeStage.GetName(),
										},
									},
									{
										Name: "temporary form stack",
										Size: 34,
										Form: &split.Form{
											StackName: stager.formStage.GetName(),
										},
									},
								},
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
