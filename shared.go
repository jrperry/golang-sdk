package iland

type SubnetParticipation struct {
	Gateway  string    `json:"gateway"`
	Netmask  string    `json:"netmask"`
	IPAddres string    `json:"ip_address"`
	IPRanges []IPRange `json:"ip_ranges"`
}

type IPRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type NATRule struct {
	ID             int    `json:"id"`
	Type           string `json:"type"`
	Enabled        bool   `json:"enabled"`
	Description    string `json:"description"`
	OriginalIP     string `json:"original_ip"`
	OriginalPort   string `json:"original_port"`
	TranslatedIP   string `json:"translated_ip"`
	TranslatedPort string `json:"translated_port"`
	Protocol       string `json:"protocol"`
	InterfaceName  string `json:"interface"`
}
