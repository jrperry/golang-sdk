package iland

import (
	"encoding/json"
	"fmt"
)

type VirtualMachine struct {
	client              *Client
	Name                string   `json:"name"`
	UUID                string   `json:"uuid"`
	Description         string   `json:"description"`
	OperatingSystem     string   `json:"os"`
	VCPU                int      `json:"cpus_number"`
	CoresPerSocket      int      `json:"cores_per_socket"`
	MemoryMB            int      `json:"memory_size"`
	StorageProfileUUIDs []string `json:"storage_profiles"`
	HardwareVersion     string   `json:"hardware_version"`
	MediaMounted        bool     `json:"media_inserted"`
	MountedMediaName    string   `json:"inserted_media_name"`
	Deployed            bool     `json:"deployed"`
	Status              string   `json:"status"`
	LocationID          string   `json:"location_id"`
	OrgUUID             string   `json:"org_uuid"`
	VdcUUID             string   `json:"vdc_uuid"`
	VAppUUID            string   `json:"vapp_uuid"`
	VCloudHref          string   `json:"vcloud_href"`
	CreatedDate         int      `json:"created_date"`
	UpdatedDate         int      `json:"updated_date"`
}

type Disk struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	Type string `json:"type"`
}

type Nic struct {
	Index            int    `json:"vnic_id"`
	IPAddress        string `json:"ip_addr"`
	MacAddress       string `json:"mac_address"`
	IPAllocationMode string `json:"address_mode"`
	Primary          bool   `json:"primary_cnx"`
	Connected        bool   `json:"connected"`
	AdapterType      string `json:"adapter_type"`
	NetworkName      string `json:"net_name"`
}

func (v VirtualMachine) GetDisks() []Disk {
	disks := []Disk{}
	data, _ := v.client.Get(fmt.Sprintf("/vm/%s/virtual-disks", v.UUID))
	json.Unmarshal(data, &disks)
	return disks
}

func (v VirtualMachine) GetNics() []Nic {
	nics := []Nic{}
	data, _ := v.client.Get(fmt.Sprintf("/vm/%s/vnics", v.UUID))
	json.Unmarshal(data, &nics)
	return nics
}

type VMwareTools struct {
	Status        string `json:"status"`
	RunningStatus string `json:"running_status"`
	Version       string `json:"version"`
}

func (v VirtualMachine) GetTools() VMwareTools {
	tools := VMwareTools{}
	data, _ := v.client.Get(fmt.Sprintf("/vm/%s/tools", v.UUID))
	json.Unmarshal(data, &tools)
	return tools
}

type HotAddConfig struct {
	CPUHotAdd    bool `json:"cpu_hot_add_enabled"`
	MemoryHotAdd bool `json:"mem_hot_add_enabled"`
}

func (v VirtualMachine) GetHotAddConfig() HotAddConfig {
	hotAdd := HotAddConfig{}
	data, _ := v.client.Get(fmt.Sprintf("/vm/%s/capabilities", v.UUID))
	json.Unmarshal(data, &hotAdd)
	return hotAdd
}

func (v VirtualMachine) SetHotAdd(cpuHotAdd, memoryHotAdd bool) (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	hotAdd := HotAddConfig{
		CPUHotAdd:    cpuHotAdd,
		MemoryHotAdd: memoryHotAdd,
	}
	output, _ := json.Marshal(&hotAdd)
	data, err := v.client.Put(fmt.Sprintf("/vm/%s/capabilities", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VirtualMachine) Delete() (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	data, err := v.client.Delete(fmt.Sprintf("/vm/%s", v.UUID))
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VirtualMachine) IsBusy() bool {
	tasks := []Task{}
	data, _ := v.client.Delete(fmt.Sprintf("/task/%s/entity/%s/active", v.LocationID, v.UUID))
	json.Unmarshal(data, &tasks)
	if len(tasks) == 0 {
		return false
	}
	return true
}

func (v VirtualMachine) Rename(newName string) (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	params := struct {
		Name string `json:"name"`
	}{
		Name: newName,
	}
	output, _ := json.Marshal(&params)
	data, err := v.client.Put(fmt.Sprintf("/vm/%s/name", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VirtualMachine) PowerOn() (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vm/%s/poweron", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VirtualMachine) Reboot() (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vm/%s/reboot", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VirtualMachine) PowerOff() (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vm/%s/poweroff", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VirtualMachine) Shutdown() (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	tools := v.GetTools()
	if tools.Version == "0" {
		return v.PowerOff()
	}
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vm/%s/shutdown", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VirtualMachine) ModifyCPU(cpuCount int) (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	params := struct {
		CPUCount       int `json:"cpus_number"`
		CoresPerSocket int `json:"cores_per_socket"`
	}{
		CPUCount:       cpuCount,
		CoresPerSocket: 1,
	}
	output, _ := json.Marshal(&params)
	data, err := v.client.Put(fmt.Sprintf("/vm/%s/cpu", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VirtualMachine) ModifyMemory(memoryMB int) (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	params := struct {
		MemoryMB int `json:"memory_size"`
	}{
		MemoryMB: memoryMB,
	}
	output, _ := json.Marshal(&params)
	data, err := v.client.Put(fmt.Sprintf("/vm/%s/mem", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VirtualMachine) ModifyNics(nics []Nic) (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	output, _ := json.Marshal(&nics)
	data, err := v.client.Put(fmt.Sprintf("/vm/%s/vnics", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VirtualMachine) ModifyDisks(disks []Disk) (Task, error) {
	v.client.waitUntilObjectIsReady(v.LocationID, v.UUID)
	task := Task{}
	output, _ := json.Marshal(&disks)
	data, err := v.client.Put(fmt.Sprintf("/vm/%s/virtual-disks", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

type ConsoleSession struct {
	VMX    string `json:"vmx"`
	Ticket string `json:"ticket"`
	Host   string `json:"host"`
	Port   string `json:"port"`
}

func (v VirtualMachine) GetConsoleSession() (ConsoleSession, error) {
	session := ConsoleSession{}
	data, err := v.client.Get(fmt.Sprintf("/vm/%s/mks-screen-ticket", v.UUID))
	if err != nil {
		return session, err
	}
	err = json.Unmarshal(data, &session)
	return session, err
}

func (v VirtualMachine) GetScreenThumbnail() ([]byte, error) {
	data, err := v.client.getBinary(fmt.Sprintf("/vm/%s/screen", v.UUID))
	if err != nil {
		return []byte{}, err
	}
	return data, err
}
