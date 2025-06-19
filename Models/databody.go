package Models

import (
	"encoding/json"
)

type CaesarTestService struct {
}

//create table application_form_data
//(
//    id                     serial
//        primary key,
//    form_id                varchar(255) not null,
//    asset_o_id             varchar(255),
//    asset_id               varchar(255),
//    category_name          varchar(255),
//    client_id              varchar(255),
//    client_ip              varchar(255),
//    client_name            varchar(255),
//    device_id              varchar(255),
//    device_name            varchar(255),
//    device_uid             varchar(255),
//    intranet               bigint,
//    reason                 text,
//    remark                 text,
//    responsible_name       varchar(255),
//    staff_id               bigint       not null,
//    staff_name             varchar(255),
//    type                   integer      not null,
//    status                 integer      not null,
//    start_time             bigint,
//    create_time            timestamp,
//    update_time            timestamp,
//    out_use_auth           integer      not null,
//    out_use_auth_max_time  integer      not null,
//    out_use_auth_valid_day varchar(255),
//    client_node_name       varchar,
//    client_node_path       varchar,
//    custom_swap_area       boolean,
//    custom_swap_size       integer,
//    device_capacity        integer,
//    device_v_id            varchar,
//    last_category_name     varchar,
//    register_type          integer,
//    last_responsible_name  varchar,
//    apply_time             bigint,
//    apply_type             integer,
//    apply_type_desc        varchar,
//    client_type            varchar,
//    device_p_id            varchar,
//    staff_node_name        varchar,
//    status_desc            varchar
//);
//
//alter table application_form_data
//    owner to postgres;

type ApplicationFormData struct {
	ID                 uint
	FormID             string //表单ID
	AssetOID           string //资产信息
	AssetID            string
	CategoryName       string //设备分类
	ClientID           string //终端ID
	ClientIP           string //终端IP
	ClientName         string //终端名称
	DeviceID           string //设备硬件编号
	DeviceName         string //产品型号
	DeviceUID          string //设备UID
	Intranet           int64  //内网有效期
	Reason             string //申请理由
	Remark             string //管理员答复
	ResponsibleName    string //责任人名称
	StaffID            int64  //用户ID
	StaffName          string //用户组织
	Type               int64  //申请类型
	Status             int32  //申请状态
	StartTime          int64  //发起时间
	CreateTime         string
	UpdateTime         string
	OutUseAuth         int32
	OutUseAuthMaxTime  int32
	OutUseAuthValidDay string
	//ApplyReason         string
	ApplyTime           int64
	ApplyType           int64
	ApplyTypeDesc       string
	ClientNodeName      string
	ClientNodePath      string
	ClientType          string
	CustomSwapArea      bool
	CustomSwapSize      int
	DeviceCapacity      int64
	DevicePID           string
	DeviceVID           string
	LastCategoryName    string
	LastResponsibleName string
	RegisterType        int
	StaffNodeName       string
	StatusDesc          string
}

func (a *ApplicationFormData) String() string {
	application, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(application)
}

type AcceptApplicationFormRequest struct {
	ID           string `json:"id"`
	AssetOID     string `json:"asset_oid"`
	AssetID      string `json:"asset_id"`
	CategoryName string `json:"category_name"`
	ClientID     string `json:"client_id"`
	ClientIP     string `json:"client_ip"`
	ClientName   string `json:"client_name"`
	ClientType   string `json:"client_type"`
	DeviceID     string `json:"device_id"`
	DeviceName   string `json:"device_name"`
	DeviceUID    string `json:"device_uid"`
	DevicePID    string `json:"device_pid"`
	Intranet     int64  `json:"intranet"`
	//ApplyReason         string `json:"apply_reason"`
	Remark              string `json:"remark"`
	ResponsibleName     string `json:"responsible_name"`
	StaffID             int64  `json:"staff_id"`
	StaffNodeName       string `json:"staff_node_name"`
	StaffName           string `json:"staff_name"`
	ApplyType           int64  `json:"apply_type"`
	ApplyTypeDesc       string `json:"apply_type_desc"`
	Status              int32  `json:"status"`
	StatusDesc          string `json:"status_desc"`
	ApplyTime           int64  `json:"apply_time"`
	OutUseAuth          int32  `json:"out_use_auth"`      // 1=>禁止 2=>读写 3=>只写 4=>只读
	OutUseValidDay      int64  `json:"out_use_valid_day"` // 外出有效期 -1 永远有效
	OutUseMaxTimes      int32  `json:"out_use_max_times"` // 外出使用最大次数
	StartTime           string `json:"start_time"`
	FormID              string `json:"form_id"` //表单ID
	Reason              string `json:"reason"`  //申请理由
	Type                int32  `json:"type"`    //申请类型
	CreateTime          string `json:"create_time"`
	UpdateTime          string `json:"update_time"`
	OutUseAuthMaxTime   int32  `json:"out_use_auth_max_time"`
	OutUseAuthValidDay  string `json:"out_use_auth_valid_day"`
	ClientNodeName      string `json:"client_node_name"`
	ClientNodePath      string `json:"client_node_path"`
	CustomSwapArea      bool   `json:"custom_swap_area"`
	CustomSwapSize      int    `json:"custom_swap_size"`
	DeviceCapacity      int64  `json:"device_capacity"`
	DeviceVID           string `json:"device_vid"`
	LastCategoryName    string `json:"last_category_name"`    // 原设备分类
	LastResponsibleName string `json:"last_responsible_name"` //原责任人名称
	RegisterType        int    `json:"register_type"`
}

func (a *AcceptApplicationFormRequest) String() string {
	acceptAFR, err := json.MarshalIndent(1, "", "  ")
	if err != nil {
		return err.Error()
	} else {
		return string(acceptAFR)
	}
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
	AppKey    string `json:"appkey"`
	AppSecret string `json:"appsecret"`
	Addr      string `json:"addr"`
	Port      string `json:"port"`
}

type CaesarMobileStoragePutCheck struct {
	Data CaesarMobileStoragePutCheckData `json:"data"`
}

type CaesarMobileStoragePutCheckData struct {
	ID       string `json:"id"`
	AssetID  string `json:"asset_id"`
	AssetOID string `json:"asset_oid"`
	Reason   string `json:"reason"`
	Status   int    `json:"status"`
}
