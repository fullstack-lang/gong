//go:build !js

package main

import (
	"flag"
	"log"
	// musicxml_stack "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/stack"
	// musicxml_static "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")

	musicxmlFile = flag.String("musicxmlFile", "", "")
)

func main() {
	log.SetPrefix("musicxml: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	// r := musicxml_static.ServeStaticFiles(*logGINFlag)

	// // setup stack
	// stack := musicxml_stack.NewStack(r, "musicxml", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	// stack.Probe.Refresh()

	// // Open the XML file
	// xmlFile, err := os.Open(*musicxmlFile)
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer xmlFile.Close()

	// // Read the XML file
	// byteValue, err := io.ReadAll(xmlFile)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }

	// // Convert UTF-16 to UTF-8
	// utf16Decoder := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
	// utf8Data, _, err := transform.Bytes(utf16Decoder, byteValue)
	// if err != nil {
	// 	fmt.Println("Error converting to UTF-8:", err)
	// 	return
	// }

	// // Unmarshal the XML into the Reqif struct
	// var scorePartwise models.Score_partwise
	// err = xml.Unmarshal(utf8Data, &scorePartwise)
	// if err != nil {
	// 	fmt.Println("Error unmarshalling XML:", err)
	// 	return
	// }

	// models.StageBranch(stack.Stage, &scorePartwise)

	// stack.Stage.Commit()
	// stack.Probe.Refresh()

	// log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	// err = r.Run(":" + strconv.Itoa(*port))
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
}
