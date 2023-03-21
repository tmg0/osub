package parser

import (
	"encoding/base64"
	"strings"
)

func Subscription(encodeStr string) (*[]string, error) {
	bytes, err := base64.StdEncoding.DecodeString(encodeStr)

	if err != nil {
		return nil, err
	}

	decode := string(bytes)
	arr := strings.Split(decode, "\n")

	return &arr, nil
}
