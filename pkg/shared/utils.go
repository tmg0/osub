package shared

import (
	"os"
	"osub/pkg/shared/types"
	"path/filepath"

	"github.com/spf13/viper"
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

	viper.AddConfigPath(filepath.Dir(dir))
	viper.AddConfigPath(".")
	viper.SetConfigName(CONFIG_FILE_NAME)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var conf *types.OsubConfig

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return conf, nil
}
