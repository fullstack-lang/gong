package probe

import (
	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

func CreateDataEditorLayout(
	treeNavigationStackName string,
	treeStackName string,
	loadStackName string,
	tableStackName string,
	notificationTableStackName string,
	formStackName string,
) *split.AsSplit {
	return &split.AsSplit{
		Name:          "Top, sidebar, table & form",
		Direction:     split.Horizontal,
		IsSizeInPixel: true,
		AsSplitAreas: []*split.AsSplitArea{
			{
				Name: "sidebar",
				Size: 525,
				AsSplit: &split.AsSplit{
					Direction:              split.Vertical,
					IsSizeInPixel:          true,
					IsWithCustomGutterSize: true,
					GutterSize:             1,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name: "sidebar tree",
							Size: 53, // to align on the top of the table
							Tree: &split.Tree{
								Name:      "Sidebar",
								StackName: treeNavigationStackName,
							},
						},
						{
							Name:  "sidebar tree",
							IsAny: true,
							Tree: &split.Tree{
								Name:      "Sidebar",
								StackName: treeStackName,
							},
						},
						{
							Name: "load",
							Size: 70,
							Load: &split.Load{
								Name:      "Table",
								StackName: loadStackName,
							},
						},
					},
				},
			},

			{
				Name:  "both tables",
				IsAny: true,
				AsSplit: &split.AsSplit{
					Direction: split.Vertical,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name: "table",
							Size: 50,
							Table: &split.Table{
								Name:      "Table",
								StackName: tableStackName,
							},
						},
						{
							Name: "notification table",
							Size: 50,
							Table: &split.Table{
								Name:      "Table",
								StackName: notificationTableStackName,
							},
						},
					},
				},
			},
			{
				Name: "form",
				Size: 525,
				Form: &split.Form{
					Name:      "Form",
					StackName: formStackName,
				},
			},
		},
	}
}

func CreateFormsLayout(
	formStackName string,
	shapeFormStackName string,
) *split.AsSplit {
	return &split.AsSplit{
		Name:      "as split forms",
		Direction: split.Vertical,
		AsSplitAreas: []*split.AsSplitArea{
			{
				Size: 75,
				Form: &split.Form{
					StackName: formStackName,
				},
			},
			{
				Size: 25,
				Form: &split.Form{
					StackName: shapeFormStackName,
				},
			},
		},
	}
}

