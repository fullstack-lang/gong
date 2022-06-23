package test

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//go:embed ng/dist/ng
var NgDistNg embed.FS

//go:embed go
var GoDir embed.FS
