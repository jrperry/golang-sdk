package iland

import "time"

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

type PerfMetric struct {
	Group string
	Name  string
	Type  string
}

type PerfResults struct {
	Summary  string       `json:"summary"`
	Interval int          `json:"interval"`
	Group    string       `json:"cpu"`
	Name     string       `json:"name"`
	Type     string       `json:"type"`
	Unit     string       `json:"unit"`
	Samples  []PerfSample `json:"samples"`
}

type PerfSample struct {
	Time  int `json:"int"`
	Value int `json:"value"`
}

type BillingSummary struct {
	TotalCost             float64           `json:"total"`
	TotalCostEstimate     float64           `json:"estimate"`
	CPUTotalCost          float64           `json:"cpu"`
	CPUUsage              float64           `json:"cpu_usage"`
	CPUBurstCost          float64           `json:"cpu_burst"`
	CPUReserveUsage       float64           `json:"cpu_res_usage"`
	CPUBurstUsage         float64           `json:"cpu_burst_usage"`
	MemoryTotalCost       float64           `json:"mem"`
	MemoryUsage           float64           `json:"mem_usage"`
	MemoryReserveUsage    float64           `json:"mem_res_usage"`
	MemoryBurstUsage      float64           `json:"mem_burst_usage"`
	MemoryBurstCost       float64           `json:"mem_burst"`
	BandwidthTotalCost    float64           `json:"bandwidth"`
	BandwidthUsage        float64           `json:"bandwidth_usage"`
	BandwidthBurstCost    float64           `json:"bandwidth_burst"`
	BandwidthBurstUsage   float64           `json:"bandwidth_burst_usage"`
	BandwidthReserveCost  float64           `json:"bandwidth_reserved_cost"`
	BandwidthReserveUsage float64           `json:"bandwidth_reserved_usage"`
	DiskTotalCost         float64           `json:"disk"`
	DiskUsage             float64           `json:"disk_usage"`
	DiskBurstUsage        float64           `json:"disk_burst_usage"`
	DiskBurstCost         float64           `json:"disk_burst"`
	HardDiskUsage         float64           `json:"hdd_usage"`
	HardDiskCost          float64           `json:"hdd_cost"`
	HardDiskBurstUsage    float64           `json:"hdd_burst_usage"`
	HardDiskBurstCost     float64           `json:"hdd_burst_cost"`
	HardDiskReserveCost   float64           `json:"hdd_reserved_cost"`
	HardDiskReserveUsage  float64           `json:"hdd_reserved_usage"`
	SSDUsage              float64           `json:"ssd_usage"`
	SSDCost               float64           `json:"ssd_cost"`
	SSDBurstUsage         float64           `json:"ssd_burst_usage"`
	SSDBurstCost          float64           `json:"ssd_burst_cost"`
	SSDReserveCost        float64           `json:"ssd_reserved_cost"`
	SSDReserveUsage       float64           `json:"ssd_reserved_usage"`
	ArchiveUsage          float64           `json:"archive_usage"`
	ArchiveCost           float64           `json:"archive_cost"`
	ArchiveBurstUsage     float64           `json:"archive_burst_usage"`
	ArchiveBurstCost      float64           `json:"archive_burst_cost"`
	ArchiveReserveCost    float64           `json:"archive_reserved_cost"`
	ArchiveReserveUsage   float64           `json:"archive_reserved_usage"`
	ZertoArchiveUsage     float64           `json:"zerto_archive_usage"`
	ZertoArchiveCost      float64           `json:"zerto_archive_cost"`
	ZertoAdvancedUsage    float64           `json:"zerto_advanced_usage"`
	ZertoAdvancedCost     float64           `json:"zerto_advanced_cost"`
	EntityUUID            string            `json:"entity_uuid"`
	EntityType            string            `json:"entity_type"`
	EntityName            string            `json:"entity_name"`
	CurrencyCode          string            `json:"currency_code"`
	CurrentTime           int               `json:"time"`
	TestDrive             bool              `json:"test_drive"`
	LineItems             []BillingLineItem `json:"line_items"`
	Discount              float64           `json:"discount"`
}

type BillingLineItem struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	ProductID string  `json:"product_id"`
}

func getPerfLimit(perfInterval string) string {
	switch perfInterval {
	case PerfIntervalSecond:
		return "500"
	case PerfIntervalHour:
		return "1000"
	case PerfIntervalDay:
		return "2000"
	case PerfIntervalWeek:
		return "3000"
	case PerfIntervalMonth:
		return "5000"
	default:
		return "1000"
	}
}

func getUnixMilliseconds(datetime time.Time) int {
	return int(datetime.UnixNano() / int64(time.Millisecond))
}
