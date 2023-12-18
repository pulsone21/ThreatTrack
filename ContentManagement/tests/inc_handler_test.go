package tests

import (
	"ContentManagement/api"
	"ContentManagement/api/models/incident"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (t *SuiteTest) TestHello() {
	fmt.Println("Hello Word")
}

func (t *SuiteTest) TestIncidentHandler() {
	fmt.Println("Starting incident handler testing")
	baseUrl := fmt.Sprintf("%s/incident", t.serverAdress)

	// Get All Incident
	err := getAllIncs(baseUrl)
	if err != nil {
		t.Error(err)
	}

	// Creating a Incident
	incID, err := createInc(baseUrl)
	if err != nil {
		t.Error(err)
	}

	// Get Specific Incident
	if err := getIncByID(baseUrl, incID); err != nil {
		t.Error(err)
	}

	// Delete Incident
	if err := deleteInc(baseUrl, incID); err != nil {
		t.Error(err)
	}

}

func (t *SuiteTest) TestIncidentTypeHandler() {
	// Get All Incident Types
	// Creating a Incident Type
	// Get specficic Type
	// Delete a Incident Type
	fmt.Println("Starting incidentType handler testing")
	baseUrl := fmt.Sprintf("%s/incidenttype", t.serverAdress)

	// Get All Incident
	err := getAllIncTypes(baseUrl)
	if err != nil {
		t.Error(err)
	}

	// Creating a Incident
	incID, err := createIncType(baseUrl)
	if err != nil {
		t.Error(err)
	}

	// Get Specific Incident
	if err := getIncTypeByID(baseUrl, incID); err != nil {
		t.Error(err)
	}

	// Delete Incident
	if err := deleteIncType(baseUrl, incID); err != nil {
		t.Error(err)
	}
}

func getAllIncs(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GetAllIncs - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	ACTUAL := string(body)
	EXPECTED := "[{\"id\":\"b595bee6-d8c6-4b3b-b072-1e45c2103002\",\"name\":\"RapidResponse Case 2\",\"severity\":\"Critical\",\"status\":\"Pending\",\"type\":{\"id\":1,\"name\":\"CSIRTaaS\"}},{\"id\":\"b595bee6-d8c7-4b3b-b071-1e45c2103002\",\"name\":\"RapidResponse Case 1\",\"severity\":\"Low\",\"status\":\"Pending\",\"type\":{\"id\":1,\"name\":\"CSIRTaaS\"}}]"
	if ACTUAL != EXPECTED {
		return fmt.Errorf("GetAllIncs - FAILED - Expected %v but got %v", EXPECTED, ACTUAL)
	}
	fmt.Println("GetAllIncs - Passed")
	return nil
}

func getAllIncTypes(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GetAllIncType - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	ACTUAL := string(body)
	EXPECTED := "[{\"id\":1,\"name\":\"CSIRTaaS\"},{\"id\":2,\"name\":\"RapidResponse\"}]"
	if ACTUAL != EXPECTED {
		return fmt.Errorf("GetAllIncs - FAILED - Expected %v but got %v", EXPECTED, ACTUAL)
	}
	fmt.Println("GetAllIncTypes - Passed")
	return nil
}

func createInc(url string) (string, error) {
	inc_marshalled, err := json.Marshal(&api.CreaIncReq{
		Name:         "IncidentTwo",
		Severity:     "Low",
		IncidentType: 1,
	})
	if err != nil {
		return "", err
	}
	res, err := http.Post(url, "application/json", bytes.NewReader(inc_marshalled))
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("CreatIncident - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var inc incident.Incident

	json.NewDecoder(res.Body).Decode(&inc)
	EXPECTED_STATUS := incident.IncOpen
	if inc.Status != EXPECTED_STATUS {
		return "", fmt.Errorf("CreateIncident - FAILED - expected Incident Status: %s but got %s", EXPECTED_STATUS, inc.Status)
	}
	fmt.Printf("CreateIncident - PASSED - Incident with ID: '%s' was created. \n", inc.Id)
	return inc.Id.String(), nil
}

func createIncType(url string) (string, error) {
	inc_marshalled, err := json.Marshal(&api.CreIncTypeReq{
		Name: "TEST_TYPE",
	})
	if err != nil {
		return "", err
	}
	res, err := http.Post(url, "application/json", bytes.NewReader(inc_marshalled))
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("CreatIncidentType - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var inc incident.IncidentType
	json.NewDecoder(res.Body).Decode(&inc)
	fmt.Printf("CreatIncidentType - PASSED - Incident with ID: '%d' was created. \n", inc.Id)
	return fmt.Sprint(inc.Id), nil
}

func getIncByID(baseUrl, incID string) error {
	res, err := http.Get(fmt.Sprintf("%s/%s", baseUrl, incID))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GetIncByID - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var inc incident.Incident

	json.NewDecoder(res.Body).Decode(&inc)

	if inc.Id.String() != incID {
		return fmt.Errorf("GetIncByID - FAILED - expected IncidentType ID: %s but got %d", incID, inc.Id)
	}
	fmt.Printf("GetIncByID - PASSED - Got Incident with ID: '%s'\n", inc.Id)
	return nil
}

func getIncTypeByID(baseUrl, incID string) error {
	res, err := http.Get(fmt.Sprintf("%s/%s", baseUrl, incID))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("getIncTypeByID - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var inc incident.IncidentType

	json.NewDecoder(res.Body).Decode(&inc)

	if fmt.Sprint(inc.Id) != incID {
		return fmt.Errorf("getIncTypeByID - FAILED - expected IncidentType ID: %s but got %d", incID, inc.Id)
	}
	fmt.Printf("getIncTypeByID - PASSED - Got Incident with ID: '%d'\n", inc.Id)
	return nil
}

func deleteInc(baseUrl, incID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", baseUrl, incID), nil)
	if err != nil {
		return err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("DeleteInc - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var ACTUAL map[string]any
	json.NewDecoder(res.Body).Decode(&ACTUAL)
	EXPECTED := fmt.Sprintf(`Incident with ID: %s was deleted`, incID)
	if ACTUAL["message"] != EXPECTED {
		return fmt.Errorf("DeleteInc - FAILED - expected DeleteMessage: %s but got %s", EXPECTED, ACTUAL["message"])
	}
	fmt.Printf("GetIncByID - PASSED - '%s'\n", ACTUAL["message"])
	return nil
}

func deleteIncType(baseUrl, incID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", baseUrl, incID), nil)
	if err != nil {
		return err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("DeleteIncType - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var ACTUAL map[string]any
	json.NewDecoder(res.Body).Decode(&ACTUAL)
	EXPECTED := fmt.Sprintf(`IncidentType with ID: %s was deleted`, incID)
	if ACTUAL["message"] != EXPECTED {
		return fmt.Errorf("DeleteIncType - FAILED - expected DeleteMessage: %s but got %s", EXPECTED, ACTUAL["message"])
	}
	fmt.Printf("DeleteIncType - PASSED - '%s'\n", ACTUAL["message"])
	return nil
}
