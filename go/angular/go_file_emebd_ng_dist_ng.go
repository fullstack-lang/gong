package angular

const EmebedNgDistNg = `//go:build !js

// generated code - do not edit
package {{pkgname}}

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed {{NgWorkspaceName}}/dist/{{NgWorkspaceName}}
var NgDistNg embed.FS
`

const EmbedNgDistNgWasm = `//go:build js && wasm

// generated code - do not edit
package {{pkgname}}

import "embed"

// NgDistNg is an empty filesystem in WASM mode since static files are not served by the WASM binary
var NgDistNg embed.FS
`

