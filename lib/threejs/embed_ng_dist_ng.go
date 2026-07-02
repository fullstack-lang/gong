package threejs

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gong-lib-threejs/dist/ng-github.com-fullstack-lang-gong-lib-threejs
var NgDistNg embed.FS
