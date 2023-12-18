package task

import (
	"ContentManagement/api/models/incident"
	"ContentManagement/api/models/user"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID
	Title       string
	Description string
	Assignee    user.User
	Incident    incident.Incident
	Status      TaskSatus
	Priority    TaskPriority
	Comments    []TaskComment
}

type TaskComment struct {
	Id       uuid.UUID
	Content  string
	WriterId uuid.UUID
	TaskId   uuid.UUID
}
type TaskPriority string
type TaskSatus string

const (
	TaskOpen   TaskSatus    = "Open"
	InProgress TaskSatus    = "In Progress"
	Done       TaskSatus    = "Done"
	tpLow      TaskPriority = "Low"
	tpMedium   TaskPriority = "Medium"
	tpHigh     TaskPriority = "High"
	tpCritical TaskPriority = "Critical"
)

func NewTask(title, description string, userId, incId uuid.UUID, prio TaskPriority, status TaskSatus) *Task {
	state := TaskOpen
	if status != "" {
		state = status
	}
	return &Task{
		Id:          uuid.New(),
		Title:       title,
		Description: description,
		Assignee: user.User{
			Id: userId,
		},
		Incident: incident.Incident{
			Id: incId,
		},
		Status:   state,
		Priority: prio,
	}
}
