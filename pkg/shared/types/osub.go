package types

type OsubSubscription struct {
	URL string `json:"url"`
}

type OsubConfig struct {
	Interval      string              `json:"interval"`
	Subscriptions []*OsubSubscription `json:"subscriptions"`
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
