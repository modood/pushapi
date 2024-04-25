package huaweipush

const (
	Host = "https://push-api.cloud.huawei.com"

	AuthURL = "https://oauth-login.cloud.huawei.com/oauth2/v2/token" // 鉴权
	SendURL = "/v1/%s/messages:send"                                 // 下行消息
)

type Code = string

const (
	CodeSuccess              Code = "80000000" // 成功。
	CodeIllegalToken         Code = "80100000" // 部分Token发送成功，返回的illegal_tokens为不合法而发送失败的Token。请检查返回值中发送失败的Token。
	CodeNotCorrectToken      Code = "80100001" // 请检查返回值中发送失败的Token。按照响应消息中的提示，请检查请求参数。
	CodeSyncCountToken       Code = "80100002" // 发送同步消息的时候，Token的数量必须为1。请检查请求参数中Token字段。
	CodeIncorrectMessage     Code = "80100003" // 消息结构体错误。按照响应消息中的提示，请检查消息结构体的参数。
	CodeTTL                  Code = "80100004" // 消息设置的过期时间小于当前时间导致。请检查消息字段ttl。
	CodeColapseKey           Code = "80100013" // 消息字段collapse_key不合法。请检查消息字段collapse_key。
	CodeSensitiveInformation Code = "80100016" // 消息里面含有敏感信息。请检查发送消息内容。
	CodeTooManyTopics        Code = "80100017" // 同时发送的Topic任务超过100个。请稍后再发送主题消息，增加主题消息发送间隔。
	CodeReviewFailed         Code = "80100018" // 消息体内容验签不通过。请检查发给三方机构审核的消息体与发给Push服务器的消息体内容是否一致。
	CodeOAuth                Code = "80200001" // Oauth认证错误。请求HTTP头中Authorization参数里面的Access Token鉴权失败，请检查Access Token。
	CodeOAuthExpired         Code = "80200003" // Oauth Token过期。请求HTTP头中Authorization参数里面的Access Token已过期，请重新申请后重试。
	CodeAppPermission        Code = "80300002" // 当前应用无权限下发推送消息。
	CodeInvalidTokens        Code = "80300007" // 所有Token都是无效的。
	CodeMessageSize          Code = "80300008" // 消息体大小超过系统设置的默认值（4096Bytes）。请求消息体大小超过默认值，请减小消息体后重新发送消息。
	CodeNumberTokens         Code = "80300010" // 消息体中的Token数量超过系统设置的默认值。请减少Token数量后分批发送消息。
	CodePriority             Code = "80300011" // 无权限发送高级别通知消息。请申请权限。
	CodeReceiptFailed        Code = "80300013" // 回执地址错误。请检查您的回执地址是否正确，回执证书是否过期。
	CodeOAuthFailed          Code = "80600003" // 请求Oauth服务失败。请检查Oauth 2.0客户端ID和客户端密钥。
	CodeInternal             Code = "81000001" // 系统内部错误。请联系华为技术支持解决。
)

// App Services > 推送服务 > 指南 基于OAuth 2.0开放鉴权
// https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/oauth2-0000001212610981#section128682386159
type AuthReq struct {
	GrantType    string `json:"grant_type,omitempty"`    // 填写为“client_credentials”，表示为客户端模式。
	ClientId     string `json:"client_id,omitempty"`     // 在接入前准备中得到的OAuth 2.0客户端ID，对于AppGallery Connect类应用，该值为应用的APP ID。
	ClientSecret string `json:"client_secret,omitempty"` // 在接入前准备中给客户端ID分配的密钥，对于AppGallery Connect类应用，该值为应用的APP SECRET。
}

type AuthRes struct {
	AccessToken string `json:"access_token"` // 应用级Access Token。
	ExpiresIn   int    `json:"expires_in"`   // Access Token的剩余有效期，单位：秒。3600 秒
	TokenType   string `json:"token_type"`   // 固定返回Bearer，标识返回Access Token的类型。
}

