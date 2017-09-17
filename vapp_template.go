package iland

import (
	"encoding/json"
	"fmt"
)

type VAppTemplate struct {
	client             *Client
	Name               string  `json:"name"`
	UUID               string  `json:"uuid"`
	Status             int     `json:"status"`
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

type VAppTemplateVirtualMachine struct {
	Name             string `json:"name"`
	UUID             string `json:"uuid"`
	LocationID       string `json:"location_id"`
	VdcUUID          string `json:"vdc_uuid"`
	VAppTemplateUUID string `json:"vapp_template_uuid"`
	SizeGB           int    `json:"size_gb"`
}

func (v VAppTemplate) Delete() (Task, error) {
	task := Task{}
	data, err := v.client.Delete(fmt.Sprintf("/vapp-template/%s", v.UUID))
	if err != nil {
		return task, err
	}
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}

func (v VAppTemplate) GetVirtualMachines() []VAppTemplateVirtualMachine {
	virtualMachines := []VAppTemplateVirtualMachine{}
	data, _ := v.client.Get(fmt.Sprintf("/vapp-template/%s/vms", v.UUID))
	json.Unmarshal(data, &virtualMachines)
	for i, virtualMachine := range virtualMachines {
		virtualMachine.LocationID = v.LocationID
		virtualMachine.VdcUUID = v.VdcUUID
		virtualMachine.VAppTemplateUUID = v.UUID
		virtualMachines[i] = virtualMachine
	}
	return virtualMachines
}

func (v VAppTemplate) Deploy(vdcUUID, NewVAppName string) (Task, error) {
	task := Task{}
	params := struct {
		VAppTemplateUUID string `json:"vapp_template_uuid"`
		Name             string `json:"name"`
	}{
		VAppTemplateUUID: v.UUID,
		Name:             NewVAppName,
	}
	vdc, err := v.client.GetVdc(vdcUUID)
	if err != nil {
		return task, err
	}
	output, _ := json.Marshal(&params)
	data, err := v.client.Post(fmt.Sprintf("/vdc/%s/vapp", vdc.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}

func (v VAppTemplate) Rename(newVAppTemplateName string) (Task, error) {
	task := Task{}
	params := struct {
		Name string `json:"name"`
	}{
		Name: newVAppTemplateName,
	}
	output, _ := json.Marshal(&params)
	data, err := v.client.Put(fmt.Sprintf("/vapp/%s", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}
