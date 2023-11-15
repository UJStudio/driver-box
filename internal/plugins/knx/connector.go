package knx

import (
	"fmt"
	"github.com/ibuilding-x/driver-box/driverbox/common"
	"github.com/ibuilding-x/driver-box/driverbox/helper"
	"github.com/ibuilding-x/driver-box/driverbox/plugin"
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/dptwrap"
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/knx"
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/knx/cemi"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	ATTRI_ADDRESS          = "address"
	ATTRI_FEEDBACK_ADDR    = "feedbackAddr"
	ATTRI_POLL_ENABLE      = "pollEnable"
	ATTRI_DPT              = "dpt"
	ATTRI_READOKAFTERWRITE = "readOkAfterWrite"
)

// KnxTunnelDeviceConfig 连接器配置
type knxTunnelDeviceConfig struct {
	Host            string        `json:"host"`
	Port            int           `json:"port"`
	SrcAddr         string        `json:"srcAddr"` //%d.%d.%d
	LocalAddress    string        `json:"localAddress"`
	PollFrequence   time.Duration `json:"pollFrequence"`
	ResponseTimeout time.Duration `json:"responseTimeout"`
	SendLocal       bool          `json:"sendLocal"`
	UseTcp          bool          `json:"useTcp"`
}

// connector 连接器
type connector struct {
	plugin *Plugin
	config knxTunnelDeviceConfig
	tunnel *knx.GroupTunnel
	mutex  sync.Mutex
}

type rawData struct {
	DeviceName     string          `json:"deviceName"`
	PointRawValues []pointRawValue `json:"points"`
}

type pointRawValue struct {
	Name      string      `json:"name"`      // 点位名称
	RawValues []uint16    `json:"rawValues"` // 原始数据  按照寄存器顺序排列
	Value     interface{} `json:"value"`     // 按照标准类型解析后的数据
}

// Send 发送数据
func (c *connector) Send(data interface{}) (err error) {
	// 数据转换
	td, _ := data.(transportationData)
	// 读操作
	if td.Mode == plugin.ReadMode {
		err = c.sendReadMode(td)
	} else {
		err = c.sendWriteMode(td)
	}
	return
}

