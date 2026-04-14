// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for TableExtraNameEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (tableextranameenum TableExtraNameEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch tableextranameenum {
	// insertion code per enum code
	case TableSelectExtraName:
		res = "tmp-picker"
	case TableSortExtraName:
		res = "tmp-sort"
	}
	return
}

func (tableextranameenum *TableExtraNameEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "tmp-picker":
		*tableextranameenum = TableSelectExtraName
		return
	case "tmp-sort":
		*tableextranameenum = TableSortExtraName
		return
	default:
		return errUnkownEnum
	}
}

func (tableextranameenum *TableExtraNameEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TableSelectExtraName":
		*tableextranameenum = TableSelectExtraName
	case "TableSortExtraName":
		*tableextranameenum = TableSortExtraName
	default:
		err = errUnkownEnum
	}
	return
}

func (tableextranameenum *TableExtraNameEnum) ToCodeString() (res string) {

	switch *tableextranameenum {
	// insertion code per enum code
	case TableSelectExtraName:
		res = "TableSelectExtraName"
	case TableSortExtraName:
		res = "TableSortExtraName"
	}
	return
}

func (tableextranameenum TableExtraNameEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TableSelectExtraName")
	res = append(res, "TableSortExtraName")

	return
}

func (tableextranameenum TableExtraNameEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "tmp-picker")
	res = append(res, "tmp-sort")

	return
}

// Utility function for TableExtraPathEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (tableextrapathenum TableExtraPathEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch tableextrapathenum {
	// insertion code per enum code
	case StackNamePostFixForTableForMainTable:
		res = "-table"
	case StackNamePostFixForTableForMainForm:
		res = "-form"
	case StackNamePostFixForTableForMainTree:
		res = "-sidebar"
	case StackNamePostFixForTableForAssociation:
		res = "-table-pick"
	case StackNamePostFixForTableForAssociationSorting:
		res = "-table-sort"
	}
	return
}

func (tableextrapathenum *TableExtraPathEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "-table":
		*tableextrapathenum = StackNamePostFixForTableForMainTable
		return
	case "-form":
		*tableextrapathenum = StackNamePostFixForTableForMainForm
		return
	case "-sidebar":
		*tableextrapathenum = StackNamePostFixForTableForMainTree
		return
	case "-table-pick":
		*tableextrapathenum = StackNamePostFixForTableForAssociation
		return
	case "-table-sort":
		*tableextrapathenum = StackNamePostFixForTableForAssociationSorting
		return
	default:
		return errUnkownEnum
	}
}

func (tableextrapathenum *TableExtraPathEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "StackNamePostFixForTableForMainTable":
		*tableextrapathenum = StackNamePostFixForTableForMainTable
	case "StackNamePostFixForTableForMainForm":
		*tableextrapathenum = StackNamePostFixForTableForMainForm
	case "StackNamePostFixForTableForMainTree":
		*tableextrapathenum = StackNamePostFixForTableForMainTree
	case "StackNamePostFixForTableForAssociation":
		*tableextrapathenum = StackNamePostFixForTableForAssociation
	case "StackNamePostFixForTableForAssociationSorting":
		*tableextrapathenum = StackNamePostFixForTableForAssociationSorting
	default:
		err = errUnkownEnum
	}
	return
}

func (tableextrapathenum *TableExtraPathEnum) ToCodeString() (res string) {

	switch *tableextrapathenum {
	// insertion code per enum code
	case StackNamePostFixForTableForMainTable:
		res = "StackNamePostFixForTableForMainTable"
	case StackNamePostFixForTableForMainForm:
		res = "StackNamePostFixForTableForMainForm"
	case StackNamePostFixForTableForMainTree:
		res = "StackNamePostFixForTableForMainTree"
	case StackNamePostFixForTableForAssociation:
		res = "StackNamePostFixForTableForAssociation"
	case StackNamePostFixForTableForAssociationSorting:
		res = "StackNamePostFixForTableForAssociationSorting"
	}
	return
}

func (tableextrapathenum TableExtraPathEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "StackNamePostFixForTableForMainTable")
	res = append(res, "StackNamePostFixForTableForMainForm")
	res = append(res, "StackNamePostFixForTableForMainTree")
	res = append(res, "StackNamePostFixForTableForAssociation")
	res = append(res, "StackNamePostFixForTableForAssociationSorting")

	return
}

func (tableextrapathenum TableExtraPathEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "-table")
	res = append(res, "-form")
	res = append(res, "-sidebar")
	res = append(res, "-table-pick")
	res = append(res, "-table-sort")

	return
}

