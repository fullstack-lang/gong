package models

// NoteShape is a UML note in a class diagram
type NoteShape struct {
	Name string

	// Identifier of the note object in the model package
	// a note is a special coment syntax https://pkg.go.dev/go/doc#Note
	// see vendor/github.com/fullstack-lang/gong/go/models/gong_note.go
	// since a note name is not an identifier, it cannot be renamed by gopls
	// unless it is understood to be the name of the identifier immediatly
	// following the comment
	//
	// Gongdoc makes this option a wrtting condition for note:
	// a comment text in gong containing a
	// 'gongnote' have to be followed by a constant identifier.
	//
	//    // GONGNOTE(ShortNodeOnModels): this is an example of a short note
	//    // It uses the DocLink convention for referencing Identifiers
	//    // In this case [Line], [Point] and [Line.Start]
	//    const ShortNodeOnModels = ""
	//
	// the following directive directs gong to manage this field as an identifier field
	//
	//gong:ident
	Identifier string

	Body string

	BodyHTML string

	X, Y          float64
	Width, Heigth float64
	Matched       bool // if a note with the same name has been found

	NoteShapeLinks []*NoteShapeLink
}
