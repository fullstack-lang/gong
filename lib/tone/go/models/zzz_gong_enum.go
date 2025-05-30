// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for Status
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (status Status) ToString() (res string) {

	// migration of former implementation of enum
	switch status {
	// insertion code per enum code
	case PLAYING:
		res = "PLAYING"
	case PAUSED:
		res = "PAUSED"
	}
	return
}

func (status *Status) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PLAYING":
		*status = PLAYING
		return
	case "PAUSED":
		*status = PAUSED
		return
	default:
		return errUnkownEnum
	}
}

func (status *Status) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PLAYING":
		*status = PLAYING
	case "PAUSED":
		*status = PAUSED
	default:
		err = errUnkownEnum
	}
	return
}

func (status *Status) ToCodeString() (res string) {

	switch *status {
	// insertion code per enum code
	case PLAYING:
		res = "PLAYING"
	case PAUSED:
		res = "PAUSED"
	}
	return
}

func (status Status) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PLAYING")
	res = append(res, "PAUSED")

	return
}

func (status Status) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PLAYING")
	res = append(res, "PAUSED")

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
