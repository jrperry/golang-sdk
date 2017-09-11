package goiland

type StorageProfile struct {
	Name          string `json:"name"`
	UUID          string `json:"uuid"`
	VdcUUID       string `json:"vdc_uuid"`
	Enabled       bool   `json:"enabled"`
	Default       bool   `json:"default"`
	StorageUnit   string `json:"unit"`
	StorageLimit  int    `json:"size_limit"`
	StorageUsedMB int    `json:"storage_used_in_mb"`
	VCloudHref    string `json:"href"`
	UpdatedDate   int    `json:"updated_date"`
}
