package iland

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"mime/multipart"
	"os"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

type Catalog struct {
	client      *Client
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Description string `json:"description"`
	Version     int    `json:"version"`
	Shared      bool   `json:"shared"`
	Public      bool   `json:"public"`
	LocationID  string `json:"location_id"`
	OrgUUID     string `json:"org_uuid"`
	VCloudHref  string `json:"vcloud_href"`
	CreatedDate int    `json:"created_date"`
	UpdatedDate int    `json:"updated_date"`
}

func (c Catalog) GetVAppTemplates() []VAppTemplate {
	vAppTemplates := []VAppTemplate{}
	data, _ := c.client.Get(fmt.Sprintf("/catalog/%s/vapp-templates", c.UUID))
	json.Unmarshal(data, &vAppTemplates)
	for i, vAppTemplate := range vAppTemplates {
		vAppTemplate.client = c.client
		vAppTemplates[i] = vAppTemplate
	}
	return vAppTemplates
}

func (c Catalog) AddVAppTemplate(sourceVAppUUID, newVAppTemplateName string) (Task, error) {
	task := Task{}
	vApp, err := c.client.GetVApp(sourceVAppUUID)
	if err != nil {
		return task, err
	}
	if newVAppTemplateName == "" {
		newVAppTemplateName = vApp.Name
	}
	existingVAppTemplates := c.GetVAppTemplates()
	for _, vAppTemplate := range existingVAppTemplates {
		if vAppTemplate.Name == newVAppTemplateName {
			return task, fmt.Errorf("vApp template with name, %s, already exists in this catalog", newVAppTemplateName)
		}
	}
	params := struct {
		Name string `json:"name"`
	}{
		Name: newVAppTemplateName,
	}
	output, _ := json.Marshal(&params)
	data, err := c.client.Post(fmt.Sprintf("/catalog/%s/add-vapp-template/%s", c.UUID, sourceVAppUUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	if err != nil {
		return task, err
	}
	task.client = c.client
	return task, nil
}

func (c Catalog) UploadVAppTemplate(ovaFilePath, vAppTemplateName, storageProfileUUID string) error {
	var storageProfile StorageProfile
	if storageProfileUUID == "" {
		org, _ := c.client.GetOrg(c.OrgUUID)
		storageProfile = org.GetDefaultStorageProfile()
	} else {
		var err error
		storageProfile, err = c.client.GetStorageProfile(storageProfileUUID)
		if err != nil {
			return fmt.Errorf("storage profile with UUID, %s, does not exist", storageProfileUUID)
		}
	}

	chunkSizeBytes := 5242880
	file, _ := os.Open(ovaFilePath)
	defer file.Close()
	fileStats, _ := file.Stat()
	fileSize := int(fileStats.Size())
	chunks := int(math.Ceil(float64(fileSize) / float64(chunkSizeBytes)))
	totalSize := strconv.Itoa(fileSize)
	totalChunks := strconv.Itoa(chunks)
	bytesSent := 0
	uploadID := uuid.NewV4().String()
	for i := 1; i <= chunks; i++ {
		if bytesSent+chunkSizeBytes > fileSize {
			chunkSizeBytes = fileSize - bytesSent
		}
		chunk := make([]byte, chunkSizeBytes)
		file.Read(chunk)

		chunkNumber := strconv.Itoa(i)
		chunkSize := strconv.Itoa(len(chunk))

		var b bytes.Buffer
		w := multipart.NewWriter(&b)
		writer, _ := w.CreateFormFile("file", fileStats.Name())
		writer.Write(chunk)
		w.WriteField("name", vAppTemplateName)
		w.WriteField("description", "test")
		w.WriteField("vdc", storageProfile.VdcUUID)
		w.WriteField("storage_profile", storageProfile.UUID)
		w.WriteField("resumableIdentifier", uploadID)
		w.WriteField("resumableChunkNumber", chunkNumber)
		w.WriteField("resumableChunkSize", chunkSize)
		w.WriteField("resumableCurrentChunkSize", chunkSize)
		w.WriteField("resumableTotalSize", totalSize)
		w.WriteField("resumableTotalChunks", totalChunks)
		w.WriteField("resumableRelativePath", fileStats.Name())
		w.WriteField("resumableFileName", fileStats.Name())
		w.WriteField("resumableType", "")
		w.Close()

		contentType := "multipart/form-data; boundary=" + w.Boundary()
		_, err := c.client.postForm(fmt.Sprintf("/catalog/%s/vapp-template/upload", c.UUID), contentType, b.Bytes())
		if err != nil {
			return err
		}
		bytesSent += len(chunk)
	}
	return nil
}
