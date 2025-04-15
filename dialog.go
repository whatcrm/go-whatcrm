package go_whatcrm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/whatcrm/go-whatcrm/models"
	"github.com/whatcrm/go-whatcrm/utils"
	"net/http"
	"net/url"
)

func (c *Client) GetDialogs(ctx context.Context, chatKey string) ([]models.Dialog, error) {
	uri := fmt.Sprintf(c.APIBase+utils.DialogsEndPoint, chatKey)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	var dialog []models.Dialog
	if err := c.Send(req, &dialog); err != nil {
		return nil, err
	}
	return dialog, nil
}

func (c *Client) GetDialogInfo(ctx context.Context, chatKey, d models.Dialog) (models.Dialog, error) {
	uri := fmt.Sprintf(c.APIBase+utils.DialogEndPoint, chatKey)

	requestBodyBytes, err := json.Marshal(d.ChatID)
	if err != nil {
		fmt.Println("error: ", err)
		return models.Dialog{}, nil
	}

	req, err := http.NewRequest("GET", uri, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return models.Dialog{}, err
	}

	var dialog models.Dialog
	if err := c.Send(req, &dialog); err != nil {
		return models.Dialog{}, err
	}
	return dialog, nil
}

func (c *Client) ReadDialog(ctx context.Context, chatKey, dialogID string) (interface{}, error) {
	uri := fmt.Sprintf(c.APIBase+utils.ReadDialogEndPoint, chatKey)

	urlWithQuery := fmt.Sprintf("%s?chatId=%s", uri, url.QueryEscape(dialogID))

	req, err := http.NewRequest("GET", urlWithQuery, nil)
	if err != nil {
		return models.Message{}, err
	}

	var response models.Response
	if err := c.Send(req, &response); err != nil {
		return models.Message{}, err
	}
	return response, nil
}

func (c *Client) PinDialog(ctx context.Context, chatKey, dialogID string) (bool, error) {
	uri := fmt.Sprintf(c.APIBase+utils.PinChatEndPoint, chatKey)

	requestBodyBytes, err := json.Marshal(dialogID)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("GET", uri, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return false, err
	}

	var response models.Response
	if err := c.Send(req, &response); err != nil {
		return false, err
	}

	return response.Success, nil
}

func (c *Client) UnpinDialog(ctx context.Context, chatKey, dialogID string) (bool, error) {
	uri := fmt.Sprintf(c.APIBase+utils.UnpinChatEndPoint, chatKey)

	requestBodyBytes, err := json.Marshal(dialogID)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("GET", uri, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return false, err
	}

	var response models.Response
	if err := c.Send(req, &response); err != nil {
		return false, err
	}

	return response.Success, nil
}

func (c *Client) DeleteDialog(ctx context.Context, chatKey string, d models.Dialog) (bool, error) {
	uri := fmt.Sprintf(c.APIBase+utils.DeleteDialogEndPoint, chatKey)

	requestBodyBytes, err := json.Marshal(d.ChatID)
	if err != nil {
		fmt.Println("error: ", err)
		return false, nil
	}

	req, err := http.NewRequest("POST", uri, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return false, err
	}

	var response models.Response
	if err := c.Send(req, &response); err != nil {
		return false, err
	}

	return response.Success, nil
}
