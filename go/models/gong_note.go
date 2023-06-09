package models

// Gongnote is a documentation note that contribues to the model documentation
//
// A Gongnote is a go comment that contains the keywork "GONGNOTE(<Name>): <body>"
// for instance
// ....
// GONGDOC(NoteOnOrganizationOfLineAndPoints): first line
// this is an example of a note that
// could be displayed on a diagram.
//
// The comment can be standalone or attached to a declation/var
// const NoteOnOrganizationOfLineAndPoints = ""
type GongNote struct {
	Name string
	Body string

	// conversion of the comment into HTML (uses of go comment )
	BodyHTML string

	Links []*GongLink
}