// App Services > 推送服务 > API参考 下行消息
// https://developer.huawei.com/consumer/cn/doc/HMSCore-References/https-send-api-0000001050986197
type SendReq struct {
	ValidateOnly bool      `json:"validate_only,omitempty"` // 控制当前是否为测试消息，测试消息只做格式合法性校验，不会推送给用户设备，取值如下：true：测试消息 false：正式消息（默认值）
	Message      *Message  `json:"message,omitempty"`       // 推送消息结构体，message结构体中必须存在有效消息负载以及有效发送目标，具体字段请参见Message的定义。
	Review       []*Review `json:"review,omitempty"`        // 第三方审核结构对推送消息体内容的审核结果信息，具体结构请参见Review的定义。 https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/android-3rd-party-review-0000001050166008
}

type SendRes struct {
	Code      Code   `json:"code"`      // 错误码。
	Message   string `json:"msg"`       // 错误码描述。
	RequestID string `json:"requestId"` // 请求标识。
}

type Message struct {
	Data         string         `json:"data,omitempty"`         // 自定义消息负载，通知栏消息支持JSON格式字符串，透传消息支持普通字符串或者JSON格式字符串。样例："your data"，"{'param1':'value1','param2':'value2'}"。消息体中有message.data，没有message.notification和message.android.notification，消息类型为透传消息如果用户发送的是网页应用的透传消息，那么接收消息中字段orignData为透传消息内容
	Notification *Notification  `json:"notification,omitempty"` // 通知栏消息内容，具体字段请参见Notification的定义。
	Android      *AndroidConfig `json:"android,omitempty"`      // Android消息推送控制参数，具体字段请参见AndroidConfig的定义。如果是Android通知栏消息，本字段必填。
	Apns         *ApnsConfig    `json:"apns,omitempty"`         // iOS消息推送控制参数，具体字段请参见ApnsConfig的定义。 如果是iOS消息，本字段必填。
	Webpush      *WebPushConfig `json:"webpush,omitempty"`      // 网页应用推送消息控制参数，具体字段请参见WebPushConfig结构体的定义。 如果是网页应用通知栏消息，本字段必填。
	Tokens       []string       `json:"token,omitempty"`        // 按照Token向目标用户推消息，token/topic/condition三者只能且必须设置一个。样例：["pushtoken1","pushtoken2"]
	Topic        string         `json:"topic,omitempty"`        // 按照Topic向订阅了本topic的用户推消息（目前只支持Android应用），token/topic/condition三者只能且必须设置一个。
	Condition    string         `json:"condition,omitempty"`    // 按照条件（主题组合表达式）向目标用户推消息（目前只支持Android应用），token/topic/condition三者只能且必须设置一个。
}

type Review struct {
	Reviewer string                 `json:"reviewer"` // 第三方审核机构名称，当前必须设置为tuibian。
	Type     int                    `json:"type"`     // 消息体内容经审核后的类型标识，当前必须设置为0。
	Result   map[string]interface{} `json:"result"`   // 第三方审核机构对消息体内容的审核结果，具体结构请参见推必安公有云接口文档。https://tuibian.mobileservice.cn/
}

type Notification struct {
	Title    string `json:"title,omitempty"` // 通知栏消息的标题。
	Body     string `json:"body,omitempty"`  // 通知栏消息的内容。
	ImageURL string `json:"image,omitempty"` // 用户自定义的通知栏消息右侧大图标URL，如果不设置，则不展示通知栏右侧图标。URL使用的协议必须是HTTPS协议，取值样例：https://example.com/image.png。
}

