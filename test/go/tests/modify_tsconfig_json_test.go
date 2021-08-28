package tests

import (
	"path/filepath"
	"testing"

	"encoding/json"
	"io/ioutil"
)

// TestModifyJSON
//
// This test covers a a gong functionality where the Stage is updated
// through a callback that is defined in the "models" package
//
//
type User struct {
	Name string
	Age  int8
}

func TestModifyJSON(t *testing.T) {

	HandleJson(filepath.Join("../../../ng", "tsconfig.json"), "tsconfig.json")
	// user := User{
	// 	Name: "tanggu",
	// 	Age:  18,
	// }
	// var jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary

	// //  serialization
	// data, err := jsoniter.Marshal(&user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(data))

	// //  deserialization
	// var people User
	// err = jsoniter.Unmarshal(data, &people)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(people)
}

func HandleJson(jsonFile string, outFile string) error {
	// Read json buffer from jsonFile
	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	// We have known the outer json object is a map, so we define result as map.
	// otherwise, result could be defined as slice if outer is an array
	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return err
	}

	// handle peers
	nodesRaw := result["compilerOptions"]
	nodes := nodesRaw.(map[string]interface{})
	for _, node := range nodes {
		m := node.(map[string]interface{})
		if name, exists := m["name"]; exists {
			if name == "node1" {
				m["location"] = "new-value1"
			} else if name == "node2" {
				m["location"] = "new-value2"
			}
		}
	}

	// Convert golang object back to byte
	byteValue, err = json.Marshal(result)
	if err != nil {
		return err
	}

	// Write back to file
	err = ioutil.WriteFile(outFile, byteValue, 0644)
	return err
}
