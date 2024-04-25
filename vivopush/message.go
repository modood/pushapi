package vivopush

// PUSH-UPS-API接口文档
// https://dev.vivo.com.cn/documentCenter/doc/362
const (
	Host = "https://api-push.vivo.com.cn"

	AuthURL = "/message/auth" // 推送鉴权接口
	SendURL = "/message/send" // 单推接口
)

type AuthReq struct {
	AppId     string `json:"appId,omitempty"`     // 用户申请推送业务时生成的appId
	AppKey    string `json:"appKey,omitempty"`    // 用户申请推送业务时获得的appKey
	Timestamp int64  `json:"timestamp,omitempty"` // Unix13位毫秒时间戳 做签名用，单位：毫秒，且在vivo服务器当前utc时间戳前后十分钟区间内。该timestamp要与生成sign时使用的timestamp值相同。
	Sign      string `json:"sign,omitempty"`      // 签名 使用MD5算法，字符串拼接（appId+appKey+timestamp+appSecret），然后通过MD5加密得到的值（字母小写）。如sign生成示例所示，可使用示例参数生成sign，与示例结果对比是否一致，如一致表示生成正确，如不一致请排查按sign生成示例指引排查。
}

type AuthRes struct {
	Result    int    `json:"result"`    // 接口调用是否成功的状态码 0成功，非0失败
	Desc      string `json:"desc"`      // 文字描述接口调用情况
	AuthToken string `json:"authToken"` // 当鉴权成功时才会有该字段，推送消息时，需要提供authToken，有效期默认为1天，过期后无法使用。一个appId可对应多个token，24小时过期，业务方做中心缓存，1-2小时更新一次。
}

type SendReq struct {
	AppId           int                      `json:"appId,omitempty"`           // 用户申请推送业务时生成的appId，用于与获取authToken时传递的appId校验，一致才可以推送
	RegId           string                   `json:"regId,omitempty"`           // 应用订阅PUSH服务器得到的id（regId，alias 两者需一个不为空，当两个不为空时，取regId）
	Alias           string                   `json:"alias,omitempty"`           // 别名 长度不超过70字符（regId，alias两者需一个不为空，当两个不为空时，取regId）
	NotifyType      int                      `json:"notifyType,omitempty"`      // 通知类型 1:无，2:响铃，3:振动，4:响铃和振动
	Title           string                   `json:"title,omitempty"`           // 通知标题（用于通知栏消息） 最大40个字符（不区分中英文）
	Content         string                   `json:"content,omitempty"`         // 通知内容（用于通知栏消息） 最大100个字符（不区分中英文）
	TimeToLive      int64                    `json:"timeToLive,omitempty"`      // 消息缓存时间，单位是秒。在用户设备没有网络时，消息在Push服务器进行缓存，在消息缓存时间内用户设备重新连接网络，消息会下发，超过缓存时间后消息会丢弃。取值至少60秒，最长一天。当值为空时，默认一天
	SkipType        int                      `json:"skipType,omitempty"`        // 点击跳转类型 1：打开APP首页 2：打开链接 3：自定义 4:打开app内指定页面
	SkipContent     string                   `json:"skipContent,omitempty"`     // 跳转内容 跳转类型为2或3或4时，跳转内容最大1024个字符，skipType传3需要在onNotificationMessageClicked回调函数中自己写处理逻辑。关于skipContent的内容可以参考 【vivo推送常见问题汇总】 https://dev.vivo.com.cn/documentCenter/doc/156
	NetworkType     int                      `json:"networkType,omitempty"`     // 网络方式 -1：不限，1：wifi下发送，不填默认为-1
	Classification  int                      `json:"classification,omitempty"`  // 消息类型 0：运营类消息，1：系统类消息。不填默认为0
	ClientCustomMap map[string]string        `json:"clientCustomMap,omitempty"` // 客户端自定义键值对 key和Value键值对总长度不能超过1024字符。app可以按照客户端SDK接入文档获取该键值对
	Extra           map[string]string        `json:"extra,omitempty"`           // 高级特性（详见目录：一.公共——4.高级特性 extra） https://dev.vivo.com.cn/documentCenter/doc/362#s-k2w30pkd
	RequestId       string                   `json:"requestId,omitempty"`       // 用户请求唯一标识 最大64字符
	PushMode        int                      `json:"pushMode,omitempty"`        // 推送模式 0：正式推送；1：测试推送，不填默认为0（测试推送，只能给web界面录入的测试用户推送；审核中应用，只能用测试推送）
	AuditReview     []map[string]interface{} `json:"auditReview,omitempty"`     // 第三方审核结果，参见：基于第三方审核结果的消息推送 https://dev.vivo.com.cn/documentCenter/doc/585
	NotifyId        int                      `json:"notifyId,omitempty"`        // 每条消息在通知显示时的唯一标识。不携带时，vpush自动为给每条消息生成一个唯一标识；当不同的消息设置为同一个notifyId，到达设备的新消息将覆盖旧消息展示在设备通知栏中。值范围：1~2147483647。
	Category        string                   `json:"category,omitempty"`        // 二级分类，传值参见：二级分类标准 中category说明 https://dev.vivo.com.cn/documentCenter/doc/359 1、填写category后，可以不填写classification，但若填写classification，请保证category与classification是正确对应关系，否则返回错误码10097； 2、赋值请按照消息分类规则填写，且必须大写；若传入错误无效的值，否则返回错误码10096；
	ProfileId       string                   `json:"profileId,omitempty"`       // 关联终端设备登录用户标识，最大长度为64，仅单推支持
	SendOnline      bool                     `json:"sendOnline"`                // 是否在线直推，设置为true表示是在线直推，false表示非直推。在线直推功能推送时在设备在线下发一次，设备离线直接丢弃。详情请参见：在线直推 https://dev.vivo.com.cn/documentCenter/doc/743
	ForegroundShow  bool                     `json:"foregroundShow"`            // 是否前台通知展示，设置为false表示应用在前台则不展示通知消息，true表示无论应用是否在前台都展示通知。
}

type SendRes struct {
	Result      int    `json:"result"` // 接口调用是否成功的状态码 0成功，非0失败
	Desc        string `json:"desc"`   // 文字描述接口调用情况
	TaskId      string `json:"taskId"` // 任务编号
	InvalidUser *struct {
		UserId int `json:"userid"` // userid为接入方传的regId或者alias
		Status int `json:"status"` // status有3种情况： 1.userid不存在(userid与appId绑定，该userid无法在当前应用中找到)； 2.卸载, 或主动触发解订阅, 或用户清除数据(用户清除数据会使客户端SDK触发解订阅) 4.非测试用户
	} `json:"invalidUser"` // 非法用户信息，包括status和userid
}
