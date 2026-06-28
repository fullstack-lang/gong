package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	gongxsd_musicxml_models "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models"
)

func main() {
	// Open the XML file
	xmlFile, err := os.Open("../../phylotaxy music 2 september 2024 0800.musicxml")
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

	// // Convert UTF-16 to UTF-8
	// utf16Decoder := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
	// utf8Data, _, err := transform.Bytes(utf16Decoder, byteValue)
	// if err != nil {
	// 	fmt.Println("Error converting to UTF-8:", err)
	// 	return
	// }

	// Unmarshal the XML into the Books struct
	var scorePartwise gongxsd_musicxml_models.Score_partwise
	err = xml.Unmarshal(byteValue, &scorePartwise)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	// // Print the book details
	// for _, book := range books.Books {
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

	// Marshal the books struct back into XML
	output, err := xml.MarshalIndent(scorePartwise, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}

	// Write the XML to a new file
	err = os.WriteFile("../../new phylotaxy music 2 september 2024 0800.musicxml", []byte(xml.Header+string(output)), 0o644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
