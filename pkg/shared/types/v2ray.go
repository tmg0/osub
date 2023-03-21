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
	Protocol    string      `json:"protocol"`
	SendThrough string      `json:"sendThrough"`
	Tag         string      `json:"tag"`
	Settings    interface{} `json:"settings"`
}

type V2rayVmessOutbound struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	UUID    string `json:"uuid"`
}
