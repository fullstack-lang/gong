package models

import (
	"log"
	"path/filepath"
)

type Docx struct {
	Name string

	Files []*File

	Document *Document
}

func NewDocx(gongdocx_stage *Stage, path string, embed bool) (docx *Docx) {

	docx = (&Docx{Name: filepath.Base(path)}).Stage(gongdocx_stage)

	if err := docx2md(docx, gongdocx_stage, path, embed); err != nil {
		log.Fatal(err)
	}

	return
}
