package iland

import (
	"encoding/json"
	"fmt"
)

type Vdc struct {
	client             *Client
	Name               string `json:"name"`
	UUID               string `json:"uuid"`
	Enabled            bool   `json:"enabled"`
	Description        string `json:"description"`
	OrgUUID            string `json:"org_uuid"`
	LocationID         string `json:"location_id"`
	AllocationModel    string `json:"allocation_model"`
	ReservedCPU        int    `json:"reserved_cpu"`
	AllocatedCPU       int    `json:"alloc_cpu"`
	ReservedMemory     int    `json:"reserved_mem"`
	AllocatedMemory    int    `json:"alloc_mem"`
	StorageLimit       int    `json:"disk_limit"`
	MaxHardwareVersion string `json:"max_hdw_version"`
	NetworkQuota       int    `json:"network_quota"`
	UsedNetworkCount   int    `json:"used_network_count"`
	VCloudHref         string `json:"vcloud_href"`
	UpdatedDate        int    `json:"updated_date"`
}

func (v Vdc) GetEdges() []Edge {
	edges := []Edge{}
	data, _ := v.client.Get(fmt.Sprintf("/vdc/%s/edges", v.UUID))
	json.Unmarshal(data, &edges)
	for i, edge := range edges {
		edge.client = v.client
		edges[i] = edge
	}
	return edges
}

func (v Vdc) GetStorageProfiles() []StorageProfile {
	storageProfiles := []StorageProfile{}
	data, _ := v.client.Get(fmt.Sprintf("/vdc/%s/storage-profiles", v.UUID))
	json.Unmarshal([]byte(data), &storageProfiles)
	return storageProfiles
}

func (v Vdc) GetVdcNetworks() []VdcNetwork {
	vdcNetworks := []VdcNetwork{}
	data, _ := v.client.Get(fmt.Sprintf("/vdc/%s/networks", v.UUID))
	json.Unmarshal([]byte(data), &vdcNetworks)
	for i, vdcNetwork := range vdcNetworks {
		vdcNetwork.client = v.client
		vdcNetworks[i] = vdcNetwork
	}
	return vdcNetworks
}

func (v Vdc) GetVApps() []VApp {
	vApps := []VApp{}
	data, _ := v.client.Get(fmt.Sprintf("/vdc/%s/vapps", v.UUID))
	json.Unmarshal([]byte(data), &vApps)
	for i, vApp := range vApps {
		vApp.client = v.client
		vApps[i] = vApp
	}
	return vApps
}

func (v Vdc) GetVirtualMachines() []VirtualMachine {
	virtualMachines := []VirtualMachine{}
	data, _ := v.client.Get(fmt.Sprintf("/org/%s/vms", v.UUID))
	json.Unmarshal([]byte(data), &virtualMachines)
	for i, virtualMachine := range virtualMachines {
		virtualMachine.client = v.client
		virtualMachines[i] = virtualMachine
	}
	return virtualMachines
}
