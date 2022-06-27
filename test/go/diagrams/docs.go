// Package diagrams implements uml diagrams for documenting the test stack.
//
// The diagrams package is a set of go files, one for each diagram. A diagram is
// a go variable of type either github.com/fullstack-lang/gongdoc/go/models.Classdiagram or
// github.com/fullstack-lang/gongdoc/go/models.Umlsc (for state chart diagrams)
//
// Diagram go files should not be manualy edited but instead be edited and then saved by the gongdoc tool (github.com/fullstack-lang/gongdoc).
// Diagram go files can be read at runtime by the stack (see https://github.com/golang/go/issues/52974 for justification)
package diagrams
