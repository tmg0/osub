package types

type OsubSubscription struct {
	URL string `json:"url"`
}

type OsubConfig struct {
	Interval      string              `json:"interval"`
	Subscriptions []*OsubSubscription `json:"subscriptions"`
}
