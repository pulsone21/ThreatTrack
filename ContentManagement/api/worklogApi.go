package api

import (
	"ContentManagement/api/models/worklog"
	"ContentManagement/api/stores"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type WorklogApi struct {
	Store *stores.WorklogStore
}

type CreatWorklogRequest struct {
	WriterId   string
	IncidentId string
	Content    string
}

func CreateWorklogApi(db *sql.DB, ser *ApiServer) *WorklogApi {
	wA := &WorklogApi{
		Store: stores.NewWorklogStore(db),
	}
	wA.createHandles(ser)
	return wA
}

func (s *WorklogApi) createHandles(ser *ApiServer) {
	//ser.HandleFunc("/worklog", createHandleFunc(s.api_CreateWorklog)).Methods("POST")
	//ser.HandleFunc("/worklog", createHandleFunc(s.api_GetWorklogs)).Methods("GET")
	//ser.HandleFunc("/worklog", createHandleFunc(s.api_DeleteWorklog)).Methods("DELETE")
	//ser.HandleFunc("/worklog/{id}", createHandleFunc(s.api_UpdateWorklog)).Methods("PUT")
}

func (a *WorklogApi) api_CreateWorklog(w http.ResponseWriter, r *http.Request) error {
	var WorReq CreatWorklogRequest
	json.NewDecoder(r.Body).Decode(&WorReq)
	var writer_id, inc_id uuid.UUID
	writer_id, err1 := uuid.Parse(WorReq.WriterId)
	inc_id, err2 := uuid.Parse(WorReq.IncidentId)
	if err := errors.Join(err1, err2); err != nil {
		return err
	}
	wl := worklog.NewWorklog(writer_id, inc_id, WorReq.Content)
	if err := a.Store.CreateWorklog(wl); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, wl)
}

func (a *WorklogApi) api_GetWorklogs(w http.ResponseWriter, r *http.Request) error {
	url_q := r.URL.Query()
	idRaw := ""
	var err error
	colName := "Id"
	if url_q.Has("incidentId") {
		idRaw = url_q.Get("incidentId")
		colName = "IncidentId"
	} else if url_q.Has("userId") {
		idRaw = url_q.Get("userId")
		colName = "UserId"
	} else if url_q.Has("worklogId") {
		idRaw = url_q.Get("worklogId")
		colName = "Id"
	} else {
		worklogs, err := a.Store.GetAllWorklogs()
		if err != nil {
			return err
		}
		return writeJSON(w, http.StatusOK, worklogs)
	}
	id, err := uuid.Parse(idRaw)
	if err != nil {
		return err
	}
	wl, err := a.Store.GetWorklogBy(id, colName)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, wl)
}

func (a *WorklogApi) api_UpdateWorklog(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("not implemented")
}

func (a *WorklogApi) api_DeleteWorklog(w http.ResponseWriter, r *http.Request) error {
	url_q := r.URL.Query()
	idRaw := ""
	var err error
	colName := "Id"
	if url_q.Has("incidentId") {
		idRaw = url_q.Get("incidentId")
		colName = "IncidentId"
	} else if url_q.Has("userId") {
		idRaw = url_q.Get("userId")
		colName = "UserId"
	} else if url_q.Has("worklogId") {
		idRaw = url_q.Get("worklogId")
		colName = "Id"
	} else {
		return writeJSON(w, http.StatusBadRequest, "Query for worklog not defined")
	}
	id, err := uuid.Parse(idRaw)
	if err != nil {
		return err
	}
	if err = a.Store.DeleteWorklog(id, colName); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, fmt.Sprintf(`Worklog with ID: %s was deleted`, id))
}
