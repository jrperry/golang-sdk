package iland

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
