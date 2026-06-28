package models

import (
	table "github.com/fullstack-lang/gong/lib/table/go/models"
)

func (stager *Stager) updateAndCommitSummaryTableStage() {

	stager.summaryTableStage.Reset()

	summary := &table.Table{
		Name: "Summary",
		DisplayedColumns: []*table.DisplayedColumn{
			{
				Name: "Property",
			},
			{
				Name: "Value",
			},
		},
		Rows: []*table.Row{
			{
				Cells: []*table.Cell{
					{
						CellString: &table.CellString{
							Value: "Filename",
						},
					},
					{
						CellString: &table.CellString{
							Value: stager.PathToReqifFile,
						},
					},
				},
			},
			{
				Cells: []*table.Cell{
					{
						CellString: &table.CellString{
							Value: "Creation time",
						},
					},
					{
						CellString: &table.CellString{
							Value: stager.rootReqif.THE_HEADER.REQ_IF_HEADER.CREATION_TIME,
						},
					},
				},
			},
			{
				Cells: []*table.Cell{
					{
						CellString: &table.CellString{
							Value: "Title",
						},
					},
					{
						CellString: &table.CellString{
							Value: stager.rootReqif.THE_HEADER.REQ_IF_HEADER.TITLE,
						},
					},
				},
			},
			{
				Cells: []*table.Cell{
					{
						CellString: &table.CellString{
							Value: "REQ-IF-TOOL-ID",
						},
					},
					{
						CellString: &table.CellString{
							Value: stager.rootReqif.THE_HEADER.REQ_IF_HEADER.REQ_IF_TOOL_ID,
						},
					},
				},
			},
			{
				Cells: []*table.Cell{
					{
						CellString: &table.CellString{
							Value: "REQ-IF-VERSION",
						},
					},
					{
						CellString: &table.CellString{
							Value: stager.rootReqif.THE_HEADER.REQ_IF_HEADER.REQ_IF_VERSION,
						},
					},
				},
			},
			{
				Cells: []*table.Cell{
					{
						CellString: &table.CellString{
							Value: "SOURCE-TOOL-ID",
						},
					},
					{
						CellString: &table.CellString{
							Value: stager.rootReqif.THE_HEADER.REQ_IF_HEADER.SOURCE_TOOL_ID,
						},
					},
				},
			},
			{
				Cells: []*table.Cell{
					{
						CellString: &table.CellString{
							Value: "COMMENT",
						},
					},
					{
						CellString: &table.CellString{
							Value: stager.rootReqif.THE_HEADER.REQ_IF_HEADER.COMMENT,
						},
					},
				},
			},
			{
				Cells: []*table.Cell{
					{
						CellString: &table.CellString{
							Value: "IDENTIFIER",
						},
					},
					{
						CellString: &table.CellString{
							Value: stager.rootReqif.THE_HEADER.REQ_IF_HEADER.IDENTIFIER,
						},
					},
				},
			},
		},
	}

	table.StageBranch(stager.summaryTableStage, summary)

	stager.summaryTableStage.Commit()
}
