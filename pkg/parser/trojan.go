package parser

type TrojanConfig struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Sni      string `json:"sni"`
}

func Trojan(link string) (*TrojanConfig, error) {

	return &TrojanConfig{
		Address:  "2",
		Port:     123,
		Password: "123",
		Sni:      "string",
	}, nil
}
