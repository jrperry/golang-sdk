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

func (e Edge) GetUplinkInterface() EdgeInterface {
	for _, edgeInterface := range e.Interfaces {
		if edgeInterface.Type == "uplink" {
			return edgeInterface
		}
	}
	return EdgeInterface{}
}

func (e Edge) GetNATConfig() (EdgeNATConfig, error) {
	natConfig := EdgeNATConfig{}
	data, err := e.client.Get(fmt.Sprintf("/edge/%s/nat", e.UUID))
	if err != nil {
		return natConfig, err
	}
	err = json.Unmarshal(data, &natConfig)
	return natConfig, err
}

func (e Edge) AddNATRule(rule NATRule) (Task, error) {
	task := Task{}
	natConfig, err := e.GetNATConfig()
	if err != nil {
		return task, err
	}
	natConfig.Rules = append(natConfig.Rules, rule)
	output, _ := json.Marshal(&natConfig)
	data, err := e.client.Put(fmt.Sprintf("/edge/%s/nat", e.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = e.client
	return task, err
}

func (e Edge) DeleteNATRule(ruleID int) (Task, error) {
	task := Task{}
	natConfig, err := e.GetNATConfig()
	if err != nil {
		return task, err
	}
	natRules := []NATRule{}
	for _, rule := range natConfig.Rules {
		if rule.ID != ruleID {
			natRules = append(natRules, rule)
		}
	}
	natConfig.Rules = natRules
	output, _ := json.Marshal(&natConfig)
	data, err := e.client.Put(fmt.Sprintf("/edge/%s/nat", e.UUID), output)
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(data, &task)
	task.client = e.client
	return task, err
}
