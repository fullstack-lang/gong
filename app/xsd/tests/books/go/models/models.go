// generated code - do not edit
package models

import (
	"encoding/xml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// insertion point for enums declaration

// From xsd simple type with enumerate restriction "bookFormatEnum"
type Enum_BookFormatEnum string

const (
	Enum_BookFormatEnum_Paperback Enum_BookFormatEnum = "Paperback"
	Enum_BookFormatEnum_Hardcover Enum_BookFormatEnum = "Hardcover"
)

// insertion point for gongstructs declarations

// BookType Named source named complex type "bookType"
//  This complex type defines the structure of a book, including title,
// author, year, and format. 
type BookType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "extendedAttributes"
	AttributeGroup_extendedAttributes

	// generated from group "bookDetailsGroup" with order 4 depth 2
	Group_bookDetailsGroup

	// generated from element "credit" of complex type credit
	Credit []*Credit `xml:"credit,omitempty"`
}

// Group_bookDetailsGroup UnNamed source named group "bookDetailsGroup"
type Group_bookDetailsGroup struct {

	// insertion point for fields

	// generated from element "title" of simple type titleType
	Title string `xml:"title,omitempty"`

	// generated from element "author" of simple type xs:string
	Author string `xml:"author,omitempty"`

	// generated from element "year" of simple type yearType
	Year int `xml:"year,omitempty"`

	// generated from element "format" of simple type bookFormatEnum
	Format Enum_BookFormatEnum `xml:"format,omitempty"`
}

// Books Named source element books within schema
type Books struct {
	Name string `xml:"-"`

	// insertion point for fields

	// necessary since it is a root element
	XMLName xml.Name `xml:"books"`

	// generated from element "book" of complex type bookType
	Book []*BookType `xml:"book,omitempty"`
}

// AttributeGroup_commonAttributes UnNamed source named attribute group "commonAttributes"
type AttributeGroup_commonAttributes struct {

	// insertion point for fields

	// generated from attribute "isbn"
	Isbn string `xml:"isbn,attr,omitempty"`

	// generated from attribute "bestseller"
	Bestseller bool `xml:"bestseller,attr,omitempty"`
}

// AttributeGroup_extendedAttributes UnNamed source named attribute group "extendedAttributes"
type AttributeGroup_extendedAttributes struct {

	// insertion point for fields

	// generated from attribute "edition"
	Edition string `xml:"edition,attr,omitempty"`

	// generated from attribute group "commonAttributes"
	AttributeGroup_commonAttributes
}

// Credit Named source named complex type "credit"
// The credit type .. 
type Credit struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "page"
	Page int `xml:"page,attr,omitempty"`

	// generated from element "credit-type" of simple type xs:string
	Credit_type string `xml:"credit-type,omitempty"`

	// generated from element "link" of complex type link
	Link []*Link `xml:"link,omitempty"`

	// generated from element "credit-words" of simple type xs:string
	Credit_words string `xml:"credit-words,omitempty"`

	// generated from element "credit-symbol" of simple type xs:string
	Credit_symbol string `xml:"credit-symbol,omitempty"`
}

// Link Named source named complex type "link"
// The link type serves as an outgoing simple XLink. If a relative link
// is used within a document that is part of a compressed MusicXML file, the link is
// relative to the root folder of the zip file.
type Link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name"
	NameXSD string `xml:"name,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}
