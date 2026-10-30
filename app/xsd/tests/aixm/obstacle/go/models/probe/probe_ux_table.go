// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	maticons "github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/xlink"
)

const TableName = "Table"

var (
	_ = fmt.Sprintf
	_ = sort.Strings
	_ = strings.Join
	_ = maticons.BUTTON_delete
)

type tableRowUpdater struct {
	onClick func()
}

func (u *tableRowUpdater) RowUpdated(stage *table_models.Stage, row, updatedRow *table_models.Row) {
	if u.onClick != nil {
		u.onClick()
	}
}

func (probe *Probe) ux_table() {
	var tableName string
	for tbl := range probe.tableStage.Tables {
		tableName = tbl.Name
	}
	switch tableName {
	case "models.AIXMBasicMessage", "AIXMBasicMessage":
		updateStageSetTable_AIXMBasicMessage_Stage(probe)
	case "models.AIXMBasicMessageType", "AIXMBasicMessageType":
		updateStageSetTable_AIXMBasicMessageType_Stage(probe)
	case "models.BasicMessageMemberAIXMPropertyType", "BasicMessageMemberAIXMPropertyType":
		updateStageSetTable_BasicMessageMemberAIXMPropertyType_Stage(probe)
	case "aixmfeatures.A_extension":
		updateStageSetTable_A_extension_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_1":
		updateStageSetTable_A_extension_1_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_10":
		updateStageSetTable_A_extension_10_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_11":
		updateStageSetTable_A_extension_11_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_12":
		updateStageSetTable_A_extension_12_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_2":
		updateStageSetTable_A_extension_2_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_3":
		updateStageSetTable_A_extension_3_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_4":
		updateStageSetTable_A_extension_4_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_5":
		updateStageSetTable_A_extension_5_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_6":
		updateStageSetTable_A_extension_6_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_7":
		updateStageSetTable_A_extension_7_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_8":
		updateStageSetTable_A_extension_8_AixmfeaturesStage(probe)
	case "aixmfeatures.A_extension_9":
		updateStageSetTable_A_extension_9_AixmfeaturesStage(probe)
	case "aixmfeatures.AbstractPropertiesWithScheduleType":
		updateStageSetTable_AbstractPropertiesWithScheduleType_AixmfeaturesStage(probe)
	case "aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType":
		updateStageSetTable_AeronauticalGroundLightTimeSlicePropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.AeronauticalGroundLightTimeSliceType":
		updateStageSetTable_AeronauticalGroundLightTimeSliceType_AixmfeaturesStage(probe)
	case "aixmfeatures.AeronauticalGroundLightType":
		updateStageSetTable_AeronauticalGroundLightType_AixmfeaturesStage(probe)
	case "aixmfeatures.AirportHeliportPropertyType":
		updateStageSetTable_AirportHeliportPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.CurveType":
		updateStageSetTable_CurveType_AixmfeaturesStage(probe)
	case "aixmfeatures.ElevatedCurvePropertyType":
		updateStageSetTable_ElevatedCurvePropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.ElevatedCurveType":
		updateStageSetTable_ElevatedCurveType_AixmfeaturesStage(probe)
	case "aixmfeatures.ElevatedPointPropertyType":
		updateStageSetTable_ElevatedPointPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.ElevatedPointType":
		updateStageSetTable_ElevatedPointType_AixmfeaturesStage(probe)
	case "aixmfeatures.ElevatedSurfacePropertyType":
		updateStageSetTable_ElevatedSurfacePropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.ElevatedSurfaceType":
		updateStageSetTable_ElevatedSurfaceType_AixmfeaturesStage(probe)
	case "aixmfeatures.GroundLightSystemPropertyType":
		updateStageSetTable_GroundLightSystemPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.LightElementPropertyType":
		updateStageSetTable_LightElementPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.LightElementStatusPropertyType":
		updateStageSetTable_LightElementStatusPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.LightElementStatusType":
		updateStageSetTable_LightElementStatusType_AixmfeaturesStage(probe)
	case "aixmfeatures.LightElementType":
		updateStageSetTable_LightElementType_AixmfeaturesStage(probe)
	case "aixmfeatures.LinguisticNotePropertyType":
		updateStageSetTable_LinguisticNotePropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.LinguisticNoteType":
		updateStageSetTable_LinguisticNoteType_AixmfeaturesStage(probe)
	case "aixmfeatures.MarkerBeaconPropertyType":
		updateStageSetTable_MarkerBeaconPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.NavaidEquipmentPropertyType":
		updateStageSetTable_NavaidEquipmentPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.NotePropertyType":
		updateStageSetTable_NotePropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.NoteType":
		updateStageSetTable_NoteType_AixmfeaturesStage(probe)
	case "aixmfeatures.ObstacleAreaTimeSlicePropertyType":
		updateStageSetTable_ObstacleAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.ObstacleAreaTimeSliceType":
		updateStageSetTable_ObstacleAreaTimeSliceType_AixmfeaturesStage(probe)
	case "aixmfeatures.ObstacleAreaType":
		updateStageSetTable_ObstacleAreaType_AixmfeaturesStage(probe)
	case "aixmfeatures.OrganisationAuthorityPropertyType":
		updateStageSetTable_OrganisationAuthorityPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.PassengerServicePropertyType":
		updateStageSetTable_PassengerServicePropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.PointType":
		updateStageSetTable_PointType_AixmfeaturesStage(probe)
	case "aixmfeatures.RunwayDirectionPropertyType":
		updateStageSetTable_RunwayDirectionPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.ServicePropertyType":
		updateStageSetTable_ServicePropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.SpecialNavigationStationPropertyType":
		updateStageSetTable_SpecialNavigationStationPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.SurfacePropertyType":
		updateStageSetTable_SurfacePropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.SurfaceType":
		updateStageSetTable_SurfaceType_AixmfeaturesStage(probe)
	case "aixmfeatures.TimesheetPropertyType":
		updateStageSetTable_TimesheetPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.TimesheetType":
		updateStageSetTable_TimesheetType_AixmfeaturesStage(probe)
	case "aixmfeatures.UnitPropertyType":
		updateStageSetTable_UnitPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.VerticalStructureLightingStatusPropertyType":
		updateStageSetTable_VerticalStructureLightingStatusPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.VerticalStructureLightingStatusType":
		updateStageSetTable_VerticalStructureLightingStatusType_AixmfeaturesStage(probe)
	case "aixmfeatures.VerticalStructurePartPropertyType":
		updateStageSetTable_VerticalStructurePartPropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.VerticalStructurePartType":
		updateStageSetTable_VerticalStructurePartType_AixmfeaturesStage(probe)
	case "aixmfeatures.VerticalStructurePropertyType":
		updateStageSetTable_VerticalStructurePropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.VerticalStructureTimeSlicePropertyType":
		updateStageSetTable_VerticalStructureTimeSlicePropertyType_AixmfeaturesStage(probe)
	case "aixmfeatures.VerticalStructureTimeSliceType":
		updateStageSetTable_VerticalStructureTimeSliceType_AixmfeaturesStage(probe)
	case "aixmfeatures.VerticalStructureType":
		updateStageSetTable_VerticalStructureType_AixmfeaturesStage(probe)
	case "aixmdatatypes.CodeColourType":
		updateStageSetTable_CodeColourType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeDayType":
		updateStageSetTable_CodeDayType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeGroundLightingType":
		updateStageSetTable_CodeGroundLightingType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeLightIntensityType":
		updateStageSetTable_CodeLightIntensityType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeLightSourceType":
		updateStageSetTable_CodeLightSourceType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeNotePurposeType":
		updateStageSetTable_CodeNotePurposeType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeObstacleAreaType":
		updateStageSetTable_CodeObstacleAreaType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeObstacleAssessmentSurfaceType":
		updateStageSetTable_CodeObstacleAssessmentSurfaceType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeStatusConstructionType":
		updateStageSetTable_CodeStatusConstructionType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeStatusOperationsType":
		updateStageSetTable_CodeStatusOperationsType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeTimeEventCombinationType":
		updateStageSetTable_CodeTimeEventCombinationType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeTimeEventType":
		updateStageSetTable_CodeTimeEventType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeTimeReferenceType":
		updateStageSetTable_CodeTimeReferenceType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeVerticalDatumType":
		updateStageSetTable_CodeVerticalDatumType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeVerticalStructureMarkingType":
		updateStageSetTable_CodeVerticalStructureMarkingType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeVerticalStructureMaterialType":
		updateStageSetTable_CodeVerticalStructureMaterialType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeVerticalStructureType":
		updateStageSetTable_CodeVerticalStructureType_AixmdatatypesStage(probe)
	case "aixmdatatypes.CodeYesNoType":
		updateStageSetTable_CodeYesNoType_AixmdatatypesStage(probe)
	case "aixmdatatypes.DateMonthDayType":
		updateStageSetTable_DateMonthDayType_AixmdatatypesStage(probe)
	case "aixmdatatypes.TextDesignatorType":
		updateStageSetTable_TextDesignatorType_AixmdatatypesStage(probe)
	case "aixmdatatypes.TextNameType":
		updateStageSetTable_TextNameType_AixmdatatypesStage(probe)
	case "aixmdatatypes.TextNoteType":
		updateStageSetTable_TextNoteType_AixmdatatypesStage(probe)
	case "aixmdatatypes.TextPropertyNameType":
		updateStageSetTable_TextPropertyNameType_AixmdatatypesStage(probe)
	case "aixmdatatypes.TimeType":
		updateStageSetTable_TimeType_AixmdatatypesStage(probe)
	case "aixmdatatypes.ValDistanceSignedType":
		updateStageSetTable_ValDistanceSignedType_AixmdatatypesStage(probe)
	case "aixmdatatypes.ValDistanceType":
		updateStageSetTable_ValDistanceType_AixmdatatypesStage(probe)
	case "aixmdatatypes.ValDistanceVerticalType":
		updateStageSetTable_ValDistanceVerticalType_AixmdatatypesStage(probe)
	case "aixmdatatypes.ValDurationType":
		updateStageSetTable_ValDurationType_AixmdatatypesStage(probe)
	case "aixmdatatypes.ValLightIntensityType":
		updateStageSetTable_ValLightIntensityType_AixmdatatypesStage(probe)
	case "aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType":
		updateStageSetTable_AbstractAIXMFeatureBaseType_AixmabstractgmlobjecttypesStage(probe)
	case "aixmabstractgmlobjecttypes.AbstractAIXMFeatureType":
		updateStageSetTable_AbstractAIXMFeatureType_AixmabstractgmlobjecttypesStage(probe)
	case "aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType":
		updateStageSetTable_AbstractAIXMMessageBaseType_AixmabstractgmlobjecttypesStage(probe)
	case "aixmabstractgmlobjecttypes.AbstractAIXMMessageType":
		updateStageSetTable_AbstractAIXMMessageType_AixmabstractgmlobjecttypesStage(probe)
	case "aixmabstractgmlobjecttypes.AbstractAIXMObjectType":
		updateStageSetTable_AbstractAIXMObjectType_AixmabstractgmlobjecttypesStage(probe)
	case "aixmabstractgmlobjecttypes.AbstractAIXMPropertyType":
		updateStageSetTable_AbstractAIXMPropertyType_AixmabstractgmlobjecttypesStage(probe)
	case "aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType":
		updateStageSetTable_AbstractAIXMTimeSliceBaseType_AixmabstractgmlobjecttypesStage(probe)
	case "aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType":
		updateStageSetTable_AbstractAIXMTimeSliceType_AixmabstractgmlobjecttypesStage(probe)
	case "aixmabstractgmlobjecttypes.AbstractExtensionType":
		updateStageSetTable_AbstractExtensionType_AixmabstractgmlobjecttypesStage(probe)
	case "aixmabstractgmlobjecttypes.FeatureMetadataPropertyType":
		updateStageSetTable_FeatureMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
	case "aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType":
		updateStageSetTable_FeatureTimeSliceMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
	case "aixmabstractgmlobjecttypes.MessageMetadataPropertyType":
		updateStageSetTable_MessageMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
	case "gml.AbstractCurveSegmentType":
		updateStageSetTable_AbstractCurveSegmentType_GmlStage(probe)
	case "gml.AbstractCurveType":
		updateStageSetTable_AbstractCurveType_GmlStage(probe)
	case "gml.AbstractFeatureMemberType":
		updateStageSetTable_AbstractFeatureMemberType_GmlStage(probe)
	case "gml.AbstractFeatureType":
		updateStageSetTable_AbstractFeatureType_GmlStage(probe)
	case "gml.AbstractGMLType":
		updateStageSetTable_AbstractGMLType_GmlStage(probe)
	case "gml.AbstractGeometricPrimitiveType":
		updateStageSetTable_AbstractGeometricPrimitiveType_GmlStage(probe)
	case "gml.AbstractGeometryType":
		updateStageSetTable_AbstractGeometryType_GmlStage(probe)
	case "gml.AbstractMetadataPropertyType":
		updateStageSetTable_AbstractMetadataPropertyType_GmlStage(probe)
	case "gml.AbstractRingPropertyType":
		updateStageSetTable_AbstractRingPropertyType_GmlStage(probe)
	case "gml.AbstractRingType":
		updateStageSetTable_AbstractRingType_GmlStage(probe)
	case "gml.AbstractSurfacePatchType":
		updateStageSetTable_AbstractSurfacePatchType_GmlStage(probe)
	case "gml.AbstractSurfaceType":
		updateStageSetTable_AbstractSurfaceType_GmlStage(probe)
	case "gml.AbstractTimeGeometricPrimitiveType":
		updateStageSetTable_AbstractTimeGeometricPrimitiveType_GmlStage(probe)
	case "gml.AbstractTimeObjectType":
		updateStageSetTable_AbstractTimeObjectType_GmlStage(probe)
	case "gml.AbstractTimePrimitiveType":
		updateStageSetTable_AbstractTimePrimitiveType_GmlStage(probe)
	case "gml.AbstractTimeSliceType":
		updateStageSetTable_AbstractTimeSliceType_GmlStage(probe)
	case "gml.AngleType":
		updateStageSetTable_AngleType_GmlStage(probe)
	case "gml.ArcByCenterPointType":
		updateStageSetTable_ArcByCenterPointType_GmlStage(probe)
	case "gml.ArcStringType":
		updateStageSetTable_ArcStringType_GmlStage(probe)
	case "gml.ArcType":
		updateStageSetTable_ArcType_GmlStage(probe)
	case "gml.BoundingShapeType":
		updateStageSetTable_BoundingShapeType_GmlStage(probe)
	case "gml.CircleByCenterPointType":
		updateStageSetTable_CircleByCenterPointType_GmlStage(probe)
	case "gml.CircleType":
		updateStageSetTable_CircleType_GmlStage(probe)
	case "gml.CodeType":
		updateStageSetTable_CodeType_GmlStage(probe)
	case "gml.CodeWithAuthorityType":
		updateStageSetTable_CodeWithAuthorityType_GmlStage(probe)
	case "gml.CompositeCurveType":
		updateStageSetTable_CompositeCurveType_GmlStage(probe)
	case "gml.CurvePropertyType":
		updateStageSetTable_CurvePropertyType_GmlStage(probe)
	case "gml.CurveSegmentArrayPropertyType":
		updateStageSetTable_CurveSegmentArrayPropertyType_GmlStage(probe)
	case "gml.CurveType":
		updateStageSetTable_CurveType_GmlStage(probe)
	case "gml.DirectPositionListType":
		updateStageSetTable_DirectPositionListType_GmlStage(probe)
	case "gml.DirectPositionType":
		updateStageSetTable_DirectPositionType_GmlStage(probe)
	case "gml.DynamicFeatureType":
		updateStageSetTable_DynamicFeatureType_GmlStage(probe)
	case "gml.EnvelopeType":
		updateStageSetTable_EnvelopeType_GmlStage(probe)
	case "gml.GeodesicStringType":
		updateStageSetTable_GeodesicStringType_GmlStage(probe)
	case "gml.GeodesicType":
		updateStageSetTable_GeodesicType_GmlStage(probe)
	case "gml.LengthType":
		updateStageSetTable_LengthType_GmlStage(probe)
	case "gml.LineStringSegmentType":
		updateStageSetTable_LineStringSegmentType_GmlStage(probe)
	case "gml.MeasureType":
		updateStageSetTable_MeasureType_GmlStage(probe)
	case "gml.OrientableCurveType":
		updateStageSetTable_OrientableCurveType_GmlStage(probe)
	case "gml.PointPropertyType":
		updateStageSetTable_PointPropertyType_GmlStage(probe)
	case "gml.PointType":
		updateStageSetTable_PointType_GmlStage(probe)
	case "gml.PolygonPatchType":
		updateStageSetTable_PolygonPatchType_GmlStage(probe)
	case "gml.RingType":
		updateStageSetTable_RingType_GmlStage(probe)
	case "gml.SurfacePatchArrayPropertyType":
		updateStageSetTable_SurfacePatchArrayPropertyType_GmlStage(probe)
	case "gml.SurfaceType":
		updateStageSetTable_SurfaceType_GmlStage(probe)
	case "gml.TimeInstantType":
		updateStageSetTable_TimeInstantType_GmlStage(probe)
	case "gml.TimePeriodType":
		updateStageSetTable_TimePeriodType_GmlStage(probe)
	case "gml.TimePositionType":
		updateStageSetTable_TimePositionType_GmlStage(probe)
	case "gml.TimePrimitivePropertyType":
		updateStageSetTable_TimePrimitivePropertyType_GmlStage(probe)
	case "xlink.Xlink_SchemaElement":
		updateStageSetTable_Xlink_SchemaElement_XlinkStage(probe)
	}
}

