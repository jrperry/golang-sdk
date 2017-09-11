package goiland

type VAppNetwork struct {
	client *Client
	Name   string `json:"name"`
	UUID   string `json:"uuid"`
}
