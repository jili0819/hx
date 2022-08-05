package hx

import (
	"github.com/guonaihong/gout"
	"sync"
)

var (
	defaultClient *Client
)

type (
	Client struct {
		Lk            sync.RWMutex
		config        *Config
		Customers     map[string]*CustomerInfo        // 多个就诊人(身份证key)
		Caches        map[string]map[string]ImageCode //验证码缓存器(身份证key)
		Hospitals     map[string]Hospital             // 院区
		HospitalAreas map[string][]DeptCategory       // 院区科室
	}
	Config struct {
		AppCode        string `json:"appCode"`
		OrganCode      string `json:"organCode"`
		ChannelCode    string `json:"channelCode"`
		Debug          bool   `json:"debug"`
		ImageCachePath string `json:"imageCachePath"`
	}
)

// IClient in(内部接口) out(调用外部接口)
type IClient interface {
	Client() *Client
	Config() *Config
	Customer(string) ICustomer
	AddCustomer(*CustomerInfo)
	// GenerateHeader in 生成就诊人header
	GenerateHeader(token string) interface{}

	// HxInitHospitalAreaRecord out 初始化所有院区
	HxInitHospitalAreaRecord(string, HospitalReq) error
	// HxInitDept 初始化院区科室
	HxInitDept(string, DeptReq) error

	HxDoctorList(idCard string, doctorListReq DoctorListReq) (doctorListResp DoctorListResp, err error)

	HxDoctorDetail(idCard string, doctorDetailReq DoctorDetailReq) (doctorDetailResp DoctorDetailResp, err error)
}

//
// NewClient
//  @Description: newClient
//  @param config
//  @return IClient
//
func NewClient(config *Config) *Client {
	if defaultClient == nil {
		return &Client{
			config:        config,
			Customers:     make(map[string]*CustomerInfo),
			Caches:        make(map[string]map[string]ImageCode),
			Hospitals:     make(map[string]Hospital),
			HospitalAreas: make(map[string][]DeptCategory),
		}
	}
	return defaultClient
}

//
// Config
//  @Description:
//  @receiver c
//  @return *Config
//
func (c *Client) Config() *Config {
	return c.config
}

//
// Client
//  @Description:
//  @receiver c
//  @return *Client
//
func (c *Client) Client() *Client {
	return defaultClient
}

//
// GenerateHeader生成token头
//  @Description:
//  @receiver c
//  @param key
//  @return header
//

func (c *Client) GenerateHeader(idCard string) (header interface{}) {
	c.Lk.RLock()
	defer c.Lk.RUnlock()
	// X-Requested-With 请求头 区分ajax请求还是普通请求
	// 如果requestedWith为null，则为同步请求。
	// 如果requestedWith为XMLHttpRequest则为Ajax请求。
	// com.tencent.mm 微信请求防止被拦截
	return gout.H{"Token": c.Customers[idCard].Token, "accessToken": c.Customers[idCard].Token, "X-Requested-With": "com.tencent.mm"}
}

//
// Customer
//  @Description: 获取就诊人
//  @receiver c
//  @param idCard
//  @return *CustomerInfo
//
func (c *Client) Customer(idCard string) *CustomerInfo {
	c.Lk.RLock()
	defer c.Lk.RUnlock()
	if _, ok := c.Customers[idCard]; ok {
		return c.Customers[idCard]
	}
	return nil
}

//
// AddCustomer
//  @Description:添加就诊人
//  @receiver customer
//  @param c
//  @param info
//
func (c *Client) AddCustomer(
	info *CustomerInfo,
) {
	c.Lk.Lock()
	defer c.Lk.Unlock()
	c.Customers[info.IDCard] = info
	return
}
