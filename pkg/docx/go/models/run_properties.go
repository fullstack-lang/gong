package models

// GONGNOTE(NoteOnRunProperties)
// The "w:rPr" node represents run properties in a Word document's XML structure.
// It is found within a "w:r" (run) node in the document.xml file.
// -
// This node defines the formatting for a specific run of text within a paragraph.
// It can include properties like font size, font type, color, highlighting,
// bolding, italics, underlining, and more.
// -
// For example, a "w:rPr" node might contain a "w:sz" element for font size, a
// "w:color" element for text color, or "w:b" for bold formatting. The presence of
// elements like "w:b" (bold), "w:i" (italic), and "w:u" (underline) indicate that
// the formatting is applied, as they are toggled by their presence alone.
// -
// When parsing a "w:rPr" node, your code should use the information it provides to
// apply the appropriate formatting to the text in the run ("w:r") that contains
// this "w:rPr" node.
const NoteOnRunProperties = ""

type RuneProperties struct {
	Name     string
	Node     *Node
	IsBold   bool
	IsStrike bool
	IsItalic bool

	Content string
}
