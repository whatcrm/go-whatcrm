package go_whatcrm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/whatcrm/go-whatcrm/utils"
)

type Client struct {
	Client  *http.Client
	APIBase string
	Token   string
}

func NewClient() (*Client, error) {
	return &Client{
		Client:  &http.Client{},
		APIBase: utils.BaseURL,
	}, nil
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	if req.Header.Get("X-Whatsapp-Token") == "" {
		return errors.New("request cannot be processed without header 'X-Whatsapp-Token'")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	var (
		err  error
		resp *http.Response
		data []byte
	)

	resp, err = c.Client.Do(req)
	if err != nil {
		fmt.Println("Do client error: ", err)
		return err
	}

	defer func(Body io.ReadCloser) error {
		return Body.Close()
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(data))
	}

	fmt.Println(string(data))

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err := io.Copy(w, resp.Body)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func buildQueryParamsString(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}
	var queryString strings.Builder
	queryString.WriteString("?")
	for key, value := range params {
		queryString.WriteString(fmt.Sprintf("%s=%v&", key, value))
	}
	query := queryString.String()
	return query[:len(query)-1] // remove the trailing '&'
}
