package models

// StageSetModel represents a developer-defined StageSet struct in a models package.
// It coordinates multiple stages across packages for unified persistence.
type StageSetModel struct {
	Name   string
	Fields []*StageSetField
}

// StageSetField represents a field of the StageSet struct that points to a Stage.
type StageSetField struct {
	Name        string // Field name on StageSet, e.g. "Stage", "XStage", "YStage"
	PackageName string // Package name, e.g. "models", "x", "y"
	PackagePath string // Full package import path, e.g. "github.com/fullstack-lang/gong/test/test2/go/models/x"
	IsLocal     bool   // true if this stage belongs to the current package being generated
	ImportAlias string // deterministic synthetic import alias, e.g. "__stage_0__"
}
