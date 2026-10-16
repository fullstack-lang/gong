// generated code by xsd2gong - do not edit
package xlink

import (
	"encoding/xml"
)

var _ xml.Attr

type EnumTypeType string

const (
	EnumTypeTypeSimple EnumTypeType = "simple"
	EnumTypeTypeExtended EnumTypeType = "extended"
	EnumTypeTypeTitle EnumTypeType = "title"
	EnumTypeTypeResource EnumTypeType = "resource"
	EnumTypeTypeLocator EnumTypeType = "locator"
	EnumTypeTypeArc EnumTypeType = "arc"
)

type EnumShowType string

const (
	EnumShowTypeNew EnumShowType = "new"
	EnumShowTypeReplace EnumShowType = "replace"
	EnumShowTypeEmbed EnumShowType = "embed"
	EnumShowTypeOther EnumShowType = "other"
	EnumShowTypeNone EnumShowType = "none"
)

type EnumActuateType string

const (
	EnumActuateTypeOnLoad EnumActuateType = "onLoad"
	EnumActuateTypeOnRequest EnumActuateType = "onRequest"
	EnumActuateTypeOther EnumActuateType = "other"
	EnumActuateTypeNone EnumActuateType = "none"
)

// Simple: Intended for use as the type of user-declared elements to make them
// simple links.
type Simple struct {
	Name string `xml:"-"`
	Type EnumTypeType `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href string `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole string `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show EnumShowType `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate EnumActuateType `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Extended: Intended for use as the type of user-declared elements to make them
// extended links.
// Note that the elements referenced in the content model are all abstract.
// The intention is that by simply declaring elements with these as their
// substitutionGroup, all the right things will happen.
type Extended struct {
	Name string `xml:"-"`
	Type EnumTypeType `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Role string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Title_1 []*TitleEltType `xml:"title,omitempty"`
	Resource []*ResourceType `xml:"resource,omitempty"`
	Locator []*LocatorType `xml:"locator,omitempty"`
	Arc []*ArcType `xml:"arc,omitempty"`
}

type TitleEltType struct {
	Name string `xml:"-"`
	Type EnumTypeType `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
	Value string `xml:",chardata"`
}

type ResourceType struct {
	Name string `xml:"-"`
	Type EnumTypeType `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Role string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Label string `xml:"http://www.w3.org/1999/xlink label,attr,omitempty"`
	Value string `xml:",chardata"`
}

type LocatorType struct {
	Name string `xml:"-"`
	Type EnumTypeType `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href string `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Label string `xml:"http://www.w3.org/1999/xlink label,attr,omitempty"`
	Title_1 []*TitleEltType `xml:"title,omitempty"`
}

type ArcType struct {
	Name string `xml:"-"`
	Type EnumTypeType `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Arcrole string `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show EnumShowType `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate EnumActuateType `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
	From string `xml:"http://www.w3.org/1999/xlink from,attr,omitempty"`
	To string `xml:"http://www.w3.org/1999/xlink to,attr,omitempty"`
	Title_1 []*TitleEltType `xml:"title,omitempty"`
}

