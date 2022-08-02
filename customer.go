package main

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/jili0819/hx/base"
)

type CardReq struct {
	AppCode     string `json:"appCode"`
	OrganCode   string `json:"organCode"`
	ChannelCode string `json:"channelCode"`
	Guidance    string `json:"guidance"`
}

type CardResp struct {
	Code    string `json:"code"`
	ErrCode string `json:"errCode"`
	Msg     string `json:"msg"`
	Data    Card   `json:"data"`
}

type Card struct {
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

// CardList 获取就诊卡列表
func (c *Client) CardList(cardReq CardReq) (cardResp CardResp, err error) {
	header := c.GenerateHeader()
	req := gout.H{
		"appCode":     cardReq.AppCode,
		"organCode":   cardReq.OrganCode,
		"channelCode": cardReq.ChannelCode,
		"guidance":    cardReq.Guidance,
	}
	cardResp = CardResp{}
	if err = gout.POST(base.HxHost + base.HxCardListUrl).
		Debug(c.config.Debug).
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
