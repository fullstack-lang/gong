package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/xlsx/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	const __write__local_time = "2025-06-18 06:50:19.700498 CEST"
	const __write__utc_time__ = "2025-06-18 04:50:19.700498 UTC"

	const __commitId__ = "0000000043"

	// Declaration of instances to stage

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_NbRows := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_RowIndex := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000003_NbSheets := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_49_07Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_XLSheet := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_XLRow := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_XLFile := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_XLCell := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_SheetCells := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_Rows := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_Cells := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000003_Sheets := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.XLSheet{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `XLSheet`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_NbRows.Name = `NbRows`
	__AttributeShape__000001_NbRows.IdentifierMeta = ref_models.XLSheet{}.NbRows
	__AttributeShape__000001_NbRows.FieldTypeAsString = ``
	__AttributeShape__000001_NbRows.Structname = `XLSheet`
	__AttributeShape__000001_NbRows.Fieldtypename = `int`

	__AttributeShape__000002_RowIndex.Name = `RowIndex`
	__AttributeShape__000002_RowIndex.IdentifierMeta = ref_models.XLRow{}.RowIndex
	__AttributeShape__000002_RowIndex.FieldTypeAsString = ``
	__AttributeShape__000002_RowIndex.Structname = `XLRow`
	__AttributeShape__000002_RowIndex.Fieldtypename = `int`

	__AttributeShape__000003_NbSheets.Name = `NbSheets`
	__AttributeShape__000003_NbSheets.IdentifierMeta = ref_models.XLFile{}.NbSheets
	__AttributeShape__000003_NbSheets.FieldTypeAsString = ``
	__AttributeShape__000003_NbSheets.Structname = `XLFile`
	__AttributeShape__000003_NbSheets.Fieldtypename = `int`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[false,false,true,true,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_49_07Z.Name = `Diagram Package created the 2025-06-18T04:49:07Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_49_07Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_49_07Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_49_07Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_XLSheet.Name = `Default-XLSheet`
	__GongStructShape__000000_Default_XLSheet.X = 466.000000
	__GongStructShape__000000_Default_XLSheet.Y = 185.000000
	__GongStructShape__000000_Default_XLSheet.IdentifierMeta = ref_models.XLSheet{}
	__GongStructShape__000000_Default_XLSheet.ShowNbInstances = false
	__GongStructShape__000000_Default_XLSheet.NbInstances = 0
	__GongStructShape__000000_Default_XLSheet.Width = 240.000000
	__GongStructShape__000000_Default_XLSheet.Height = 103.000000
	__GongStructShape__000000_Default_XLSheet.IsSelected = false

	__GongStructShape__000001_Default_XLRow.Name = `Default-XLRow`
	__GongStructShape__000001_Default_XLRow.X = 875.000000
	__GongStructShape__000001_Default_XLRow.Y = 199.000000
	__GongStructShape__000001_Default_XLRow.IdentifierMeta = ref_models.XLRow{}
	__GongStructShape__000001_Default_XLRow.ShowNbInstances = false
	__GongStructShape__000001_Default_XLRow.NbInstances = 0
	__GongStructShape__000001_Default_XLRow.Width = 240.000000
	__GongStructShape__000001_Default_XLRow.Height = 83.000000
	__GongStructShape__000001_Default_XLRow.IsSelected = false

	__GongStructShape__000002_Default_XLFile.Name = `Default-XLFile`
	__GongStructShape__000002_Default_XLFile.X = 64.000000
	__GongStructShape__000002_Default_XLFile.Y = 196.000000
	__GongStructShape__000002_Default_XLFile.IdentifierMeta = ref_models.XLFile{}
	__GongStructShape__000002_Default_XLFile.ShowNbInstances = false
	__GongStructShape__000002_Default_XLFile.NbInstances = 0
	__GongStructShape__000002_Default_XLFile.Width = 240.000000
	__GongStructShape__000002_Default_XLFile.Height = 83.000000
	__GongStructShape__000002_Default_XLFile.IsSelected = false

	__GongStructShape__000003_Default_XLCell.Name = `Default-XLCell`
	__GongStructShape__000003_Default_XLCell.X = 1244.000000
	__GongStructShape__000003_Default_XLCell.Y = 91.000000
	__GongStructShape__000003_Default_XLCell.IdentifierMeta = ref_models.XLCell{}
	__GongStructShape__000003_Default_XLCell.ShowNbInstances = false
	__GongStructShape__000003_Default_XLCell.NbInstances = 0
	__GongStructShape__000003_Default_XLCell.Width = 240.000000
	__GongStructShape__000003_Default_XLCell.Height = 120.000000
	__GongStructShape__000003_Default_XLCell.IsSelected = false

	__LinkShape__000000_SheetCells.Name = `SheetCells`
	__LinkShape__000000_SheetCells.IdentifierMeta = ref_models.XLSheet{}.SheetCells
	__LinkShape__000000_SheetCells.FieldTypeIdentifierMeta = ref_models.XLCell{}
	__LinkShape__000000_SheetCells.FieldOffsetX = 0.000000
	__LinkShape__000000_SheetCells.FieldOffsetY = 0.000000
	__LinkShape__000000_SheetCells.TargetMultiplicity = models.MANY
	__LinkShape__000000_SheetCells.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_SheetCells.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_SheetCells.SourceMultiplicity = models.MANY
	__LinkShape__000000_SheetCells.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_SheetCells.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_SheetCells.X = 540.000000
	__LinkShape__000000_SheetCells.Y = 149.000000
	__LinkShape__000000_SheetCells.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_SheetCells.StartRatio = 0.500000
	__LinkShape__000000_SheetCells.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_SheetCells.EndRatio = 0.000000
	__LinkShape__000000_SheetCells.CornerOffsetRatio = 1.380000

	__LinkShape__000001_Rows.Name = `Rows`
	__LinkShape__000001_Rows.IdentifierMeta = ref_models.XLSheet{}.Rows
	__LinkShape__000001_Rows.FieldTypeIdentifierMeta = ref_models.XLRow{}
	__LinkShape__000001_Rows.FieldOffsetX = 0.000000
	__LinkShape__000001_Rows.FieldOffsetY = 0.000000
	__LinkShape__000001_Rows.TargetMultiplicity = models.MANY
	__LinkShape__000001_Rows.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Rows.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Rows.SourceMultiplicity = models.MANY
	__LinkShape__000001_Rows.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Rows.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Rows.X = 432.000000
	__LinkShape__000001_Rows.Y = 93.000000
	__LinkShape__000001_Rows.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Rows.StartRatio = 0.500000
	__LinkShape__000001_Rows.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Rows.EndRatio = 0.500000
	__LinkShape__000001_Rows.CornerOffsetRatio = 1.380000

	__LinkShape__000002_Cells.Name = `Cells`
	__LinkShape__000002_Cells.IdentifierMeta = ref_models.XLRow{}.Cells
	__LinkShape__000002_Cells.FieldTypeIdentifierMeta = ref_models.XLCell{}
	__LinkShape__000002_Cells.FieldOffsetX = 0.000000
	__LinkShape__000002_Cells.FieldOffsetY = 0.000000
	__LinkShape__000002_Cells.TargetMultiplicity = models.MANY
	__LinkShape__000002_Cells.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Cells.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Cells.SourceMultiplicity = models.MANY
	__LinkShape__000002_Cells.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Cells.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Cells.X = 565.000000
	__LinkShape__000002_Cells.Y = 126.500000
	__LinkShape__000002_Cells.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Cells.StartRatio = 0.464943
	__LinkShape__000002_Cells.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Cells.EndRatio = 0.500000
	__LinkShape__000002_Cells.CornerOffsetRatio = 1.380000

	__LinkShape__000003_Sheets.Name = `Sheets`
	__LinkShape__000003_Sheets.IdentifierMeta = ref_models.XLFile{}.Sheets
	__LinkShape__000003_Sheets.FieldTypeIdentifierMeta = ref_models.XLSheet{}
	__LinkShape__000003_Sheets.FieldOffsetX = 0.000000
	__LinkShape__000003_Sheets.FieldOffsetY = 0.000000
	__LinkShape__000003_Sheets.TargetMultiplicity = models.MANY
	__LinkShape__000003_Sheets.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000003_Sheets.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000003_Sheets.SourceMultiplicity = models.MANY
	__LinkShape__000003_Sheets.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000003_Sheets.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000003_Sheets.X = 394.500000
	__LinkShape__000003_Sheets.Y = 98.000000
	__LinkShape__000003_Sheets.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_Sheets.StartRatio = 0.500000
	__LinkShape__000003_Sheets.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_Sheets.EndRatio = 0.500000
	__LinkShape__000003_Sheets.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_XLSheet)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_XLRow)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_XLFile)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_XLCell)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_49_07Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_49_07Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_18T04_49_07Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_XLSheet.AttributeShapes = append(__GongStructShape__000000_Default_XLSheet.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_XLSheet.AttributeShapes = append(__GongStructShape__000000_Default_XLSheet.AttributeShapes, __AttributeShape__000001_NbRows)
	__GongStructShape__000000_Default_XLSheet.LinkShapes = append(__GongStructShape__000000_Default_XLSheet.LinkShapes, __LinkShape__000000_SheetCells)
	__GongStructShape__000000_Default_XLSheet.LinkShapes = append(__GongStructShape__000000_Default_XLSheet.LinkShapes, __LinkShape__000001_Rows)
	__GongStructShape__000001_Default_XLRow.AttributeShapes = append(__GongStructShape__000001_Default_XLRow.AttributeShapes, __AttributeShape__000002_RowIndex)
	__GongStructShape__000001_Default_XLRow.LinkShapes = append(__GongStructShape__000001_Default_XLRow.LinkShapes, __LinkShape__000002_Cells)
	__GongStructShape__000002_Default_XLFile.AttributeShapes = append(__GongStructShape__000002_Default_XLFile.AttributeShapes, __AttributeShape__000003_NbSheets)
	__GongStructShape__000002_Default_XLFile.LinkShapes = append(__GongStructShape__000002_Default_XLFile.LinkShapes, __LinkShape__000003_Sheets)
	// setup of LinkShape instances pointers
}

