package oppopush

const (
	Host = "https://api.push.oppomobile.com"

	AuthURL = "/server/v1/auth"                         // 鉴权
	SendURL = "/server/v1/message/notification/unicast" // 单推-通知栏消息推送
)

// SendReq 单推-通知栏消息推送
// https://open.oppomobile.com/new/developmentDoc/info?id=11238
type SendReq struct {
	TargetType           int           `json:"target_type,omitempty"`  // 目标类型 2: registration_id  5:别名
	TargetValue          string        `json:"target_value,omitempty"` // 推送目标用户: registration_id或alias
	Notification         *Notification `json:"notification,omitempty"` // 请参见通知栏消息
	VerifyRegistrationId bool          `json:"verify_registration_id"` // 消息到达客户端后是否校验registration_id。 true表示推送目标与客户端registration_id进行比较，如果一致则继续展示，不一致则就丢弃；false表示不校验
}

// 问：OPPO Push 推送消息是否可以提供声音、震动等提醒选项设置？
// 答：消息推送时目前没有提供提醒方式的选择，所以通知都是使用系统默认的提醒方式。
// https://open.oppomobile.com/new/developmentDoc/info?id=11256

// Notification 通知栏消息
// https://open.oppomobile.com/new/developmentDoc/info?id=11236
type Notification struct {
	AppMessageID        string      `json:"app_message_id,omitempty"`        // App开发者自定义消息Id，主要用于消息去重。对于广播消息，相同app_message_id只会保存一条；对于单推消息，相同app_message_id的消息只会对同一个目标推送一次。
	Style               int         `json:"style,omitempty"`                 // 通知栏样式 1. 标准样式 2. 长文本样式（ColorOS版本>5.0可用，通知栏第一条消息可展示全部内容，非第一条消息只展示一行内容） 3. 大图样式（ColorOS版本>5.0可用，通知栏第一条消息展示大图，非第一条消息不显示大图，推送方式仅支持广播，且不支持定速功能）
	BigPictureId        string      `json:"big_picture_id,omitempty"`        // 大图id【style为3时，必填】,通过上传大图接口获得大图id后可使用。上传大图接口请参考服务端API介绍章节
	SmallPictureId      string      `json:"small_picture_id,omitempty"`      // 通知图标id,通过上传小图接口获得小图id后可使用。上传小图接口请参考服务端API介绍章节。
	Title               string      `json:"title,omitempty"`                 // 设置在通知栏展示的通知栏标题, 【字数串长度限制在50个字符内，中英文字符及特殊符号（如emoji）均视为一个字符】
	SubTitle            string      `json:"sub_title,omitempty"`             // 子标题，设置在通知栏展示的通知栏标题, 【字符串长度限制在10个字符以内，中英文字符及特殊符号（如emoji）均视为一个字符计算】
	Content             string      `json:"content,omitempty"`               // 设置在通知栏展示的通知的正文内容 1）当选择标准样式（style 设置为 1）时，内容字符串长度限制在50以内； 2）当选择长文本样式（style设置 为 2）时，内容字符串长度限制在128以内； 3）当选择大图样式（style 设置为 3）时，内容字符串长度限制在50以内。 【字符串长度计算说明：中英文字符及特殊符号（如emoji）均视作一个字符计算】
	ClickActionType     int         `json:"click_action_type,omitempty"`     // 点击通知栏后触发的动作类型。点击动作类型值的定义和含义如下：0.启动应用；1.跳转指定应用内页（action标签名）；2.跳转网页；4.跳转指定应用内页（全路径类名）；【非必填，默认值为0】;5.跳转Intent scheme URL
	ClickActionActivity string      `json:"click_action_activity,omitempty"` // 当设置click_action_type为1或者4时，需要配置本参数。应用内页地址【click_action_type为1/4/时必填，长度500】
	ClickActionURL      string      `json:"click_action_url,omitempty"`      // 跳转URL，当跳转的形式为URL时，click_action_type参数需要设置为2或5，同时设置本参数。本参数接受最大长度2000以内的URL。
	ActionParameters    string      `json:"action_parameters,omitempty"`     // 跳转动作参数。打开应用内页或网页时传递给应用或网页的附加参数【JSON格式】，字符串长度不超过4000。当跳转类型是URL类型时，参数会以URL参数直接拼接在URL后面。示例：{“key1”:“value1”,“key2”:“value2”}
	ShowTimeType        int         `json:"show_time_type,omitempty"`        // 通知栏展示类型。展示类型如下 0：即时展示 1：定时展示，配置该参数后定时展示开始时间（show_start_time）及定时展示的结束时间（show_end_time）为必填
	ShowStartTime       int64       `json:"show_start_time,omitempty"`       // 定时展示的开始时间。选择定时展示后，消息将于设定的开始时间到结束时间之内展示，展示开始时间不能大于等于展示结束时间。 本参数接受13位的unix时间戳。
	ShowEndTime         int64       `json:"show_end_time,omitempty"`         // 定时展示的结束时间。选择定时展示后，消息将于设定的开始时间到结束时间之内展示。本参数接受13位的unix时间戳。
	OffLine             bool        `json:"off_line"`                        // 是否是离线消息。如果是离线消息，OPPO PUSH在设备离线期间缓存消息一段时间，等待设备上线接收。
	OffLineTTL          int         `json:"off_line_ttl,omitempty"`          // 离线消息的存活时间，单位是秒。存活时间最大允许设置为10天，参数超过10天以10天传入。
	PushTimeType        int         `json:"push_time_type,omitempty"`        // 定时推送 (0, “即时”),(1, “定时”), 【只对全部用户推送生效】
	PushStartTime       int64       `json:"push_start_time,omitempty"`       // 定时推送开始时间（根据time_zone转换成当地时间）, 【push_time_type 为1必填】，时间的毫秒数
	TimeZone            string      `json:"time_zone,omitempty"`             // 时区，默认值：（GMT+08:00）北京，香港，新加坡
	FixSpeed            bool        `json:"fix_speed"`                       // 是否定速推送。广播类型消息专用，如果设置定速推送，消息将会以给定的速度均匀下发。
	FixSpeedRate        int64       `json:"fix_speed_rate,omitempty"`        // 定速推送的速率，单位为条每秒。指定消息为定速推送消息时，需要指定本参数。定速推送速率范围在[1000, 10000]。
	NetworkType         int         `json:"network_type,omitempty"`          // 推送的网络环境类型。本参数将影响用户设备仅在指定类型的网络环境下接收消息。参数定义如下：0：不限联网方式；1：仅wifi推送，设置后，消息只会在用户处于WiFi环境下才下发。
	CallBackURL         string      `json:"call_back_url,omitempty"`         // 回执功能详见回执一章仅支持registrationId推送方式开发者接收消息送达的回执消息的URL地址。https://open.oppomobile.com/new/developmentDoc/info?id=11239
	CallBackParameter   string      `json:"call_back_parameter,omitempty"`   // 开发者指定的自定义回执参数。参数字符串长度限制在100以内，OPPO PUSH将这个参数设置在回执请求体单个JSON结构的param字段中。
	ChannelID           string      `json:"channel_id,omitempty"`            // 指定下发的通道ID。通知栏通道（NotificationChannel），从Android9开始，Android设备发送通知栏消息必须要指定通道ID，（如果是快应用，必须带置顶的通道Id:OPPO PUSH推送）
	ShowTtl             int         `json:"show_ttl,omitempty"`              // 限时展示时间(单位：秒)。消息在通知栏展示后开始计时，展示时长超过展示事件后，消息会从通知栏中消失。限时展示的时间范围：公信0-12小时，私信0-24小时。
	NotifyId            int         `json:"notify_id,omitempty"`             // 每条消息在通知显示时的唯一标识，主要用于新旧消息的覆盖。不设置本参数时，PUSH自动为给每条消息生成一个唯一标识；当不同的消息设置为同一个notify_id，到达设备的新消息将覆盖旧消息展示在设备通知栏中。
	AuditResponse       interface{} `json:"auditResponse,omitempty"`         // 推必安信息审核api响应内容，详见《基于第三方审核结果的消息推送》 https://open.oppomobile.com/new/developmentDoc/info?id=11344
}

