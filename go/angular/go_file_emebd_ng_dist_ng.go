package angular

const EmebedNgDistNg = `package {{pkgname}}

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed {{NgWorkspaceName}}/dist/{{NgWorkspaceName}}
var NgDistNg embed.FS
`
