package models

// A NoteShapeLink is a visual link from a shape to
// a Link or to a Classshape
//
// # The end point of the link is computed from the Type
//
// Because of[issue](https://github.com/golang/go/issues/57559)
// the referenced identifiers in the note comments are not
// renamed. Unlinke for identifiers used for referencing
// stuff, there is workaround this limitation
type NoteShapeLink struct {
	Name string

	// Identifier of the target shape / link of the note link
	//gong:ident
	Identifier string

	// Type of the target shape / link of the note link
	Type NoteShapeLinkType
}
