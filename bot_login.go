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
