package stores

import (
	"ContentManagement/api/models/incident"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type IncidentStore struct {
	DB *sql.DB
}

func NewIncidentStore(db *sql.DB) *IncidentStore {
	iS := &IncidentStore{
		DB: db,
	}
	iS.createTables()
	return iS
}

func (s *IncidentStore) createTables() {
	fmt.Println("Try to create incident.Incident Table")
	incTable, err1 := LoadSQL("incident/CreateTable.sql")
	incType, err2 := LoadSQL("incident/type/CreateTable.sql")
	if err := errors.Join(err1, err2); err != nil {
		panic(err.Error())
	}
	if _, err := s.DB.Exec(incType); err != nil {
		panic(err.Error())
	}
	if _, err := s.DB.Exec(incTable); err != nil {
		panic(err.Error())
	}
}

func (s *IncidentStore) GetAllIncidents() (*[]incident.Incident, error) {
	var incidents []incident.Incident
	query, err := LoadSQL("incident/GetAll.sql")
	if err != nil {
		return nil, err
	}
	res, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		inci, err := s.scanIncident(res.Scan)
		if err != nil {
			return nil, err
		}
		incidents = append(incidents, *inci)
	}
	return &incidents, nil
}

func (s *IncidentStore) GetIncidentByID(value string) (*incident.Incident, error) {
	query, err := LoadSQL("incident/GetById.sql")
	if err != nil {
		return nil, err
	}
	res := s.DB.QueryRow(query, value)

	inc, err := s.scanIncident(res.Scan)
	if err != nil {
		return nil, err
	}
	return inc, nil
}

func (s *IncidentStore) CreateIncident(inc *incident.Incident) error {
	query, err := LoadSQL("incident/Create.sql")
	if err != nil {
		return err
	}
	_, err = s.DB.Exec(query, inc.Id, inc.Name, inc.Severity, inc.IncidentType.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *IncidentStore) UpdateIncident(w http.ResponseWriter, r *http.Request) (*incident.Incident, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *IncidentStore) DeleteIncident(id string) error {
	sql, err := LoadSQL("incident/Delete.sql")
	if err != nil {
		return err
	}
	_, err = s.DB.Exec(sql, id)

	return err
}

// incident.Incident Types

func (s *IncidentStore) GetAllIncidentTypes() (*[]incident.IncidentType, error) {
	var iTs []incident.IncidentType
	query, err := LoadSQL("incident/type/GetAll.sql")
	if err != nil {
		return nil, err
	}
	res, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		var incT incident.IncidentType

		err := res.Scan(&incT.Id, &incT.Name)
		if err != nil {
			return nil, err
		}
		iTs = append(iTs, incT)
	}
	return &iTs, nil
}

func (s *IncidentStore) GetIncidentTypeBy(column, value string) (*incident.IncidentType, error) {
	var iT incident.IncidentType
	rawQ, err := LoadSQL("incident/type/GetBy.sql")
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(rawQ, column)
	fmt.Println(query)
	err = s.DB.QueryRow(query, value).Scan(&iT.Id, &iT.Name)
	if err != nil {
		fmt.Println("Found an Error")
		return nil, err
	}
	return &iT, nil
}

func (s *IncidentStore) CreateIncidentType(value string) (int64, error) {
	query, err := LoadSQL("incident/type/Create.sql")
	if err != nil {
		return -1, err
	}
	res, err := s.DB.Exec(query, value)
	if err != nil {
		return -1, err
	}
	_id, err := res.LastInsertId()
	return _id, err
}

func (s *IncidentStore) DeleteIncidentType(id string) error {
	query, err := LoadSQL("incident/type/Delete.sql")
	if err != nil {
		return err
	}
	_, err = s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *IncidentStore) scanIncident(scan ScanFunc) (*incident.Incident, error) {
	var inci incident.Incident
	err := scan(
		&inci.Id,
		&inci.Name,
		&inci.Severity,
		&inci.Status,
		&inci.IncidentType.Id,
		&inci.IncidentType.Name)
	if err != nil {
		return nil, err
	}
	return &inci, nil
}
