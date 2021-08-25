package models

const OrmFileSetupTemplate = `// generated from OrmFileSetupTemplate
package orm

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// SetupModels connects to the sqlite database
func SetupModels(logMode bool, filepath string) *gorm.DB {
	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "{{PkgPathRootWithoutSlashes}}_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filepath), gormConfig)

	if err != nil {
		panic("Failed to connect to database!")
	}

	AutoMigrate(db)

	return db
}

// AutoMigrate migrates db with with orm Struct
func AutoMigrate(db *gorm.DB) {
	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "{{PkgPathRootWithoutSlashes}}_", // table name prefix
	}

	err := db.AutoMigrate( // insertion point for reference to structs{{` + string(OrmSetupRefToStructDB) + `}}
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package {{PkgPathRoot}}")
	}
	log.Printf("Database Migration of package {{PkgPathRoot}} is OK")

	BackRepo.init(db)
}

func ResetDB(db *gorm.DB) { // insertion point for reference to structs{{` + string(OrmSetupDelete) + `}}
}`

type OrmSetupCumulSubTemplate string

const (
	OrmSetupRefToStructDB OrmSetupCumulSubTemplate = "OrmSetupRefToStructDB"
	OrmSetupDelete        OrmSetupCumulSubTemplate = "OrmSetupDelete"
)

var OrmSetupCumulSubTemplateCode map[string]string = // new line
map[string]string{

	string(OrmSetupRefToStructDB): `
		&{{Structname}}DB{},`,

	string(OrmSetupDelete): `
	db.Delete(&{{Structname}}DB{})`,
}
