package golang

const EmebedGoDirTemplate = `package {{pkgname}}

import "embed"

//go:embed go/models go/diagrams
var GoDir embed.FS
`
