package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func Load(mode, srcPath string, cfg interface{}) error {
	cfgFileName := fmt.Sprintf("application.%s.json", mode)
	cfgFilePath := path.Join(srcPath, "conf", cfgFileName)
	if _, err := os.Stat(cfgFilePath); err != nil {
		cfgFilePath = path.Join(srcPath, "conf", "application.json")
	}

	data, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &cfg)
}
