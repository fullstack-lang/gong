package markdown

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gong-lib-markdown/dist/ng-github.com-fullstack-lang-gong-lib-markdown
var NgDistNg embed.FS
