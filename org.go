package iland

import (
	"encoding/json"
	"fmt"
)

type Org struct {
	client      *Client
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	CRM         string `json:"crm"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	VCloudHref  string `json:"vcloud_href"`
	LocationID  string `json:"location_id"`
	ZertoTarget bool   `json:"zerto_target"`
	Enabled     bool   `json:"enabled"`
	UpdatedDate int    `json:"updated_date"`
}

func (o Org) GetCatalogs() []Catalog {
	catalogs := []Catalog{}
	data, _ := o.client.Get(fmt.Sprintf("/org/%s/catalogs", o.UUID))
	json.Unmarshal(data, &catalogs)
	for i, catalog := range catalogs {
		catalog.client = o.client
		catalogs[i] = catalog
	}
	return catalogs
}

func (o Org) GetVdcs() []Vdc {
	vdcs := []Vdc{}
	data, _ := o.client.Get(fmt.Sprintf("/org/%s/vdcs", o.UUID))
	json.Unmarshal(data, &vdcs)
	for i, vdc := range vdcs {
		vdc.client = o.client
		vdcs[i] = vdc
	}
	return vdcs
}

func (o Org) GetEdges() []Edge {
	edges := []Edge{}
	data, _ := o.client.Get(fmt.Sprintf("/org/%s/edges", o.UUID))
	json.Unmarshal(data, &edges)
	for i, edge := range edges {
		edge.client = o.client
		edges[i] = edge
	}
	return edges
}

func (o Org) GetVdcNetworks() []VdcNetwork {
	vdcNetworks := []VdcNetwork{}
	data, _ := o.client.Get(fmt.Sprintf("/org/%s/vdc-networks", o.UUID))
	json.Unmarshal(data, &vdcNetworks)
	for i, vdcNetwork := range vdcNetworks {
		vdcNetwork.client = o.client
		vdcNetworks[i] = vdcNetwork
	}
	return vdcNetworks
}

func (o Org) GetVAppTemplates() []VAppTemplate {
	vAppTemplates := []VAppTemplate{}
	data, _ := o.client.Get(fmt.Sprintf("/org/%s/vapp-templates", o.UUID))
	json.Unmarshal(data, &vAppTemplates)
	for i, vAppTemplate := range vAppTemplates {
		vAppTemplate.client = o.client
		vAppTemplates[i] = vAppTemplate
	}
	return vAppTemplates
}

func (o Org) GetMedias() []Media {
	medias := []Media{}
	data, _ := o.client.Get(fmt.Sprintf("/org/%s/medias", o.UUID))
	json.Unmarshal(data, &medias)
	for i, media := range medias {
		media.client = o.client
		medias[i] = media
	}
	return medias
}

func (o Org) GetVApps() []VApp {
	vApps := []VApp{}
	data, _ := o.client.Get(fmt.Sprintf("/org/%s/vapps", o.UUID))
	json.Unmarshal(data, &vApps)
	for i, vApp := range vApps {
		vApp.client = o.client
		vApps[i] = vApp
	}
	return vApps
}

func (o Org) GetVAppNetworks() []VAppNetwork {
	vAppNetworks := []VAppNetwork{}
	data, _ := o.client.Get(fmt.Sprintf("/org/%s/vapp-networks", o.UUID))
	json.Unmarshal(data, &vAppNetworks)
	for i, vAppNetwork := range vAppNetworks {
		vAppNetwork.client = o.client
		vAppNetworks[i] = vAppNetwork
	}
	return vAppNetworks
}

func (o Org) GetVirtualMachines() []VirtualMachine {
	virtualMachines := []VirtualMachine{}
	data, _ := o.client.Get(fmt.Sprintf("/org/%s/vms", o.UUID))
	json.Unmarshal(data, &virtualMachines)
	for i, virtualMachine := range virtualMachines {
		virtualMachine.client = o.client
		virtualMachines[i] = virtualMachine
	}
	return virtualMachines
}

func (o Org) GetActiveTasks() []Task {
	tasks := []Task{}
	data, _ := o.client.Get(fmt.Sprintf("/task/%s/org/%s/active", o.LocationID, o.UUID))
	json.Unmarshal(data, &tasks)
	for i, task := range tasks {
		task.client = o.client
		tasks[i] = task
	}
	return tasks
}

func (o Org) GetDefaultStorageProfile() StorageProfile {
	for _, vdc := range o.GetVdcs() {
		for _, storageProfile := range vdc.GetStorageProfiles() {
			if storageProfile.Default {
				return storageProfile
			}
		}
	}
	return StorageProfile{}
}
