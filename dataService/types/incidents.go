package types

import (
	"time"

	"github.com/google/uuid"
)

type Incident struct {
	Id           uuid.UUID      `json:"id"`
	Name         string         `json:"name"`
	Severity     Priority       `json:"severity"`
	Status       IncidentStatus `json:"status"`
	IncidentType IncidentType   `json:"type"`
	Creationdate uint           `json:"creationdate"`
	Tasks        []Task         `json:"tasks"`
}

type IncidentType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type IncidentStatus string
type Priority string

var IncidentQueryParams = []string{"id", "name", "severity", "status", "type", "creationdate"}
var IncidentTypeQueryParams = []string{"name", "id"}

const (
	Low      Priority       = "Low"
	Medium   Priority       = "Medium"
	High     Priority       = "High"
	Critical Priority       = "Critical"
	Pending  IncidentStatus = "Pending"
	Open     IncidentStatus = "Open"
	Active   IncidentStatus = "Active"
	Closed   IncidentStatus = "Closed"
)

func NewIncident(name string, severity Priority, incType IncidentType) *Incident {
	return &Incident{
		Id:           uuid.New(),
		Name:         name,
		Severity:     severity,
		IncidentType: incType,
		Status:       Pending,
		Creationdate: uint(time.Now().Unix()),
	}
}

func NewIncidentType(name string) *IncidentType {
	return &IncidentType{
		Name: name,
	}
}

func (i *Incident) ScanTo(scan ScanFunc) error {
	return scan(
		&i.Id,
		&i.Name,
		&i.Severity,
		&i.Status,
		&i.Creationdate,
		&i.IncidentType.Id,
		&i.IncidentType.Name)
}

func (iT *IncidentType) ScanTo(scan ScanFunc) error {
	return scan(
		&iT.Id,
		&iT.Name)
}
