package gong

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gong/dist/ng-github.com-fullstack-lang-gong
var NgDistNg embed.FS
