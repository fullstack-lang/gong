package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type InlineStruct struct {
	FieldA string `xml:"FieldA"`
	FieldB string `xml:"FieldB"`
}

type Root struct {
	Inlines []*InlineStruct `xml:",inline"`
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
	fmt.Println(string(output))

	// Optionally, save the XML to a file
	file, err := os.Create("output.xml")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer file.Close()

	file.WriteString(xml.Header)
	file.Write(output)
}
