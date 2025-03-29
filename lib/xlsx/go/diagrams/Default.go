package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/doc/go/models"
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

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__Field__000000_Name := (&models.Field{}).Stage(stage)
	__Field__000001_RowIndex := (&models.Field{}).Stage(stage)
	__Field__000002_Name := (&models.Field{}).Stage(stage)
	__Field__000003_X := (&models.Field{}).Stage(stage)
	__Field__000004_Y := (&models.Field{}).Stage(stage)

	__GongStructShape__000000_Default_XLSheet := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_XLRow := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_XLFile := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_XLCell := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000004_Default_DisplaySelection := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_Rows := (&models.Link{}).Stage(stage)
	__Link__000001_SheetCells := (&models.Link{}).Stage(stage)
	__Link__000002_Cells := (&models.Link{}).Stage(stage)
	__Link__000003_Sheets := (&models.Link{}).Stage(stage)
	__Link__000004_XLFile := (&models.Link{}).Stage(stage)
	__Link__000005_XLSheet := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_Default_XLSheet := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_XLRow := (&models.Position{}).Stage(stage)
	__Position__000002_Pos_Default_XLFile := (&models.Position{}).Stage(stage)
	__Position__000003_Pos_Default_XLCell := (&models.Position{}).Stage(stage)
	__Position__000004_Pos_Default_DisplaySelection := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_XLSheet_and_Default_XLRow := (&models.Vertice{}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_XLSheet_and_Default_XLCell := (&models.Vertice{}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_XLRow_and_Default_XLCell := (&models.Vertice{}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_XLFile_and_Default_XLSheet := (&models.Vertice{}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_DisplaySelection_and_Default_XLFile := (&models.Vertice{}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_DisplaySelection_and_Default_XLSheet := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Name.Name = `Name`

	//gong:ident [ref_models.XLSheet.Name] comment added to overcome the problem with the comment map association
	__Field__000000_Name.Identifier = `ref_models.XLSheet.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `XLSheet`
	__Field__000000_Name.Fieldtypename = `string`

	__Field__000001_RowIndex.Name = `RowIndex`

	//gong:ident [ref_models.XLRow.RowIndex] comment added to overcome the problem with the comment map association
	__Field__000001_RowIndex.Identifier = `ref_models.XLRow.RowIndex`
	__Field__000001_RowIndex.FieldTypeAsString = ``
	__Field__000001_RowIndex.Structname = `XLRow`
	__Field__000001_RowIndex.Fieldtypename = `int`

	__Field__000002_Name.Name = `Name`

	//gong:ident [ref_models.XLCell.Name] comment added to overcome the problem with the comment map association
	__Field__000002_Name.Identifier = `ref_models.XLCell.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `XLCell`
	__Field__000002_Name.Fieldtypename = `string`

	__Field__000003_X.Name = `X`

	//gong:ident [ref_models.XLCell.X] comment added to overcome the problem with the comment map association
	__Field__000003_X.Identifier = `ref_models.XLCell.X`
	__Field__000003_X.FieldTypeAsString = ``
	__Field__000003_X.Structname = `XLCell`
	__Field__000003_X.Fieldtypename = `int`

	__Field__000004_Y.Name = `Y`

	//gong:ident [ref_models.XLCell.Y] comment added to overcome the problem with the comment map association
	__Field__000004_Y.Identifier = `ref_models.XLCell.Y`
	__Field__000004_Y.FieldTypeAsString = ``
	__Field__000004_Y.Structname = `XLCell`
	__Field__000004_Y.Fieldtypename = `int`

	__GongStructShape__000000_Default_XLSheet.Name = `Default-XLSheet`

	//gong:ident [ref_models.XLSheet] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_XLSheet.Identifier = `ref_models.XLSheet`
	__GongStructShape__000000_Default_XLSheet.ShowNbInstances = false
	__GongStructShape__000000_Default_XLSheet.NbInstances = 0
	__GongStructShape__000000_Default_XLSheet.Width = 240.000000
	__GongStructShape__000000_Default_XLSheet.Height = 78.000000
	__GongStructShape__000000_Default_XLSheet.IsSelected = false

	__GongStructShape__000001_Default_XLRow.Name = `Default-XLRow`

	//gong:ident [ref_models.XLRow] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_XLRow.Identifier = `ref_models.XLRow`
	__GongStructShape__000001_Default_XLRow.ShowNbInstances = false
	__GongStructShape__000001_Default_XLRow.NbInstances = 0
	__GongStructShape__000001_Default_XLRow.Width = 240.000000
	__GongStructShape__000001_Default_XLRow.Height = 78.000000
	__GongStructShape__000001_Default_XLRow.IsSelected = false

	__GongStructShape__000002_Default_XLFile.Name = `Default-XLFile`

	//gong:ident [ref_models.XLFile] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_XLFile.Identifier = `ref_models.XLFile`
	__GongStructShape__000002_Default_XLFile.ShowNbInstances = false
	__GongStructShape__000002_Default_XLFile.NbInstances = 0
	__GongStructShape__000002_Default_XLFile.Width = 240.000000
	__GongStructShape__000002_Default_XLFile.Height = 63.000000
	__GongStructShape__000002_Default_XLFile.IsSelected = false

	__GongStructShape__000003_Default_XLCell.Name = `Default-XLCell`

	//gong:ident [ref_models.XLCell] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_XLCell.Identifier = `ref_models.XLCell`
	__GongStructShape__000003_Default_XLCell.ShowNbInstances = false
	__GongStructShape__000003_Default_XLCell.NbInstances = 0
	__GongStructShape__000003_Default_XLCell.Width = 240.000000
	__GongStructShape__000003_Default_XLCell.Height = 108.000000
	__GongStructShape__000003_Default_XLCell.IsSelected = false

	__GongStructShape__000004_Default_DisplaySelection.Name = `Default-DisplaySelection`

	//gong:ident [ref_models.DisplaySelection] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_DisplaySelection.Identifier = `ref_models.DisplaySelection`
	__GongStructShape__000004_Default_DisplaySelection.ShowNbInstances = false
	__GongStructShape__000004_Default_DisplaySelection.NbInstances = 0
	__GongStructShape__000004_Default_DisplaySelection.Width = 240.000000
	__GongStructShape__000004_Default_DisplaySelection.Height = 63.000000
	__GongStructShape__000004_Default_DisplaySelection.IsSelected = false

	__Link__000000_Rows.Name = `Rows`

	//gong:ident [ref_models.XLSheet.Rows] comment added to overcome the problem with the comment map association
	__Link__000000_Rows.Identifier = `ref_models.XLSheet.Rows`

	//gong:ident [ref_models.XLRow] comment added to overcome the problem with the comment map association
	__Link__000000_Rows.Fieldtypename = `ref_models.XLRow`
	__Link__000000_Rows.FieldOffsetX = 0.000000
	__Link__000000_Rows.FieldOffsetY = 0.000000
	__Link__000000_Rows.TargetMultiplicity = models.MANY
	__Link__000000_Rows.TargetMultiplicityOffsetX = 0.000000
	__Link__000000_Rows.TargetMultiplicityOffsetY = 0.000000
	__Link__000000_Rows.SourceMultiplicity = models.MANY
	__Link__000000_Rows.SourceMultiplicityOffsetX = 0.000000
	__Link__000000_Rows.SourceMultiplicityOffsetY = 0.000000
	__Link__000000_Rows.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Rows.StartRatio = 0.500000
	__Link__000000_Rows.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Rows.EndRatio = 0.500000
	__Link__000000_Rows.CornerOffsetRatio = 1.221688

	__Link__000001_SheetCells.Name = `SheetCells`

	//gong:ident [ref_models.XLSheet.SheetCells] comment added to overcome the problem with the comment map association
	__Link__000001_SheetCells.Identifier = `ref_models.XLSheet.SheetCells`

	//gong:ident [ref_models.XLCell] comment added to overcome the problem with the comment map association
	__Link__000001_SheetCells.Fieldtypename = `ref_models.XLCell`
	__Link__000001_SheetCells.FieldOffsetX = 0.000000
	__Link__000001_SheetCells.FieldOffsetY = 0.000000
	__Link__000001_SheetCells.TargetMultiplicity = models.MANY
	__Link__000001_SheetCells.TargetMultiplicityOffsetX = 0.000000
	__Link__000001_SheetCells.TargetMultiplicityOffsetY = 0.000000
	__Link__000001_SheetCells.SourceMultiplicity = models.MANY
	__Link__000001_SheetCells.SourceMultiplicityOffsetX = 0.000000
	__Link__000001_SheetCells.SourceMultiplicityOffsetY = 0.000000
	__Link__000001_SheetCells.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_SheetCells.StartRatio = 0.500000
	__Link__000001_SheetCells.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_SheetCells.EndRatio = 0.500000
	__Link__000001_SheetCells.CornerOffsetRatio = 1.380000

	__Link__000002_Cells.Name = `Cells`

	//gong:ident [ref_models.XLRow.Cells] comment added to overcome the problem with the comment map association
	__Link__000002_Cells.Identifier = `ref_models.XLRow.Cells`

	//gong:ident [ref_models.XLCell] comment added to overcome the problem with the comment map association
	__Link__000002_Cells.Fieldtypename = `ref_models.XLCell`
	__Link__000002_Cells.FieldOffsetX = 0.000000
	__Link__000002_Cells.FieldOffsetY = 0.000000
	__Link__000002_Cells.TargetMultiplicity = models.MANY
	__Link__000002_Cells.TargetMultiplicityOffsetX = 0.000000
	__Link__000002_Cells.TargetMultiplicityOffsetY = 0.000000
	__Link__000002_Cells.SourceMultiplicity = models.MANY
	__Link__000002_Cells.SourceMultiplicityOffsetX = 0.000000
	__Link__000002_Cells.SourceMultiplicityOffsetY = 0.000000
	__Link__000002_Cells.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Cells.StartRatio = 0.500000
	__Link__000002_Cells.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Cells.EndRatio = 0.500000
	__Link__000002_Cells.CornerOffsetRatio = 1.380000

	__Link__000003_Sheets.Name = `Sheets`

	//gong:ident [ref_models.XLFile.Sheets] comment added to overcome the problem with the comment map association
	__Link__000003_Sheets.Identifier = `ref_models.XLFile.Sheets`

	//gong:ident [ref_models.XLSheet] comment added to overcome the problem with the comment map association
	__Link__000003_Sheets.Fieldtypename = `ref_models.XLSheet`
	__Link__000003_Sheets.FieldOffsetX = 0.000000
	__Link__000003_Sheets.FieldOffsetY = 0.000000
	__Link__000003_Sheets.TargetMultiplicity = models.MANY
	__Link__000003_Sheets.TargetMultiplicityOffsetX = 0.000000
	__Link__000003_Sheets.TargetMultiplicityOffsetY = 0.000000
	__Link__000003_Sheets.SourceMultiplicity = models.MANY
	__Link__000003_Sheets.SourceMultiplicityOffsetX = 0.000000
	__Link__000003_Sheets.SourceMultiplicityOffsetY = 0.000000
	__Link__000003_Sheets.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Sheets.StartRatio = 0.234188
	__Link__000003_Sheets.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Sheets.EndRatio = 0.213354
	__Link__000003_Sheets.CornerOffsetRatio = 1.512866

	__Link__000004_XLFile.Name = `XLFile`

	//gong:ident [ref_models.DisplaySelection.XLFile] comment added to overcome the problem with the comment map association
	__Link__000004_XLFile.Identifier = `ref_models.DisplaySelection.XLFile`

	//gong:ident [ref_models.XLFile] comment added to overcome the problem with the comment map association
	__Link__000004_XLFile.Fieldtypename = `ref_models.XLFile`
	__Link__000004_XLFile.FieldOffsetX = 0.000000
	__Link__000004_XLFile.FieldOffsetY = 0.000000
	__Link__000004_XLFile.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_XLFile.TargetMultiplicityOffsetX = 0.000000
	__Link__000004_XLFile.TargetMultiplicityOffsetY = 0.000000
	__Link__000004_XLFile.SourceMultiplicity = models.MANY
	__Link__000004_XLFile.SourceMultiplicityOffsetX = 0.000000
	__Link__000004_XLFile.SourceMultiplicityOffsetY = 0.000000
	__Link__000004_XLFile.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_XLFile.StartRatio = 0.500000
	__Link__000004_XLFile.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_XLFile.EndRatio = 0.449374
	__Link__000004_XLFile.CornerOffsetRatio = -0.686646

	__Link__000005_XLSheet.Name = `XLSheet`

	//gong:ident [ref_models.DisplaySelection.XLSheet] comment added to overcome the problem with the comment map association
	__Link__000005_XLSheet.Identifier = `ref_models.DisplaySelection.XLSheet`

	//gong:ident [ref_models.XLSheet] comment added to overcome the problem with the comment map association
	__Link__000005_XLSheet.Fieldtypename = `ref_models.XLSheet`
	__Link__000005_XLSheet.FieldOffsetX = 0.000000
	__Link__000005_XLSheet.FieldOffsetY = 0.000000
	__Link__000005_XLSheet.TargetMultiplicity = models.ZERO_ONE
	__Link__000005_XLSheet.TargetMultiplicityOffsetX = 0.000000
	__Link__000005_XLSheet.TargetMultiplicityOffsetY = 0.000000
	__Link__000005_XLSheet.SourceMultiplicity = models.MANY
	__Link__000005_XLSheet.SourceMultiplicityOffsetX = 0.000000
	__Link__000005_XLSheet.SourceMultiplicityOffsetY = 0.000000
	__Link__000005_XLSheet.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000005_XLSheet.StartRatio = 0.584188
	__Link__000005_XLSheet.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000005_XLSheet.EndRatio = 0.955021
	__Link__000005_XLSheet.CornerOffsetRatio = 2.100167

	__Position__000000_Pos_Default_XLSheet.X = 65.000000
	__Position__000000_Pos_Default_XLSheet.Y = 203.000000
	__Position__000000_Pos_Default_XLSheet.Name = `Pos-Default-XLSheet`

	__Position__000001_Pos_Default_XLRow.X = 581.000000
	__Position__000001_Pos_Default_XLRow.Y = 339.000000
	__Position__000001_Pos_Default_XLRow.Name = `Pos-Default-XLRow`

	__Position__000002_Pos_Default_XLFile.X = 56.000000
	__Position__000002_Pos_Default_XLFile.Y = 1.000000
	__Position__000002_Pos_Default_XLFile.Name = `Pos-Default-XLFile`

	__Position__000003_Pos_Default_XLCell.X = 575.000000
	__Position__000003_Pos_Default_XLCell.Y = 189.000000
	__Position__000003_Pos_Default_XLCell.Name = `Pos-Default-XLCell`

	__Position__000004_Pos_Default_DisplaySelection.X = 706.000000
	__Position__000004_Pos_Default_DisplaySelection.Y = 3.000000
	__Position__000004_Pos_Default_DisplaySelection.Name = `Pos-Default-DisplaySelection`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_XLSheet_and_Default_XLRow.X = 435.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_XLSheet_and_Default_XLRow.Y = 120.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_XLSheet_and_Default_XLRow.Name = `Verticle in class diagram Default in middle between Default-XLSheet and Default-XLRow`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_XLSheet_and_Default_XLCell.X = 679.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_XLSheet_and_Default_XLCell.Y = 186.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_XLSheet_and_Default_XLCell.Name = `Verticle in class diagram Default in middle between Default-XLSheet and Default-XLCell`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_XLRow_and_Default_XLCell.X = 930.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_XLRow_and_Default_XLCell.Y = 317.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_XLRow_and_Default_XLCell.Name = `Verticle in class diagram Default in middle between Default-XLRow and Default-XLCell`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_XLFile_and_Default_XLSheet.X = 715.500000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_XLFile_and_Default_XLSheet.Y = 97.500000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_XLFile_and_Default_XLSheet.Name = `Verticle in class diagram Default in middle between Default-XLFile and Default-XLSheet`

	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_DisplaySelection_and_Default_XLFile.X = 666.500000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_DisplaySelection_and_Default_XLFile.Y = 35.000000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_DisplaySelection_and_Default_XLFile.Name = `Verticle in class diagram Default in middle between Default-DisplaySelection and Default-XLFile`

	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_DisplaySelection_and_Default_XLSheet.X = 745.500000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_DisplaySelection_and_Default_XLSheet.Y = 134.500000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_DisplaySelection_and_Default_XLSheet.Name = `Verticle in class diagram Default in middle between Default-DisplaySelection and Default-XLSheet`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_XLSheet)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_XLRow)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_XLFile)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_XLCell)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_DisplaySelection)
	// setup of Field instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_XLSheet.Position = __Position__000000_Pos_Default_XLSheet
	__GongStructShape__000000_Default_XLSheet.Fields = append(__GongStructShape__000000_Default_XLSheet.Fields, __Field__000000_Name)
	__GongStructShape__000000_Default_XLSheet.Links = append(__GongStructShape__000000_Default_XLSheet.Links, __Link__000000_Rows)
	__GongStructShape__000000_Default_XLSheet.Links = append(__GongStructShape__000000_Default_XLSheet.Links, __Link__000001_SheetCells)
	__GongStructShape__000001_Default_XLRow.Position = __Position__000001_Pos_Default_XLRow
	__GongStructShape__000001_Default_XLRow.Fields = append(__GongStructShape__000001_Default_XLRow.Fields, __Field__000001_RowIndex)
	__GongStructShape__000001_Default_XLRow.Links = append(__GongStructShape__000001_Default_XLRow.Links, __Link__000002_Cells)
	__GongStructShape__000002_Default_XLFile.Position = __Position__000002_Pos_Default_XLFile
	__GongStructShape__000002_Default_XLFile.Links = append(__GongStructShape__000002_Default_XLFile.Links, __Link__000003_Sheets)
	__GongStructShape__000003_Default_XLCell.Position = __Position__000003_Pos_Default_XLCell
	__GongStructShape__000003_Default_XLCell.Fields = append(__GongStructShape__000003_Default_XLCell.Fields, __Field__000002_Name)
	__GongStructShape__000003_Default_XLCell.Fields = append(__GongStructShape__000003_Default_XLCell.Fields, __Field__000003_X)
	__GongStructShape__000003_Default_XLCell.Fields = append(__GongStructShape__000003_Default_XLCell.Fields, __Field__000004_Y)
	__GongStructShape__000004_Default_DisplaySelection.Position = __Position__000004_Pos_Default_DisplaySelection
	__GongStructShape__000004_Default_DisplaySelection.Links = append(__GongStructShape__000004_Default_DisplaySelection.Links, __Link__000004_XLFile)
	__GongStructShape__000004_Default_DisplaySelection.Links = append(__GongStructShape__000004_Default_DisplaySelection.Links, __Link__000005_XLSheet)
	// setup of Link instances pointers
	__Link__000000_Rows.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_XLSheet_and_Default_XLRow
	__Link__000001_SheetCells.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_XLSheet_and_Default_XLCell
	__Link__000002_Cells.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_XLRow_and_Default_XLCell
	__Link__000003_Sheets.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_XLFile_and_Default_XLSheet
	__Link__000004_XLFile.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_DisplaySelection_and_Default_XLFile
	__Link__000005_XLSheet.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_DisplaySelection_and_Default_XLSheet
	// setup of Position instances pointers
	// setup of Vertice instances pointers
}
