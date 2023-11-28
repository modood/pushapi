package huaweipush

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/modood/pushapi/httputil"
)

type Client struct {
	httpClient        *http.Client
	appId             string
	appSecret         string
	authToken         string
	authTokenExpireAt int64
	sendURL           string
}

func NewClient(appId, appSecret string) *Client {
	return &Client{
		appId:     appId,
		appSecret: appSecret,
		sendURL:   fmt.Sprintf(Host+SendURL, appId),
	}
}

func (c *Client) SetHost(host string) {
	c.sendURL = fmt.Sprintf(host+SendURL, c.appId)
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

func (c *Client) auth(ctx context.Context) (string, error) {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	if c.authToken != "" && c.authTokenExpireAt > now {
		return c.authToken, nil
	}

	req := &AuthReq{
		GrantType:    "client_credentials",
		ClientId:     c.appId,
		ClientSecret: c.appSecret,
	}
	res := &AuthRes{}

	params := httputil.StructToUrlValues(req)
	code, resBody, err := httputil.PostForm(ctx, c.httpClient, AuthURL, params, res, nil)
	if err != nil {
		return "", fmt.Errorf("code=%d body=%s err=%v", code, resBody, err)
	}

	if code != http.StatusOK || res.AccessToken == "" {
		return "", fmt.Errorf("code=%d body=%s", code, resBody)
	}

	c.authToken = fmt.Sprintf("%s %s", res.TokenType, res.AccessToken)
	c.authTokenExpireAt = now + 60*1000 // 一分钟后更新
	return c.authToken, nil
}

func (c *Client) Send(req *SendReq) (*SendRes, error) {
	return c.SendWithContext(context.Background(), req)
}

func (c *Client) SendWithContext(ctx context.Context, req *SendReq) (*SendRes, error) {
	res := &SendRes{}

	token, err := c.auth(ctx)
	if err != nil {
		return nil, err
	}

	code, resBody, err := httputil.PostJSON(ctx, c.httpClient, c.sendURL, req, res, map[string]string{"Authorization": token})
	if err != nil {
		return nil, fmt.Errorf("code=%d body=%s err=%v", code, resBody, err)
	}

	if code != http.StatusOK || res.Code != CodeSuccess {
		return nil, fmt.Errorf("code=%d body=%s", code, resBody)
	}

	return res, nil
}
