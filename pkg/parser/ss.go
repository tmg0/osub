package parser

import (
	"encoding/base64"
	"fmt"
	"osub/pkg/shared"
	"regexp"
	"strings"
)

type SsConfig struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Method   string `json:"method"`
}

func Ss(link string) (*SsConfig, error) {
	var ssConfig SsConfig

	// Remove vmess prefix
	link = strings.TrimPrefix(link, shared.SS_PREFIX)
	link = strings.Split(link, "#")[0]

	// Decode Base64
	data, err := base64.RawStdEncoding.DecodeString(link)
	if err != nil {
		return nil, err
	}

	// Parse JSON string to SsConfig struct
	pattern := regexp.MustCompile(`^(.*?):(.*?)@(.*?):(\d+)$`)
	matches := pattern.FindStringSubmatch(string(data))

	if len(matches) != 5 {
		return nil, fmt.Errorf("invalidate ss link")
	}

	ssConfig.Method = matches[1]
	ssConfig.Password = matches[2]
	ssConfig.Server = matches[3]
	ssConfig.Port = matches[4]

	return &ssConfig, nil
}
