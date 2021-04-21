package xiaomipush

import (
	"fmt"
	"net/http"

	"github.com/modood/pushapi/httputil"
)

type Client struct {
	host      string
	appSecret string
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

func (c *Client) Send(req *SendReq) (*SendRes, error) {
	res := &SendRes{}

	params := httputil.StructToUrlValues(req)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded;charset=UTF-8",
		"Authorization": fmt.Sprintf("key=%s", c.appSecret),
	}

	code, resBody, err := httputil.PostForm(c.host+SendURL, params, res, headers)
	if err != nil {
		return nil, fmt.Errorf("code=%d body=%s err=%v", code, resBody, err)
	}

	if code != http.StatusOK || res.Code != 0 {
		return nil, fmt.Errorf("code=%d body=%s", code, resBody)
	}

	return res, nil
}
