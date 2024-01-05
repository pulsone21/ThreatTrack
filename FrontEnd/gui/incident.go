package gui

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"threattrack/entities"
	"threattrack/utils"
)

type IncidentHandler struct {
	backendAdress string
	templatePath  string
}

type incTableViewData struct {
	Incidents []entities.Incident
}


func (iH *IncidentHandler) createHandles(s *Server) {
	s.Router.HandleFunc("/incidentTable/", utils.CreateHandleFunc(iH.serveIncidentTable)).Methods("GET")
	s.Router.HandleFunc("/incident/summary", utils.CreateHandleFunc(iH.serveIncidentPage)).Methods("GET")
	s.Router.HandleFunc("/incident/worklog", utils.CreateHandleFunc(iH.serveIncidentWorklog)).Methods("GET")
	s.Router.HandleFunc("/incident/planing", utils.CreateHandleFunc(iH.serveIncidentPlaning)).Methods("GET")
	s.Router.HandleFunc("/incident/iocView", utils.CreateHandleFunc(iH.serveIncidentiocView)).Methods("GET")
}

func CreateIncidentHandler(ser *Server, backendBase string) *IncidentHandler {
	wd, _ := os.Getwd()
	iH := &IncidentHandler{
		backendAdress: fmt.Sprintf("%s/incidents", backendBase),
		templatePath:  "templates/incident",
	}
	fmt.Printf("%s/%s\n", wd, iH.templatePath)
	iH.createHandles(ser)
	return iH
}

func (iH *IncidentHandler) serveIncidentTable(ctx context.Context, w http.ResponseWriter, r *http.Request) (*entities.ApiResponse,  *entities.ApiError) {
	fmt.Printf("\nrequesting backend with %s \n", iH.backendAdress)
	res, err := http.Get(iH.backendAdress)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		// Handle the Error
		return http.ErrAbortHandler
	}
	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/incidentTable.html", iH.templatePath))
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	var data struct {
		StatusCode int
		RequestUrl string
		Data []entities.Incident
	}
	fmt.Println("defining the struct")
	if err = json.Unmarshal(resbody, &data); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("struct unmarshaled")
	fmt.Println(data)
	return tmpl.Execute(w, incTableViewData{
		Incidents: data.Data,
	})
}

func (iH *IncidentHandler) serveIncidentPage(ctx context.Context, w http.ResponseWriter, r *http.Request) (*entities.ApiResponse,  *entities.ApiError) {
	incId := r.URL.Query().Get("id")
	url := fmt.Sprintf("%s/%s", iH.backendAdress, incId)
	fmt.Printf("\nrequesting backend with %s \n", url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	var inc entities.Incident
	if err = json.NewDecoder(res.Body).Decode(&inc); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/incident.html", iH.templatePath))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, inc)
}

func (iH *IncidentHandler) serveIncidentWorklog(ctx context.Context, w http.ResponseWriter, r *http.Request) (*entities.ApiResponse,  *entities.ApiError) {
	incId := r.URL.Query().Get("id")
	url := fmt.Sprintf("%s/%s", iH.backendAdress, incId)
	fmt.Printf("\nrequesting backend with %s \n", url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	var inc entities.Incident
	if err = json.NewDecoder(res.Body).Decode(&inc); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/incidentWorklogs.html", iH.templatePath))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, inc)
}

func (iH *IncidentHandler) serveIncidentPlaning(ctx context.Context, w http.ResponseWriter, r *http.Request) (*entities.ApiResponse,  *entities.ApiError) {
	incId := r.URL.Query().Get("id")
	url := fmt.Sprintf("%s/%s", iH.backendAdress, incId)
	fmt.Printf("\nrequesting backend with %s \n", url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	var inc entities.Incident
	if err = json.NewDecoder(res.Body).Decode(&inc); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/incidentPlaning.html", iH.templatePath))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, inc)
}

func (iH *IncidentHandler) serveIncidentiocView(ctx context.Context, w http.ResponseWriter, r *http.Request) (*entities.ApiResponse,  *entities.ApiError) {
	incId := r.URL.Query().Get("id")
	url := fmt.Sprintf("%s/%s", iH.backendAdress, incId)
	fmt.Printf("\nrequesting backend with %s \n", url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	var inc entities.Incident
	if err = json.NewDecoder(res.Body).Decode(&inc); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/incidentIOCView.html", iH.templatePath))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, inc)
}
