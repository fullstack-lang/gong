// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"strings"

	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
	maticons "github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
)

type tableRowUpdater struct {
	onClick func()
}

func (u *tableRowUpdater) RowUpdated(stage *table_models.Stage, row, updatedRow *table_models.Row) {
	if u.onClick != nil {
		u.onClick()
	}
}

func (probe *StageSetProbe) ux_table() {
	var tableName string
	for tbl := range probe.tableStage.Tables {
		tableName = tbl.Name
	}
	switch tableName {
	case "models.Diagram":
		updateStageSetTable_Diagram_Stage(probe)
	case "models.Library":
		updateStageSetTable_Library_Stage(probe)
	case "models.Note":
		updateStageSetTable_Note_Stage(probe)
	case "models.NoteProductShape":
		updateStageSetTable_NoteProductShape_Stage(probe)
	case "models.NoteResourceShape":
		updateStageSetTable_NoteResourceShape_Stage(probe)
	case "models.NoteShape":
		updateStageSetTable_NoteShape_Stage(probe)
	case "models.NoteTaskShape":
		updateStageSetTable_NoteTaskShape_Stage(probe)
	case "models.Product":
		updateStageSetTable_Product_Stage(probe)
	case "models.ProductCompositionShape":
		updateStageSetTable_ProductCompositionShape_Stage(probe)
	case "models.ProductReferenceShape":
		updateStageSetTable_ProductReferenceShape_Stage(probe)
	case "models.ProductShape":
		updateStageSetTable_ProductShape_Stage(probe)
	case "models.Resource":
		updateStageSetTable_Resource_Stage(probe)
	case "models.ResourceCompositionShape":
		updateStageSetTable_ResourceCompositionShape_Stage(probe)
	case "models.ResourceShape":
		updateStageSetTable_ResourceShape_Stage(probe)
	case "models.ResourceTaskShape":
		updateStageSetTable_ResourceTaskShape_Stage(probe)
	case "models.Task":
		updateStageSetTable_Task_Stage(probe)
	case "models.TaskCompositionShape":
		updateStageSetTable_TaskCompositionShape_Stage(probe)
	case "models.TaskGroup":
		updateStageSetTable_TaskGroup_Stage(probe)
	case "models.TaskGroupShape":
		updateStageSetTable_TaskGroupShape_Stage(probe)
	case "models.TaskInputShape":
		updateStageSetTable_TaskInputShape_Stage(probe)
	case "models.TaskOutputShape":
		updateStageSetTable_TaskOutputShape_Stage(probe)
	case "models.TaskPredecessorShape":
		updateStageSetTable_TaskPredecessorShape_Stage(probe)
	case "models.TaskShape":
		updateStageSetTable_TaskShape_Stage(probe)
	}
}

