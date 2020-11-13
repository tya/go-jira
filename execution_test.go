package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestExecutionService_Create_Success(t *testing.T) {
	setup()
	defer teardown()

	raw, err := ioutil.ReadFile("./mocks/execution_create.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(executionEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, executionEndpoint)
		fmt.Fprint(w, string(raw))
	})

	exe := new(Execution)
	exe, _, err = testClient.Execution.Create(exe)
	if err != nil {
		t.Errorf("Error given: %v", err)
	}
	if exe == nil {
		t.Error("Expected execution. Reply is nil")
		return
	}
	if exe.ID != 13377 {
		t.Errorf("Expected id 13377 but got %d", exe.ID)
	}
}

func TestExecutionService_Create_RequestError(t *testing.T) {
	setup()
	backup := executionEndpoint
	defer func() {
		executionEndpoint = backup
		teardown()
	}()

	// set an invalid executionEndpoint to trigger a request error
	executionEndpoint = "\r"

	execution := new(Execution)
	reply, _, err := testClient.Execution.Create(execution)
	if reply != nil {
		t.Errorf("Expected execution create reply to be nil, %v", reply)
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestExecutionService_Create_HttpError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(executionEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, executionEndpoint)
		w.WriteHeader(http.StatusNotFound)
	})

	execution := new(Execution)
	reply, _, err := testClient.Execution.Create(execution)
	if reply != nil {
		t.Errorf("Expected execution create reply to be nil, %v", reply)
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestExecutionService_Execute_Success(t *testing.T) {
	setup()
	defer teardown()

	id := 13377
	status := "-1"

	raw, err := ioutil.ReadFile("./mocks/execution_execute.json")
	if err != nil {
		t.Error(err.Error())
	}
	endpoint := fmt.Sprintf(executeEndpointFormat, id)
	testMux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testRequestURL(t, r, endpoint)

		// verify put
		es := new(ExecutionStatus)
		err := json.NewDecoder(r.Body).Decode(es)
		if err != nil {
			t.Error(err.Error())
		}
		if status != es.Status {
			t.Errorf("Server expected status %s but got %s", status, es.Status)
		}

		// send reply
		fmt.Fprint(w, string(raw))
	})

	es := &ExecutionStatus{Status: status}
	exe, _, err := testClient.Execution.Execute(id, es)
	if err != nil {
		t.Errorf("Error given: %v", err)
	}
	if exe == nil {
		t.Error("Expected execution. Reply is nil")
		return
	}
	if exe.ID != id {
		t.Errorf("Expected id %d but got %d", id, exe.ID)
	}
	if exe.ExecutionStatus != status {
		t.Errorf("Expected ExecutionStatus %s but got %s", status, exe.ExecutionStatus)
	}
}

func TestExecutionService_Execute_RequestError(t *testing.T) {
	setup()
	backup := executeEndpointFormat
	defer func() {
		executeEndpointFormat = backup
		teardown()
	}()

	// set an invalid executionEndpoint to trigger a request error
	executeEndpointFormat = "%d\r"

	id := 1000
	status := &ExecutionStatus{Status: "1"}
	reply, _, err := testClient.Execution.Execute(id, status)

	if reply != nil {
		t.Errorf("Expected execute reply to be nil, %v", reply)
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestExecutionService_Execute_HttpError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(executionEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, executionEndpoint)
		w.WriteHeader(http.StatusNotFound)
	})

	id := 1000
	status := &ExecutionStatus{Status: "1"}
	reply, _, err := testClient.Execution.Execute(id, status)

	if reply != nil {
		t.Errorf("Expected execute reply to be nil, %v", reply)
	}
	if err == nil {
		t.Errorf("No error given")
	}
}
