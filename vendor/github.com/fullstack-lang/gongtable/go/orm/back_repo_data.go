// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	CellAPIs []*CellAPI

	CellBooleanAPIs []*CellBooleanAPI

	CellFloat64APIs []*CellFloat64API

	CellIconAPIs []*CellIconAPI

	CellIntAPIs []*CellIntAPI

	CellStringAPIs []*CellStringAPI

	CheckBoxAPIs []*CheckBoxAPI

	DisplayedColumnAPIs []*DisplayedColumnAPI

	FormDivAPIs []*FormDivAPI

	FormEditAssocButtonAPIs []*FormEditAssocButtonAPI

	FormFieldAPIs []*FormFieldAPI

	FormFieldDateAPIs []*FormFieldDateAPI

	FormFieldDateTimeAPIs []*FormFieldDateTimeAPI

	FormFieldFloat64APIs []*FormFieldFloat64API

	FormFieldIntAPIs []*FormFieldIntAPI

	FormFieldSelectAPIs []*FormFieldSelectAPI

	FormFieldStringAPIs []*FormFieldStringAPI

	FormFieldTimeAPIs []*FormFieldTimeAPI

	FormGroupAPIs []*FormGroupAPI

	FormSortAssocButtonAPIs []*FormSortAssocButtonAPI

	OptionAPIs []*OptionAPI

	RowAPIs []*RowAPI

	TableAPIs []*TableAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, cellDB := range backRepo.BackRepoCell.Map_CellDBID_CellDB {

		var cellAPI CellAPI
		cellAPI.ID = cellDB.ID
		cellAPI.CellPointersEncoding = cellDB.CellPointersEncoding
		cellDB.CopyBasicFieldsToCell_WOP(&cellAPI.Cell_WOP)

		backRepoData.CellAPIs = append(backRepoData.CellAPIs, &cellAPI)
	}

	for _, cellbooleanDB := range backRepo.BackRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB {

		var cellbooleanAPI CellBooleanAPI
		cellbooleanAPI.ID = cellbooleanDB.ID
		cellbooleanAPI.CellBooleanPointersEncoding = cellbooleanDB.CellBooleanPointersEncoding
		cellbooleanDB.CopyBasicFieldsToCellBoolean_WOP(&cellbooleanAPI.CellBoolean_WOP)

		backRepoData.CellBooleanAPIs = append(backRepoData.CellBooleanAPIs, &cellbooleanAPI)
	}

	for _, cellfloat64DB := range backRepo.BackRepoCellFloat64.Map_CellFloat64DBID_CellFloat64DB {

		var cellfloat64API CellFloat64API
		cellfloat64API.ID = cellfloat64DB.ID
		cellfloat64API.CellFloat64PointersEncoding = cellfloat64DB.CellFloat64PointersEncoding
		cellfloat64DB.CopyBasicFieldsToCellFloat64_WOP(&cellfloat64API.CellFloat64_WOP)

		backRepoData.CellFloat64APIs = append(backRepoData.CellFloat64APIs, &cellfloat64API)
	}

	for _, celliconDB := range backRepo.BackRepoCellIcon.Map_CellIconDBID_CellIconDB {

		var celliconAPI CellIconAPI
		celliconAPI.ID = celliconDB.ID
		celliconAPI.CellIconPointersEncoding = celliconDB.CellIconPointersEncoding
		celliconDB.CopyBasicFieldsToCellIcon_WOP(&celliconAPI.CellIcon_WOP)

		backRepoData.CellIconAPIs = append(backRepoData.CellIconAPIs, &celliconAPI)
	}

	for _, cellintDB := range backRepo.BackRepoCellInt.Map_CellIntDBID_CellIntDB {

		var cellintAPI CellIntAPI
		cellintAPI.ID = cellintDB.ID
		cellintAPI.CellIntPointersEncoding = cellintDB.CellIntPointersEncoding
		cellintDB.CopyBasicFieldsToCellInt_WOP(&cellintAPI.CellInt_WOP)

		backRepoData.CellIntAPIs = append(backRepoData.CellIntAPIs, &cellintAPI)
	}

	for _, cellstringDB := range backRepo.BackRepoCellString.Map_CellStringDBID_CellStringDB {

		var cellstringAPI CellStringAPI
		cellstringAPI.ID = cellstringDB.ID
		cellstringAPI.CellStringPointersEncoding = cellstringDB.CellStringPointersEncoding
		cellstringDB.CopyBasicFieldsToCellString_WOP(&cellstringAPI.CellString_WOP)

		backRepoData.CellStringAPIs = append(backRepoData.CellStringAPIs, &cellstringAPI)
	}

	for _, checkboxDB := range backRepo.BackRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB {

		var checkboxAPI CheckBoxAPI
		checkboxAPI.ID = checkboxDB.ID
		checkboxAPI.CheckBoxPointersEncoding = checkboxDB.CheckBoxPointersEncoding
		checkboxDB.CopyBasicFieldsToCheckBox_WOP(&checkboxAPI.CheckBox_WOP)

		backRepoData.CheckBoxAPIs = append(backRepoData.CheckBoxAPIs, &checkboxAPI)
	}

	for _, displayedcolumnDB := range backRepo.BackRepoDisplayedColumn.Map_DisplayedColumnDBID_DisplayedColumnDB {

		var displayedcolumnAPI DisplayedColumnAPI
		displayedcolumnAPI.ID = displayedcolumnDB.ID
		displayedcolumnAPI.DisplayedColumnPointersEncoding = displayedcolumnDB.DisplayedColumnPointersEncoding
		displayedcolumnDB.CopyBasicFieldsToDisplayedColumn_WOP(&displayedcolumnAPI.DisplayedColumn_WOP)

		backRepoData.DisplayedColumnAPIs = append(backRepoData.DisplayedColumnAPIs, &displayedcolumnAPI)
	}

	for _, formdivDB := range backRepo.BackRepoFormDiv.Map_FormDivDBID_FormDivDB {

		var formdivAPI FormDivAPI
		formdivAPI.ID = formdivDB.ID
		formdivAPI.FormDivPointersEncoding = formdivDB.FormDivPointersEncoding
		formdivDB.CopyBasicFieldsToFormDiv_WOP(&formdivAPI.FormDiv_WOP)

		backRepoData.FormDivAPIs = append(backRepoData.FormDivAPIs, &formdivAPI)
	}

	for _, formeditassocbuttonDB := range backRepo.BackRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB {

		var formeditassocbuttonAPI FormEditAssocButtonAPI
		formeditassocbuttonAPI.ID = formeditassocbuttonDB.ID
		formeditassocbuttonAPI.FormEditAssocButtonPointersEncoding = formeditassocbuttonDB.FormEditAssocButtonPointersEncoding
		formeditassocbuttonDB.CopyBasicFieldsToFormEditAssocButton_WOP(&formeditassocbuttonAPI.FormEditAssocButton_WOP)

		backRepoData.FormEditAssocButtonAPIs = append(backRepoData.FormEditAssocButtonAPIs, &formeditassocbuttonAPI)
	}

	for _, formfieldDB := range backRepo.BackRepoFormField.Map_FormFieldDBID_FormFieldDB {

		var formfieldAPI FormFieldAPI
		formfieldAPI.ID = formfieldDB.ID
		formfieldAPI.FormFieldPointersEncoding = formfieldDB.FormFieldPointersEncoding
		formfieldDB.CopyBasicFieldsToFormField_WOP(&formfieldAPI.FormField_WOP)

		backRepoData.FormFieldAPIs = append(backRepoData.FormFieldAPIs, &formfieldAPI)
	}

	for _, formfielddateDB := range backRepo.BackRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB {

		var formfielddateAPI FormFieldDateAPI
		formfielddateAPI.ID = formfielddateDB.ID
		formfielddateAPI.FormFieldDatePointersEncoding = formfielddateDB.FormFieldDatePointersEncoding
		formfielddateDB.CopyBasicFieldsToFormFieldDate_WOP(&formfielddateAPI.FormFieldDate_WOP)

		backRepoData.FormFieldDateAPIs = append(backRepoData.FormFieldDateAPIs, &formfielddateAPI)
	}

	for _, formfielddatetimeDB := range backRepo.BackRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB {

		var formfielddatetimeAPI FormFieldDateTimeAPI
		formfielddatetimeAPI.ID = formfielddatetimeDB.ID
		formfielddatetimeAPI.FormFieldDateTimePointersEncoding = formfielddatetimeDB.FormFieldDateTimePointersEncoding
		formfielddatetimeDB.CopyBasicFieldsToFormFieldDateTime_WOP(&formfielddatetimeAPI.FormFieldDateTime_WOP)

		backRepoData.FormFieldDateTimeAPIs = append(backRepoData.FormFieldDateTimeAPIs, &formfielddatetimeAPI)
	}

	for _, formfieldfloat64DB := range backRepo.BackRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB {

		var formfieldfloat64API FormFieldFloat64API
		formfieldfloat64API.ID = formfieldfloat64DB.ID
		formfieldfloat64API.FormFieldFloat64PointersEncoding = formfieldfloat64DB.FormFieldFloat64PointersEncoding
		formfieldfloat64DB.CopyBasicFieldsToFormFieldFloat64_WOP(&formfieldfloat64API.FormFieldFloat64_WOP)

		backRepoData.FormFieldFloat64APIs = append(backRepoData.FormFieldFloat64APIs, &formfieldfloat64API)
	}

	for _, formfieldintDB := range backRepo.BackRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntDB {

		var formfieldintAPI FormFieldIntAPI
		formfieldintAPI.ID = formfieldintDB.ID
		formfieldintAPI.FormFieldIntPointersEncoding = formfieldintDB.FormFieldIntPointersEncoding
		formfieldintDB.CopyBasicFieldsToFormFieldInt_WOP(&formfieldintAPI.FormFieldInt_WOP)

		backRepoData.FormFieldIntAPIs = append(backRepoData.FormFieldIntAPIs, &formfieldintAPI)
	}

	for _, formfieldselectDB := range backRepo.BackRepoFormFieldSelect.Map_FormFieldSelectDBID_FormFieldSelectDB {

		var formfieldselectAPI FormFieldSelectAPI
		formfieldselectAPI.ID = formfieldselectDB.ID
		formfieldselectAPI.FormFieldSelectPointersEncoding = formfieldselectDB.FormFieldSelectPointersEncoding
		formfieldselectDB.CopyBasicFieldsToFormFieldSelect_WOP(&formfieldselectAPI.FormFieldSelect_WOP)

		backRepoData.FormFieldSelectAPIs = append(backRepoData.FormFieldSelectAPIs, &formfieldselectAPI)
	}

	for _, formfieldstringDB := range backRepo.BackRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB {

		var formfieldstringAPI FormFieldStringAPI
		formfieldstringAPI.ID = formfieldstringDB.ID
		formfieldstringAPI.FormFieldStringPointersEncoding = formfieldstringDB.FormFieldStringPointersEncoding
		formfieldstringDB.CopyBasicFieldsToFormFieldString_WOP(&formfieldstringAPI.FormFieldString_WOP)

		backRepoData.FormFieldStringAPIs = append(backRepoData.FormFieldStringAPIs, &formfieldstringAPI)
	}

	for _, formfieldtimeDB := range backRepo.BackRepoFormFieldTime.Map_FormFieldTimeDBID_FormFieldTimeDB {

		var formfieldtimeAPI FormFieldTimeAPI
		formfieldtimeAPI.ID = formfieldtimeDB.ID
		formfieldtimeAPI.FormFieldTimePointersEncoding = formfieldtimeDB.FormFieldTimePointersEncoding
		formfieldtimeDB.CopyBasicFieldsToFormFieldTime_WOP(&formfieldtimeAPI.FormFieldTime_WOP)

		backRepoData.FormFieldTimeAPIs = append(backRepoData.FormFieldTimeAPIs, &formfieldtimeAPI)
	}

	for _, formgroupDB := range backRepo.BackRepoFormGroup.Map_FormGroupDBID_FormGroupDB {

		var formgroupAPI FormGroupAPI
		formgroupAPI.ID = formgroupDB.ID
		formgroupAPI.FormGroupPointersEncoding = formgroupDB.FormGroupPointersEncoding
		formgroupDB.CopyBasicFieldsToFormGroup_WOP(&formgroupAPI.FormGroup_WOP)

		backRepoData.FormGroupAPIs = append(backRepoData.FormGroupAPIs, &formgroupAPI)
	}

	for _, formsortassocbuttonDB := range backRepo.BackRepoFormSortAssocButton.Map_FormSortAssocButtonDBID_FormSortAssocButtonDB {

		var formsortassocbuttonAPI FormSortAssocButtonAPI
		formsortassocbuttonAPI.ID = formsortassocbuttonDB.ID
		formsortassocbuttonAPI.FormSortAssocButtonPointersEncoding = formsortassocbuttonDB.FormSortAssocButtonPointersEncoding
		formsortassocbuttonDB.CopyBasicFieldsToFormSortAssocButton_WOP(&formsortassocbuttonAPI.FormSortAssocButton_WOP)

		backRepoData.FormSortAssocButtonAPIs = append(backRepoData.FormSortAssocButtonAPIs, &formsortassocbuttonAPI)
	}

	for _, optionDB := range backRepo.BackRepoOption.Map_OptionDBID_OptionDB {

		var optionAPI OptionAPI
		optionAPI.ID = optionDB.ID
		optionAPI.OptionPointersEncoding = optionDB.OptionPointersEncoding
		optionDB.CopyBasicFieldsToOption_WOP(&optionAPI.Option_WOP)

		backRepoData.OptionAPIs = append(backRepoData.OptionAPIs, &optionAPI)
	}

	for _, rowDB := range backRepo.BackRepoRow.Map_RowDBID_RowDB {

		var rowAPI RowAPI
		rowAPI.ID = rowDB.ID
		rowAPI.RowPointersEncoding = rowDB.RowPointersEncoding
		rowDB.CopyBasicFieldsToRow_WOP(&rowAPI.Row_WOP)

		backRepoData.RowAPIs = append(backRepoData.RowAPIs, &rowAPI)
	}

	for _, tableDB := range backRepo.BackRepoTable.Map_TableDBID_TableDB {

		var tableAPI TableAPI
		tableAPI.ID = tableDB.ID
		tableAPI.TablePointersEncoding = tableDB.TablePointersEncoding
		tableDB.CopyBasicFieldsToTable_WOP(&tableAPI.Table_WOP)

		backRepoData.TableAPIs = append(backRepoData.TableAPIs, &tableAPI)
	}

}
