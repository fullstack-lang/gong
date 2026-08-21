package angular

const EmebedNgDistNg = `// generated code - do not edit
package {{pkgname}}

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed {{NgWorkspaceName}}/dist/{{NgWorkspaceName}}
var NgDistNg embed.FS
`
