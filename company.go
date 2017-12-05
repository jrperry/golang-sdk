package iland

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	client      *Client
	CRM         string `json:"uuid"`
	Name        string `json:"name"`
	HasIAAS     bool   `json:"has_iaas"`
	HasVCC      bool   `json:"has_vcc"`
	Deleted     bool   `json:"deleted"`
	UpdatedDate int    `json:"updated_date"`
	DeletedDate int    `json:"deleted_date"`
}

func (c *Company) GetUsers() []User {
	users := []User{}
	data, _ := c.client.Get(fmt.Sprintf("/companies/%s/users", c.CRM))
	json.Unmarshal(data, &users)
	return users
}

func (c *Company) GetCloudTenants() []CloudTenant {
	cloudTenants := []CloudTenant{}
	data, _ := c.client.Get(fmt.Sprintf("/companies/%s/cloud-tenants", c.CRM))
	json.Unmarshal(data, &cloudTenants)
	return cloudTenants
}

func (c *Company) GetSupportTickets() []SupportTicket {
	tickets := []SupportTicket{}
	data, _ := c.client.Get(fmt.Sprintf("/companies/%s/support-tickets", c.CRM))
	err := json.Unmarshal(data, &tickets)
	if err != nil {
		fmt.Println(err)
	}
	for i, ticket := range tickets {
		ticket.client = c.client
		tickets[i] = ticket
	}
	return tickets
}

func (c *Company) GetSupportTicket(ticketID int) (SupportTicket, error) {
	ticket := SupportTicket{}
	data, err := c.client.Get(fmt.Sprintf("/companies/%s/support-tickets/%d", c.CRM, ticketID))
	if err != nil {
		return ticket, err
	}
	err = json.Unmarshal(data, &ticket)
	ticket.client = c.client
	return ticket, err
}
