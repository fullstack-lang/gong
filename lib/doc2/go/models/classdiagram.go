package models

// Classdiagram diagram struct store a class diagram
// swagger:model Classdiagram
type Classdiagram struct {
	Name string

	// gong:text width:400 height:400
	Description string

	// IsIncludedInStaticWebSite is true if the diagram
	// is present in the generated static web site documenting
	// the package
	IsIncludedInStaticWebSite bool

	// list of gongstructshapes in the diagram
	GongStructShapes []*GongStructShape

	GongEnumShapes []*GongEnumShape

	// list of notes in the diagram
	GongNoteShapes []*GongNoteShape

	// IsInRenameMode means the user can edit the node
	IsInRenameMode bool

	// IsExpanded is true if the corresponding Classdiagram node itself is expanded in a tree.
	IsExpanded bool

	// NodeGongStructsIsExpanded stores the expansion state of the "GongStructs" category node.
	NodeGongStructsIsExpanded bool

	// NodeGongStructNodeExpansion stores the expanded status per individual named struct node
	// within the "GongStructs" category, as a JSON string (e.g., "[true, false, true]").
	NodeGongStructNodeExpansion string // Refactored from NodeGongStructNodeExpansionBinaryEncoding int

	// NodeGongEnumsIsExpanded stores the expansion state of the "GongEnums" category node.
	NodeGongEnumsIsExpanded bool

	// NodeGongEnumNodeExpansion stores the expanded status per individual enum node
	// within the "GongEnums" category, as a JSON string.
	NodeGongEnumNodeExpansion string // Refactored from NodeGongEnumNodeExpansionBinaryEncoding int

	// NodeGongNotesIsExpanded stores the expansion state of the "GongNotes" category node.
	NodeGongNotesIsExpanded bool

	// NodeGongNoteNodeExpansion stores the expanded status per individual note node
	// within the "GongNotes" category, as a JSON string.
	NodeGongNoteNodeExpansion string // Refactored from NodeGongNoteNodeExpansionBinaryEncoding int
}

// DuplicateDiagram generates a new diagram with duplicated shapes
func (classdiagram *Classdiagram) DuplicateDiagram() (newClassdiagram *Classdiagram) {

	newClassdiagram = CopyBranch(classdiagram)

	// All fields, including boolean ...IsExpanded and string ...NodeExpansion JSON strings,
	// are copied by CopyBranch.
	return
}
