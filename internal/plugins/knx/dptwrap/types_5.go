package dptwrap

import (
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/knx/dpt"
	"github.com/spf13/cast"
)

// DPT_5001 represents DPT 5.001 / Scaling.
// type DPT_5001 float32
type DPT_5001_Wrap struct {
	src dpt.DPT_5001
}

func (d *DPT_5001_Wrap) Name() string {
	return "Percent(0..100)"
}

func (d *DPT_5001_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_5001_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_5001_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_5001_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_5001_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_5001(b)
}

func (d *DPT_5001_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_5003 represents DPT 5.003 / Angle.
// type DPT_5003 float32
type DPT_5003_Wrap struct {
	src dpt.DPT_5003
}

func (d *DPT_5003_Wrap) Name() string {
	return "Angle(0..360)"
}

func (d *DPT_5003_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_5003_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_5003_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_5003_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_5003_Wrap) Set(i interface{}) {
	b := cast.ToFloat32(i)
	d.src = dpt.DPT_5003(b)
}

func (d *DPT_5003_Wrap) Get() interface{} {
	return float32(d.src)
}

// DPT_5004 represents DPT 5.004 / Percent_U8.
// type DPT_5004 uint8
type DPT_5004_Wrap struct {
	src dpt.DPT_5004
}

func (d *DPT_5004_Wrap) Name() string {
	return "Percent(0..255)"
}

func (d *DPT_5004_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_5004_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_5004_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_5004_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_5004_Wrap) Set(i interface{}) {
	b := cast.ToUint8(i)
	d.src = dpt.DPT_5004(b)
}

func (d *DPT_5004_Wrap) Get() interface{} {
	return uint8(d.src)
}

// DPT_5005 represents DPT 5.005 / Ratio (0..255).
// type DPT_5005 uint8
type DPT_5005_Wrap struct {
	src dpt.DPT_5005
}

func (d *DPT_5005_Wrap) Name() string {
	return "Ratio(0..255)"
}

func (d *DPT_5005_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_5005_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_5005_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_5005_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_5005_Wrap) Set(i interface{}) {
	b := cast.ToUint8(i)
	d.src = dpt.DPT_5005(b)
}

func (d *DPT_5005_Wrap) Get() interface{} {
	return uint8(d.src)
}
