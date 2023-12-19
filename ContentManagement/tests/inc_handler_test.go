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

func (t *SuiteTest) Test1GetAllInc() {
	res, err := http.Get(fmt.Sprintf("%s/incident", t.ApiServer))
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(fmt.Errorf("GetAllIncs - FAILED - expected status code 200 but got %d\n", res.StatusCode))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	ACTUAL := string(body)
	EXPECTED := "[{\"id\":\"b595bee6-d8c6-4b3b-b072-1e45c2103002\",\"name\":\"RapidResponse Case 2\",\"severity\":\"Critical\",\"status\":\"Pending\",\"type\":{\"id\":1,\"name\":\"CSIRTaaS\"}},{\"id\":\"b595bee6-d8c7-4b3b-b071-1e45c2103002\",\"name\":\"RapidResponse Case 1\",\"severity\":\"Low\",\"status\":\"Pending\",\"type\":{\"id\":1,\"name\":\"CSIRTaaS\"}}]"
	if ACTUAL != EXPECTED {
		t.Error(fmt.Errorf("GetAllIncs - FAILED - Expected %v but got %v", EXPECTED, ACTUAL))
	}
}

func (t *SuiteTest) Test2CreateIncident() {
	inc_marshalled, err := json.Marshal(&api.CreaIncReq{
		Name:         "IncidentTwo",
		Severity:     "Low",
		IncidentType: 1,
	})
	if err != nil {
		t.Error(err)
	}
	url := fmt.Sprintf("%s/incident", t.ApiServer)
	res, err := http.Post(url, "application/json", bytes.NewReader(inc_marshalled))
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(fmt.Errorf("CreatIncident - FAILED - expected status code 200 but got %d\n", res.StatusCode))
	}
	defer res.Body.Close()
	var inc incident.Incident

	json.NewDecoder(res.Body).Decode(&inc)
	EXPECTED_STATUS := incident.Open
	if inc.Status != EXPECTED_STATUS {
		t.Error(fmt.Errorf("CreateIncident - FAILED - expected Incident Status: %s but got %s", EXPECTED_STATUS, inc.Status))
	}
}

func (t *SuiteTest) Test3GetIncByID() {
	incID := "b595bee6-d8c7-4b3b-b071-1e45c2103002"
	url := fmt.Sprintf("%s/incident/%s", t.ApiServer, incID)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		t.Error(err)
		return
	}
	t.Assertions.True(res.StatusCode == 200, fmt.Sprintf("GetIncByID - FAILED - expected status code 200 but got %d\n", res.StatusCode))

	defer res.Body.Close()
	var inc incident.Incident

	json.NewDecoder(res.Body).Decode(&inc)
	t.Assertions.Equal(fmt.Sprint(inc.Id), incID, fmt.Sprintf("GetIncByID - FAILED - expected IncidentType ID: %s but got %d", incID, inc.Id))
}

func (t *SuiteTest) Test4DeleteIncident() {
	incID := "b595bee6-d8c7-4b3b-b071-1e45c2103002"
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/incident/%s", t.ApiServer, incID), nil)
	if err != nil {
		t.Error(err)
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Assertions.True(res.StatusCode == 200, fmt.Sprintf("DeleteInc - FAILED - expected status code 200 but got %d\n", res.StatusCode))

	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}
	var ACTUAL map[string]any
	json.Unmarshal(bytes, &ACTUAL)
	EXPECTED := fmt.Sprintf(`Incident with ID: %s was deleted`, incID)
	t.Assertions.Equal(EXPECTED, ACTUAL["Message"], fmt.Sprintf("DeleteInc - FAILED - expected DeleteMessage: '%s' but got %s", EXPECTED, ACTUAL["Message"]))
}

func (t *SuiteTest) Test1GetAllIncTypes() {
	url := fmt.Sprintf("%s/incidenttype", t.ApiServer)
	res, err := http.Get(url)
	if err != nil {
		t.Error(err)
		return
	}

	t.Assertions.Equal(res.StatusCode, http.StatusOK, fmt.Sprintf("GetAllIncType - FAILED - expected status code 200 but got %d\n", res.StatusCode))

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}

	ACTUAL := string(body)
	EXPECTED := "[{\"id\":1,\"name\":\"CSIRTaaS\"},{\"id\":2,\"name\":\"RapidResponse\"}]\n"
	t.Assertions.Equal(EXPECTED, ACTUAL, fmt.Sprintf("GetAllIncs - FAILED - Expected %v but got %v", EXPECTED, ACTUAL))
}

func (t *SuiteTest) Test2CreateIncType() {
	url := fmt.Sprintf("%s/incidenttype", t.ApiServer)
	test_name := "TEST_TYPE"
	inc_marshalled, err := json.Marshal(&api.CreIncTypeReq{
		Name: test_name,
	})
	if err != nil {
		t.Error(err)
		return
	}
	res, err := http.Post(url, "application/json", bytes.NewReader(inc_marshalled))
	if err != nil {
		t.Error(err)
		return
	}
	t.Assertions.Equal(res.StatusCode, http.StatusOK, fmt.Sprintf("CreatIncidentType - FAILED - expected status code 200 but got %d\n", res.StatusCode))

	defer res.Body.Close()
	var inc incident.IncidentType
	json.NewDecoder(res.Body).Decode(&inc)
	t.Assertions.Equal(inc.Name, test_name, fmt.Sprintf("CreatIncidentType - FAILED - expected type name %s but got %d\n", test_name, res.StatusCode))
}

func (t *SuiteTest) Test3GetIncTypeByID() {
	url := fmt.Sprintf("%s/incidenttype", t.ApiServer)
	typeID := "1"
	res, err := http.Get(fmt.Sprintf("%s/%s", url, typeID))
	if err != nil {
		t.Error(err)
	}
	t.Assertions.Equal(res.StatusCode, http.StatusOK, fmt.Sprintf("GetIncTypeByID - FAILED - expected status code 200 but got %d\n", res.StatusCode))

	defer res.Body.Close()
	var inc incident.IncidentType

	json.NewDecoder(res.Body).Decode(&inc)

	t.Assertions.Equal(fmt.Sprint(inc.Id), typeID, fmt.Sprintf("getIncTypeByID - FAILED - expected IncidentType ID: %s but got %d", typeID, inc.Id))

}

func (t *SuiteTest) Test4DeleteIncType() {
	url := fmt.Sprintf("%s/incidenttype", t.ApiServer)
	typeID := "3"
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", url, typeID), nil)
	if err != nil {
		t.Error(err)
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Assertions.Equal(res.StatusCode, http.StatusOK, fmt.Sprintf("DeleteIncType - FAILED - expected status code 200 but got %d\n", res.StatusCode))

	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}
	var ACTUAL map[string]any
	json.Unmarshal(bytes, &ACTUAL)
	EXPECTED := fmt.Sprintf(`IncidentType with ID: %s was deleted`, typeID)
	t.Assertions.Equal(EXPECTED, ACTUAL["Message"], fmt.Sprintf("DeleteIncType - FAILED - expected DeleteMessage: '%s' but got %s", EXPECTED, ACTUAL["Message"]))
}
