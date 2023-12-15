package stores

import (
	"ContentManagement/api/models"
	"database/sql"
	"errors"
	"fmt"
)

type IocStore struct {
	DB *sql.DB
}

func NewIocStore(db *sql.DB) *IocStore {
	iS := &IocStore{
		DB: db,
	}
	iS.createTables()
	return iS
}

func (s *IocStore) createTables() {
	fmt.Println("Starting to Create models.IOC Tables")
	iocTable, err1 := LoadSQL("iocs/CreateTable.sql")
	iocType, err2 := LoadSQL("iocs/type/CreateTable.sql")
	iocInc, err3 := LoadSQL("iocs/relations/CreateTable.sql")

	if err := errors.Join(err1, err2, err3); err != nil {
		panic(err.Error())
	}
	if _, err := s.DB.Exec(iocTable); err != nil {
		panic(err.Error())
	}
	if _, err := s.DB.Exec(iocType); err != nil {
		panic(err.Error())
	}
	if _, err := s.DB.Exec(iocInc); err != nil {
		panic(err.Error())
	}
}

// TODO Replace inline SQL with loading function
// models.IOC Functions

func (s *IocStore) GetAllIocs() (*[]models.IOC, error) {
	fmt.Println("Requesting all iocs")
	var iocs []models.IOC
	query := `SELECT iocs.id, iocs.value ,iocs.iocType, ioc_types.name AS iocTypeName, iocs.verdict
	FROM iocs
	LEFT JOIN ioc_types ON iocs.iocType = ioc_types.id;`

	res, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		var ioc models.IOC
		err := res.Scan(&ioc.Id, &ioc.Value, &ioc.IocType.Id, &ioc.IocType.Name, &ioc.Verdict)
		if err != nil {
			return nil, err
		}
		iocs = append(iocs, ioc)
	}
	return &iocs, nil
}

func (s *IocStore) CreateIOC(ioc models.IOC, incIds []string) (*models.IOC, error) {

	query := `INSERT INTO iocs (id, value, iocType) VALUES (?, ?, ?)`
	res, err := s.DB.Exec(query, ioc.Id, ioc.Value, ioc.IocType.Id)
	if err != nil {
		return nil, err
	}
	iocId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	fmt.Println(iocId)
	if len(incIds) > 0 {

		query = `INSERT INTO iocs_incidents (iocId, incidentId) VALUES`
		vals := []interface{}{}
		for _, row := range incIds {
			query += "(?, ?),"
			vals = append(vals, iocId, row)
		}
		// trim last comma
		query = query[0 : len(query)-1]
		//prepare the statement
		stmt, _ := s.DB.Prepare(query)
		//format all vals at once
		_, err = stmt.Exec(vals...)
		if err != nil {
			return nil, err
		}
	}
	return &ioc, err
}

func (s *IocStore) UpdateIoc() error {
	return fmt.Errorf("NOT IMPLEMENTED")
}

func (s *IocStore) DeleteIoc(name string) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}

func (s *IocStore) GetGetIocById(id string) (*models.IOC, error) {
	var ioc models.IOC
	query := `SELECT iocs.id, iocs.value ,iocs.iocType, ioc_types.name AS iocTypeName, iocs.verdict
				FROM iocs
				LEFT JOIN ioc_types ON iocs.iocType = ioc_types.id
				WHERE iocs.id = (?);`

	res := s.DB.QueryRow(query, id)

	err := res.Scan(&ioc.Id, &ioc.Value, &ioc.IocType.Id, &ioc.IocType.Name, &ioc.Verdict)
	if err != nil {
		return nil, err
	}
	fmt.Println(ioc)
	return &ioc, nil
}

// models.IOC Type Functions

func (s *IocStore) GetAllIocTypes() ([]models.IOCType, error) {
	var iTs []models.IOCType
	query := `select * from ioc_types;`
	res, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		var iT models.IOCType
		err = res.Scan(&iT.Id, &iT.Name)
		if err != nil {
			return nil, err
		}
		iTs = append(iTs, iT)
	}
	return iTs, nil
}

func (s *IocStore) GetIocTypeBy(col, value string) (*models.IOCType, error) {
	query := "select * from ioc_types where id = ?;" //fmt.Sprintf("select * from ioc_types where %s = ?;", col)
	fmt.Println(query)
	res := s.DB.QueryRow(query, value)
	fmt.Println(res)
	var iT models.IOCType
	err := res.Scan(&iT.Id, &iT.Name)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf(fmt.Sprintf("Couldn't found models.IOC with combination %s and %s", col, value))
		}
		return nil, err
	}
	return &iT, nil
}

func (s *IocStore) CreateIOCType(value string) (*models.IOCType, error) {
	query := `INSERT INTO ioc_types (name) VALUES (?)`
	//TODO maybe implement a check if we have the value already present?
	res, err := s.DB.Exec(query, value)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &models.IOCType{
		Id:   int(id),
		Name: value,
	}, nil
}

func (s *IocStore) DeleteIOCType(id string) error {
	_, err := s.DB.Exec(`DELETE FROM ioc_types WHERE id = ?`, id)
	return err
}

// models.IOC Incident Relations

func (s *IocStore) GetAllRelations() ([]models.Ioc_Incident, error) {
	var iocs []models.Ioc_Incident
	query := `select * from iocs_incidents;`
	res, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var ioc models.Ioc_Incident
		err := res.Scan(&ioc.Id, &ioc.IocId, &ioc.IncidentId)
		if err != nil {
			return nil, err
		}
		iocs = append(iocs, ioc)
	}
	return iocs, nil
}

func (s *IocStore) GetRelationBy(col, value string) ([]models.Ioc_Incident, error) {
	var iocs []models.Ioc_Incident
	query := fmt.Sprintf(`select * from iocs_incidents WHERE %s = ?;`, col)
	res, err := s.DB.Query(query, value)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var ioc models.Ioc_Incident
		err := res.Scan(&ioc.Id, &ioc.IocId, &ioc.IncidentId)
		if err != nil {
			return nil, err
		}
		iocs = append(iocs, ioc)
	}
	return iocs, nil
}

func (s *IocStore) DeleteRelationsBy(col, id string) error {
	query := fmt.Sprintf(`DELETE FROM iocs_incidents WHERE %s = ?`, col)
	_, err := s.DB.Exec(query, id)
	return err
}
