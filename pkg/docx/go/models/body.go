package models

type Body struct {
	Name       string
	Paragraphs []*Paragraph
	Tables     []*Table

	// for enabling navigation between paragraphs
	LastParagraph *Paragraph
}
