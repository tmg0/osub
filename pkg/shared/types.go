package shared

type V2rayConfig struct {
	LogLevel  string          `json:"loglevel"`
	Inbounds  []V2rayInbound  `json:"inbounds"`
	Outbounds []V2rayOutbound `json:"outbounds"`
	Routing   V2rayRouting    `json:"routing"`
}

type V2rayInbound struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

type V2rayOutbound struct {
	Protocol string      `json:"protocol"`
	Settings interface{} `json:"settings"`
}

type V2rayRouting struct {
	Rules []V2rayRule `json:"rules"`
}

type V2rayRule struct {
	Type        string `json:"type"`
	InboundTag  string `json:"inboundTag"`
	OutboundTag string `json:"outboundTag"`
}

type OsubSubscription struct {
	URL string `json:"url"`
}

type OsubConfig struct {
	Interval      string              `json:"interval"`
	Subscriptions []*OsubSubscription `json:"subscriptions"`
}