type SendRes struct {
	Code    int    `json:"code"`    // 返回码,请参考公共返回码与接口返回码
	Message string `json:"message"` // 错误详细信息，不存在则不填
	Data    struct {
		MessageID string `json:"messageId"` // 消息 ID
	} `json:"data"` // 返回值，JSON类型
}

type AuthReq struct {
	AppKey    string `json:"app_key,omitempty"`   // OPPO PUSH发放给合法应用的AppKey。
	Sign      string `json:"sign,omitempty"`      // 加密签名。是用AppKey、当前时间戳毫秒数、MasterSecret拼接而成的字符串并用SHA256加密而成的字符串。MasterSecret是注册应用时OPPO PUSH发放的服务端密钥，与AppKey对应
	Timestamp string `json:"timestamp,omitempty"` // 当前时间的unix时间戳。格式为13位时间毫秒数，时区采用GMT+8。需要使用最近10分钟内的时间戳，否则会导致鉴权失败
}

type AuthRes struct {
	Code    int    `json:"code"`    // 返回码,请参考公共返回码与接口返回码
	Message string `json:"message"` // 错误详细信息，不存在则不填
	Data    struct {
		AuthToken  string `json:"auth_token"`  // 权限令牌，推送消息时，需要提供auth_token，有效期默认为24小时，过期后无法使用
		CreateTime int64  `json:"create_time"` // 时间毫秒数
	} `json:"data"` // 返回值，JSON类型，包含响应结构体
}
