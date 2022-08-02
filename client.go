package main

import (
	"github.com/guonaihong/gout"
)

type (
	Client struct {
		config Config
	}
	Config struct {
		Token string
		Debug bool
	}
)

type IClient interface {
	GetToken() string
	GenerateHeader() interface{}
	CardList(req CardReq) (resp CardResp, err error)
	DoctorList(req DoctorListReq) (resp DoctorListResp, err error)
	DoctorDetail(req DoctorDetailReq) (resp DoctorDetailResp, err error)
}

func NewClient(config Config) IClient {
	return &Client{
		config: config,
	}
}

func (c *Client) GetToken() string {
	return c.config.Token
}

// GenerateHeader header token&timespan
func (c *Client) GenerateHeader() (header interface{}) {
	// X-Requested-With 请求头 区分ajax请求还是普通请求
	// 如果requestedWith为null，则为同步请求。
	// 如果requestedWith为XMLHttpRequest则为Ajax请求。
	// com.tencent.mm 微信请求防止被拦截
	return gout.H{"Token": c.config.Token, "accessToken": c.config.Token, "X-Requested-With": "com.tencent.mm"}

}
