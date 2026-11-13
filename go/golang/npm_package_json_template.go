package golang

const NpmPackageJsonTemplate = `{
    "name": "{{pkgname}}",
    "workspaces": [
        "{{NgWorkspaceName}}",
        "vendor/github.com/fullstack-lang/gong/lib/button/ng-github.com-fullstack-lang-gong-lib-button",
        "vendor/github.com/fullstack-lang/gong/lib/cesium/ng-github.com-fullstack-lang-gong-lib-cesium",
        "vendor/github.com/fullstack-lang/gong/lib/cursor/ng-github.com-fullstack-lang-gong-lib-cursor",
        "vendor/github.com/fullstack-lang/gong/lib/form/ng-github.com-fullstack-lang-gong-lib-form",
        "vendor/github.com/fullstack-lang/gong/lib/load/ng-github.com-fullstack-lang-gong-lib-load",
        "vendor/github.com/fullstack-lang/gong/lib/maplibre/ng-github.com-fullstack-lang-gong-lib-maplibre",
        "vendor/github.com/fullstack-lang/gong/lib/markdown/ng-github.com-fullstack-lang-gong-lib-markdown",
        "vendor/github.com/fullstack-lang/gong/lib/sim/ng-github.com-fullstack-lang-gong-lib-sim",
        "vendor/github.com/fullstack-lang/gong/lib/slider/ng-github.com-fullstack-lang-gong-lib-slider",
        "vendor/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split",
        "vendor/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg",
        "vendor/github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table",
        "vendor/github.com/fullstack-lang/gong/lib/threejs/ng-github.com-fullstack-lang-gong-lib-threejs",
        "vendor/github.com/fullstack-lang/gong/lib/tone/ng-github.com-fullstack-lang-gong-lib-tone",
        "vendor/github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree"
    ],
    "dependencies": {
        "@cesium/engine": "^26.5.0",
        "@types/katex": "^0.16.8",
        "katex": "^0.16.21",
        "maplibre-gl": "^6.17.0",
        "marked": "^18.0.0",
        "marked-katex-extension": "^5.1.12"
    }
}
`
