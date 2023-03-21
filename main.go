package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"osub/pkg/parser"
	"osub/pkg/shared"
	"strings"
)

func main() {

	resp, err := http.Get("")

	if err != nil {
		fmt.Println("Error fetching subscription: ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error request subscription link: %v", err)
	}

	links, err := parser.Subscription(string(body))

	if err != nil {
		log.Fatalf("Error parsing Subscription link: %v", err)
	}

	for _, link := range *links {
		if strings.HasPrefix(link, shared.VMESS_PREFIX) {
			config, err := parser.Vmess(link)

			if err != nil {
				log.Fatalf("Error parsing Vmess link: %v", err)
			}

			fmt.Println(config)
		}
	}

	// v2rayConfig := generateV2rayConfig(config)

	// v2rayConfigJson, err := json.MarshalIndent(v2rayConfig, "", "  ")
	// if err != nil {
	// 	log.Fatalf("Error marshaling V2ray config: %v", err)
	// }
	// fmt.Println(string(v2rayConfigJson))
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
