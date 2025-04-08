package go_whatcrm

import (
	"context"
	"fmt"
	"github.com/whatcrm/go-whatcrm/models"
	"github.com/whatcrm/go-whatcrm/utils"
	"net/http"
)

func (c *Client) GetInstance(ctx context.Context, domain, chatKey string) (*models.Instance, error) {
	data, err := c.GetInstances(ctx, domain)
	if err != nil {
		return nil, err
	}

	if len(data) <= 0 {
		return nil, fmt.Errorf("connection is not found")
	}

	d := &models.Instance{}
	for i := range data {
		if data[i].ChatKey == chatKey {
			d = &data[i]
			break
		}

		if i == len(data)-1 {
			return nil, fmt.Errorf("connection is not found")
		}
	}

	return d, nil
}

func (c *Client) GetInstances(ctx context.Context, chatKey string) ([]models.Instance, error) {
	url := fmt.Sprintf(c.APIBase+utils.ConnectionEndPoint, chatKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var v []models.Instance
	if err := c.Send(req, &v); err != nil {
		return nil, err
	}

	return v, nil
}

func (c *Client) GetConnectionStatus(ctx context.Context, chatKey string) (string, error) {
	url := fmt.Sprintf(c.APIBase+utils.ConnectionStatusEndPoint, chatKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	var status models.Status
	if err := c.Send(req, &status); err != nil {
		return "", err
	}

	return status.State, nil
}
