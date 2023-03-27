package shared

import (
	"fmt"
	"os"
	"osub/pkg/shared/types"
	"path/filepath"

	"github.com/spf13/viper"
)

var OSUB_CONFIG = viper.New()
var V2RAY_CONFIG = viper.New()

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

	OSUB_CONFIG.AddConfigPath(filepath.Dir(dir))
	OSUB_CONFIG.AddConfigPath(".")
	OSUB_CONFIG.SetConfigName(CONFIG_FILE_NAME)
	OSUB_CONFIG.SetConfigType("json")

	if err := OSUB_CONFIG.ReadInConfig(); err != nil {
		return nil, err
	}

	var conf *types.OsubConfig

	if err := OSUB_CONFIG.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func ReadV2rayConfig(v2ray *types.OsubV2rayConfig) (*types.V2rayConfig, error) {
	if v2ray == nil {
		return nil, fmt.Errorf("invalidate v2ray config file path")
	}

	if v2ray.Config == nil {
		return nil, fmt.Errorf("invalidate v2ray config file path")
	}

	V2RAY_CONFIG.AddConfigPath("../")
	V2RAY_CONFIG.SetConfigFile("v2ray.config")
	V2RAY_CONFIG.SetConfigType("json")

	if err := V2RAY_CONFIG.ReadInConfig(); err != nil {
		return nil, err
	}

	var conf *types.V2rayConfig

	if err := V2RAY_CONFIG.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return conf, nil
}
