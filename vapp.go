package iland

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"
)

type VApp struct {
	client              *Client
	Name                string   `json:"name"`
	UUID                string   `json:"uuid"`
	Description         string   `json:"description"`
	Deployed            bool     `json:"deployed"`
	Expired             bool     `json:"is_expired"`
	Status              string   `json:"status"`
	LocationID          string   `json:"location_id"`
	OrgUUID             string   `json:"org_uuid"`
	VdcUUID             string   `json:"vdc_uuid"`
	StorageProfileUUIDs []string `json:"storage_profiles"`
	VCloudHref          string   `json:"vcloud_href"`
	CreatedDate         int      `json:"created_date"`
	UpdatedDate         int      `json:"updated_date"`
}

func (v VApp) GetVirtualMachines() []VirtualMachine {
	virtualMachines := []VirtualMachine{}
	data, _ := v.client.Get(fmt.Sprintf("/vapp/%s/vms", v.UUID))
	json.Unmarshal(data, &virtualMachines)
	for i, virtualMachine := range virtualMachines {
		virtualMachine.client = v.client
		virtualMachines[i] = virtualMachine
	}
	return virtualMachines
}

func (v VApp) GetVAppNetworks() []VAppNetwork {
	vAppNetworks := []VAppNetwork{}
	data, _ := v.client.Get(fmt.Sprintf("/vapp/%s/networks", v.UUID))
	json.Unmarshal(data, &vAppNetworks)
	for i, vAppNetwork := range vAppNetworks {
		vAppNetwork.client = v.client
		vAppNetworks[i] = vAppNetwork
	}
	return vAppNetworks
}

func (v VApp) Delete() (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	data, err := v.client.Delete(fmt.Sprintf("/vapp/%s", v.UUID))
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VApp) PowerOn() (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/poweron", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VApp) PowerOff() (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/poweroff", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VApp) Suspend() (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/suspend", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VApp) Rename(newVAppName string) (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	params := struct {
		Name string `json:"name"`
	}{
		Name: newVAppName,
	}
	output, _ := json.Marshal(&params)
	data, err := v.client.Put(fmt.Sprintf("/vapp/%s/name", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VApp) HasSnapshot() bool {
	check := struct {
		HasSnapshot bool `json:"has_snapshot"`
	}{}
	data, err := v.client.Get(fmt.Sprintf("/vapp/%s/snapshot/check", v.UUID))
	if err != nil {
		return false
	}
	json.Unmarshal(data, &check)
	return check.HasSnapshot
}

func (v VApp) GetSnapshot() (Snapshot, error) {
	snapshot := Snapshot{}
	data, err := v.client.Get(fmt.Sprintf("/vapp/%s/snapshot", v.UUID))
	if err != nil {
		return snapshot, err
	}
	err = json.Unmarshal(data, &snapshot)
	return snapshot, err
}

func (v VApp) TakeSnapshot() (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	params := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Memory      bool   `json:"memory"`
		Quiesce     bool   `json:"quiesce"`
	}{
		Name:    time.Now().UTC().String(),
		Memory:  false,
		Quiesce: false,
	}
	output, _ := json.Marshal(&params)
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/snapshot", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VApp) RevertSnapshot() (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/snapshot/restore", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VApp) Clone(targetVdcUUID, newVAppName string) (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	params := struct {
		Name string `json:"name"`
	}{
		Name: newVAppName,
	}
	output, _ := json.Marshal(&params)
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/copy/%s", v.UUID, targetVdcUUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

type AddVAppNetworkParams struct {
	Name         string    `json:"name"`
	Gateway      string    `json:"gateway_address"`
	Netmask      string    `json:"network_mask"`
	PrimaryDNS   string    `json:"primary_dns"`
	SecondaryDNS string    `json:"secondary_dns"`
	IPRanges     []IPRange `json:"ip_ranges"`
}

func (v VApp) AddVAppNetwork(params AddVAppNetworkParams) (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	gateway := net.ParseIP(params.Gateway)
	if gateway == nil {
		return task, errors.New("invalid gateway address")
	}
	netmask := net.ParseIP(params.Netmask)
	if netmask == nil {
		return task, errors.New("invalid netmask")
	}
	for _, ipRange := range params.IPRanges {
		startAddress := net.ParseIP(ipRange.Start)
		if startAddress == nil {
			return task, errors.New("invalid ip range start address")
		}
		endAddress := net.ParseIP(ipRange.End)
		if endAddress == nil {
			return task, errors.New("invalid ip range end address")
		}
	}
	output, _ := json.Marshal(&params)
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/vapp-network", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VApp) RemoveNetwork(vAppNetworkUUID string) (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	data, err := v.client.Delete(fmt.Sprintf("/vapp/%s/network/%s", v.UUID, vAppNetworkUUID))
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

type AddVirtualMachineFromVAppTemplateParams struct {
	NewVirtualMachineName    string
	SourceVAppTemplateUUID   string
	SourceVirtualMachineUUID string
}

type addVirtualMachinesFromVAppTemplateParams struct {
	NewVirtualMachineName    string `json:"name"`
	SourceVAppTemplateUUID   string `json:"vapp_template_uuid"`
	SourceVirtualMachineUUID string `json:"vm_template_uuid"`
	IPAddressMode            string `json:"ip_address_mode"`
	NetworkUUID              string `json:"network_uuid"`
	IPAddress                string `json:"ip_address"`
}

func (v VApp) AddVirtualMachinesFromVAppTemplates(params []AddVirtualMachineFromVAppTemplateParams) (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	virtualMachineParams := []addVirtualMachinesFromVAppTemplateParams{}
	networks := v.GetVAppNetworks()
	for _, param := range params {
		virtualMachineParam := addVirtualMachinesFromVAppTemplateParams{
			NewVirtualMachineName:    param.NewVirtualMachineName,
			SourceVAppTemplateUUID:   param.SourceVAppTemplateUUID,
			SourceVirtualMachineUUID: param.SourceVirtualMachineUUID,
			IPAddressMode:            "DHCP",
			NetworkUUID:              networks[0].UUID,
			IPAddress:                "",
		}
		virtualMachineParams = append(virtualMachineParams, virtualMachineParam)
	}
	task := Task{}
	output, _ := json.Marshal(&virtualMachineParams)
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/vms", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VApp) GetPerformance(start, end time.Time, perfInterval string, metric PerfMetric) (PerfResults, error) {
	results := PerfResults{}
	limit := getPerfLimit(perfInterval)
	queryParams := fmt.Sprintf("?group=%s&name=%s&type=%s&start=%d&end=%d&interval=%s&limit=%s", metric.Group, metric.Name, metric.Type, getUnixMilliseconds(start), getUnixMilliseconds(end), perfInterval, limit)
	data, err := v.client.Get(fmt.Sprintf("/vapp/%s/p%s", v.UUID, queryParams))
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(data, &results)
	return results, err
}

func (v VApp) GetCurrentBill() (BillingSummary, error) {
	billing := BillingSummary{}
	data, err := v.client.Get(fmt.Sprintf("/vapp/%s/bill", v.UUID))
	if err != nil {
		return billing, err
	}
	err = json.Unmarshal(data, &billing)
	return billing, err
}

func (v VApp) GetPrevBill(month, year int) (BillingSummary, error) {
	billing := BillingSummary{}
	queryParams := fmt.Sprintf("?month=%d&year=%d", month, year)
	data, err := v.client.Get(fmt.Sprintf("/vapp/%s/bill%s", v.UUID, queryParams))
	if err != nil {
		return billing, err
	}
	err = json.Unmarshal(data, &billing)
	return billing, err
}
