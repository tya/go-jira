package jira

import (
	"context"
	"fmt"
	"net/http"
)

const (
	ExecutionRequestError = "Execution Request Error"
	ExecutionCreateError  = "Execution Create Error"
	ExecuteRequestError   = "Execute Request Error"
	ExecuteError          = "Execute Error"
)

var (
	executionEndpoint     = "/rest/zapi/latest/execution"
	executeEndpointFormat = "/rest/zapi/latest/execution/%d/execute"
)

type ExecutionService struct {
	client *Client
}

type Execution struct {
	ID              int    `json:"id"`
	Name            string `json:"name,omitempty"`
	AssignedTo      string `json:"assignedTo,omitempty"`
	Component       string `json:"component,omitempty"`
	CycleID         int    `json:"cycleId,omitempty"`
	CycleName       string `json:"cycleName,omitempty"`
	ExecutionStatus string `json:"executionStatus,omitempty"`
	FolderID        int    `json:"folderId,omitempty"`
	FolderName      string `json:"folderName,omitempty"`
	IssueID         int    `json:"issueId,omitempty"`
	IssueKey        string `json:"issueKey,omitempty"`
	Label           string `json:"label,omitempty"`
	ProjectID       int    `json:"projectId,omitempty"`
	ProjectKey      string `json:"projectKey,omitempty"`
	Summary         string `json:"summary,omitempty"`
	VersionID       int    `json:"versionId,omitempty"`
	VersionName     string `json:"versionName,omitempty"`
}

type ExecutionStatus struct {
	Status   string `json:"status,ommitempty"`
	Assignee string `json:"assignee,ommitempty"`
}

func (s *ExecutionService) CreateWithContext(ctx context.Context, exe *Execution) (*Execution, *Response, error) {
	// create request
	req, err := s.client.NewRequestWithContext(ctx, http.MethodPost, executionEndpoint, exe)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", ExecutionRequestError, err)
	}

	// do request, unmarshal reply
	//tlm := make(map[string]*Execution)
	var tlm map[string]*Execution
	resp, err := s.client.Do(req, &tlm)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", ExecutionCreateError, err)
	}

	// api suggestion: do not return a map here
	// return first execution
	for _, exe = range tlm {
		break
	}
	return exe, resp, nil
}

func (s *ExecutionService) Create(exe *Execution) (*Execution, *Response, error) {
	return s.CreateWithContext(context.Background(), exe)
}

func (s *ExecutionService) ExecuteWithContext(ctx context.Context, exeId int, status *ExecutionStatus) (*Execution, *Response, error) {
	// create request
	endpoint := fmt.Sprintf(executeEndpointFormat, exeId)
	req, err := s.client.NewRequestWithContext(ctx, http.MethodPut, endpoint, status)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", ExecuteRequestError, err)
	}

	exe := new(Execution)
	resp, err := s.client.Do(req, exe)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", ExecuteError, err)
	}
	return exe, resp, nil

}

func (s *ExecutionService) Execute(executionId int, status *ExecutionStatus) (*Execution, *Response, error) {
	return s.ExecuteWithContext(context.Background(), executionId, status)
}
