package golang

const NpmPackageJsonTemplate = `{
    "name": "{{pkgname}}",
    "workspaces": [
        "{{NgWorkspaceName}}",
        "vendor/github.com/fullstack-lang/gong/lib/button/ng-github.com-fullstack-lang-gong-lib-button",
        "vendor/github.com/fullstack-lang/gong/lib/cursor/ng-github.com-fullstack-lang-gong-lib-cursor",
        "vendor/github.com/fullstack-lang/gong/lib/doc/ng-github.com-fullstack-lang-gong-lib-doc",
        "vendor/github.com/fullstack-lang/gong/lib/doc2/ng-github.com-fullstack-lang-gong-lib-doc2",
        "vendor/github.com/fullstack-lang/gong/lib/gantt/ng-github.com-fullstack-lang-gong-lib-gantt",
        "vendor/github.com/fullstack-lang/gong/lib/load/ng-github.com-fullstack-lang-gong-lib-load",
        "vendor/github.com/fullstack-lang/gong/lib/probe/ng-github.com-fullstack-lang-gong-lib-probe",
        "vendor/github.com/fullstack-lang/gong/lib/slider/ng-github.com-fullstack-lang-gong-lib-slider",
        "vendor/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split",
        "vendor/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg",
        "vendor/github.com/fullstack-lang/gong/lib/sim/ng-github.com-fullstack-lang-gong-lib-sim",
        "vendor/github.com/fullstack-lang/gong/lib/ssg/ng-github.com-fullstack-lang-gong-lib-ssg",
        "vendor/github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table",
        "vendor/github.com/fullstack-lang/gong/lib/tone/ng-github.com-fullstack-lang-gong-lib-tone",
        "vendor/github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree",
        "vendor/github.com/fullstack-lang/gong/lib/xlsx/ng-github.com-fullstack-lang-gong-lib-xlsx",
        "vendor/github.com/fullstack-lang/gong/ng-github.com-fullstack-lang-gong",
        ""
    ]
}
`
