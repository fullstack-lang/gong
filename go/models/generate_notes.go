package models

import (
	"bytes"
	"go/doc"
	"log"
)

// GenerateDocs populates modelPkg with Gongnote presnt in the docPackage
func (modelPkg *ModelPkg) GenerateDocs(docPackage *doc.Package) {

	for noteName, notes := range docPackage.Notes {

		for _, note := range notes {
			log.Println("note uid : ", note.UID)
			log.Println("note : ", note.Body)
			docBuf := bytes.Buffer{}
			indent := "	"
			indentedWidth := 1
			doc.ToText(&docBuf, note.Body, "", indent, indentedWidth)
			gongNote := (&GongNote{Name: note.UID,
				Body: note.Body,
			})
			modelPkg.GongNotes[note.UID] = gongNote
		}
		log.Println("documenting ", noteName)

	}
}
