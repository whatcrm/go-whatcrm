package go_whatcrm

import (
	"context"
	"github.com/whatcrm/go-whatcrm/models"
	"net/http"
)

func (c *Client) GetDialogs(ctx context.Context, chatID string) ([]models.Dialog, error) {
	url := c.APIBase + utils.DialogsEndPoint + chatID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var dialogResponse models.DialogResponse
	if err := c.Send(req, dialogResponse); err != nil {
		return nil, err
	}
	return dialogResponse.Data, nil
}
