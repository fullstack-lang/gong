package models

// Gongnote is a documentation note that contribues to the model documentation
//
// A Gongnote is a go comment that contains the keywork "GONGNOTE(<Name>): <body>"
// for instance
// ....
// GONGDOC(note on the organization between line and points): first line
// this is an example of a note that
// could be displayed on a diagram.
//
// The comment can be standalone or attached to a declation/var
type GongNote struct {
	Name string
	Body string
}
