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
	default:
		return errUnkownEnum
	}
	return
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
	case "password":
		*inputtypeenum = Password
	case "number":
		*inputtypeenum = Number
	case "email":
		*inputtypeenum = Email
	case "tel":
		*inputtypeenum = Tel
	case "date":
		*inputtypeenum = Date
	case "datetime-local":
		*inputtypeenum = Datetime
	case "time":
		*inputtypeenum = Time
	case "url":
		*inputtypeenum = URL
	case "search":
		*inputtypeenum = Search
	case "range":
		*inputtypeenum = Range
	case "color":
		*inputtypeenum = Color
	case "file":
		*inputtypeenum = File
	case "hidden":
		*inputtypeenum = Hidden
	case "month":
		*inputtypeenum = Month
	case "week":
		*inputtypeenum = Week
	default:
		return errUnkownEnum
	}
	return
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
	case "tmp-sort":
		*tableextranameenum = TableSortExtraName
	default:
		return errUnkownEnum
	}
	return
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

// Utility function for TableExtraPathEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (tableextrapathenum TableExtraPathEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch tableextrapathenum {
	// insertion code per enum code
	case StackNamePostFixForTableForAssociation:
		res = "-table"
	case StackNamePostFixForTableForAssociationSorting:
		res = "-table-sort"
	}
	return
}

func (tableextrapathenum *TableExtraPathEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "-table":
		*tableextrapathenum = StackNamePostFixForTableForAssociation
	case "-table-sort":
		*tableextrapathenum = StackNamePostFixForTableForAssociationSorting
	default:
		return errUnkownEnum
	}
	return
}

func (tableextrapathenum *TableExtraPathEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
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
	case StackNamePostFixForTableForAssociation:
		res = "StackNamePostFixForTableForAssociation"
	case StackNamePostFixForTableForAssociationSorting:
		res = "StackNamePostFixForTableForAssociationSorting"
	}
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
	default:
		return errUnkownEnum
	}
	return
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

// Last line of the template
