package huaweipush_test

import (
	"strconv"
	"testing"

	"github.com/modood/pushapi/huaweipush"
)

var appId = "your app id"
var appSecret = "your app secret"
var regId = "your reg id"
var badgeClass = "your badge class. example: com.example.hmstest.MainActivity"

func TestSend(t *testing.T) {
	client := huaweipush.NewClient(appId, appSecret)

	sendReq := &huaweipush.SendReq{
		Message: &huaweipush.Message{
			Android: &huaweipush.AndroidConfig{
				FastAppTarget: 2,
				Notification: &huaweipush.AndroidNotification{
					Title: "test push title",
					Body:  "test push content",
					ClickAction: &huaweipush.ClickAction{
						Type: 3,
					},
					Sound: strconv.Itoa(1),
					Badge: &huaweipush.BadgeNotification{
						AddNum: 1,
						Class:  badgeClass,
					},
				},
			},
			Tokens: []string{regId},
		},
	}
	sendRes, err := client.Send(sendReq)
	t.Log(sendRes, err)
}
