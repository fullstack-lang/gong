package models

// PathToGoSubDirectory for instance "/tmp"
var PathToGoSubDirectory string

// OrmPkgGenPath is target path for orm package, for instance "/tmp/libraryorm"
var OrmPkgGenPath string

// ControllersPkgGenPath is target path for controllers package
var ControllersPkgGenPath string

// FullstackPkgGenPath is target path for Fullstack package
var FullstackPkgGenPath string

// StaticPkgGenPath is target path for Static package
var StaticPkgGenPath string

// NgWorkspacePath is the path to the Ng Workspace
var NgWorkspacePath string

// NgDataLibrarySourceCodeDirectory is the "<pkgName>/src/lib" directory where,
// by angular CLI convention,
// the source code for the library's components,
// services, modules, and other features are located.
//
// gongc regenerates at each compilation a material angular library "<pkgName>" for
// having the code related to data manipulation of objects of the
// stack
var NgDataLibrarySourceCodeDirectory string

// NgSpecificLibrarySourceCodeDirectory is the "<pkgName>specific/src/lib"
// where the developper stores its specific code for the front component
//
// This library is generated once at the stack creation
var NgSpecificLibrarySourceCodeDirectory string

// MaterialLibDatamodelTargetPath is the "<pkgName>datamodel/src/lib"
//
// gongc regenerates at each compilation a material angular library "<pkgName>datamodel" for
// having the code related to inclusion of the gongdoc stack and the gong stack
// for docmenting the stack with UML diagrams (gongdoc) and providing a meta inspection
// of the declarations of the stack
var MaterialLibDatamodelTargetPath string
