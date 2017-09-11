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
	Index            int    `json:"nic_id"`
	IPAddress        string `json:"ip_address"`
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

func (v VirtualMachine) Delete() (Task, error) {
	task := Task{}
	data, err := v.client.Delete(fmt.Sprintf("/vm/%s", v.UUID))
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = v.client
	return task, err
}

func (v VirtualMachine) Rename(newName string) (Task, error) {
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
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vm/%s/poweron", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	return task, err
}

func (v VirtualMachine) Reboot() (Task, error) {
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vm/%s/reboot", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	return task, err
}

func (v VirtualMachine) PowerOff() (Task, error) {
	task := Task{}
	data, err := v.client.Post(fmt.Sprintf("/vm/%s/poweroff", v.UUID), []byte{})
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	return task, err
}

func (v VirtualMachine) ModifyCPU(cpuCount int) (Task, error) {
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

func (v VirtualMachine) AddNic(nic Nic) (Task, error) {
	task := Task{}
	nics := v.GetNics()
	nics = append(nics, nic)
	output, _ := json.Marshal(&nics)
	data, err := v.client.Put(fmt.Sprintf("/vm/%s/vnics", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	return task, err
}

func (v VirtualMachine) UpdateNics(nics []Nic) (Task, error) {
	task := Task{}
	output, _ := json.Marshal(&nics)
	data, err := v.client.Put(fmt.Sprintf("/vm/%s/vnics", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	return task, err
}

func (v VirtualMachine) DeleteNic(nicIndex int) (Task, error) {
	task := Task{}
	data, err := v.client.Delete(fmt.Sprintf("/vm/%s/vnics/%d", v.UUID, nicIndex))
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	return task, err
}

func (v VirtualMachine) AddDisk(diskSizeMB int) (Task, error) {
	task := Task{}
	disk := Disk{
		Size: diskSizeMB,
		Type: "LSI_LOGIC",
	}
	output, _ := json.Marshal(&disk)
	data, err := v.client.Post(fmt.Sprintf("/vm/%s/virtual-disk", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	return task, err
}

func (v VirtualMachine) ModifyDisk(disk Disk) (Task, error) {
	task := Task{}
	output, _ := json.Marshal(&disk)
	data, err := v.client.Put(fmt.Sprintf("/vm/%s/virtual-disk", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	return task, err
}

func (v VirtualMachine) RemoveDisk(diskName string) (Task, error) {
	task := Task{}
	disks := v.GetDisks()
	updatedDisks := []Disk{}
	for _, disk := range disks {
		if disk.Name != diskName {
			updatedDisks = append(updatedDisks, disk)
		}
	}
	output, _ := json.Marshal(&updatedDisks)
	data, err := v.client.Put(fmt.Sprintf("/vm/%s/virtual-disks", v.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
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
