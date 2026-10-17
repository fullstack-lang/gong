package probe

import (
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	splitlite "github.com/fullstack-lang/gong/lib/splitlite/go/models"
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

func CreateDataEditorLayoutSplitlite(
	treeNavigationStackName string,
	treeStackName string,
	loadStackName string,
	tableStackName string,
	notificationTableStackName string,
	formStackName string,
) *splitlite.AsSplit {
	return &splitlite.AsSplit{
		Name:          "Top, sidebar, table & form",
		Direction:     splitlite.Horizontal,
		IsSizeInPixel: true,
		AsSplitAreas: []*splitlite.AsSplitArea{
			{
				Name: "sidebar",
				Size: 525,
				AsSplit: &splitlite.AsSplit{
					Direction:              splitlite.Vertical,
					IsSizeInPixel:          true,
					IsWithCustomGutterSize: true,
					GutterSize:             1,
					AsSplitAreas: []*splitlite.AsSplitArea{
						{
							Name: "sidebar tree",
							Size: 53, // to align on the top of the table
							Tree: &splitlite.Tree{
								Name:      "Sidebar",
								StackName: treeNavigationStackName,
							},
						},
						{
							Name:  "sidebar tree",
							IsAny: true,
							Tree: &splitlite.Tree{
								Name:      "Sidebar",
								StackName: treeStackName,
							},
						},
						{
							Name: "load",
							Size: 70,
							Load: &splitlite.Load{
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
				AsSplit: &splitlite.AsSplit{
					Direction: splitlite.Vertical,
					AsSplitAreas: []*splitlite.AsSplitArea{
						{
							Name: "table",
							Size: 50,
							Table: &splitlite.Table{
								Name:      "Table",
								StackName: tableStackName,
							},
						},
						{
							Name: "notification table",
							Size: 50,
							Table: &splitlite.Table{
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
				Form: &splitlite.Form{
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

func CreateFormsLayoutSplitlite(
	formStackName string,
	shapeFormStackName string,
) *splitlite.AsSplit {
	return &splitlite.AsSplit{
		Name:      "as split forms",
		Direction: splitlite.Vertical,
		AsSplitAreas: []*splitlite.AsSplitArea{
			{
				Size: 75,
				Form: &splitlite.Form{
					StackName: formStackName,
				},
			},
			{
				Size: 25,
				Form: &splitlite.Form{
					StackName: shapeFormStackName,
				},
			},
		},
	}
}

