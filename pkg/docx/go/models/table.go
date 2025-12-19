package models

// GONGNOTE(NoteOnTable)
// The "w:tbl" node represents a table within a Word document's XML structure. It
// is found in the document.xml file.
// -
// This node defines the structure and formatting of a table in the document. It
// contains child elements that represent the table's properties ("w:tblPr"),
// grid ("w:tblGrid"), and rows ("w:tr").
// -
// The "w:tblPr" node defines the table's overall properties, such as its width,
// alignment, borders, and shading.
// -
// The "w:tblGrid" node defines the table's grid structure - specifically, the
// number and width of the columns.
// -
// The "w:tr" nodes represent table rows, and each row contains "w:tc" nodes that
// represent the individual cells within that row.
// -
// When parsing a "w:tbl" node, your code should handle the table structure and
// formatting information it provides to correctly represent the table in your
// data structure or output format.
const NoteOnTable = ""

type Table struct {
	Name            string
	Node            *Node
	Content         string
	TableProperties *TableProperties
	TableRows       []*TableRow
}
