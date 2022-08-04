package cache

type (
	MemoryCache struct {
		ImageCodes map[string]map[string]ImageCode // 缓存多个就诊人的验证码
	}

	ImageCodeReq struct {
		AppCode     string `json:"appCode"`
		OrganCode   string `json:"organCode"`
		ChannelCode string `json:"channelCode"`
		Type        string `json:"type"`
	}

	ImageCodeResp struct {
		Code    string `json:"code"`
		ErrCode string `json:"errCode"`
		Msg     string `json:"msg"`
		Data    struct {
			BizSeq    string `json:"bizSeq"`
			ImageData string `json:"imageData"`
			Type      string `json:"type"`
		} `json:"data"`
	}

	ImageCode struct {
		CardNo    string `json:"cardNo"`    // 身份证号码
		BizSeq    string `json:"bizSeq"`    // 验证码ID
		ImageData string `json:"imageData"` // base64 图片
		CheckNum  string `json:"checkNum"`  // 实际验证码
		ImageUrl  string `json:"imageUrl"`  // 图片地址
	}
)
