package iland

import (
	"encoding/json"
	"fmt"
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
	task := Task{}
	data, err := v.client.Delete(fmt.Sprintf("/vapp/%s", v.UUID))
	if err != nil {
		return task, err
	}
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}

func (v VApp) PowerOn() (Task, error) {
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/poweron", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}

func (v VApp) PowerOff() (Task, error) {
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/poweroff", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}

func (v VApp) Suspend() (Task, error) {
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/suspend", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}

func (v VApp) Rename(newVAppName string) (Task, error) {
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
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}

func (v VApp) TakeSnapshot() (Task, error) {
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
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}

func (v VApp) RevertSnapshot() (Task, error) {
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/snapshot/restore", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}

func (v VApp) Clone(targetVdcUUID, newVAppName string) (Task, error) {
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
	err = json.Unmarshal([]byte(data), &task)
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
	task := Task{}
	output, _ := json.Marshal(&params)
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/vapp-network", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}

func (v VApp) RemoveNetwork(vAppNetworkUUID string) (Task, error) {
	task := Task{}
	data, err := v.client.Delete(fmt.Sprintf("/vapp/%s/network/%s", v.UUID, vAppNetworkUUID))
	if err != nil {
		return task, err
	}
	err = json.Unmarshal([]byte(data), &task)
	task.client = v.client
	return task, err
}

type AddVirtualMachineFromVAppTemplateParams struct {
	NewVirtualMachineName    string `json:"name"`
	SourceVAppTemplateUUID   string `json:"vapp_template_uuid"`
	SourceVirtualMachineUUID string `json:"vm_template_uuid"`
}

func (v VApp) AddVirtualMachinesFromVAppTemplates(params []AddVirtualMachineFromVAppTemplateParams) (Task, error) {
	task := Task{}
	output, _ := json.Marshal(&params)
	data, err := v.client.Post(fmt.Sprintf("/vapp/%s/vms", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal([]byte(data), &task)
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
