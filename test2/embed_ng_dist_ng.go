package test2

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng/dist/ng
var NgDistNg embed.FS
