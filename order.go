package hx

type (
	Order struct {
	}
	OrderCreateReq struct {
		AppCode          string `json:"appCode"`
		OrganCode        string `json:"organCode"`
		ChannelCode      string `json:"channelCode"`
		AppointmentType  int    `json:"appointmentType"`
		CardId           string `json:"cardId"`
		HospitalCode     string `json:"hospitalCode"`
		HospitalAreaCode string `json:"hospitalAreaCode"`
		SysScheduleId    string `json:"sysScheduleId"`
		Type             string `json:"type"`
		ImageId          string `json:"imageId"`
	}
	OrderListReq struct {
	}
	OrderListResp struct {
	}
	OrderRefundReq struct {
	}
	OrderRefundResp struct {
	}
)

func (c *Client) HxRob(req OrderCreateReq) {

}

func (c *Client) HxRefund() {

}
