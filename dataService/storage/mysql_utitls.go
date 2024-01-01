package storage

import (
	"data-service/types"
	"fmt"
)

func (s *MySqlStorage) CheckWhitelist(entity, key, value string) bool {
	switch entity {
	case "incidents":
		whiteList := s.createIncidentWhitelist()
		return CheckWhitelist(key, value, whiteList)
	default:
		fmt.Println(fmt.Errorf("entity: %s not implemented", entity))
		return false
	}
}

func (s *MySqlStorage) createIncidentWhitelist() Whitelist {
	sql, err := LoadRawSQL("incidenttypes/GetAll.sql")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := s.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Close()
	iTs := []string{}

	for res.Next() {
		var iT types.IncidentType
		err = iT.ScanTo(res.Scan)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		iTs = append(iTs, iT.Name)
	}
	incWhitelist := map[string][]string{
		"Severity": {string(types.Low), string(types.Medium), string(types.High), string(types.Critical)},
		"Status":   {string(types.Pending), string(types.Open), string(types.Active), string(types.Closed)},
		"Type":     iTs,
	}
	return incWhitelist
}
