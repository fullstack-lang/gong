package button

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gong-lib-button/dist/ng-github.com-fullstack-lang-gong-lib-button
var NgDistNg embed.FS
