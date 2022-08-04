package base

const (
	// HytHost 华医通host
	HytHost         = "https://hytapiv2.cd120.com"
	HytDoctorList   = "/cloud/hosplatcustomer/call/appointment/doctorListModel/selDoctorListByMoreTerm" // 预约日医生列表
	HytDoctorDetail = "/cloud/hosplatcustomer/call/appointment/selDoctorDetailsTwo"                     // 预约医生详情
	HytHospitalArea = "/cloud/hosplatcustomer/call/appointment/hospitalConfig/selHospitalAreaRecord"    // 获取医院所有院区
	HytDept         = "/cloud/hosplatcustomer/call/appointment/dept/selFirstDept"                       // 院区科室分类
	// HxHost 华西
	HxHost         = "https://hxgyapiv2.cd120.info"
	HxCardListUrl  = "/cloud/hosplatcustomer/cardservice/cardlist"        // 就诊卡列表
	HxImageCodeUrl = "/cloud/hosplatcustomer/customer/image/getimagecode" // 获取验证码
)