func (probe *Probe) UpdateAndCommitNotificationTable() {
	probe.notificationTableStage.Reset()

	table := new(table_models.Table)
	table.Name = TableName
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	notificationsResetButton := &table_models.Button{
		Name:            "NotificationsResetButton",
		Icon:            string(tree_buttons.BUTTON_playlist_remove),
		HasToolTip:      true,
		ToolTipText:     "Reset notification table",
		ToolTipPosition: table_models.Below,
		OnClick: func() {
			probe.ResetNotifications()
		},
	}
	table.Buttons = append(table.Buttons, notificationsResetButton)

	notificationsDownloadCSVButton := &table_models.Button{
		Name:            "NotificationsDownloadCSVButton",
		Icon:            string(tree_buttons.BUTTON_download),
		HasToolTip:      true,
		ToolTipText:     "Download notifications as CSV",
		ToolTipPosition: table_models.Below,
		OnClick: func() {
			probe.DownloadNotificationsCSV()
		},
	}
	table.Buttons = append(table.Buttons, notificationsDownloadCSVButton)

	column := new(table_models.DisplayedColumn)
	column.Name = "Date"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(table_models.DisplayedColumn)
	column.Name = "Message"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	sort.Slice(probe.notification, func(i, j int) bool {
		return probe.notification[i].Date.After(probe.notification[j].Date)
	})

	for _, notification := range probe.notification {
		row := new(table_models.Row)
		table.Rows = append(table.Rows, row)

		{
			cell := new(table_models.Cell)
			cellString := new(table_models.CellString)
			cell.CellString = cellString
			cellString.Value = notification.Date.Format(time.StampMicro)
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := new(table_models.Cell)
			cellString := new(table_models.CellString)
			cell.CellString = cellString
			cellString.Value = notification.Message
			row.Cells = append(row.Cells, cell)
		}
	}

	probe.notificationTableStage.StageBranch(table)
	probe.notificationTableStage.Commit()
}

