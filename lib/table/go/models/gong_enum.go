// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for FormGroupName
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (formgroupname FormGroupName) ToString() (res string) {

	// migration of former implementation of enum
	switch formgroupname {
	// insertion code per enum code
	case FormGroupDefaultName:
		res = "Form"
	}
	return
}

func (formgroupname *FormGroupName) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Form":
		*formgroupname = FormGroupDefaultName
		return
	default:
		return errUnkownEnum
	}
}

func (formgroupname *FormGroupName) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FormGroupDefaultName":
		*formgroupname = FormGroupDefaultName
	default:
		return errUnkownEnum
	}
	return
}

func (formgroupname *FormGroupName) ToCodeString() (res string) {

	switch *formgroupname {
	// insertion code per enum code
	case FormGroupDefaultName:
		res = "FormGroupDefaultName"
	}
	return
}

func (formgroupname FormGroupName) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FormGroupDefaultName")

	return
}

func (formgroupname FormGroupName) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Form")

	return
}

// Utility function for InputTypeEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (inputtypeenum InputTypeEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch inputtypeenum {
	// insertion code per enum code
	case Text:
		res = "text"
	case Password:
		res = "password"
	case Number:
		res = "number"
	case Email:
		res = "email"
	case Tel:
		res = "tel"
	case Date:
		res = "date"
	case Datetime:
		res = "datetime-local"
	case Time:
		res = "time"
	case URL:
		res = "url"
	case Search:
		res = "search"
	case Range:
		res = "range"
	case Color:
		res = "color"
	case File:
		res = "file"
	case Hidden:
		res = "hidden"
	case Month:
		res = "month"
	case Week:
		res = "week"
	}
	return
}

func (inputtypeenum *InputTypeEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "text":
		*inputtypeenum = Text
		return
	case "password":
		*inputtypeenum = Password
		return
	case "number":
		*inputtypeenum = Number
		return
	case "email":
		*inputtypeenum = Email
		return
	case "tel":
		*inputtypeenum = Tel
		return
	case "date":
		*inputtypeenum = Date
		return
	case "datetime-local":
		*inputtypeenum = Datetime
		return
	case "time":
		*inputtypeenum = Time
		return
	case "url":
		*inputtypeenum = URL
		return
	case "search":
		*inputtypeenum = Search
		return
	case "range":
		*inputtypeenum = Range
		return
	case "color":
		*inputtypeenum = Color
		return
	case "file":
		*inputtypeenum = File
		return
	case "hidden":
		*inputtypeenum = Hidden
		return
	case "month":
		*inputtypeenum = Month
		return
	case "week":
		*inputtypeenum = Week
		return
	default:
		return errUnkownEnum
	}
}

func (inputtypeenum *InputTypeEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Text":
		*inputtypeenum = Text
	case "Password":
		*inputtypeenum = Password
	case "Number":
		*inputtypeenum = Number
	case "Email":
		*inputtypeenum = Email
	case "Tel":
		*inputtypeenum = Tel
	case "Date":
		*inputtypeenum = Date
	case "Datetime":
		*inputtypeenum = Datetime
	case "Time":
		*inputtypeenum = Time
	case "URL":
		*inputtypeenum = URL
	case "Search":
		*inputtypeenum = Search
	case "Range":
		*inputtypeenum = Range
	case "Color":
		*inputtypeenum = Color
	case "File":
		*inputtypeenum = File
	case "Hidden":
		*inputtypeenum = Hidden
	case "Month":
		*inputtypeenum = Month
	case "Week":
		*inputtypeenum = Week
	default:
		return errUnkownEnum
	}
	return
}

func (inputtypeenum *InputTypeEnum) ToCodeString() (res string) {

	switch *inputtypeenum {
	// insertion code per enum code
	case Text:
		res = "Text"
	case Password:
		res = "Password"
	case Number:
		res = "Number"
	case Email:
		res = "Email"
	case Tel:
		res = "Tel"
	case Date:
		res = "Date"
	case Datetime:
		res = "Datetime"
	case Time:
		res = "Time"
	case URL:
		res = "URL"
	case Search:
		res = "Search"
	case Range:
		res = "Range"
	case Color:
		res = "Color"
	case File:
		res = "File"
	case Hidden:
		res = "Hidden"
	case Month:
		res = "Month"
	case Week:
		res = "Week"
	}
	return
}

func (inputtypeenum InputTypeEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Text")
	res = append(res, "Password")
	res = append(res, "Number")
	res = append(res, "Email")
	res = append(res, "Tel")
	res = append(res, "Date")
	res = append(res, "Datetime")
	res = append(res, "Time")
	res = append(res, "URL")
	res = append(res, "Search")
	res = append(res, "Range")
	res = append(res, "Color")
	res = append(res, "File")
	res = append(res, "Hidden")
	res = append(res, "Month")
	res = append(res, "Week")

	return
}

func (inputtypeenum InputTypeEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "text")
	res = append(res, "password")
	res = append(res, "number")
	res = append(res, "email")
	res = append(res, "tel")
	res = append(res, "date")
	res = append(res, "datetime-local")
	res = append(res, "time")
	res = append(res, "url")
	res = append(res, "search")
	res = append(res, "range")
	res = append(res, "color")
	res = append(res, "file")
	res = append(res, "hidden")
	res = append(res, "month")
	res = append(res, "week")

	return
}

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
		return errUnkownEnum
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
		return errUnkownEnum
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
		return errUnkownEnum
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
		return errUnkownEnum
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
	
	FromCodeString(input string) (err error)
}

// Last line of the template
