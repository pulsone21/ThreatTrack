package storage

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strconv"
)

type QueryParameter struct {
	Limit  int
	Offset int
	Query  map[string]string
	// TODO maybe sorting is a good idea to implement
}

const (
	MAX_LIMIT int = 200
)

func NewQueryParameter(urlQuery url.Values, withParams bool) *QueryParameter {
	limit, offset, query := ExtractUrlQueries(urlQuery, withParams)
	return &QueryParameter{
		Limit:  limit,
		Offset: offset,
		Query:  query,
	}
}

func ExtractUrlQueries(uV url.Values, withAdditionalParams bool) (int, int, map[string]string) {
	limit, _ := strconv.Atoi(uV.Get("limit"))
	if limit == 0 {
		limit = MAX_LIMIT
	}
	offset, _ := strconv.Atoi(uV.Get("offset"))
	params := make(map[string]string)
	if withAdditionalParams {
		for k, v := range uV {
			if k == "offset" || k == "limit" {
				continue
			}
			params[k] = v[0]
		}
	}
	return limit, offset, params
}

// Reads an sql file from the given path location
//
// Subpath starts without an "/"
func LoadRawSQL(subpath string) (string, error) {
	fmt.Println(os.Getwd())
	path := path.Join("queries/", subpath)
	fmt.Println(path)
	f, ioErr := os.ReadFile(path)
	if ioErr != nil {
		// handle error.
		return "", ioErr
	}
	sql := string(f)
	return sql, nil
}

func FinalizeSQL(rawSql, entity string, queryParam QueryParameter) string {
	whereClause := "WHERE"
	i := 0
	for k, v := range queryParam.Query {
		if i > 0 {
			whereClause = fmt.Sprintf("%s AND %s.%s=%s", whereClause, entity, k, v)
		} else {
			whereClause = fmt.Sprintf("%s %s.%s=%s", whereClause, entity, k, v)
		}
		i++
	}
	sqlAddtion := ""
	if whereClause == "WHERE" {
		sqlAddtion = fmt.Sprintf("LIMIT %d OFFSET %d;", queryParam.Limit, queryParam.Offset)
	} else {
		sqlAddtion = fmt.Sprintf("%s\nLIMIT %d OFFSET %d;", whereClause, queryParam.Limit, queryParam.Offset)
	}
	return fmt.Sprintf(rawSql, sqlAddtion)
}

func CheckWhitelist(key, value string, whitelist Whitelist) bool {
	if whitelist == nil {
		return false
	}
	if whitelist[key] == nil {

	}

	return false
}
