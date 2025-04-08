package go_whatcrm

import (
	"context"
	"fmt"
	"github.com/whatcrm/go-whatcrm/models"
	"github.com/whatcrm/go-whatcrm/utils"
	"net/http"
)

func (c *Client) GetContactInfo(ctx context.Context, chatKey, userID string) (*models.User, error) {
	baseURL := fmt.Sprintf(c.APIBase, chatKey)
	url := baseURL + fmt.Sprintf(utils.UserEndPoint, userID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := c.Send(req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
