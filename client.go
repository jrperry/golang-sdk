package iland

import (
	"encoding/json"
	"fmt"
	"time"
)

type Client struct {
	username        string
	password        string
	clientID        string
	clientSecret    string
	Token           Token
	tokenExpiration time.Time
}

func NewClient(Username, Password, ClientID, ClientSecret string) *Client {
	client := Client{
		username:     Username,
		password:     Password,
		clientID:     ClientID,
		clientSecret: ClientSecret,
	}
	return &client
}

func (c *Client) Get(endpoint string) ([]byte, error) {
	return c.request(endpoint, "GET", []byte{})
}

func (c *Client) Post(endpoint string, body []byte) ([]byte, error) {
	return c.request(endpoint, "POST", body)
}

func (c *Client) Put(endpoint string, body []byte) ([]byte, error) {
	return c.request(endpoint, "PUT", body)
}

func (c *Client) Delete(endpoint string) ([]byte, error) {
	return c.request(endpoint, "DELETE", []byte{})
}

func (c *Client) GetLocations() []Location {
	locations := []Location{}
	data, _ := c.Get(fmt.Sprintf("/user/%s/inventory", c.username))
	json.Unmarshal(data, &locations)
	for i, location := range locations {
		location.client = c
		locations[i] = location
	}
	return locations
}

func (c *Client) GetLocation(locationID string) (Location, error) {
	for _, location := range c.GetLocations() {
		if location.ID == locationID {
			location.client = c
			return location, nil
		}
	}
	return Location{}, fmt.Errorf("location with ID, %s, not found", locationID)
}

func (c *Client) GetOrgs() []Org {
	orgs := []Org{}
	for _, location := range c.GetLocations() {
		locationOrgs := []Org{}
		data, _ := c.Get(fmt.Sprintf("/location/%s/orgs", location.ID))
		json.Unmarshal(data, &locationOrgs)
		orgs = append(orgs, locationOrgs...)
	}
	for i, org := range orgs {
		org.client = c
		orgs[i] = org
	}
	return orgs
}

func (c *Client) GetOrg(orgUUID string) (Org, error) {
	org := Org{}
	data, err := c.Get(fmt.Sprintf("/org/%s", orgUUID))
	if err != nil {
		return org, err
	}
	err = json.Unmarshal(data, &org)
	if err != nil {
		return org, err
	}
	org.client = c
	return org, nil
}

func (c *Client) GetCatalogs() []Catalog {
	catalogs := []Catalog{}
	for _, org := range c.GetOrgs() {
		orgCatalogs := []Catalog{}
		data, _ := c.Get(fmt.Sprintf("/org/%s/catalogs", org.UUID))
		json.Unmarshal(data, &orgCatalogs)
		catalogs = append(catalogs, orgCatalogs...)
	}
	for i, catalog := range catalogs {
		catalog.client = c
		catalogs[i] = catalog
	}
	return catalogs
}

func (c *Client) GetCatalog(catalogUUID string) (Catalog, error) {
	catalog := Catalog{}
	data, err := c.Get(fmt.Sprintf("/catalog/%s", catalogUUID))
	if err != nil {
		return catalog, err
	}
	err = json.Unmarshal(data, &catalog)
	if err != nil {
		return catalog, err
	}
	catalog.client = c
	return catalog, nil
}

func (c *Client) GetVdcs() []Vdc {
	vdcs := []Vdc{}
	for _, location := range c.GetLocations() {
		locationVdcs := []Vdc{}
		data, _ := c.Get(fmt.Sprintf("/location/%s/vdcs", location.ID))
		json.Unmarshal(data, &locationVdcs)
		vdcs = append(vdcs, locationVdcs...)
	}
	for i, vdc := range vdcs {
		vdc.client = c
		vdcs[i] = vdc
	}
	return vdcs
}

