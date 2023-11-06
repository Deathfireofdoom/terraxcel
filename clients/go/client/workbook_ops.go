package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Deathfireofdoom/terraxcel/common/models"
)

func (c *Client) CreateWorkbook(workbook *models.Workbook) (*models.Workbook, error) {
	req, err := c.NewRequest(http.MethodPost, "/workbook", workbook)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		createdWorkbook := &models.Workbook{}
		err := json.NewDecoder(resp.Body).Decode(createdWorkbook)
		if err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}
		return createdWorkbook, nil
	}

	return nil, fmt.Errorf("received non-201 status code: %d", resp.StatusCode)
}

func (c *Client) ReadWorkbook(WorkbookID string) (*models.Workbook, error) {
	req, err := c.NewRequest(http.MethodGet, "/workbook/"+WorkbookID, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		Workbook := &models.Workbook{}
		err := json.NewDecoder(resp.Body).Decode(Workbook)
		if err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}
		return Workbook, nil
	}
	return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
}

func (c *Client) DeleteWorkbook(workbook models.Workbook) error {
	req, err := c.NewRequest(http.MethodDelete, "/workbook/"+workbook.ID, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
}

func (c *Client) UpdateWorkbook(workbook *models.Workbook) (*models.Workbook, error) {
	req, err := c.NewRequest(http.MethodPut, "/workbook/"+workbook.ID, workbook)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		workbook := &models.Workbook{}
		err := json.NewDecoder(resp.Body).Decode(&workbook)
		if err != nil {
			return nil, fmt.Errorf("failed decoding response: %w", err)
		}
		return workbook, nil
	}

	return nil, fmt.Errorf("receieved non-200 status code: %d", resp.StatusCode)
}
