package main

type (
	OrderListReq struct {
	}
	OrderListResp struct {
	}
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
	// OrderRefundReq 取消订单
	OrderRefundReq struct {
	}
)
