package config

import "time"

type NewBrokerResponse struct {
	GUID string `json:"guid"`
}

type Plan struct {
	Name            string           `json:"name" validate:"printascii,min=5"`
	ID              string           `json:"id,omitempty"`
	Description     string           `json:"description,omitempty"`
	MaintenanceInfo *MaintenanceInfo `json:"maintenance_info,omitempty"`
}

type Service struct {
	Name        string `json:"name" validate:"printascii,min=5"`
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Plans       []Plan `json:"plans" validate:"min=1,dive"`
}

type BrokerConfiguration struct {
	Services            []Service     `json:"services" validate:"min=1,dive"`
	Username            string        `json:"username" validate:"printascii,min=5"`
	Password            string        `json:"password" validate:"printascii,min=5"`
	CatalogResponse     int           `json:"catalog_response,omitempty" validate:"min=0,max=600"`
	ProvisionResponse   int           `json:"provision_response,omitempty" validate:"min=0,max=600"`
	DeprovisionResponse int           `json:"deprovision_response,omitempty" validate:"min=0,max=600"`
	AsyncResponseDelay  time.Duration `json:"async_response_delay"`
}

type MaintenanceInfo struct {
	Version     string `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
}
