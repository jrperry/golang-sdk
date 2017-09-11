package iland

type VdcNetwork struct {
	client            *Client
	Name              string    `json:"name"`
	UUID              string    `json:"uuid"`
	Description       string    `json:"description"`
	Gateway           string    `json:"gateway"`
	Netmask           string    `json:"netmask"`
	IPRanges          []IPRange `json:"ip_ranges"`
	PrimaryDNS        string    `json:"primary_dns"`
	SecondaryDNS      string    `json:"secondary_dns"`
	DNSSuffix         string    `json:"dns_suffix"`
	FenceMode         string    `json:"fence_mode"`
	LocationID        string    `json:"location_id"`
	OrgUUID           string    `json:"org_uuid"`
	VdcUUID           string    `json:"vdc_uuid"`
	EdgeUUID          string    `json:"edge_uuid"`
	ParentNetworkUUID string    `json:"parent_network_uuid"`
	Shared            bool      `json:"shared"`
	Inherited         bool      `json:"inherited"`
	UpdatedDate       int       `json:"updated_date"`
}
