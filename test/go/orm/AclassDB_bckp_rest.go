package orm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/fullstack-lang/gong/test/go/models"
)

// Backup serialize the Aclass array into a json file (while nullng pointer before)
func (backRepoAclass *BackRepoAclassStruct) Backup(stage *models.StageStruct, dirPath string) (Error error) {

	// backup file name
	filename := filepath.Join(dirPath, "Aclass.json")

	file, err := json.MarshalIndent(backRepoAclass.Map_AclassDBID_AclassDB, "", " ")

	if err != nil {
		log.Panic("Error with Aclass json marshalling ", err.Error())
	}

	_ = ioutil.WriteFile(filename, file, 0644)

	return
}
