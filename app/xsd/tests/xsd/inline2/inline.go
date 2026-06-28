package main

import (
	"encoding/xml"
	"fmt"
)

type InlineStruct struct {
	FieldA string `xml:"FieldA"`
	FieldB string `xml:"FieldB"`
}

type Root struct {
	Inlines []*InlineStruct
}

// Custom MarshalXML implementation for Root
func (r *Root) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Root"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, inline := range r.Inlines {
		if err := e.EncodeElement(inline.FieldA, xml.StartElement{Name: xml.Name{Local: "FieldA"}}); err != nil {
			return err
		}
		if err := e.EncodeElement(inline.FieldB, xml.StartElement{Name: xml.Name{Local: "FieldB"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

// Custom UnmarshalXML implementation for Root
func (r *Root) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inline *InlineStruct

	for {
		t, err := d.Token()
		if err != nil {
			return err
		}

		switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "FieldA":
				if inline == nil {
					inline = new(InlineStruct)
				}
				if err := d.DecodeElement(&inline.FieldA, &se); err != nil {
					return err
				}
			case "FieldB":
				if inline != nil {
					if err := d.DecodeElement(&inline.FieldB, &se); err != nil {
						return err
					}
					r.Inlines = append(r.Inlines, inline)
					inline = nil
				}
			}
		case xml.EndElement:
			if se.Name.Local == start.Name.Local {
				return nil
			}
		}
	}
}

func main() {
	// Create an instance of Root with some data
	root := new(Root)

	inline1 := new(InlineStruct)
	inline1.FieldA = "A1"
	inline1.FieldB = "B1"

	inline2 := new(InlineStruct)
	inline2.FieldA = "A2"
	inline2.FieldB = "B2"

	root.Inlines = append(root.Inlines, inline1, inline2)

	// Marshal the Root struct to XML
	output, err := xml.MarshalIndent(root, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// Print the XML output
	fmt.Println("Marshalled XML:")
	fmt.Println(string(output))

	// Simulate unmarshaling from the marshalled XML
	var unmarshalledRoot Root
	err = xml.Unmarshal(output, &unmarshalledRoot)
	if err != nil {
		fmt.Printf("error unmarshalling: %v\n", err)
		return
	}

	// Print the unmarshalled data
	fmt.Println("\nUnmarshalled Data:")
	for _, inline := range unmarshalledRoot.Inlines {
		fmt.Printf("FieldA: %s, FieldB: %s\n", inline.FieldA, inline.FieldB)
	}
}
