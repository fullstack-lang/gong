// generated code - do not edit
package xlink

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type ArcType_WOP struct {
	// insertion point

	Name string

	Type EnumTypeType

	Arcrole string

	Title string

	Show EnumShowType

	Actuate EnumActuateType

	From string

	To string
}

func (from *ArcType) GongCopyBasicFields(to *ArcType) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
	to.From = from.From
	to.To = from.To
}

type Extended_WOP struct {
	// insertion point

	Name string

	Type EnumTypeType

	Role string

	Title string
}

func (from *Extended) GongCopyBasicFields(to *Extended) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Role = from.Role
	to.Title = from.Title
}

type LocatorType_WOP struct {
	// insertion point

	Name string

	Type EnumTypeType

	Href string

	Role string

	Title string

	Label string
}

func (from *LocatorType) GongCopyBasicFields(to *LocatorType) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.Href = from.Href
	to.Role = from.Role
	to.Title = from.Title
	to.Label = from.Label
}

type ResourceType_WOP struct {
	// insertion point

	Name string

	Type EnumTypeType

	Role string

	Title string

	Label string

	Value string
}

func (from *ResourceType) GongCopyBasicFields(to *ResourceType) {
	// insertion point
	*to = *from
}

type Simple_WOP struct {
	// insertion point

	Name string

	Type EnumTypeType

	Href string

	Role string

	Arcrole string

	Title string

	Show EnumShowType

	Actuate EnumActuateType

	Value string
}

func (from *Simple) GongCopyBasicFields(to *Simple) {
	// insertion point
	*to = *from
}

type TitleEltType_WOP struct {
	// insertion point

	Name string

	Type EnumTypeType

	Lang string

	Value string
}

func (from *TitleEltType) GongCopyBasicFields(to *TitleEltType) {
	// insertion point
	*to = *from
}

// end of insertion point
