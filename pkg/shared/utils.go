package shared

import (
	"encoding/json"
	"os"
	"osub/pkg/shared/types"
	"path/filepath"
)

func OptionalArg[T any](args []T) *T {
	if len(args) == 1 {
		return &args[0]
	}

	return nil
}

func WithDefault[T any](arg *T, defaultValue func() (*T, error)) (T, error) {
	if arg == nil {
		res, err := defaultValue()
		if err != nil {
			return *arg, err
		}
		return *res, nil
	}
	return *arg, nil
}

func Cwd() (*string, error) {
	dir, err := os.Executable()
	if err != nil {
		return nil, err
	}
	return &dir, nil
}

func ReadConfig(args ...string) (config *types.OsubConfig, err error) {

	dir, err := WithDefault(OptionalArg(args), Cwd)
	if err != nil {
		return nil, err
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
