package golang

const NpmPackageJsonTemplate = `{
    "name": "{{pkgname}}",
    "workspaces": [
        "{{NgWorkspaceName}}",
        "vendor/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree",
        "vendor/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg",
        "vendor/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable",
        "vendor/github.com/fullstack-lang/gongdoc/ng-github.com-fullstack-lang-gongdoc",
        ""
    ]
}
`