type AndroidConfig struct {
	BiTag string `json:"bi_tag,omitempty"` // 批量任务消息标识，消息回执时会返回给应用服务器，应用服务器可以识别bi_tag对消息的下发情况进行统计分析。

	// 作用一：完成自分类权益申请后，用于标识消息类型，确定消息提醒方式，对特定类型消息加快发送，取值如下：
	// https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001652651372__section893184112272
	// https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/message-classification-0000001149358835#ZH-CN_TOPIC_0000001652651372__p3850133955718
	// *   IM：即时聊天
	// *   VOIP：音视频通话
	// *   SUBSCRIPTION：订阅
	// *   TRAVEL：出行
	// *   HEALTH：健康
	// *   WORK：工作事项提醒
	// *   ACCOUNT：帐号动态
	// *   EXPRESS：订单&物流
	// *   FINANCE：财务
	// *   DEVICE_REMINDER：设备提醒
	// *   MAIL：邮件
	// *   PLAY_VOICE：语音播报（仅透传消息支持）
	// *   MARKETING：内容推荐、新闻、财经动态、生活资讯、社交动态、调研、产品促销、功能推荐、运营活动（仅对内容进行标识，不会加快消息发送）
	//
	// 作用二：申请特殊权限后，用于标识高优先级透传场景，取值如下：
	// https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/faq-0000001050042183#section037425218509
	// *   VOIP：音视频通话
	// *   PLAY_VOICE：语音播报
	Category string `json:"category,omitempty"`

	// 用户设备离线时，Push服务器对离线消息缓存机制的控制方式，用户设备上线后缓存消息会再次下发，取值如下：
	// *   0：对每个应用发送到该用户设备的离线消息只会缓存最新的一条
	// *   -1：对所有离线消息都缓存（默认值）
	// *   1~100：离线消息缓存分组标识，对离线消息进行分组缓存，每个应用每一组最多缓存一条离线消息
	// 如果您发送了10条消息，其中前5条的collapse_key为1，后5条的collapse_key为2，那么待用户上线后collapse_key为1和2的分别下发最新的一条消息给最终用户。
	CollapseKey    int                  `json:"collapse_key,omitempty"`
	Data           string               `json:"data,omitempty"`             // 自定义消息负载，此处如果设置了data，则会覆盖message.data字段。
	FastAppTarget  int                  `json:"fast_app_target,omitempty"`  // 快应用发送透传消息时，指定小程序的模式类型，小程序有两种模式开发态和生产态，取值如下：1：开发态 2：生产态（默认值）
	Notification   *AndroidNotification `json:"notification,omitempty"`     // Android通知栏消息结构体，具体字段请参见AndroidNotification结构体的定义。
	ReceiptId      string               `json:"receipt_id,omitempty"`       // 输入一个唯一的回执ID指定本次下行消息的回执地址及配置，该回执ID可以在回执参数配置中查看。 https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/msg-receipt-guide-0000001050040176#ZH-CN_TOPIC_0000001700731529__li15263162510251
	TargetUserType int                  `json:"target_user_type,omitempty"` // 0：普通消息（默认值） 1：测试消息。每个应用每日可发送该测试消息500条且不受每日单设备推送数量上限要求 https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/message-restriction-description-0000001361648361#section104849311415
	TTL            string               `json:"ttl,omitempty"`              // 消息缓存时间，单位是秒。在用户设备没有网络时，消息在Push服务器进行缓存，在消息缓存时间内用户设备重新连接网络，消息会下发，超过缓存时间后消息会丢弃，默认值为“86400s”（1天），最大值为“1296000s”（15天）。
	Urgency        string               `json:"urgency,omitempty"`          // 透传消息投递优先级，取值如下：HIGH NORMAL（默认值） 设置为HIGH时需要申请权限，请参见申请特殊权限。 HIGH级别消息到达用户手机时可强制拉起应用进程。 https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/faq-0000001050042183#section037425218509
}

