package dptwrap

import (
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/knx/dpt"
	"github.com/spf13/cast"
)

// DPT_6010 represents DPT 6.010 / counter pulses (-128..127).
// type DPT_6010 int8
type DPT_6010_Wrap struct {
	src dpt.DPT_6010
}

func (d *DPT_6010_Wrap) Name() string {
	return "counter pulses(-128..127)"
}

func (d *DPT_6010_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_6010_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_6010_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_6010_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_6010_Wrap) Set(i interface{}) {
	b := cast.ToInt8(i)
	d.src = dpt.DPT_6010(b)
}

func (d *DPT_6010_Wrap) Get() interface{} {
	return int8(d.src)
}
