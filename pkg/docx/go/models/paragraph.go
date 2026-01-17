package models

// GONGNOTE(NoteOnParagraph)
// The "w:p" node represents a [models.Paragraph] in a Word document's XML structure.
// -
// It is one of the primary building blocks of a document's content within the
// document.xml file.
// -
// Each "w:p" element contains a series of "w:r" (run) elements, which represent
// sections of text within the paragraph that have consistent formatting. These
// "w:r" elements, in turn, contain "w:t" elements that hold the actual text
// content.
// -
// The "w:p" element may also contain various other child elements that provide
// additional information about the paragraph. For example, "w:pPr" specifies
// paragraph properties like alignment, indentation, and spacing.
// -
// As you parse a "w:p" node, you would typically create a new paragraph object in
// your code, then parse the child nodes to fill in the content and properties of
// that paragraph.
const NoteOnParagraph = ""

type Paragraph struct {
	Name                string
	Content             string
	Node                *Node
	ParagraphProperties *ParagraphProperties
	Runes               []*Rune

	CollatedText string // collated text of all runes

	// navigation
	Next                 *Paragraph
	Previous             *Paragraph
	EnclosingBody        *Body
	EnclosingTableColumn *TableColumn
}
