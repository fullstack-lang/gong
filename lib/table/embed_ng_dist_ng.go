package table

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gong-lib-table/dist/ng-github.com-fullstack-lang-gong-lib-table
var NgDistNg embed.FS
