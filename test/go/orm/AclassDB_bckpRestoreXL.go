package orm

import (
	"log"
	"sort"

	"github.com/tealeg/xlsx/v3"
)

// Backup generates a json file from a slice of all AclassDB instances in the backrepo
func (backRepoAclass *BackRepoAclassStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*AclassDB, 0)
	for _, aclassDB := range *backRepoAclass.Map_AclassDBID_AclassDB {
		forBackup = append(forBackup, aclassDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Aclass")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	for _, aclassDB := range forBackup {

		row := sh.AddRow()
		row.WriteStruct(aclassDB, -1)
	}
}
