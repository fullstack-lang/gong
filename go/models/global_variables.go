package models

// PkgGoPath for generation
var PkgGoPath string

// PkgName is generated package name (for back) and generated project elements for front
var PkgName string

// PathToGoSubDirectory for instance "/tmp"
var PathToGoSubDirectory string

// OrmPkgGenPath is target path for orm package, for instance "/tmp/libraryorm"
var OrmPkgGenPath string

// ApiPkgGenPath is target path for api package
var ApiPkgGenPath string

// ControllersPkgGenPath is target path for controllers package, for instance "/tmp/librarycontrollers"
var ControllersPkgGenPath string

// // ModulesTargetPath is where ng modules are generated
// var ModulesTargetPath string

// MatTargetPath is where the ng components are generated
var MatTargetPath string

// NgWorkspacePath is the path to the Ng Workspace
var NgWorkspacePath string

// ADDR is the network address addr where the angular generated service will lookup the server
var ADDR string
