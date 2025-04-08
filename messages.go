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

func (c *Client) GetMessagesFromDialog(ctx context.Context, chatKey, dialogID string) ([]*models.Message, error) {
	baseURL := fmt.Sprintf(c.APIBase+utils.MessagesEndPoint, chatKey)

	urlWithQuery := fmt.Sprintf("%s?chatId=%s", baseURL, url.QueryEscape(dialogID))

	req, err := http.NewRequestWithContext(ctx, "GET", urlWithQuery, nil)
	if err != nil {
		return nil, err
	}

	var messages []*models.Message
	if err := c.Send(req, &messages); err != nil {
		return nil, err
	}

	for i, msg := range messages {
		fmt.Printf("Message #%d: %+v\n", i+1, msg)
	}

	return messages, nil
}

func (c *Client) AddReactionToMessage(ctx context.Context, chatKey string, reaction models.Reaction) (bool, error) {
	url := fmt.Sprintf(c.APIBase+utils.ReactionEndPoint, chatKey)

	requestBodyBytes, err := json.Marshal(reaction)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return false, err
	}

	var response models.Response
	if err := c.Send(req, &response); err != nil {
		return false, err
	}

	return response.Success, nil
}

func (c *Client) EditMessage(ctx context.Context, chatKey string, message models.Message) (bool, error) {
	url := fmt.Sprintf(c.APIBase+utils.EditMessageEndPoint, chatKey)

	requestBodyBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("error: ", err)
		return false, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		fmt.Println("error: ", err)
		return false, err
	}

	var response models.Response
	if err := c.Send(req, &response); err != nil {
		return false, err
	}

	return response.Success, nil
}

func (c *Client) SendMessage(ctx context.Context, chatKey string, message models.MessageInput) (models.MessageResponse, error) {
	uri := fmt.Sprintf(c.APIBase+utils.SendMessageEndPoint, chatKey)

	requestBodyBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("error: ", err)
		return models.MessageResponse{}, err
	}

	req, err := http.NewRequest("POST", uri, bytes.NewReader(requestBodyBytes))
	if err != nil {
		fmt.Println("error: ", err)
		return models.MessageResponse{}, err
	}

	var msg models.MessageResponse
	if err := c.Send(req, &msg); err != nil {
		return models.MessageResponse{}, err
	}

	return msg, nil
}

func (c *Client) SendFile(ctx context.Context, chatKey string, message models.MessageInput) (models.MessageResponse, error) {
	uri := fmt.Sprintf(c.APIBase+utils.SendMessageEndPoint, chatKey)

	requestBodyBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("error: ", err)
		return models.MessageResponse{}, err
	}

	req, err := http.NewRequest("POST", uri, bytes.NewReader(requestBodyBytes))
	if err != nil {
		fmt.Println("error: ", err)
		return models.MessageResponse{}, err
	}

	var msg models.MessageResponse
	if err := c.Send(req, &msg); err != nil {
		return models.MessageResponse{}, err
	}

	return msg, nil
}
