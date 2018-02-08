package iland

import (
	"encoding/json"
	"fmt"
	"time"
)

type Task struct {
	client         *Client
	UUID           string `json:"uuid"`
	EntityUUID     string `json:"entity_uuid"`
	Status         string `json:"status"`
	Progress       int    `json:"progress"`
	Active         bool   `json:"active"`
	Synchronized   bool   `json:"synchronized"`
	Message        string `json:"message"`
	TaskType       string `json:"task_type"`
	Operation      string `json:"operation"`
	Description    string `json:"operation_description"`
	LocationID     string `json:"location_id"`
	OrgUUID        string `json:"org_uuid"`
	VCloudTaskID   string `json:"task_id"`
	Username       string `json:"username"`
	InitiationTime int    `json:"initiation_time"`
	StartTime      int    `json:"start_time"`
	EndTime        int    `json:"end_time"`
}

func (t Task) Refresh() Task {
	task := Task{}
	data, _ := t.client.Get(fmt.Sprintf("/task/%s/%s", t.LocationID, t.UUID))
	json.Unmarshal(data, &task)
	task.client = t.client
	return task
}

func (t Task) Track() Task {
	for {
		task := Task{}
		data, _ := t.client.Get(fmt.Sprintf("/task/%s/%s", t.LocationID, t.UUID))
		json.Unmarshal(data, &task)
		if !task.Active && task.Synchronized {
			task.client = t.client
			return task
		}
		time.Sleep(time.Second * 10)
	}
}
