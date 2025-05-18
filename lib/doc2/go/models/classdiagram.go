package models

// Classdiagram diagram struct store a class diagram
// temporary here
// swagger:model Classdiagram
type Classdiagram struct {
	Name string

	// gong:text width:400 height:400
	Description string

	// IsIncludedInStaticWebSite is true is the diagram
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

	// IsExpanded is true if the corresponding node is expanded
	IsExpanded bool

	NodeGongStructsIsExpanded bool

	// encoding of the the expanded status per named struct node
	// the bit position is the NodeNamedStruct position
	// beyond 63, there is no storage of expanded status
	//
	// bit 0 at 0 means first gong struct is not encoded
	// bit 1 at 0 means second gong struct is not encoded
	// etc...
	NodeGongStructNodeExpansionBinaryEncoding int

	NodeGongEnumsIsExpanded                 bool
	NodeGongEnumNodeExpansionBinaryEncoding int

	NodeGongNotesIsExpanded                 bool
	NodeGongNoteNodeExpansionBinaryEncoding int
}

// DuplicateDiagram generates a new diagram with duplicated shapes
func (classdiagram *Classdiagram) DuplicateDiagram() (newClassdiagram *Classdiagram) {

	newClassdiagram = CopyBranch(classdiagram)

	return
}

func IsNodeExpanded(binaryEncoding, nodeRank int) bool {
	if nodeRank < 0 || nodeRank > 63 {
		// Rank must be between 0 and 63 for a 64-bit integer
		return false
	}
	// Check if the rank-th note is played
	return binaryEncoding&(1<<nodeRank) != 0
}

func ToggleNodeExpanded(binaryEncoding *int, nodeRank int) {
	if nodeRank < 0 || nodeRank > 63 {
		return // Ignore invalid ranks
	}

	*binaryEncoding ^= 1 << nodeRank
}
