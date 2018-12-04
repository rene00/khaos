package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	UNAUTHORIZED:                    "unauthorized",
	INVALID_PARAMS:                  "è¯·æ±åæ°éè¯¯",
	ERROR_EXIST_TAG:                 "å·²å­å¨è¯¥æ ç­¾åç§°",
	ERROR_EXIST_TAG_FAIL:            "è·åå·²å­å¨æ ç­¾å¤±è´¥",
	ERROR_NOT_EXIST_TAG:             "è¯¥æ ç­¾ä¸å­å¨",
	ERROR_GET_TAGS_FAIL:             "è·åæææ ç­¾å¤±è´¥",
	ERROR_COUNT_TAG_FAIL:            "ç»è®¡æ ç­¾å¤±è´¥",
	ERROR_ADD_TAG_FAIL:              "æ°å¢æ ç­¾å¤±è´¥",
	ERROR_EDIT_TAG_FAIL:             "ä¿®æ¹æ ç­¾å¤±è´¥",
	ERROR_DELETE_TAG_FAIL:           "å é¤æ ç­¾å¤±è´¥",
	ERROR_EXPORT_TAG_FAIL:           "å¯¼åºæ ç­¾å¤±è´¥",
	ERROR_IMPORT_TAG_FAIL:           "å¯¼å¥æ ç­¾å¤±è´¥",
	ERROR_NOT_EXIST_ARTICLE:         "è¯¥æç« ä¸å­å¨",
	ERROR_ADD_ARTICLE_FAIL:          "æ°å¢æç« å¤±è´¥",
	ERROR_DELETE_ARTICLE_FAIL:       "å é¤æç« å¤±è´¥",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "æ£æ¥æç« æ¯å¦å­å¨å¤±è´¥",
	ERROR_EDIT_ARTICLE_FAIL:         "ä¿®æ¹æç« å¤±è´¥",
	ERROR_COUNT_ARTICLE_FAIL:        "ç»è®¡æç« å¤±è´¥",
	ERROR_GET_ARTICLES_FAIL:         "è·åå¤ä¸ªæç« å¤±è´¥",
	ERROR_GET_ARTICLE_FAIL:          "è·ååä¸ªæç« å¤±è´¥",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "çææç« æµ·æ¥å¤±è´¥",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Tokené´æå¤±è´¥",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Tokenå·²è¶æ¶",
	ERROR_AUTH_TOKEN:                "Tokençæå¤±è´¥",
	ERROR_AUTH:                      "Tokenéè¯¯",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "ä¿å­å¾çå¤±è´¥",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "æ£æ¥å¾çå¤±è´¥",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "æ ¡éªå¾çéè¯¯ï¼å¾çæ ¼å¼æå¤§å°æé®é¢",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
