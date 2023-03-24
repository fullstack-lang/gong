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

			// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
			// to be recognized as a proper identifier.
			// While this was introduced in go 1.19, it is not yet implemented in
			// gopls (see [issue](https://github.com/golang/go/issues/57559)
			p.LookupPackage = func(name string) (importPath string, ok bool) {
				if name == "models" {
					return "models", true
				}
				return comment.DefaultLookupPackage(name)
			}
			p.LookupSym = func(recv, name string) (ok bool) {
				if recv == "" {
					return true
				}
				return false
			}

			doc := p.Parse(note.Body)

			var pr comment.Printer
			gongNote.BodyHTML = string(pr.HTML(doc))

			for _, block := range doc.Content {

				switch paragraph := block.(type) {
				case *comment.Paragraph:
					_ = paragraph
					for _, text := range paragraph.Text {
						switch docLink := text.(type) {
						case *comment.DocLink:

							link := (&GongLink{
								Name:       docLink.Name,
								Recv:       docLink.Recv,
								ImportPath: docLink.ImportPath,
							}).Stage(modelPkg.GetStage())

							gongNote.Links = append(gongNote.Links, link)
						}
					}
				}
			}
		}
		log.Println("documenting ", noteName)

	}
}
