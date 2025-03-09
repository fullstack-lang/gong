// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for StacksNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stacksnames StacksNames) ToString() (res string) {

	// migration of former implementation of enum
	switch stacksnames {
	// insertion code per enum code
	case Tone:
		res = "gongtone"
	case GongsvgStackName:
		res = "gongsvg"
	case SidebarTree:
		res = "sidebar tree"
	case GongtreeStackName:
		res = "gongtree"
	case GongtableStackName:
		res = "gongtable"
	case GongsimStackName:
		res = "gongsim"
	}
	return
}

func (stacksnames *StacksNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "gongtone":
		*stacksnames = Tone
		return
	case "gongsvg":
		*stacksnames = GongsvgStackName
		return
	case "sidebar tree":
		*stacksnames = SidebarTree
		return
	case "gongtree":
		*stacksnames = GongtreeStackName
		return
	case "gongtable":
		*stacksnames = GongtableStackName
		return
	case "gongsim":
		*stacksnames = GongsimStackName
		return
	default:
		return errUnkownEnum
	}
}

func (stacksnames *StacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Tone":
		*stacksnames = Tone
	case "GongsvgStackName":
		*stacksnames = GongsvgStackName
	case "SidebarTree":
		*stacksnames = SidebarTree
	case "GongtreeStackName":
		*stacksnames = GongtreeStackName
	case "GongtableStackName":
		*stacksnames = GongtableStackName
	case "GongsimStackName":
		*stacksnames = GongsimStackName
	default:
		return errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) ToCodeString() (res string) {

	switch *stacksnames {
	// insertion code per enum code
	case Tone:
		res = "Tone"
	case GongsvgStackName:
		res = "GongsvgStackName"
	case SidebarTree:
		res = "SidebarTree"
	case GongtreeStackName:
		res = "GongtreeStackName"
	case GongtableStackName:
		res = "GongtableStackName"
	case GongsimStackName:
		res = "GongsimStackName"
	}
	return
}

func (stacksnames StacksNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Tone")
	res = append(res, "GongsvgStackName")
	res = append(res, "SidebarTree")
	res = append(res, "GongtreeStackName")
	res = append(res, "GongtableStackName")
	res = append(res, "GongsimStackName")

	return
}

func (stacksnames StacksNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "gongtone")
	res = append(res, "gongsvg")
	res = append(res, "sidebar tree")
	res = append(res, "gongtree")
	res = append(res, "gongtable")
	res = append(res, "gongsim")

	return
}

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
		return errUnkownEnum
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
