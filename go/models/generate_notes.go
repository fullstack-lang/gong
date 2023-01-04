package models

import (
	"bytes"
	"go/doc"
	"go/doc/comment"
	"log"
)

// GenerateDocs populates modelPkg with Gongnote presnt in the docPackage
func (modelPkg *ModelPkg) GenerateDocs(docPackage *doc.Package) {

	for noteName, notes := range docPackage.Notes {

		for _, note := range notes {
			// log.Println("note uid : ", note.UID)
			// log.Println("note : ", note.Body)
			docBuf := bytes.Buffer{}
			indent := "	"
			indentedWidth := 1
			doc.ToText(&docBuf, note.Body, "", indent, indentedWidth)
			gongNote := (&GongNote{Name: note.UID,
				Body: note.Body,
			})
			modelPkg.GongNotes[note.UID] = gongNote

			// parse body to find doclinks
			var p comment.Parser
			doc := p.Parse(note.Body)

			for _, docLink := range doc.Links {

				link := (&GongLink{Name: docLink.Text})
				gongNote.Links = append(gongNote.Links, link)
			}

		}
		log.Println("documenting ", noteName)

	}
}
