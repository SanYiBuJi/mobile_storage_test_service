package Models

type CaesarTestService struct {
}

type ApplicationFormData struct {
	ID              uint
	FormID          string //表单ID
	AssetOID        string //资产信息
	AssetID         string
	CategoryName    string //设备分类
	ClientID        string //终端ID
	ClientIP        string //终端IP
	ClientName      string //终端名称
	DeviceID        string //设备硬件编号
	DeviceName      string //产品型号
	DeviceUID       string //设备UID
	Intranet        int    //内网有效期
	Reason          string //申请理由
	Remark          string //管理员答复
	ResponsibleName string //责任人名称
	StaffID         int    //用户ID
	StaffName       string //用户组织
	Type            int    //申请类型
	Status          int    //申请状态
	StartTime       string //发起时间
	CreateTime      string
	UpdateTime      string
}

type AcceptApplicationFormRequest struct {
	ID              string `json:"id"`        //表单ID
	AssetOID        string `json:"asset_oid"` //资产信息
	AssetID         string `json:"asset_id"`
	CategoryName    string `json:"category_name"`    //设备分类
	ClientID        string `json:"client_id"`        //终端ID
	ClientIP        string `json:"client_ip"`        //终端IP
	ClientName      string `json:"client_name"`      //终端名称
	DeviceID        string `json:"device_id"`        //设备硬件编号
	DeviceName      string `json:"device_name"`      //产品型号
	DeviceUID       string `json:"device_uid"`       //设备UID
	Intranet        int    `json:"intranet"`         //内网有效期
	Reason          string `json:"reason"`           //申请理由
	Remark          string `json:"remark"`           //管理员答复
	ResponsibleName string `json:"responsible_name"` //责任人名称
	StaffID         int    `json:"staff_id"`         //用户ID
	StaffName       string `json:"staff_name"`       //用户组织
	Type            int    `json:"type"`             //申请类型
	Status          int    `json:"status"`           //申请状态
	StartTime       string `json:"start_time"`       //发起时间
}

type BatchGetApplicationFormRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type BatchGetApplicationFormResponse struct {
	Data *[]ApplicationFormData
}

type UpdateApplicationFormTestV1Request struct {
	ID        string `json:"id"`
	AssetId   string `json:"asset_id"`
	AssetOID  string `json:"asset_oid"`
	Status    int    `json:"status"`
	Reason    string `json:"reason"`
	Appkey    string `json:"appkey"`
	Appsecret string `json:"appsecret"`
	Addr      string `json:"addr"`
	Port      string `json:"port"`
}

type CaesarMobileStoragePutCheck struct {
	ID       string `json:"id"`
	AssetID  string `json:"asset_id"`
	AssetOID string `json:"asset_oid"`
	Reason   string `json:"reason"`
	Status   int    `json:"status"`
}
