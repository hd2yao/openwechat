package openwechat

import "errors"

var (
    // ErrForbidden define login forbidden error
    ErrForbidden = errors.New("login forbidden")

    // ErrInvalidStorage define invalid storage error
    ErrInvalidStorage = errors.New("invalid storage")

    // NetworkErr define wechat network error
    NetworkErr = errors.New("wechat network error")

    // ErrNoSuchUserFoundError define no such user found error
    ErrNoSuchUserFoundError = errors.New("no such user found")

    // ErrLoginTimeout define login timeout error
    ErrLoginTimeout = errors.New("login timeout")

    // ErrWebWxDataTicketNotFound define webwx_data_ticket not found error
    ErrWebWxDataTicketNotFound = errors.New("webwx_data_ticket not found")

    // ErrUserLogout define user logout error
    ErrUserLogout = errors.New("user logout")
)

func (r Ret) Error() string {
    return r.String()
}
