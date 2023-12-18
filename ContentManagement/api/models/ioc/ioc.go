package ioc

import (
	"github.com/google/uuid"
)

type IOCType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type IOC struct {
	Id      uuid.UUID `json:"id"`
	Value   string    `json:"value"`
	IocType IOCType   `json:"iocType"`
	Verdict Verdict   `json:"verdict"`
}

type Ioc_Incident struct {
	Id         int    `json:"id"`
	IocId      string `json:"iocId"`
	IncidentId string `json:"incidentId"`
}

type Verdict string

const (
	Benigne   Verdict = "Benigne"
	Malicious Verdict = "Malicious"
	Neutral   Verdict = "Neutral"
)

func NewIOC(value string, iT IOCType) *IOC {
	return &IOC{
		Id:      uuid.New(),
		Value:   value,
		IocType: iT,
	}
}
