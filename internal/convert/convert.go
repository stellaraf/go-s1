package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/stellaraf/go-utils"
)

const SPEC2_FILE string = "spec.oapi2.json"
const SPEC3_FILE string = "spec.oapi3.json"

func getSpec() (string, error) {
	root, err := utils.FindProjectRoot(4)
	if err != nil {
		return "", err
	}
	specPath := filepath.Join(root, SPEC2_FILE)
	return specPath, nil
}

func main() {
	specPath, err := getSpec()
	if err != nil {
		panic(err)
	}
	specFile, err := os.ReadFile(specPath)
	if err != nil {
		panic(err)
	}
	var spec2 *openapi2.T
	if err = json.Unmarshal(specFile, &spec2); err != nil {
		panic(err)
	}
	spec3, err := openapi2conv.ToV3(spec2)
	if err != nil {
		panic(err)
	}
	paths := spec3.Paths.Map()
	for k, v := range paths {
		ops := v.Operations()
		for ok, ov := range ops {
			opID := ov.OperationID
			opID = strings.ReplaceAll(opID, "_web_api_", "")
			ov.OperationID = opID
			paths[k].SetOperation(ok, ov)
		}
		spec3.Paths.Set(k, paths[k])
	}
	root, err := utils.FindProjectRoot(4)
	if err != nil {
		panic(err)
	}
	spec3Bytes, err := json.Marshal(&spec3)
	if err != nil {
		panic(err)
	}

	spec3Path := filepath.Join(root, SPEC3_FILE)
	spec3File, err := os.Create(spec3Path)
	if err != nil {
		panic(err)
	}
	_, err = spec3File.Write(spec3Bytes)
	if err != nil {
		panic(err)
	}
}
