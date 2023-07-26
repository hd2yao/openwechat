package openwechat

import "strconv"

const (
	_Ret_name_0 = "ticket error"
	_Ret_name_1 = "login errorsys error"
	_Ret_name_2 = "param error"
	_Ret_name_3 = "failed login warnfailed login checkcookie invalid"
	_Ret_name_4 = "login environmental abnormality"
	_Ret_name_5 = "operate too often"
)

var (
	_Ret_index_1 = [...]uint8{0, 11, 20}
	_Ret_index_3 = [...]uint8{0, 17, 35, 49}
)

func (r Ret) String() string {
	switch {
	case r == -14:
		return _Ret_name_0
	case -2 <= r && r <= -1:
		return _Ret_name_1[_Ret_index_1[r]:_Ret_index_1[r+1]]
	case r == 1:
		return _Ret_name_2
	case 1100 <= r && r <= 1102:
		r -= 1100
		return _Ret_name_3[_Ret_index_3[r]:_Ret_index_3[r+1]]
	case r == 1203:
		return _Ret_name_4
	case r == 1205:
		return _Ret_name_5
	default:
		return "Ret(" + strconv.FormatInt(int64(r), 10) + ")"
	}
}
