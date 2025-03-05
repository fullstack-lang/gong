package golang

const NpmPackageJsonTemplate = `{
    "name": "{{pkgname}}",
    "workspaces": [
        "{{NgWorkspaceName}}",
        "vendor/github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree",
        "vendor/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg",
        "vendor/github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table",
        "vendor/github.com/fullstack-lang/gong/lib/doc/ng-github.com-fullstack-lang-gong-lib-doc",
        ""
    ]
}
`
