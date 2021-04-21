package xiaomipush_test

import (
	"testing"

	"github.com/modood/pushapi/xiaomipush"
)

var appSecret = "your app secret"
var regId = "your reg id"
var channelId = "your channel id"
var channelName = "your channel name"

func TestSend(t *testing.T) {
	client := xiaomipush.NewClient(appSecret)

	sendReq := &xiaomipush.SendReq{
		RegistrationId: regId,
		Title:          "test push title",
		Description:    "test push content",
		NotifyType:     2,
		Extra: &xiaomipush.Extra{
			NotifyEffect: "1",
			ChannelId:    channelId,
			ChannelName:  channelName,
		},
	}
	sendRes, err := client.Send(sendReq)
	t.Log(sendRes, err)
}
