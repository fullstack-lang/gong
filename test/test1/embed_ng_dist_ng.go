package test1

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gong-test-test1/dist/ng-github.com-fullstack-lang-gong-test-test1
var NgDistNg embed.FS
