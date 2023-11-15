package dptwrap

import (
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/knx/dpt"
	"github.com/spf13/cast"
)

// DPT_9001 represents DPT 9.001 / Temperature °C.
// type DPT_9001 float32
type DPT_9001_Wrap struct {
	src dpt.DPT_9001
}

func (d *DPT_9001_Wrap) Name() string {
	return "Temperature(°C)"
}

func (d *DPT_9001_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9001_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9001_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9001_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9001_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9001(b)
}

func (d *DPT_9001_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9002 represents DPT 9.002 / Temperature K.
// type DPT_9002 float32
type DPT_9002_Wrap struct {
	src dpt.DPT_9002
}

func (d *DPT_9002_Wrap) Name() string {
	return "Temperature(K)"
}

func (d *DPT_9002_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9002_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9002_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9002_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9002_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9002(b)
}

func (d *DPT_9002_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9003 represents DPT 9.003 / Temperature K/h.
// type DPT_9003 float32
type DPT_9003_Wrap struct {
	src dpt.DPT_9003
}

func (d *DPT_9003_Wrap) Name() string {
	return "Temperature(K/h)"
}

func (d *DPT_9003_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9003_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9003_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9003_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9003_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9003(b)
}

func (d *DPT_9003_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9004 represents DPT 9.004 / Illumination lux.
// type DPT_9004 float32
type DPT_9004_Wrap struct {
	src dpt.DPT_9004
}

func (d *DPT_9004_Wrap) Name() string {
	return "Illumination(lux)"
}

func (d *DPT_9004_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9004_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9004_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9004_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9004_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9004(b)
}

func (d *DPT_9004_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9005 represents DPT 9.005 / Wind Speed m/s.
// type DPT_9005 float32
type DPT_9005_Wrap struct {
	src dpt.DPT_9005
}

func (d *DPT_9005_Wrap) Name() string {
	return "Wind Speed(m/s)"
}

func (d *DPT_9005_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9005_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9005_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9005_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9005_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9005(b)
}

func (d *DPT_9005_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9006 represents DPT 9.006 / Pressure Pa.
// type DPT_9006 float32
type DPT_9006_Wrap struct {
	src dpt.DPT_9006
}

func (d *DPT_9006_Wrap) Name() string {
	return "Pressure(Pa)"
}

func (d *DPT_9006_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9006_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9006_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9006_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9006_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9006(b)
}

func (d *DPT_9006_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9007 represents DPT 9.007 / Humidity %
// type DPT_9007 float32
type DPT_9007_Wrap struct {
	src dpt.DPT_9007
}

func (d *DPT_9007_Wrap) Name() string {
	return "Humidity(%)"
}

func (d *DPT_9007_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9007_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9007_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9007_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9007_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9007(b)
}

func (d *DPT_9007_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9008 represents DPT 9.008 / Air quality ppm
// type DPT_9008 float32
type DPT_9008_Wrap struct {
	src dpt.DPT_9008
}

func (d *DPT_9008_Wrap) Name() string {
	return "Air quality(ppm)"
}

func (d *DPT_9008_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9008_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9008_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9008_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9008_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9008(b)
}

func (d *DPT_9008_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9010 represents DPT 9.010 / Time s.
// type DPT_9010 float32
type DPT_9010_Wrap struct {
	src dpt.DPT_9010
}

func (d *DPT_9010_Wrap) Name() string {
	return "Time(s)"
}

func (d *DPT_9010_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9010_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9010_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9010_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9010_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9010(b)
}

func (d *DPT_9010_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9011 represents DPT 9.011 / Time ms.
// type DPT_9011 float32
type DPT_9011_Wrap struct {
	src dpt.DPT_9011
}

func (d *DPT_9011_Wrap) Name() string {
	return "Time(ms)"
}

func (d *DPT_9011_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9011_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9011_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9011_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9011_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9011(b)
}

func (d *DPT_9011_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9020 represents DPT 9.020 / Volt mV.
// type DPT_9020 float32
type DPT_9020_Wrap struct {
	src dpt.DPT_9020
}

func (d *DPT_9020_Wrap) Name() string {
	return "Volt(mV)"
}

func (d *DPT_9020_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9020_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9020_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9020_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9020_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9020(b)
}

func (d *DPT_9020_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9021 represents DPT 9.021 / Current mA.
// type DPT_9021 float32
type DPT_9021_Wrap struct {
	src dpt.DPT_9021
}

func (d *DPT_9021_Wrap) Name() string {
	return "Current(mA)"
}

func (d *DPT_9021_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9021_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9021_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9021_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9021_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9021(b)
}

func (d *DPT_9021_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9022 represents DPT 9.022 / Power Density W/m2.
// type DPT_9022 float32
type DPT_9022_Wrap struct {
	src dpt.DPT_9022
}

func (d *DPT_9022_Wrap) Name() string {
	return "Power Density(W/m2)"
}

func (d *DPT_9022_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9022_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9022_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9022_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9022_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9022(b)
}

func (d *DPT_9022_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9024 represents DPT 9.024 / Power kW.
// type DPT_9024 float32
type DPT_9024_Wrap struct {
	src dpt.DPT_9024
}

func (d *DPT_9024_Wrap) Name() string {
	return "Power(kW)"
}

func (d *DPT_9024_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9024_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9024_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9024_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9024_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9024(b)
}

func (d *DPT_9024_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9025 represents DPT 9.025 / Volume Flow l/h.
// type DPT_9025 float32
type DPT_9025_Wrap struct {
	src dpt.DPT_9025
}

func (d *DPT_9025_Wrap) Name() string {
	return "Volume Flow(l/h)"
}

func (d *DPT_9025_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9025_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9025_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9025_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9025_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9025(b)
}

func (d *DPT_9025_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_9028 represents DPT 9.028 / Wind Speed km/h.
// type DPT_9028 float32
type DPT_9028_Wrap struct {
	src dpt.DPT_9028
}

func (d *DPT_9028_Wrap) Name() string {
	return "Wind Speed(km/h)"
}

func (d *DPT_9028_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_9028_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_9028_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_9028_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_9028_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_9028(b)
}

func (d *DPT_9028_Wrap) Get() interface{} {
	return float32(d.src)
}
