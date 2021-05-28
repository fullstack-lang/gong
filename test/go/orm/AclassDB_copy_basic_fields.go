package orm

import (
	"time"

	"github.com/fullstack-lang/gong/test/go/models"
)

// declaration in order to justify use of the models import
var __Aclass__dummysDeclaration__ models.Aclass
var __Aclass_time__copyfields_dummyDeclaration time.Duration

// CopyBasicFieldsToAclassDB is used to copy basic fields between the Stage or the CRUD to the back repo
func (aclassDB *AclassDB) CopyBasicFieldsFromAclass(aclass *models.Aclass) {
	aclassDB.Name_Data.String = aclass.Name
	aclassDB.Name_Data.Valid = true

	aclassDB.Date_Data.Time = aclass.Date
	aclassDB.Date_Data.Valid = true

	aclassDB.Booleanfield_Data.Bool = aclass.Booleanfield
	aclassDB.Booleanfield_Data.Valid = true

	aclassDB.Aenum_Data.String = string(aclass.Aenum)
	aclassDB.Aenum_Data.Valid = true

	aclassDB.Aenum_2_Data.String = string(aclass.Aenum_2)
	aclassDB.Aenum_2_Data.Valid = true

	aclassDB.Benum_Data.String = string(aclass.Benum)
	aclassDB.Benum_Data.Valid = true

	aclassDB.CName_Data.String = aclass.CName
	aclassDB.CName_Data.Valid = true

	aclassDB.CFloatfield_Data.Float64 = aclass.CFloatfield
	aclassDB.CFloatfield_Data.Valid = true

	aclassDB.Floatfield_Data.Float64 = aclass.Floatfield
	aclassDB.Floatfield_Data.Valid = true

	aclassDB.Intfield_Data.Int64 = int64(aclass.Intfield)
	aclassDB.Intfield_Data.Valid = true

	aclassDB.Anotherbooleanfield_Data.Bool = aclass.Anotherbooleanfield
	aclassDB.Anotherbooleanfield_Data.Valid = true

	aclassDB.Duration1_Data.Int64 = int64(aclass.Duration1)
	aclassDB.Duration1_Data.Valid = true
}

// CopyBasicFieldsToAclassDB is used to copy basic fields between the Stage or the CRUD to the back repo
func (aclassDB *AclassDB) CopyBasicFieldsToAclass(aclass *models.Aclass) {
	// insertion point for fields value set from nullable fields
	if aclassDB.Name_Data.Valid {
		aclass.Name = aclassDB.Name_Data.String
	}

	if aclassDB.Date_Data.Valid {
		aclass.Date = aclassDB.Date_Data.Time
	}

	aclass.Booleanfield = aclassDB.Booleanfield_Data.Bool

	if aclassDB.Aenum_Data.Valid {
		aclass.Aenum = models.AEnumType(aclassDB.Aenum_Data.String)
	}

	if aclassDB.Aenum_2_Data.Valid {
		aclass.Aenum_2 = models.AEnumType(aclassDB.Aenum_2_Data.String)
	}

	if aclassDB.Benum_Data.Valid {
		aclass.Benum = models.BEnumType(aclassDB.Benum_Data.String)
	}

	if aclassDB.CName_Data.Valid {
		aclass.CName = aclassDB.CName_Data.String
	}

	if aclassDB.CFloatfield_Data.Valid {
		aclass.CFloatfield = aclassDB.CFloatfield_Data.Float64
	}

	if aclassDB.Floatfield_Data.Valid {
		aclass.Floatfield = aclassDB.Floatfield_Data.Float64
	}

	if aclassDB.Intfield_Data.Valid {
		aclass.Intfield = int(aclassDB.Intfield_Data.Int64)
	}

	aclass.Anotherbooleanfield = aclassDB.Anotherbooleanfield_Data.Bool

	if aclassDB.Duration1_Data.Valid {
		aclass.Duration1 = time.Duration(aclassDB.Duration1_Data.Int64)
	}
}
