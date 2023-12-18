package api

import (
	"ContentManagement/api/models/task"
	"ContentManagement/api/stores"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type TaskApi struct {
	Store *stores.TaskStore
}

type CreTask struct {
	Title       string
	Description string
	Assignee    string
	Incident    string
	Status      string
	Priority    string
}

func CreateTaskApi(db *sql.DB, ser *ApiServer) *TaskApi {
	tA := &TaskApi{
		Store: stores.NewTaskStore(db),
	}
	tA.createHandles(ser)
	fmt.Println("Created Task Api")
	return tA
}

func (api *TaskApi) createHandles(ser *ApiServer) {
	taskRouter := ser.Router.PathPrefix("/task").Subrouter()
	taskRouter.HandleFunc("", createHandleFunc(api.api_GetAllTasks)).Methods(http.MethodGet)
	taskRouter.HandleFunc("", createHandleFunc(api.api_CreateTask)).Methods(http.MethodPost)
	taskRouter.HandleFunc("/{id}", createHandleFunc(api.api_GetTaskByID)).Methods(http.MethodGet)
	taskRouter.HandleFunc("/{id}", createHandleFunc(api.api_UpdateTask)).Methods(http.MethodPut)
	taskRouter.HandleFunc("/{id}", createHandleFunc(api.api_DeleteTask)).Methods(http.MethodDelete)
}

// API FUNCTIONS FOR TASKS

func (s *TaskApi) api_GetAllTasks(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Requesting All Tasks API Handler")
	tasks, err := s.Store.GetAllTasks()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, tasks)
}

func (s *TaskApi) api_GetTaskByID(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Request for Getting Task by ID")
	id := mux.Vars(r)["id"]
	task, err := s.Store.GetTaskByID(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, task)
}

func (s *TaskApi) api_CreateTask(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Request for Creating Task")
	var tR CreTask
	json.NewDecoder(r.Body).Decode(&tR)
	incId, err1 := uuid.Parse(tR.Incident)
	usrId, err2 := uuid.Parse(tR.Assignee)
	if err := errors.Join(err1, err2); err != nil {
		return err
	}
	task := task.NewTask(tR.Title, tR.Description, usrId, incId, task.TaskPriority(tR.Priority), task.TaskSatus(tR.Status))
	return s.Store.CreateTask(task)

}

func (s *TaskApi) api_UpdateTask(w http.ResponseWriter, r *http.Request) error {
	// TODO Implement
	return fmt.Errorf("not implemented")
	// fmt.Println("Request for Updating Task")
	// id := mux.Vars(r)["id"]
	// var taskReq CreTask
	// json.NewDecoder(r.Body).Decode(&taskReq)
	// tT, err := s.Store.GetTaskTypeBy("id", fmt.Sprint(taskReq.TaskType))
	// if err != nil {
	// fmt.Println("error in get TaskType request")
	// return err
	// }
	// task, err := s.Store.UpdateTask(id, taskReq.Value, *tT, taskReq.IncidentIDs)
	// if err != nil {
	// return err
	// }
	// return writeJSON(w, http.StatusOK, task)
}

func (s *TaskApi) api_DeleteTask(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Request for Deleting Task")
	id := mux.Vars(r)["id"]
	err := s.Store.DeleteTask(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, "Task deleted successfully")
}
