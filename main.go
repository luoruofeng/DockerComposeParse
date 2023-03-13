package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/docker/cli/cli/compose/loader"
	"github.com/docker/cli/cli/compose/types"
)

func YamlToMap(filePath string) (map[string]interface{}, error) {
	// 读取docker-compose.yml文件
	composeFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open compose file: %v", err)
	}
	defer composeFile.Close()

	composeBytes, err := ioutil.ReadAll(composeFile)
	if err != nil {
		log.Fatalf("Failed to read compose file: %v", err)
	}

	// 解析docker-compose.yml文件
	return loader.ParseYAML(composeBytes)
}

func buildConfigDetails(source map[string]interface{}, env map[string]string) types.ConfigDetails {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return types.ConfigDetails{
		WorkingDir: path.Join(workingDir, "demo"),
		ConfigFiles: []types.ConfigFile{
			{Filename: "compose.yaml", Config: source},
		},
		Environment: env,
	}
}

func MapToConfig(dict map[string]interface{}) (*types.Config, error) {
	env := map[string]string{
		"ENV_VAR_1": "value_1",
		"ENV_VAR_2": "value_2",
	}
	details := buildConfigDetails(dict, env)

	// Create a new Compose file loader.
	return loader.Load(details)
}

func main() {
	m, err := YamlToMap("demo/compose.yaml")
	if err != nil {
		panic(err)
	}
	config, err := MapToConfig(m)
	if err != nil {
		panic(err)
	}
	//config includ all infomation about compose file: services, network ,volume.....
	fmt.Println(config)
}
