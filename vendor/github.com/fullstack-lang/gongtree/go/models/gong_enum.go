// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for TreeStackName
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (treestackname TreeStackName) ToString() (res string) {

	// migration of former implementation of enum
	switch treestackname {
	// insertion code per enum code
	case TreeStackDefaultName:
		res = "Tree"
	}
	return
}

func (treestackname *TreeStackName) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Tree":
		*treestackname = TreeStackDefaultName
	default:
		return errUnkownEnum
	}
	return
}

func (treestackname *TreeStackName) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TreeStackDefaultName":
		*treestackname = TreeStackDefaultName
	default:
		return errUnkownEnum
	}
	return
}

func (treestackname *TreeStackName) ToCodeString() (res string) {

	switch *treestackname {
	// insertion code per enum code
	case TreeStackDefaultName:
		res = "TreeStackDefaultName"
	}
	return
}

func (treestackname TreeStackName) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TreeStackDefaultName")

	return
}

func (treestackname TreeStackName) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Tree")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | TreeStackName
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*TreeStackName
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
