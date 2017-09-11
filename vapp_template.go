package iland

import (
	"encoding/json"
	"fmt"
)

type VAppTemplate struct {
	client             *Client
	Name               string  `json:"name"`
	UUID               string  `json:"uuid"`
	Description        string  `json:"description"`
	SizeGB             float64 `json:"size"`
	GoldMaster         bool    `json:"gold_master"`
	Public             bool    `json:"public"`
	IsExpired          bool    `json:"is_expired"`
	LocationID         string  `json:"location_id"`
	OrgUUID            string  `json:"org_uuid"`
	VdcUUID            string  `json:"vdc_uuid"`
	CatalogUUID        string  `json:"catalog_uuid"`
	StorageProfileUUID string  `json:"storage_profile_uuid"`
	VCloudHref         string  `json:"vcloud_href"`
	CreatedDate        int     `json:"created_date"`
	UpdatedDate        int     `json:"updated_date"`
}

func (v VAppTemplate) GetVirtualMachines() []VirtualMachine {
	virtualMachines := []VirtualMachine{}
	data, _ := v.client.Get(fmt.Sprintf("/vapp-template/%s/vms", v.UUID))
	json.Unmarshal(data, &virtualMachines)
	for i, virtualMachine := range virtualMachines {
		virtualMachine.client = v.client
		virtualMachines[i] = virtualMachine
	}
	return virtualMachines
}
