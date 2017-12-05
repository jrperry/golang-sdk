package iland

import (
	"encoding/json"
	"fmt"
	"io"
)

type SupportTicket struct {
	client           *Client
	ID               int      `json:"id"`
	Summary          string   `json:"summary"`
	Status           string   `json:"status"`
	CRM              string   `json:"crm"`
	CreatorFullName  string   `json:"creator_full_name"`
	CreatorUsername  string   `json:"creator_user_name"`
	CCEmailsEnabled  bool     `json:"cc_emails_enabled"`
	CCEmailAddresses []string `json:"cc_email_addresses"`
	CreationDate     int      `json:"creation_date"`
}

type TicketAttachment struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Filename        string `json:"filename"`
	CreatorFullName string `json:"creator_full_name"`
	CreatorUsername string `json:"creator_user_name"`
	CreationDate    int    `json:"creation_date"`
}

type TicketComment struct {
	ID              int    `json:"id"`
	TicketID        int    `json:"ticket_id"`
	Text            string `json:"text"`
	Type            string `json:"comment_type"`
	CreatorFullName string `json:"creator_full_name"`
	CreatorUsername string `json:"creator_username"`
	CreationDate    int    `json:"creation_date"`
}

func (t *SupportTicket) GetAttachments() []TicketAttachment {
	attachments := []TicketAttachment{}
	data, _ := t.client.Get(fmt.Sprintf("/companies/%s/support-tickets/%d/attachments", t.CRM, t.ID))
	err := json.Unmarshal(data, &attachments)
	if err != nil {
		fmt.Println(err)
	}
	return attachments
}

func (t *SupportTicket) DownloadAttachment(attachmentID int) (io.ReadCloser, error) {
	reader, err := t.client.getBinaryStream(fmt.Sprintf("/companies/%s/support-tickets/%d/attachments/%d", t.CRM, t.ID, attachmentID))
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func (t *SupportTicket) GetComments() []TicketComment {
	comments := []TicketComment{}
	data, _ := t.client.Get(fmt.Sprintf("/companies/%s/support-tickets/%d/comments", t.CRM, t.ID))
	json.Unmarshal(data, &comments)
	return comments
}
