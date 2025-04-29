package doc2

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gong-lib-doc2/dist/ng-github.com-fullstack-lang-gong-lib-doc2
var NgDistNg embed.FS