type AndroidNotification struct {
	Title             string                 `json:"title,omitempty"`               // Android通知栏消息标题，如果此处设置了title则会覆盖message.notification.title字段，且发送通知栏消息时，此处title和message.notification.title两者最少需要设置一个。
	Body              string                 `json:"body,omitempty"`                // Android通知栏消息内容，如果此处设置了body则会覆盖message.notification.body字段，且发送通知栏消息时，此处body和message.notification.body两者最少需要设置一个。
	Icon              string                 `json:"icon,omitempty"`                // 自定义通知栏消息左侧小图标，此处设置的图标文件必须存放在应用的/res/raw路径下，例如“/raw/ic_launcher”，对应应用本地的“/res/raw/ic_launcher.xxx”文件。支持的文件格式目前包括PNG、JPG。自定义小图标规格规范请参见通知图标规范。 https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/notificattion_spec-0000001052845223#section1626375315119
	Color             string                 `json:"color,omitempty"`               // 自定义通知栏按钮颜色，以#RRGGBB格式，其中RR代表红色的16进制色素，GG代表绿色的16进制色素，BB代表蓝色的16进制色素，样例：#FFEEFF。
	Sound             string                 `json:"sound,omitempty"`               // 自定义消息通知铃声。在新创建渠道时有效，此处设置的铃声文件必须存放在应用的/res/raw路径下，例如设置为“/raw/shake”，对应应用本地的“/res/raw/shake.xxx”文件。支持的文件格式包括MP3、WAV、MPEG等，如果不设置，则用默认系统铃声。
	DefaultSound      bool                   `json:"default_sound,omitempty"`       // 默认铃声控制开关，取值如下：true：使用系统默认铃声（默认值） false：使用sound自定义铃声
	Tag               string                 `json:"tag,omitempty"`                 // 消息标签，同一应用下使用同一个消息标签的消息会相互覆盖，只展示最新的一条。
	ClickAction       *ClickAction           `json:"click_action,omitempty"`        // 消息点击行为，具体字段请参见ClickAction结构体的定义。 如果是Android通知栏消息时，则该参数必选。
	BodyLocKey        string                 `json:"body_loc_key,omitempty"`        // 显示本地化body的StringId，具体使用请参见通知栏消息语言本地化。 https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/android-noti-local-0000001050042073
	BodyLocArgs       []string               `json:"body_loc_args,omitempty"`       // 本地化body的可变参数，具体使用请参见通知栏消息语言本地化。样例："body_loc_args":["1","2","3"]
	TitleLocKey       string                 `json:"title_loc_key,omitempty"`       // 显示本地化title的StringId，具体使用请参见通知栏消息语言本地化。
	TitleLocArgs      []string               `json:"title_loc_args,omitempty"`      // 本地化title的可变参数，具体使用请参见通知栏消息语言本地化。样例："title_loc_args":["1","2","3"]
	MultiLangKey      map[string]interface{} `json:"multi_lang_key,omitempty"`      // 消息国际化多语言参数，body_loc_key，title_loc_key优先从multi_lang_key读取内容，如果key不存在，则从APK本地字符串资源读，具体使用请参见通知栏消息语言本地化。最多设置3种语言。
	ChannelID         string                 `json:"channel_id,omitempty"`          // 自Android O版本后可以支持通知栏自定义渠道，指定消息要展示在哪个通知渠道上，详情请参见自定义通知渠道。 https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/android-custom-chan-0000001050040122
	NotifySummary     string                 `json:"notify_summary,omitempty"`      // Android通知栏消息简要描述。
	Image             string                 `json:"image,omitempty"`               // 自定义通知栏消息右侧小图片URL，功能和message.notification.image字段一样，如果此处设置，则覆盖message.notification.image中的值。URL使用的协议必须是HTTPS协议，取值样例：https://example.com/image.png。
	Style             int                    `json:"style,omitempty"`               // 通知栏样式，取值如下： 0：默认样式 1：大文本样式 3：Inbox样式
	BigTitle          string                 `json:"big_title,omitempty"`           // Android通知栏消息大文本标题，当style为1时必选，设置big_title后通知栏展示时，使用big_title而不用title。
	BigBody           string                 `json:"big_body,omitempty"`            // Android通知栏消息大文本内容，当style为1时必选，设置big_body后通知栏展示时，使用big_body而不用body。
	AutoClear         int                    `json:"auto_clear,omitempty"`          // 消息展示时长，超过后自动清除，单位：毫秒。
	NotifyID          int                    `json:"notify_id,omitempty"`           // 每条消息在通知显示时的唯一标识。不携带时或者设置-1时，Push NC自动为给每条消息生成一个唯一标识；不同的通知栏消息可以拥有相同的notifyId，实现新的消息覆盖上一条消息功能。
	Group             string                 `json:"group,omitempty"`               // 消息分组，例如发送10条带有同样group字段的消息，手机上只会展示该组消息中最新的一条和当前该组接收到的消息总数目，不会展示10条消息。
	Badge             *BadgeNotification     `json:"badge,omitempty"`               // Android通知消息角标控制，具体字段请参见BadgeNotification结构体的定义。
	Ticker            string                 `json:"ticker,omitempty"`              // 设备收到通知消息后状态栏上显示的内容提示。受Android系统原生机制的限制，在Android 5.0版本（API Level 21）之后的设备上，设置了该字段也不会显示。
	AutoCancel        bool                   `json:"auto_cancel,omitempty"`         // 通知消息常驻标识，用户点击通知中心消息后，消息是否仍驻留在通知中心上，取值如下： true：设置为true，用户点击消息后从通知中心清理掉。 false：用户点击消息后消息仍常驻通知中心，需要开通权益，请参见申请特殊权限。
	When              string                 `json:"when,omitempty"`                // 消息的排序时间，Android通知栏消息根据这个值将消息排序，同时将转换后的时间在通知栏上显示。样例：2014-10-02T15:01:23.045123456Z
	Importance        string                 `json:"importance,omitempty"`          // Android通知栏消息优先级，决定用户设备消息通知行为，取值如下： LOW：一般（静默）消息 NORMAL：重要消息 HIGH：非常重要消息
	UseDefaultVibrate bool                   `json:"use_default_vibrate,omitempty"` // 是否使用系统默认振动模式控制开关。
	UseDefaultLight   bool                   `json:"use_default_light,omitempty"`   // 是否使用默认呼吸灯模式控制开关。
	VibrateConfig     []string               `json:"vibrate_config,omitempty"`      // Android自定义通知消息振动模式，每个数组元素按照“[0-9]+|[0-9]+[sS]|[0-9]+[.][0-9]{1,9}|[0-9]+[.][0-9]{1,9}[sS]”格式，取值样例["3.5S","2S","1S","1.5S"]，数组元素最多支持10个，每个元素数值整数大于0小于等于60。暂不支持EMUI 11。
	Visibility        string                 `json:"visibility,omitempty"`          // Android通知栏消息可见性，取值如下： “VISIBILITY_UNSPECIFIED”：未指定“visibility”，效果等同于设置了“PRIVATE”。 “PUBLIC”：锁屏时收到通知栏消息，显示消息内容。 “SECRET”：锁屏时收到通知栏消息，不提示收到通知消息。 “PRIVATE”：设置了锁屏密码，“锁屏通知”（导航：“设置 > 通知 > 隐藏通知内容”）选择“隐藏通知内容”时收到通知消息，不显示消息内容。
	LightSettings     *LightSettings         `json:"light_settings,omitempty"`      // 自定义呼吸灯模式，具体字段请参见LightSettings结构体的定义。
	ForegroundShow    bool                   `json:"foreground_show,omitempty"`     // 应用在前台时通知栏消息是否前台展示开关，具体使用请参见基于前台应用的通知展示。
	ProfileId         string                 `json:"profile_id,omitempty"`          // 关联终端设备登录用户标识，最大长度为64。
	InboxContent      []string               `json:"inbox_content,omitempty"`       // 当style为3时，Inbox样式的内容（必选），支持最大5条内容，每条最大长度1024。展示效果请参见Inbox样式。
	Buttons           []Button               `json:"buttons,omitempty"`             // 通知栏消息动作按钮，最多设置3个。具体字段请参见Button结构体的定义。
}

