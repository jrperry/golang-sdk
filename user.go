package iland

import (
	"encoding/json"
	"fmt"
)

type User struct {
	client      *Client
	Name        string `json:"name"`
	FullName    string `json:"fullname"`
	CRM         string `json:"crm"`
	Type        string `json:"user_type"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Company     string `json:"company"`
	Country     string `json:"country"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Locked      bool   `json:"locked"`
	Deleted     bool   `json:"deleted"`
	CreatedDate int    `json:"created_date"`
	DeletedDate int    `json:"deleted_date"`
}

type UserRole struct {
	Role     string `json:"role"`
	Type     string `json:"type"`
	Username string `json:"username"`
	OrgUUID  string `json:"org_uuid"`
}

func (u *User) GetRoles() []UserRole {
	roles := []UserRole{}
	data, _ := u.client.Get(fmt.Sprintf("/user/%s/roles", u.Name))
	json.Unmarshal(data, &roles)
	return roles
}

func (u *User) GetAlerts() []Alert {
	alerts := []Alert{}
	data, _ := u.client.Get(fmt.Sprintf("/user/%s/alerts", u.Name))
	json.Unmarshal(data, &alerts)
	return alerts
}
