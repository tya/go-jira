package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCycleService_GetList_Success(t *testing.T) {
	setup()
	defer teardown()

	raw, err := ioutil.ReadFile("./mocks/all_cycles.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(cycleEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, cycleEndpoint)
		fmt.Fprint(w, string(raw))
	})

	opts := &CycleListOptions{0, 0}
	cycles, _, err := testClient.Cycle.GetList(opts)
	if err != nil {
		t.Errorf("Error given: %v", err)
	}
	if len(cycles) != 2 {
		t.Errorf("Expected %d cycles but got %d", 2, len(cycles))
	}
}

func TestCycleService_GetList_NoList(t *testing.T) {
	setup()
	defer teardown()

	raw, err := ioutil.ReadFile("./mocks/no_cycle.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(cycleEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, cycleEndpoint)
		fmt.Fprint(w, string(raw))
	})

	cycles, _, err := testClient.Cycle.GetList(nil)
	if cycles != nil {
		t.Errorf("Actual cycle list has %d entries but expected nil", len(cycles))
	}
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestCycleService_GetList_UrlError(t *testing.T) {
	setup()
	backup := cycleEndpoint
	defer func() {
		cycleEndpoint = backup
		teardown()
	}()

	// set an invlid cycleEndpoint to trigger a url error
	cycleEndpoint = "\r"

	opts := new(CycleListOptions)
	cycles, _, err := testClient.Cycle.GetList(opts)
	if cycles != nil {
		t.Errorf("Actual cycle list has %d entries but expected nil", len(cycles))
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestCycleService_GetList_RequestError(t *testing.T) {
	setup()
	backup := cycleEndpoint
	defer func() {
		cycleEndpoint = backup
		teardown()
	}()

	// set an invalid cycleEndpoint to trigger a request error
	cycleEndpoint = "\r"

	// options must be nil to trigger error in request creation
	cycles, _, err := testClient.Cycle.GetList(nil)
	if cycles != nil {
		t.Errorf("Actual cycle list has %d entries but expected nil", len(cycles))
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestCycleService_GetList_HttpError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(cycleEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, cycleEndpoint)
		w.WriteHeader(http.StatusNotFound)
	})

	cycles, _, err := testClient.Cycle.GetList(nil)
	if cycles != nil {
		t.Errorf("Actual cycle list has %d entries but expected nil", len(cycles))
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestCycleService_GetList_UnmarshalError(t *testing.T) {
	setup()
	defer teardown()

	raw, err := ioutil.ReadFile("./mocks/error_cycles.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(cycleEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, cycleEndpoint)
		fmt.Fprint(w, string(raw))
	})

	opts := &CycleListOptions{-1, -1}
	cycles, _, err := testClient.Cycle.GetList(opts)
	if cycles != nil {
		t.Errorf("Actual cycle list has %d entries but expected nil", len(cycles))
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestCycleService_Create_RequestError(t *testing.T) {
	setup()
	backup := cycleEndpoint
	defer func() {
		cycleEndpoint = backup
		teardown()
	}()

	// set an invalid cycleEndpoint to trigger a request error
	cycleEndpoint = "\r"

	cycle := new(Cycle)
	reply, _, err := testClient.Cycle.Create(cycle)
	if reply != nil {
		t.Errorf("Expected cycle create reply to be nil, %v", reply)
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestCycleService_Create_HttpError(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc(cycleEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, cycleEndpoint)
		w.WriteHeader(http.StatusNotFound)
	})

	cycle := new(Cycle)
	reply, _, err := testClient.Cycle.Create(cycle)
	if reply != nil {
		t.Errorf("Expected cycle create reply to be nil, %v", reply)
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestCycleService_Create(t *testing.T) {
	setup()
	defer teardown()

	raw, err := ioutil.ReadFile("./mocks/cycle_create.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(cycleEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, cycleEndpoint)
		fmt.Fprint(w, string(raw))
	})

	cycle := new(Cycle)
	reply, _, err := testClient.Cycle.Create(cycle)
	if err != nil {
		t.Errorf("Error given: %v", err)
	}
	if reply == nil {
		t.Error("Expected creation reply. Reply is nil")
		return
	}
	if reply.ID != "54" {
		t.Errorf("Expected id 54 but got %s", reply.ID)
	}
}