func updateStageSetTable_Diagram_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Diagram"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DefaultBoxWidth"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DefaultBoxHeigth"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DateFormat"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Width"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Height"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsTimeDiagram"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedStart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedEnd"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedDuration"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DrawVerticalTimeLines"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DrawSecondaryVerticalTimeLines"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HideWeekendsPeriod"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "UseManualStartAndEndDates"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ManualStart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ManualEnd"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeStep"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeStepScale"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SecondaryTimeStep"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SecondaryTimeStepScale"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LaneHeight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RatioBarToLaneHeight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "YTopMargin"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "XLeftText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TextHeight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "XLeftLanes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "XRightMargin"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ArrowLengthToTheRightOfStartBar"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ArrowTipLenght"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeLine_Color"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeLine_FillOpacity"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeLine_Stroke"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeLine_StrokeWidth"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Group_Stroke"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Group_StrokeWidth"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Group_StrokeDashArray"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DateYOffset"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AlignOnBeginningOfTimeScale"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsChecked"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsEditable_"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsShowPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsInAutoLayoutMode"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Product_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ProductsWhoseNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsPBSNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ProductComposition_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ProductReference_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsWBSNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Task_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TasksWhoseNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TasksWhoseInputNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TasksWhoseOutputNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TasksWhosePredecessorNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsTaskGroupsNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TaskGroupShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TaskGroupsWhoseNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TaskComposition_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TaskInputShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TaskOutputShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TaskPredecessorShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Note_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NotesWhoseNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsNotesNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NoteProductShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NoteTaskShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NoteResourceShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Resource_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ResourcesWhoseNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsResourcesNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ResourceComposition_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ResourceTaskShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Library) -> Diagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Diagram]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Diagram_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DefaultBoxWidth"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DefaultBoxWidth)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DefaultBoxHeigth"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DefaultBoxHeigth)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DateFormat"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.DateFormat)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Width"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Width)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Height"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Height)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsTimeDiagram"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsTimeDiagram}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedStart"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedStart)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedEnd"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedEnd)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedDuration"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedDuration)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DrawVerticalTimeLines"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.DrawVerticalTimeLines}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DrawSecondaryVerticalTimeLines"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.DrawSecondaryVerticalTimeLines}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HideWeekendsPeriod"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.HideWeekendsPeriod}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "UseManualStartAndEndDates"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.UseManualStartAndEndDates}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ManualStart"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ManualStart)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ManualEnd"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ManualEnd)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeStep"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.TimeStep)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeStepScale"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.TimeStepScale)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SecondaryTimeStep"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.SecondaryTimeStep)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SecondaryTimeStepScale"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SecondaryTimeStepScale)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LaneHeight"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.LaneHeight)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RatioBarToLaneHeight"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RatioBarToLaneHeight)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "YTopMargin"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.YTopMargin)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "XLeftText"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.XLeftText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TextHeight"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.TextHeight)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "XLeftLanes"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.XLeftLanes)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "XRightMargin"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.XRightMargin)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ArrowLengthToTheRightOfStartBar"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.ArrowLengthToTheRightOfStartBar)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ArrowTipLenght"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.ArrowTipLenght)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeLine_Color"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.TimeLine_Color)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeLine_FillOpacity"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.TimeLine_FillOpacity)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeLine_Stroke"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.TimeLine_Stroke)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeLine_StrokeWidth"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.TimeLine_StrokeWidth)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Group_Stroke"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Group_Stroke)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Group_StrokeWidth"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Group_StrokeWidth)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Group_StrokeDashArray"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Group_StrokeDashArray)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DateYOffset"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DateYOffset)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AlignOnBeginningOfTimeScale"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.AlignOnBeginningOfTimeScale}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsChecked"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsChecked}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsEditable_"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsEditable_}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsShowPrefix"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsShowPrefix}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsInAutoLayoutMode"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsInAutoLayoutMode}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Product_Shapes"}
			var names []string
			for _, elem := range structInstance.Product_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ProductsWhoseNodeIsExpanded"}
			var names []string
			for _, elem := range structInstance.ProductsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsPBSNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsPBSNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ProductComposition_Shapes"}
			var names []string
			for _, elem := range structInstance.ProductComposition_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ProductReference_Shapes"}
			var names []string
			for _, elem := range structInstance.ProductReference_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsWBSNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsWBSNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Task_Shapes"}
			var names []string
			for _, elem := range structInstance.Task_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TasksWhoseNodeIsExpanded"}
			var names []string
			for _, elem := range structInstance.TasksWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TasksWhoseInputNodeIsExpanded"}
			var names []string
			for _, elem := range structInstance.TasksWhoseInputNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TasksWhoseOutputNodeIsExpanded"}
			var names []string
			for _, elem := range structInstance.TasksWhoseOutputNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TasksWhosePredecessorNodeIsExpanded"}
			var names []string
			for _, elem := range structInstance.TasksWhosePredecessorNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsTaskGroupsNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsTaskGroupsNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TaskGroupShapes"}
			var names []string
			for _, elem := range structInstance.TaskGroupShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TaskGroupsWhoseNodeIsExpanded"}
			var names []string
			for _, elem := range structInstance.TaskGroupsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TaskComposition_Shapes"}
			var names []string
			for _, elem := range structInstance.TaskComposition_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TaskInputShapes"}
			var names []string
			for _, elem := range structInstance.TaskInputShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TaskOutputShapes"}
			var names []string
			for _, elem := range structInstance.TaskOutputShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TaskPredecessorShapes"}
			var names []string
			for _, elem := range structInstance.TaskPredecessorShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Note_Shapes"}
			var names []string
			for _, elem := range structInstance.Note_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NotesWhoseNodeIsExpanded"}
			var names []string
			for _, elem := range structInstance.NotesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsNotesNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsNotesNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NoteProductShapes"}
			var names []string
			for _, elem := range structInstance.NoteProductShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NoteTaskShapes"}
			var names []string
			for _, elem := range structInstance.NoteTaskShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NoteResourceShapes"}
			var names []string
			for _, elem := range structInstance.NoteResourceShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Resource_Shapes"}
			var names []string
			for _, elem := range structInstance.Resource_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ResourcesWhoseNodeIsExpanded"}
			var names []string
			for _, elem := range structInstance.ResourcesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsResourcesNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsResourcesNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ResourceComposition_Shapes"}
			var names []string
			for _, elem := range structInstance.ResourceComposition_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ResourceTaskShapes"}
			var names []string
			for _, elem := range structInstance.ResourceTaskShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Library) -> Diagrams"}
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.Diagrams {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Library_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Library"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SubLibraries"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NbPixPerCharacter"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LogoSVGFile"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsRootLibrary"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RootProducts"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RootTasks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RootTaskGroups"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RootResources"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Notes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Diagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Library) -> SubLibraries"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Library]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Library_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SubLibraries"}
			var names []string
			for _, elem := range structInstance.SubLibraries {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NbPixPerCharacter"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.NbPixPerCharacter)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LogoSVGFile"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.LogoSVGFile)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsRootLibrary"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsRootLibrary}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RootProducts"}
			var names []string
			for _, elem := range structInstance.RootProducts {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RootTasks"}
			var names []string
			for _, elem := range structInstance.RootTasks {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RootTaskGroups"}
			var names []string
			for _, elem := range structInstance.RootTaskGroups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RootResources"}
			var names []string
			for _, elem := range structInstance.RootResources {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Notes"}
			var names []string
			for _, elem := range structInstance.Notes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Diagrams"}
			var names []string
			for _, elem := range structInstance.Diagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Library) -> SubLibraries"}
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.SubLibraries {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Note_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Note"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LayoutDirection"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Products"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Tasks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Resources"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> NotesWhoseNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Library) -> Notes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.NoteProductShape) -> Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.NoteResourceShape) -> Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.NoteShape) -> Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.NoteTaskShape) -> Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Note]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Note_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LayoutDirection"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.LayoutDirection)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Products"}
			var names []string
			for _, elem := range structInstance.Products {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Tasks"}
			var names []string
			for _, elem := range structInstance.Tasks {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Resources"}
			var names []string
			for _, elem := range structInstance.Resources {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> NotesWhoseNodeIsExpanded"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.NotesWhoseNodeIsExpanded {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Library) -> Notes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.Notes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.NoteProductShape) -> Note"}
			var refNames []string
			for src := range probe.stageSet.Stage.NoteProductShapes {
				if src.Note == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.NoteResourceShape) -> Note"}
			var refNames []string
			for src := range probe.stageSet.Stage.NoteResourceShapes {
				if src.Note == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.NoteShape) -> Note"}
			var refNames []string
			for src := range probe.stageSet.Stage.NoteShapes {
				if src.Note == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.NoteTaskShape) -> Note"}
			var refNames []string
			for src := range probe.stageSet.Stage.NoteTaskShapes {
				if src.Note == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_NoteProductShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.NoteProductShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CornerOffsetRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> NoteProductShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.NoteProductShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_NoteProductShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Note"}
			val := ""
			if structInstance.Note != nil {
				val = structInstance.Note.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Product"}
			val := ""
			if structInstance.Product != nil {
				val = structInstance.Product.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StartRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EndRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StartOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EndOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CornerOffsetRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.CornerOffsetRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> NoteProductShapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.NoteProductShapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_NoteResourceShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.NoteResourceShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Resource"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CornerOffsetRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> NoteResourceShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.NoteResourceShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_NoteResourceShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Note"}
			val := ""
			if structInstance.Note != nil {
				val = structInstance.Note.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Resource"}
			val := ""
			if structInstance.Resource != nil {
				val = structInstance.Resource.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StartRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EndRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StartOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EndOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CornerOffsetRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.CornerOffsetRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> NoteResourceShapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.NoteResourceShapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_NoteShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.NoteShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsLayoutDirectionDifferent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "X"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Y"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Width"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Height"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> Note_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.NoteShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_NoteShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Note"}
			val := ""
			if structInstance.Note != nil {
				val = structInstance.Note.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsLayoutDirectionDifferent"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsLayoutDirectionDifferent}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "X"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.X)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Y"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Y)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Width"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Width)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Height"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Height)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> Note_Shapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Note_Shapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_NoteTaskShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.NoteTaskShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CornerOffsetRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> NoteTaskShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.NoteTaskShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_NoteTaskShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Note"}
			val := ""
			if structInstance.Note != nil {
				val = structInstance.Note.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Task"}
			val := ""
			if structInstance.Task != nil {
				val = structInstance.Task.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StartRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EndRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StartOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EndOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CornerOffsetRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.CornerOffsetRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> NoteTaskShapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.NoteTaskShapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Product_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Product"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SubProducts"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsProducersNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsConsumersNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsImport"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ReferencedProduct"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LayoutDirection"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> ProductsWhoseNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Library) -> RootProducts"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Note) -> Products"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.NoteProductShape) -> Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Product) -> SubProducts"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Product) -> ReferencedProduct"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.ProductCompositionShape) -> Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.ProductReferenceShape) -> Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.ProductReferenceShape) -> ReferencedProduct"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.ProductShape) -> Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Task) -> Inputs"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Task) -> Outputs"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TaskInputShape) -> Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TaskOutputShape) -> Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Product]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Product_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SubProducts"}
			var names []string
			for _, elem := range structInstance.SubProducts {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsProducersNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsProducersNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsConsumersNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsConsumersNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsImport"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsImport}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ReferencedProduct"}
			val := ""
			if structInstance.ReferencedProduct != nil {
				val = structInstance.ReferencedProduct.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LayoutDirection"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.LayoutDirection)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> ProductsWhoseNodeIsExpanded"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ProductsWhoseNodeIsExpanded {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Library) -> RootProducts"}
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootProducts {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Note) -> Products"}
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				for _, target := range src.Products {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.NoteProductShape) -> Product"}
			var refNames []string
			for src := range probe.stageSet.Stage.NoteProductShapes {
				if src.Product == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Product) -> SubProducts"}
			var refNames []string
			for src := range probe.stageSet.Stage.Products {
				for _, target := range src.SubProducts {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Product) -> ReferencedProduct"}
			var refNames []string
			for src := range probe.stageSet.Stage.Products {
				if src.ReferencedProduct == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.ProductCompositionShape) -> Product"}
			var refNames []string
			for src := range probe.stageSet.Stage.ProductCompositionShapes {
				if src.Product == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.ProductReferenceShape) -> Product"}
			var refNames []string
			for src := range probe.stageSet.Stage.ProductReferenceShapes {
				if src.Product == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.ProductReferenceShape) -> ReferencedProduct"}
			var refNames []string
			for src := range probe.stageSet.Stage.ProductReferenceShapes {
				if src.ReferencedProduct == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.ProductShape) -> Product"}
			var refNames []string
			for src := range probe.stageSet.Stage.ProductShapes {
				if src.Product == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Task) -> Inputs"}
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.Inputs {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Task) -> Outputs"}
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.Outputs {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TaskInputShape) -> Product"}
			var refNames []string
			for src := range probe.stageSet.Stage.TaskInputShapes {
				if src.Product == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TaskOutputShape) -> Product"}
			var refNames []string
			for src := range probe.stageSet.Stage.TaskOutputShapes {
				if src.Product == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_ProductCompositionShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.ProductCompositionShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CornerOffsetRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> ProductComposition_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.ProductCompositionShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_ProductCompositionShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Product"}
			val := ""
			if structInstance.Product != nil {
				val = structInstance.Product.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StartRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EndRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StartOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EndOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CornerOffsetRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.CornerOffsetRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> ProductComposition_Shapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ProductComposition_Shapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_ProductReferenceShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.ProductReferenceShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ReferencedProduct"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CornerOffsetRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> ProductReference_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.ProductReferenceShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_ProductReferenceShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Product"}
			val := ""
			if structInstance.Product != nil {
				val = structInstance.Product.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ReferencedProduct"}
			val := ""
			if structInstance.ReferencedProduct != nil {
				val = structInstance.ReferencedProduct.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StartRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EndRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StartOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EndOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CornerOffsetRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.CornerOffsetRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> ProductReference_Shapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ProductReference_Shapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_ProductShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.ProductShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsShowType"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsLayoutDirectionDifferent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "X"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Y"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Width"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Height"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> Product_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.ProductShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_ProductShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Product"}
			val := ""
			if structInstance.Product != nil {
				val = structInstance.Product.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsShowType"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsShowType}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsLayoutDirectionDifferent"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsLayoutDirectionDifferent}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "X"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.X)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Y"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Y)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Width"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Width)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Height"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Height)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> Product_Shapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Product_Shapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Resource_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Resource"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Tasks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SubResources"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LayoutDirection"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsImport"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ReferencedResource"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> ResourcesWhoseNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Library) -> RootResources"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Note) -> Resources"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.NoteResourceShape) -> Resource"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Resource) -> SubResources"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Resource) -> ReferencedResource"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.ResourceCompositionShape) -> Resource"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.ResourceShape) -> Resource"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.ResourceTaskShape) -> Resource"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Resource]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Resource_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Tasks"}
			var names []string
			for _, elem := range structInstance.Tasks {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SubResources"}
			var names []string
			for _, elem := range structInstance.SubResources {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LayoutDirection"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.LayoutDirection)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsImport"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsImport}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ReferencedResource"}
			val := ""
			if structInstance.ReferencedResource != nil {
				val = structInstance.ReferencedResource.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> ResourcesWhoseNodeIsExpanded"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ResourcesWhoseNodeIsExpanded {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Library) -> RootResources"}
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootResources {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Note) -> Resources"}
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				for _, target := range src.Resources {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.NoteResourceShape) -> Resource"}
			var refNames []string
			for src := range probe.stageSet.Stage.NoteResourceShapes {
				if src.Resource == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Resource) -> SubResources"}
			var refNames []string
			for src := range probe.stageSet.Stage.Resources {
				for _, target := range src.SubResources {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Resource) -> ReferencedResource"}
			var refNames []string
			for src := range probe.stageSet.Stage.Resources {
				if src.ReferencedResource == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.ResourceCompositionShape) -> Resource"}
			var refNames []string
			for src := range probe.stageSet.Stage.ResourceCompositionShapes {
				if src.Resource == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.ResourceShape) -> Resource"}
			var refNames []string
			for src := range probe.stageSet.Stage.ResourceShapes {
				if src.Resource == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.ResourceTaskShape) -> Resource"}
			var refNames []string
			for src := range probe.stageSet.Stage.ResourceTaskShapes {
				if src.Resource == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_ResourceCompositionShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.ResourceCompositionShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Resource"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CornerOffsetRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> ResourceComposition_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.ResourceCompositionShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_ResourceCompositionShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Resource"}
			val := ""
			if structInstance.Resource != nil {
				val = structInstance.Resource.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StartRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EndRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StartOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EndOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CornerOffsetRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.CornerOffsetRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> ResourceComposition_Shapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ResourceComposition_Shapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_ResourceShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.ResourceShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Resource"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsLayoutDirectionDifferent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "X"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Y"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Width"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Height"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> Resource_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.ResourceShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_ResourceShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Resource"}
			val := ""
			if structInstance.Resource != nil {
				val = structInstance.Resource.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsLayoutDirectionDifferent"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsLayoutDirectionDifferent}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "X"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.X)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Y"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Y)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Width"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Width)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Height"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Height)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> Resource_Shapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Resource_Shapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_ResourceTaskShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.ResourceTaskShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Resource"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CornerOffsetRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> ResourceTaskShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.ResourceTaskShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_ResourceTaskShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Resource"}
			val := ""
			if structInstance.Resource != nil {
				val = structInstance.Resource.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Task"}
			val := ""
			if structInstance.Task != nil {
				val = structInstance.Task.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StartRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EndRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StartOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EndOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CornerOffsetRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.CornerOffsetRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> ResourceTaskShapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ResourceTaskShapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Task_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Task"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Start"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "End"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsAllDay"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsMilestone"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TaskGroups"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Predecessors"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DependencyType"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DependencyDurationYears"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DependencyDurationMonths"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DependencyDurationWeeks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DependencyDurationDays"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DependencyDurationHours"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DurationYears"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DurationMonths"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DurationWeeks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DurationDays"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DurationHours"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsEndDateComputedFromDuration"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Inputs"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Outputs"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SubTasks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsWithCompletion"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Completion"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TaskGroupsToDisplay"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TextPosition"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "XOffset"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "YOffset"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsImport"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ReferencedTask"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsInputsNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsOutputsNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LayoutDirection"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> TasksWhoseNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> TasksWhoseInputNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> TasksWhoseOutputNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> TasksWhosePredecessorNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Library) -> RootTasks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Note) -> Tasks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.NoteTaskShape) -> Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Resource) -> Tasks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.ResourceTaskShape) -> Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Task) -> Predecessors"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Task) -> SubTasks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Task) -> ReferencedTask"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TaskCompositionShape) -> Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TaskGroup) -> Tasks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TaskInputShape) -> Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TaskOutputShape) -> Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TaskPredecessorShape) -> Predecessor"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TaskPredecessorShape) -> Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TaskShape) -> Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Task]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Task_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Start"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Start)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "End"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.End)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsAllDay"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsAllDay}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsMilestone"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsMilestone}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TaskGroups"}
			var names []string
			for _, elem := range structInstance.TaskGroups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Predecessors"}
			var names []string
			for _, elem := range structInstance.Predecessors {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DependencyType"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.DependencyType)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DependencyDurationYears"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DependencyDurationYears)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DependencyDurationMonths"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DependencyDurationMonths)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DependencyDurationWeeks"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DependencyDurationWeeks)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DependencyDurationDays"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DependencyDurationDays)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DependencyDurationHours"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DependencyDurationHours)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DurationYears"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DurationYears)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DurationMonths"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DurationMonths)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DurationWeeks"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DurationWeeks)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DurationDays"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DurationDays)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DurationHours"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.DurationHours)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsEndDateComputedFromDuration"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsEndDateComputedFromDuration}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Inputs"}
			var names []string
			for _, elem := range structInstance.Inputs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Outputs"}
			var names []string
			for _, elem := range structInstance.Outputs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SubTasks"}
			var names []string
			for _, elem := range structInstance.SubTasks {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsWithCompletion"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsWithCompletion}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Completion"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Completion)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TaskGroupsToDisplay"}
			var names []string
			for _, elem := range structInstance.TaskGroupsToDisplay {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TextPosition"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.TextPosition)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "XOffset"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.XOffset)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "YOffset"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.YOffset)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsImport"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsImport}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ReferencedTask"}
			val := ""
			if structInstance.ReferencedTask != nil {
				val = structInstance.ReferencedTask.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsInputsNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsInputsNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsOutputsNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsOutputsNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LayoutDirection"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.LayoutDirection)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> TasksWhoseNodeIsExpanded"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TasksWhoseNodeIsExpanded {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> TasksWhoseInputNodeIsExpanded"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TasksWhoseInputNodeIsExpanded {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> TasksWhoseOutputNodeIsExpanded"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TasksWhoseOutputNodeIsExpanded {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> TasksWhosePredecessorNodeIsExpanded"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TasksWhosePredecessorNodeIsExpanded {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Library) -> RootTasks"}
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootTasks {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Note) -> Tasks"}
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				for _, target := range src.Tasks {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.NoteTaskShape) -> Task"}
			var refNames []string
			for src := range probe.stageSet.Stage.NoteTaskShapes {
				if src.Task == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Resource) -> Tasks"}
			var refNames []string
			for src := range probe.stageSet.Stage.Resources {
				for _, target := range src.Tasks {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.ResourceTaskShape) -> Task"}
			var refNames []string
			for src := range probe.stageSet.Stage.ResourceTaskShapes {
				if src.Task == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Task) -> Predecessors"}
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.Predecessors {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Task) -> SubTasks"}
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.SubTasks {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Task) -> ReferencedTask"}
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				if src.ReferencedTask == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TaskCompositionShape) -> Task"}
			var refNames []string
			for src := range probe.stageSet.Stage.TaskCompositionShapes {
				if src.Task == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TaskGroup) -> Tasks"}
			var refNames []string
			for src := range probe.stageSet.Stage.TaskGroups {
				for _, target := range src.Tasks {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TaskInputShape) -> Task"}
			var refNames []string
			for src := range probe.stageSet.Stage.TaskInputShapes {
				if src.Task == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TaskOutputShape) -> Task"}
			var refNames []string
			for src := range probe.stageSet.Stage.TaskOutputShapes {
				if src.Task == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TaskPredecessorShape) -> Predecessor"}
			var refNames []string
			for src := range probe.stageSet.Stage.TaskPredecessorShapes {
				if src.Predecessor == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TaskPredecessorShape) -> Task"}
			var refNames []string
			for src := range probe.stageSet.Stage.TaskPredecessorShapes {
				if src.Task == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TaskShape) -> Task"}
			var refNames []string
			for src := range probe.stageSet.Stage.TaskShapes {
				if src.Task == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TaskCompositionShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.TaskCompositionShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CornerOffsetRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> TaskComposition_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.TaskCompositionShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_TaskCompositionShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Task"}
			val := ""
			if structInstance.Task != nil {
				val = structInstance.Task.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StartRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EndRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StartOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EndOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CornerOffsetRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.CornerOffsetRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> TaskComposition_Shapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskComposition_Shapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TaskGroup_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.TaskGroup"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Tasks"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> TaskGroupsWhoseNodeIsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Library) -> RootTaskGroups"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Task) -> TaskGroups"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Task) -> TaskGroupsToDisplay"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TaskGroupShape) -> TaskGroup"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.TaskGroup]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_TaskGroup_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Tasks"}
			var names []string
			for _, elem := range structInstance.Tasks {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> TaskGroupsWhoseNodeIsExpanded"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskGroupsWhoseNodeIsExpanded {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Library) -> RootTaskGroups"}
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootTaskGroups {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Task) -> TaskGroups"}
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.TaskGroups {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Task) -> TaskGroupsToDisplay"}
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.TaskGroupsToDisplay {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TaskGroupShape) -> TaskGroup"}
			var refNames []string
			for src := range probe.stageSet.Stage.TaskGroupShapes {
				if src.TaskGroup == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TaskGroupShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.TaskGroupShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TaskGroup"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "X"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Y"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Width"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Height"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> TaskGroupShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.TaskGroupShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_TaskGroupShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TaskGroup"}
			val := ""
			if structInstance.TaskGroup != nil {
				val = structInstance.TaskGroup.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "X"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.X)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Y"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Y)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Width"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Width)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Height"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Height)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> TaskGroupShapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskGroupShapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TaskInputShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.TaskInputShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CornerOffsetRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> TaskInputShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.TaskInputShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_TaskInputShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Product"}
			val := ""
			if structInstance.Product != nil {
				val = structInstance.Product.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Task"}
			val := ""
			if structInstance.Task != nil {
				val = structInstance.Task.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StartRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EndRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StartOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EndOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CornerOffsetRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.CornerOffsetRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> TaskInputShapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskInputShapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TaskOutputShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.TaskOutputShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Product"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CornerOffsetRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> TaskOutputShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.TaskOutputShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_TaskOutputShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Task"}
			val := ""
			if structInstance.Task != nil {
				val = structInstance.Task.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Product"}
			val := ""
			if structInstance.Product != nil {
				val = structInstance.Product.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StartRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EndRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StartOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EndOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CornerOffsetRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.CornerOffsetRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> TaskOutputShapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskOutputShapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TaskPredecessorShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.TaskPredecessorShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Predecessor"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndOrientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CornerOffsetRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> TaskPredecessorShapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.TaskPredecessorShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_TaskPredecessorShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Predecessor"}
			val := ""
			if structInstance.Predecessor != nil {
				val = structInstance.Predecessor.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Task"}
			val := ""
			if structInstance.Task != nil {
				val = structInstance.Task.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StartRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EndRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StartOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndOrientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EndOrientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CornerOffsetRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.CornerOffsetRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> TaskPredecessorShapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskPredecessorShapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TaskShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.TaskShape"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Task"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsShowDate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalOffset"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DisplayVerticalBar"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsLayoutDirectionDifferent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "X"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Y"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Width"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Height"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHidden"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Diagram) -> Task_Shapes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.TaskShape]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_TaskShape_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Task"}
			val := ""
			if structInstance.Task != nil {
				val = structInstance.Task.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsShowDate"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsShowDate}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalOffset"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.VerticalOffset)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DisplayVerticalBar"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.DisplayVerticalBar}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsLayoutDirectionDifferent"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsLayoutDirectionDifferent}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "X"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.X)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Y"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Y)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Width"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Width)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Height"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Height)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHidden"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHidden}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Diagram) -> Task_Shapes"}
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Task_Shapes {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

