package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFolderService_GetList_Success(t *testing.T) {
	setup()
	defer teardown()

	cycleId := 1
	endpoint := fmt.Sprintf(folderEndpointFormat, cycleId)

	raw, err := ioutil.ReadFile("./mocks/all_folders.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, endpoint)
		fmt.Fprint(w, string(raw))
	})

	//opts := new(FolderListOptions)
	opts := &FolderListOptions{1, 1}
	folders, _, err := testClient.Folder.GetList(cycleId, opts)
	if err != nil {
		t.Errorf("Error given: %v", err)
	}
	if len(folders) != 2 {
		t.Errorf("Expected %d folders but got %d", 2, len(folders))
	}
}

func TestFolderService_GetList_NoList(t *testing.T) {
	setup()
	defer teardown()

	cycleId := 1
	endpoint := fmt.Sprintf(folderEndpointFormat, cycleId)

	raw, err := ioutil.ReadFile("./mocks/no_folder.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, endpoint)
		fmt.Fprint(w, string(raw))
	})

	folders, _, err := testClient.Folder.GetList(cycleId, nil)
	if err != nil {
		t.Errorf("Error given: %v", err)
	}
	if len(folders) != 0 {
		t.Errorf("Expected %d folders but got %d", 0, len(folders))
	}
}

func TestFolderService_GetList_UrlError(t *testing.T) {
	setup()
	backup := folderEndpointFormat
	defer func() {
		folderEndpointFormat = backup
		teardown()
	}()

	// set an invlid endpoint to trigger a url error
	folderEndpointFormat = "%d\r"
	cycleId := 1

	opts := new(FolderListOptions)
	folders, _, err := testClient.Folder.GetList(cycleId, opts)
	if folders != nil {
		t.Errorf("Actual folder list has %d entries but expected nil", len(folders))
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestFolderService_GetList_RequestError(t *testing.T) {
	backup := folderEndpointFormat
	defer func() {
		folderEndpointFormat = backup
		teardown()
	}()

	// set an invlid endpoint to trigger a url error
	folderEndpointFormat = "%d\r"
	cycleId := 1000

	// options must be nil to trigger error in request creation
	folders, _, err := testClient.Folder.GetList(cycleId, nil)
	if folders != nil {
		t.Errorf("Actual folder list has %d entries but expected nil", len(folders))
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestFolderService_GetList_HttpError(t *testing.T) {
	setup()
	defer teardown()

	cycleId := 1
	endpoint := fmt.Sprintf(folderEndpointFormat, cycleId)

	testMux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, endpoint)
		w.WriteHeader(http.StatusNotFound)
	})

	folders, _, err := testClient.Folder.GetList(cycleId, nil)
	if folders != nil {
		t.Errorf("Actual folder list has %d entries but expected nil", len(folders))
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

// func TestFolderService_Create_RequestError(t *testing.T) {
// 	setup()
// 	backup := endpoint
// 	defer func() {
// 		endpoint = backup
// 		teardown()
// 	}()

// 	// set an invalid endpoint to trigger a request error
// 	endpoint = "\r"

// 	folder := new(Folder)
// 	reply, _, err := testClient.Folder.Create(folder)
// 	if reply != nil {
// 		t.Errorf("Expected folder create reply to be nil, %v", reply)
// 	}
// 	if err == nil {
// 		t.Errorf("No error given")
// 	}
// }

// func TestFolderService_Create_HttpError(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	testMux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, http.MethodPost)
// 		testRequestURL(t, r, endpoint)
// 		w.WriteHeader(http.StatusNotFound)
// 	})

// 	folder := new(Folder)
// 	reply, _, err := testClient.Folder.Create(folder)
// 	if reply != nil {
// 		t.Errorf("Expected folder create reply to be nil, %v", reply)
// 	}
// 	if err == nil {
// 		t.Errorf("No error given")
// 	}
// }

// func TestFolderService_Create(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	raw, err := ioutil.ReadFile("./mocks/folder_create.json")
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
// 	testMux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, http.MethodPost)
// 		testRequestURL(t, r, endpoint)
// 		fmt.Fprint(w, string(raw))
// 	})

// 	folder := new(Folder)
// 	reply, _, err := testClient.Folder.Create(folder)
// 	if err != nil {
// 		t.Errorf("Error given: %v", err)
// 	}
// 	if reply == nil {
// 		t.Error("Expected creation reply. Reply is nil")
// 		return
// 	}
// 	if reply.ID != "54" {
// 		t.Errorf("Expected id 54 but got %s", reply.ID)
// 	}
// }
