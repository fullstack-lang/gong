package test

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng/dist/ng/browser
var NgDistNg embed.FS
