package untils

type WXBizDataCrypt struct {
	appid string
	sessionKey string
}

const (
	ok = 0
	IllegalAesKey = -41001
	IllegalIv = -41002
	IllegalBuffer = 41003
	DecodeBase64Error = 41004
)
