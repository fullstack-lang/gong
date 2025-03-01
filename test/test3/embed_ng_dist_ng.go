package test3

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gong-test3/dist/ng-github.com-fullstack-lang-gong-test3
var NgDistNg embed.FS
