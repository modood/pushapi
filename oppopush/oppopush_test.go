package oppopush_test

import (
	"testing"

	"github.com/modood/pushapi/oppopush"
)

var appKey = "your app key"
var masterSecret = "your master secret"
var regId = "your reg id"
var channelId = "your channel id"

func TestSend(t *testing.T) {
	client := oppopush.NewClient(appKey, masterSecret)

	sendReq := &oppopush.SendReq{
		Notification: &oppopush.Notification{
			Title:     "test push title",
			Content:   "test push content",
			ChannelID: channelId,
		},
		TargetType:  2,
		TargetValue: regId,
	}
	sendRes, err := client.Send(sendReq)
	t.Log(sendRes, err)
}
