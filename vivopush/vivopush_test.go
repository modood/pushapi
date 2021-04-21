package vivopush_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/modood/pushapi/vivopush"
)

var appId = "your app id"
var appKey = "your app key"
var appSecret = "your app secret"
var regId = "your reg id"

func TestSend(t *testing.T) {
	client := vivopush.NewClient(appId, appKey, appSecret)

	sendReq := &vivopush.SendReq{
		RegId:          regId,
		NotifyType:     4,
		Title:          "test push title",
		Content:        "test push content",
		TimeToLive:     24 * 60 * 60,
		SkipType:       1,
		NetworkType:    -1,
		Classification: 1,
		RequestId:      strconv.Itoa(int(time.Now().UnixNano())),
	}
	sendRes, err := client.Send(sendReq)
	t.Log(sendRes, err)
}
