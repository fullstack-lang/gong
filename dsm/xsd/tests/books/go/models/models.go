// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// insertion point for enums declaration

// From xsd simple type with enumerate restriction "bookFormatEnum"
type Enum_BookFormatEnum string

const (
	Enum_BookFormatEnum_Paperback Enum_BookFormatEnum = "Paperback"
	Enum_BookFormatEnum_Hardcover Enum_BookFormatEnum = "Hardcover"
)

// From xsd simple type with enumerate restriction "titleType"
type Enum_TitleType string

const ()

// insertion point for gongstructs declarations

// BookType Named source named complex type "bookType"
//
//	This complex type defines the structure of a book, including title,
//
// author, year, and format.
type BookType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "extendedAttributes
	AttributeGroup_extendedAttributes

	// generated from group with order 3 depth 1
	Group_bookDetailsGroup

	// generated from element "credit" of type credit order 4 depth 1
	Credit []*Credit `xml:"credit,omitempty"`
}

// Group_bookDetailsGroup UnNamed source named group "bookDetailsGroup"
type Group_bookDetailsGroup struct {

	// insertion point for fields

	// generated from element "title" of type titleType order 7 depth 1
	Title string `xml:"title,omitempty"`

	// generated from element "author" of type string order 8 depth 1
	Author string `xml:"author,omitempty"`

	// generated from element "year" of type yearType order 9 depth 1
	Year int `xml:"year,omitempty"`

	// generated from element "format" of type bookFormatEnum order 10 depth 1
	Format string `xml:"format,omitempty"`
}

// Books Named source element books within root schema
type Books struct {
	Name string `xml:"-"`

	// insertion point for fields

	// necessary since it is a root element
	XMLName xml.Name `xml:"books"`

	// generated from inline complex type
	A_books
}

// A_books Named source within outer element "books"
type A_books struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "book" of type bookType order 15 depth 2
	Book []*BookType `xml:"book,omitempty"`
}

// AttributeGroup_commonAttributes UnNamed source named attribute group "commonAttributes"
type AttributeGroup_commonAttributes struct {

	// insertion point for fields

	// generated from attribute "isbn
	Isbn string `xml:"isbn,attr,omitempty"`

	// generated from attribute "bestseller
	Bestseller bool `xml:"bestseller,attr,omitempty"`
}

// AttributeGroup_extendedAttributes UnNamed source named attribute group "extendedAttributes"
type AttributeGroup_extendedAttributes struct {

	// insertion point for fields

	// generated from attribute "edition
	Edition string `xml:"edition,attr,omitempty"`

	// generated from attribute group "commonAttributes
	AttributeGroup_commonAttributes
}

// Credit Named source named complex type "credit"
// The credit type ..
type Credit struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "page
	Page int `xml:"page,attr,omitempty"`

	// generated from element "credit-type" of type string order 20 depth 1
	Credit_type string `xml:"credit-type,omitempty"`

	// generated from element "link" of type link order 24 depth 1
	Link []*Link `xml:"link,omitempty"`

	// generated from element "credit-words" of type string order 25 depth 1
	Credit_words string `xml:"credit-words,omitempty"`

	// generated from element "credit-symbol" of type string order 26 depth 1
	Credit_symbol string `xml:"credit-symbol,omitempty"`
}

// Link Named source named complex type "link"
// The link type serves as an outgoing simple XLink. If a relative link
// is used within a document that is part of a compressed MusicXML file, the link is
// relative to the root folder of the zip file.
type Link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}
