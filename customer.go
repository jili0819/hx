package hx

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/jili0819/hx/base"
	"os"
)

type (
	CustomerInfo struct {
		Cards  []Card // 就诊卡列表
		Token  string `json:"token"`  // token
		IDCard string `json:"idCard"` // 身份证号码
		Name   string `json:"name"`   // 姓名
	}
	CardReq struct {
		AppCode     string `json:"appCode"`
		OrganCode   string `json:"organCode"`
		ChannelCode string `json:"channelCode"`
		Guidance    string `json:"guidance"`
	}
	CardResp struct {
		Code    string `json:"code"`
		ErrCode string `json:"errCode"`
		Msg     string `json:"msg"`
		Data    Card   `json:"data"`
	}
	Card struct {
		CanBindCount int         `json:"canBindCount"`
		OpenEhCard   interface{} `json:"openEhCard"`
		UserCardList []struct {
			Age             int           `json:"age"`
			BankCardBottom  interface{}   `json:"bankCardBottom"`
			BankCardIcon    interface{}   `json:"bankCardIcon"`
			BankName        interface{}   `json:"bankName"`
			CardId          string        `json:"cardId"`
			CardInsChannel  []interface{} `json:"cardInsChannel"`
			CardNo          string        `json:"cardNo"`
			CardType        string        `json:"cardType"`
			CardTypeDesc    string        `json:"cardTypeDesc"`
			ChannelCode     string        `json:"channelCode"`
			CredNo          string        `json:"credNo"`
			DetailAddress   string        `json:"detailAddress"`
			Gender          int           `json:"gender"`
			GuarderRelation interface{}   `json:"guarderRelation"`
			IndexDefault    bool          `json:"indexDefault"`
			IsSelf          int           `json:"isSelf"`
			ListDefault     bool          `json:"listDefault"`
			MemberRelation  string        `json:"memberRelation"`
			OpenEhCard      interface{}   `json:"openEhCard"`
			OrganCode       string        `json:"organCode"`
			OrganInsChannel interface{}   `json:"organInsChannel"`
			PatientId       string        `json:"patientId"`
			PatientName     string        `json:"patientName"`
			Pmi             string        `json:"pmi"`
			PmiNo           string        `json:"pmiNo"`
			QrCode          string        `json:"qrCode"`
			RealName        bool          `json:"realName"`
			Status          int           `json:"status"`
			Tel             string        `json:"tel"`
		} `json:"userCardList"`
	}
)

type ICustomer interface {
	HxGetCard(*Client, CardReq) (CardResp, error)
	HxCardList(*Client) []Card
	HxGetToken(*Client) string
	HxUpdateToken(*Client, string)
	HxGenerateImageCode(*Client, ImageCodeReq) error
	HxGetImageCodeRand(c *Client) (ImageCode, error)
}

func (customer *CustomerInfo) HxGetCard(
	c *Client,
	cardReq CardReq,
) (
	cardResp CardResp,
	err error,
) {
	header := c.GenerateHeader(customer.IDCard)
	req := gout.H{
		"appCode":     cardReq.AppCode,
		"organCode":   cardReq.OrganCode,
		"channelCode": cardReq.ChannelCode,
		"guidance":    cardReq.Guidance,
	}
	cardResp = CardResp{}
	if err = gout.POST(base.HxHost + base.HxCardListUrl).
		Debug(c.Config().Debug).
		SetHeader(header).
		SetJSON(req).
		BindJSON(&cardResp).
		Do(); err != nil {
		return
	}
	if cardResp.Code == "0" {
		err = fmt.Errorf("url:%s,code:%s,errCode:%s,msg:%s", base.HxHost+base.HxCardListUrl, cardResp.Code, cardResp.ErrCode, cardResp.Msg)
		return
	}
	return
}

func (customer *CustomerInfo) HxCardList(
	c *Client,
) (
	cards []Card,
) {
	c.Lk.RLock()
	defer c.Lk.RUnlock()
	if _, ok := c.Customers[customer.IDCard]; ok {
		tempCustomer := c.Customers[customer.IDCard]
		return tempCustomer.Cards
	}
	return nil
}

