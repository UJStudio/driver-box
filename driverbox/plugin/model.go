package plugin

import (
	"github.com/ibuilding-x/driver-box/driverbox/event"
)

// 触发 ExportTo 的类型
type ExportType string

// EncodeMode 编码模式
type EncodeMode string

const (
	ReadMode       EncodeMode = "read"           // 读模式
	WriteMode      EncodeMode = "write"          // 写模式
	RealTimeExport ExportType = "realTimeExport" //实时上报
)

// PointData 点位数据
type PointData struct {
	PointName string      `json:"name"`  // 点位名称
	Value     interface{} `json:"value"` // 点位值
}

// DeviceData 设备数据
type DeviceData struct {
	ID         string       `json:"id"`
	Values     []PointData  `json:"values"`
	Events     []event.Data `json:"events"`
	ExportType ExportType   //上报类型，底层的变化上报和实时上报等同于RealTimeExport
}

// PointReadValue 点位读操作的结构体
type PointReadValue struct {
	//设备 ID
	ID string `json:"id"`
	// PointName 点位名称
	PointName string `json:"pointName"`
	// Value 点位值
	Value interface{} `json:"value"`
}

// PointWriteValue 点位写操作的结构体
type PointWriteValue struct {
	// PointName 点位名称
	PointName string `json:"pointName"`
	// Value 点位值
	Value interface{} `json:"value"`
	//模型名称，某些驱动解析需要根据模型作区分
	ModelName string `json:"modelName"`
	//前置操作，例如空开要先解锁，空调要先开机
	PreOp []PointWriteValue `json:"preOp"`
}

// 连接配置基础模型
type BaseConnection struct {
	ConnectionKey string //连接标识
	//ScriptEnable  bool   //是否存在动态脚本
	////当前连接的 lua 虚拟机
	//Ls          *lua.LState
	ProtocolKey string `json:"protocolKey"` //协议驱动库标识
	Discover    bool   `json:"discover"`    //是否支持设备发现
	Enable      bool   `json:"enable"`      //是否启用
	Virtual     bool   `json:"virtual"`     //虚拟设备功能
}

// 简单的编码结构体,对于无Encode实现的插件，可以使用该结构体
type SimpleEncodeStruct struct {
	// 设备 ID
	ID string `json:"id"`
	// 点位名称
	Mode EncodeMode `json:"mode"`
	// 点位值
	Values []PointData `json:"values"`
}
