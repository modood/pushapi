# pushapi

[![Build Status](https://travis-ci.org/modood/pushapi.png)](https://travis-ci.org/modood/pushapi)
[![Coverage Status](https://coveralls.io/repos/github/modood/pushapi/badge.svg?branch=master)](https://coveralls.io/github/modood/pushapi?branch=master)
[![GoDoc](https://pkg.go.dev/badge/github.com/modood/pushapi)](https://pkg.go.dev/github.com/modood/pushapi)

各手机厂商推送 api 接入

*   vivo （最后更新时间：2021-02-08 17:34:37）
*   oppo （最后更新时间：2021-04-21 16:03:09）
*   小米 （最后更新时间：2022-08-02 10:00:00）
*   华为 （最后更新时间：2021-04-21 10:51:00）

## 调用示例

### vivo

```go
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/modood/pushapi/vivopush"
)

var appId = "your app id"
var appKey = "your app key"
var appSecret = "your app secret"
var regId = "your reg id"

func main() {
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
	fmt.Println(sendRes, err)
}
```

### oppo

```go
package main

import (
	"fmt"

	"github.com/modood/pushapi/oppopush"
)

var appKey = "your app key"
var masterSecret = "your master secret"
var regId = "your reg id"
var channelId = "your channel id"

func main() {
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
	fmt.Println(sendRes, err)
}
```

### 小米

```go
package main

import (
	"fmt"

	"github.com/modood/pushapi/xiaomipush"
)

var appSecret = "your app secret"
var regId = "your reg id"
var channelId = "your channel id"
var channelName = "your channel name"

func main() {
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
	fmt.Println(sendRes, err)
}
```

### 华为

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/modood/pushapi/huaweipush"
)

var appId = "your app id"
var appSecret = "your app secret"
var regId = "your reg id"
var badgeClass = "your badge class. example: com.example.hmstest.MainActivity"

func main() {
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
	fmt.Println(sendRes, err)
}
```

## License

this repo is released under the [MIT License](https://github.com/modood/pushapi/blob/master/LICENSE).

