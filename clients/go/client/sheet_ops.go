package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Deathfireofdoom/terraxcel/common/models"
)

func (c *Client) CreateSheet(sheet *models.Sheet) (*models.Sheet, error) {
	req, err := c.NewRequest(http.MethodPost, "/workbook/"+sheet.WorkbookID+"/sheet", sheet)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		var sheet *models.Sheet
		err := json.NewDecoder(resp.Body).Decode(&sheet)
		if err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}
		return sheet, nil
	}

	return nil, fmt.Errorf("receieved non-200 status code: %d", resp.StatusCode)
}

func (c *Client) ReadSheet(sheetID, workbookID string) (*models.Sheet, error) {
	req, err := c.NewRequest(http.MethodGet, "/workbook/"+workbookID+"/sheet/"+sheetID, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var sheet *models.Sheet
		err := json.NewDecoder(resp.Body).Decode(&sheet)
		if err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}
		return sheet, nil
	}

	return nil, fmt.Errorf("receieved non-200 status code: %d", resp.StatusCode)
}

func (c *Client) DeleteSheet(sheet *models.Sheet) error {
	req, err := c.NewRequest(http.MethodDelete, "/workbook/"+sheet.WorkbookID+"/sheet/"+sheet.ID, nil)
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

	return fmt.Errorf("receieved non-200 status code: %d", resp.StatusCode)
}

func (c *Client) UpdateSheet(sheet *models.Sheet) (*models.Sheet, error) {
	req, err := c.NewRequest(http.MethodPut, "/workbook/"+sheet.WorkbookID+"/sheet/"+sheet.ID, sheet)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var sheet *models.Sheet
		err := json.NewDecoder(resp.Body).Decode(&sheet)
		if err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}
		return sheet, nil
	}

	return nil, fmt.Errorf("receieved non-200 status code: %d", resp.StatusCode)
}
