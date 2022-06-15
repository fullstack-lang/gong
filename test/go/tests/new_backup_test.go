package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/xuri/excelize/v2"

	"github.com/stretchr/testify/assert"
)

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

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
	line := 1

	// write headers
	var _astruct models.Astruct
	for index, fieldName := range _astruct.GetFields() {
		assert.NoError(t, f.SetCellStr(astructSheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName))
	}

	for astruct := range set {

		line = line + 1

		for index, fieldName := range _astruct.GetFields() {
			assert.NoError(t, f.SetCellStr(astructSheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), astruct.GetFieldStringValue(fieldName)))
		}
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