type Button struct {
	Name       string `json:"name,omitempty"`        // 按钮名称，最大长度40。
	ActionType int    `json:"action_type,omitempty"` // 按钮动作类型： 0：打开应用首页 1：打开应用自定义页面 2：打开指定的网页 3：清除通知 4：华为分享功能
	IntentType int    `json:"intent_type,omitempty"` //	打开自定义页面的方式： 0：设置通过intent打开应用自定义页面 1：设置通过action打开应用自定义页面 当action_type为1时，该字段必填。
	Intent     string `json:"intent,omitempty"`      // 当action_type为1，此字段按照intent_type字段设置应用页面的uri或者action，具体设置方式参见打开应用自定义页面。当action_type为2，此字段设置打开指定网页的URL，URL使用的协议必须是HTTPS协议，取值样例：https://example.com/image.png。
	Data       string `json:"data,omitempty"`        // 最大长度1024。 当字段action_type为0或1时，该字段用于在点击按钮后给应用透传数据，选填，格式必须为key-value形式：{"key1":"value1","key2":"value2",…}。 当action_type为4时，此字段必选，为分享的内容。
}

type ClickAction struct {
	Type   int    `json:"type,omitempty"`   // 消息点击行为类型，取值如下： 1：打开应用自定义页面 2：点击后打开特定URL 3：点击后打开应用
	Intent string `json:"intent,omitempty"` // 自定义页面中intent的实现，请参见指定intent参数​。 当type为1时，字段intent和action至少二选一。
	URL    string `json:"url,omitempty"`    // 设置打开特定URL，本字段填写需要打开的URL，URL使用的协议必须是HTTPS协议，取值样例：https://example.com/image.png。 当type为2时必选。 如果是游戏类应用，不支持设置特定URL。
	Action string `json:"action,omitempty"` // 设置通过action打开应用自定义页面时，本字段填写要打开的页面activity对应的action。 当type为1（打开自定义页面）时，字段intent和action至少二选一。
}

