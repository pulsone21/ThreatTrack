package stores

import (
	"ContentManagement/api/models/worklog"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type WorklogStore struct {
	DB *sql.DB
}

func NewWorklogStore(db *sql.DB) *WorklogStore {
	iS := &WorklogStore{
		DB: db,
	}
	iS.createTable()
	return iS
}

func (s *WorklogStore) createTable() {
	fmt.Println("Try to create Worklog Table")
	query, err := LoadSQL("worklog/CreateTable.sql")
	if err != nil {
		panic(err.Error())
	}
	_, err = s.DB.Exec(query)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("No Issues found on table creation")
}

func (s *WorklogStore) GetAllWorklogs() (*[]worklog.Worklog, error) {
	var worklogs []worklog.Worklog
	query, err := LoadSQL("worklog/GetAll.sql")
	if err != nil {
		return nil, err
	}
	res, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		wl, err := s.scanWorklog(res.Scan)
		if err == nil {
			worklogs = append(worklogs, *wl)
		} else {
			fmt.Println(fmt.Errorf("%s", err.Error()))
		}
	}
	return &worklogs, nil
}

func (s *WorklogStore) GetWorklogBy(id uuid.UUID, colName string) (*worklog.Worklog, error) {
	query, err := LoadSQL(fmt.Sprintf("worklogs/GetBy%s.sql", colName))
	if err != nil {
		return nil, err
	}
	row := s.DB.QueryRow(query, id)
	wl, err := s.scanWorklog(row.Scan)
	if err != nil {
		return nil, err
	}
	return wl, nil
}

func (s *WorklogStore) CreateWorklog(wl *worklog.Worklog) error {

	query, err := LoadSQL("worklogs/Create.sql")
	if err != nil {
		return err
	}
	_, err = s.DB.Exec(query, wl.Id, wl.WriterId, wl.IncidentId, wl.Content, wl.Created_at)
	return err
}

func (s *WorklogStore) UpdateWorklog(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("not implemented")
}

func (s *WorklogStore) DeleteWorklog(id uuid.UUID, colName string) error {
	q, err := LoadSQL(fmt.Sprintf("worklogs/DeleteBy%s.sql", colName))
	if err != nil {
		return err
	}
	_, err = s.DB.Exec(q, id)
	return err
}

func (s *WorklogStore) scanWorklog(scan ScanFunc) (*worklog.Worklog, error) {
	var wl worklog.Worklog
	if err := scan(
		&wl.Id,
		&wl.WriterId,
		&wl.IncidentId,
		&wl.Content,
		&wl.Created_at,
	); err != nil {
		return nil, err
	}
	return &wl, nil
}
