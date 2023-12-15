package models

import (
	"github.com/google/uuid"
)

type Incident struct {
	Id           uuid.UUID        `json:"id"`
	Name         string           `json:"name"`
	Severity     IncidentSeverity `json:"severity"`
	Status       Status           `json:"status"`
	IncidentType IncidentType     `json:"type"`
}

type IncidentType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Status string
type IncidentSeverity string

const (
	Low      IncidentSeverity = "Low"
	Medium   IncidentSeverity = "Medium"
	High     IncidentSeverity = "High"
	Critical IncidentSeverity = "Critical"
	Pending  Status           = "Pending"
	Open     Status           = "Open"
	Active   Status           = "Active"
	Closed   Status           = "Closed"
)

func NewIncident(name string, severity IncidentSeverity, incType IncidentType) *Incident {
	return &Incident{
		Id:           uuid.New(),
		Name:         name,
		Severity:     severity,
		IncidentType: incType,
	}
}

func NewIncidentType(name string) *IncidentType {
	return &IncidentType{
		Name: name,
	}
}
