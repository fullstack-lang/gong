package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	models "github.com/fullstack-lang/gong/dsm/xsd/tests/reqif/go/models"
)

func main() {
	// Open the XML file
	xmlFile, err := os.Open("../../samples/Sample.reqif")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	// Read the XML file
	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the XML into the Reqif struct
	var reqif models.REQ_IF
	err = xml.Unmarshal(byteValue, &reqif)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	// // Print the book details
	// for _, book := range reqif.Reqif {
	// 	fmt.Printf("ISBN: %s\n", book.ISBN)
	// 	fmt.Printf("Bestseller: %t\n", book.Bestseller)
	// 	if book.Edition != "" {
	// 		fmt.Printf("Edition: %s\n", book.Edition)
	// 	}
	// 	fmt.Printf("Title: %s\n", book.Title)
	// 	fmt.Printf("Author: %s\n", book.Author)
	// 	fmt.Printf("Year: %d\n", book.Year)
	// 	fmt.Printf("Format: %s\n", book.Format)
	// 	fmt.Println()
	// }

	// Marshal the reqif struct back into XML
	output, err := xml.MarshalIndent(reqif, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}

	// Write the XML to a new file
	err = os.WriteFile("../../samples/Sample_out.reqif", []byte(xml.Header+string(output)), 0o644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
