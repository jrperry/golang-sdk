package iland

type Alert struct {
	UUID     string        `json:"uuid"`
	Username string        `json:"username"`
	Enabled  bool          `json:"enabled"`
	Type     string        `json:"type"`
	Delivery AlertDelivery `json:"delivery"`
}

type AlertDelivery struct {
	PushEnabled  bool `json:"push_enabled"`
	InboxEnabled bool `json:"inbox_enabled"`
	EmailEnabled bool `json:"email_enabled"`
}
