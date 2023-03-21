package shared

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func ReadConfig() (*OsubConfig, error) {
	path, err := os.Executable()

	if err != nil {
		return nil, err
	}

	path = filepath.Join(filepath.Dir(path), CONFIG_FILE_NAME)

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var conf OsubConfig

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
