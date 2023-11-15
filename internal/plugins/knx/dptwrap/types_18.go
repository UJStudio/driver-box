package dptwrap

import (
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/knx/dpt"
	"github.com/spf13/cast"
)

// DPT_18001 represents DPT 18.001 / Scene Control.
// type DPT_18001 uint8
type DPT_18001_Wrap struct {
	src dpt.DPT_18001
}

func (d *DPT_18001_Wrap) Name() string {
	return "Scene Control"
}

func (d *DPT_18001_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_18001_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_18001_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_18001_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_18001_Wrap) Set(i interface{}) {
	b := cast.ToUint8(i)
	d.src = dpt.DPT_18001(b)
}

func (d *DPT_18001_Wrap) Get() interface{} {
	return uint8(d.src)
}
