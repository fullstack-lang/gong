package models

// GONGNOTE(NoteOnColumn)
// The "w:tr" node represents a table row within a Word document's XML structure.
// It is found as a child of the "w:tbl" (table) node in the document.xml file.
// -
// This node contains child elements that represent the individual cells ("w:tc")
// in the row, as well as the row properties ("w:trPr"), which include attributes
// like height, header status, and more.
// -
// The "w:tc" node, on the other hand, represents an individual table cell. It is
// found as a child of the "w:tr" (table row) node.
// -
// This node contains the content of the cell, which can include text,
// paragraphs, or even other tables. It also contains the cell properties
// ("w:tcPr"), which include attributes like cell width, vertical alignment,
// borders, shading, and more.
// -
// When parsing "w:tr" and "w:tc" nodes, your code should correctly map the
// structure of the table, row, and cell, and apply the appropriate properties
// to each element.
const NoteOnColumn = ""

type TableColumn struct {
	Name       string
	Content    string
	Node       *Node
	Paragraphs []*Paragraph
}
