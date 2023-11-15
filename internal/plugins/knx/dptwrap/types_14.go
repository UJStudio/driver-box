package dptwrap

import (
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/knx/dpt"
	"github.com/spf13/cast"
)

// DPT_14033 represents DPT 14.033 / Frequency
// type DPT_14033 float32
type DPT_14033_Wrap struct {
	src dpt.DPT_14033
}

func (d *DPT_14033_Wrap) Name() string {
	return "Frequency(Hz)"
}

func (d *DPT_14033_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_14033_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_14033_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_14033_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_14033_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_14033(b)
}

func (d *DPT_14033_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_14056 represents DPT 14.056 / Power
// type DPT_14056 float32
type DPT_14056_Wrap struct {
	src dpt.DPT_14056
}

func (d *DPT_14056_Wrap) Name() string {
	return "Power(W)"
}

func (d *DPT_14056_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_14056_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_14056_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_14056_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_14056_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_14056(b)
}

func (d *DPT_14056_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_14058 represents DPT 14.058 / Pressure
// type DPT_14058 float32
type DPT_14058_Wrap struct {
	src dpt.DPT_14058
}

func (d *DPT_14058_Wrap) Name() string {
	return "Pressure(Pa)"
}

func (d *DPT_14058_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_14058_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_14058_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_14058_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_14058_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_14058(b)
}

func (d *DPT_14058_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_14059 represents DPT 14.059 / Reactance
// type DPT_14059 float32
type DPT_14059_Wrap struct {
	src dpt.DPT_14059
}

func (d *DPT_14059_Wrap) Name() string {
	return "Reactance(Ω)"
}

func (d *DPT_14059_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_14059_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_14059_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_14059_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_14059_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_14059(b)
}

func (d *DPT_14059_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_14060 represents DPT 14.060 / Resistance
// type DPT_14060 float32
type DPT_14060_Wrap struct {
	src dpt.DPT_14060
}

func (d *DPT_14060_Wrap) Name() string {
	return "Resistance(Ω)"
}

func (d *DPT_14060_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_14060_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_14060_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_14060_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_14060_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_14060(b)
}

func (d *DPT_14060_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_14065 represents DPT 14.065 / Speed
// type DPT_14065 float32
type DPT_14065_Wrap struct {
	src dpt.DPT_14065
}

func (d *DPT_14065_Wrap) Name() string {
	return "Speed(m/s)"
}

func (d *DPT_14065_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_14065_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_14065_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_14065_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_14065_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_14065(b)
}

func (d *DPT_14065_Wrap) Get() interface{} {
	return float32(d.src)
}
