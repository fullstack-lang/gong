//go:build !js

package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/fullstack-lang/gong/dsm/xsd/tests/books/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/xsd/tests/books/go/models"
)

var (
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {
	log.SetPrefix("books: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("books", "", "", true, *embeddedDiagrams)

	if len(os.Args) < 2 {
		fmt.Println("Please provide a path to an XML file as the first argument.")
		return
	}

	// Open the XML file
	xmlFile, err := os.Open(os.Args[1])
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

	// Unmarshal the XML into the Books struct
	var books models.Books
	err = xml.Unmarshal(byteValue, &books)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	for _, book := range books.Book {
		book.Name = book.Title
		for _, credit := range book.Credit {
			credit.Name = credit.Credit_type
		}
	}

	models.StageBranch(stack.Stage, &books)
	stack.Stage.Commit()

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err = stack.R.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
