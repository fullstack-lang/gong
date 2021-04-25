package models

const RootFileDocsTemplate = `// Package {{PkgName}} {{PkgName}} backend API
//
// {{PkgName}} exposes types of a package.
//
//     Schemes: http
//     BasePath: /api/{{PkgPathRoot}}/v1
//     Version: 0.1
//     Host: localhost:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package {{PkgName}}
`
