//go:build !js

package models

// the following blank imports will force the vendoring of the front end code needed for splitlite
// this allows the front end compilation of the splitlite front end
import (
	_ "github.com/fullstack-lang/gong/lib/button/ng-github.com-fullstack-lang-gong-lib-button"
	_ "github.com/fullstack-lang/gong/lib/form/ng-github.com-fullstack-lang-gong-lib-form"
	_ "github.com/fullstack-lang/gong/lib/load/ng-github.com-fullstack-lang-gong-lib-load"
	_ "github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg"
	_ "github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table"
	_ "github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree"
)