// HxGetToken
//  @Description: 根据身份证号码获取token
//  @receiver c
//  @param idCard
//  @return string
//
func (customer *CustomerInfo) HxGetToken(
	c *Client,
) string {
	c.Lk.RLock()
	defer c.Lk.RUnlock()
	return c.Customers[customer.IDCard].Token
}

//
// HxUpdateToken
//  @Description: 更新就诊人token
//  @receiver customer
//  @param c
//  @param idCard
//  @param token
//
func (customer *CustomerInfo) HxUpdateToken(
	c *Client,
	token string,
) {
	c.Lk.Lock()
	defer c.Lk.Unlock()
	temp := c.Customers[customer.IDCard]
	temp.Token = token
	c.Customers[customer.IDCard] = temp
}

//
// HxGenerateImageCode
//  @Description: 生成验证码
//  @receiver i
//  @param c
//  @param idCard
//  @param imageCodeReq
//  @param
//  @return err
//
func (customer *CustomerInfo) HxGenerateImageCode(
	c *Client,
	imageCodeReq ImageCodeReq,
) (err error) {
	header := c.GenerateHeader(customer.IDCard)
	req := gout.H{
		"appCode":     imageCodeReq.AppCode,
		"organCode":   imageCodeReq.OrganCode,
		"channelCode": imageCodeReq.ChannelCode,
		"type":        imageCodeReq.Type,
	}
	imageCodeResp := ImageCodeResp{}
	if err = gout.POST(base.HxHost + base.HxImageCodeUrl).
		Debug(c.Config().Debug).
		SetHeader(header).
		SetJSON(req).
		BindJSON(&imageCodeResp).
		Do(); err != nil {
		return
	}
	if imageCodeResp.Code == "0" {
		err = fmt.Errorf("url:%s,code:%s,errCode:%s,msg:%s", base.HxHost+base.HxImageCodeUrl, imageCodeResp.Code, imageCodeResp.ErrCode, imageCodeResp.Msg)
		return
	}
	// 生成验证码图片，并保存
	// todo
	url := "url"
	code := ImageCode{
		CardNo:    "",
		BizSeq:    imageCodeResp.Data.BizSeq,
		ImageData: imageCodeResp.Data.ImageData,
		ImageUrl:  url,
	}
	setImageCodeCache(c, customer.IDCard, code)
	return
}

//
// HxGetImageCodeRand
//  @Description: 随机获取用户所属图片验证码
//  @receiver i
//  @param c
//  @return err
//
func (customer *CustomerInfo) HxGetImageCodeRand(
	c *Client,
) (
	imageCode ImageCode,
	err error,
) {
	c.Lk.RLock()
	defer c.Lk.RUnlock()
	if _, ok := c.Caches[customer.IDCard]; !ok {
		// 还未给就诊人生成验证码
		if err = customer.HxGenerateImageCode(c, ImageCodeReq{
			AppCode:     c.Config().AppCode,
			OrganCode:   c.Config().OrganCode,
			ChannelCode: "PATIENT_WECHAT",
			Type:        "WEB",
		}); err != nil {
			return
		}
		if imageCode, err = customer.HxGetImageCodeRand(c); err != nil {
			return
		}
	} else if len(c.Caches[customer.IDCard]) == 0 {
		// 就诊人验证码使用完了
		if err = customer.HxGenerateImageCode(c, ImageCodeReq{
			AppCode: c.Config().AppCode,
			//OrganCode:   c.GetClient().Config.,
			ChannelCode: "PATIENT_WECHAT",
			Type:        "WEB",
		}); err != nil {
			return
		}
		if imageCode, err = customer.HxGetImageCodeRand(c); err != nil {
			return
		}
	} else {
		for _, value := range c.Caches[customer.IDCard] {
			imageCode = value
			if err = os.Remove(value.ImageUrl); err != nil {
				return
			}
			delete(c.Caches[customer.IDCard], value.BizSeq)
			break
		}
	}
	return
}

//
//  setImageCodeCache
//  @Description: 保存验证码到缓存
//  @param c
//  @param idCard
//  @param code
//
func setImageCodeCache(c *Client, idCard string, code ImageCode) {
	c.Lk.Lock()
	defer c.Lk.Unlock()
	if _, ok := c.Caches[idCard]; !ok {
		c.Caches[idCard] = make(map[string]ImageCode)
	}
	c.Caches[idCard][code.BizSeq] = code
}
