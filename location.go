package goiland

import (
	"encoding/json"
	"fmt"
)

type Location struct {
	client      *Client
	ID          string `json:"location_id"`
	CRM         string `json:"crm"`
	UpdatedDate int    `json:"updated_date"`
}

func (l Location) GetEntityActiveTasks(entityUUID string) ([]Task, error) {
	tasks := []Task{}
	data, err := l.client.Get(fmt.Sprintf("/task/%s/entity/%s", l.ID, entityUUID))
	if err != nil {
		return tasks, err
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return tasks, err
	}
	for i, task := range tasks {
		task.client = l.client
		tasks[i] = task
	}
	return tasks, nil
}
