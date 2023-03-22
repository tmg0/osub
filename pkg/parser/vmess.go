package parser

import (
	"encoding/base64"
	"encoding/json"
	"osub/pkg/shared"
	"osub/pkg/shared/types"
	"strconv"
	"strings"
)

type VmessConfig struct {
	Ps   string `json:"ps"`
	Add  string `json:"add"`
	Port string `json:"port"`
	ID   string `json:"id"`
	Aid  int    `json:"aid"`
	Net  string `json:"net"`
	Type string `json:"type"`
	Tls  string `json:"tls"`
}

func Vmess(link string) (*types.OsubServerConfig, error) {
	var vmessConfig VmessConfig
	var conf types.OsubServerConfig

	// Remove vmess prefix
	link = strings.TrimPrefix(link, shared.VMESS_PREFIX)

	// Decode Base64
	data, err := base64.RawStdEncoding.DecodeString(link)
	if err != nil {
		return nil, err
	}

	// Parse JSON string to VmessConfig struct
	err = json.Unmarshal([]byte(data), &vmessConfig)
	if err != nil {
		return nil, err
	}

	conf.Type = shared.VMESS
	conf.Address = vmessConfig.Add
	conf.Port, err = strconv.Atoi(vmessConfig.Port)
	conf.Password = vmessConfig.ID

	if err != nil {
		return nil, err
	}

	return &conf, nil
}
