package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "invalid params",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token timeout",
	ERROR_AUTH_TOKEN:               "token error",
	ERROR_AUTH:                     "auth error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}