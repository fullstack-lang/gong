package golang

const GitIgnoreTempl = `node_modules

go/cmd/{{pkgname}}/{{pkgname}}
go/cmd/{{pkgname}}/{{pkgname}}.exe
go/cmd/{{pkgname}}/__debug_bin*
`