func updateStageSetTable_AIXMBasicMessage_Stage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.AIXMBasicMessage"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SequenceNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "MessageMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AggregationType"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HasMember"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GetInstancesByOrder[*models.AIXMBasicMessage]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AIXMBasicMessage_Stage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BoundedBy"}
			val := ""
			if structInstance.BoundedBy != nil {
				val = structInstance.BoundedBy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SequenceNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.SequenceNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "MessageMetadata"}
			val := ""
			if structInstance.MessageMetadata != nil {
				val = structInstance.MessageMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AggregationType"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.AggregationType)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HasMember"}
			var names []string
			for _, elem := range structInstance.HasMember {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AIXMBasicMessageType_Stage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.AIXMBasicMessageType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SequenceNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "MessageMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AggregationType"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HasMember"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GetInstancesByOrder[*models.AIXMBasicMessageType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AIXMBasicMessageType_Stage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BoundedBy"}
			val := ""
			if structInstance.BoundedBy != nil {
				val = structInstance.BoundedBy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SequenceNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.SequenceNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "MessageMetadata"}
			val := ""
			if structInstance.MessageMetadata != nil {
				val = structInstance.MessageMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AggregationType"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.AggregationType)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HasMember"}
			var names []string
			for _, elem := range structInstance.HasMember {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_BasicMessageMemberAIXMPropertyType_Stage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.BasicMessageMemberAIXMPropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalStructure"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ObstacleArea"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AeronauticalGroundLight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AIXMBasicMessage) -> HasMember"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AIXMBasicMessageType) -> HasMember"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GetInstancesByOrder[*models.BasicMessageMemberAIXMPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_BasicMessageMemberAIXMPropertyType_Stage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalStructure"}
			val := ""
			if structInstance.VerticalStructure != nil {
				val = structInstance.VerticalStructure.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ObstacleArea"}
			val := ""
			if structInstance.ObstacleArea != nil {
				val = structInstance.ObstacleArea.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AeronauticalGroundLight"}
			val := ""
			if structInstance.AeronauticalGroundLight != nil {
				val = structInstance.AeronauticalGroundLight.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AIXMBasicMessage) -> HasMember"}
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessages {
				for _, target := range src.HasMember {
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
			cell := &table_models.Cell{Name: "(models.AIXMBasicMessageType) -> HasMember"}
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessageTypes {
				for _, target := range src.HasMember {
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

func updateStageSetTable_A_extension_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractElevatedCurveExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedCurveType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractElevatedCurveExtension"}
			var names []string
			for _, elem := range structInstance.AbstractElevatedCurveExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedCurveType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_1_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_1"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractElevatedPointExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedPointType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_1]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_1_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractElevatedPointExtension"}
			var names []string
			for _, elem := range structInstance.AbstractElevatedPointExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedPointType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_10_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_10"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractLightElementExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_10]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_10_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractLightElementExtension"}
			var names []string
			for _, elem := range structInstance.AbstractLightElementExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_11_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_11"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractPropertiesWithScheduleExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractLightElementStatusExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementStatusType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_11]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_11_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractPropertiesWithScheduleExtension"}
			var names []string
			for _, elem := range structInstance.AbstractPropertiesWithScheduleExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractLightElementStatusExtension"}
			var names []string
			for _, elem := range structInstance.AbstractLightElementStatusExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementStatusType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_12_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_12"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractTimesheetExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_12]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_12_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractTimesheetExtension"}
			var names []string
			for _, elem := range structInstance.AbstractTimesheetExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_2_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_2"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractElevatedSurfaceExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedSurfaceType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_2]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_2_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractElevatedSurfaceExtension"}
			var names []string
			for _, elem := range structInstance.AbstractElevatedSurfaceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedSurfaceType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_3_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_3"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractAeronauticalGroundLightExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_3]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_3_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractAeronauticalGroundLightExtension"}
			var names []string
			for _, elem := range structInstance.AbstractAeronauticalGroundLightExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_4_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_4"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractLinguisticNoteExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LinguisticNoteType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_4]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_4_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractLinguisticNoteExtension"}
			var names []string
			for _, elem := range structInstance.AbstractLinguisticNoteExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LinguisticNoteType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LinguisticNoteTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_5_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_5"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractNoteExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.NoteType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_5]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_5_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractNoteExtension"}
			var names []string
			for _, elem := range structInstance.AbstractNoteExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.NoteType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_6_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_6"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractObstacleAreaExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_6]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_6_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractObstacleAreaExtension"}
			var names []string
			for _, elem := range structInstance.AbstractObstacleAreaExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_7_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_7"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractVerticalStructureExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_7]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_7_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractVerticalStructureExtension"}
			var names []string
			for _, elem := range structInstance.AbstractVerticalStructureExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_8_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_8"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractPropertiesWithScheduleExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractVerticalStructureLightingStatusExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureLightingStatusType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_8]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_8_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractPropertiesWithScheduleExtension"}
			var names []string
			for _, elem := range structInstance.AbstractPropertiesWithScheduleExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractVerticalStructureLightingStatusExtension"}
			var names []string
			for _, elem := range structInstance.AbstractVerticalStructureLightingStatusExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureLightingStatusType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_A_extension_9_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.A_extension_9"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractPropertiesWithScheduleExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AbstractVerticalStructurePartExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_9]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_A_extension_9_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractPropertiesWithScheduleExtension"}
			var names []string
			for _, elem := range structInstance.AbstractPropertiesWithScheduleExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AbstractVerticalStructurePartExtension"}
			var names []string
			for _, elem := range structInstance.AbstractVerticalStructurePartExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> Extension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.Extension {
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

func updateStageSetTable_AbstractPropertiesWithScheduleType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.AbstractPropertiesWithScheduleType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AbstractPropertiesWithScheduleType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractPropertiesWithScheduleType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AeronauticalGroundLightTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AeronauticalGroundLightTimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightType) -> TimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AeronauticalGroundLightTimeSlicePropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AeronauticalGroundLightTimeSlice"}
			val := ""
			if structInstance.AeronauticalGroundLightTimeSlice != nil {
				val = structInstance.AeronauticalGroundLightTimeSlice.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightType) -> TimeSlice"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTypes {
				for _, target := range src.TimeSlice {
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

func updateStageSetTable_AeronauticalGroundLightTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.AeronauticalGroundLightTimeSliceType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpretation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SequenceNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CorrectionNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeSliceMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "FeatureLifetime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Colour"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Flashing"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StructureBeacon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AerodromeBeacon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Location"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType) -> AeronauticalGroundLightTimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AeronauticalGroundLightTimeSliceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AeronauticalGroundLightTimeSliceType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ValidTime"}
			val := ""
			if structInstance.ValidTime != nil {
				val = structInstance.ValidTime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpretation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpretation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SequenceNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.SequenceNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CorrectionNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.CorrectionNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeSliceMetadata"}
			val := ""
			if structInstance.TimeSliceMetadata != nil {
				val = structInstance.TimeSliceMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "FeatureLifetime"}
			val := ""
			if structInstance.FeatureLifetime != nil {
				val = structInstance.FeatureLifetime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			val := ""
			if structInstance.Type_XSD != nil {
				val = structInstance.Type_XSD.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Colour"}
			val := ""
			if structInstance.Colour != nil {
				val = structInstance.Colour.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Flashing"}
			val := ""
			if structInstance.Flashing != nil {
				val = structInstance.Flashing.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StructureBeacon"}
			val := ""
			if structInstance.StructureBeacon != nil {
				val = structInstance.StructureBeacon.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AerodromeBeacon"}
			val := ""
			if structInstance.AerodromeBeacon != nil {
				val = structInstance.AerodromeBeacon.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Location"}
			val := ""
			if structInstance.Location != nil {
				val = structInstance.Location.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType) -> AeronauticalGroundLightTimeSlice"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSlicePropertyTypes {
				if src.AeronauticalGroundLightTimeSlice == structInstance {
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

func updateStageSetTable_AeronauticalGroundLightType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.AeronauticalGroundLightType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "FeatureMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.BasicMessageMemberAIXMPropertyType) -> AeronauticalGroundLight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AeronauticalGroundLightType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AeronauticalGroundLightType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BoundedBy"}
			val := ""
			if structInstance.BoundedBy != nil {
				val = structInstance.BoundedBy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "FeatureMetadata"}
			val := ""
			if structInstance.FeatureMetadata != nil {
				val = structInstance.FeatureMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeSlice"}
			var names []string
			for _, elem := range structInstance.TimeSlice {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.BasicMessageMemberAIXMPropertyType) -> AeronauticalGroundLight"}
			var refNames []string
			for src := range probe.stageOfInterest.BasicMessageMemberAIXMPropertyTypes {
				if src.AeronauticalGroundLight == structInstance {
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

func updateStageSetTable_AirportHeliportPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.AirportHeliportPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> AerodromeBeacon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> Reference_ownerAirport"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirportHeliportPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AirportHeliportPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> AerodromeBeacon"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.AerodromeBeacon == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> Reference_ownerAirport"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.Reference_ownerAirport == structInstance {
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

func updateStageSetTable_CurveType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.CurveType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Segments"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.CurveType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CurveType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Segments"}
			val := ""
			if structInstance.Segments != nil {
				val = structInstance.Segments.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HorizontalAccuracy"}
			val := ""
			if structInstance.HorizontalAccuracy != nil {
				val = structInstance.HorizontalAccuracy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_ElevatedCurvePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.ElevatedCurvePropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ElevatedCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> HorizontalProjection_linearExtent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevatedCurvePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ElevatedCurvePropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ElevatedCurve"}
			val := ""
			if structInstance.ElevatedCurve != nil {
				val = structInstance.ElevatedCurve.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> HorizontalProjection_linearExtent"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.HorizontalProjection_linearExtent == structInstance {
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

func updateStageSetTable_ElevatedCurveType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.ElevatedCurveType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Segments"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Elevation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "GeoidUndulation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalDatum"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedCurvePropertyType) -> ElevatedCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevatedCurveType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ElevatedCurveType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Segments"}
			val := ""
			if structInstance.Segments != nil {
				val = structInstance.Segments.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HorizontalAccuracy"}
			val := ""
			if structInstance.HorizontalAccuracy != nil {
				val = structInstance.HorizontalAccuracy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Elevation"}
			val := ""
			if structInstance.Elevation != nil {
				val = structInstance.Elevation.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "GeoidUndulation"}
			val := ""
			if structInstance.GeoidUndulation != nil {
				val = structInstance.GeoidUndulation.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalDatum"}
			val := ""
			if structInstance.VerticalDatum != nil {
				val = structInstance.VerticalDatum.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalAccuracy"}
			val := ""
			if structInstance.VerticalAccuracy != nil {
				val = structInstance.VerticalAccuracy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedCurvePropertyType) -> ElevatedCurve"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurvePropertyTypes {
				if src.ElevatedCurve == structInstance {
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

func updateStageSetTable_ElevatedPointPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.ElevatedPointPropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ElevatedPoint"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Location"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementType) -> Location"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> HorizontalProjection_location"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevatedPointPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ElevatedPointPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ElevatedPoint"}
			val := ""
			if structInstance.ElevatedPoint != nil {
				val = structInstance.ElevatedPoint.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Location"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.Location == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementType) -> Location"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.Location == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> HorizontalProjection_location"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.HorizontalProjection_location == structInstance {
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

func updateStageSetTable_ElevatedPointType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.ElevatedPointType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Elevation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "GeoidUndulation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalDatum"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedPointPropertyType) -> ElevatedPoint"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevatedPointType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ElevatedPointType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Pos"}
			val := ""
			if structInstance.Pos != nil {
				val = structInstance.Pos.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HorizontalAccuracy"}
			val := ""
			if structInstance.HorizontalAccuracy != nil {
				val = structInstance.HorizontalAccuracy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Elevation"}
			val := ""
			if structInstance.Elevation != nil {
				val = structInstance.Elevation.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "GeoidUndulation"}
			val := ""
			if structInstance.GeoidUndulation != nil {
				val = structInstance.GeoidUndulation.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalDatum"}
			val := ""
			if structInstance.VerticalDatum != nil {
				val = structInstance.VerticalDatum.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalAccuracy"}
			val := ""
			if structInstance.VerticalAccuracy != nil {
				val = structInstance.VerticalAccuracy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedPointPropertyType) -> ElevatedPoint"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointPropertyTypes {
				if src.ElevatedPoint == structInstance {
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

func updateStageSetTable_ElevatedSurfacePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.ElevatedSurfacePropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ElevatedSurface"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> HorizontalProjection_surfaceExtent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevatedSurfacePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ElevatedSurfacePropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ElevatedSurface"}
			val := ""
			if structInstance.ElevatedSurface != nil {
				val = structInstance.ElevatedSurface.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> HorizontalProjection_surfaceExtent"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.HorizontalProjection_surfaceExtent == structInstance {
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

func updateStageSetTable_ElevatedSurfaceType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.ElevatedSurfaceType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Patches"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Elevation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "GeoidUndulation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalDatum"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedSurfacePropertyType) -> ElevatedSurface"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.SurfacePropertyType) -> ElevatedSurface"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevatedSurfaceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ElevatedSurfaceType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Patches"}
			val := ""
			if structInstance.Patches != nil {
				val = structInstance.Patches.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HorizontalAccuracy"}
			val := ""
			if structInstance.HorizontalAccuracy != nil {
				val = structInstance.HorizontalAccuracy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Elevation"}
			val := ""
			if structInstance.Elevation != nil {
				val = structInstance.Elevation.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "GeoidUndulation"}
			val := ""
			if structInstance.GeoidUndulation != nil {
				val = structInstance.GeoidUndulation.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalDatum"}
			val := ""
			if structInstance.VerticalDatum != nil {
				val = structInstance.VerticalDatum.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalAccuracy"}
			val := ""
			if structInstance.VerticalAccuracy != nil {
				val = structInstance.VerticalAccuracy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedSurfacePropertyType) -> ElevatedSurface"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfacePropertyTypes {
				if src.ElevatedSurface == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.SurfacePropertyType) -> ElevatedSurface"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfacePropertyTypes {
				if src.ElevatedSurface == structInstance {
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

func updateStageSetTable_GroundLightSystemPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.GroundLightSystemPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> SupportedGroundLight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.GroundLightSystemPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_GroundLightSystemPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> SupportedGroundLight"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.SupportedGroundLight {
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

func updateStageSetTable_LightElementPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.LightElementPropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LightElement"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> Lighting"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.LightElementPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_LightElementPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LightElement"}
			val := ""
			if structInstance.LightElement != nil {
				val = structInstance.LightElement.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> Lighting"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.Lighting {
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

func updateStageSetTable_LightElementStatusPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.LightElementStatusPropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LightElementStatus"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementType) -> Availability"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.LightElementStatusPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_LightElementStatusPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LightElementStatus"}
			val := ""
			if structInstance.LightElementStatus != nil {
				val = structInstance.LightElementStatus.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementType) -> Availability"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				for _, target := range src.Availability {
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

func updateStageSetTable_LightElementStatusType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.LightElementStatusType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeInterval"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SpecialDateAuthority"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Status"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementStatusPropertyType) -> LightElementStatus"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.LightElementStatusType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_LightElementStatusType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeInterval"}
			var names []string
			for _, elem := range structInstance.TimeInterval {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SpecialDateAuthority"}
			var names []string
			for _, elem := range structInstance.SpecialDateAuthority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Status"}
			val := ""
			if structInstance.Status != nil {
				val = structInstance.Status.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementStatusPropertyType) -> LightElementStatus"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusPropertyTypes {
				if src.LightElementStatus == structInstance {
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

func updateStageSetTable_LightElementType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.LightElementType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Colour"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IntensityLevel"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Intensity"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Location"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Availability"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementPropertyType) -> LightElement"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.LightElementType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_LightElementType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Colour"}
			val := ""
			if structInstance.Colour != nil {
				val = structInstance.Colour.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IntensityLevel"}
			val := ""
			if structInstance.IntensityLevel != nil {
				val = structInstance.IntensityLevel.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Intensity"}
			val := ""
			if structInstance.Intensity != nil {
				val = structInstance.Intensity.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			val := ""
			if structInstance.Type_XSD != nil {
				val = structInstance.Type_XSD.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Location"}
			val := ""
			if structInstance.Location != nil {
				val = structInstance.Location.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Availability"}
			var names []string
			for _, elem := range structInstance.Availability {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementPropertyType) -> LightElement"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementPropertyTypes {
				if src.LightElement == structInstance {
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

func updateStageSetTable_LinguisticNotePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.LinguisticNotePropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LinguisticNote"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.NoteType) -> TranslatedNote"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.LinguisticNotePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_LinguisticNotePropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LinguisticNote"}
			val := ""
			if structInstance.LinguisticNote != nil {
				val = structInstance.LinguisticNote.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.NoteType) -> TranslatedNote"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				for _, target := range src.TranslatedNote {
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

func updateStageSetTable_LinguisticNoteType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.LinguisticNoteType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LinguisticNotePropertyType) -> LinguisticNote"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.LinguisticNoteType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_LinguisticNoteType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
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
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LinguisticNotePropertyType) -> LinguisticNote"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LinguisticNotePropertyTypes {
				if src.LinguisticNote == structInstance {
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

func updateStageSetTable_MarkerBeaconPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.MarkerBeaconPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> Marker"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.MarkerBeaconPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_MarkerBeaconPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> Marker"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Marker == structInstance {
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

func updateStageSetTable_NavaidEquipmentPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.NavaidEquipmentPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> HostedNavaidEquipment"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidEquipmentPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_NavaidEquipmentPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> HostedNavaidEquipment"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.HostedNavaidEquipment {
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

func updateStageSetTable_NotePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.NotePropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.CurveType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedCurveType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedPointType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedSurfaceType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementStatusType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.PointType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.SurfaceType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureLightingStatusType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NotePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_NotePropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
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
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.CurveType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurveTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedCurveType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedPointType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedSurfaceType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementStatusType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.PointType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.SurfaceType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureLightingStatusType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.Annotation {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> Annotation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.Annotation {
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

func updateStageSetTable_NoteType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.NoteType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PropertyName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Purpose"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TranslatedNote"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.NotePropertyType) -> Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NoteType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_NoteType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PropertyName"}
			val := ""
			if structInstance.PropertyName != nil {
				val = structInstance.PropertyName.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Purpose"}
			val := ""
			if structInstance.Purpose != nil {
				val = structInstance.Purpose.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TranslatedNote"}
			var names []string
			for _, elem := range structInstance.TranslatedNote {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.NotePropertyType) -> Note"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NotePropertyTypes {
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

func updateStageSetTable_ObstacleAreaTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.ObstacleAreaTimeSlicePropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ObstacleAreaTimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaType) -> TimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ObstacleAreaTimeSlicePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ObstacleAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ObstacleAreaTimeSlice"}
			val := ""
			if structInstance.ObstacleAreaTimeSlice != nil {
				val = structInstance.ObstacleAreaTimeSlice.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaType) -> TimeSlice"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTypes {
				for _, target := range src.TimeSlice {
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

func updateStageSetTable_ObstacleAreaTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.ObstacleAreaTimeSliceType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpretation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SequenceNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CorrectionNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeSliceMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "FeatureLifetime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ObstructionIdSurfaceCondition"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Reference_ownerOrganisation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Reference_ownerRunway"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Reference_ownerAirport"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SurfaceExtent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Obstacle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSlicePropertyType) -> ObstacleAreaTimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ObstacleAreaTimeSliceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ObstacleAreaTimeSliceType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ValidTime"}
			val := ""
			if structInstance.ValidTime != nil {
				val = structInstance.ValidTime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpretation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpretation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SequenceNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.SequenceNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CorrectionNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.CorrectionNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeSliceMetadata"}
			val := ""
			if structInstance.TimeSliceMetadata != nil {
				val = structInstance.TimeSliceMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "FeatureLifetime"}
			val := ""
			if structInstance.FeatureLifetime != nil {
				val = structInstance.FeatureLifetime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			val := ""
			if structInstance.Type_XSD != nil {
				val = structInstance.Type_XSD.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ObstructionIdSurfaceCondition"}
			val := ""
			if structInstance.ObstructionIdSurfaceCondition != nil {
				val = structInstance.ObstructionIdSurfaceCondition.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Reference_ownerOrganisation"}
			val := ""
			if structInstance.Reference_ownerOrganisation != nil {
				val = structInstance.Reference_ownerOrganisation.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Reference_ownerRunway"}
			val := ""
			if structInstance.Reference_ownerRunway != nil {
				val = structInstance.Reference_ownerRunway.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Reference_ownerAirport"}
			val := ""
			if structInstance.Reference_ownerAirport != nil {
				val = structInstance.Reference_ownerAirport.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SurfaceExtent"}
			val := ""
			if structInstance.SurfaceExtent != nil {
				val = structInstance.SurfaceExtent.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Obstacle"}
			var names []string
			for _, elem := range structInstance.Obstacle {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSlicePropertyType) -> ObstacleAreaTimeSlice"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSlicePropertyTypes {
				if src.ObstacleAreaTimeSlice == structInstance {
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

func updateStageSetTable_ObstacleAreaType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.ObstacleAreaType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "FeatureMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.BasicMessageMemberAIXMPropertyType) -> ObstacleArea"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ObstacleAreaType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ObstacleAreaType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BoundedBy"}
			val := ""
			if structInstance.BoundedBy != nil {
				val = structInstance.BoundedBy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "FeatureMetadata"}
			val := ""
			if structInstance.FeatureMetadata != nil {
				val = structInstance.FeatureMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeSlice"}
			var names []string
			for _, elem := range structInstance.TimeSlice {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.BasicMessageMemberAIXMPropertyType) -> ObstacleArea"}
			var refNames []string
			for src := range probe.stageOfInterest.BasicMessageMemberAIXMPropertyTypes {
				if src.ObstacleArea == structInstance {
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

func updateStageSetTable_OrganisationAuthorityPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.OrganisationAuthorityPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementStatusType) -> SpecialDateAuthority"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> Reference_ownerOrganisation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureLightingStatusType) -> SpecialDateAuthority"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> SpecialDateAuthority"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> HostedOrganisation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.OrganisationAuthorityPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_OrganisationAuthorityPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementStatusType) -> SpecialDateAuthority"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				for _, target := range src.SpecialDateAuthority {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> Reference_ownerOrganisation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.Reference_ownerOrganisation == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureLightingStatusType) -> SpecialDateAuthority"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				for _, target := range src.SpecialDateAuthority {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> SpecialDateAuthority"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.SpecialDateAuthority {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> HostedOrganisation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.HostedOrganisation {
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

func updateStageSetTable_PassengerServicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.PassengerServicePropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> HostedPassengerService"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PassengerServicePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_PassengerServicePropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> HostedPassengerService"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.HostedPassengerService {
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

func updateStageSetTable_PointType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.PointType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PointType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_PointType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Pos"}
			val := ""
			if structInstance.Pos != nil {
				val = structInstance.Pos.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HorizontalAccuracy"}
			val := ""
			if structInstance.HorizontalAccuracy != nil {
				val = structInstance.HorizontalAccuracy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_RunwayDirectionPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.RunwayDirectionPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> Reference_ownerRunway"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RunwayDirectionPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_RunwayDirectionPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> Reference_ownerRunway"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.Reference_ownerRunway == structInstance {
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

func updateStageSetTable_ServicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.ServicePropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> SupportedService"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ServicePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ServicePropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> SupportedService"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.SupportedService {
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

func updateStageSetTable_SpecialNavigationStationPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.SpecialNavigationStationPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> HostedSpecialNavStation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SpecialNavigationStationPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_SpecialNavigationStationPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> HostedSpecialNavStation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.HostedSpecialNavStation {
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

func updateStageSetTable_SurfacePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.SurfacePropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Surface"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ElevatedSurface"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> SurfaceExtent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SurfacePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_SurfacePropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Surface"}
			val := ""
			if structInstance.Surface != nil {
				val = structInstance.Surface.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ElevatedSurface"}
			val := ""
			if structInstance.ElevatedSurface != nil {
				val = structInstance.ElevatedSurface.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> SurfaceExtent"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.SurfaceExtent == structInstance {
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

func updateStageSetTable_SurfaceType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.SurfaceType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Patches"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.SurfacePropertyType) -> Surface"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SurfaceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_SurfaceType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Patches"}
			val := ""
			if structInstance.Patches != nil {
				val = structInstance.Patches.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HorizontalAccuracy"}
			val := ""
			if structInstance.HorizontalAccuracy != nil {
				val = structInstance.HorizontalAccuracy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.SurfacePropertyType) -> Surface"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfacePropertyTypes {
				if src.Surface == structInstance {
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

func updateStageSetTable_TimesheetPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.TimesheetPropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Timesheet"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementStatusType) -> TimeInterval"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureLightingStatusType) -> TimeInterval"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> TimeInterval"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.TimesheetPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_TimesheetPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Timesheet"}
			val := ""
			if structInstance.Timesheet != nil {
				val = structInstance.Timesheet.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementStatusType) -> TimeInterval"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				for _, target := range src.TimeInterval {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureLightingStatusType) -> TimeInterval"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				for _, target := range src.TimeInterval {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> TimeInterval"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.TimeInterval {
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

func updateStageSetTable_TimesheetType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.TimesheetType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeReference"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartDate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndDate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Day"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DayTil"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartEvent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartTimeRelativeEvent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartEventInterpretation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndEvent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndTimeRelativeEvent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndEventInterpretation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DaylightSavingAdjust"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Excluded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetPropertyType) -> Timesheet"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.TimesheetType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_TimesheetType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeReference"}
			val := ""
			if structInstance.TimeReference != nil {
				val = structInstance.TimeReference.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartDate"}
			val := ""
			if structInstance.StartDate != nil {
				val = structInstance.StartDate.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndDate"}
			val := ""
			if structInstance.EndDate != nil {
				val = structInstance.EndDate.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Day"}
			val := ""
			if structInstance.Day != nil {
				val = structInstance.Day.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DayTil"}
			val := ""
			if structInstance.DayTil != nil {
				val = structInstance.DayTil.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartTime"}
			val := ""
			if structInstance.StartTime != nil {
				val = structInstance.StartTime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartEvent"}
			val := ""
			if structInstance.StartEvent != nil {
				val = structInstance.StartEvent.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartTimeRelativeEvent"}
			val := ""
			if structInstance.StartTimeRelativeEvent != nil {
				val = structInstance.StartTimeRelativeEvent.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartEventInterpretation"}
			val := ""
			if structInstance.StartEventInterpretation != nil {
				val = structInstance.StartEventInterpretation.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndTime"}
			val := ""
			if structInstance.EndTime != nil {
				val = structInstance.EndTime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndEvent"}
			val := ""
			if structInstance.EndEvent != nil {
				val = structInstance.EndEvent.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndTimeRelativeEvent"}
			val := ""
			if structInstance.EndTimeRelativeEvent != nil {
				val = structInstance.EndTimeRelativeEvent.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndEventInterpretation"}
			val := ""
			if structInstance.EndEventInterpretation != nil {
				val = structInstance.EndEventInterpretation.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DaylightSavingAdjust"}
			val := ""
			if structInstance.DaylightSavingAdjust != nil {
				val = structInstance.DaylightSavingAdjust.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Excluded"}
			val := ""
			if structInstance.Excluded != nil {
				val = structInstance.Excluded.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetPropertyType) -> Timesheet"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetPropertyTypes {
				if src.Timesheet == structInstance {
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

func updateStageSetTable_UnitPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.UnitPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> HostedUnit"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.UnitPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_UnitPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> HostedUnit"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.HostedUnit {
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

func updateStageSetTable_VerticalStructureLightingStatusPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.VerticalStructureLightingStatusPropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalStructureLightingStatus"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> LightingAvailability"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VerticalStructureLightingStatusPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_VerticalStructureLightingStatusPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalStructureLightingStatus"}
			val := ""
			if structInstance.VerticalStructureLightingStatus != nil {
				val = structInstance.VerticalStructureLightingStatus.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> LightingAvailability"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.LightingAvailability {
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

func updateStageSetTable_VerticalStructureLightingStatusType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.VerticalStructureLightingStatusType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeInterval"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SpecialDateAuthority"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Status"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureLightingStatusPropertyType) -> VerticalStructureLightingStatus"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VerticalStructureLightingStatusType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_VerticalStructureLightingStatusType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeInterval"}
			var names []string
			for _, elem := range structInstance.TimeInterval {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SpecialDateAuthority"}
			var names []string
			for _, elem := range structInstance.SpecialDateAuthority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Status"}
			val := ""
			if structInstance.Status != nil {
				val = structInstance.Status.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureLightingStatusPropertyType) -> VerticalStructureLightingStatus"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusPropertyTypes {
				if src.VerticalStructureLightingStatus == structInstance {
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

func updateStageSetTable_VerticalStructurePartPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.VerticalStructurePartPropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalStructurePart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> Part"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VerticalStructurePartPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_VerticalStructurePartPropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalStructurePart"}
			val := ""
			if structInstance.VerticalStructurePart != nil {
				val = structInstance.VerticalStructurePart.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> Part"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.Part {
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

func updateStageSetTable_VerticalStructurePartType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.VerticalStructurePartType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeInterval"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SpecialDateAuthority"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalExtent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalExtentAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ConstructionStatus"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "MarkingPattern"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "MarkingFirstColour"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "MarkingSecondColour"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Mobile"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Frangible"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VisibleMaterial"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Designator"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HorizontalProjection_surfaceExtent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HorizontalProjection_linearExtent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HorizontalProjection_location"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Lighting"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartPropertyType) -> VerticalStructurePart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VerticalStructurePartType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_VerticalStructurePartType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeInterval"}
			var names []string
			for _, elem := range structInstance.TimeInterval {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SpecialDateAuthority"}
			var names []string
			for _, elem := range structInstance.SpecialDateAuthority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalExtent"}
			val := ""
			if structInstance.VerticalExtent != nil {
				val = structInstance.VerticalExtent.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalExtentAccuracy"}
			val := ""
			if structInstance.VerticalExtentAccuracy != nil {
				val = structInstance.VerticalExtentAccuracy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			val := ""
			if structInstance.Type_XSD != nil {
				val = structInstance.Type_XSD.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ConstructionStatus"}
			val := ""
			if structInstance.ConstructionStatus != nil {
				val = structInstance.ConstructionStatus.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "MarkingPattern"}
			val := ""
			if structInstance.MarkingPattern != nil {
				val = structInstance.MarkingPattern.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "MarkingFirstColour"}
			val := ""
			if structInstance.MarkingFirstColour != nil {
				val = structInstance.MarkingFirstColour.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "MarkingSecondColour"}
			val := ""
			if structInstance.MarkingSecondColour != nil {
				val = structInstance.MarkingSecondColour.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Mobile"}
			val := ""
			if structInstance.Mobile != nil {
				val = structInstance.Mobile.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Frangible"}
			val := ""
			if structInstance.Frangible != nil {
				val = structInstance.Frangible.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VisibleMaterial"}
			val := ""
			if structInstance.VisibleMaterial != nil {
				val = structInstance.VisibleMaterial.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Designator"}
			val := ""
			if structInstance.Designator != nil {
				val = structInstance.Designator.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HorizontalProjection_surfaceExtent"}
			val := ""
			if structInstance.HorizontalProjection_surfaceExtent != nil {
				val = structInstance.HorizontalProjection_surfaceExtent.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HorizontalProjection_linearExtent"}
			val := ""
			if structInstance.HorizontalProjection_linearExtent != nil {
				val = structInstance.HorizontalProjection_linearExtent.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HorizontalProjection_location"}
			val := ""
			if structInstance.HorizontalProjection_location != nil {
				val = structInstance.HorizontalProjection_location.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Lighting"}
			var names []string
			for _, elem := range structInstance.Lighting {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartPropertyType) -> VerticalStructurePart"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartPropertyTypes {
				if src.VerticalStructurePart == structInstance {
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

func updateStageSetTable_VerticalStructurePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.VerticalStructurePropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> StructureBeacon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> Obstacle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VerticalStructurePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_VerticalStructurePropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> StructureBeacon"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.StructureBeacon == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> Obstacle"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				for _, target := range src.Obstacle {
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

func updateStageSetTable_VerticalStructureTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.VerticalStructureTimeSlicePropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalStructureTimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureType) -> TimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VerticalStructureTimeSlicePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_VerticalStructureTimeSlicePropertyType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalStructureTimeSlice"}
			val := ""
			if structInstance.VerticalStructureTimeSlice != nil {
				val = structInstance.VerticalStructureTimeSlice.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureType) -> TimeSlice"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTypes {
				for _, target := range src.TimeSlice {
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

func updateStageSetTable_VerticalStructureTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.VerticalStructureTimeSliceType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpretation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SequenceNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CorrectionNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeSliceMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "FeatureLifetime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Lighted"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "MarkingICAOStandard"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Group"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Length"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Width"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Radius"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LightingICAOStandard"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SynchronisedLighting"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Marker"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Part"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HostedPassengerService"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SupportedGroundLight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HostedNavaidEquipment"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HostedSpecialNavStation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HostedUnit"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HostedOrganisation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SupportedService"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Annotation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LightingAvailability"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Extension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSlicePropertyType) -> VerticalStructureTimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VerticalStructureTimeSliceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_VerticalStructureTimeSliceType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ValidTime"}
			val := ""
			if structInstance.ValidTime != nil {
				val = structInstance.ValidTime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpretation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpretation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SequenceNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.SequenceNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CorrectionNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.CorrectionNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeSliceMetadata"}
			val := ""
			if structInstance.TimeSliceMetadata != nil {
				val = structInstance.TimeSliceMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "FeatureLifetime"}
			val := ""
			if structInstance.FeatureLifetime != nil {
				val = structInstance.FeatureLifetime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			val := ""
			if structInstance.Type_XSD != nil {
				val = structInstance.Type_XSD.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Lighted"}
			val := ""
			if structInstance.Lighted != nil {
				val = structInstance.Lighted.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "MarkingICAOStandard"}
			val := ""
			if structInstance.MarkingICAOStandard != nil {
				val = structInstance.MarkingICAOStandard.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Group"}
			val := ""
			if structInstance.Group != nil {
				val = structInstance.Group.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Length"}
			val := ""
			if structInstance.Length != nil {
				val = structInstance.Length.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Width"}
			val := ""
			if structInstance.Width != nil {
				val = structInstance.Width.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Radius"}
			val := ""
			if structInstance.Radius != nil {
				val = structInstance.Radius.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LightingICAOStandard"}
			val := ""
			if structInstance.LightingICAOStandard != nil {
				val = structInstance.LightingICAOStandard.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SynchronisedLighting"}
			val := ""
			if structInstance.SynchronisedLighting != nil {
				val = structInstance.SynchronisedLighting.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Marker"}
			val := ""
			if structInstance.Marker != nil {
				val = structInstance.Marker.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Part"}
			var names []string
			for _, elem := range structInstance.Part {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HostedPassengerService"}
			var names []string
			for _, elem := range structInstance.HostedPassengerService {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SupportedGroundLight"}
			var names []string
			for _, elem := range structInstance.SupportedGroundLight {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HostedNavaidEquipment"}
			var names []string
			for _, elem := range structInstance.HostedNavaidEquipment {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HostedSpecialNavStation"}
			var names []string
			for _, elem := range structInstance.HostedSpecialNavStation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HostedUnit"}
			var names []string
			for _, elem := range structInstance.HostedUnit {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HostedOrganisation"}
			var names []string
			for _, elem := range structInstance.HostedOrganisation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SupportedService"}
			var names []string
			for _, elem := range structInstance.SupportedService {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Annotation"}
			var names []string
			for _, elem := range structInstance.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LightingAvailability"}
			var names []string
			for _, elem := range structInstance.LightingAvailability {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Extension"}
			var names []string
			for _, elem := range structInstance.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSlicePropertyType) -> VerticalStructureTimeSlice"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSlicePropertyTypes {
				if src.VerticalStructureTimeSlice == structInstance {
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

func updateStageSetTable_VerticalStructureType_AixmfeaturesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmfeatures.VerticalStructureType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "FeatureMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeSlice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.BasicMessageMemberAIXMPropertyType) -> VerticalStructure"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VerticalStructureType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmfeaturesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmfeaturesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_VerticalStructureType_AixmfeaturesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BoundedBy"}
			val := ""
			if structInstance.BoundedBy != nil {
				val = structInstance.BoundedBy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "FeatureMetadata"}
			val := ""
			if structInstance.FeatureMetadata != nil {
				val = structInstance.FeatureMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeSlice"}
			var names []string
			for _, elem := range structInstance.TimeSlice {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.BasicMessageMemberAIXMPropertyType) -> VerticalStructure"}
			var refNames []string
			for src := range probe.stageOfInterest.BasicMessageMemberAIXMPropertyTypes {
				if src.VerticalStructure == structInstance {
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

func updateStageSetTable_CodeColourType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeColourType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Colour"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementType) -> Colour"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> MarkingFirstColour"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> MarkingSecondColour"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeColourType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeColourType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Colour"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.Colour == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementType) -> Colour"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.Colour == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> MarkingFirstColour"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.MarkingFirstColour == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> MarkingSecondColour"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.MarkingSecondColour == structInstance {
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

func updateStageSetTable_CodeDayType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeDayType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> Day"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> DayTil"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeDayType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeDayType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> Day"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.Day == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> DayTil"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.DayTil == structInstance {
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

func updateStageSetTable_CodeGroundLightingType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeGroundLightingType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeGroundLightingType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeGroundLightingType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Type_XSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.Type_XSD == structInstance {
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

func updateStageSetTable_CodeLightIntensityType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeLightIntensityType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementType) -> IntensityLevel"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeLightIntensityType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeLightIntensityType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementType) -> IntensityLevel"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.IntensityLevel == structInstance {
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

func updateStageSetTable_CodeLightSourceType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeLightSourceType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementType) -> Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeLightSourceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeLightSourceType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementType) -> Type_XSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.Type_XSD == structInstance {
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

func updateStageSetTable_CodeNotePurposeType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeNotePurposeType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.NoteType) -> Purpose"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeNotePurposeType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeNotePurposeType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.NoteType) -> Purpose"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				if src.Purpose == structInstance {
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

func updateStageSetTable_CodeObstacleAreaType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeObstacleAreaType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeObstacleAreaType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeObstacleAreaType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> Type_XSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.Type_XSD == structInstance {
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

func updateStageSetTable_CodeObstacleAssessmentSurfaceType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeObstacleAssessmentSurfaceType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> ObstructionIdSurfaceCondition"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeObstacleAssessmentSurfaceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeObstacleAssessmentSurfaceType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> ObstructionIdSurfaceCondition"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.ObstructionIdSurfaceCondition == structInstance {
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

func updateStageSetTable_CodeStatusConstructionType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeStatusConstructionType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> ConstructionStatus"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeStatusConstructionType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeStatusConstructionType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> ConstructionStatus"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.ConstructionStatus == structInstance {
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

func updateStageSetTable_CodeStatusOperationsType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeStatusOperationsType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementStatusType) -> Status"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureLightingStatusType) -> Status"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeStatusOperationsType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeStatusOperationsType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementStatusType) -> Status"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				if src.Status == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureLightingStatusType) -> Status"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				if src.Status == structInstance {
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

func updateStageSetTable_CodeTimeEventCombinationType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeTimeEventCombinationType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> StartEventInterpretation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> EndEventInterpretation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeTimeEventCombinationType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeTimeEventCombinationType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> StartEventInterpretation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.StartEventInterpretation == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> EndEventInterpretation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.EndEventInterpretation == structInstance {
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

func updateStageSetTable_CodeTimeEventType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeTimeEventType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> StartEvent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> EndEvent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeTimeEventType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeTimeEventType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> StartEvent"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.StartEvent == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> EndEvent"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.EndEvent == structInstance {
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

func updateStageSetTable_CodeTimeReferenceType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeTimeReferenceType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> TimeReference"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeTimeReferenceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeTimeReferenceType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> TimeReference"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.TimeReference == structInstance {
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

func updateStageSetTable_CodeVerticalDatumType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeVerticalDatumType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedCurveType) -> VerticalDatum"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedPointType) -> VerticalDatum"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedSurfaceType) -> VerticalDatum"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeVerticalDatumType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeVerticalDatumType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedCurveType) -> VerticalDatum"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.VerticalDatum == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedPointType) -> VerticalDatum"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.VerticalDatum == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedSurfaceType) -> VerticalDatum"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.VerticalDatum == structInstance {
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

func updateStageSetTable_CodeVerticalStructureMarkingType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeVerticalStructureMarkingType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> MarkingPattern"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeVerticalStructureMarkingType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeVerticalStructureMarkingType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> MarkingPattern"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.MarkingPattern == structInstance {
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

func updateStageSetTable_CodeVerticalStructureMaterialType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeVerticalStructureMaterialType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> VisibleMaterial"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeVerticalStructureMaterialType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeVerticalStructureMaterialType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> VisibleMaterial"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.VisibleMaterial == structInstance {
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

func updateStageSetTable_CodeVerticalStructureType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeVerticalStructureType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeVerticalStructureType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeVerticalStructureType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> Type_XSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.Type_XSD == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> Type_XSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Type_XSD == structInstance {
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

func updateStageSetTable_CodeYesNoType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.CodeYesNoType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Flashing"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> DaylightSavingAdjust"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> Excluded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> Mobile"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> Frangible"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> Lighted"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> MarkingICAOStandard"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> Group"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> LightingICAOStandard"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> SynchronisedLighting"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeYesNoType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeYesNoType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Flashing"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.Flashing == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> DaylightSavingAdjust"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.DaylightSavingAdjust == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> Excluded"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.Excluded == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> Mobile"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.Mobile == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> Frangible"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.Frangible == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> Lighted"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Lighted == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> MarkingICAOStandard"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.MarkingICAOStandard == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> Group"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Group == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> LightingICAOStandard"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.LightingICAOStandard == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> SynchronisedLighting"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.SynchronisedLighting == structInstance {
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

func updateStageSetTable_DateMonthDayType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.DateMonthDayType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> StartDate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> EndDate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.DateMonthDayType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_DateMonthDayType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> StartDate"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.StartDate == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> EndDate"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.EndDate == structInstance {
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

func updateStageSetTable_TextDesignatorType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.TextDesignatorType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> Designator"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TextDesignatorType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_TextDesignatorType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> Designator"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.Designator == structInstance {
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

func updateStageSetTable_TextNameType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.TextNameType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TextNameType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_TextNameType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TextNoteType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.TextNoteType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Lang"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LinguisticNoteType) -> Note"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TextNoteType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_TextNoteType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Lang"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Lang)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LinguisticNoteType) -> Note"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LinguisticNoteTypes {
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

func updateStageSetTable_TextPropertyNameType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.TextPropertyNameType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.NoteType) -> PropertyName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TextPropertyNameType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_TextPropertyNameType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.NoteType) -> PropertyName"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				if src.PropertyName == structInstance {
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

func updateStageSetTable_TimeType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.TimeType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> StartTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> EndTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TimeType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_TimeType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> StartTime"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.StartTime == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> EndTime"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.EndTime == structInstance {
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

func updateStageSetTable_ValDistanceSignedType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.ValDistanceSignedType"
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
		col.Name = "Uom"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedCurveType) -> GeoidUndulation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedPointType) -> GeoidUndulation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedSurfaceType) -> GeoidUndulation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValDistanceSignedType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ValDistanceSignedType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "Uom"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Uom)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedCurveType) -> GeoidUndulation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.GeoidUndulation == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedPointType) -> GeoidUndulation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.GeoidUndulation == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedSurfaceType) -> GeoidUndulation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.GeoidUndulation == structInstance {
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

func updateStageSetTable_ValDistanceType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.ValDistanceType"
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
		col.Name = "Uom"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.CurveType) -> HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedCurveType) -> HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedCurveType) -> VerticalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedPointType) -> HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedPointType) -> VerticalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedSurfaceType) -> HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedSurfaceType) -> VerticalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.PointType) -> HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.SurfaceType) -> HorizontalAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> VerticalExtent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> VerticalExtentAccuracy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> Length"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> Width"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> Radius"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValDistanceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ValDistanceType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "Uom"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Uom)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.CurveType) -> HorizontalAccuracy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurveTypes {
				if src.HorizontalAccuracy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedCurveType) -> HorizontalAccuracy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.HorizontalAccuracy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedCurveType) -> VerticalAccuracy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.VerticalAccuracy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedPointType) -> HorizontalAccuracy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.HorizontalAccuracy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedPointType) -> VerticalAccuracy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.VerticalAccuracy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedSurfaceType) -> HorizontalAccuracy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.HorizontalAccuracy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedSurfaceType) -> VerticalAccuracy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.VerticalAccuracy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.PointType) -> HorizontalAccuracy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointTypes {
				if src.HorizontalAccuracy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.SurfaceType) -> HorizontalAccuracy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes {
				if src.HorizontalAccuracy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> VerticalExtent"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.VerticalExtent == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> VerticalExtentAccuracy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.VerticalExtentAccuracy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> Length"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Length == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> Width"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Width == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> Radius"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Radius == structInstance {
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

func updateStageSetTable_ValDistanceVerticalType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.ValDistanceVerticalType"
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
		col.Name = "Uom"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedCurveType) -> Elevation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedPointType) -> Elevation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedSurfaceType) -> Elevation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValDistanceVerticalType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ValDistanceVerticalType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "Uom"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Uom)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedCurveType) -> Elevation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.Elevation == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedPointType) -> Elevation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.Elevation == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedSurfaceType) -> Elevation"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.Elevation == structInstance {
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

func updateStageSetTable_ValDurationType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.ValDurationType"
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
		col.Name = "Uom"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> StartTimeRelativeEvent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> EndTimeRelativeEvent"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValDurationType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ValDurationType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "Uom"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Uom)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> StartTimeRelativeEvent"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.StartTimeRelativeEvent == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> EndTimeRelativeEvent"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.EndTimeRelativeEvent == structInstance {
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

func updateStageSetTable_ValLightIntensityType_AixmdatatypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmdatatypes.ValLightIntensityType"
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
		col.Name = "Uom"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementType) -> Intensity"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValLightIntensityType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmdatatypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmdatatypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ValLightIntensityType_AixmdatatypesStage(probe)
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
			cell := &table_models.Cell{Name: "Uom"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Uom)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementType) -> Intensity"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.Intensity == structInstance {
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

func updateStageSetTable_AbstractAIXMFeatureBaseType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractAIXMFeatureBaseType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BoundedBy"}
			val := ""
			if structInstance.BoundedBy != nil {
				val = structInstance.BoundedBy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractAIXMFeatureType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.AbstractAIXMFeatureType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "FeatureMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMFeatureType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractAIXMFeatureType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BoundedBy"}
			val := ""
			if structInstance.BoundedBy != nil {
				val = structInstance.BoundedBy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "FeatureMetadata"}
			val := ""
			if structInstance.FeatureMetadata != nil {
				val = structInstance.FeatureMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractAIXMMessageBaseType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SequenceNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "MessageMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractAIXMMessageBaseType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BoundedBy"}
			val := ""
			if structInstance.BoundedBy != nil {
				val = structInstance.BoundedBy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SequenceNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.SequenceNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "MessageMetadata"}
			val := ""
			if structInstance.MessageMetadata != nil {
				val = structInstance.MessageMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractAIXMMessageType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.AbstractAIXMMessageType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SequenceNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "MessageMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AggregationType"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMMessageType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractAIXMMessageType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BoundedBy"}
			val := ""
			if structInstance.BoundedBy != nil {
				val = structInstance.BoundedBy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SequenceNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.SequenceNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "MessageMetadata"}
			val := ""
			if structInstance.MessageMetadata != nil {
				val = structInstance.MessageMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AggregationType"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.AggregationType)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractAIXMObjectType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.AbstractAIXMObjectType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMObjectType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractAIXMObjectType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractAIXMPropertyType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.AbstractAIXMPropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractAIXMPropertyType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractAIXMTimeSliceBaseType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractAIXMTimeSliceBaseType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ValidTime"}
			val := ""
			if structInstance.ValidTime != nil {
				val = structInstance.ValidTime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractAIXMTimeSliceType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpretation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SequenceNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CorrectionNumber"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeSliceMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "FeatureLifetime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractAIXMTimeSliceType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ValidTime"}
			val := ""
			if structInstance.ValidTime != nil {
				val = structInstance.ValidTime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpretation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpretation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SequenceNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.SequenceNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CorrectionNumber"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.CorrectionNumber)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeSliceMetadata"}
			val := ""
			if structInstance.TimeSliceMetadata != nil {
				val = structInstance.TimeSliceMetadata.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "FeatureLifetime"}
			val := ""
			if structInstance.FeatureLifetime != nil {
				val = structInstance.FeatureLifetime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractExtensionType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.AbstractExtensionType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension) -> AbstractElevatedCurveExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_1) -> AbstractElevatedPointExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_10) -> AbstractLightElementExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_11) -> AbstractPropertiesWithScheduleExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_11) -> AbstractLightElementStatusExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_12) -> AbstractTimesheetExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_2) -> AbstractElevatedSurfaceExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_3) -> AbstractAeronauticalGroundLightExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_4) -> AbstractLinguisticNoteExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_5) -> AbstractNoteExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_6) -> AbstractObstacleAreaExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_7) -> AbstractVerticalStructureExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_8) -> AbstractPropertiesWithScheduleExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_8) -> AbstractVerticalStructureLightingStatusExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_9) -> AbstractPropertiesWithScheduleExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.A_extension_9) -> AbstractVerticalStructurePartExtension"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractExtensionType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractExtensionType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension) -> AbstractElevatedCurveExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extensions {
				for _, target := range src.AbstractElevatedCurveExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_1) -> AbstractElevatedPointExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_1s {
				for _, target := range src.AbstractElevatedPointExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_10) -> AbstractLightElementExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_10s {
				for _, target := range src.AbstractLightElementExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_11) -> AbstractPropertiesWithScheduleExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_11s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_11) -> AbstractLightElementStatusExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_11s {
				for _, target := range src.AbstractLightElementStatusExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_12) -> AbstractTimesheetExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_12s {
				for _, target := range src.AbstractTimesheetExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_2) -> AbstractElevatedSurfaceExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_2s {
				for _, target := range src.AbstractElevatedSurfaceExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_3) -> AbstractAeronauticalGroundLightExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_3s {
				for _, target := range src.AbstractAeronauticalGroundLightExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_4) -> AbstractLinguisticNoteExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_4s {
				for _, target := range src.AbstractLinguisticNoteExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_5) -> AbstractNoteExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_5s {
				for _, target := range src.AbstractNoteExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_6) -> AbstractObstacleAreaExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_6s {
				for _, target := range src.AbstractObstacleAreaExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_7) -> AbstractVerticalStructureExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_7s {
				for _, target := range src.AbstractVerticalStructureExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_8) -> AbstractPropertiesWithScheduleExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_8s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_8) -> AbstractVerticalStructureLightingStatusExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_8s {
				for _, target := range src.AbstractVerticalStructureLightingStatusExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_9) -> AbstractPropertiesWithScheduleExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_9s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.A_extension_9) -> AbstractVerticalStructurePartExtension"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_9s {
				for _, target := range src.AbstractVerticalStructurePartExtension {
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

func updateStageSetTable_FeatureMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.FeatureMetadataPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightType) -> FeatureMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaType) -> FeatureMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureType) -> FeatureMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureType) -> FeatureMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_FeatureMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightType) -> FeatureMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTypes {
				if src.FeatureMetadata == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaType) -> FeatureMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTypes {
				if src.FeatureMetadata == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureType) -> FeatureMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTypes {
				if src.FeatureMetadata == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureType) -> FeatureMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureTypes {
				if src.FeatureMetadata == structInstance {
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

func updateStageSetTable_FeatureTimeSliceMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> TimeSliceMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> TimeSliceMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> TimeSliceMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType) -> TimeSliceMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_FeatureTimeSliceMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> TimeSliceMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.TimeSliceMetadata == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> TimeSliceMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.TimeSliceMetadata == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> TimeSliceMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.TimeSliceMetadata == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType) -> TimeSliceMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes {
				if src.TimeSliceMetadata == structInstance {
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

func updateStageSetTable_MessageMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "aixmabstractgmlobjecttypes.MessageMetadataPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AIXMBasicMessage) -> MessageMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AIXMBasicMessageType) -> MessageMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType) -> MessageMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMMessageType) -> MessageMetadata"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.MessageMetadataPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_MessageMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AIXMBasicMessage) -> MessageMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessages {
				if src.MessageMetadata == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AIXMBasicMessageType) -> MessageMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessageTypes {
				if src.MessageMetadata == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType) -> MessageMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageBaseTypes {
				if src.MessageMetadata == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMMessageType) -> MessageMetadata"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageTypes {
				if src.MessageMetadata == structInstance {
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

func updateStageSetTable_AbstractCurveSegmentType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractCurveSegmentType"
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
		col.Name = "NumDerivativesAtStart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativesAtEnd"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativeInterior"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractCurveSegmentType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractCurveSegmentType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NumDerivativesAtStart"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtStart)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativesAtEnd"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtEnd)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativeInterior"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativeInterior)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractCurveType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractCurveType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractCurveType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractCurveType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractFeatureMemberType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractFeatureMemberType"
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

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractFeatureMemberType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractFeatureMemberType_GmlStage(probe)
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

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractFeatureType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractFeatureType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractFeatureType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractFeatureType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BoundedBy"}
			val := ""
			if structInstance.BoundedBy != nil {
				val = structInstance.BoundedBy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractGMLType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractGMLType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractGMLType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractGMLType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractGeometricPrimitiveType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractGeometricPrimitiveType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractGeometricPrimitiveType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractGeometricPrimitiveType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractGeometryType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractGeometryType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractGeometryType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractGeometryType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractMetadataPropertyType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractMetadataPropertyType"
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
		col.Name = "Owns"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractMetadataPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractMetadataPropertyType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Owns"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.Owns}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractRingPropertyType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractRingPropertyType"
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
		col.Name = "Ring"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.PolygonPatchType) -> Exterior"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractRingPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractRingPropertyType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Ring"}
			val := ""
			if structInstance.Ring != nil {
				val = structInstance.Ring.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.PolygonPatchType) -> Exterior"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.PolygonPatchTypes {
				if src.Exterior == structInstance {
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

func updateStageSetTable_AbstractRingType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractRingType"
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

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractRingType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractRingType_GmlStage(probe)
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

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractSurfacePatchType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractSurfacePatchType"
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

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractSurfacePatchType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractSurfacePatchType_GmlStage(probe)
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

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractSurfaceType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractSurfaceType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractSurfaceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractSurfaceType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractTimeGeometricPrimitiveType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractTimeGeometricPrimitiveType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Frame"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractTimeGeometricPrimitiveType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractTimeGeometricPrimitiveType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Frame"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Frame)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractTimeObjectType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractTimeObjectType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractTimeObjectType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractTimeObjectType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractTimePrimitiveType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractTimePrimitiveType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractTimePrimitiveType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractTimePrimitiveType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AbstractTimeSliceType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AbstractTimeSliceType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractTimeSliceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AbstractTimeSliceType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ValidTime"}
			val := ""
			if structInstance.ValidTime != nil {
				val = structInstance.ValidTime.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_AngleType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.AngleType"
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
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcByCenterPointType) -> StartAngle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcByCenterPointType) -> EndAngle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AngleType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_AngleType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.ArcByCenterPointType) -> StartAngle"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.StartAngle == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.ArcByCenterPointType) -> EndAngle"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.EndAngle == structInstance {
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

func updateStageSetTable_ArcByCenterPointType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.ArcByCenterPointType"
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
		col.Name = "NumDerivativesAtStart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativesAtEnd"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativeInterior"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpolation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumArc"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Radius"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartAngle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndAngle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurveSegmentArrayPropertyType) -> ArcByCenterPoint"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.ArcByCenterPointType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ArcByCenterPointType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NumDerivativesAtStart"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtStart)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativesAtEnd"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtEnd)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativeInterior"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativeInterior)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpolation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpolation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumArc"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumArc)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Pos"}
			val := ""
			if structInstance.Pos != nil {
				val = structInstance.Pos.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PointProperty"}
			val := ""
			if structInstance.PointProperty != nil {
				val = structInstance.PointProperty.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PosList"}
			val := ""
			if structInstance.PosList != nil {
				val = structInstance.PosList.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Radius"}
			val := ""
			if structInstance.Radius != nil {
				val = structInstance.Radius.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartAngle"}
			val := ""
			if structInstance.StartAngle != nil {
				val = structInstance.StartAngle.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndAngle"}
			val := ""
			if structInstance.EndAngle != nil {
				val = structInstance.EndAngle.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurveSegmentArrayPropertyType) -> ArcByCenterPoint"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.ArcByCenterPoint {
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

func updateStageSetTable_ArcStringType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.ArcStringType"
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
		col.Name = "NumDerivativesAtStart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativesAtEnd"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativeInterior"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpolation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumArc"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurveSegmentArrayPropertyType) -> ArcString"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.ArcStringType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ArcStringType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NumDerivativesAtStart"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtStart)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativesAtEnd"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtEnd)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativeInterior"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativeInterior)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpolation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpolation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumArc"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumArc)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Pos"}
			var names []string
			for _, elem := range structInstance.Pos {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PointProperty"}
			var names []string
			for _, elem := range structInstance.PointProperty {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PosList"}
			val := ""
			if structInstance.PosList != nil {
				val = structInstance.PosList.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurveSegmentArrayPropertyType) -> ArcString"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.ArcString {
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

func updateStageSetTable_ArcType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.ArcType"
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
		col.Name = "NumDerivativesAtStart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativesAtEnd"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativeInterior"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpolation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumArc"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurveSegmentArrayPropertyType) -> Arc"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.ArcType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_ArcType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NumDerivativesAtStart"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtStart)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativesAtEnd"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtEnd)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativeInterior"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativeInterior)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpolation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpolation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumArc"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumArc)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Pos"}
			val := ""
			if structInstance.Pos != nil {
				val = structInstance.Pos.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PointProperty"}
			val := ""
			if structInstance.PointProperty != nil {
				val = structInstance.PointProperty.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PosList"}
			val := ""
			if structInstance.PosList != nil {
				val = structInstance.PosList.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurveSegmentArrayPropertyType) -> Arc"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.Arc {
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

func updateStageSetTable_BoundingShapeType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.BoundingShapeType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Envelope"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AIXMBasicMessage) -> BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AIXMBasicMessageType) -> BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightType) -> BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaType) -> BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureType) -> BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType) -> BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureType) -> BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType) -> BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMMessageType) -> BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractFeatureType) -> BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.DynamicFeatureType) -> BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.BoundingShapeType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_BoundingShapeType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Envelope"}
			val := ""
			if structInstance.Envelope != nil {
				val = structInstance.Envelope.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AIXMBasicMessage) -> BoundedBy"}
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessages {
				if src.BoundedBy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AIXMBasicMessageType) -> BoundedBy"}
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessageTypes {
				if src.BoundedBy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightType) -> BoundedBy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTypes {
				if src.BoundedBy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaType) -> BoundedBy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTypes {
				if src.BoundedBy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureType) -> BoundedBy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTypes {
				if src.BoundedBy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType) -> BoundedBy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureBaseTypes {
				if src.BoundedBy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureType) -> BoundedBy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureTypes {
				if src.BoundedBy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType) -> BoundedBy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageBaseTypes {
				if src.BoundedBy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMMessageType) -> BoundedBy"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageTypes {
				if src.BoundedBy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractFeatureType) -> BoundedBy"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractFeatureTypes {
				if src.BoundedBy == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.DynamicFeatureType) -> BoundedBy"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.DynamicFeatureTypes {
				if src.BoundedBy == structInstance {
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

func updateStageSetTable_CircleByCenterPointType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.CircleByCenterPointType"
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
		col.Name = "NumDerivativesAtStart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativesAtEnd"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativeInterior"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpolation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumArc"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Radius"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurveSegmentArrayPropertyType) -> CircleByCenterPoint"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CircleByCenterPointType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CircleByCenterPointType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NumDerivativesAtStart"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtStart)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativesAtEnd"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtEnd)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativeInterior"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativeInterior)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpolation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpolation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumArc"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumArc)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Pos"}
			val := ""
			if structInstance.Pos != nil {
				val = structInstance.Pos.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PointProperty"}
			val := ""
			if structInstance.PointProperty != nil {
				val = structInstance.PointProperty.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PosList"}
			val := ""
			if structInstance.PosList != nil {
				val = structInstance.PosList.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Radius"}
			val := ""
			if structInstance.Radius != nil {
				val = structInstance.Radius.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurveSegmentArrayPropertyType) -> CircleByCenterPoint"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.CircleByCenterPoint {
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

func updateStageSetTable_CircleType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.CircleType"
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
		col.Name = "NumDerivativesAtStart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativesAtEnd"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativeInterior"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpolation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumArc"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurveSegmentArrayPropertyType) -> Circle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CircleType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CircleType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NumDerivativesAtStart"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtStart)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativesAtEnd"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtEnd)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativeInterior"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativeInterior)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpolation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpolation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumArc"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumArc)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Pos"}
			var names []string
			for _, elem := range structInstance.Pos {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PointProperty"}
			var names []string
			for _, elem := range structInstance.PointProperty {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PosList"}
			val := ""
			if structInstance.PosList != nil {
				val = structInstance.PosList.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurveSegmentArrayPropertyType) -> Circle"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.Circle {
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

func updateStageSetTable_CodeType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.CodeType"
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
		col.Name = "CodeSpace"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AIXMBasicMessage) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AIXMBasicMessageType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AbstractPropertiesWithScheduleType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.CurveType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedCurveType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedPointType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedSurfaceType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementStatusType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LinguisticNoteType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.NoteType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.PointType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.SurfaceType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureLightingStatusType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMMessageType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractExtensionType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractCurveType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractFeatureType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractGMLType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractGeometricPrimitiveType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractGeometryType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractSurfaceType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractTimeGeometricPrimitiveType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractTimeObjectType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractTimePrimitiveType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractTimeSliceType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CompositeCurveType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurveType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.DynamicFeatureType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.OrientableCurveType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.PointType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.SurfaceType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.TimeInstantType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.TimePeriodType) -> NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CodeType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "CodeSpace"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.CodeSpace)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AIXMBasicMessage) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessages {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(models.AIXMBasicMessageType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessageTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.AbstractPropertiesWithScheduleType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractPropertiesWithScheduleTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.CurveType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurveTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedCurveType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedPointType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedSurfaceType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementStatusType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.LinguisticNoteType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LinguisticNoteTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.NoteType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.PointType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.SurfaceType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureLightingStatusType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureBaseTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageBaseTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMMessageType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractExtensionType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractExtensionTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.AbstractCurveType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractCurveTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.AbstractFeatureType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractFeatureTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.AbstractGMLType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGMLTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.AbstractGeometricPrimitiveType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGeometricPrimitiveTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.AbstractGeometryType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGeometryTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.AbstractSurfaceType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractSurfaceTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.AbstractTimeGeometricPrimitiveType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeGeometricPrimitiveTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.AbstractTimeObjectType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeObjectTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.AbstractTimePrimitiveType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimePrimitiveTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.AbstractTimeSliceType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeSliceTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.CompositeCurveType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CompositeCurveTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.CurveType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.DynamicFeatureType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.DynamicFeatureTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.OrientableCurveType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.OrientableCurveTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.PointType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.PointTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.SurfaceType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.SurfaceTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.TimeInstantType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimeInstantTypes {
				for _, target := range src.NameXSD {
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
			cell := &table_models.Cell{Name: "(gml.TimePeriodType) -> NameXSD"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePeriodTypes {
				for _, target := range src.NameXSD {
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

func updateStageSetTable_CodeWithAuthorityType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.CodeWithAuthorityType"
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
		col.Name = "CodeSpace"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AIXMBasicMessage) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AIXMBasicMessageType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AbstractPropertiesWithScheduleType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.CurveType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedCurveType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedPointType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedSurfaceType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementStatusType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LightElementType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.LinguisticNoteType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.NoteType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.PointType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.SurfaceType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.TimesheetType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureLightingStatusType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructurePartType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMMessageType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractExtensionType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractCurveType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractFeatureType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractGMLType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractGeometricPrimitiveType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractGeometryType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractSurfaceType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractTimeGeometricPrimitiveType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractTimeObjectType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractTimePrimitiveType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractTimeSliceType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CompositeCurveType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurveType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.DynamicFeatureType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.OrientableCurveType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.PointType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.SurfaceType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.TimeInstantType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.TimePeriodType) -> Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CodeWithAuthorityType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CodeWithAuthorityType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "CodeSpace"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.CodeSpace)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AIXMBasicMessage) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessages {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AIXMBasicMessageType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessageTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AbstractPropertiesWithScheduleType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractPropertiesWithScheduleTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.CurveType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurveTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedCurveType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedPointType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedSurfaceType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementStatusType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LightElementType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.LinguisticNoteType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LinguisticNoteTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.NoteType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.PointType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.SurfaceType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.TimesheetType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureLightingStatusType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructurePartType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureBaseTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMFeatureType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageBaseTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMMessageType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractExtensionType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractExtensionTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractCurveType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractCurveTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractFeatureType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractFeatureTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractGMLType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGMLTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractGeometricPrimitiveType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGeometricPrimitiveTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractGeometryType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGeometryTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractSurfaceType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractSurfaceTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractTimeGeometricPrimitiveType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeGeometricPrimitiveTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractTimeObjectType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeObjectTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractTimePrimitiveType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimePrimitiveTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractTimeSliceType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeSliceTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CompositeCurveType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CompositeCurveTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurveType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.DynamicFeatureType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.DynamicFeatureTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.OrientableCurveType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.OrientableCurveTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.PointType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.PointTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.SurfaceType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.SurfaceTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.TimeInstantType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimeInstantTypes {
				if src.Identifier == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.TimePeriodType) -> Identifier"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePeriodTypes {
				if src.Identifier == structInstance {
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

func updateStageSetTable_CompositeCurveType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.CompositeCurveType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AggregationType"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CurveMember"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurvePropertyType) -> CompositeCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CompositeCurveType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CompositeCurveType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AggregationType"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.AggregationType)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CurveMember"}
			var names []string
			for _, elem := range structInstance.CurveMember {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurvePropertyType) -> CompositeCurve"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurvePropertyTypes {
				if src.CompositeCurve == structInstance {
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

func updateStageSetTable_CurvePropertyType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.CurvePropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Curve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "OrientableCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CompositeCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CompositeCurveType) -> CurveMember"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.OrientableCurveType) -> BaseCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.RingType) -> CurveMember"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CurvePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CurvePropertyType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Curve"}
			val := ""
			if structInstance.Curve != nil {
				val = structInstance.Curve.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "OrientableCurve"}
			val := ""
			if structInstance.OrientableCurve != nil {
				val = structInstance.OrientableCurve.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CompositeCurve"}
			val := ""
			if structInstance.CompositeCurve != nil {
				val = structInstance.CompositeCurve.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CompositeCurveType) -> CurveMember"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CompositeCurveTypes {
				for _, target := range src.CurveMember {
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
			cell := &table_models.Cell{Name: "(gml.OrientableCurveType) -> BaseCurve"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.OrientableCurveTypes {
				if src.BaseCurve == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.RingType) -> CurveMember"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.RingTypes {
				for _, target := range src.CurveMember {
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

func updateStageSetTable_CurveSegmentArrayPropertyType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.CurveSegmentArrayPropertyType"
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
		col.Name = "LineStringSegment"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ArcString"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arc"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Circle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ArcByCenterPoint"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CircleByCenterPoint"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "GeodesicString"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Geodesic"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.CurveType) -> Segments"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedCurveType) -> Segments"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurveType) -> Segments"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CurveSegmentArrayPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CurveSegmentArrayPropertyType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "LineStringSegment"}
			var names []string
			for _, elem := range structInstance.LineStringSegment {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ArcString"}
			var names []string
			for _, elem := range structInstance.ArcString {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arc"}
			var names []string
			for _, elem := range structInstance.Arc {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Circle"}
			var names []string
			for _, elem := range structInstance.Circle {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ArcByCenterPoint"}
			var names []string
			for _, elem := range structInstance.ArcByCenterPoint {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CircleByCenterPoint"}
			var names []string
			for _, elem := range structInstance.CircleByCenterPoint {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "GeodesicString"}
			var names []string
			for _, elem := range structInstance.GeodesicString {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Geodesic"}
			var names []string
			for _, elem := range structInstance.Geodesic {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.CurveType) -> Segments"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurveTypes {
				if src.Segments == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedCurveType) -> Segments"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.Segments == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurveType) -> Segments"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveTypes {
				if src.Segments == structInstance {
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

func updateStageSetTable_CurveType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.CurveType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Segments"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurvePropertyType) -> Curve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CurveType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_CurveType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Segments"}
			val := ""
			if structInstance.Segments != nil {
				val = structInstance.Segments.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurvePropertyType) -> Curve"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurvePropertyTypes {
				if src.Curve == structInstance {
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

func updateStageSetTable_DirectPositionListType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.DirectPositionListType"
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
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcByCenterPointType) -> PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcStringType) -> PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcType) -> PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CircleByCenterPointType) -> PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CircleType) -> PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.GeodesicStringType) -> PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.GeodesicType) -> PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.LineStringSegmentType) -> PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.DirectPositionListType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_DirectPositionListType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.ArcByCenterPointType) -> PosList"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.PosList == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.ArcStringType) -> PosList"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcStringTypes {
				if src.PosList == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.ArcType) -> PosList"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcTypes {
				if src.PosList == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CircleByCenterPointType) -> PosList"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleByCenterPointTypes {
				if src.PosList == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CircleType) -> PosList"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleTypes {
				if src.PosList == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.GeodesicStringType) -> PosList"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicStringTypes {
				if src.PosList == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.GeodesicType) -> PosList"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicTypes {
				if src.PosList == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.LineStringSegmentType) -> PosList"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.LineStringSegmentTypes {
				if src.PosList == structInstance {
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

func updateStageSetTable_DirectPositionType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.DirectPositionType"
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
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedPointType) -> Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.PointType) -> Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcByCenterPointType) -> Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcStringType) -> Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcType) -> Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CircleByCenterPointType) -> Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CircleType) -> Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.EnvelopeType) -> LowerCorner"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.EnvelopeType) -> UpperCorner"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.GeodesicStringType) -> Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.GeodesicType) -> Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.LineStringSegmentType) -> Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.PointType) -> Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.DirectPositionType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_DirectPositionType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedPointType) -> Pos"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.Pos == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.PointType) -> Pos"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointTypes {
				if src.Pos == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.ArcByCenterPointType) -> Pos"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.Pos == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.ArcStringType) -> Pos"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcStringTypes {
				for _, target := range src.Pos {
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
			cell := &table_models.Cell{Name: "(gml.ArcType) -> Pos"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcTypes {
				if src.Pos == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CircleByCenterPointType) -> Pos"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleByCenterPointTypes {
				if src.Pos == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CircleType) -> Pos"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleTypes {
				for _, target := range src.Pos {
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
			cell := &table_models.Cell{Name: "(gml.EnvelopeType) -> LowerCorner"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.EnvelopeTypes {
				if src.LowerCorner == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.EnvelopeType) -> UpperCorner"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.EnvelopeTypes {
				if src.UpperCorner == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.GeodesicStringType) -> Pos"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicStringTypes {
				if src.Pos == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.GeodesicType) -> Pos"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicTypes {
				if src.Pos == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.LineStringSegmentType) -> Pos"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.LineStringSegmentTypes {
				for _, target := range src.Pos {
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
			cell := &table_models.Cell{Name: "(gml.PointType) -> Pos"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.PointTypes {
				if src.Pos == structInstance {
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

func updateStageSetTable_DynamicFeatureType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.DynamicFeatureType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BoundedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.DynamicFeatureType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_DynamicFeatureType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BoundedBy"}
			val := ""
			if structInstance.BoundedBy != nil {
				val = structInstance.BoundedBy.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_EnvelopeType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.EnvelopeType"
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
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LowerCorner"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "UpperCorner"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.BoundingShapeType) -> Envelope"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.EnvelopeType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_EnvelopeType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LowerCorner"}
			val := ""
			if structInstance.LowerCorner != nil {
				val = structInstance.LowerCorner.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "UpperCorner"}
			val := ""
			if structInstance.UpperCorner != nil {
				val = structInstance.UpperCorner.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.BoundingShapeType) -> Envelope"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.BoundingShapeTypes {
				if src.Envelope == structInstance {
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

func updateStageSetTable_GeodesicStringType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.GeodesicStringType"
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
		col.Name = "NumDerivativesAtStart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativesAtEnd"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativeInterior"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpolation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurveSegmentArrayPropertyType) -> GeodesicString"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.GeodesicStringType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_GeodesicStringType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NumDerivativesAtStart"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtStart)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativesAtEnd"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtEnd)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativeInterior"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativeInterior)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpolation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpolation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PosList"}
			val := ""
			if structInstance.PosList != nil {
				val = structInstance.PosList.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Pos"}
			val := ""
			if structInstance.Pos != nil {
				val = structInstance.Pos.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PointProperty"}
			val := ""
			if structInstance.PointProperty != nil {
				val = structInstance.PointProperty.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurveSegmentArrayPropertyType) -> GeodesicString"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.GeodesicString {
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

func updateStageSetTable_GeodesicType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.GeodesicType"
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
		col.Name = "NumDerivativesAtStart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativesAtEnd"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativeInterior"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpolation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurveSegmentArrayPropertyType) -> Geodesic"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.GeodesicType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_GeodesicType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NumDerivativesAtStart"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtStart)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativesAtEnd"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtEnd)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativeInterior"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativeInterior)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpolation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpolation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PosList"}
			val := ""
			if structInstance.PosList != nil {
				val = structInstance.PosList.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Pos"}
			val := ""
			if structInstance.Pos != nil {
				val = structInstance.Pos.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PointProperty"}
			val := ""
			if structInstance.PointProperty != nil {
				val = structInstance.PointProperty.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurveSegmentArrayPropertyType) -> Geodesic"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.Geodesic {
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

func updateStageSetTable_LengthType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.LengthType"
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
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcByCenterPointType) -> Radius"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CircleByCenterPointType) -> Radius"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.LengthType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_LengthType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.ArcByCenterPointType) -> Radius"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.Radius == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CircleByCenterPointType) -> Radius"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleByCenterPointTypes {
				if src.Radius == structInstance {
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

func updateStageSetTable_LineStringSegmentType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.LineStringSegmentType"
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
		col.Name = "NumDerivativesAtStart"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativesAtEnd"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NumDerivativeInterior"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Interpolation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PosList"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurveSegmentArrayPropertyType) -> LineStringSegment"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.LineStringSegmentType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_LineStringSegmentType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NumDerivativesAtStart"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtStart)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativesAtEnd"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativesAtEnd)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NumDerivativeInterior"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumDerivativeInterior)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Interpolation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpolation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Pos"}
			var names []string
			for _, elem := range structInstance.Pos {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PointProperty"}
			var names []string
			for _, elem := range structInstance.PointProperty {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PosList"}
			val := ""
			if structInstance.PosList != nil {
				val = structInstance.PosList.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurveSegmentArrayPropertyType) -> LineStringSegment"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.LineStringSegment {
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

func updateStageSetTable_MeasureType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.MeasureType"
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
		col.Name = "Uom"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.MeasureType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_MeasureType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Uom"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Uom)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_OrientableCurveType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.OrientableCurveType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Orientation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BaseCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CurvePropertyType) -> OrientableCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.OrientableCurveType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_OrientableCurveType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Orientation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Orientation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BaseCurve"}
			val := ""
			if structInstance.BaseCurve != nil {
				val = structInstance.BaseCurve.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CurvePropertyType) -> OrientableCurve"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurvePropertyTypes {
				if src.OrientableCurve == structInstance {
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

func updateStageSetTable_PointPropertyType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.PointPropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Point"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcByCenterPointType) -> PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcStringType) -> PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.ArcType) -> PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CircleByCenterPointType) -> PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.CircleType) -> PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.GeodesicStringType) -> PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.GeodesicType) -> PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.LineStringSegmentType) -> PointProperty"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.PointPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_PointPropertyType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Point"}
			val := ""
			if structInstance.Point != nil {
				val = structInstance.Point.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.ArcByCenterPointType) -> PointProperty"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.PointProperty == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.ArcStringType) -> PointProperty"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcStringTypes {
				for _, target := range src.PointProperty {
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
			cell := &table_models.Cell{Name: "(gml.ArcType) -> PointProperty"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcTypes {
				if src.PointProperty == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CircleByCenterPointType) -> PointProperty"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleByCenterPointTypes {
				if src.PointProperty == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.CircleType) -> PointProperty"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleTypes {
				for _, target := range src.PointProperty {
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
			cell := &table_models.Cell{Name: "(gml.GeodesicStringType) -> PointProperty"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicStringTypes {
				if src.PointProperty == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.GeodesicType) -> PointProperty"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicTypes {
				if src.PointProperty == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.LineStringSegmentType) -> PointProperty"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.LineStringSegmentTypes {
				for _, target := range src.PointProperty {
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

func updateStageSetTable_PointType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.PointType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Pos"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.PointPropertyType) -> Point"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.PointType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_PointType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Pos"}
			val := ""
			if structInstance.Pos != nil {
				val = structInstance.Pos.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.PointPropertyType) -> Point"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.PointPropertyTypes {
				if src.Point == structInstance {
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

func updateStageSetTable_PolygonPatchType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.PolygonPatchType"
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
		col.Name = "Interpolation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Exterior"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.SurfacePatchArrayPropertyType) -> PolygonPatch"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.PolygonPatchType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_PolygonPatchType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Interpolation"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Interpolation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Exterior"}
			val := ""
			if structInstance.Exterior != nil {
				val = structInstance.Exterior.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.SurfacePatchArrayPropertyType) -> PolygonPatch"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.SurfacePatchArrayPropertyTypes {
				for _, target := range src.PolygonPatch {
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

func updateStageSetTable_RingType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.RingType"
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
		col.Name = "AggregationType"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CurveMember"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractRingPropertyType) -> Ring"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.RingType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_RingType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "AggregationType"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.AggregationType)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CurveMember"}
			var names []string
			for _, elem := range structInstance.CurveMember {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractRingPropertyType) -> Ring"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractRingPropertyTypes {
				if src.Ring == structInstance {
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

func updateStageSetTable_SurfacePatchArrayPropertyType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.SurfacePatchArrayPropertyType"
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
		col.Name = "PolygonPatch"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ElevatedSurfaceType) -> Patches"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.SurfaceType) -> Patches"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.SurfaceType) -> Patches"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.SurfacePatchArrayPropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_SurfacePatchArrayPropertyType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "PolygonPatch"}
			var names []string
			for _, elem := range structInstance.PolygonPatch {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ElevatedSurfaceType) -> Patches"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.Patches == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.SurfaceType) -> Patches"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes {
				if src.Patches == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.SurfaceType) -> Patches"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.SurfaceTypes {
				if src.Patches == structInstance {
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

func updateStageSetTable_SurfaceType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.SurfaceType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SrsName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Patches"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.SurfaceType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_SurfaceType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SrsName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SrsName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Patches"}
			val := ""
			if structInstance.Patches != nil {
				val = structInstance.Patches.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TimeInstantType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.TimeInstantType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Frame"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimePosition"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.TimePrimitivePropertyType) -> TimeInstant"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.TimeInstantType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_TimeInstantType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Frame"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Frame)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimePosition"}
			val := ""
			if structInstance.TimePosition != nil {
				val = structInstance.TimePosition.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.TimePrimitivePropertyType) -> TimeInstant"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePrimitivePropertyTypes {
				if src.TimeInstant == structInstance {
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

func updateStageSetTable_TimePeriodType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.TimePeriodType"
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
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Description"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Identifier"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NameXSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Frame"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BeginPosition"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndPosition"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.TimePrimitivePropertyType) -> TimePeriod"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.TimePeriodType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_TimePeriodType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Description"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Description)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Identifier"}
			val := ""
			if structInstance.Identifier != nil {
				val = structInstance.Identifier.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NameXSD"}
			var names []string
			for _, elem := range structInstance.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Frame"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Frame)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BeginPosition"}
			val := ""
			if structInstance.BeginPosition != nil {
				val = structInstance.BeginPosition.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndPosition"}
			val := ""
			if structInstance.EndPosition != nil {
				val = structInstance.EndPosition.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.TimePrimitivePropertyType) -> TimePeriod"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePrimitivePropertyTypes {
				if src.TimePeriod == structInstance {
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

func updateStageSetTable_TimePositionType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.TimePositionType"
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
		col.Name = "Frame"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IndeterminatePosition"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EnclosedText"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.TimeInstantType) -> TimePosition"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.TimePeriodType) -> BeginPosition"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.TimePeriodType) -> EndPosition"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.TimePositionType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_TimePositionType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "Frame"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Frame)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IndeterminatePosition"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.IndeterminatePosition)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EnclosedText"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.EnclosedText)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.TimeInstantType) -> TimePosition"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimeInstantTypes {
				if src.TimePosition == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.TimePeriodType) -> BeginPosition"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePeriodTypes {
				if src.BeginPosition == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.TimePeriodType) -> EndPosition"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePeriodTypes {
				if src.EndPosition == structInstance {
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

func updateStageSetTable_TimePrimitivePropertyType_GmlStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "gml.TimePrimitivePropertyType"
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
		col.Name = "NilReason"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Type_XSD"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Href"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Role"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Arcrole"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Title"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Show"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Actuate"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimeInstant"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TimePeriod"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> FeatureLifetime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.ObstacleAreaTimeSliceType) -> FeatureLifetime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmfeatures.VerticalStructureTimeSliceType) -> FeatureLifetime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType) -> ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType) -> ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType) -> FeatureLifetime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(gml.AbstractTimeSliceType) -> ValidTime"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.TimePrimitivePropertyType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GmlStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.GmlStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_TimePrimitivePropertyType_GmlStage(probe)
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
			cell := &table_models.Cell{Name: "NilReason"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.NilReason)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Type_XSD"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Type_XSD)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Href"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Href)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Role"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Role)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Arcrole"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Arcrole)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Title"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Title)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Show"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Show)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Actuate"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Actuate)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimeInstant"}
			val := ""
			if structInstance.TimeInstant != nil {
				val = structInstance.TimeInstant.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TimePeriod"}
			val := ""
			if structInstance.TimePeriod != nil {
				val = structInstance.TimePeriod.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> ValidTime"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.ValidTime == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.AeronauticalGroundLightTimeSliceType) -> FeatureLifetime"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.FeatureLifetime == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> ValidTime"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.ValidTime == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.ObstacleAreaTimeSliceType) -> FeatureLifetime"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.FeatureLifetime == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> ValidTime"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.ValidTime == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmfeatures.VerticalStructureTimeSliceType) -> FeatureLifetime"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.FeatureLifetime == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType) -> ValidTime"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceBaseTypes {
				if src.ValidTime == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType) -> ValidTime"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes {
				if src.ValidTime == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType) -> FeatureLifetime"}
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes {
				if src.FeatureLifetime == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(gml.AbstractTimeSliceType) -> ValidTime"}
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeSliceTypes {
				if src.ValidTime == structInstance {
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

func updateStageSetTable_Xlink_SchemaElement_XlinkStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "xlink.Xlink_SchemaElement"
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

	instances := probe.stageOfInterest.XlinkStage.GetInstancesByOrder[*xlink.Xlink_SchemaElement]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.XlinkStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageOfInterest.XlinkStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_Xlink_SchemaElement_XlinkStage(probe)
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

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

