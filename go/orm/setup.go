// generated from OrmFileSetupTemplate
package orm

import (
	"fmt"

	"github.com/fullstack-lang/gong/go/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// SetupModels connects to the sqlite database
func SetupModels(stage *models.StageStruct, logMode bool, filepath string) *gorm.DB {
	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filepath), gormConfig)

	if err != nil {
		panic("Failed to connect to database!")
	}

	AutoMigrate(stage, db)

	return db
}

// AutoMigrate migrates db with with orm Struct
func AutoMigrate(stage *models.StageStruct, db *gorm.DB) {
	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_go_", // table name prefix
	}

	err := db.AutoMigrate( // insertion point for reference to structs
		&GongBasicFieldDB{},
		&GongEnumDB{},
		&GongEnumValueDB{},
		&GongLinkDB{},
		&GongNoteDB{},
		&GongStructDB{},
		&GongTimeFieldDB{},
		&MetaDB{},
		&MetaReferenceDB{},
		&ModelPkgDB{},
		&PointerToGongStructFieldDB{},
		&SliceOfPointerToGongStructFieldDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/go")
	}
	// log.Printf("Database Migration of package github.com/fullstack-lang/gong/go is OK")

	BackRepo.init(stage, db)
}

func ResetDB(db *gorm.DB) { // insertion point for reference to structs
	db.Delete(&GongBasicFieldDB{})
	db.Delete(&GongEnumDB{})
	db.Delete(&GongEnumValueDB{})
	db.Delete(&GongLinkDB{})
	db.Delete(&GongNoteDB{})
	db.Delete(&GongStructDB{})
	db.Delete(&GongTimeFieldDB{})
	db.Delete(&MetaDB{})
	db.Delete(&MetaReferenceDB{})
	db.Delete(&ModelPkgDB{})
	db.Delete(&PointerToGongStructFieldDB{})
	db.Delete(&SliceOfPointerToGongStructFieldDB{})
}
