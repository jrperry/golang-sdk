package iland

type VAppNetwork struct {
	client            *Client
	Name              string    `json:"name"`
	UUID              string    `json:"uuid"`
	Description       string    `json:"description"`
	Enabled           bool      `json:"enabled"`
	Inherited         bool      `json:"inherited"`
	Shared            bool      `json:"shared"`
	VAppNetwork       bool      `json:"vapp_network"`
	FenceMode         string    `json:"fence_mode"`
	Gateway           string    `json:"gateway"`
	Netmask           string    `json:"netmask"`
	PrimaryDNS        string    `json:"primary_dns"`
	SecondaryDNS      string    `json:"secondary_dns"`
	DNSSuffix         string    `json:"dns_suffix"`
	IPRanges          []IPRange `json:"ip_ranges"`
	ParentNetworkName string    `json:"parent_network_name"`
	ParentNetworkUUID string    `json:"parent_network_uuid"`
	ParentEntityUUID  string    `json:"parent_entity_uuid"`
	EdgeUUID          string    `json:"edge_uuid"`
	RouterExternalIP  string    `json:"router_external_ip"`
}
