package stores

import (
	"ContentManagement/api/models/task"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type TaskStore struct {
	DB *sql.DB
}

func NewTaskStore(db *sql.DB) *TaskStore {
	iS := &TaskStore{
		DB: db,
	}
	iS.createTable()
	return iS
}

func (s *TaskStore) createTable() {
	fmt.Println("Try to create Task Table")
	task, err1 := LoadSQL("task/CreateTable.sql")
	taskComment, err2 := LoadSQL("taskComment/CreateTable.sql")
	if err := errors.Join(err1, err2); err != nil {
		panic(err.Error())
	}
	_, err1 = s.DB.Exec(task)
	_, err2 = s.DB.Exec(taskComment)
	if err := errors.Join(err1, err2); err != nil {
		panic(err.Error())
	}
	fmt.Println("No Issues found on table creation")
}

func (s *TaskStore) GetAllTasks() (*[]task.Task, error) {
	var tasks []task.Task
	query, err := LoadSQL("task/GetAll.sql")
	if err != nil {
		return nil, err
	}
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		task, err := s.scanTask(rows.Scan)
		if err != nil {
			tasks = append(tasks, *task)
		} else {
			fmt.Println(err.Error())
		}
	}
	return &tasks, nil
}

func (s *TaskStore) GetTaskByID(id string) (*task.Task, error) {
	query, err := LoadSQL("task/GetById.sql")
	if err != nil {
		return nil, err
	}
	row := s.DB.QueryRow(query, id)
	task, err := s.scanTask(row.Scan)
	if err != nil {
		if err == sql.ErrNoRows {
			// No such task exists
			return nil, fmt.Errorf("no task found with id: %s", id)
		} else {
			return nil, err
		}
	}
	return task, nil
}

func (s *TaskStore) CreateTask(task *task.Task) error {
	query, err := LoadSQL("task/Create.sql")
	if err != nil {
		return err
	}
	_, err = s.DB.Exec(query,
		&task.Id,
		&task.Title,
		&task.Description,
		&task.Assignee.Id,
		&task.Assignee.FirstName,
		&task.Assignee.LastName,
		&task.Assignee.Email,
		&task.Assignee.Fullanme,
		&task.Incident.Id,
		&task.Incident.Name,
		&task.Incident.Severity,
		&task.Incident.Status,
		&task.Status,
		&task.Priority,
	)
	return err
}

func (s *TaskStore) UpdateTask(task *task.Task) error {
	// TODO Implement
	return fmt.Errorf("not implemented")
}

func (s *TaskStore) DeleteTask(id string) error {
	q, err := LoadSQL("task/Delete.sql")
	if err != nil {
		return err
	}
	_, err = s.DB.Exec(q, id)
	return err
}

func (s *TaskStore) getCommentsByTask(id uuid.UUID) *[]task.TaskComment {
	var comments []task.TaskComment
	commQuery, err := LoadSQL("taskComment/GetByTask.sql")
	if err == nil {
		rows, err := s.DB.Query(commQuery, id)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var com task.TaskComment
				err := rows.Scan(
					&com.Id,
					&com.Content,
					&com.WriterId,
					&com.TaskId,
				)
				if err == nil {
					comments = append(comments, com)
				}
			}
		}

	}
	return &comments
}

func (s *TaskStore) scanTask(scan ScanFunc) (*task.Task, error) {
	var task task.Task
	err := scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.Assignee.Id,
		&task.Assignee.FirstName,
		&task.Assignee.LastName,
		&task.Assignee.Email,
		&task.Assignee.Fullanme,
		&task.Incident.Id,
		&task.Incident.Name,
		&task.Incident.Severity,
		&task.Incident.Status,
		&task.Status,
		&task.Priority,
	)
	if err != nil {
		return nil, err
	}
	comments := s.getCommentsByTask(task.Id)
	task.Comments = *comments
	return &task, nil
}
