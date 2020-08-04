package yaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ReadYaml(configPath string, cfg interface{}) error {
	f, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error on ReadConfig: %s", err)
	}
	if f == nil {
		return fmt.Errorf("error on ReadConfig: file is nil")
	}

	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		return fmt.Errorf("error on ReadConfig.Unmarshal: %v", err)
	}
	return nil
}