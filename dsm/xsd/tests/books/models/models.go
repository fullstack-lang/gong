package models

import "encoding/xml"

// Book is generated from named complex type "bookType"
type Book struct {
	ISBN       string `xml:"isbn,attr"`
	Bestseller bool   `xml:"bestseller,attr"`
	Edition    string `xml:"edition,attr,omitempty"`
	Title      string `xml:"title"`
	Author     string `xml:"author"`
	Year       int    `xml:"year"`
	Format     string `xml:"format"`
}

// Books represents the root element containing multiple books
// Books is generated from inlined complex type within element "books"
type Books struct {
	XMLNSXSI       string `xml:"xmlns:xsi,attr"`
	SchemaLocation string `xml:"xsi:noNamespaceSchemaLocation,attr"`

	XMLName xml.Name `xml:"books"`
	Books   []Book   `xml:"book"`
}

func (b *Books) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "xsi":
			b.XMLNSXSI = attr.Value
		case "noNamespaceSchemaLocation":
			b.SchemaLocation = attr.Value
		}
	}
	type alias Books
	return d.DecodeElement((*alias)(b), &start)
}
