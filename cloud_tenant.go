package iland

type CloudTenant struct {
	UUID                 string               `json:"uuid"`
	Name                 string               `json:"name"`
	UID                  string               `json:"uid"`
	CRM                  string               `json:"crm"`
	OwnerName            string               `json:"owner_name"`
	ContractUUID         string               `json:"contract_uuid"`
	LocationID           string               `json:"location_id"`
	Enabled              bool                 `json:"enabled"`
	Resources            CloudTenantResources `json:"resources"`
	LastResult           string               `json:"last_result"`
	LastActive           int                  `json:"last_active"`
	ThrottlingEnabled    bool                 `json:"throttling_enabled"`
	ThrottlingSpeedLimit int                  `json:"throttling_speed_limit"`
	ThrottlingSpeedUnit  string               `json:"throttling_speed_unit"`
	PublicIPCount        int                  `json:"public_ip_count"`
	BackupCount          int                  `json:"backup_count"`
	Deleted              bool                 `json:"deleted"`
	UpdatedDate          int                  `json:"updated_date"`
	DeletedDate          int                  `json:"deleted_date"`
}

type CloudTenantResources struct {
	Resource []CloudTenantResource `json:"resources"`
}

type CloudTenantResource struct {
	Repository CloudTenantRepository `json:"repository"`
}

type CloudTenantRepository struct {
	Name        string `json:"display_name"`
	QuotaMB     int    `json:"quota"`
	UsedQuotaMB int    `json:"used_quota"`
}
