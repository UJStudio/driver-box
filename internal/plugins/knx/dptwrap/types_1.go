package dptwrap

import (
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/knx/dpt"
	"github.com/spf13/cast"
)

// DPT_1001 represents DPT 1.001 (G) / DPT_Switch.
type DPT_1001_Wrap struct {
	src dpt.DPT_1001
}

func (d *DPT_1001_Wrap) Name() string {
	return "Switch"
}

func (d *DPT_1001_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1001_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1001_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1001_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1001_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1001(b)
}

func (d *DPT_1001_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1002 represents DPT 1.002 (G) / DPT_Bool.
type DPT_1002_Wrap struct {
	src dpt.DPT_1002
}

func (d *DPT_1002_Wrap) Name() string {
	return "Bool"
}

func (d *DPT_1002_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1002_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1002_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1002_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1002_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1002(b)
}

func (d *DPT_1002_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1003 represents DPT 1.003 (G) / DPT_Enable.
type DPT_1003_Wrap struct {
	src dpt.DPT_1003
}

func (d *DPT_1003_Wrap) Name() string {
	return "Enable"
}

func (d *DPT_1003_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1003_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1003_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1003_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1003_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1003(b)
}

func (d *DPT_1003_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1004 represents DPT 1.004 (FB) / DPT_Ramp.
// type DPT_1004 bool
type DPT_1004_Wrap struct {
	src dpt.DPT_1004
}

func (d *DPT_1004_Wrap) Name() string {
	return "Ramp"
}

func (d *DPT_1004_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1004_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1004_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1004_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1004_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1004(b)
}

func (d *DPT_1004_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1005 represents DPT 1.005 (FB) / DPT_Alarm.
// type DPT_1005 bool
type DPT_1005_Wrap struct {
	src dpt.DPT_1005
}

func (d *DPT_1005_Wrap) Name() string {
	return "Alarm"
}

func (d *DPT_1005_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1005_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1005_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1005_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1005_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1005(b)
}

func (d *DPT_1005_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1006 represents DPT 1.006 (FB) / DPT_BinaryValue.
// type DPT_1006 bool
type DPT_1006_Wrap struct {
	src dpt.DPT_1006
}

func (d *DPT_1006_Wrap) Name() string {
	return "BinaryValue"
}

func (d *DPT_1006_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1006_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1006_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1006_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1006_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1006(b)
}

func (d *DPT_1006_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1007 represents DPT 1.007 (FB) / DPT_Step.
// type DPT_1007 bool
type DPT_1007_Wrap struct {
	src dpt.DPT_1007
}

func (d *DPT_1007_Wrap) Name() string {
	return "Step"
}

func (d *DPT_1007_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1007_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1007_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1007_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1007_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1007(b)
}

func (d *DPT_1007_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1008 represents DPT 1.008 (G) / DPT_UpDown.
// type DPT_1008 bool
type DPT_1008_Wrap struct {
	src dpt.DPT_1008
}

func (d *DPT_1008_Wrap) Name() string {
	return "UpDown"
}

func (d *DPT_1008_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1008_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1008_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1008_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1008_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1008(b)
}

func (d *DPT_1008_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1009 represents DPT 1.009 (G) / DPT_OpenClose.
// type DPT_1009 bool
type DPT_1009_Wrap struct {
	src dpt.DPT_1009
}

func (d *DPT_1009_Wrap) Name() string {
	return "OpenClose"
}

func (d *DPT_1009_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1009_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1009_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1009_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1009_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1009(b)
}

func (d *DPT_1009_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1010 represents DPT 1.010 (G) / DPT_Start.
// type DPT_1010 bool
type DPT_1010_Wrap struct {
	src dpt.DPT_1010
}

func (d *DPT_1010_Wrap) Name() string {
	return "Start"
}

func (d *DPT_1010_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1010_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1010_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1010_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1010_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1010(b)
}

func (d *DPT_1010_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1011 represents DPT 1.011 (FB) / DPT_State.
// type DPT_1011 bool
type DPT_1011_Wrap struct {
	src dpt.DPT_1011
}

func (d *DPT_1011_Wrap) Name() string {
	return "State"
}

func (d *DPT_1011_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1011_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1011_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1011_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1011_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1011(b)
}

func (d *DPT_1011_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1012 represents DPT 1.012 (FB) / DPT_Invert.
// type DPT_1012 bool
type DPT_1012_Wrap struct {
	src dpt.DPT_1012
}

func (d *DPT_1012_Wrap) Name() string {
	return "Invert"
}

func (d *DPT_1012_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1012_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1012_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1012_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1012_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1012(b)
}

func (d *DPT_1012_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1013 represents DPT 1.013 (FB) / DPT_DimSendStyle.
// type DPT_1013 bool
type DPT_1013_Wrap struct {
	src dpt.DPT_1013
}

func (d *DPT_1013_Wrap) Name() string {
	return "DimSendStyle"
}

func (d *DPT_1013_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1013_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1013_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1013_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1013_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1013(b)
}

func (d *DPT_1013_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1014 represents DPT 1.014 (FB) / DPT_InputSource.
// type DPT_1014 bool
type DPT_1014_Wrap struct {
	src dpt.DPT_1014
}

func (d *DPT_1014_Wrap) Name() string {
	return "InputSource"
}

func (d *DPT_1014_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1014_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1014_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1014_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1014_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1014(b)
}

func (d *DPT_1014_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1015 represents DPT 1.015 (G) / DPT_Reset.
// type DPT_1015 bool
type DPT_1015_Wrap struct {
	src dpt.DPT_1015
}

func (d *DPT_1015_Wrap) Name() string {
	return "Reset"
}

func (d *DPT_1015_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1015_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1015_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1015_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1015_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1015(b)
}

func (d *DPT_1015_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1016 represents DPT 1.016 (G) / DPT_Ack.
// type DPT_1016 bool
type DPT_1016_Wrap struct {
	src dpt.DPT_1016
}

func (d *DPT_1016_Wrap) Name() string {
	return "Ack"
}

func (d *DPT_1016_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1016_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1016_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1016_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1016_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1016(b)
}

func (d *DPT_1016_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1017 represents DPT 1.017 (G) / DPT_Trigger.
// type DPT_1017 bool
type DPT_1017_Wrap struct {
	src dpt.DPT_1017
}

func (d *DPT_1017_Wrap) Name() string {
	return "Trigger"
}

func (d *DPT_1017_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1017_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1017_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1017_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1017_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1017(b)
}

func (d *DPT_1017_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1018 represents DPT 1.018 (G) / DPT_Occupancy.
// type DPT_1018 bool
type DPT_1018_Wrap struct {
	src dpt.DPT_1018
}

func (d *DPT_1018_Wrap) Name() string {
	return "Occupancy"
}

func (d *DPT_1018_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1018_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1018_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1018_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1018_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1018(b)
}

func (d *DPT_1018_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1019 represents DPT 1.019 (G) / DPT_Window_Door.
// type DPT_1019 bool
type DPT_1019_Wrap struct {
	src dpt.DPT_1019
}

func (d *DPT_1019_Wrap) Name() string {
	return "Window_Door"
}

func (d *DPT_1019_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1019_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1019_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1019_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1019_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1019(b)
}

func (d *DPT_1019_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1021 represents DPT 1.021 (FB) / DPT_LogicalFunction.
// type DPT_1021 bool
type DPT_1021_Wrap struct {
	src dpt.DPT_1021
}

func (d *DPT_1021_Wrap) Name() string {
	return "LogicalFunction"
}

func (d *DPT_1021_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1021_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1021_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1021_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1021_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1021(b)
}

func (d *DPT_1021_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1022 represents DPT 1.022 (FB) / DPT_Scene_AB.
// type DPT_1022 bool
type DPT_1022_Wrap struct {
	src dpt.DPT_1022
}

func (d *DPT_1022_Wrap) Name() string {
	return "Scene_AB"
}

func (d *DPT_1022_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1022_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1022_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1022_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1022_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1022(b)
}

func (d *DPT_1022_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1023 represents DPT 1.023 (FB) / DPT_ShutterBlinds_Mode.
// type DPT_1023 bool
type DPT_1023_Wrap struct {
	src dpt.DPT_1023
}

func (d *DPT_1023_Wrap) Name() string {
	return "ShutterBlinds_Mode"
}

func (d *DPT_1023_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1023_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1023_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1023_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1023_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1023(b)
}

func (d *DPT_1023_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1024 represents DPT 1.024 (G) / DPT_DayNight.
// type DPT_1024 bool
type DPT_1024_Wrap struct {
	src dpt.DPT_1024
}

func (d *DPT_1024_Wrap) Name() string {
	return "DayNight"
}

func (d *DPT_1024_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1024_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1024_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1024_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1024_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1024(b)
}

func (d *DPT_1024_Wrap) Get() interface{} {
	return bool(d.src)
}

// DPT_1100 represents DPT 1.100 (FB) / DPT_Heat/Cool.
// type DPT_1100 bool
type DPT_1100_Wrap struct {
	src dpt.DPT_1100
}

func (d *DPT_1100_Wrap) Name() string {
	return "Heat/Cool"
}

func (d *DPT_1100_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_1100_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_1100_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_1100_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_1100_Wrap) Set(i interface{}) {
	b := cast.ToBool(i)
	d.src = dpt.DPT_1100(b)
}

func (d *DPT_1100_Wrap) Get() interface{} {
	return bool(d.src)
}