func (c *Client) GetVdc(vdcUUID string) (Vdc, error) {
	vdc := Vdc{}
	data, err := c.Get(fmt.Sprintf("/vdc/%s", vdcUUID))
	if err != nil {
		return vdc, err
	}
	err = json.Unmarshal(data, &vdc)
	if err != nil {
		return vdc, err
	}
	vdc.client = c
	return vdc, nil
}

func (c *Client) GetStorageProfiles() []StorageProfile {
	storageProfiles := []StorageProfile{}
	for _, vdc := range c.GetVdcs() {
		vdcStorageProfiles := []StorageProfile{}
		data, _ := c.Get(fmt.Sprintf("/vdc/%s/storage-profiles", vdc.UUID))
		json.Unmarshal(data, &vdcStorageProfiles)
		storageProfiles = append(storageProfiles, vdcStorageProfiles...)
	}
	return storageProfiles
}

func (c *Client) GetStorageProfile(storageProfileUUID string) (StorageProfile, error) {
	for _, storageProfile := range c.GetStorageProfiles() {
		if storageProfile.UUID == storageProfileUUID {
			return storageProfile, nil
		}
	}
	return StorageProfile{}, fmt.Errorf("storage profile with UUID, %s, not found", storageProfileUUID)
}

func (c *Client) GetEdges() []Edge {
	edges := []Edge{}
	for _, org := range c.GetOrgs() {
		orgEdges := []Edge{}
		data, _ := c.Get(fmt.Sprintf("/org/%s/edges", org.UUID))
		fmt.Println(string(data))
		json.Unmarshal(data, &orgEdges)
		edges = append(edges, orgEdges...)
	}
	for i, edge := range edges {
		edge.client = c
		edges[i] = edge
	}
	return edges
}

func (c *Client) GetEdge(edgeUUID string) (Edge, error) {
	edge := Edge{}
	data, err := c.Get(fmt.Sprintf("/edge/%s", edgeUUID))
	if err != nil {
		return edge, err
	}
	err = json.Unmarshal(data, &edge)
	if err != nil {
		return edge, err
	}
	edge.client = c
	return edge, nil
}

func (c *Client) GetVdcNetworks() []VdcNetwork {
	vdcNetworks := []VdcNetwork{}
	for _, org := range c.GetOrgs() {
		networks := []VdcNetwork{}
		data, _ := c.Get(fmt.Sprintf("/org/%s/vdc-networks", org.UUID))
		json.Unmarshal(data, &networks)
		vdcNetworks = append(vdcNetworks, networks...)
	}
	for i, vdcNetwork := range vdcNetworks {
		vdcNetwork.client = c
		vdcNetworks[i] = vdcNetwork
	}
	return vdcNetworks
}

func (c *Client) GetVdcNetwork(vdcNetworkUUID string) (VdcNetwork, error) {
	for _, network := range c.GetVdcNetworks() {
		if network.UUID == vdcNetworkUUID {
			network.client = c
			return network, nil
		}
	}
	return VdcNetwork{}, fmt.Errorf("vdc network with uuid, %s, not found", vdcNetworkUUID)
}

func (c *Client) GetVAppTemplates() []VAppTemplate {
	vAppTemplates := []VAppTemplate{}
	for _, catalog := range c.GetCatalogs() {
		catalogTemplates := []VAppTemplate{}
		data, _ := c.Get(fmt.Sprintf("/catalog/%s/vapp-templates", catalog.UUID))
		json.Unmarshal(data, &catalogTemplates)
		vAppTemplates = append(vAppTemplates, catalogTemplates...)
	}
	locations := c.GetLocations()
	for _, location := range locations {
		publicTemplates, _ := location.GetPublicVAppTemplates()
		vAppTemplates = append(vAppTemplates, publicTemplates...)
	}
	for i, vAppTemplate := range vAppTemplates {
		vAppTemplate.client = c
		vAppTemplates[i] = vAppTemplate
	}
	return vAppTemplates
}

