package openwechat

const (
    appId           = "wx782c26e4c19acffb"
    appMessageAppId = "wxeb7ec651dd0aefa9"

    jsonContentType       = "application/json; charset=utf-8"
    uosPatchClientVersion = "2.0.0"
    uosPatchExtspam       = "Go8FCIkFEokFCggwMDAwMDAwMRAGGvAESySibk50w5Wb3uTl2c2h64jVVrV7gNs06GFlWplHQbY/5FfiO++1yH4ykC" +
        "yNPWKXmco+wfQzK5R98D3so7rJ5LmGFvBLjGceleySrc3SOf2Pc1gVehzJgODeS0lDL3/I/0S2SSE98YgKleq6Uqx6ndTy9yaL9qFxJL7eiA/R" +
        "3SEfTaW1SBoSITIu+EEkXff+Pv8NHOk7N57rcGk1w0ZzRrQDkXTOXFN2iHYIzAAZPIOY45Lsh+A4slpgnDiaOvRtlQYCt97nmPLuTipOJ8Qc5p" +
        "M7ZsOsAPPrCQL7nK0I7aPrFDF0q4ziUUKettzW8MrAaiVfmbD1/VkmLNVqqZVvBCtRblXb5FHmtS8FxnqCzYP4WFvz3T0TcrOqwLX1M/DQvcHa" +
        "GGw0B0y4bZMs7lVScGBFxMj3vbFi2SRKbKhaitxHfYHAOAa0X7/MSS0RNAjdwoyGHeOepXOKY+h3iHeqCvgOH6LOifdHf/1aaZNwSkGotYnYSc" +
        "W8Yx63LnSwba7+hESrtPa/huRmB9KWvMCKbDThL/nne14hnL277EDCSocPu3rOSYjuB9gKSOdVmWsj9Dxb/iZIe+S6AiG29Esm+/eUacSba0k8" +
        "wn5HhHg9d4tIcixrxveflc8vi2/wNQGVFNsGO6tB5WF0xf/plngOvQ1/ivGV/C1Qpdhzznh0ExAVJ6dwzNg7qIEBaw+BzTJTUuRcPk92Sn6QDn" +
        "2Pu3mpONaEumacjW4w6ipPnPw+g2TfywJjeEcpSZaP4Q3YV5HG8D6UjWA4GSkBKculWpdCMadx0usMomsSS/74QgpYqcPkmamB4nVv1JxczYIT" +
        "IqItIKjD35IGKAUwAA=="
)

// MessageType 以 Go 惯用形式定义了 PC 微信所有的官方消息类型
// 详见 message_test.go
type MessageType int

// AppMessageType 以 Go 惯用形式定义了 PC 微信所有的官方APP消息类型
type AppMessageType int

const (
    MsgTypeText           MessageType = 1     // 文本消息
    MsgTypeImage          MessageType = 3     // 图片消息
    MsgTypeVoice          MessageType = 34    // 语音消息
    MsgTypeVerify         MessageType = 37    // 认证消息
    MsgTypePossibleFriend MessageType = 40    // 好友推荐消息
    MsgTypeShareCard      MessageType = 42    // 名片消息
    MsgTypeVideo          MessageType = 43    // 视频消息
    MsgTypeEmotion        MessageType = 47    // 表情消息
    MsgTypeLocation       MessageType = 48    // 地理位置消息
    MsgTypeApp            MessageType = 49    // APP 消息
    MsgTypeVoip           MessageType = 50    // VOIP消息
    MsgTypeVoipNotify     MessageType = 52    // VOIP 结束消息
    MsgTypeVoipInvite     MessageType = 53    // VOIP 邀请
    MsgTypeMicroVideo     MessageType = 63    // 小视频消息
    MsgTypeSys            MessageType = 10000 // 系统消息
    MsgTypeRecalled       MessageType = 10002 // 消息撤回
)

const (
    AppMsgTypeText                  AppMessageType = 1      // 文本消息
    AppMsgTypeImg                   AppMessageType = 2      // 图片消息
    AppMsgTypeAudio                 AppMessageType = 3      // 语音消息
    AppMsgTypeVideo                 AppMessageType = 4      // 视频消息
    AppMsgTypeUrl                   AppMessageType = 5      // 文章消息
    AppMsgTypeAttach                AppMessageType = 6      // 附件消息
    AppMsgTypeOpen                  AppMessageType = 7      // Open
    AppMsgTypeEmoji                 AppMessageType = 8      // 表情消息
    AppMsgTypeVoiceRemind           AppMessageType = 9      // VoiceRemind
    AppMsgTypeScanGood              AppMessageType = 10     // ScanGood
    AppMsgTypeGood                  AppMessageType = 13     // Good
    AppMsgTypeEmotion               AppMessageType = 15     // Emotion
    AppMsgTypeCardTicket            AppMessageType = 16     // 名片消息
    AppMsgTypeRealtimeShareLocation AppMessageType = 17     // 地理位置消息
    AppMsgTypeTransfers             AppMessageType = 2000   // 转账消息
    AppMsgTypeRedEnvelopes          AppMessageType = 2001   // 红包消息
    AppMsgTypeReaderType            AppMessageType = 100001 //自定义的消息
)
