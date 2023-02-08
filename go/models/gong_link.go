package models

// GongLink stores the result of the parsing of a Doc Link
//
// It is appended to a [GongNote]
// see [Doc Links](https://tip.golang.org/doc/comment) which has been added to the go in 1.19 and
// is accessible in the [AST](https://pkg.go.dev/go/doc/comment@go1.19.4#DocLink).
type GongLink struct {

	// store the link without the brackets, for instance or "Astruct" or "AstructBstructUse" in
	// "Astruct.AstructBstructUse"
	Name string

	Recv       string // "Astruct" if docLink is "Astruct.AstructBstructUse"
	ImportPath string
}
