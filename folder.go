package jira

import (
	"context"
	"fmt"
	"net/http"
)

const (
	FolderListError   = "Folder List Error"
	FolderCreateError = "Folder Create Error"
)

var (
	folderEndpointFormat = "/rest/zapi/latest/cycle/%d/folders"
	folderEndpoint       = "/rest/zapi/latest/folder/create"
)

type FolderService struct {
	client *Client
}

type Folder struct {
	CycleID     int    `json:"cycleId,omitempty"`
	CycleName   string `json:"cycleName,omitempty"`
	Description string `json:"folderDescription,omitempty"`
	FolderID    int    `json:"folderId,omitempty"`
	FolderName  string `json:"folderName,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ProjectID   int    `json:"projectId,omitempty"`
	ProjectKey  string `json:"projectKey,omitempty"`
	VersionID   int    `json:"versionId,omitempty"`
	VersionName string `json:"versionName,omitempty"`
}

// FolderListOptions parameters to the FolderService.GetList
type FolderListOptions struct {
	ProjectID int `url:"projectId"`
	VersionID int `url:"versionId"`
}

// GetListWithContext gets a list of folders
func (s *FolderService) GetListWithContext(ctx context.Context, cycleId int, opts *FolderListOptions) ([]Folder, *Response, error) {
	// setup url
	endpoint := fmt.Sprintf(folderEndpointFormat, cycleId)
	url, err := addOptions(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	// setup request from url
	req, err := s.client.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	// do request
	var folders []Folder
	resp, err := s.client.Do(req, &folders)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}
	return folders, resp, nil
}

// GetList wraps GetListWithContext using the background context
func (s *FolderService) GetList(cycleId int, opt *FolderListOptions) ([]Folder, *Response, error) {
	return s.GetListWithContext(context.Background(), cycleId, opt)
}

// CreateWithContext creates a folders
func (s *FolderService) CreateWithContext(ctx context.Context, folder *Folder) (*Folder, *Response, error) {
	// create request
	req, err := s.client.NewRequestWithContext(ctx, http.MethodPost, folderEndpoint, folder)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", FolderCreateError, err)
	}

	// do request, unmarshal reply
	reply := new(Folder)
	resp, err := s.client.Do(req, reply)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", FolderCreateError, err)
	}
	return reply, resp, nil
}

// Create a folder wraps CreateWithContext using context.Background()
func (s *FolderService) Create(folder *Folder) (*Folder, *Response, error) {
	return s.CreateWithContext(context.Background(), folder)
}
