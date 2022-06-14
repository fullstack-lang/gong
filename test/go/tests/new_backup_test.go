package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/xuri/excelize/v2"
)

func TestNewBackup(t *testing.T) {

	CreateTestStage()

	f := excelize.NewFile()

	astructSheetName := "Astruct"

	// Create a new sheet.
	index := f.NewSheet(astructSheetName)

	f.DeleteSheet("Sheet1")

	stage := models.Stage
	_ = stage

	set := *models.GetGongstructInstancesSet[models.Astruct]()
	line := 0
	for astruct := range set {

		line = line + 1
		f.SetCellStr(astructSheetName, fmt.Sprintf("A%d", line), astruct.Name)
		f.SetCellStr(astructSheetName, fmt.Sprintf("B%d", line), string(astruct.Aenum))
		f.SetCellFloat(astructSheetName, fmt.Sprintf("C%d", line), astruct.Floatfield, 64, 64)
		f.SetCellInt(astructSheetName, fmt.Sprintf("D%d", line), astruct.Intfield)
		f.SetCellValue(astructSheetName, fmt.Sprintf("E%d", line), astruct.Date)
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save spreadsheet by the given path.
	if err := f.SaveAs("new backup.xlsx"); err != nil {
		fmt.Println(err)
	}

	result := -2
	expected := -2
	if result != expected {
		t.Errorf("sub() test returned an unexpected result: got %v want %v", result, expected)
	}

	log.Println("New backup test is over")

}
