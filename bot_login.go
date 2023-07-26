package openwechat

// LoginCode 定义登录状态码
type LoginCode string

const (
    // LoginCodeSuccess 登录成功
    LoginCodeSuccess LoginCode = "200"
    // LoginCodeScanned 已扫码
    LoginCodeScanned LoginCode = "201"
    // LoginCodeTimeout 登录超时
    LoginCodeTimeout LoginCode = "400"
    // LoginCodeWait 等待扫码
    LoginCodeWait LoginCode = "408"
)

func (l LoginCode) String() string {
    switch l {
    case LoginCodeSuccess:
        return "登录成功"
    case LoginCodeScanned:
        return "已扫码"
    case LoginCodeTimeout:
        return "登录超时"
    case LoginCodeWait:
        return "等待扫码"
    default:
        return "未知状态"
    }
}

type BotPreparer interface {
    Prepare(*Bot)
}

type BotLoginOption interface {
    BotPreparer
    OnError(*Bot, error) error
    OnSuccess(*Bot) error
}

// BotOptionGroup 是一个 BotLoginOption 的集合
// 用户将多个 BotLoginOption 组合成一个 BotLoginOption
type BotOptionGroup []BotLoginOption

func (g BotOptionGroup) Prepare(bot *Bot) {
    for _, option := range g {
        option.Prepare(bot)
    }
}

func (g BotOptionGroup) OnError(bot *Bot, err error) error {
    // 当有一个 BotLoginOption 的 OnError 返回的 error == nil 时，就会停止执行后续的 BotLoginOption
    for _, option := range g {
        currentErr := option.OnError(bot, err)
        if currentErr == nil {
            return nil
        }
        if currentErr != nil {
            return currentErr
        }
    }
    return err
}

func (g BotOptionGroup) OnSuccess(bot *Bot) error {
    for _, option := range g {
        if err := option.OnSuccess(bot); err != nil {
            return err
        }
    }
    return nil
}
