package xiaomipush

import (
	"context"
	"fmt"
	"net/http"

	"github.com/modood/pushapi/httputil"
)

type Client struct {
	httpClient *http.Client
	host       string
	appSecret  string
}

func NewClient(appSecret string) *Client {
	return &Client{
		host:      Host,
		appSecret: appSecret,
	}
}

func (c *Client) SetHost(host string) {
	c.host = host
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

func (c *Client) Send(req *SendReq) (*SendRes, error) {
	return c.SendWithContext(context.Background(), req)
}

func (c *Client) SendWithContext(ctx context.Context, req *SendReq) (*SendRes, error) {
	res := &SendRes{}

	params := httputil.StructToUrlValues(req)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded;charset=UTF-8",
		"Authorization": fmt.Sprintf("key=%s", c.appSecret),
	}

	code, resBody, err := httputil.PostForm(ctx, c.httpClient, c.host+SendURL, params, res, headers)
	if err != nil {
		return nil, fmt.Errorf("code=%d body=%s err=%v", code, resBody, err)
	}

	if code != http.StatusOK || res.Code != 0 {
		return nil, fmt.Errorf("code=%d body=%s", code, resBody)
	}

	return res, nil
}
