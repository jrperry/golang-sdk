package iland

type Media struct {
	client             *Client
	Name               string  `json:"name"`
	UUID               string  `json:"uuid"`
	Description        string  `json:"description"`
	Public             bool    `json:"public"`
	SizeGB             float64 `json:"size"`
	LocationID         string  `json:"location_id"`
	OrgUUID            string  `json:"org_uuid"`
	VdcUUID            string  `json:"vdc_uuid"`
	CatalogUUID        string  `json:"catalog_uuid"`
	StorageProfileUUID string  `json:"storage_profile_uuid"`
	VCloudHref         string  `json:"vcloud_href"`
	CreatedDate        int     `json:"created_date"`
	UpdatedDate        int     `json:"updated_date"`
}
