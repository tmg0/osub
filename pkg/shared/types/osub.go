package types

type OsubSubscription struct {
	URL string `json:"url"`
}

type OsubConfig struct {
	Interval      string              `json:"interval"`
	Subscriptions []*OsubSubscription `json:"subscriptions"`
	V2ray         *OsubV2rayConfig    `json:"v2ray"`
	Clash         *OsubClashConfig    `json:"clash"`
}

type OsubServerConfig struct {
	Type     string
	Address  string
	Port     int
	Password string
	Method   *string
	UUID     *string
	AlterID  *string
	Remarks  *string
	SNI      *string
}

type OsubV2rayConfig struct {
	Config *string `json:"config"`
}

type OsubClashConfig struct {
	Config *string `json:"config"`
}
