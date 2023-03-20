package main

import (
	"encoding/json"
	"fmt"
	"log"
	"osub/pkg/parser"
)

func main() {
	// 定义 Vmess 链接
	link := "vmess://"

	// 解析 Vmess 链接
	config, err := parser.Vmess(link)
	if err != nil {
		log.Fatalf("Error parsing Vmess link: %v", err)
	}

	// 生成 V2Ray 配置
	v2rayConfig := generateV2rayConfig(config)

	// 输出 V2Ray 配置
	v2rayConfigJson, err := json.MarshalIndent(v2rayConfig, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling V2ray config: %v", err)
	}
	fmt.Println(string(v2rayConfigJson))
}

func generateV2rayConfig(vmess *parser.VmessConfig) string {
	var outbounds []map[string]interface{}

	outbound := map[string]interface{}{
		"protocol": "vmess",
		"settings": map[string]interface{}{
			"vnext": []map[string]interface{}{
				{
					"address": vmess.Add,
					"port":    vmess.Port,
					"users": []map[string]interface{}{
						{
							"id":      vmess.ID,
							"alterId": vmess.Aid,
						},
					},
				},
			},
		},
	}
	outbounds = append(outbounds, outbound)

	config := map[string]interface{}{
		"outbounds": outbounds,
	}
	configBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal v2ray config:", err)
		return ""
	}
	return string(configBytes)
}
