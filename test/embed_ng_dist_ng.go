package test

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gong-test-go-models/dist/ng-github.com-fullstack-lang-gong-test-go-models
var NgDistNg embed.FS
