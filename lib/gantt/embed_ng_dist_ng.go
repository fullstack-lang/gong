package gantt

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gong-lib-gantt/dist/ng-github.com-fullstack-lang-gong-lib-gantt
var NgDistNg embed.FS
