// generated from OrmFileSetupTemplate
package orm

import (
	"fmt"

	"github.com/glebarez/sqlite"
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
			TablePrefix: "github_com_fullstack_lang_gongdoc_go_", // table name prefix
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
		TablePrefix: "github_com_fullstack_lang_gongdoc_go_", // table name prefix
	}

	err := db.AutoMigrate( // insertion point for reference to structs
		&ClassdiagramDB{},
		&DiagramPackageDB{},
		&FieldDB{},
		&GongEnumShapeDB{},
		&GongEnumValueEntryDB{},
		&GongStructShapeDB{},
		&LinkDB{},
		&NodeDB{},
		&NoteShapeDB{},
		&NoteShapeLinkDB{},
		&PositionDB{},
		&TreeDB{},
		&UmlStateDB{},
		&UmlscDB{},
		&VerticeDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gongdoc/go")
	}
	// log.Printf("Database Migration of package github.com/fullstack-lang/gongdoc/go is OK")

	BackRepo.init(db)
}

func ResetDB(db *gorm.DB) { // insertion point for reference to structs
	db.Delete(&ClassdiagramDB{})
	db.Delete(&DiagramPackageDB{})
	db.Delete(&FieldDB{})
	db.Delete(&GongEnumShapeDB{})
	db.Delete(&GongEnumValueEntryDB{})
	db.Delete(&GongStructShapeDB{})
	db.Delete(&LinkDB{})
	db.Delete(&NodeDB{})
	db.Delete(&NoteShapeDB{})
	db.Delete(&NoteShapeLinkDB{})
	db.Delete(&PositionDB{})
	db.Delete(&TreeDB{})
	db.Delete(&UmlStateDB{})
	db.Delete(&UmlscDB{})
	db.Delete(&VerticeDB{})
}