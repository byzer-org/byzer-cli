package utils

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/jsonq"
	"os"
	"strings"
)

func CreateTmpFile(content string) (*os.File, error) {
	//fsys := os.DirFS(".")
	tmpfile, _ := os.CreateTemp(".", "*")
	tmpfileName := tmpfile.Name()
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		return nil, fmt.Errorf("Fail to create tmp file [%s] ", tmpfileName)
	}
	return tmpfile, nil
}

func BuildJsonQueryFromStr(jsonStr string) *jsonq.JsonQuery {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	return jq

}
