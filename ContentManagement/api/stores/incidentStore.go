package stores

import (
	"ContentManagement/api/models"
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
	fmt.Println("Try to create models.Incident Table")
	incTable, err1 := LoadSQL("incident/CreateTable.sql")
	incType, err2 := LoadSQL("incident/type/CreateTable.sql")
	if err := errors.Join(err1, err2); err != nil {
		panic(err.Error())
	}
	if _, err := s.DB.Exec(incTable); err != nil {
		panic(err.Error())
	}
	if _, err := s.DB.Exec(incType); err != nil {
		panic(err.Error())
	}
}

func (s *IncidentStore) GetAllIncidents() (*[]models.Incident, error) {
	var incidents []models.Incident
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
		var inci models.Incident

		err := res.Scan(&inci.Id, &inci.Name, &inci.Severity, &inci.Status, &inci.IncidentType.Id, &inci.IncidentType.Name)
		if err != nil {
			return nil, err
		}
		incidents = append(incidents, inci)
	}
	return &incidents, nil
}

func (s *IncidentStore) GetIncidentByID(value string) (*models.Incident, error) {
	var inc models.Incident
	query, err := LoadSQL("incident/GetById.sql")
	if err != nil {
		return nil, err
	}
	res := s.DB.QueryRow(query, value)

	err = res.Scan(&inc.Id, &inc.Name, &inc.Severity, &inc.IncidentType.Id, &inc.IncidentType.Name)
	if err != nil {
		return nil, err
	}
	return &inc, nil
}

func (s *IncidentStore) CreateIncident(inc *models.Incident) error {
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

func (s *IncidentStore) UpdateIncident(w http.ResponseWriter, r *http.Request) (*models.Incident, error) {
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

// models.Incident Types

func (s *IncidentStore) GetAllIncidentTypes() (*[]models.IncidentType, error) {
	var iTs []models.IncidentType
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
		var incT models.IncidentType

		err := res.Scan(&incT.Id, &incT.Name)
		if err != nil {
			return nil, err
		}
		iTs = append(iTs, incT)
	}
	return &iTs, nil
}

func (s *IncidentStore) GetIncidentTypeBy(column, value string) (*models.IncidentType, error) {
	var iT models.IncidentType
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
