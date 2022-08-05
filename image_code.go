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
		HxDelImageCode(*Client, string, ...string) error
		HxUpdateImageCodeCheck(*Client, string, string, string)
	}
)

// HxDelImageCode 删除验证码
func (i *ImageCode) HxDelImageCode(
	c *Client,
	idCard string,
	bizSecs ...string,
) (err error) {
	c.Lk.Lock()
	defer c.Lk.Unlock()
	if _, ok := c.Caches[idCard]; !ok {
		return
	}

	if len(bizSecs) == 0 {
		// 清空
		for index, value := range c.Caches[idCard] {
			if err = os.Remove(value.ImageUrl); err != nil {
				return
			}
			delete(c.Caches[idCard], index)
		}
		delete(c.Caches, idCard)
	} else {
		for _, bizSeq := range bizSecs {
			if _, ok := c.Caches[idCard][bizSeq]; ok {
				tempCode := c.Caches[idCard][bizSeq]
				if err = os.Remove(tempCode.ImageUrl); err != nil {
					return
				}
				delete(c.Caches[idCard], bizSeq)
			}
		}
	}
	return
}

// HxUpdateImageCodeCheck 设置验证码校验数字
func (i *ImageCode) HxUpdateImageCodeCheck(
	c *Client,
	idCard, bizSeq, check string,
) {
	c.Lk.Lock()
	defer c.Lk.Unlock()
	if _, ok := c.Caches[idCard]; !ok {
		return
	}
	if _, ok := c.Caches[idCard][bizSeq]; !ok {
		return
	}
	tempCode := c.Caches[idCard][bizSeq]
	tempCode.CheckNum = check
	c.Caches[idCard][bizSeq] = tempCode
}