// Utility function for TableName
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (tablename TableName) ToString() (res string) {

	// migration of former implementation of enum
	switch tablename {
	// insertion code per enum code
	case TableDefaultName:
		res = "Table"
	}
	return
}

func (tablename *TableName) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Table":
		*tablename = TableDefaultName
		return
	default:
		return errUnkownEnum
	}
}

func (tablename *TableName) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TableDefaultName":
		*tablename = TableDefaultName
	default:
		err = errUnkownEnum
	}
	return
}

func (tablename *TableName) ToCodeString() (res string) {

	switch *tablename {
	// insertion code per enum code
	case TableDefaultName:
		res = "TableDefaultName"
	}
	return
}

func (tablename TableName) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TableDefaultName")

	return
}

func (tablename TableName) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Table")

	return
}

// Utility function for TableTestNameEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (tabletestnameenum TableTestNameEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch tabletestnameenum {
	// insertion code per enum code
	case ManualyEditedTableStackName:
		res = "manualy edited table"
	case ManualyEditedFormStackName:
		res = "manualy edited form"
	case GeneratedTableStackName:
		res = "generated table"
	}
	return
}

func (tabletestnameenum *TableTestNameEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "manualy edited table":
		*tabletestnameenum = ManualyEditedTableStackName
		return
	case "manualy edited form":
		*tabletestnameenum = ManualyEditedFormStackName
		return
	case "generated table":
		*tabletestnameenum = GeneratedTableStackName
		return
	default:
		return errUnkownEnum
	}
}

func (tabletestnameenum *TableTestNameEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ManualyEditedTableStackName":
		*tabletestnameenum = ManualyEditedTableStackName
	case "ManualyEditedFormStackName":
		*tabletestnameenum = ManualyEditedFormStackName
	case "GeneratedTableStackName":
		*tabletestnameenum = GeneratedTableStackName
	default:
		err = errUnkownEnum
	}
	return
}

func (tabletestnameenum *TableTestNameEnum) ToCodeString() (res string) {

	switch *tabletestnameenum {
	// insertion code per enum code
	case ManualyEditedTableStackName:
		res = "ManualyEditedTableStackName"
	case ManualyEditedFormStackName:
		res = "ManualyEditedFormStackName"
	case GeneratedTableStackName:
		res = "GeneratedTableStackName"
	}
	return
}

func (tabletestnameenum TableTestNameEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ManualyEditedTableStackName")
	res = append(res, "ManualyEditedFormStackName")
	res = append(res, "GeneratedTableStackName")

	return
}

func (tabletestnameenum TableTestNameEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "manualy edited table")
	res = append(res, "manualy edited form")
	res = append(res, "generated table")

	return
}

// Utility function for ToolTipPositionEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (tooltippositionenum ToolTipPositionEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch tooltippositionenum {
	// insertion code per enum code
	case Below:
		res = "below"
	case Above:
		res = "above"
	case Left:
		res = "left"
	case Right:
		res = "right"
	}
	return
}

func (tooltippositionenum *ToolTipPositionEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "below":
		*tooltippositionenum = Below
		return
	case "above":
		*tooltippositionenum = Above
		return
	case "left":
		*tooltippositionenum = Left
		return
	case "right":
		*tooltippositionenum = Right
		return
	default:
		return errUnkownEnum
	}
}

func (tooltippositionenum *ToolTipPositionEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Below":
		*tooltippositionenum = Below
	case "Above":
		*tooltippositionenum = Above
	case "Left":
		*tooltippositionenum = Left
	case "Right":
		*tooltippositionenum = Right
	default:
		err = errUnkownEnum
	}
	return
}

func (tooltippositionenum *ToolTipPositionEnum) ToCodeString() (res string) {

	switch *tooltippositionenum {
	// insertion code per enum code
	case Below:
		res = "Below"
	case Above:
		res = "Above"
	case Left:
		res = "Left"
	case Right:
		res = "Right"
	}
	return
}

func (tooltippositionenum ToolTipPositionEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Below")
	res = append(res, "Above")
	res = append(res, "Left")
	res = append(res, "Right")

	return
}

func (tooltippositionenum ToolTipPositionEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "below")
	res = append(res, "above")
	res = append(res, "left")
	res = append(res, "right")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type PointerToGongstructEnumStringField interface {
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	//insertion point for pointers to enum int types
	FromCodeString(input string) (err error)
}

// Last line of the template