func (c *Client) GetVAppTemplate(vAppTemplateUUID string) (VAppTemplate, error) {
	vAppTemplate := VAppTemplate{}
	data, err := c.Get(fmt.Sprintf("/vapp-template/%s", vAppTemplateUUID))
	if err != nil {
		return vAppTemplate, err
	}
	err = json.Unmarshal(data, &vAppTemplate)
	if err != nil {
		return vAppTemplate, err
	}
	vAppTemplate.client = c
	return vAppTemplate, nil
}

func (c *Client) GetVApps() []VApp {
	vApps := []VApp{}
	for _, location := range c.GetLocations() {
		locationVApps := []VApp{}
		data, _ := c.Get(fmt.Sprintf("/location/%s/vapps", location.ID))
		json.Unmarshal(data, &locationVApps)
		vApps = append(vApps, locationVApps...)
	}
	for i, vApp := range vApps {
		vApp.client = c
		vApps[i] = vApp
	}
	return vApps
}

func (c *Client) GetVApp(vAppUUID string) (VApp, error) {
	vApp := VApp{}
	data, err := c.Get(fmt.Sprintf("/vapp/%s", vAppUUID))
	if err != nil {
		return vApp, err
	}
	err = json.Unmarshal(data, &vApp)
	if err != nil {
		return vApp, err
	}
	vApp.client = c
	return vApp, nil
}

func (c *Client) GetVirtualMachines() []VirtualMachine {
	virtualMachines := []VirtualMachine{}
	for _, location := range c.GetLocations() {
		locationVirtualMachines := []VirtualMachine{}
		data, _ := c.Get(fmt.Sprintf("/location/%s/vms", location.ID))
		json.Unmarshal(data, &locationVirtualMachines)
		virtualMachines = append(virtualMachines, locationVirtualMachines...)
	}
	for i, virtualMachine := range virtualMachines {
		virtualMachine.client = c
		virtualMachines[i] = virtualMachine
	}
	return virtualMachines
}

func (c *Client) GetVirtualMachine(virtualMachineUUID string) (VirtualMachine, error) {
	virtualMachine := VirtualMachine{}
	data, err := c.Get(fmt.Sprintf("/vm/%s", virtualMachineUUID))
	if err != nil {
		return virtualMachine, err
	}
	err = json.Unmarshal(data, &virtualMachine)
	if err != nil {
		return virtualMachine, err
	}
	virtualMachine.client = c
	return virtualMachine, nil
}

func (c *Client) GetMedias() []Media {
	medias := []Media{}
	for _, catalog := range c.GetCatalogs() {
		catalogMedia := []Media{}
		data, _ := c.Get(fmt.Sprintf("/catalog/%s/medias", catalog.UUID))
		json.Unmarshal(data, &catalogMedia)
		medias = append(medias, catalogMedia...)
	}
	for i, media := range medias {
		media.client = c
		medias[i] = media
	}
	return medias
}

func (c *Client) GetMedia(mediaUUID string) (Media, error) {
	media := Media{}
	data, err := c.Get(fmt.Sprintf("/media/%s", mediaUUID))
	if err != nil {
		return media, err
	}
	err = json.Unmarshal(data, &media)
	if err != nil {
		return media, err
	}
	media.client = c
	return media, nil
}

func (c *Client) GetTask(locationID, taskUUID string) (Task, error) {
	task := Task{}
	data, err := c.Get(fmt.Sprintf("/task/%s/%s", locationID, taskUUID))
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	if err != nil {
		return task, err
	}
	task.client = c
	return task, nil
}

func (c Client) waitUntilObjectIsReady(locationID, objectUUID string) {
	tasks := []Task{}
	for {
		data, _ := c.Delete(fmt.Sprintf("/task/%s/entity/%s/active", locationID, objectUUID))
		json.Unmarshal(data, &tasks)
		if len(tasks) == 0 {
			return
		}
		time.Sleep(time.Second * 5)
	}
}
