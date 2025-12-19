package models

// GONGNOTE(NoteOnRune)
// for [models.Rune]
// -
// The "w:r" node, known as a run, represents a continuous run of text within a
// paragraph ("w:p" node) in a Word document's XML structure. It is found within
// the document.xml file.
// -
// Runs are segments of text within a paragraph that share the same formatting.
// This can include properties like bolding, italics, underlining, color, size,
// font, and more. The specific formatting is defined in a "w:rPr" (Run Properties)
// element within the "w:r" node.
// -
// A "w:r" node contains one or more "w:t" nodes, which hold the actual text
// content of the run. It can also contain other types of nodes like "w:br" for a
// line break or "w:tab" for a tab character.
// -
// When parsing a "w:r" node, your code should handle the formatting information
// provided in the "w:rPr" node (if present) and apply it to the text found within
// the "w:t" nodes.
const NoteOnRune = ""

type Rune struct {
	Name    string
	Content string

	Node           *Node
	Text           *Text
	RuneProperties *RuneProperties

	// navigation
	EnclosingParagraph *Paragraph
}