type BadgeNotification struct {
	AddNum int    `json:"add_num,omitempty"` // 应用角标累加数字非应用角标实际显示数字，为大于0小于100的整数。 例如，某应用当前有N条未读消息，若add_num设置为3，则每发一次消息，应用角标显示的数字累加3，为N+3。
	Class  string `json:"class,omitempty"`   // 应用入口Activity类全路径。 样例：com.example.hmstest.MainActivity
	SetNum int    `json:"set_num,omitempty"` // 角标设置数字，大于等于0小于100的整数。 例如，set_num设置为10，则不论发了多少次消息，应用角标显示的数字都是10。 如果set_num与add_num同时存在时，以set_num为准。
}

type LightSettings struct {
	Color            Color  `json:"color,omitempty"`              // 呼吸灯颜色，当设置light_settings时，该字段必选。具体字段请参见Color结构体的定义。
	LightOnDuration  string `json:"light_on_duration,omitempty"`  // 呼吸灯点亮时间间隔，当设置light_settings时，该字段必选，格式按照“\d+|\d+[sS]|\d+.\d{1,9}|\d+.\d{1,9}[sS]”。
	LightOffDuration string `json:"light_off_duration,omitempty"` // 呼吸灯熄灭时间间隔，当设置light_settings时，该字段必选，格式按照“\d+|\d+[sS]|\d+.\d{1,9}|\d+.\d{1,9}[sS]”。
}

type Color struct {
	Alpha float64 `json:"alpha,omitempty"` // RGB颜色中的alpha设置，默认值为1，取值范围[0,1]。
	Red   float64 `json:"red,omitempty"`   // RGB颜色中的red设置，默认值为0，取值范围[0,1]。
	Green float64 `json:"green,omitempty"` // RGB颜色中的green设置，默认值为0，取值范围[0,1]。
	Blue  float64 `json:"blue,omitempty"`  // RGB颜色中的blue设置，默认值为0，取值范围[0,1]。
}

