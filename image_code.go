package hx

import (
	"os"
)

type (
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

	IImageCache interface {
		HxDelImageCode(IClient, string, ...string) error
		HxUpdateImageCodeCheck(IClient, string, string, string)
	}
)

// HxDelImageCode 删除验证码
func (i ImageCode) HxDelImageCode(
	c IClient,
	idCard string,
	bizSecs ...string,
) (err error) {
	c.Client().Lk.Lock()
	defer c.Client().Lk.Unlock()
	if _, ok := c.Client().Caches[idCard]; !ok {
		return
	}

	if len(bizSecs) == 0 {
		// 清空
		for index, value := range c.Client().Caches[idCard] {
			if err = os.Remove(value.ImageUrl); err != nil {
				return
			}
			delete(c.Client().Caches[idCard], index)
		}
		delete(c.Client().Caches, idCard)
	} else {
		for _, bizSeq := range bizSecs {
			if _, ok := c.Client().Caches[idCard][bizSeq]; ok {
				tempCode := c.Client().Caches[idCard][bizSeq]
				if err = os.Remove(tempCode.ImageUrl); err != nil {
					return
				}
				delete(c.Client().Caches[idCard], bizSeq)
			}
		}
	}
	return
}

// HxUpdateImageCodeCheck 设置验证码校验数字
func (i ImageCode) HxUpdateImageCodeCheck(c IClient, idCard, bizSeq, check string) {
	client := c.Client()
	client.Lk.Lock()
	defer client.Lk.Unlock()
	if _, ok := client.Caches[idCard]; !ok {
		return
	}
	if _, ok := client.Caches[idCard][bizSeq]; !ok {
		return
	}
	tempCode := client.Caches[idCard][bizSeq]
	tempCode.CheckNum = check
	client.Caches[idCard][bizSeq] = tempCode
}
