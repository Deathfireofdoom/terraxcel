package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Deathfireofdoom/terraxcel/common/models"
)

func (c *Client) CreateCell(cell *models.Cell) (*models.Cell, error) {
	req, err := c.NewRequest(http.MethodPost, "/workbook/"+cell.WorkbookID+"/sheet/"+cell.SheetID+"/cell", cell)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		var cell *models.Cell
		err := json.NewDecoder(resp.Body).Decode(&cell)
		if err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}
		return cell, nil
	}

	return nil, fmt.Errorf("receieved non-200 status code: %d", resp.StatusCode)
}

func (c *Client) ReadCell(cellID, sheetID, workbookID string) (*models.Cell, error) {
	req, err := c.NewRequest(http.MethodGet, "/workbook/"+workbookID+"/sheet/"+sheetID+"/cell/"+cellID, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var cell *models.Cell
		err := json.NewDecoder(resp.Body).Decode(&cell)
		if err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}
		return cell, nil
	}

	return nil, fmt.Errorf("receieved non-200 status code: %d", resp.StatusCode)
}

func (c *Client) DeleteCell(cell *models.Cell) error {
	req, err := c.NewRequest(http.MethodDelete, "/workbook/"+cell.WorkbookID+"/sheet/"+cell.SheetID+"/cell/"+cell.ID, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending requets: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var cell *models.Cell
		err := json.NewDecoder(resp.Body).Decode(&cell)
		if err != nil {
			return fmt.Errorf("error decoding response: %w", err)
		}

		return nil
	}

	return fmt.Errorf("recieved non-200 status code: %d", resp.StatusCode)
}

func (c *Client) UpdateCell(cell *models.Cell) (*models.Cell, error) {
	req, err := c.NewRequest(http.MethodPut, "/workbook/"+cell.WorkbookID+"/sheet/"+cell.SheetID+"/cell/"+cell.ID, cell)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending requets: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var cell *models.Cell
		err := json.NewDecoder(resp.Body).Decode(&cell)
		if err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}

		return cell, nil
	}

	return nil, fmt.Errorf("recieved non-200 status code: %d", resp.StatusCode)
}
