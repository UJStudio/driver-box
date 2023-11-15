package dptwrap

import (
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/knx/dpt"
	"github.com/spf13/cast"
)

// DPT_7001 represents DPT 7.001 / Value 2 Ucount.
// type DPT_7001 uint16
type DPT_7001_Wrap struct {
	src dpt.DPT_7001
}

func (d *DPT_7001_Wrap) Name() string {
	return "pulses"
}

func (d *DPT_7001_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_7001_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_7001_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_7001_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_7001_Wrap) Set(i interface{}) {
	b := cast.ToUint16(i)
	d.src = dpt.DPT_7001(b)
}

func (d *DPT_7001_Wrap) Get() interface{} {
	return uint16(d.src)
}

// DPT_7002 represents DPT 7.002 / Time Period MSec.
// type DPT_7002 uint16
type DPT_7002_Wrap struct {
	src dpt.DPT_7002
}

func (d *DPT_7002_Wrap) Name() string {
	return "Time(ms)"
}

func (d *DPT_7002_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_7002_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_7002_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_7002_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_7002_Wrap) Set(i interface{}) {
	b := cast.ToUint16(i)
	d.src = dpt.DPT_7002(b)
}

func (d *DPT_7002_Wrap) Get() interface{} {
	return uint16(d.src)
}

// DPT_7003 represents DPT 7.003 / Time Period 10 MSec.
// type DPT_7003 uint16
type DPT_7003_Wrap struct {
	src dpt.DPT_7003
}

func (d *DPT_7003_Wrap) Name() string {
	return "Time(10ms)"
}

func (d *DPT_7003_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_7003_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_7003_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_7003_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_7003_Wrap) Set(i interface{}) {
	b := cast.ToUint16(i)
	d.src = dpt.DPT_7003(b)
}

func (d *DPT_7003_Wrap) Get() interface{} {
	return uint16(d.src)
}

// DPT_7004 represents DPT 7.004 / Time Period 100 MSec.
// type DPT_7004 uint16
type DPT_7004_Wrap struct {
	src dpt.DPT_7004
}

func (d *DPT_7004_Wrap) Name() string {
	return "Time(100ms)"
}

func (d *DPT_7004_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_7004_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_7004_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_7004_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_7004_Wrap) Set(i interface{}) {
	b := cast.ToUint16(i)
	d.src = dpt.DPT_7004(b)
}

func (d *DPT_7004_Wrap) Get() interface{} {
	return uint16(d.src)
}

// DPT_7005 represents DPT 7.005 / Time Period Sec.
// type DPT_7005 uint16
type DPT_7005_Wrap struct {
	src dpt.DPT_7005
}

func (d *DPT_7005_Wrap) Name() string {
	return "Time(s)"
}

func (d *DPT_7005_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_7005_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_7005_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_7005_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_7005_Wrap) Set(i interface{}) {
	b := cast.ToUint16(i)
	d.src = dpt.DPT_7005(b)
}

func (d *DPT_7005_Wrap) Get() interface{} {
	return uint16(d.src)
}

// DPT_7006 represents DPT 7.006 / Time Period Min.
// type DPT_7006 uint16
type DPT_7006_Wrap struct {
	src dpt.DPT_7006
}

func (d *DPT_7006_Wrap) Name() string {
	return "Time(min)"
}

func (d *DPT_7006_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_7006_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_7006_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_7006_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_7006_Wrap) Set(i interface{}) {
	b := cast.ToUint16(i)
	d.src = dpt.DPT_7006(b)
}

func (d *DPT_7006_Wrap) Get() interface{} {
	return uint16(d.src)
}

// DPT_7007 represents DPT 7.007 / Time Period Hrs.
// type DPT_7007 uint16
type DPT_7007_Wrap struct {
	src dpt.DPT_7007
}

func (d *DPT_7007_Wrap) Name() string {
	return "Time(hr)"
}

func (d *DPT_7007_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_7007_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_7007_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_7007_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_7007_Wrap) Set(i interface{}) {
	b := cast.ToUint16(i)
	d.src = dpt.DPT_7007(b)
}

func (d *DPT_7007_Wrap) Get() interface{} {
	return uint16(d.src)
}

// DPT_7010 represents DPT 7.010 / Property DataType.
// type DPT_7010 uint16
type DPT_7010_Wrap struct {
	src dpt.DPT_7010
}

func (d *DPT_7010_Wrap) Name() string {
	return "Property DataType"
}

func (d *DPT_7010_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_7010_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_7010_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_7010_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_7010_Wrap) Set(i interface{}) {
	b := cast.ToUint16(i)
	d.src = dpt.DPT_7010(b)
}

func (d *DPT_7010_Wrap) Get() interface{} {
	return uint16(d.src)
}

// DPT_7011 represents DPT 7.011 / Length mm.
// type DPT_7011 uint16
type DPT_7011_Wrap struct {
	src dpt.DPT_7011
}

func (d *DPT_7011_Wrap) Name() string {
	return "Length(mm)"
}

func (d *DPT_7011_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_7011_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_7011_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_7011_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_7011_Wrap) Set(i interface{}) {
	b := cast.ToUint16(i)
	d.src = dpt.DPT_7011(b)
}

func (d *DPT_7011_Wrap) Get() interface{} {
	return uint16(d.src)
}

// DPT_7012 represents DPT 7.012 / Current mA.
// type DPT_7012 uint16
type DPT_7012_Wrap struct {
	src dpt.DPT_7012
}

func (d *DPT_7012_Wrap) Name() string {
	return "Current(mA)"
}

func (d *DPT_7012_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_7012_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_7012_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_7012_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_7012_Wrap) Set(i interface{}) {
	b := cast.ToUint16(i)
	d.src = dpt.DPT_7012(b)
}

func (d *DPT_7012_Wrap) Get() interface{} {
	return uint16(d.src)
}

// DPT_7013 represents DPT 7.013 / Brightness lux.
// type DPT_7013 uint16
type DPT_7013_Wrap struct {
	src dpt.DPT_7013
}

func (d *DPT_7013_Wrap) Name() string {
	return "Brightness(lux)"
}

func (d *DPT_7013_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_7013_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_7013_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_7013_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_7013_Wrap) Set(i interface{}) {
	b := cast.ToUint16(i)
	d.src = dpt.DPT_7013(b)
}

func (d *DPT_7013_Wrap) Get() interface{} {
	return uint16(d.src)
}
