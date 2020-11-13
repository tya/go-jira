package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	CycleListError   = "Cycle List Error"
	CycleCreateError = "Cycle Create Error"
)

var (
	cycleEndpoint = "/rest/zapi/latest/cycle"
)

type CycleService struct {
	client *Client
}

type CycleCreateReply struct {
	ID              string `json:"id"`
	ResponseMessage string `json:"responseMessage"`
}

type Cycle struct {
	Build                string `json:"build,omitempty"`
	CreatedBy            string `json:"createdBy,omitempty"`
	CreatedByDisplay     string `json:"createdByDisplay,omitempty"`
	CreatedDate          string `json:"createdDate,omitempty"`
	CycleOrderID         int    `json:"cycleOrderId,omitempty"`
	Description          string `json:"description,omitempty"`
	EndDate              string `json:"endDate,omitempty"`
	Ended                string `json:"ended,omitempty"`
	Environment          string `json:"environment,omitempty"`
	Expand               string `json:"expand,omitempty"`
	ID                   int    `json:"id,omitempty,omitempty"`
	ModifiedBy           string `json:"modifiedBy,omitempty"`
	Name                 string `json:"name,omitempty"`
	ProjectID            int    `json:"projectId,omitempty"`
	ProjectKey           string `json:"projectKey,omitempty"`
	StartDate            string `json:"startDate,omitempty"`
	Started              string `json:"started,omitempty"`
	TotalCycleExecutions int    `json:"totalCycleExecutions,omitempty"`
	TotalDefects         int    `json:"totalDefects,omitempty"`
	TotalExecuted        int    `json:"totalExecuted,omitempty"`
	TotalExecutions      int    `json:"totalExecutions,omitempty"`
	TotalFolders         int    `json:"totalFolders,omitempty"`
	VersionID            int    `json:"versionId,omitempty"`
	VersionName          string `json:"versionName,omitempty"`
	ExecutionSummaries   struct {
		ExecutionSummary []interface{} `json:"executionSummary,omitempty"`
	} `json:"executionSummaries,omitempty"`
}

type ExecutionSummary struct {
	Count       int    `json:"count"`
	StatusKey   int    `json:"statusKey"`
	StatusName  string `json:"statusName"`
	StatusColor string `json:"statusColor"`
}

type ExecutionSummaries struct {
	ExecutionSummary []ExecutionSummary `json:"executionSummary"`
}

// CycleListOptions parameters to the CycleService.GetList
type CycleListOptions struct {
	ProjectID int `url:"projectId"`
	VersionID int `url:"versionId"`
}

// GetListWithContext gets a list of cycles
func (s *CycleService) GetListWithContext(ctx context.Context, opts *CycleListOptions) ([]Cycle, *Response, error) {
	url, err := addOptions(cycleEndpoint, opts)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", CycleListError, err)
	}
	req, err := s.client.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", CycleListError, err)
	}

	// get top level map (tlm)
	var tlm map[string]json.RawMessage
	resp, err := s.client.Do(req, &tlm)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", CycleListError, err)
	}
	defer resp.Body.Close()

	// loop over top level map (tml), if has int key and decoedes, return it
	var cycles []Cycle
	for k, rawJson := range tlm {
		// all cycle keys will convert to intergers (aka skip recordCount key)
		cycleId, err := strconv.Atoi(k)
		if err != nil {
			continue
		}

		// deocde cycle
		var cycle Cycle
		err = json.Unmarshal(rawJson, &cycle)
		if err != nil {
			return nil, nil, fmt.Errorf("%s: %w", CycleListError, err)
		}

		// rawJson did not include the cycle id, so set it
		cycle.ID = cycleId
		cycles = append(cycles, cycle)
	}
	return cycles, resp, nil
}

// GetList wraps GetListWithContext using the background context
func (s *CycleService) GetList(opts *CycleListOptions) ([]Cycle, *Response, error) {
	return s.GetListWithContext(context.Background(), opts)
}

// CreateWithContext creates a cycles
func (s *CycleService) CreateWithContext(ctx context.Context, cycle *Cycle) (*CycleCreateReply, *Response, error) {
	// create request
	req, err := s.client.NewRequestWithContext(ctx, http.MethodPost, cycleEndpoint, cycle)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", CycleCreateError, err)
	}

	// do request, unmarshal reply
	reply := new(CycleCreateReply)
	resp, err := s.client.Do(req, reply)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", CycleCreateError, err)
	}
	return reply, resp, nil
}

// Create a cycle wraps CreateWithContext using context.Background()
func (s *CycleService) Create(cycle *Cycle) (*CycleCreateReply, *Response, error) {
	return s.CreateWithContext(context.Background(), cycle)
}
