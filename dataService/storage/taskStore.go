package storage

import (
	"context"
	"data-service/types"
	"database/sql"
	"fmt"
)

type RequestTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnerId     string `json:"owner_id"`
	State       string `json:"state"`
	Priority    string `json:"priority"`
	IncidentID  string `json:"incident_id"`
}

type TaskStore struct {
	storage *MySqlStorage
	EntityStore[*types.Task]
	db *sql.DB
}

func NewTaskStore(storage *MySqlStorage) *TaskStore {
	return &TaskStore{
		storage: storage,
		db:      storage.Db,
	}
}

func (i *TaskStore) Get(ctx context.Context, id string) (*types.Task, *types.ApiError) {
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("tasks/GetById.sql")
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	res := i.db.QueryRow(loadedSql, id)
	if res.Err() != nil {
		if res.Err() == sql.ErrNoRows {
			return nil, types.NotFoundError(fmt.Errorf("no task found"), uri)
		}
		return nil, types.InternalServerError(res.Err(), uri)
	}
	var task types.Task
	if err := task.ScanTo(res.Scan); err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	return &task, nil
}

func (i *TaskStore) GetAll(ctx context.Context, qP QueryParameter) (*[]types.Task, *types.ApiError) {
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("tasks/GetAll.sql")
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	res, err := i.db.Query(loadedSql, qP.Limit, qP.Offset)
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	if res.Err() != nil {
		if res.Err() == sql.ErrNoRows {
			return nil, types.NotFoundError(fmt.Errorf("no tasks found"), uri)
		}
		return nil, types.InternalServerError(res.Err(), uri)
	}
	defer res.Close()
	tasks := []types.Task{}
	for res.Next() {
		var task types.Task
		if err := task.ScanTo(res.Scan); err != nil {
			return nil, types.InternalServerError(err, uri)
		} else {
			tasks = append(tasks, task)
		}
	}
	return &tasks, nil
}

func (i *TaskStore) GetQuery(ctx context.Context, qP QueryParameter) (*[]types.Task, *types.ApiError) {
	panic("not implemented") // TODO: Implement
}

func (i *TaskStore) Create(ctx context.Context, task *types.Task) (*types.Task, *types.ApiError) {
	fmt.Println("creating new task from ", task)
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("tasks/Create.sql")
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	_, err = i.db.Exec(loadedSql)
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	return task, nil
}

func (i *TaskStore) Update(entity types.Incident) (*types.Task, *types.ApiError) {
	panic("not implemented") // TODO: Implement
}

func (i *TaskStore) Delete(ctx context.Context, id string) *types.ApiError {
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("tasks/Delete.sql")
	if err != nil {
		return types.InternalServerError(err, uri)
	}
	_, err = i.db.Exec(loadedSql, id)
	if err != nil {
		return types.InternalServerError(err, uri)
	}
	return nil
}

func (i *TaskStore) createWhitelist() Whitelist {
	incWhitelist := map[string][]string{
		"Priority": {string(types.Low), string(types.Medium), string(types.High), string(types.Critical)},
		"State":    {string(types.Backlog), string(types.Doing), string(types.Review), string(types.Done)},
	}
	return incWhitelist
}
