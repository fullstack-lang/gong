package tests

import (
	"log"
	"testing"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

// TestStageCallBack
//
// This test covers a a gong functionality where the Stage is updated
// through a callback that is defined in the "models" package
//
//
func TestBclassOrder(t *testing.T) {

	// setup GORM
	db := orm.SetupModels(false, "../../test.db")
	db.DB().SetMaxOpenConns(1)

	// initiate back repo a callback functions
	orm.BackRepo.Init(db)

	var aclasss []orm.AclassDB
	{
		query := db.Find(&aclasss)
		if query.Error != nil {
			t.Errorf("Find error %s", query.Error.Error())
			return
		}
	}

	var bclasss []orm.BclassDB
	{
		query := db.Find(&bclasss)
		if query.Error != nil {
			t.Errorf("Find error %s", query.Error.Error())
			return
		}
	}
	models.Stage.Checkout()

	// get the aclass instance with name "A1_1"
	var aclass_A1_1 *models.Aclass
	for aclass := range models.Stage.Aclasss {
		if aclass.Name == "A1_1" {
			aclass_A1_1 = aclass
			log.Println("found A1_1")

			for _, bclass := range aclass.Anarrayofb {
				log.Printf("bclass name %s ", bclass.Name)
			}
		}
	}

	// make a permutation of
	tmp := aclass_A1_1.Anarrayofb[2]
	aclass_A1_1.Anarrayofb[2] = aclass_A1_1.Anarrayofb[3]
	aclass_A1_1.Anarrayofb[3] = tmp
	models.Stage.Commit()

	db.Close()
}
