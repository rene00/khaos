package e

var MsgFlags = map[int]string{

	SUCCESS:      "OK",
	ERROR:        "Fail",
	UNAUTHORIZED: "Unauthorized",
	BAD_REQUEST:  "Bad Request",

	ERROR_AUTH:                     "Authentication Error",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token Fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token Timeout",
	ERROR_AUTH_TOKEN:               "Token Error",
	ERROR_AUTH_FAILED_TO_DECODE:    "Failed to Decode Authentication Header",
	ERROR_AUTH_INVALID_HEADER:      "Invalid Authentication Header",
	ERROR_PING:                     "Ping Error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
