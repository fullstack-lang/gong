package tree

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gong-lib-tree/dist/ng-github.com-fullstack-lang-gong-lib-tree
var NgDistNg embed.FS