func (c *connector) sendWriteMode(td transportationData) error {
	dest, err := cemi.NewGroupAddrString(td.Address)
	if err != nil {
		return fmt.Errorf("invalid address: %v", td.Address)
	}
	idpt, ok := dptwrap.Produce(td.Dpt)
	if !ok {
		return fmt.Errorf("no such dpt: %v", td.Dpt)
	}
	idpt.Set(td.Value)
	err = c.tunnel.Send(knx.GroupEvent{
		Command:     knx.GroupWrite,
		Destination: dest,
		Data:        idpt.Pack(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *connector) sendReadMode(td transportationData) error {
	ext, err := getExtendProps(td.DeviceName, td.PointName)
	if err != nil {
		return err
	}
	return c.tunnel.Send(knx.GroupEvent{
		Command:     knx.GroupRead,
		Destination: ext.addr,
	})
}

func getExtendProps(deviceName, pointName string) (ext extendProps, err error) {
	point, ok := helper.CoreCache.GetPointByDevice(deviceName, pointName)
	if !ok {
		err = fmt.Errorf("point %s not found for device %s", pointName, deviceName)
	}
	err = helper.Map2Struct(point.Extends, &ext)
	address, registerType, err := castStartingAddress(point.Extends["startAddress"])
	if err != nil {
		return
	}
	ext.StartAddress = address
	if registerType != "" {
		ext.RegisterType = string(registerType)
	}
	if err != nil {
		return
	}
	return
}

// castStartingAddress modbus 地址转换
func castStartingAddress(i interface{}) (address uint16, registerType primaryTable, err error) {
	s := cast.ToString(i)
	if strings.HasPrefix(s, "0x") { //check hex format
		addr, err := strconv.ParseInt(s[2:], 16, 32)
		if err != nil {
			return 0, "", err
		}
		return cast.ToUint16(addr), "", nil
	} else if len(s) == 5 { //handle modbus format
		res, err := cast.ToUint16E(s)
		if err != nil {
			return 0, "", err
		}
		if res > 0 && res < 10000 {
			return res - 1, Coil, nil
		} else if res > 10000 && res < 20000 {
			return res - 10001, DiscreteInput, nil
		} else if res > 30000 && res < 40000 {
			return res - 30001, InputRegister, nil
		} else if res > 40000 && res < 50000 {
			return res - 40001, HoldingRegister, nil
		} else {
			return 0, "", err
		}
	}

	res, err := cast.ToUint16E(i)
	if err != nil {
		return 0, "", err
	}
	return res, "", nil
}

func (prv *pointRawValue) uint16sToValue(originalData []uint16, registerType string, rawType string,
	byteSwap bool, wordSwap bool) error {
	e := Endianness(byteSwap)
	w := WordOrder(wordSwap)
	bytes := Uint16sToBytes(e, originalData)
	length := len(originalData)
	switch strings.ToUpper(registerType) {
	case string(Coil), string(DiscreteInput):
		if length != 1 {
			return fmt.Errorf("illegal length %s", prv.Name)
		}
		prv.Value = originalData[0] == 1
		return nil
	case string(HoldingRegister), string(InputRegister):
		switch strings.ToUpper(rawType) {
		case strings.ToUpper(common.ValueTypeUint16):
			if length != 1 {
				return fmt.Errorf("illegal length %s", prv.Name)
			}
			prv.Value = BytesToUint16(e, bytes)
			return nil
		case strings.ToUpper(common.ValueTypeInt16):
			if length != 1 {
				return fmt.Errorf("illegal length %s", prv.Name)
			}
			prv.Value = cast.ToInt16(BytesToUint16(e, bytes))
			return nil
		case strings.ToUpper(common.ValueTypeUint32):
			if length != 2 {
				return fmt.Errorf("illegal length %s", prv.Name)
			}
			prv.Value = BytesToUint32s(e, w, bytes)[0]
			return nil
		case strings.ToUpper(common.ValueTypeInt32):
			if length != 2 {
				return fmt.Errorf("illegal length %s", prv.Name)
			}
			prv.Value = cast.ToInt32(BytesToUint32s(e, w, bytes)[0])
			return nil
		case strings.ToUpper(common.ValueTypeFloat32):
			if length != 2 {
				return fmt.Errorf("illegal length %s", prv.Name)
			}
			prv.Value = BytesToFloat32s(e, w, bytes)[0]
			return nil
		case strings.ToUpper(common.ValueTypeUint64):
			if length != 4 {
				return fmt.Errorf("illegal length %s", prv.Name)
			}
			prv.Value = BytesToUint64s(e, w, bytes)[0]
			return nil
		case strings.ToUpper(common.ValueTypeInt64):
			if length != 4 {
				return fmt.Errorf("illegal length %s", prv.Name)
			}
			prv.Value = cast.ToInt64(BytesToUint64s(e, w, bytes)[0])
			return nil
		case strings.ToUpper(common.ValueTypeFloat64):
			if length != 4 {
				return fmt.Errorf("illegal length %s", prv.Name)
			}
			prv.Value = BytesToFloat64s(e, w, bytes)[0]
			return nil
		case strings.ToUpper(common.ValueTypeString):
			prv.Value = cast.ToString(bytes)
			return nil
		default:
			// 未填获取填错的情况不做处理
			return nil
		}
	default:
		return fmt.Errorf("unknown register type: %s", rawType)
	}
}

// Release 释放资源
// 不释放连接资源，经测试该包不支持频繁创建连接
func (c *connector) Release() (err error) {
	defer func() {
		c.mutex.Unlock()
	}()
	c.tunnel.Close()
	return
}

// connect 建立并打开连接
func (c *connector) connect() {
	cfg := c.config
	tunnel, err := knx.NewGroupTunnel(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port), knx.TunnelConfig{
		ResendInterval:    500 * time.Millisecond,
		HeartbeatInterval: 10 * time.Second,
		ResponseTimeout:   cfg.ResponseTimeout,
		SendLocalAddress:  cfg.SendLocal,
		LocalAddress:      cfg.LocalAddress,
		UseTCP:            cfg.UseTcp,
	})
	if err != nil {
		c.plugin.logger.Error(fmt.Sprintf("knx connect:%v:%v error", cfg.Host, cfg.Port), zap.Error(err))
		return
	}
	c.tunnel = &tunnel
	//启动读监听
	go func() {
		for msg := range tunnel.Inbound() {
			c.plugin.callback(c.plugin, msg)
		}
	}()
}