type ApnsConfig struct {
	Headers    map[string]string     `json:"headers,omitempty"`     // APNs消息头。具体字段请参见iOS开发者网站。
	Payload    map[string]string     `json:"payload,omitempty"`     // APNs消息负载。如果消息负载中设置了title、body则会覆盖message.notification.title、body字段，且发送消息时，此处title、body和message.notification.title、body两者最少需要设置一个。具体字段请参见iOS开发者网站。
	HMSOptions *ApnsConfigHmsOptions `json:"hms_options,omitempty"` // APNs的hms参数，具体字段请参见ApnsConfig.HmsOptions结构体的定义。
}

type ApnsConfigHmsOptions struct {
	TargetUserType int `json:"target_user_type,omitempty"` // 目标用户类型，取值如下： 1：测试用户 2：正式用户 3：VoIP用户
}

type WebPushConfig struct {
	Headers      Headers                  `json:"headers,omitempty"`      // Web推送消息头，具体字段请参见Headers结构体的定义。
	Notification *WebNotification         `json:"notification,omitempty"` // Web推送通知栏消息结构体，具体字段请参见WebNotification结构体的定义。
	HmsOptions   *WebPushConfigHmsOptions `json:"hms_options,omitempty"`  // Web推送的参数，具体字段请参见WebPushConfig.HmsOptions结构体的定义。
}

type Headers struct {
	TTL     string `json:"ttl,omitempty"`     // 消息缓存时间，单位是秒，示例：20或者20s或者20S。
	Topic   string `json:"topic,omitempty"`   // 消息标识，可用于覆盖未送达的消息。
	Urgency string `json:"urgency,omitempty"` // 消息紧急程度，取值如下： very-low：不紧急 low：一般 normal：紧急 high：非常紧急
}

type WebNotification struct {
	Title              string        `json:"title,omitempty"`              // 网页应用通知消息标题，如果此处设置了title则会覆盖message.notification.title字段，且发送消息时，此处title和message.notification.title两者最少需要设置一个。
	Body               string        `json:"body,omitempty"`               // 网页应用通知消息文本，如果此处设置了body则会覆盖message.notification.body字段，且发送消息时，此处body和message.notification.body两者最少需要设置一个。
	Icon               string        `json:"icon,omitempty"`               // 小图标URL。
	Image              string        `json:"image,omitempty"`              // 大图URL。
	Lang               string        `json:"lang,omitempty"`               // 语言。
	Tag                string        `json:"tag,omitempty"`                // 通知消息分组覆盖标签，多条相同tag折叠显示，显示最新一条，仅仅用于手机端浏览器。
	Badge              string        `json:"badge,omitempty"`              // 浏览器图标URL，仅用于手机端浏览器，用于替换默认情况下显示的浏览器图标。
	Dir                string        `json:"dir,omitempty"`                // 文字方向，取值如下：auto：默认方向，从左向右。 ltr：方向从左向右。 rtl：方向从右向左。
	Vibrate            []int         `json:"vibrate,omitempty"`            // 振动间隔时间，单位毫秒。样例：[100,200,300]
	Renotify           bool          `json:"renotify,omitempty"`           // 消息重新提醒标识。
	RequireInteraction bool          `json:"requireInteraction,omitempty"` // 通知应保持活动状态，直到用户点击或将其关闭为止，而不是自动关闭。
	Silent             bool          `json:"silent,omitempty"`             // 消息免声音、振动提醒标识。
	TimestampMillis    *int64        `json:"timestamp,omitempty"`          // 标准的unix时间戳。
	Actions            []*WebActions `json:"actions,omitempty"`            // 消息动作定义，具体字段请参见WebActions结构体的定义。
}

type WebActions struct {
	Action string `json:"action,omitempty"` // 动作的名称。
	Icon   string `json:"icon,omitempty"`   // 动作的按钮图标URL。
	Title  string `json:"title,omitempty"`  // 动作显示的标题。
}

type WebPushConfigHmsOptions struct {
	Link string `json:"link,omitempty"` // 没有action情况下，点击跳转的默认URI。
}
