package shared

import (
	"encoding/json"
	"os"
	"osub/pkg/shared/types"
	"path/filepath"
)

func ReadConfig(args ...string) (config *types.OsubConfig, err error) {
	dir := ""
	if len(args) > 0 {
		dir = args[0]
	}

	if len(args) == 0 {
		dir, err = os.Executable()
		if err != nil {
			return nil, err
		}
	}

	path := filepath.Join(filepath.Dir(dir), CONFIG_FILE_NAME)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var conf = new(types.OsubConfig)
	err = json.NewDecoder(file).Decode(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
