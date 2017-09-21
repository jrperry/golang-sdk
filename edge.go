package iland

import (
	"encoding/json"
	"fmt"
)

type Edge struct {
	client                    *Client
	Name                      string          `json:"name"`
	UUID                      string          `json:"uuid"`
	Description               string          `json:"description"`
	BackwardCompatabilityMode bool            `json:"backward_compatibility_mode"`
	GatewayBackingConfig      string          `json:"gateway_backing_config"`
	HighAvailabilityEnabled   bool            `json:"high_availability_enabled"`
	DefaultDNSRelayRoute      bool            `json:"default_dns_relay_route"`
	Interfaces                []EdgeInterface `json:"interfaces"`
	LocationID                string          `json:"location_id"`
	OrgUUID                   string          `json:"org_uuid"`
	VdcUUID                   string          `json:"vdc_uuid"`
	UpdatedDate               int             `json:"updated_date"`
}

type EdgeInterface struct {
	Name                string                `json:"name"`
	DisplayName         string                `json:"display_name"`
	Type                string                `json:"type"`
	SubnetParticipation []SubnetParticipation `json:"subnet_participation"`
	DefaultRoute        bool                  `json:"default_route"`
	LimitEnabled        bool                  `json:"apply_rate_limit"`
	InboundLimit        float64               `json:"in_rate_limit"`
	OutboundLimit       float64               `json:"out_rate_limit"`
	NetworkName         string                `json:"network"`
	NetworkUUID         string                `json:"network_uuid"`
}

type EdgeNATConfig struct {
	Enabled bool      `json:"enabled"`
	Rules   []NATRule `json:"rules"`
}

type NATRule struct {
	ID             int    `json:"id,omitempty"`
	Type           string `json:"type"`
	Enabled        bool   `json:"enabled"`
	Description    string `json:"description"`
	OriginalIP     string `json:"original_ip"`
	OriginalPort   string `json:"original_port,omitempty"`
	TranslatedIP   string `json:"translated_ip"`
	TranslatedPort string `json:"translated_port,omitempty"`
	Protocol       string `json:"protocol"`
	InterfaceName  string `json:"interface"`
}

type EdgeFirewallConfig struct {
	EdgeUUID      string         `json:"edge_uuid"`
	Version       int            `json:"version,omitempty"`
	DefaultAction string         `json:"default_action"`
	Enabled       bool           `json:"enabled"`
	Log           bool           `json:"log"`
	Rules         []FirewallRule `json:"rules"`
}

type FirewallRule struct {
	ID                   string   `json:"id,omitempty"`
	IDX                  int      `json:"idx,omitempty"`
	Enabled              bool     `json:"enabled"`
	Description          string   `json:"description"`
	DestinationIP        string   `json:"destination_ip"`
	DestinationPortRange string   `json:"destination_port_range"`
	Direction            string   `json:"direction,omitempty"`
	ICMPSubType          string   `json:"icmp_sub_type,omitempty"`
	Logging              bool     `json:"logging"`
	MatchOnTranslate     bool     `json:"match_on_translate"`
	Policy               string   `json:"policy"`
	Protocol             []string `json:"protocol"`
	SourceIP             string   `json:"source_ip"`
	SourcePort           int      `json:"source_port"`
	SourcePortRange      string   `json:"source_port_range"`
}

func (e Edge) GetUplinkInterface() EdgeInterface {
	for _, edgeInterface := range e.Interfaces {
		if edgeInterface.Type == "uplink" {
			return edgeInterface
		}
	}
	return EdgeInterface{}
}

func (e Edge) GetFirewallConfig() (EdgeFirewallConfig, error) {
	config := EdgeFirewallConfig{}
	data, err := e.client.Get(fmt.Sprintf("/edge/%s/firewall", e.UUID))
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}

func (e Edge) GetExternalInterface() EdgeInterface {
	edgeInterface := EdgeInterface{}
	data, _ := e.client.Get(fmt.Sprintf("/edge/%s/edge-interface", e.UUID))
	fmt.Println(string(data))
	err := json.Unmarshal(data, &edgeInterface)
	if err != nil {
		fmt.Println(err)
	}
	return edgeInterface
}

func (e Edge) UpdateExternalInterface(edgeInterface EdgeInterface) (Task, error) {
	task := Task{}
	output, err := json.Marshal(&edgeInterface)
	if err != nil {
		return task, err
	}
	data, err := e.client.Put(fmt.Sprintf("/edge/%s/edge-interface", e.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	return task, err
}

func (e Edge) UpdateFirewallConfig(config EdgeFirewallConfig) (Task, error) {
	e.client.waitUntilObjectIsReady(e.LocationID, e.UUID)
	task := Task{}
	output, err := json.Marshal(&config)
	if err != nil {
		return task, err
	}
	fmt.Println(string(output))
	data, err := e.client.Put(fmt.Sprintf("/edge/%s/firewall", e.UUID), output)
	if err != nil {
		return task, err
	}
	fmt.Println(string(data))
	err = json.Unmarshal(data, &task)
	task.client = e.client
	return task, err
}

func (e Edge) GetNATConfig() (EdgeNATConfig, error) {
	config := EdgeNATConfig{}
	data, err := e.client.Get(fmt.Sprintf("/edge/%s/nat", e.UUID))
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}

func (e Edge) UpdateNATConfig(config EdgeNATConfig) (Task, error) {
	e.client.waitUntilObjectIsReady(e.LocationID, e.UUID)
	task := Task{}
	output, err := json.Marshal(&config)
	if err != nil {
		return task, err
	}
	data, err := e.client.Put(fmt.Sprintf("/edge/%s/nat", e.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = e.client
	return task, err
}
