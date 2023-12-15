package models

import "time"

type SLA struct {
	StartTime time.Time
	EndTime   time.Time
	Breached  bool
	Status    SLAStatus
	SlaTime   int32
}

type SLAStatus string

const (
	Idle    SLAStatus = "Idle"
	Running SLAStatus = "Running"
	Paused  SLAStatus = "Paused"
	Ended   SLAStatus = "Ended"
)
