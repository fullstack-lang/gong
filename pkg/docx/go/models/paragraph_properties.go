package models

// GONGNOTE(NoteOnParagraphProperties) [models.ParagraphProperties]
// -
// The "w:pPr" node represents paragraph properties within a Word document's XML
// structure, and is found within a "w:p" (paragraph) node in the document.xml
// file.
// -
// It contains information about the formatting and layout of a paragraph. This can
// include properties like text alignment (left, right, centered, or justified),
// indentation, spacing before or after the paragraph, line spacing, and more.
// -
// The "w:pPr" node may also contain a "w:numPr" element for numbered or bulleted
// lists, and a "w:sectPr" element for section properties (if this paragraph marks
// the end of a section).
// -
// When parsing a "w:pPr" node, your code should use the information it provides to
// format the paragraph appropriately in your data structure or output format.
const NoteOnParagraphProperties = ""

type ParagraphProperties struct {
	Name    string
	Content string

	ParagraphStyle *ParagraphStyle

	Node *Node
}
