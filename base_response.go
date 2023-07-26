package openwechat

type Ret int

const (
    ticketError         Ret = -14  // ticket error
    loginError          Ret = -2   // login error
    sysError            Ret = -1   // sys error
    paramError          Ret = 1    // param error
    failedLoginWarn     Ret = 1100 // failed login warn
    failedLoginCheck    Ret = 1101 // failed login check
    cookieInvalid       Ret = 1102 // cookie invalid
    loginEnvAbnormality Ret = 1203 // login environmental abnormality
    optTooOften         Ret = 1205 // operate too oftn
)
