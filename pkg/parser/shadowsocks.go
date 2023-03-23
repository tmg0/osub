package parser

import (
	"encoding/base64"
	"fmt"
	"osub/pkg/shared"
	"osub/pkg/shared/types"
	"regexp"
	"strconv"
	"strings"
)

func Shadowsocks(link string) (*types.OsubServerConfig, error) {
	var conf types.OsubServerConfig

	// Remove vmess prefix
	link = strings.TrimPrefix(link, shared.SS_PREFIX)
	hashHost := strings.Split(link, "#")

	encode := hashHost[0]

	// Decode Base64
	data, err := base64.RawStdEncoding.DecodeString(encode)
	if err != nil {
		return nil, err
	}

	// Parse JSON string to SsConfig struct
	pattern := regexp.MustCompile(`^(.*?):(.*?)@(.*?):(\d+)$`)
	matches := pattern.FindStringSubmatch(string(data))

	if len(matches) != 5 {
		return nil, fmt.Errorf("invalidate ss link")
	}

	add := matches[3]

	if len(hashHost) > 1 {
		add = hashHost[1]
	}

	conf.Type = shared.SHADOWSOCKS
	conf.Method = &matches[1]
	conf.Password = matches[2]
	conf.Address = add
	conf.Port, err = strconv.Atoi(matches[4])

	if err != nil {
		return nil, err
	}

	return &conf, nil
}
