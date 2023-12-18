package worklog

import (
	"time"

	"github.com/google/uuid"
)

type CreatWorklogRequest struct {
	WriterId   string
	IncidentId string
	Content    string
}

type Worklog struct {
	Id         uuid.UUID
	WriterId   uuid.UUID
	IncidentId uuid.UUID
	Content    string
	Created_at string
}

func NewWorklog(writerId uuid.UUID, inc_id uuid.UUID, content string) *Worklog {
	return &Worklog{
		Id:         uuid.New(),
		WriterId:   writerId,
		IncidentId: inc_id,
		Content:    content,
		Created_at: time.Now().UTC().String(),
	}
}
