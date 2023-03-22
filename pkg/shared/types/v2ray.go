package types

type V2rayConfig struct {
	Log       string          `json:"url"`
	Outbounds []V2rayOutbound `json:"outbounds"`
}

type V2rayLog struct {
	Access string `json:"access"`
	Error  string `json:"error"`
}

type V2rayOutbound struct {
	Protocol    string                `json:"protocol"`
	SendThrough *string               `json:"sendThrough"`
	Tag         string                `json:"tag"`
	Settings    V2rayOutboundSettings `json:"settings"`
}

type V2rayOutboundSettings struct {
	Vnext []V2rayServerObject `json:"vnext"`
}

type V2rayServerObject struct {
	Address string            `json:"address"`
	Port    int               `json:"port"`
	Users   []V2rayUserObject `json:"users"`
}

type V2rayUserObject struct {
	ID       string `json:"id"`
	AlterID  int    `json:"alterId"`
	Security string `json:"security"`
	Level    string `json:"level"`
}
